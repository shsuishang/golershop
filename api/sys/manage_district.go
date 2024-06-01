package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

// start fo front

// start fo manage
type DistrictBaseAdd struct {
	DistrictId       uint    `json:"district_id"        ` // 地区编号
	DistrictName     string  `json:"district_name"      ` // 地区名称
	DistrictParentId uint    `json:"district_parent_id" ` // 上级编号
	DistrictLevel    uint    `json:"district_level"     ` // 地区等级
	DistrictCitycode string  `json:"district_citycode"  ` // 区号
	DistrictZipcode  string  `json:"district_zipcode"   ` // 邮编
	DistrictLng      float64 `json:"district_lng"       ` // 经度
	DistrictLat      float64 `json:"district_lat"       ` // 维度
	DistrictSort     uint    `json:"district_sort"      ` // 地区排序
}
type DistrictBaseEditReq struct {
	g.Meta `path:"/manage/sys/districtBase/edit" tags:"区域管理" method:"post" summary:"分类编辑接口"`

	DistrictId uint `json:"district_id"        ` // 地区编号
	DistrictBaseAdd
}

type DistrictBaseEditRes struct {
	DistrictId interface{} `json:"district_id"   dc:"区域管理信息"`
}

type DistrictBaseAddReq struct {
	g.Meta `path:"/manage/sys/districtBase/add" tags:"区域管理" method:"post" summary:"区域管理添加接口"`

	DistrictBaseAdd
}

type DistrictBaseRemoveReq struct {
	g.Meta     `path:"/manage/sys/districtBase/remove" tags:"区域管理" method:"post" summary:"区域管理删除接口"`
	DistrictId []uint `json:"district_id" v:"required#请输入区域管理编号"   dc:"区域管理信息"`
}

type DistrictBaseRemoveRes struct {
}

type DistrictBaseListReq struct {
	g.Meta `path:"/manage/sys/districtBase/list" tags:"区域管理" method:"get" summary:"区域管理列表接口"`
	Page   int `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size   int `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`
}

type DistrictBaseListRes struct {
	Items   interface{} `json:"items"    dc:"分类列表"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
type DistrictBaseTreeReq struct {
	g.Meta `path:"/manage/sys/districtBase/tree" tags:"区域管理" method:"get" summary:"获得地址区域"`

	DistrictName string `json:"district_name"      ` // 地区名称
	DistrictBaseAdd
}

type DistrictBaseTreeRes []*model.DistrictTreeNode
