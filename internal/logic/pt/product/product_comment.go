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

package product

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/trade"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"strings"
	"time"
)

type sProductComment struct{}

func init() {
	service.RegisterProductComment(NewProductComment())
}

func NewProductComment() *sProductComment {
	return &sProductComment{}
}

// Find 查询数据
func (s *sProductComment) Find(ctx context.Context, in *do.ProductCommentListInput) (out []*entity.ProductComment, err error) {
	out, err = dao.ProductComment.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sProductComment) List(ctx context.Context, in *do.ProductCommentListInput) (out *do.ProductCommentListOutput, err error) {
	out, err = dao.ProductComment.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sProductComment) Add(ctx context.Context, in *do.ProductComment) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ProductComment.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sProductComment) Edit(ctx context.Context, in *do.ProductComment) (affected int64, err error) {
	_, err = dao.ProductComment.Edit(ctx, in.CommentId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sProductComment) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.ProductComment.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetList 获取商品评论列表
func (s *sProductComment) GetList(ctx context.Context, productCommentListReq *do.ProductCommentListInput) (productCommentPage *do.ProductCommentListOutput, err error) {
	productCommentPage, err = s.List(ctx, productCommentListReq)
	if err != nil {
		return nil, err
	}

	if !g.IsEmpty(productCommentPage.Items) {
		commentList := productCommentPage.Items
		productIds := array.Column(commentList, "ProductId")
		productBases, err := dao.ProductBase.Gets(ctx, productIds)
		if err != nil {
			return nil, err
		}

		productNameMap := make(map[uint64]string)
		if !g.IsEmpty(productBases) {
			for _, productBase := range productBases {
				productNameMap[productBase.ProductId] = productBase.ProductName
			}
		}

		for i := range commentList {
			commentImage := commentList[i].CommentImage
			if !g.IsEmpty(commentImage) {
				commentList[i].CommentImages = gconv.Strings(commentImage)
			}

			// 产品名称
			if !g.IsEmpty(productNameMap) {
				commentList[i].ProductName = productNameMap[commentList[i].ProductId]
			}
		}
	}

	return productCommentPage, nil
}

// StoreEvaluationWithContent 读取订单商品评价
func (s *sProductComment) StoreEvaluationWithContent(ctx context.Context, orderId string) (out *model.OrderCommentOutput, err error) {
	// 获取当前登录用户ID
	userId := service.BizCtx().GetUser(ctx).UserId

	/*	// 获取订单信息
		orderInfo, err := dao.OrderInfo.Get(ctx, orderId)
		if err != nil {
			return nil, err
		}

		// 检查数据权限
		if !CheckDataRights(userId, orderInfo.UserId) {
			return nil, gerror.New("无权限访问")
		}
	*/
	// 订单评价状态
	orderItemEvaluationStatus := []uint{consts.ORDER_ITEM_EVALUATION_NO, consts.ORDER_ITEM_EVALUATION_YES}

	// 获取评价项目
	commentRes, err := service.OrderBase().GetEvaluationItem(ctx, &model.EvaluationVo{
		UserId:                    userId,
		OrderItemEvaluationStatus: orderItemEvaluationStatus,
		OrderId:                   orderId,
	})
	if err != nil {
		return nil, err
	}

	// 查询产品评论
	commentList, err := dao.ProductComment.Find(ctx, &do.ProductCommentListInput{
		Where: do.ProductComment{
			OrderId: orderId,
			UserId:  userId,
		},
	})
	if err != nil {
		return nil, err
	}

	// 遍历订单项并匹配评论
	orderItems := commentRes.Items
	for i := range orderItems {
		itemId := orderItems[i].ItemId
		for _, productComment := range commentList {
			if productComment.ItemId == itemId {
				gconv.Scan(productComment, &orderItems[i])

				// 评论图片
				if productComment.CommentImage != "" {
					orderItems[i].CommentImage = gstr.Split(productComment.CommentImage, ",")
				}

				// 评论回复
				replyList, err := dao.ProductCommentReply.Find(ctx, &do.ProductCommentReplyListInput{
					Where: do.ProductCommentReply{
						CommentId:          productComment.CommentId,
						UserIdTo:           userId,
						CommentReplyEnable: true,
					},
					BaseList: ml.BaseList{},
				})
				if err != nil {
					return nil, err
				}

				if len(replyList) > 0 {
					orderItems[i].ProductCommentReplyList = replyList
				}
			}
		}
	}

	// 查询订单评价信息
	orderComment, err := dao.OrderComment.Get(ctx, orderId)
	if err != nil {
		return nil, err
	}

	if orderComment != nil {
		var orderCommentVo model.OrderCommentVo
		gconv.Scan(orderComment, &orderCommentVo)

		if orderCommentVo.CommentImage != "" {
			orderCommentVo.CommentImages = gstr.Split(orderCommentVo.CommentImage, ",")
		}

		commentRes.OrderEvaluation = orderCommentVo
	} else {
		commentRes.OrderEvaluation = model.OrderCommentVo{}
	}

	// 返回结果
	return commentRes, nil
}

// AddOrderComment 添加订单评论
func (s *sProductComment) AddOrderComment(ctx context.Context, commentReq *trade.OrderCommentReq) error {
	// 获取当前用户
	user := service.BizCtx().GetUser(ctx)
	if user == nil {
		return gerror.New("用户未登录")
	}

	// 检查该订单是否已评论
	orderComment, err := dao.OrderComment.Get(ctx, commentReq.OrderId)
	if err == nil && orderComment != nil {
		return gerror.New("该订单已评论！")
	}

	orderId := commentReq.OrderId
	orderBase, err := dao.OrderBase.Get(ctx, orderId)
	if err != nil {
		return err
	}

	orderInfo, err := dao.OrderInfo.Get(ctx, orderId)
	if err != nil {
		return err
	}

	itemArray := gjson.New(commentReq.Item)
	orderCommentReqItems := make([]*model.OrderCommentItemVo, 0)
	err = itemArray.Scan(&orderCommentReqItems)
	if err != nil {
		return err
	}

	for _, commentItemReq := range orderCommentReqItems {
		comment := new(entity.ProductComment)
		gconv.Struct(commentItemReq, comment)

		if commentItemReq.CommentIsAnonymous {
			comment.UserName = fmt.Sprintf("匿名用户%d", time.Now().Unix())
			comment.CommentIsAnonymous = true
		} else {
			comment.UserName = user.UserNickname
			comment.CommentIsAnonymous = false
		}

		/*		// 提交时候，判断没有违禁词才可以通行
				commentContent := commentItemReq.CommentContent
				if service.filterKeyword.HasKeyword(commentContent) {
					return gerror.New("评论中包含非法词汇！")
				}*/

		comment.UserId = user.UserId
		comment.CommentEnable = true
		comment.StoreId = orderBase.StoreId
		comment.StoreName = orderBase.StoreName
		comment.ChainId = orderInfo.ChainId
		comment.CommentScores = commentItemReq.CommentScores

		if !g.IsEmpty(commentItemReq.CommentImage) {
			comment.CommentImage = strings.Join(commentItemReq.CommentImage, ",")
		} else {
			comment.CommentImage = ""
		}

		// 商品评论信息
		if err := s.AddProductComment(ctx, comment, commentItemReq.OrderItemId); err != nil {
			return err
		}

		commentInput := &model.OrderCommentInput{
			OrderId:             orderId,
			OrderBase:           orderBase,
			OrderCommentItem:    commentItemReq,
			CommentImage:        commentItemReq.CommentImage,
			StoreDesccredit:     commentReq.StoreDesccredit,
			StoreServicecredit:  commentReq.StoreServicecredit,
			StoreDeliverycredit: commentReq.StoreDeliverycredit,
		}

		// 添加订单评论
		err := service.OrderComment().AddOrderComment(ctx, commentInput)
		if err != nil {
			return err
		}
	}

	return nil
}

// addProductComment 添加商品评论
func (s *sProductComment) AddProductComment(ctx context.Context, comment *entity.ProductComment, orderItemId int) (err error) {
	orderId := comment.OrderId

	var orderItems []*entity.OrderItem
	// 根据 orderItemId 获取订单商品
	if !g.IsEmpty(orderItemId) {
		orderItems, err = dao.OrderItem.Gets(ctx, orderItemId)
	} else {
		orderItems, err = dao.OrderItem.Find(ctx, &do.OrderItemListInput{
			Where: do.OrderItem{
				OrderId: orderId,
				ItemId:  comment.ItemId,
			},
		})
	}

	if err != nil || len(orderItems) == 0 {
		return gerror.New("保存评论失败！")
	}

	// 获取第一个订单商品
	orderItem := orderItems[0]
	productId := orderItem.ProductId
	comment.ProductId = productId
	comment.ItemId = orderItem.ItemId
	comment.ItemName = orderItem.ItemName

	// 获取商品基本信息
	productBase, err := dao.ProductBase.Get(ctx, productId)
	if err != nil || productBase == nil {
		return gerror.New("商品已不存在，不可评论！")
	}

	// 检查订单商品的评价状态
	orderItemEvaluationStatus := orderItem.OrderItemEvaluationStatus
	if orderItemEvaluationStatus == 0 {
		// 获得积分 经验
		pointsEvaluateGood := service.ConfigBase().GetFloat(ctx, "points_evaluate_good", 0)
		//expEvaluateGood := service.ConfigBase().GetFloat(ctx, "exp_evaluate_good", 0)

		// 添加评论积分
		comment.CommentPoints = pointsEvaluateGood

		// 分站
		orderInfo, err := dao.OrderInfo.Get(ctx, orderId)
		if err == nil {
			comment.SubsiteId = orderInfo.SubsiteId
		}

		// 添加评论
		input := &do.ProductComment{}
		gconv.Scan(comment, input)
		_, err = dao.ProductComment.Add(ctx, input)
		if err != nil {
			return gerror.New("保存评论失败！")
		}

		// 更新订单商品评价状态
		itemQueryWrapper := &do.OrderItem{
			OrderItemId: array.Column(orderItems, "OrderItemId"),
		}
		shopOrderItem := &do.OrderItem{
			OrderItemEvaluationStatus: true,
		}

		_, err = dao.OrderItem.Edit(ctx, itemQueryWrapper, shopOrderItem)
		if err != nil {
			return err
		}

		// 更新产品评论数
		shopProductIndex, err := dao.ProductIndex.Get(ctx, productId)
		if err == nil {
			shopProductIndex.ProductEvaluationNum += 1
			in := &do.ProductIndex{}
			gconv.Scan(shopProductIndex, in)
			_, err := dao.ProductIndex.Edit(ctx, shopProductIndex.ProductId, in)
			if err != nil {
				return gerror.New("产品评论更新失败！")
			}

		}

		// 触发评论完成时间
		// 差评提醒
		//if comment.CommentScores == 1 {
		//	messageId := "bad-review-reminder-notification"
		//	args := map[string]interface{}{
		//		"order_id": orderId,
		//	}
		//	adminUserId := service.ConfigBase().GetInt(ctx, "message_notice_user_id", 10001)
		//	service.Message().SendNoticeMsg(ctx, adminUserId, messageId, args)
		//}
	} else {
		return gerror.New("该订单商品已进行评价！")
	}

	return nil
}
