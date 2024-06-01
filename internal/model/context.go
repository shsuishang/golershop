package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"golershop.cn/internal/model/entity"
)

type Context struct {
	Session   *ghttp.Session // Session in context.
	User      *ContextUser   // User in context.
	Tx        int
	CacheKeys []string
}

type ContextUser struct {
	UserId       uint   // User ID.
	UserAccount  string // User Account.
	UserNickname string // User Nickname.
	UserSalt     string // User Salt .
	RoleId       uint   // 角色编号  0-用户;2-商家;3-门店;9-平台;
	SiteId       uint   // 分站编号
	StoreId      uint   // 店铺编号
	ChainId      uint   // 门店编号
	ClientId     uint   // 后台管理-admin=1;移动端front=0

	Roles       []*entity.UserRole //角色列表
	Authorities []*entity.MenuBase //权限列表
}
