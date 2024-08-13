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

package trade

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

type sOrderComment struct{}

func init() {
	service.RegisterOrderComment(NewOrderComment())
}

func NewOrderComment() *sOrderComment {
	return &sOrderComment{}
}

// Get 读取订单
func (s *sOrderComment) Get(ctx context.Context, id any) (out *entity.OrderComment, err error) {
	var list []*entity.OrderComment
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 读取多条订单
func (s *sOrderComment) Gets(ctx context.Context, id any) (list []*entity.OrderComment, err error) {
	err = dao.OrderComment.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sOrderComment) Find(ctx context.Context, in *do.OrderCommentListInput) (out []*entity.OrderComment, err error) {
	out, err = dao.OrderComment.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sOrderComment) List(ctx context.Context, in *do.OrderCommentListInput) (out *do.OrderCommentListOutput, err error) {
	out, err = dao.OrderComment.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sOrderComment) Add(ctx context.Context, in *do.OrderComment) (lastInsertId int64, err error) {
	lastInsertId, err = dao.OrderComment.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sOrderComment) Edit(ctx context.Context, in *do.OrderComment) (affected int64, err error) {
	_, err = dao.OrderComment.Edit(ctx, in.OrderId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sOrderComment) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.OrderComment.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// AddOrderComment 添加订单评论
func (s *sOrderComment) AddOrderComment(ctx context.Context, input *model.OrderCommentInput) error {
	// 获取当前用户
	user := service.BizCtx().GetUser(ctx)
	if user == nil {
		return gerror.New("用户未登录")
	}

	orderId := input.OrderId
	orderBase := input.OrderBase
	commentItemReq := input.OrderCommentItem
	commentImage := input.CommentImage

	// 添加订单评论
	orderComment := &do.OrderComment{}
	gconv.Scan(commentItemReq, orderComment)
	pointsEvaluateGood := service.ConfigBase().GetFloat(ctx, "points_evaluate_good", 0.0)
	orderComment.CommentPoints = pointsEvaluateGood
	orderComment.StoreId = orderBase.StoreId
	orderComment.StoreName = orderBase.StoreName
	orderComment.UserId = user.UserId
	orderComment.UserName = user.UserNickname
	orderComment.CommentContent = orderComment.CommentContent

	if !g.IsEmpty(commentImage) {
		orderComment.CommentImage = gstr.Join(commentImage, ",")
	}

	orderComment.CommentEnable = true
	orderComment.CommentTime = gtime.Now()
	orderComment.CommentStoreDescCredit = input.StoreDesccredit
	orderComment.CommentStoreServiceCredit = input.StoreServicecredit
	orderComment.CommentStoreDeliveryCredit = input.StoreDeliverycredit
	orderComment.CommentScores = (input.StoreDesccredit + input.StoreServicecredit + input.StoreDeliverycredit) / 3

	if _, err := dao.OrderComment.Save(ctx, orderComment); err != nil {
		return gerror.New("订单评论信息保存失败！")
	}

	// 添加店铺分析数据
	/* storeAnalytics, err := dao.StoreAnalytics.Get(ctx, orderBase.StoreId)
	if err != nil {
		return gerror.New("获取店铺分析数据失败！")
	}

	storeAnalytics.StoreDesccredit += input.StoreDesccredit
	storeAnalytics.StoreServicecredit += input.StoreServicecredit
	storeAnalytics.StoreDeliverycredit += input.StoreDeliverycredit
	storeAnalytics.StoreEvaluationNum++

	if _, err := dao.StoreAnalytics.Save(ctx, storeAnalytics); err != nil {
		return gerror.New("店铺评论信息修改失败！")
	} */

	// 修改订单信息
	shopOrderInfo := &do.OrderInfo{
		OrderId:                    orderId,
		OrderBuyerEvaluationStatus: consts.ORDER_EVALUATION_YES,
	}

	if _, err := dao.OrderInfo.Save(ctx, shopOrderInfo); err != nil {
		return gerror.New("订单评论状态修改失败！")
	}

	// 差评提醒
	if commentItemReq.CommentScores == 1 {
		// TODO: 发送站内信
	}

	return nil
}
