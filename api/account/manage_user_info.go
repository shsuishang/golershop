package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type UserInfoAdd struct {
	UserAccount          string      `json:"user_account" v:"required#请输入用户账号"        `     // 用户账号
	UserNickname         string      `json:"user_nickname" v:"required#请输入用户昵称"       `     // 用户昵称
	UserAvatar           string      `json:"user_avatar"                                  ` // 用户头像
	UserState            uint        `json:"user_state"                                   ` // 状态(ENUM):0-锁定;1-已激活;2-未激活;
	UserMobile           string      `json:"user_mobile" v:"required#请输入手机号码"         `     // 手机号码(mobile)
	UserIntl             string      `json:"user_intl"              `                       // 国家编码
	UserGender           uint        `json:"user_gender"            `                       // 性别(ENUM):0-保密;1-男;  2-女;
	UserBirthday         *gtime.Time `json:"user_birthday"          `                       // 生日(DATE)
	UserEmail            string      `json:"user_email"             `                       // 用户邮箱(email)
	UserLevelId          uint        `json:"user_level_id" v:"required#请选择用户等级"      `      // 用户等级ID
	UserIsAuthentication uint        `json:"user_is_authentication" `                       // 0-未认证 1-待审核 2-认证通过 3-认证失败

	Password string `json:"password"     ` // 用户密码
}
type UserInfoEditReq struct {
	g.Meta `path:"/manage/account/userInfo/edit" tags:"会员管理" method:"post" summary:"用户编辑接口"`

	UserId uint `json:"user_id"   v:"required#请输入会员编号"    dc:"会员编号"     `
	UserInfoAdd
}

type UserInfoAddRes struct {
	UserId interface{} `json:"user_id"   dc:"会员信息"`
}

type UserInfoEditRes struct {
	UserId interface{} `json:"user_id"   dc:"会员信息"`
}

type UserInfoAddReq struct {
	g.Meta `path:"/manage/account/userInfo/add" tags:"会员管理" method:"post" summary:"用户编辑接口"`

	UserInfoAdd
}

type UserInfoRemoveReq struct {
	g.Meta `path:"/manage/account/userInfo/remove" tags:"会员管理" method:"post" summary:"用户删除接口"`
	UserId []uint `json:"user_id" v:"required#请输入会员编号"   dc:"会员信息"`
}

type UserInfoRemoveRes struct {
}

type UserInfoListReq struct {
	g.Meta `path:"/manage/account/userInfo/list" tags:"会员管理" method:"get" summary:"用户列表接口"`
	ml.BaseList

	UserId       uint   `json:"user_id"                `            // 用户编号
	UserAccount  string `json:"user_account"   type:"LIKE"        ` // 用户账号
	UserNickname string `json:"user_nickname"  type:"LIKE"        ` // 用户昵称
	UserMobile   string `json:"user_mobile"    type:"LIKE"        ` //用户号码
}

type UserInfoListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
