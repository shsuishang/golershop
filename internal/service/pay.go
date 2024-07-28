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

	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/shopspring/decimal"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
)

type (
	IConsumeRecord interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.ConsumeRecordListInput) (out []*entity.ConsumeRecord, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ConsumeRecordListInput) (out *do.ConsumeRecordListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ConsumeRecord) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ConsumeRecord) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		GetList(ctx context.Context, in *do.ConsumeRecordListInput) (out *model.ConsumeRecordOutput, err error)
	}
	IConsumeReturn interface {
		// DoRefund 执行退款操作
		DoRefund(ctx context.Context, orderReturns []*entity.OrderReturn) bool
		// SetReturnPaidYes 修改为退款已支付状态
		SetReturnPaidYes(ctx context.Context, returnIds []string) (bool, error)
		// DoRefundOrder 操作退款数据
		DoRefundOrder(ctx context.Context, orderRefundFlag bool, userId, storeId uint, userResource *entity.UserResource, orderId string, buyerUserMoney, buyerUserPoints decimal.Decimal, returnId string) (bool, error)
		// doOnLineRefund 执行线上支付退款
		DoOnlineRefund(ctx context.Context, returnId string) error
	}
	IConsumeTrade interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.ConsumeTradeListInput) (out []*entity.ConsumeTrade, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ConsumeTradeListInput) (out *do.ConsumeTradeListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ConsumeTrade) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ConsumeTrade) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// ProcessPay 余额支付
		ProcessPay(ctx context.Context, ids string, deposit model.PayMetVo) (out model.ProcessPayOutput, err error)
	}
	IPaymentAlipay interface {
		// GetClient 初始化支付宝客户端并做配置
		GetClient(ctx context.Context, tradeNo string) (client *alipay.Client, err error)
	}
	IConsumeDeposit interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.ConsumeDepositListInput) (out []*entity.ConsumeDeposit, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ConsumeDepositListInput) (out *do.ConsumeDepositListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ConsumeDeposit) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// ProcessDeposit 新增
		ProcessDeposit(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error)
		// GetTradeNo 获取交易编号
		GetTradeNo(ctx context.Context, orderIdList []string) (tradeNo string, err error)
		// GetOrderId 根据交易编号 获取 订单编号
		GetOrderId(ctx context.Context, tradeNo string) (orderIds string, err error)
		// WechatApplet 微信小程序
		WechatApplet(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error)
		// WechatApp App
		WechatApp(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error)
		// WechatJSAPI 微信JSAPI
		WechatJSAPI(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error)
		// OfflinePay 离线支付
		OfflinePay(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error)
		// 获取业务返回数据
		GetPayResult(ctx context.Context, in *model.PaymentInput) (*model.PaymentOutput, error)
	}
	IPaymentWechat interface {
		// GetClient 初始化微信v3客户端并做配置
		GetClient(ctx context.Context) (client *wechat.ClientV3, err error)
	}
	IUserPay interface {
		// Get 读取信息
		Get(ctx context.Context, id any) (out *entity.UserPay, err error)
		// Gets 读取多条信息
		Gets(ctx context.Context, id any) (list []*entity.UserPay, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.UserPayListInput) (out []*entity.UserPay, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserPayListInput) (out *do.UserPayListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserPay) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserPay) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// GetPayPasswd 获取支付密码
		GetPayPasswd(ctx context.Context, userId uint) (*entity.UserPay, error)
		// ChangePayPassword 修改支付密码
		ChangePayPassword(ctx context.Context, oldPayPassword, newPayPassword, payPassword string, userId uint) (bool, error)
	}
	IUserPointsHistory interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.UserPointsHistoryListInput) (out []*entity.UserPointsHistory, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserPointsHistoryListInput) (out *do.UserPointsHistoryListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserPointsHistory) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserPointsHistory) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		GetList(ctx context.Context, in *do.UserPointsHistoryListInput) (out *model.UserPointsHistoryOutput, err error)
	}
	IUserResource interface {
		// Get 读取信息
		Get(ctx context.Context, id any) (out *entity.UserResource, err error)
		// Gets 读取多条信息
		Gets(ctx context.Context, id any) (list []*entity.UserResource, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.UserResourceListInput) (out []*entity.UserResource, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserResourceListInput) (out *do.UserResourceListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserResource) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserResource) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// InitUserPoints 初始化用户积分
		InitUserPoints(ctx context.Context, userId uint) error
		// InitUserExperience 初始化用户经验等级
		InitUserExperience(ctx context.Context, userId uint) error
		// Experience 操作用户经验
		Experience(ctx context.Context, experienceVo *model.ExperienceVo) error
		// Points 积分操作
		Points(ctx context.Context, vo *model.UserPointsVo) (flag bool, err error)
		// 获取签到基本信息
		GetSignInfo(ctx context.Context, userId uint) (res *model.SignInfoOutput, err error)
		// 签到
		SignIn(ctx context.Context, userId uint) (flag bool, err error)
		// GetSignState 当天是否签到
		GetSignState(ctx context.Context, userId uint) (flag bool, err error)
		GetList(ctx context.Context, in *do.UserResourceListInput) (out *model.UserResourceOutput, err error)
	}
)

