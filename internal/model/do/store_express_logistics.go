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

// StoreExpressLogistics is the golang structure of table shop_store_express_logistics for DAO operations like Where/Data.
type StoreExpressLogistics struct {
	g.Meta             `orm:"table:shop_store_express_logistics, do:true"`
	LogisticsId        interface{} // 物流编号
	StoreId            interface{} // 店铺编号
	ChainId            interface{} // 门店编号
	LogisticsName      interface{} // 物流名称
	LogisticsPinyin    interface{} // 物流拼音
	LogisticsNumber    interface{} // 公司编号
	LogisticsState     interface{} // 面单状态(BOOL):1-启用;0-禁用
	ExpressId          interface{} // 快递编号
	ExpressName        interface{} // 快递名称
	LogisticsIsDefault interface{} // 是否为默认(BOOL):1-默认;0-非默认
	LogisticsIntl      interface{} // 国家编码
	LogisticsMobile    interface{} // 联系手机
	LogisticsContacter interface{} // 联系人
	LogisticsAddress   interface{} // 联系地址
	LogisticsFee       interface{} // 物流运费
	LogisticsIsEnable  interface{} // 是否启用(BOOL):1-启用;0-禁用
	CreateTime         *gtime.Time // 创建时间
}

type StoreExpressLogisticsListInput struct {
	ml.BaseList
	Where StoreExpressLogistics // 查询条件
}

type StoreExpressLogisticsListOutput struct {
	Items   []*entity.StoreExpressLogistics // 列表
	Page    int                             // 分页号码
	Total   int                             // 总页数
	Records int                             // 数据总数
	Size    int                             // 单页数量
}

type StoreExpressLogisticsListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
