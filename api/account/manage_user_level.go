package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type UserLevelAdd struct {
	UserLevelId        string      `json:"user_level_id"`                            // 等级编号
	UserLevelName      string      `json:"user_level_name" v:"required#请输入等级名称"    ` // 等级名称
	UserLevelExp       uint        `json:"user_level_exp"  v:"required#请输入升级经验值"   ` // 升级经验值
	UserLevelSpend     string      `json:"user_level_spend" v:"required#请输入累计消费" `   // 累计消费
	UserLevelLogo      string      `json:"user_level_logo" v:"required#请上传等级图标"   `  // LOGO
	UserLevelRate      string      `json:"user_level_rate"`                          // 折扣率百分比
	UserLevelTime      *gtime.Time `json:"user_level_time"`                          // 修改时间
	UserLevelIsBuildin bool        `json:"user_level_is_buildin"`                    // 系统内置(BOOL):0-否;1-是
}
type UserLevelEditReq struct {
	g.Meta `path:"/manage/account/userLevel/edit" tags:"会员等级" method:"post" summary:"等级编辑接口"`

	UserLevelId uint `json:"user_level_id"   v:"required#请输入会员等级编号"    dc:"等级编号"     `
	UserLevelAdd
}

type UserLevelAddRes struct {
	UserLevelId interface{} `json:"user_level_id"   dc:"等级信息"`
}

type UserLevelEditRes struct {
	UserLevelId interface{} `json:"user_level_id"   dc:"等级信息"`
}

type UserLevelAddReq struct {
	g.Meta `path:"/manage/account/userLevel/add" tags:"会员等级" method:"post" summary:"等级编辑接口"`

	UserLevelAdd
}

type UserLevelRemoveReq struct {
	g.Meta      `path:"/manage/account/userLevel/remove" tags:"会员等级" method:"post" summary:"等级删除接口"`
	UserLevelId []uint `json:"user_level_id" v:"required#请输入会员等级编号"   dc:"等级信息"`
}

type UserLevelRemoveRes struct {
}

type UserLevelListReq struct {
	g.Meta `path:"/manage/account/userLevel/list" tags:"会员等级" method:"get" summary:"等级列表接口"`
	ml.BaseList

	UserLevelName string `json:"user_level_name" type:"LIKE"  ` // 等级名称
}

type UserLevelListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
