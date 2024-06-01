package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

type UploadReq struct {
	g.Meta       `path:"/front/sys/upload/index" tags:"upload" method:"post" summary:"上传" dc:"上传结果"`
	GalleryId    uint64 `json:"gallery_id"         ` // 分类编号
	MaterialType string `json:"material_type"      ` // 素材类型
}

type UploadRes struct {
	model.FileInfo
}

type UploadImageReq struct {
	g.Meta       `path:"/front/sys/upload/image" tags:"upload" method:"post" summary:"上传" dc:"上传结果"`
	GalleryId    uint64 `json:"gallery_id"         ` // 分类编号
	MaterialType string `json:"material_type"      ` // 素材类型
}

type UploadVideoReq struct {
	g.Meta       `path:"/front/sys/upload/video" tags:"upload" method:"post" summary:"上传" dc:"上传结果"`
	GalleryId    uint64 `json:"gallery_id"         ` // 分类编号
	MaterialType string `json:"material_type"      ` // 素材类型
}

type UploadFileReq struct {
	g.Meta       `path:"/front/sys/upload/file" tags:"upload" method:"post" summary:"上传" dc:"上传结果"`
	GalleryId    uint64 `json:"gallery_id"         ` // 分类编号
	MaterialType string `json:"material_type"      ` // 素材类型
}
