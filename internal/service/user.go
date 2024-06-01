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

	"golershop.cn/api/pt"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
)

type (
	IUserTagGroup interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.UserTagGroupListInput) (out []*entity.UserTagGroup, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserTagGroupListInput) (out *do.UserTagGroupListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserTagGroup) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserTagGroup) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IUserVoucher interface {
		// Get 读取用户优惠券
		Get(ctx context.Context, id any) (out *entity.UserVoucher, err error)
		// Gets 读取多条用户优惠券
		Gets(ctx context.Context, id any) (list []*entity.UserVoucher, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.UserVoucherListInput) (out []*entity.UserVoucher, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserVoucherListInput) (out *do.UserVoucherListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserVoucher) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserVoucher) (affected int64, err error)
		// Edit 编辑
		EditWhere(ctx context.Context, where *do.UserVoucherListInput, in *do.UserVoucher) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// Save 保存
		Save(ctx context.Context, in *do.UserVoucher) (affected int64, err error)
		// List 分页读取
		GetList(ctx context.Context, in *do.UserVoucherListInput) (output *model.UserVoucherListOutput, err error)
	}
	IUserDistribution interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.UserDistributionListInput) (out []*entity.UserDistribution, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserDistributionListInput) (out *do.UserDistributionListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserDistribution) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserDistribution) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// GetList 获取用户分销列表
		GetList(ctx context.Context, input *do.UserDistributionListInput) (*model.UserDistributionOutput, error)
		// GetChildNum 获取子用户数量
		GetChildNum(ctx context.Context, userId uint, startTime, endTime int64) (int, error)
		// InitDistributionUser 添加
		// 添加分销用户记录 - 推广员记录
		InitDistributionUser(ctx context.Context, userParentId uint, userActive bool) bool
		// 添加用户关系
		AddPlantformUser(ctx context.Context, plantformUser *do.UserDistribution) bool
		// 添加分销用户
		AddDistribution(ctx context.Context, userDistribution *do.UserDistribution) bool
	}
	IUserTagBase interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.UserTagBaseListInput) (out []*entity.UserTagBase, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserTagBaseListInput) (out *do.UserTagBaseListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserTagBase) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserTagBase) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IUserInfo interface {
		// Get 根据编号读取活动信息
		Get(ctx context.Context, id any) (out *entity.UserInfo, err error)
		// Gets 根据编号读取读取多条活动信息
		Gets(ctx context.Context, id any) (list []*entity.UserInfo, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.UserInfoListInput) (out []*entity.UserInfo, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserInfoListInput) (out *do.UserInfoListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserInfo) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserInfo) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IUser interface {
		// 登录用户信息
		GetUserInfo(ctx context.Context) (out *model.UserInfoOutput, err error)
		UserEdit(ctx context.Context, userinfo *do.UserInfo) (affected int64, err error)
		// BindMobile 绑定手机号
		BindMobile(ctx context.Context, user *model.ContextUser, userIntl string, mobile uint64, newPassword string) (*model.LoginOutput, error)
		// UnBindMobile 解绑手机号
		UnBindMobile(ctx context.Context, user *model.ContextUser, userIntl string, mobile uint64) (bool, error)
		// ReBindMobile 重新绑定手机号
		ReBindMobile(ctx context.Context, user *model.ContextUser, userIntl string, mobile uint64, newPassword string) (*model.LoginOutput, error)
		// SaveCertificate 保存用户认证信息
		//
		// @param ctx      上下文
		// @param userInfo 用户信息
		// @return
		SaveCertificate(ctx context.Context, userInfo *do.UserInfo) (bool, error)
	}
	IUserAdmin interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.UserAdminListInput) (out []*entity.UserAdmin, err error)
		// FindOne 查询活动数据
		FindOne(ctx context.Context, in *do.UserAdminListInput) (out *entity.UserAdmin, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserAdminListInput) (out *do.UserAdminListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserAdmin) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserAdmin) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IUserMessage interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.UserMessageListInput) (out []*entity.UserMessage, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserMessageListInput) (out *do.UserMessageListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserMessage) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserMessage) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// GetMsgCount 信消息数量
		GetMsgCount(ctx context.Context, in *model.MsgCountInput) (out *model.UserMessageVo, err error)
		// GetList 获取用户消息列表
		GetList(ctx context.Context, input *do.UserMessageListInput) (output *model.UserMessageOutput, err error)
		// getById 获取用户消息详情
		GetById(ctx context.Context, messageId, userId uint) (*entity.UserMessage, error)
		// setRead 设置消息为已读
		SetRead(ctx context.Context, messageId, userOtherId, userId uint) (bool, error)
		// addMessage 添加消息
		AddMessage(ctx context.Context, messageAddInput *model.UserMessageAddInput, userId uint) (*model.UserMessageVo, error)
		// ListChatMsg 读取聊天消息
		ListChatMsg(ctx context.Context, req *do.UserMessageListInput) (messageResPage *model.MessageOutput, err error)
		// SendSysNotice 发送系统通知
		SendSysNotice(ctx context.Context, adminUserId, userId int, msgContent string, msgType int) error
		// 获取通知消息
		GetPlantFromNotice(ctx context.Context, userId uint) (*model.MessageNoticeVo, error)
		// 清除通知消息
		ClearNotice(ctx context.Context, userId uint) (bool, error)
		// 获取消息数量
		GetMessageNum(ctx context.Context, input *do.UserMessageListInput) (*model.UserMessageVo, error)
	}
	IUserProductBrowse interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.UserProductBrowseListInput) (out []*entity.UserProductBrowse, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserProductBrowseListInput) (out *do.UserProductBrowseListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserProductBrowse) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserProductBrowse) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// AddBrowser 添加浏览记录
		AddBrowser(ctx context.Context, itemId uint64, userId uint) (productBrowses []*entity.UserProductBrowse, err error)
		// RemoveBrowser 删除浏览记录
		RemoveBrowser(ctx context.Context, userProductBrowseListReq *pt.UserProductBrowseListReq) (success bool, err error)
	}
	IUserSearchHistory interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.UserSearchHistoryListInput) (out []*entity.UserSearchHistory, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserSearchHistoryListInput) (out *do.UserSearchHistoryListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserSearchHistory) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserSearchHistory) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// GetSearchInfo 返回搜索关键词
		GetSearchInfo(ctx context.Context) (*pt.SearchInfoRes, error)
	}
	IUserDeliveryAddress interface {
		// Get 读取用户地址
		Get(ctx context.Context, id any) (out *entity.UserDeliveryAddress, err error)
		// Gets 读取多条用户地址
		Gets(ctx context.Context, id any) (list []*entity.UserDeliveryAddress, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.UserDeliveryAddressListInput) (out []*entity.UserDeliveryAddress, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserDeliveryAddressListInput) (out *do.UserDeliveryAddressListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserDeliveryAddress) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserDeliveryAddress) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IUserInvoice interface {
		// Get 读取兑换码
		Get(ctx context.Context, id any) (out *entity.UserInvoice, err error)
		// Gets 读取多条兑换码
		Gets(ctx context.Context, id any) (list []*entity.UserInvoice, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.UserInvoiceListInput) (out []*entity.UserInvoice, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserInvoiceListInput) (out *do.UserInvoiceListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserInvoice) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserInvoice) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IUserFavoritesItem interface {
		// Get 读取兑换码
		Get(ctx context.Context, id any) (out *entity.UserFavoritesItem, err error)
		// Gets 读取多条兑换码
		Gets(ctx context.Context, id any) (list []*entity.UserFavoritesItem, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.UserFavoritesItemListInput) (out []*entity.UserFavoritesItem, err error)
		// Find 查询数据
		FindOne(ctx context.Context, in *do.UserFavoritesItemListInput) (out *entity.UserFavoritesItem, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserFavoritesItemListInput) (out *do.UserFavoritesItemListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserFavoritesItem) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserFavoritesItem) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IUserLevel interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.UserLevelListInput) (out []*entity.UserLevel, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserLevelListInput) (out *do.UserLevelListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserLevel) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserLevel) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		GetUserLevelRateMap(ctx context.Context) map[uint]float64
	}
)

