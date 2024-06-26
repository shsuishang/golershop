package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model/entity"
)

// =========================== 管理端使用 =============================

type ConfigTypeAdd struct {
	ConfigTypeId     uint   `json:"config_type_id"   dc:"主键编号"   `
	ConfigTypeName   string `json:"config_type_name" dc:"分组名称"   `
	ConfigTypeSort   uint   `json:"config_type_sort" dc:"分组排序:从小到大"   `
	ConfigTypeModule uint   `json:"config_type_module"  ` // 所属模块(ENUM):1001-站点设置;1002-上传设置;1003-运营设置;1004-财务设置;
}

type ConfigTypeAddReq struct {
	g.Meta `path:"/manage/sys/config/addType" tags:"配置管理" method:"post" summary:"配置类型详情接口"`

	ConfigTypeAdd
}

type ConfigTypeEditReq struct {
	g.Meta       `path:"/manage/sys/config/editType" tags:"配置管理" method:"post" summary:"配置类型详情接口"`
	ConfigTypeId uint `json:"config_type_id" v:"required#请输入配置类型编号"   dc:"配置类型信息"`
	ConfigTypeAdd
}

type ConfigTypeEditRes struct {
	ConfigTypeId uint `json:"config_type_id"   dc:"主键编号"   `
}

type ConfigTypeRemoveReq struct {
	g.Meta       `path:"/manage/sys/config/removeType" tags:"配置管理" method:"post" summary:"配置类型删除接口"`
	ConfigTypeId []uint `json:"config_type_id" v:"required#请输入配置类型编号"   dc:"配置类型信息"`
}

type ConfigTypeRemoveRes struct {
}

type ConfigTypeListReq struct {
	g.Meta `path:"/manage/sys/config/listType" tags:"配置管理" method:"get" summary:"配置类型列表接口"`
	Page   int `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size   int `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`
}

type ConfigTypeListRes struct {
	Items   interface{} `json:"items"    dc:"配置分类列表页"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

// ---------------------------- 配置项 -------------------------------

type ConfigBaseAdd struct {
	ConfigKey      string `json:"config_key"   v:"required#请输入配置编码"    dc:"配置编码"     `
	ConfigTitle    string `json:"config_title"   dc:"配置标题" `
	ConfigValue    string `json:"config_value"    dc:"配置值"`
	ConfigOptions  string `json:"config_options"  dc:"配置项"`
	ConfigTypeId   uint   `json:"config_type_id"  v:"required#请输入所属分类" `
	ConfigDatatype string `json:"config_datatype" dc:"数据类型"`
	ConfigNote     string `json:"config_note"     dc:"配置注释"`
	ConfigSort     uint   `json:"config_sort"    dc:"配置排序:从小到大" `
	ConfigEnable   bool   `json:"config_enable"  dc:"是否启用(BOOL):0-禁用;1-启用" `
}

type ConfigBaseAddReq struct {
	g.Meta `path:"/manage/sys/config/add" tags:"配置管理" method:"post" summary:"配置项目详情接口"`

	ConfigBaseAdd
}

type ConfigBaseEditReq struct {
	g.Meta `path:"/manage/sys/config/edit" tags:"配置管理" method:"post" summary:"配置项目详情接口"`

	ConfigBaseAdd
}

type ConfigBaseEditRes struct {
	ConfigKey string `json:"config_key"   v:"required#请输入配置编码"    dc:"配置编码"     `
}

type ConfigBaseRemoveReq struct {
	g.Meta    `path:"/manage/sys/config/remove" tags:"配置管理" method:"post" summary:"配置项目删除接口"`
	ConfigKey []string `json:"config_key" v:"required#请输入配置编码"   dc:"配置编码"`
}

type ConfigBaseRemoveRes struct {
}

type ConfigBaseListReq struct {
	g.Meta       `path:"/manage/sys/config/list" tags:"配置管理" method:"get" summary:"配置项目列表接口"`
	Page         int  `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size         int  `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`
	ConfigTypeId uint `json:"config_type_id" d:"0" v:"required#请输入配置类型编号"   dc:"配置类型信息"`
}

type ConfigBaseListRes struct {
	Items   interface{} `json:"items"    dc:"配置列表页"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ConfigTreeReq struct {
	g.Meta `path:"/manage/sys/config/tree" tags:"配置" method:"get" summary:"站点配置Tree"`
}

type ConfigNode struct {
	entity.ConfigType
	Items []*entity.ConfigType `json:"items"`
}

type ConfigTreeRes []*ConfigNode

type ConfigBaseIndexReq struct {
	g.Meta           `path:"/manage/sys/config/index" tags:"配置管理" method:"get" summary:"配置项目列表接口"`
	ConfigTypeModule uint `json:"config_type_module"  ` // 所属模块(ENUM):1001-站点设置;1002-上传设置;1003-运营设置;1004-财务设置;
}

type ConfigBaseIndexVo struct {
	ConfigTypeId   uint        `json:"config_type_id"   dc:"主键编号"   `
	ConfigTypeName string      `json:"config_type_name" dc:"分组名称"   `
	Items          interface{} `json:"items"    dc:"配置列表页"`
}

type ConfigBaseIndexRes struct {
	Items []ConfigBaseIndexVo `json:"items"    dc:"站点配置数据"`
}

type ConfigBaseEditSiteReq struct {
	g.Meta  `path:"/manage/sys/config/editSite" tags:"配置管理" method:"post" summary:"编辑站点配置"`
	Configs map[string]interface{} `json:"configs"    dc:"配置数据"`
}

type ConfigBaseEditSiteRes struct {
}

type CleanCacheReq struct {
	g.Meta `path:"/manage/sys/cache/clean" tags:"配置管理" method:"post" summary:"清理缓存"`
}

type CleanCacheRes struct {
}
