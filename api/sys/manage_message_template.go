package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type MessageTemplateAdd struct {
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

type MessageTemplateEditReq struct {
	g.Meta `path:"/manage/sys/messageTemplate/edit" tags:"消息模板" method:"post" summary:"消息模板编辑接口"`

	MessageId string `json:"message_id"            ` // 模板编号
	MessageTemplateAdd
}

type MessageTemplateEditRes struct {
	MessageId string `json:"message_id"            ` // 模板编号
}

type MessageTemplateAddReq struct {
	g.Meta `path:"/manage/sys/messageTemplate/add" tags:"消息模板" method:"post" summary:"消息模板编辑接口"`

	MessageTemplateAdd
}

type MessageTemplateRemoveReq struct {
	g.Meta `path:"/manage/sys/messageTemplate/remove" tags:"消息模板" method:"post" summary:"消息模板删除接口"`

	MessageId string `json:"message_id"            ` // 模板编号
}

type MessageTemplateRemoveRes struct {
}

type MessageTemplateListReq struct {
	g.Meta `path:"/manage/sys/messageTemplate/list" tags:"消息模板" method:"get" summary:"消息模板列表接口"`
	ml.BaseList

	MessageId string `json:"message_id"            ` // 模板编号
}

type MessageTemplateListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type MessageTemplateEditStateReq struct {
	g.Meta `path:"/manage/sys/messageTemplate/editState" tags:"消息模板" method:"post" summary:"消息模板状态编辑接口"`

	MessageId     string `json:"message_id"            ` // 模板编号
	MessageEnable bool   `json:"message_enable"        ` // 站内通知(BOOL):0-禁用;1-启用
}

type MessageTemplateEditStateRes struct {
	MessageId string `json:"message_id"            ` // 模板编号
}
