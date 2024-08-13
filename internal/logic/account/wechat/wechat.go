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

package wechat

import (
	"context"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/ArtisanCloud/PowerWeChat/v3/test/testLogDriver"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
	"os"
)

var OfficialAccountApp *officialAccount.OfficialAccount
var MiniProgramApp *miniProgram.MiniProgram

const TIMEZONE = "asia/shanghai"
const DATETIME_FORMAT = "20060102"

type sWechat struct{}

func init() {
	service.RegisterWechat(New())
}

func New() *sWechat {
	s := &sWechat{}

	return s
}

func (s *sWechat) MiniMiniProgram(ctx context.Context) (*miniProgram.MiniProgram, error) {
	if MiniProgramApp == nil {
		var cache kernel.CacheInterface
		redisAddrs, _ := g.Cfg().Get(ctx, "redis.default.address")
		redisDb, _ := g.Cfg().Get(ctx, "redis.default.db")
		redisPass, _ := g.Cfg().Get(ctx, "redis.default.pass")

		if redisAddrs.String() != "" {
			cache = kernel.NewRedisClient(&kernel.UniversalOptions{
				Addrs: []string{redisAddrs.String()},
				//Addrs: []string{
				//	"47.108.182.200:7000",
				//	"47.108.182.200:7001",
				//	"47.108.182.200:7002",
				//},
				DB:       redisDb.Int(),
				Password: redisPass.String(),
			})
		}

		WechatXcxAppId := service.ConfigBase().GetStr(ctx, "wechat_xcx_app_id", "")
		WechatXcxAppSecret := service.ConfigBase().GetStr(ctx, "wechat_xcx_app_secret", "")

		app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
			AppID:        WechatXcxAppId,     // 小程序、公众号或者企业微信的appid
			Secret:       WechatXcxAppSecret, // 商户号 appID
			ResponseType: response.TYPE_MAP,
			//Token:        conf.MiniProgram.MessageToken,
			//AESKey:       conf.MiniProgram.MessageAesKey,
			//
			//AppKey:  conf.MiniProgram.VirtualPayAppKey,
			//OfferID: conf.MiniProgram.VirtualPayOfferID,
			Http: miniProgram.Http{},
			Log: miniProgram.Log{
				Driver: &testLogDriver.SimpleLogger{},
				Level:  "debug",
				File:   "./wechat.log",
			},
			//"sandbox": true,
			Cache:     cache,
			HttpDebug: true,
			Debug:     false,
		})

		MiniProgramApp = app
		return MiniProgramApp, err
	}

	return MiniProgramApp, nil
}

func (s *sWechat) OfficialAccountApp(ctx context.Context) (*officialAccount.OfficialAccount, error) {
	if OfficialAccountApp == nil {
		var cache kernel.CacheInterface

		redisAddrs, _ := g.Cfg().Get(ctx, "redis.default.address")
		redisDb, _ := g.Cfg().Get(ctx, "redis.default.db")
		redisPass, _ := g.Cfg().Get(ctx, "redis.default.pass")

		if redisAddrs.String() != "" {
			cache = kernel.NewRedisClient(&kernel.UniversalOptions{
				Addrs:    []string{redisAddrs.String()},
				DB:       redisDb.Int(),
				Password: redisPass.String(),
			})
		}

		WechatAppId := service.ConfigBase().GetStr(ctx, "wechat_app_id", "")
		WechatAppSecret := service.ConfigBase().GetStr(ctx, "wechat_app_secret", "")

		app, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{

			AppID:  WechatAppId,     // 小程序、公众号或者企业微信的appid
			Secret: WechatAppSecret, // 商户号 appID

			//Token:        conf.OffiAccount.MessageToken,
			//AESKey:       conf.OffiAccount.MessageAesKey,
			ResponseType: os.Getenv("response_type"),
			Log: officialAccount.Log{
				Driver: &testLogDriver.SimpleLogger{},
				Level:  "error",
				File:   "./wechat.log",
			},
			Cache:     cache,
			HttpDebug: true,
			Debug:     true,
			//"sandbox": true,
		})

		OfficialAccountApp = app
		return OfficialAccountApp, err
	}

	return OfficialAccountApp, nil
}

func (s *sWechat) CallbackMp(ctx context.Context, code string, activityId uint) (out model.LoginOutput, err error) {
	if g.IsEmpty(code) {
		return out, errors.New("传入的code值为空")
	}

	app, _ := s.OfficialAccountApp(ctx)
	app.OAuth.SetScopes([]string{"snsapi_userinfo"})
	data, err := app.OAuth.UserFromCode(code)

	if err != nil || data == nil {
		return out, errors.New("数据异常")
	}

	fmt.Printf("%+v\n", data)

	vxUserBindConnect := &do.UserBindConnect{}
	//sessionKey := data.SessionKey
	vxUserBindConnect.BindId = data.GetOpenID()
	vxUserBindConnect.BindOpenid = data.GetOpenID()

	//fix unionid
	if data.Attributes["token_response"] != nil {
		tokenResponse := gjson.New(data.Attributes["token_response"].(string))

		if tokenResponse.Get("unionid") != nil {
			vxUserBindConnect.BindUnionid = tokenResponse.Get("unionid").String()
		}
	}

	vxUserBindConnect.BindNickname = data.GetNickname()
	vxUserBindConnect.BindIcon = data.GetAvatar()
	vxUserBindConnect.BindType = consts.WEIXIN

	userId, err := service.User().DoUserBind(ctx, vxUserBindConnect, activityId, true)
	if err != nil {
		return out, err
	}
	return service.Login().LoginById(ctx, userId)
}

