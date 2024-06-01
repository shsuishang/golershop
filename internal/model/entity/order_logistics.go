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

package entity

// OrderLogistics is the golang structure for table order_logistics.
type OrderLogistics struct {
	OrderLogisticsId    uint64 `json:"order_logistics_id"    ` // 订单物流编号
	OrderId             string `json:"order_id"              ` // 订单编号
	StockBillId         string `json:"stock_bill_id"         ` // 出入库单据id=stock_bill_id
	OrderTrackingNumber string `json:"order_tracking_number" ` // 订单物流单号AIRWAY BILL number
	LogisticsExplain    string `json:"logistics_explain"     ` // 卖家备注发货备忘
	LogisticsTime       uint64 `json:"logistics_time"        ` // 发货时间配送时间
	LogisticsId         uint   `json:"logistics_id"          ` // 对应快递公司
	ExpressName         string `json:"express_name"          ` // 快递名称
	ExpressId           int    `json:"express_id"            ` // 快递编号
	SsId                uint   `json:"ss_id"                 ` // 送货编号
	LogisticsPhone      string `json:"logistics_phone"       ` // 送货联系电话
	LogisticsMobile     string `json:"logistics_mobile"      ` // 送货联系手机
	LogisticsContacter  string `json:"logistics_contacter"   ` // 送货联系人
	LogisticsAddress    string `json:"logistics_address"     ` // 送货联系地址
	LogisticsPostcode   string `json:"logistics_postcode"    ` // 邮政编码
	LogisticsEnable     bool   `json:"logistics_enable"      ` // 是否有效(BOOL):1-有效; 0-无效
}
