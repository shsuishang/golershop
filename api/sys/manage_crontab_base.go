package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo manage
type CrontabBaseAdd struct {
	CrontabName        string `json:"crontab_name"          ` // 任务名称
	CrontabFile        string `json:"crontab_file"          ` // 任务脚本
	CrontabLastExeTime uint   `json:"crontab_last_exe_time" ` // 上次执行时间
	CrontabNextExeTime uint   `json:"crontab_next_exe_time" ` // 下次执行时间
	CrontabMinute      string `json:"crontab_minute"        ` // 分钟(LIST):*-每分;1-1; 2-2; 3-3; 4-4; 5-5; 6-6; 7-7; 8-8; 9-9; 10-10; 11-11; 12-12; 13-13; 14-14; 15-15; 16-16; 17-17; 18-18; 19-19; 20-20; 21-21; 22-22; 23-23; 24-24; 25-25; 26-26; 27-27; 28-28; 29-29; 30-30; 31-31; 32-32; 33-33; 34-34; 35-35; 36-36; 37-37; 38-38; 39-39; 40-40; 41-41; 42-42; 43-43; 44-44; 45-45; 46-46; 47-47; 48-48; 49-49; 50-50; 51-51; 52-52; 53-53; 54-54; 55-55; 56-56; 57-57; 58-58; 59-59; 60-60
	CrontabHour        string `json:"crontab_hour"          ` // 小时(LIST):*-任意;1-1; 2-2; 3-3; 4-4; 5-5; 6-6; 7-7; 8-8; 9-9; 10-10; 11-11; 12-12; 13-13; 14-14; 15-15; 16-16; 17-17; 18-18; 19-19; 20-20; 21-21; 22-22; 23-23; 24-24
	CrontabDay         string `json:"crontab_day"           ` // 每天(LIST):*-任意;1-1; 2-2; 3-3; 4-4; 5-5; 6-6; 7-7; 8-8; 9-9; 10-10; 11-11; 12-12; 13-13; 14-14; 15-15; 16-16; 17-17; 18-18; 19-19; 20-20; 21-21; 22-22; 23-23; 24-24; 25-25; 26-26; 27-27; 28-28; 29-29; 30-30
	CrontabMonth       string `json:"crontab_month"         ` // 每月(LIST):*-任意;1-1月; 2-2月; 3-3月; 4-4月; 5-5月; 6-6月; 7-7月; 8-8月; 9-9月; 10-10月; 11-11月; 12-12月
	CrontabWeek        string `json:"crontab_week"          ` // 每周(LIST):*-每周; 1-周一;2-周二;3-周三;4-周四;5-周五;6-周六;7-周日
	CrontabEnable      bool   `json:"crontab_enable"        ` // 是否启用(ENUM):0-禁用; 1-启用
	CrontabBuildin     bool   `json:"crontab_buildin"       ` // 是否内置(ENUM):0-否; 1-是
	CrontabRemark      string `json:"crontab_remark"        ` // 任务备注
}

type CrontabBaseEditReq struct {
	g.Meta `path:"/manage/sys/crontabBase/edit" tags:"任务管理" method:"post" summary:"任务编辑接口"`

	CrontabId uint `json:"crontab_id"            ` // 任务编号
	CrontabBaseAdd
}

type CrontabBaseEditRes struct {
	CrontabId interface{} `json:"crontab_id"        ` // 任务编号
}

type CrontabBaseAddReq struct {
	g.Meta `path:"/manage/sys/crontabBase/add" tags:"任务管理" method:"post" summary:"任务新增接口"`

	CrontabBaseAdd
}

type CrontabBaseRemoveReq struct {
	g.Meta `path:"/manage/sys/crontabBase/remove" tags:"任务管理" method:"post" summary:"任务删除接口"`

	CrontabId uint `json:"crontab_id"            ` // 任务编号
}

type CrontabBaseRemoveRes struct {
}

type CrontabBaseListReq struct {
	g.Meta `path:"/manage/sys/crontabBase/list" tags:"任务管理" method:"get" summary:"任务列表接口"`
	ml.BaseList

	CrontabName        string `json:"crontab_name"          ` // 任务名称
	CrontabFile        string `json:"crontab_file"          ` // 任务脚本
	CrontabLastExeTime uint   `json:"crontab_last_exe_time" ` // 上次执行时间
	CrontabNextExeTime uint   `json:"crontab_next_exe_time" ` // 下次执行时间
}

type CrontabBaseListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type CrontabBaseEditStateReq struct {
	g.Meta `path:"/manage/sys/crontabBase/editState" tags:"任务管理" method:"post" summary:"任务编辑接口"`

	CrontabId     uint `json:"crontab_id"            ` // 任务编号
	CrontabEnable bool `json:"crontab_enable"        ` // 是否启用(ENUM):0-禁用; 1-启用
}

type CrontabBaseEditStateRes struct {
	CrontabId interface{} `json:"crontab_id"   dc:"任务编号"`
}
