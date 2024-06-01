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

// MessageTemplate is the golang structure for table message_template.
type MessageTemplate struct {
	MessageId           string `json:"message_id"            ` // 模板编号
	MessageCode         string `json:"message_code"          ` // 模板编码
	MessageName         string `json:"message_name"          ` // 模板名称
	MessageEmailTitle   string `json:"message_email_title"   ` // 邮件标题
	MessageEmailContent string `json:"message_email_content" ` // 邮件内容
	MessageContent      string `json:"message_content"       ` // 站内消息
	MessageSms          string `json:"message_sms"           ` // 短信内容
	MessageApp          string `json:"message_app"           ` // APP内容
	MessageType         uint   `json:"message_type"          ` // 消息类型(ENUM):1-用户;2-商家;3-平台;
	MessageEnable       bool   `json:"message_enable"        ` // 站内通知(BOOL):0-禁用;1-启用
	MessageSmsEnable    bool   `json:"message_sms_enable"    ` // 短息通知(BOOL):0-禁用;1-启用
	MessageEmailEnable  bool   `json:"message_email_enable"  ` // 邮件通知(BOOL):0-禁用;1-启用
	MessageWechatEnable bool   `json:"message_wechat_enable" ` // 微信通知(BOOL):0-禁用;1-启用
	MessageXcxEnable    bool   `json:"message_xcx_enable"    ` // 小程序通知(BOOL):0-禁用;1-启用
	MessageAppEnable    bool   `json:"message_app_enable"    ` // APP推送(BOOL):0-禁用;1-启用
	MessageSmsForce     bool   `json:"message_sms_force"     ` // 手机短信(BOOL):0-不强制;1-强制
	MessageEmailForce   bool   `json:"message_email_force"   ` // 邮件(BOOL):0-不强制;1-强制
	MessageAppForce     bool   `json:"message_app_force"     ` // APP(BOOL):0-不强制;1-强制
	MessageForce        bool   `json:"message_force"         ` // 站内信(BOOL):0-不强制;1-强制
	MessageCategory     uint   `json:"message_category"      ` // 消息分组(ENUM):0-默认消息;1-公告消息;2-订单消息;3-商品消息;4-余额卡券;5-服务消息
	MessageOrder        uint   `json:"message_order"         ` // 消息排序
	MessageTplId        string `json:"message_tpl_id"        ` // 模板编号
}
