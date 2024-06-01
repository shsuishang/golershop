package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
)

// start fo front

// start fo manage
type UserMessageAdd struct {
	MessageParentId   uint   `json:"message_parent_id"   ` // 上级编号
	UserId            uint   `json:"user_id"             ` // 所属用户:发送者或者接收者，如果message_kind=1则为当前用户发送的消息。
	UserNickname      string `json:"user_nickname"       ` // 用户昵称
	MessageKind       uint   `json:"message_kind"        ` // 消息种类(ENUM):1-发送消息;2-接收消息
	UserOtherId       uint   `json:"user_other_id"       ` // 相关用户:发送者或者接收者
	UserOtherNickname string `json:"user_other_nickname" ` // 相关昵称:发送者或者接收者
	MessageTitle      string `json:"message_title"       ` // 消息标题
	MessageContent    string `json:"message_content"     ` // 消息内容
	MessageTime       uint64 `json:"message_time"        ` // 发送时间
	MessageIsRead     bool   `json:"message_is_read"     ` // 是否读取(BOOL):0-未读;1-已读
	MessageIsDelete   bool   `json:"message_is_delete"   ` // 是否删除(BOOL):0-正常状态;1-删除状态
	MessageType       uint   `json:"message_type"        ` // 消息类型(ENUM):1-系统消息;2-用户消息
	MessageCat        string `json:"message_cat"         ` // 消息类型(ENUM):text-文本消息;img-图片消息;video-视频消息;file:文件;location:位置;redpack:红包
	MessageDataType   uint   `json:"message_data_type"   ` // 消息分类:0-默认消息;1-公告消息;2-订单消息;3-商品消息
	MessageDataId     string `json:"message_data_id"     ` // 消息数据:商品编号|订单编号
	MessageLength     int    `json:"message_length"      ` // 消息长度
	MessageW          uint   `json:"message_w"           ` // 图片宽度
	MessageH          uint   `json:"message_h"           ` // 图片高度
}
type UserMessageEditReq struct {
	g.Meta `path:"/manage/account/userMessage/edit" tags:"用户消息" method:"post" summary:"用户消息编辑接口"`

	MessageId uint `json:"message_id"          ` // 消息编号   `
	UserMessageAdd
}

type UserMessageAddRes model.UserMessageVo

type UserMessageEditRes struct {
	MessageId interface{} `json:"message_id"   dc:"用户消息信息"`
}
type UserMessageAddReq struct {
	g.Meta `path:"/manage/account/userMessage/add" tags:"用户消息" method:"post" summary:"用户消息编辑接口"`

	UserMessageAdd
}

type UserMessageRemoveReq struct {
	g.Meta    `path:"/manage/account/userMessage/remove" tags:"用户消息" method:"post" summary:"用户消息删除接口"`
	MessageId []uint `json:"message_id" v:"required#请输入用户消息编号"   dc:"用户消息信息"`
}

type UserMessageRemoveRes struct {
}

type UserMessageListReq struct {
	g.Meta `path:"/manage/account/userMessage/list" tags:"用户消息" method:"get" summary:"用户消息列表接口"`
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
}

type UserMessageListRes struct {
	Items   interface{} `json:"items"    dc:"消息列表"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type MessageNoticeReq struct {
	g.Meta `path:"/manage/account/userMessage/getNotice" tags:"用户消息" method:"get" summary:"用户消息列表接口"`
}

type MessageNoticeRes struct {
	Items []*model.UserMessageVo `json:"items"    dc:"站内信数量"` //站内信数量
	Total int                    `json:"total"`                    // 消息数量
}

