package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	Material = cMaterial{}
)

type cMaterial struct{}

// =========================== 用户端使用 =============================

// =========================== 管理端使用 =============================

// ---------------------------- 素材分类 -------------------------------
// List 素材分类列表
func (c *cMaterial) ListGallery(ctx context.Context, req *sys.MaterialGalleryListReq) (res *sys.MaterialGalleryListRes, err error) {
	var result, error = service.MaterialGallery().List(ctx, &do.MaterialGalleryListInput{
		BaseList: ml.BaseList{
			Page: req.Page,
			Size: req.Size,
			Sidx: dao.MaterialGallery.Columns().GallerySort,
			Sort: "ASC"},
		Where: do.MaterialGallery{},
	})

	if error != nil {
		err = error
	}

	res = &sys.MaterialGalleryListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 添加素材分类
func (c *cMaterial) AddGallery(ctx context.Context, req *sys.MaterialGalleryAddReq) (res *sys.MaterialGalleryEditRes, err error) {
	input := do.MaterialGallery{}
	gconv.Scan(req, &input)

	var result, error = service.MaterialGallery().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.MaterialGalleryEditRes{
		GalleryId: uint64(result),
	}

	return
}

// Edit 编辑素材分类
func (c *cMaterial) EditGallery(ctx context.Context, req *sys.MaterialGalleryEditReq) (res *sys.MaterialGalleryEditRes, err error) {
	input := do.MaterialGallery{}
	gconv.Scan(req, &input)

	_, error := service.MaterialGallery().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.MaterialGalleryEditRes{
		GalleryId: req.GalleryId,
	}

	return
}

// Remove 删除素材分类
func (c *cMaterial) RemoveGallery(ctx context.Context, req *sys.MaterialGalleryRemoveReq) (res *sys.MaterialGalleryRemoveRes, err error) {
	var _, error = service.MaterialGallery().Remove(ctx, req.GalleryId)

	if error != nil {
		err = error
	}

	res = &sys.MaterialGalleryRemoveRes{}

	return
}

// List 素材项目列表
func (c *cMaterial) ListBase(ctx context.Context, req *sys.MaterialBaseListReq) (res *sys.MaterialBaseListRes, err error) {
	galleryId := req.GalleryId
	materialType := req.MaterialType
	item := do.MaterialBase{GalleryId: galleryId}

	if galleryId == 0 {
		item.MaterialId = nil
	}

	if materialType != "" {
		item.MaterialType = materialType
	}

	var likes []*ml.WhereExt

	if req.MaterialName != "" {
		likes = []*ml.WhereExt{{
			Column: dao.MaterialBase.Columns().MaterialName,
			Val:    "%" + req.MaterialName + "%",
		}}
	}

	var result, error = service.MaterialBase().List(ctx, &do.MaterialBaseListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size:      req.Size,
			WhereLike: likes,
			Sidx:      dao.MaterialBase.Columns().MaterialSort,
			Sort:      "ASC"},
		Where: item,
	})

	if error != nil {
		err = error
	}

	res = &sys.MaterialBaseListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 添加素材项目
func (c *cMaterial) AddBase(ctx context.Context, req *sys.MaterialBaseAddReq) (res *sys.MaterialBaseEditRes, err error) {
	input := do.MaterialBase{}
	gconv.Scan(req, &input)

	var result, error = service.MaterialBase().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.MaterialBaseEditRes{
		MaterialId: uint64(result),
	}

	return
}

// Edit 编辑素材项目
func (c *cMaterial) EditBase(ctx context.Context, req *sys.MaterialBaseEditReq) (res *sys.MaterialBaseEditRes, err error) {
	input := do.MaterialBase{}
	gconv.Scan(req, &input)

	var _, error = service.MaterialBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.MaterialBaseEditRes{
		MaterialId: req.MaterialId,
	}

	return
}

// Remove 删除素材项目
func (c *cMaterial) RemoveBase(ctx context.Context, req *sys.MaterialBaseRemoveReq) (res *sys.MaterialBaseRemoveRes, err error) {
	var _, error = service.MaterialBase().Remove(ctx, req.MaterialId)

	if error != nil {
		err = error
	}

	res = &sys.MaterialBaseRemoveRes{}

	return
}
