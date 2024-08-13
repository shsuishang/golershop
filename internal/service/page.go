// +----------------------------------------------------------------------
// | ShopSuite商城系统 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 随商信息技术（上海）有限公司
// +----------------------------------------------------------------------
// | 未获商业授权前，不得将本软件用于商业用途。禁止整体或任何部分基础上以发展任何派生版本、
// | 修改版本或第三方版本用于重新分发。
// +----------------------------------------------------------------------
// | 官方网站: https://www.shopsuite.cn  https://www.golershop.cn
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本公司对该软件产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细见https://www.golershop.cn/policy
// +----------------------------------------------------------------------

package service

import (
	"context"

	"golershop.cn/api/sys"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
)

type (
	IPageBase interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.PageBaseListInput) (out []*entity.PageBase, err error)
		// 查询主键
		FindKey(ctx context.Context, in *do.PageBaseListInput) (out []interface{}, err error)
		// List 分页读取
		List(ctx context.Context, in *do.PageBaseListInput) (out *do.PageBaseListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.PageBase) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.PageBase) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// GetUserCenterMenu
		GetUserCenterMenu(ctx context.Context) (out map[string]interface{}, err error)
		// Detail 获取页面详情
		Detail(ctx context.Context, pageId any) (out *model.PageDetail, err error)
		// GetDataInfo
		GetDataInfo(ctx context.Context, pageDataReq *sys.PageBaseGetDataInfoReq) (*sys.PageBaseGetDataInfoRes, error)
	}
	IPageModule interface {
		Get(ctx context.Context, id any) (out *entity.PageModule, err error)
		Gets(ctx context.Context, id any) (list []*entity.PageModule, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.PageModuleListInput) (out []*entity.PageModule, err error)
		// List 分页读取
		List(ctx context.Context, in *do.PageModuleListInput) (out *do.PageModuleListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.PageModule) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.PageModule) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// FixPcPageModuleData
		FixPcPageModuleData(ctx context.Context, pageData []*entity.PageModule) ([]map[string]interface{}, error)
		// GetModuleTpl
		GetModuleTpl(ctx context.Context) (map[string]interface{}, error)
		// GetLists
		GetLists(ctx context.Context, req *do.PageModuleListInput) (pageModuleVo *model.PageModuleVoOutput, err error)
	}
	IPageCategoryNav interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.PageCategoryNavListInput) (out []*entity.PageCategoryNav, err error)
		// List 分页读取
		List(ctx context.Context, in *do.PageCategoryNavListInput) (out *do.PageCategoryNavListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.PageCategoryNav) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.PageCategoryNav) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// GetPcLayout
		GetPcLayout(ctx context.Context) (resultSlice []interface{}, err error)
	}
	IPagePcNav interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.PagePcNavListInput) (out []*entity.PagePcNav, err error)
		// List 分页读取
		List(ctx context.Context, in *do.PagePcNavListInput) (out *do.PagePcNavListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.PagePcNav) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.PagePcNav) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
)

var (
	localPageBase        IPageBase
	localPageModule      IPageModule
	localPageCategoryNav IPageCategoryNav
	localPagePcNav       IPagePcNav
)

func PageBase() IPageBase {
	if localPageBase == nil {
		panic("implement not found for interface IPageBase, forgot register?")
	}
	return localPageBase
}

func RegisterPageBase(i IPageBase) {
	localPageBase = i
}

func PageModule() IPageModule {
	if localPageModule == nil {
		panic("implement not found for interface IPageModule, forgot register?")
	}
	return localPageModule
}

func RegisterPageModule(i IPageModule) {
	localPageModule = i
}
func PageCategoryNav() IPageCategoryNav {
	if localPageCategoryNav == nil {
		panic("implement not found for interface IPageCategoryNav, forgot register?")
	}
	return localPageCategoryNav
}

func RegisterPageCategoryNav(i IPageCategoryNav) {
	localPageCategoryNav = i
}
func PagePcNav() IPagePcNav {
	if localPagePcNav == nil {
		panic("implement not found for interface IPagePcNav, forgot register?")
	}
	return localPagePcNav
}

func RegisterPagePcNav(i IPagePcNav) {
	localPagePcNav = i
}