var (
	localUserPay           IUserPay
	localUserPointsHistory IUserPointsHistory
	localUserResource      IUserResource
	localConsumeDeposit    IConsumeDeposit
	localPaymentWechat     IPaymentWechat
	localConsumeTrade      IConsumeTrade
	localPaymentAlipay     IPaymentAlipay
	localConsumeRecord     IConsumeRecord
	localConsumeReturn     IConsumeReturn
)

func PaymentAlipay() IPaymentAlipay {
	if localPaymentAlipay == nil {
		panic("implement not found for interface IPaymentAlipay, forgot register?")
	}
	return localPaymentAlipay
}

func RegisterPaymentAlipay(i IPaymentAlipay) {
	localPaymentAlipay = i
}

func ConsumeRecord() IConsumeRecord {
	if localConsumeRecord == nil {
		panic("implement not found for interface IConsumeRecord, forgot register?")
	}
	return localConsumeRecord
}

func RegisterConsumeRecord(i IConsumeRecord) {
	localConsumeRecord = i
}

func ConsumeReturn() IConsumeReturn {
	if localConsumeReturn == nil {
		panic("implement not found for interface IConsumeReturn, forgot register?")
	}
	return localConsumeReturn
}

func RegisterConsumeReturn(i IConsumeReturn) {
	localConsumeReturn = i
}

func ConsumeTrade() IConsumeTrade {
	if localConsumeTrade == nil {
		panic("implement not found for interface IConsumeTrade, forgot register?")
	}
	return localConsumeTrade
}

func RegisterConsumeTrade(i IConsumeTrade) {
	localConsumeTrade = i
}

func UserPointsHistory() IUserPointsHistory {
	if localUserPointsHistory == nil {
		panic("implement not found for interface IUserPointsHistory, forgot register?")
	}
	return localUserPointsHistory
}

func RegisterUserPointsHistory(i IUserPointsHistory) {
	localUserPointsHistory = i
}

func UserResource() IUserResource {
	if localUserResource == nil {
		panic("implement not found for interface IUserResource, forgot register?")
	}
	return localUserResource
}

func RegisterUserResource(i IUserResource) {
	localUserResource = i
}

func ConsumeDeposit() IConsumeDeposit {
	if localConsumeDeposit == nil {
		panic("implement not found for interface IConsumeDeposit, forgot register?")
	}
	return localConsumeDeposit
}

func RegisterConsumeDeposit(i IConsumeDeposit) {
	localConsumeDeposit = i
}

func PaymentWechat() IPaymentWechat {
	if localPaymentWechat == nil {
		panic("implement not found for interface IPaymentWechat, forgot register?")
	}
	return localPaymentWechat
}

func RegisterPaymentWechat(i IPaymentWechat) {
	localPaymentWechat = i
}

func UserPay() IUserPay {
	if localUserPay == nil {
		panic("implement not found for interface IUserPay, forgot register?")
	}
	return localUserPay
}

func RegisterUserPay(i IUserPay) {
	localUserPay = i
}
