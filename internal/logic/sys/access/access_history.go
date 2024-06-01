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

package access

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mileusna/useragent"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
	"net/url"
)

type sAccessHistory struct {
	Pool *grpool.Pool
}

func init() {
	service.RegisterAccessHistory(NewAccessHistory())
}

func NewAccessHistory() *sAccessHistory {
	pool := grpool.New()
	return &sAccessHistory{Pool: pool}
}

// List 分页读取
func (s *sAccessHistory) List(ctx context.Context, in *do.AccessHistoryListInput) (out *do.AccessHistoryListOutput, err error) {
	out, err = dao.AccessHistory.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sAccessHistory) Add(ctx context.Context, in *do.AccessHistory) (lastInsertId int64, err error) {
	lastInsertId, err = dao.AccessHistory.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// OperateAccess  异步记录日志
func (s *sAccessHistory) OperateAccess(r *ghttp.Request) {
	now := gtime.Now()
	u, _ := url.Parse(r.RequestURI)
	user := service.BizCtx().GetUser(r.GetCtx())

	var userId uint = 0
	var AccessClientId string = ""

	if !g.IsEmpty(user) {
		userId = user.UserId
	} else {
	}

	ua := useragent.Parse(r.UserAgent())

	// 获取浏览器名称和版本号
	browserName := ua.Name
	browserVersion := ua.Version

	fmt.Println("Browser:", browserName, browserVersion)

	// 获取操作系统名称
	osName := ua.OS
	fmt.Println("Operating System:", osName)

	AccessClientId = r.GetSessionId()

	action := do.AccessHistory{
		//UserAccount: user.UserAccount,
		//UserName:    user.UserNickname,
		////Menu:         menu,
		//LogUrl:    u.Path,
		//LogParam:  r.GetMap(),
		//LogMethod: r.Method,
		//LogIp:     utility.GetClientIp(r),
		//LogDate:   gtime.Now(),
		//LogTime:   gtime.Now(),

		UserId:               userId,         // 用户编号
		AccessClientId:       AccessClientId, // 唯一客户编号
		AccessOs:             osName,         // 操作系统
		AccessBrowserName:    browserName,    // 浏览器名称
		AccessBrowserVersion: browserVersion, // 浏览器版本
		AccessSpider:         ua.Bot,         // 搜索引擎
		//AccessCountry        interface{} // 国家
		//AccessProvince       interface{} // 省份
		//AccessCity           interface{} // 市
		//AccessCounty         interface{} // 区
		//AccessLang           interface{} // 语言
		AccessIp:    r.GetClientIp(),      // 访问IP
		AccessUrl:   u.Path,               // 访问地址
		AccessTime:  now.TimestampMilli(), // 访问时间
		AccessYear:  now.Format("Y"),      // 年
		AccessMonth: now.Format("n"),      // 月
		AccessDay:   now.Format("j"),      // 日
		AccessHour:  now.Format("G"),      // 时
		//AccessDate           *gtime.Time // 年月日
		//AccessDatetime       *gtime.Time // 时间
		//AccessReferDomain    interface{} // 来源
		AccessReferUrl: r.GetReferer(), // 来源
		AccessMobile:   ua.Mobile,      // 是否手机
		//AccessPad            interface{} // 是否平板
		//AccessPc             interface{} // 是否PC
		//AccessDevice         interface{} // 终端(ENUM):1-Phone;2-Pad;3-Pc
		//AccessType           interface{} // 终端来源(ENUM):2310-其它;2311-pc;2312-H5;2313-APP;2314-小程序
		//AccessFrom           interface{} // 终端来源(ENUM):2320-其它;2321-微信;2322-百度;2323-支付宝;2324-头条
		AccessData: r.GetMap(), // 请求数据
	}

	s.Pool.Add(r.GetCtx(), func(ctx context.Context) {
		//写入日志数据
		s.Add(r.GetCtx(), &action)
	})
}
