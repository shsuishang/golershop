package sys

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	Page = cPage{}
)

type cPage struct{}

// =========================== 管理端使用 =============================

// List 页面项目列表
func (c *cPage) List(ctx context.Context, req *sys.PageBaseListReq) (res *sys.PageBaseListRes, err error) {
	item := do.PageBase{PageType: req.PageType}

	var likes []*ml.WhereExt

	var result, error = service.PageBase().List(ctx, &do.PageBaseListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size:      req.Size,
			WhereLike: likes,
			Sidx:      dao.PageBase.Columns().PageId,
			Sort:      "ASC"},
		Where: item,
	})

	if error != nil {
		err = error
	}

	res = &sys.PageBaseListRes{}
	gconv.Scan(result, &res)

	return
}

// List 页面项目列表
func (c *cPage) ListMobile(ctx context.Context, req *sys.PageMobileListReq) (res *sys.PageMobileListRes, err error) {
	item := do.PageBase{PageType: req.PageType}

	var likes []*ml.WhereExt

	var result, error = service.PageBase().List(ctx, &do.PageBaseListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size:      req.Size,
			WhereLike: likes,
			Sidx:      dao.PageBase.Columns().PageId,
			Sort:      "ASC"},
		Where: item,
	})

	if error != nil {
		err = error
	}

	res = &sys.PageMobileListRes{}
	gconv.Scan(result, &res)

	//用户中心菜单
	menus, err := service.PageBase().GetUserCenterMenu(ctx)

	pageMobileVo := &model.PageMobileVo{}
	pageMobileVo.PageCode = gjson.New(menus).String()
	pageMobileVo.PageId = 100
	pageMobileVo.PageName = "个人中心"
	pageMobileVo.AppId = 8

	res.Center, _ = gjson.EncodeString(pageMobileVo)

	return
}

// Edit 编辑页面项目
func (c *cPage) MobileEdit(ctx context.Context, req *sys.MobileEditReq) (res *sys.MobileEditRes, err error) {
	var pageMobileList []model.PageMobileVo
	gconv.Scan(req.AppPageList, &pageMobileList)

	for _, pageVo := range pageMobileList {
		//fmt.Printf("Index: %d Value: %d\n", index, pageVo)

		input := do.PageBase{}
		gconv.Scan(pageVo, &input)

		input.PageId = pageVo.PageId
		input.PageName = pageVo.PageName
		input.StoreId = pageVo.StoreId
		input.AppId = pageVo.AppId
		input.PageCode = pageVo.PageCode
		input.PageNav = pageVo.PageNav
		input.PageConfig = pageVo.PageConfig
		input.PageShareTitle = pageVo.PageShareTitle
		input.PageShareImage = pageVo.PageShareImage
		input.PageQrcode = pageVo.PageQrcode
		input.PageIndex = pageVo.PageIndex
		input.PageGb = pageVo.PageGb
		input.PageActivity = pageVo.PageActivity
		input.PagePoint = pageVo.PagePoint
		input.PageSns = pageVo.PageSns
		input.PageArticle = pageVo.PageArticle
		input.PageSecondkill = pageVo.PageSecondkill
		input.PageUpgrade = pageVo.PageUpgrade
		input.PageRelease = pageVo.PageRelease

		var _, error = service.PageBase().Edit(ctx, &input)

		if error != nil {
			err = error
		}
	}

	//UsetCenterMenu处理
	//存入config  app_member_center
	configBase := &do.ConfigBase{}
	configBase.ConfigKey = "app_member_center"
	configBase.ConfigValue = req.AppMemberCenter
	configBase.ConfigTypeId = 0
	configBase.ConfigDatatype = "text"
	configBase.ConfigBuildin = true

	service.ConfigBase().Save(ctx, configBase)

	res = &sys.MobileEditRes{}

	return
}

// Edit 编辑页面项目
func (c *cPage) EditBase(ctx context.Context, req *sys.PageBaseEditReq) (res *sys.PageBaseEditRes, err error) {
	input := do.PageBase{}
	gconv.Scan(req, &input)

	var _, error = service.PageBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.PageBaseEditRes{
		PageId: req.PageId,
	}

	return
}
