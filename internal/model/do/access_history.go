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

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// AccessHistory is the golang structure of table sys_access_history for DAO operations like Where/Data.
type AccessHistory struct {
	g.Meta               `orm:"table:sys_access_history, do:true"`
	AccessId             interface{} // 编号
	UserId               interface{} // 用户编号
	ItemId               interface{} // SKU编号
	AccessClientId       interface{} // 唯一客户编号
	AccessOs             interface{} // 操作系统
	AccessBrowserName    interface{} // 浏览器名称
	AccessBrowserVersion interface{} // 浏览器版本
	AccessSpider         interface{} // 搜索引擎
	AccessCountry        interface{} // 国家
	AccessProvince       interface{} // 省份
	AccessCity           interface{} // 市
	AccessCounty         interface{} // 区
	AccessLang           interface{} // 语言
	AccessIp             interface{} // 访问IP
	AccessMethod         interface{} // 请求方法
	AccessUrl            interface{} // 访问地址
	AccessTime           interface{} // 访问时间
	AccessYear           interface{} // 年
	AccessMonth          interface{} // 月
	AccessDay            interface{} // 日
	AccessHour           interface{} // 时
	AccessDate           *gtime.Time // 年月日
	AccessDatetime       *gtime.Time // 时间
	AccessReferDomain    interface{} // 来源
	AccessReferUrl       interface{} // 来源
	AccessMobile         interface{} // 是否手机
	AccessPad            interface{} // 是否平板
	AccessPc             interface{} // 是否PC
	AccessDevice         interface{} // 终端(ENUM):1-Phone;2-Pad;3-Pc
	AccessType           interface{} // 终端来源(ENUM):2310-其它;2311-pc;2312-H5;2313-APP;2314-小程序
	AccessFrom           interface{} // 终端来源(ENUM):2320-其它;2321-微信;2322-百度;2323-支付宝;2324-头条
	AccessData           interface{} // 请求数据
}

type AccessHistoryListInput struct {
	ml.BaseList
	Where AccessHistory // 查询条件
}

type AccessHistoryListOutput struct {
	Items   []*entity.AccessHistory // 列表
	Page    int                     // 分页号码
	Total   int                     // 总页数
	Records int                     // 数据总数
	Size    int                     // 单页数量
}

type AccessHistoryListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