var (
	localUserDistribution    IUserDistribution
	localUserTagBase         IUserTagBase
	localUserTagGroup        IUserTagGroup
	localUserVoucher         IUserVoucher
	localUser                IUser
	localUserAdmin           IUserAdmin
	localUserInfo            IUserInfo
	localUserSearchHistory   IUserSearchHistory
	localUserDeliveryAddress IUserDeliveryAddress
	localUserInvoice         IUserInvoice
	localUserMessage         IUserMessage
	localUserProductBrowse   IUserProductBrowse
	localUserFavoritesItem   IUserFavoritesItem
	localUserLevel           IUserLevel
)

func UserDeliveryAddress() IUserDeliveryAddress {
	if localUserDeliveryAddress == nil {
		panic("implement not found for interface IUserDeliveryAddress, forgot register?")
	}
	return localUserDeliveryAddress
}

func RegisterUserDeliveryAddress(i IUserDeliveryAddress) {
	localUserDeliveryAddress = i
}

func UserInvoice() IUserInvoice {
	if localUserInvoice == nil {
		panic("implement not found for interface IUserInvoice, forgot register?")
	}
	return localUserInvoice
}

func RegisterUserInvoice(i IUserInvoice) {
	localUserInvoice = i
}

func UserMessage() IUserMessage {
	if localUserMessage == nil {
		panic("implement not found for interface IUserMessage, forgot register?")
	}
	return localUserMessage
}

