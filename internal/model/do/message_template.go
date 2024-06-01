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
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// MessageTemplate is the golang structure of table sys_message_template for DAO operations like Where/Data.
type MessageTemplate struct {
	g.Meta              `orm:"table:sys_message_template, do:true"`
	MessageId           interface{} // 模板编号
	MessageCode         interface{} // 模板编码
	MessageName         interface{} // 模板名称
	MessageEmailTitle   interface{} // 邮件标题
	MessageEmailContent interface{} // 邮件内容
	MessageContent      interface{} // 站内消息
	MessageSms          interface{} // 短信内容
	MessageApp          interface{} // APP内容
	MessageType         interface{} // 消息类型(ENUM):1-用户;2-商家;3-平台;
	MessageEnable       interface{} // 站内通知(BOOL):0-禁用;1-启用
	MessageSmsEnable    interface{} // 短息通知(BOOL):0-禁用;1-启用
	MessageEmailEnable  interface{} // 邮件通知(BOOL):0-禁用;1-启用
	MessageWechatEnable interface{} // 微信通知(BOOL):0-禁用;1-启用
	MessageXcxEnable    interface{} // 小程序通知(BOOL):0-禁用;1-启用
	MessageAppEnable    interface{} // APP推送(BOOL):0-禁用;1-启用
	MessageSmsForce     interface{} // 手机短信(BOOL):0-不强制;1-强制
	MessageEmailForce   interface{} // 邮件(BOOL):0-不强制;1-强制
	MessageAppForce     interface{} // APP(BOOL):0-不强制;1-强制
	MessageForce        interface{} // 站内信(BOOL):0-不强制;1-强制
	MessageCategory     interface{} // 消息分组(ENUM):0-默认消息;1-公告消息;2-订单消息;3-商品消息;4-余额卡券;5-服务消息
	MessageOrder        interface{} // 消息排序
	MessageTplId        interface{} // 模板编号
}

type MessageTemplateListInput struct {
	ml.BaseList
	Where MessageTemplate // 查询条件
}

type MessageTemplateListOutput struct {
	Items   []*entity.MessageTemplate // 列表
	Page    int                       // 分页号码
	Total   int                       // 总页数
	Records int                       // 数据总数
	Size    int                       // 单页数量
}

type MessageTemplateListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
