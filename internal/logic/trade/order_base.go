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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
)

type sOrderBase struct{}

func init() {
	service.RegisterOrderBase(NewOrderBase())
}

func NewOrderBase() *sOrderBase {
	return &sOrderBase{}
}

// Get 读取订单
func (s *sOrderBase) Get(ctx context.Context, id any) (out *entity.OrderBase, err error) {
	var list []*entity.OrderBase
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
func (s *sOrderBase) Gets(ctx context.Context, id any) (list []*entity.OrderBase, err error) {
	err = dao.OrderBase.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sOrderBase) Find(ctx context.Context, in *do.OrderBaseListInput) (out []*entity.OrderBase, err error) {
	out, err = dao.OrderBase.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sOrderBase) List(ctx context.Context, in *do.OrderBaseListInput) (out *do.OrderBaseListOutput, err error) {
	out, err = dao.OrderBase.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sOrderBase) Add(ctx context.Context, in *do.OrderBase) (lastInsertId int64, err error) {
	lastInsertId, err = dao.OrderBase.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sOrderBase) Edit(ctx context.Context, in *do.OrderBase) (affected int64, err error) {
	_, err = dao.OrderBase.Edit(ctx, in.OrderId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sOrderBase) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.OrderBase.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetEvaluationItem 获取订单评价项
func (s *sOrderBase) GetEvaluationItem(ctx context.Context, evaluationVo *model.EvaluationVo) (comment *model.OrderCommentOutput, err error) {
	comment = &model.OrderCommentOutput{}

	orderId := evaluationVo.OrderId
	userId := evaluationVo.UserId

	input := &do.OrderItemListInput{
		Where: do.OrderItem{
			UserId: userId,
		},
	}
	input.WhereExt = []*ml.WhereExt{{
		Column: dao.OrderItem.Columns().OrderItemEvaluationStatus,
		Val:    evaluationVo.OrderItemEvaluationStatus,
		Symbol: ml.IN,
	}}

	// 如果传入订单则单个订单进行查询，否则查询出已完成的订单ID
	var orderIds []interface{}
	if !g.IsEmpty(orderId) {
		input = &do.OrderItemListInput{
			Where: do.OrderItem{
				OrderId: orderId,
			},
		}
	} else {
		stateId := []uint{consts.ORDER_STATE_FINISH, consts.ORDER_STATE_RECEIVED}
		infoinput := &do.OrderInfoListInput{
			Where: do.OrderInfo{
				UserId: userId,
			},
		}
		infoinput.WhereExt = []*ml.WhereExt{{
			Column: dao.OrderInfo.Columns().OrderStateId,
			Val:    stateId,
			Symbol: ml.IN,
		}}
		infoids, err := dao.OrderInfo.FindKey(ctx, infoinput)
		if err != nil {
			return nil, err
		}

		if len(infoids) > 0 {
			input.WhereExt = []*ml.WhereExt{{
				Column: dao.OrderItem.Columns().OrderId,
				Val:    orderIds,
				Symbol: ml.IN,
			}}
		}
	}
	orderItems, err := dao.OrderItem.Find(ctx, input)

	if err != nil {
		return nil, err
	}

	var orderItemVos []*model.OrderItemVo
	for _, orderItem := range orderItems {
		orderItemVo := &model.OrderItemVo{}
		gconv.Struct(orderItem, orderItemVo)
		orderItemVos = append(orderItemVos, orderItemVo)
	}

	comment.Items = orderItemVos
	comment.No = 0
	comment.Yes = 0

	var itemIds []uint64
	for _, orderItem := range orderItems {
		itemId := orderItem.ItemId
		if !array.InArray(itemIds, itemId) {
			itemIds = append(itemIds, itemId)
		}
	}

	if g.IsEmpty(orderId) && len(orderIds) > 0 {

		evaluationNo := &do.OrderItemListInput{
			Where: do.OrderItem{
				UserId:                    userId,
				OrderItemEvaluationStatus: consts.ORDER_ITEM_EVALUATION_NO,
			},
		}
		evaluationNo.WhereExt = []*ml.WhereExt{{
			Column: dao.OrderItem.Columns().OrderId,
			Val:    orderIds,
			Symbol: ml.IN,
		}}
		noNum, err := dao.OrderItem.Count(ctx, evaluationNo)
		if err != nil {
			return nil, err
		}

		evaluationYes := &do.OrderItemListInput{
			Where: do.OrderItem{
				UserId:                    userId,
				OrderItemEvaluationStatus: consts.ORDER_ITEM_EVALUATION_YES,
			},
		}
		evaluationNo.WhereExt = []*ml.WhereExt{{
			Column: dao.OrderItem.Columns().OrderId,
			Val:    orderIds,
			Symbol: ml.IN,
		}}
		yesNum, err := dao.OrderItem.Count(ctx, evaluationYes)
		if err != nil {
			return nil, err
		}
		comment.No = noNum
		comment.Yes = yesNum
	}

	return comment, nil
}
