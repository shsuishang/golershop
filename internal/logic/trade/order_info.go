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
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
)

type sOrderInfo struct{}

func init() {
	service.RegisterOrderInfo(NewOrderInfo())
}

func NewOrderInfo() *sOrderInfo {
	return &sOrderInfo{}
}

// Get 读取订单
func (s *sOrderInfo) Get(ctx context.Context, id any) (out *entity.OrderInfo, err error) {
	var list []*entity.OrderInfo
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
func (s *sOrderInfo) Gets(ctx context.Context, id any) (list []*entity.OrderInfo, err error) {
	err = dao.OrderInfo.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sOrderInfo) Find(ctx context.Context, in *do.OrderInfoListInput) (out []*entity.OrderInfo, err error) {
	out, err = dao.OrderInfo.Find(ctx, in)

	return out, err
}

// GetList 订单搜索查询列表
func (s *sOrderInfo) GetList(ctx context.Context, in *do.OrderInfoListInput) (out *model.OrderListOutput, err error) {
	orderInfoList, err := dao.OrderInfo.List(ctx, in)
	gconv.Scan(orderInfoList, &out)

	//补全商品基础表信息
	ids := array.Column(orderInfoList.Items, dao.OrderInfo.Columns().OrderId)
	orderBaseList, err := dao.OrderBase.Gets(ctx, ids)

	//读取订单商品
	orderItem, err := dao.OrderItem.Find(ctx, &do.OrderItemListInput{Where: do.OrderItem{OrderId: ids}})

	//处理为map
	orderItemMap := make(map[string][]*entity.OrderItem)

	for _, item := range orderItem {
		orderItemMap[item.OrderId] = append(orderItemMap[item.OrderId], item)
	}

	for _, item := range orderBaseList {
		for _, vo := range out.Items {
			if item.OrderId == vo.OrderId {
				vo.OrderNumber = item.OrderNumber
				vo.OrderTime = item.OrderTime
				vo.OrderProductAmount = item.OrderProductAmount
				vo.OrderPaymentAmount = item.OrderPaymentAmount
				vo.CurrencyId = item.CurrencyId
				vo.CurrencySymbolLeft = item.CurrencySymbolLeft
				vo.StoreName = item.StoreName
				vo.UserNickname = item.UserNickname
			}
		}
	}

	for _, vo := range out.Items {
		gconv.Scan(orderItemMap[vo.OrderId], &vo.Items)
	}

	return out, err
}

// List 分页读取
func (s *sOrderInfo) List(ctx context.Context, in *do.OrderInfoListInput) (out *do.OrderInfoListOutput, err error) {
	out, err = dao.OrderInfo.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sOrderInfo) Add(ctx context.Context, in *do.OrderInfo) (lastInsertId int64, err error) {
	lastInsertId, err = dao.OrderInfo.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sOrderInfo) Edit(ctx context.Context, in *do.OrderInfo) (affected int64, err error) {
	_, err = dao.OrderInfo.Edit(ctx, in.OrderId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sOrderInfo) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.OrderInfo.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}
