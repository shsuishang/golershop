package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

var (
	Feedback = cFeedback{}
)

type cFeedback struct{}

func (c *cFeedback) GetCategory(ctx context.Context, req *sys.FeedbackReq) (res []sys.FeedbackRes, err error) {

	var feedbackTypes *do.FeedbackTypeListInput
	// 查询所有启用的反馈类型
	feedbackTypes = &do.FeedbackTypeListInput{
		Where: do.FeedbackType{
			FeedbackTypeEnable: 1,
		}}

	var result, _ = service.FeedbackType().Find(ctx, feedbackTypes)

	// 提取反馈类型ID
	typeIDs := make([]uint, 0, len(result))
	for _, ft := range result {
		typeIDs = append(typeIDs, ft.FeedbackTypeId)
	}

	// 查询与这些类型ID关联的启用的反馈分类
	var feedbackCategories *do.FeedbackCategoryListInput

	feedbackCategories = &do.FeedbackCategoryListInput{
		Where: do.FeedbackCategory{
			FeedbackTypeId:         typeIDs,
			FeedbackCategoryEnable: 1,
		}}
	var CategoryList, _ = service.FeedbackCategory().Find(ctx, feedbackCategories)

	// 构建FeedbackRes列表
	feedbackResList := sys.FeedbackRes{}
	for _, backType := range result {
		feedbackResList.FeedbackType = backType
		Rows := make([]*entity.FeedbackCategory, 0)
		for _, Category := range CategoryList {
			if Category.FeedbackTypeId == backType.FeedbackTypeId {
				Rows = append(Rows, Category)
			}
			feedbackResList.Rows = Rows
		}
		res = append(res, feedbackResList)
	}

	return
}

func (c *cFeedback) List(ctx context.Context, req *sys.BaseListReq) (res *sys.BaseListRes, err error) {

	input := do.FeedbackBaseListInput{}
	gconv.Scan(req, &input)

	var result, error = service.FeedbackBase().List(ctx, &do.FeedbackBaseListInput{
		BaseList: ml.BaseList{
			Page: req.Page,
			Size: req.Size,
			Sidx: dao.FeedbackBase.Columns().FeedbackId,
			Sort: "DESC",
		},
	})

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

func (c *cFeedback) Add(ctx context.Context, req *sys.BaseAddReq) (res *sys.BaseEditRes, err error) {

	input := do.FeedbackBase{}
	gconv.Scan(req, &input)

	var result, error = service.FeedbackBase().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.BaseEditRes{
		FeedbackId: uint(result),
	}

	return
}
