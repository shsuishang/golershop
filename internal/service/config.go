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

	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
)

type (
	IConfigBase interface {
		GetBool(ctx context.Context, id any, defaultValue bool) (res bool)
		GetStr(ctx context.Context, id any, defaultValue string) (res string)
		GetInt(ctx context.Context, id any, defaultValue int) (res int)
		GetInt64(ctx context.Context, id any, defaultValue int64) (res int64)
		GetUint(ctx context.Context, id any, defaultValue uint) (res uint)
		GetFloat(ctx context.Context, id any, defaultValue float64) (res float64)
		// Get 读取属性
		Get(ctx context.Context, id any) (config *entity.ConfigBase, err error)
		// Gets 读取多条属性
		Gets(ctx context.Context, id any) (list []*entity.ConfigBase, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.ConfigBaseListInput) (out []*entity.ConfigBase, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ConfigBaseListInput) (out *do.ConfigBaseListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ConfigBase) (lastInsertId interface{}, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ConfigBase) (affected int64, err error)
		// Save 保存
		Save(ctx context.Context, in *do.ConfigBase) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// GetDefaultImage 系统默认图片
		GetDefaultImage(ctx context.Context) (defaultImage string)
		// getOrderStateList 读取订单状态选项
		GetOrderStateList(ctx context.Context) (optionsList []*model.SelectVo, err error)
		// getNextOrderStateId 读取配置，根据当前orderStateId获得下一状态 sc_order_process
		GetNextOrderStateId(ctx context.Context, orderStateId uint) (nextOrderStateId uint, err error)
		// getAllNextOrderStateId 读取配置，根据当前orderStateId获得下一状态 sc_order_process
		GetAllNextOrderStateId(ctx context.Context, orderStateId uint) (nextOrderStateId uint, err error)
		// GetNextReturnStateId 读取配置，根据当前returnStateId获得下一状态 sc_return_process
		GetNextReturnStateId(ctx context.Context, returnStateId uint) (uint, error)
		// GetPaymentChannelList 取得支付渠道
		GetPaymentChannelList(ctx context.Context) (optionsList []*model.SelectVo, err error)
		// GetPaymentChannelCode 读取配置，根据paymentChannelId读取paymentChannelCode
		GetPaymentChannelCode(ctx context.Context, paymentChannelId uint) (string, error)
		// GetReturnStateList 读取订单状态选项
		GetReturnStateList(ctx context.Context) (optionsList []*model.SelectVo, err error)
		// GetErrStateList 读取错误状态选项
		GetErrStateList(ctx context.Context) (optionsList []*model.SelectVo, err error)
		// GetStateIdList 订单启用状态List
		GetStateIdList(ctx context.Context) (res []uint, err error)
		// Init 初始化
		Init(ctx context.Context) (res bool, err error)
		// IfInvoicing 是否启用进销存管理
		IfInvoicing(ctx context.Context) (res bool)
		// IfPlantformFx 是否启用分销
		IfPlantformFx(ctx context.Context) (res bool)
		// IfSupplierMarket 是否启用供应商市场
		IfSupplierMarket(ctx context.Context) (res bool)
		// GetSiteInfo 初始化
		GetSiteInfo(ctx context.Context, sourceUccCode string) (res map[string]interface{}, err error)
		// CleanCache 清理缓存
		CleanCache(ctx context.Context) (res bool, err error)
		// ifIm
		IfIm(ctx context.Context) (res bool)
		IfB2B(ctx context.Context) (res bool)
		IfChain(ctx context.Context) (res bool)
		IfEdu(ctx context.Context) (res bool)
		IfB2bHall(ctx context.Context) (res bool)
		IfMultilang(ctx context.Context) (res bool)
		IfSns(ctx context.Context) (res bool)
		IfSubsite(ctx context.Context) (res bool)
	}
	IConfigType interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.ConfigTypeListInput) (out []*entity.ConfigType, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ConfigTypeListInput) (out *do.ConfigTypeListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ConfigType) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ConfigType) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
)

var (
	localConfigBase IConfigBase
	localConfigType IConfigType
)

func ConfigBase() IConfigBase {
	if localConfigBase == nil {
		panic("implement not found for interface IConfigBase, forgot register?")
	}
	return localConfigBase
}

func RegisterConfigBase(i IConfigBase) {
	localConfigBase = i
}

func ConfigType() IConfigType {
	if localConfigType == nil {
		panic("implement not found for interface IConfigType, forgot register?")
	}
	return localConfigType
}

func RegisterConfigType(i IConfigType) {
	localConfigType = i
}
