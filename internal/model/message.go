package model

import (
	"golershop.cn/internal/model/entity"
)

type UserMessageVo struct {
	entity.UserMessage
	UserAvatar      string `json:"user_avatar"     description:"用户头像"`
	UserOtherAvatar string `json:"user_other_avatar"     description:"其他用户头像"`

	Num            int                 `json:"num"            description:"站内信数量"`
	MsgRow         *entity.UserMessage `json:"msg_row"        description:"聊天网址"`
	MessageOtherId uint                `json:"message_other_id" description:"消息编号"`
	Userinfo       *MessageUserInfoVo  `json:"userinfo"       description:"站内信userInfo"`
	Content        *MessageContentVo   `json:"content"        description:"站内信content"`
	RedNumber      int                 `json:"red_number"     description:"已读信息数"`
	UnreadNumber   int                 `json:"unread_number"  description:"未读信息数"`
}

type UserMessageOutput struct {
	Items   []*UserMessageVo `json:"items"    dc:"消息列表"`
	Page    int              `json:"page"`    // 分页号码
	Total   int              `json:"total"`   // 总页数
	Records int              `json:"records"` // 数据总数
	Size    int              `json:"size"`    // 单页数量
}

type MessageUserInfoVo struct {
	Username string `json:"username" description:"用户昵称"`
	Uid      int    `json:"uid"      description:"uid"`
	Face     string `json:"face"     description:"uid"`
}

type MessageContentVo struct {
	MessageContent  string  `json:"message_content" description:"消息内容"`
	MessageLength   int     `json:"message_length"  description:"消息长度"`
	MessageW        uint    `json:"message_w"       description:"图片宽度"`
	MessageH        uint    `json:"message_h"       description:"图片高度"`
	ItemUnitPrice   float64 `json:"item_unit_price" description:"商品价格单价"`
	ProductItemName string  `json:"product_item_name" description:"商品名称"`
	ProductImage    string  `json:"product_image"     description:"商品图片"`
	ItemId          uint64  `json:"item_id"           description:"商品编号"`
}

type MsgCountInput struct {
	RecentlyFlag bool `json:"recently_flag"` // 最近消息
	UserId       uint `json:"user_id" `      // 买家编号
}

// UserMessageAddInput 站内信添加
type UserMessageAddInput struct {
	MessageCat     string `json:"messageCat"`     // 消息类型
	UserNickname   string `json:"userNickname"`   // 相关昵称:发送者或者接收者
	UserOtherId    uint   `json:"userOtherId"`    // 相关用户:发送者或者接收者
	To             string `json:"to"`             // json对象
	MessageContent string `json:"messageContent"` // 消息内容
	MessageTitle   string `json:"messageTitle"`   // 消息标题
	Mine           string `json:"mine"`           // json对象
	MessageLength  int    `json:"messageLength"`  // 消息长度
	MessageW       int    `json:"messageW"`       // 图片宽度
	MessageH       int    `json:"messageH"`       // 图片高度
}

type ChatMessageVo struct {
	entity.UserMessage
	MsgType string         `json:"type"` //
	Msg     *UserMessageVo `json:"msg"`  //
}

type MessageOutput struct {
	Items   []*ChatMessageVo `json:"items"    dc:"消息列表"`
	Page    int              `json:"page"`    // 分页号码
	Total   int              `json:"total"`   // 总页数
	Records int              `json:"records"` // 数据总数
	Size    int              `json:"size"`    // 单页数量
}

// MessageNoticeVo 后台通知对象
type MessageNoticeVo struct {
	Total int              `json:"total"` // 消息数量
	Items []*UserMessageVo `json:"items"` // 站内信数量
}
