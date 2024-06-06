package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/entity"
)

type ImConfigReq struct {
	g.Meta `path:"/front/account/userMessage/getImConfig" tags:"用户消息" method:"get" summary:"IM配置"`

	UserOtherId uint `json:"user_other_id"       ` // 相关用户:发送者或者接收者
}

type ImConfigRes struct {
}

type KefuConfigReq struct {
	g.Meta `path:"/front/account/userMessage/getKefuConfig" tags:"用户消息" method:"get" summary:"客服配置"`

	UserOtherId uint `json:"user_other_id"       ` // 相关用户:发送者或者接收者
}

type KefuConfigRes struct {
}

type MessageAddReq struct {
	g.Meta `path:"/front/account/userMessage/add" tags:"用户消息" method:"post" summary:"用户消息编辑接口"`

	UserMessageAdd
}

type MessageAddRes model.UserMessageVo

type NoticeListReq struct {
	g.Meta `path:"/front/account/userMessage/getNotice" tags:"用户消息" method:"get" summary:"用户消息列表接口"`
	ml.BaseList

	MessageId       uint   `json:"message_id"`           // 用户消息编号
	MessageTitle    string `json:"message_title"       ` // 消息标题
	MessageContent  string `json:"message_content"     ` // 消息内容
	UserId          uint   `json:"user_id"             ` // 所属用户:发送者或者接收者，如果message_kind=1则为当前用户发送的消息。
	MessageKind     uint   `json:"message_kind"        ` // 消息种类(ENUM):1-发送消息;2-接收消息
	UserOtherId     uint   `json:"user_other_id"       ` // 相关用户:发送者或者接收者
	MessageIsDelete bool   `json:"message_is_delete"   ` // 是否删除(BOOL):0-正常状态;1-删除状态
	MessageType     uint   `json:"message_type"        ` // 消息类型(ENUM):1-系统消息;2-用户消息
	MessageCat      string `json:"message_cat"         ` // 消息类型(ENUM):text-文本消息;img-图片消息;video-视频消息;file:文件;location:位置;redpack:红包
	MessageDataType uint   `json:"message_data_type"   ` // 消息分类:0-默认消息;1-公告消息;2-订单消息;3-商品消息
}

type NoticeListRes struct {
	Items   interface{} `json:"items"    dc:"消息列表"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type MsgCountReq struct {
	g.Meta `path:"/front/account/userMessage/getMsgCount" tags:"用户消息" method:"get" summary:"用户通知消息数量"`

	UserId       uint `json:"user_id" `      // 买家编号
	RecentlyFlag bool `json:"recently_flag"` // 最近消息
}

type MsgCountRes struct {
	Items   interface{} `json:"items"    dc:"消息列表"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type MessageListReq struct {
	g.Meta `path:"/front/account/userMessage/list" tags:"用户消息" method:"get" summary:"用户消息列表接口"`
	ml.BaseList

	MessageId       string `json:"message_id"`                      // 用户消息编号
	MessageTitle    string `json:"message_title"  type:"LIKE"     ` // 消息标题
	MessageContent  string `json:"message_content"     `            // 消息内容
	UserId          uint   `json:"user_id"             `            // 所属用户:发送者或者接收者，如果message_kind=1则为当前用户发送的消息。
	MessageKind     uint   `json:"message_kind"        `            // 消息种类(ENUM):1-发送消息;2-接收消息
	UserOtherId     uint   `json:"user_other_id"       `            // 相关用户:发送者或者接收者
	MessageIsDelete bool   `json:"message_is_delete"   `            // 是否删除(BOOL):0-正常状态;1-删除状态
	MessageType     uint   `json:"message_type"        `            // 消息类型(ENUM):1-系统消息;2-用户消息
	MessageCat      string `json:"message_cat"         `            // 消息类型(ENUM):text-文本消息;img-图片消息;video-视频消息;file:文件;location:位置;redpack:红包
	MessageDataType uint   `json:"message_data_type"   `            // 消息分类:0-默认消息;1-公告消息;2-订单消息;3-商品消息

	SourceType uint `json:"source_type"   ` // 来源

}

type MessageListRes struct {
	Items   interface{} `json:"items"    dc:"消息列表"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type MessageGetReq struct {
	g.Meta `path:"/front/account/userMessage/get" tags:"用户消息" method:"get" summary:"读取短消息"`

	MessageId uint `json:"message_id"` // 用户消息编号
}

type MessageGetRes entity.UserMessage

type MsgReadReq struct {
	g.Meta `path:"/front/account/userMessage/setRead" tags:"用户消息" method:"get" summary:"设置为已读"`

	MessageId   uint `json:"message_id"`    // 消息编号
	UserOtherId uint `json:"user_other_id"` // 用户编号
}

type MsgReadRes struct {
}

type ChatMsgListReq struct {
	g.Meta `path:"/front/account/userMessage/listChatMsg" tags:"用户消息" method:"get" summary:"用户消息列表接口"`
	ml.BaseList

	MessageId       string `json:"message_id"`           // 用户消息编号
	MessageTitle    string `json:"message_title"       ` // 消息标题
	MessageContent  string `json:"message_content"     ` // 消息内容
	UserId          uint   `json:"user_id"             ` // 所属用户:发送者或者接收者，如果message_kind=1则为当前用户发送的消息。
	MessageKind     uint   `json:"message_kind"        ` // 消息种类(ENUM):1-发送消息;2-接收消息
	UserOtherId     uint   `json:"user_other_id"       ` // 相关用户:发送者或者接收者
	MessageIsDelete bool   `json:"message_is_delete"   ` // 是否删除(BOOL):0-正常状态;1-删除状态
	MessageType     uint   `json:"message_type"        ` // 消息类型(ENUM):1-系统消息;2-用户消息
	MessageCat      string `json:"message_cat"         ` // 消息类型(ENUM):text-文本消息;img-图片消息;video-视频消息;file:文件;location:位置;redpack:红包
	MessageDataType uint   `json:"message_data_type"   ` // 消息分类:0-默认消息;1-公告消息;2-订单消息;3-商品消息

	SourceType uint `json:"source_type"   ` // 来源
}

type ChatMsgListRes model.MessageOutput

type MessageNumReq struct {
	g.Meta `path:"/front/account/userMessage/getMessageNum" tags:"用户消息" method:"get" summary:"信息数"`

	MessageId       string `json:"message_id"`           // 用户消息编号
	MessageTitle    string `json:"message_title"       ` // 消息标题
	MessageContent  string `json:"message_content"     ` // 消息内容
	UserId          uint   `json:"user_id"             ` // 所属用户:发送者或者接收者，如果message_kind=1则为当前用户发送的消息。
	MessageKind     uint   `json:"message_kind"        ` // 消息种类(ENUM):1-发送消息;2-接收消息
	UserOtherId     uint   `json:"user_other_id"       ` // 相关用户:发送者或者接收者
	MessageIsDelete bool   `json:"message_is_delete"   ` // 是否删除(BOOL):0-正常状态;1-删除状态
	MessageType     uint   `json:"message_type"        ` // 消息类型(ENUM):1-系统消息;2-用户消息
	MessageCat      string `json:"message_cat"         ` // 消息类型(ENUM):text-文本消息;img-图片消息;video-视频消息;file:文件;location:位置;redpack:红包
	MessageDataType uint   `json:"message_data_type"   ` // 消息分类:0-默认消息;1-公告消息;2-订单消息;3-商品消息

	SourceType uint `json:"source_type"   ` // 来源
}

type MessageNumRes model.UserMessageVo