func RegisterUserMessage(i IUserMessage) {
	localUserMessage = i
}

func UserProductBrowse() IUserProductBrowse {
	if localUserProductBrowse == nil {
		panic("implement not found for interface IUserProductBrowse, forgot register?")
	}
	return localUserProductBrowse
}

func RegisterUserProductBrowse(i IUserProductBrowse) {
	localUserProductBrowse = i
}

func UserSearchHistory() IUserSearchHistory {
	if localUserSearchHistory == nil {
		panic("implement not found for interface IUserSearchHistory, forgot register?")
	}
	return localUserSearchHistory
}

func RegisterUserSearchHistory(i IUserSearchHistory) {
	localUserSearchHistory = i
}

func UserFavoritesItem() IUserFavoritesItem {
	if localUserFavoritesItem == nil {
		panic("implement not found for interface IUserFavoritesItem, forgot register?")
	}
	return localUserFavoritesItem
}

func RegisterUserFavoritesItem(i IUserFavoritesItem) {
	localUserFavoritesItem = i
}

func UserLevel() IUserLevel {
	if localUserLevel == nil {
		panic("implement not found for interface IUserLevel, forgot register?")
	}
	return localUserLevel
}

func RegisterUserLevel(i IUserLevel) {
	localUserLevel = i
}

func UserDistribution() IUserDistribution {
	if localUserDistribution == nil {
		panic("implement not found for interface IUserDistribution, forgot register?")
	}
	return localUserDistribution
}

func RegisterUserDistribution(i IUserDistribution) {
	localUserDistribution = i
}

func UserTagBase() IUserTagBase {
	if localUserTagBase == nil {
		panic("implement not found for interface IUserTagBase, forgot register?")
	}
	return localUserTagBase
}

func RegisterUserTagBase(i IUserTagBase) {
	localUserTagBase = i
}

func UserTagGroup() IUserTagGroup {
	if localUserTagGroup == nil {
		panic("implement not found for interface IUserTagGroup, forgot register?")
	}
	return localUserTagGroup
}

func RegisterUserTagGroup(i IUserTagGroup) {
	localUserTagGroup = i
}

func UserVoucher() IUserVoucher {
	if localUserVoucher == nil {
		panic("implement not found for interface IUserVoucher, forgot register?")
	}
	return localUserVoucher
}

func RegisterUserVoucher(i IUserVoucher) {
	localUserVoucher = i
}

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("implement not found for interface IUserAdmin, forgot register?")
	}
	return localUserAdmin
}

func RegisterUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement not found for interface IUserInfo, forgot register?")
	}
	return localUserInfo
}

func RegisterUserInfo(i IUserInfo) {
	localUserInfo = i
}
