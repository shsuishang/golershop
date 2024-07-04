package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao/global"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

var (
	PageModule = cPageModule{}
)

type cPageModule struct{}

func (c *cPageModule) List(ctx context.Context, req *sys.PageModuleListReq) (res *sys.PageModuleListRes, err error) {
	input := do.PageModuleListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)
	input.Where.PageId = req.PageId

	var result, error = service.PageModule().GetLists(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.PageModuleListRes{}
	gconv.Scan(result, res)

	return
}

// Add 页面模块表-添加
func (c *cPageModule) Add(ctx context.Context, req *sys.PageModuleAddReq) (res *sys.PageModuleAddRes, err error) {
	// 将请求参数转换为实体对象
	input := &entity.PageModule{}
	gconv.Scan(req, input)

	// 设置默认值
	input.PmName = ""
	input.PmEnable = 0

	// 处理 JSON 字符串中的 URL 和 link
	pmJson := input.PmJson
	url := "\"url\":\"//test.shopsuite.cn"
	urlNew := "\"url\":\"" + global.UrlPc
	link := "\"link\":\"//test.shopsuite.cn"
	linkNew := "\"link\":\"" + global.UrlPc
	if !g.IsEmpty(pmJson) {
		pmJson = gstr.Replace(pmJson, url, urlNew)
		pmJson = gstr.Replace(pmJson, link, linkNew)
		input.PmJson = pmJson
	}

	pagemodule := &do.PageModule{}
	gconv.Scan(input, pagemodule)
	// 调用服务添加页面模块
	_, err = service.PageModule().Add(ctx, pagemodule)

	if err != nil {
		return nil, err
	}

	res = &sys.PageModuleAddRes{}
	gconv.Scan(pagemodule, res)

	return res, nil
}

// Edit 编辑反馈
func (c *cPageModule) Edit(ctx context.Context, req *sys.PageModuleEditReq) (res *sys.PageModuleEditRes, err error) {

	input := do.PageModule{}
	gconv.Scan(req, &input)

	var _, error = service.PageModule().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	return
}

// Remove 删除反馈
func (c *cPageModule) Remove(ctx context.Context, req *sys.PageModuleRemoveReq) (res *sys.PageModuleRemoveRes, err error) {

	var _, error = service.PageModule().Remove(ctx, req.PmId)

	if error != nil {
		err = error
	}

	res = &sys.PageModuleRemoveRes{}

	return
}

// Enable 启用或禁用页面模块
func (c *cPageModule) Enable(ctx context.Context, req *sys.PageModuleEnableReq) (res *sys.PageModuleEnableRes, err error) {
	// 检查 usable 参数是否非空
	if g.IsEmpty(req.PmEnable) {
		return nil, nil
	}

	// 构造 PageModule 对象
	pageModule := &do.PageModule{
		PmId:     req.PmId,
		PageId:   req.PageId,
		PmEnable: gconv.Bool(req.PmEnable == "usable"),
	}

	// 调用服务层的 edit 方法
	success, err := service.PageModule().Edit(ctx, pageModule)
	if err != nil {
		return nil, err
	}

	// 返回结果
	if success == 0 {
		return
	}

	return nil, gerror.New("操作失败")
}

// GetModuleTpl 获取PC楼层模板
func (c *cPageModule) GetModuleTpl(ctx context.Context, req *sys.PageModuleReq) (res *sys.PageModuleRes, err error) {
	// 调用 service 获取模块模板
	res = &sys.PageModuleRes{}
	moduleTpl, err := service.PageModule().GetModuleTpl(ctx)
	if err != nil {
		return nil, gerror.New("获取PC楼层模板失败: " + err.Error())
	}

	// 检查 moduleTpl 是否为空
	if !g.IsEmpty(moduleTpl) {
		gconv.Scan(moduleTpl, res)
		return res, nil
	}

	return nil, gerror.New("获取PC楼层模板失败: 模板为空")
}

// Sort 页面模块表-拖动
func (c *cPageModule) Sort(ctx context.Context, req *sys.PageModuleSortReq) (res *sys.PageModuleSortRes, err error) {
	// 获取并处理pmIds
	pmIds := gconv.SliceUint(gstr.Split(req.PmIdString, ","))

	// 获取页面模块列表
	pageModuleList, err := service.PageModule().Gets(ctx, pmIds)
	if err != nil {
		return nil, err
	}

	if !g.IsEmpty(pageModuleList) {
		// 设置排序
		for i, pageModule := range pageModuleList {
			pageModule.PmOrder = uint(i)
			input := &do.PageModule{}
			gconv.Scan(pageModule, input)
			_, err = service.PageModule().Edit(ctx, input)
			if err != nil {
				return nil, gerror.New("操作失败")
			}
		}
	}

	return nil, err
}
