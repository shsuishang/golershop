package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// =========================== 用户端使用 =============================

// =========================== 管理端使用 =============================

type MaterialGalleryAdd struct {
	GalleryName string `json:"gallery_name"       ` // 分类名
	GallerySort uint   `json:"gallery_sort"       ` // 分类排序
}

type MaterialGalleryAddReq struct {
	g.Meta `path:"/manage/sys/material/addGallery" tags:"素材管理" method:"post" summary:"素材类型详情接口"`

	MaterialGalleryAdd
}

type MaterialGalleryEditReq struct {
	g.Meta    `path:"/manage/sys/material/editGallery" tags:"素材管理" method:"post" summary:"素材类型详情接口"`
	GalleryId uint64 `json:"gallery_id" v:"required#请输入素材分类编号"   dc:"素材类型信息"`
	MaterialGalleryAdd
}

type MaterialGalleryEditRes struct {
	GalleryId uint64 `json:"gallery_id"   dc:"主键编号"   `
}

type MaterialGalleryRemoveReq struct {
	g.Meta    `path:"/manage/sys/material/removeGallery" tags:"素材管理" method:"post" summary:"素材类型删除接口"`
	GalleryId []uint `json:"gallery_id" v:"required#请输入素材类型编号"   dc:"素材类型信息"`
}

type MaterialGalleryRemoveRes struct {
}

type MaterialGalleryListReq struct {
	g.Meta `path:"/manage/sys/material/listGallery" tags:"素材管理" method:"get" summary:"素材类型列表接口"`
	Page   int `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size   int `json:"size" d:"10" v:"max:5000#分页数量最大5000条"  dc:"分页数量"`
}

type MaterialGalleryListRes struct {
	Items   interface{} `json:"items"    dc:"素材分类列表页"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

// ---------------------------- 素材项 -------------------------------

type MaterialBaseAdd struct {
	GalleryId    uint64 `json:"gallery_id"         ` // 分类编号
	StoreId      uint   `json:"store_id"           ` // 店铺编号
	MaterialUrl  string `json:"material_url"       ` // 文件URL
	MaterialSort uint   `json:"material_sort"      ` // 素材排序
	MaterialType string `json:"material_type"      ` // 素材类型
	MaterialName string `json:"material_name"      ` // 素材标题
	MaterialDesc string `json:"material_desc"      ` // 素材描述
}

type MaterialBaseAddReq struct {
	g.Meta `path:"/manage/sys/material/add" tags:"素材管理" method:"post" summary:"素材项目详情接口"`

	MaterialBaseAdd
}

type MaterialBaseEditReq struct {
	g.Meta `path:"/manage/sys/material/edit" tags:"素材管理" method:"post" summary:"素材项目详情接口"`

	MaterialId uint64 `json:"material_id"        ` // 素材编号
	MaterialBaseAdd
}

type MaterialBaseEditRes struct {
	MaterialId uint64 `json:"material_id"        ` // 素材编号 `
}

type MaterialBaseRemoveReq struct {
	g.Meta     `path:"/manage/sys/material/remove" tags:"素材管理" method:"post" summary:"素材项目删除接口"`
	MaterialId []uint64 `json:"material_id" v:"required#请输入素材编码"   dc:"素材编码"`
}

type MaterialBaseRemoveRes struct {
}

type MaterialBaseListReq struct {
	g.Meta       `path:"/manage/sys/material/list" tags:"素材管理" method:"get" summary:"素材项目列表接口"`
	Page         int    `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size         int    `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`
	GalleryId    uint   `json:"gallery_id" d:"0" v:"required#请输入素材类型编号"   dc:"素材类型信息"`
	MaterialType string `json:"material_type"      ` // 素材类型
	MaterialName string `json:"material_name"      ` // 素材标题
}

type MaterialBaseListRes struct {
	Items   interface{} `json:"items"    dc:"素材列表页"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
