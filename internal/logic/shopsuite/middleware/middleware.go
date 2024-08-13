// +----------------------------------------------------------------------
// | ShopSuite商城系统 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 随商信息技术（上海）有限公司
// +----------------------------------------------------------------------
// | 未获商业授权前，不得将本软件用于商业用途。禁止整体或任何部分基础上以发展任何派生版本、
// | 修改版本或第三方版本用于重新分发。
// +----------------------------------------------------------------------
// | 官方网站: https://www.shopsuite.cn  https://www.golershop.cn
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本公司对该软件产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细见https://www.golershop.cn/policy
// +----------------------------------------------------------------------

package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/text/gstr"
	"golershop.cn/internal/model"
	"golershop.cn/internal/service"
	"golershop.cn/utility"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

// I18N I18N
func (s *sMiddleware) I18N(r *ghttp.Request) {
	//r.SetCtx(gi18n.WithLanguage(r.Context(), r.GetString("source_lang", "zh-CN")))

	r.Middleware.Next()
}

// Ctx injects custom business context variable into context of current request.
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	service.BizCtx().Init(r, customCtx)

	//gwt 信息解码存入，方便直接读取使用
	// 从请求头中获取Token
	tokenHeader, _ := gcfg.Instance().Get(r.GetCtx(), "jwt.tokenHeader")
	token := r.GetHeader(tokenHeader.String())

	// 字符串替换
	tokenPrefix, _ := gcfg.Instance().Get(r.GetCtx(), "jwt.tokenPrefix")
	token = gstr.Replace(token, tokenPrefix.String(), "")
	token = gstr.Trim(token)

	if token == "" {
		token = r.Get("perm_key").String()
	}

	if token != "" {
		claim, err := utility.ParseToken(r.GetCtx(), token)
		if err != nil {
			fmt.Println("解析token出现错误：", err)
			//errors.New("Token已过期")
			//r.SetError(gerror.NewCode(gcode.New(101, "", nil), "Token已过期"))
			//r.Response.WriteHeader(http.StatusUnauthorized)

			//r.SetError(gerror.NewCode(gcode.New(http.StatusUnauthorized, "Token已过期", nil)))
		} else if time.Now().Unix() > claim.ExpiresAt {
			fmt.Println("时间超时")

			//r.SetError(gerror.NewCode(gcode.New(http.StatusUnauthorized, "Token时间超时", nil)))
		} else {
			//记录登录用户信息
			//gwt 信息解码存入，方便直接读取使用

			// todo claim.UserSalt读取用户信息，查看UserSalt是否变更，变更则失效
			clientId := uint(0)
			if strings.HasPrefix(r.RequestURI, "/manage") {
				clientId = 1
			} else {
				if strings.HasPrefix(r.RequestURI, "/front/account/user/info") {
					clientId = 1
				}
			}

			// todo 读取其它信息，存入
			service.BizCtx().SetUser(r.Context(), &model.ContextUser{
				UserId:      claim.UserId,
				UserAccount: claim.UserAccount,
				UserSalt:    claim.UserSalt,
				ClientId:    clientId,
			})
		}
	}

	//session 信息存入
	/*
		if user := service.Session().GetUser(r.Context()); user != nil {
			customCtx.User = &model.ContextUser{
				UserId:       user.UserId,
				Passport: user.UserToken,
				Nickname: user.UserAccount,
			}
		}
	*/

	// Continue execution of next middleware.
	r.Middleware.Next()
}

// Auth validates the request to allow only signed-in users visit.
func (s *sMiddleware) Auth(r *ghttp.Request) {
	if service.Login().IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// NeverDoneCtx
func (s *sMiddleware) NeverDoneCtx(r *ghttp.Request) {
	r.SetCtx(r.GetNeverDoneCtx())
	r.Middleware.Next()
}

// 登录验证中间件
func (s *sMiddleware) CheckLogin(r *ghttp.Request) {
	//fmt.Println("登录验证中间件")
	//fmt.Println(r.RequestURI)
	//fmt.Println(r.URL)

	u, _ := url.Parse(r.RequestURI) //将string解析成*URL格式

	//fmt.Println(u.Path)

	// 放行设置 通过配置
	urlItem := []string{}
	si, _ := gcfg.Instance().Get(r.GetCtx(), "secure.ignore")

	for _, url := range si.Interfaces() {
		urlItem = append(urlItem, url.(string))
	}

	if !gstr.InArray(urlItem, u.Path) {
		if service.Login().IsSignedIn(r.Context()) {
			r.Middleware.Next()
		} else {
			//r.SetError(gerror.NewCode(gcode.New(http.StatusUnauthorized, "需要登录", nil)))

			//返回固定的友好信息
			r.Response.ClearBuffer()
			r.Response.WriteJsonExit(model.DefaultHandlerResponse{
				Code:    http.StatusUnauthorized,
				Message: "需要登录",
				Data:    nil,
				Status:  250,
			})
		}
	} else {
		// 前置中间件
		r.Middleware.Next()
	}
}

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func (s *sMiddleware) MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg    string
		err    = r.GetError()
		res    = r.GetHandlerResponse()
		code   = gerror.Code(err)
		status = 250
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusNotFound:
			code = gcode.CodeNotFound
		case http.StatusForbidden:
			code = gcode.CodeNotAuthorized
		default:
			code = gcode.CodeUnknown
		}
	} else {
		code = gcode.CodeOK

		//处理用户层级提示,都正确
		msg = "操作成功"
		status = 200
	}

	r.Response.WriteJson(model.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
		Status:  status,
	})
}

func (s *sMiddleware) MiddlewareErrorHandler(r *ghttp.Request) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		r.Response.WriteJson(g.Map{
	//			"code":    500,
	//			"message": "Internal Server Error",
	//			"data":    nil,
	//			"error":   fmt.Sprintf("%v", err),
	//		})
	//	}
	//}()

	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		//g.Log("exception").Error(r.GetCtx(), err)

		//返回固定的友好信息
		r.Response.ClearBuffer()
		//r.Response.Writeln("服务器居然开小差了，请稍后再试吧！")
		r.Response.WriteJsonExit(model.DefaultHandlerResponse{
			Code:    gerror.Code(err).Code(),
			Message: gstr.Replace(errorToString(err), "exception recovered: ", ""),
			Data:    nil,
			Status:  250,
		})
	}
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