// CheckAppLogin checks the app login status based on the provided code.
func (s *sWechat) CheckAppLogin(ctx context.Context, code string) (out model.LoginOutput, err error) {
	if g.IsEmpty(code) {
		return out, errors.New("传入的code值为空")
	}

	app, _ := s.MiniMiniProgram(ctx)
	data, err := app.Auth.Session(ctx, code)

	if err != nil || data == nil {
		return out, errors.New("数据异常")
	}

	userBindConnect, err := dao.UserBindConnect.Get(ctx, data.OpenID)
	if err != nil {
		return out, err
	}

	// 判断是否存在bind_unionid
	if userBindConnect == nil && !g.IsEmpty(data.UnionID) {
		userBindConnect, err = dao.UserBindConnect.FindOne(ctx, &do.UserBindConnectListInput{
			Where: do.UserBindConnect{BindUnionid: data.UnionID},
		})
		if err != nil {
			return out, err
		}
	}

	var userId uint
	if userBindConnect != nil {
		userId = userBindConnect.UserId
		return service.Login().LoginById(ctx, userId)
	} else {
		return out, errors.New("用户不存在")
	}

	return out, nil
}

// JsCode2SessionReq
func (s *sWechat) JsCode2Session(ctx context.Context, code, encryptedData, iv, userInfo string, activityId, sourceUserId uint) (out model.LoginOutput, err error) {
	vxUserBindConnect, err := s.GetVxMiniAppUserBindConnect(ctx, code, encryptedData, iv, userInfo)
	if err != nil {
		return out, err
	}

	userId, err := service.User().DoUserBind(ctx, vxUserBindConnect, activityId, true)
	if err != nil {
		return out, err
	}

	return service.Login().LoginById(ctx, userId)
}

// GetOpenIdByCode
func (s *sWechat) GetOpenIdByCode(ctx context.Context, code string, userId uint) (out map[string]interface{}, err error) {
	out = make(map[string]interface{})

	if !g.IsEmpty(userId) {
		userBindConnect, err := service.User().GetBind(ctx, userId, consts.WEIXIN_XCX)
		if err != nil {
			return out, err
		}

		if userBindConnect != nil && !g.IsEmpty(userBindConnect.BindOpenid) {
			out["openid"] = userBindConnect.BindOpenid
		}
	}

	//不存在openid , 接口获取
	// 使用索引操作符
	if _, exists := out["openid"]; !exists {
		if g.IsEmpty(code) {
			return out, errors.New("传入的code值为空")
		}

		app, _ := s.MiniMiniProgram(ctx)
		data, err := app.Auth.Session(ctx, code)

		if err != nil || data == nil {
			return out, errors.New("数据异常")
		}

		out["openid"] = data.OpenID
	}

	return out, err
}

// 获取微信小程序用户绑定信息
func (s *sWechat) GetVxMiniAppUserBindConnect(ctx context.Context, code, encryptedData, iv, userInfo string) (*do.UserBindConnect, error) {
	// 解析用户基本信息

	userInfoJson := gjson.New(userInfo)
	userInfoJson.Map()
	nickName := userInfoJson.Get("nickName").String()
	//gender := userInfoJson.Get("gender").Int()
	//language := userInfoJson.Get("language").String()
	//city := userInfoJson.Get("city").String()
	//province := userInfoJson.Get("province").String()
	//country := userInfoJson.Get("country").String()
	avatarUrl := userInfoJson.Get("avatarUrl").String()

	// 初始化 UserBindConnect 对象
	accountUserBindConnect := &do.UserBindConnect{}

	// 调用 getJsCode2Session 方法获取 sessionKey, openid 等信息
	app, _ := s.MiniMiniProgram(ctx)
	data, err := app.Auth.Session(ctx, code)

	if err != nil {
		return nil, err
	}

	//sessionKey := data.SessionKey
	accountUserBindConnect.BindId = data.OpenID
	accountUserBindConnect.BindOpenid = data.OpenID
	accountUserBindConnect.BindUnionid = data.UnionID
	accountUserBindConnect.BindNickname = nickName
	accountUserBindConnect.BindIcon = avatarUrl
	accountUserBindConnect.BindType = consts.WEIXIN_XCX

	return accountUserBindConnect, nil
}
