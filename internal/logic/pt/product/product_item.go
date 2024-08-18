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
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/pt"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"math"
)

type sProductItem struct{}

func init() {
	service.RegisterProductItem(NewProductItem())
}

func NewProductItem() *sProductItem {
	return &sProductItem{}
}

// Get 读取SKU
func (s *sProductItem) Get(ctx context.Context, id any) (out *entity.ProductItem, err error) {
	var list []*entity.ProductItem
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 读取多条SKU
func (s *sProductItem) Gets(ctx context.Context, id any) (list []*entity.ProductItem, err error) {
	list, err = dao.ProductItem.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sProductItem) Find(ctx context.Context, in *do.ProductItemListInput) (out []*entity.ProductItem, err error) {
	out, err = dao.ProductItem.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sProductItem) List(ctx context.Context, in *do.ProductItemListInput) (out *do.ProductItemListOutput, err error) {
	out, err = dao.ProductItem.List(ctx, in)

	return out, err
}

// Edit 编辑
func (s *sProductItem) Edit(ctx context.Context, in *do.ProductItem) (affected int64, err error) {
	_, err = dao.ProductItem.Edit(ctx, in.ItemId, in)

	if err != nil {
		return 0, err
	}
	return
}

// LockSkuStock 锁库存
func (s *sProductItem) LockSkuStock(ctx context.Context, itemId uint64, cartQuantity uint) (affected int64, err error) {

	res, err := dao.ProductItem.Ctx(ctx).WherePri(itemId).Where("item_quantity - item_quantity_frozen >= ?", cartQuantity).Increment(dao.ProductItem.Columns().ItemQuantityFrozen, cartQuantity)
	//dao.ProductItem.Increment(ctx, itemId, dao.ProductItem.Columns().ItemQuantityFrozen, cartQuantity)

	if err != nil {
		return 0, err
	}

	affected, err = res.RowsAffected()

	return
}

// EditStock 编辑库存
func (s *sProductItem) EditStock(ctx context.Context, input *model.ProductEditStockInput) (success bool, err error) {
	// 根据商品项ID获取商品项信息
	v, err := dao.ProductItem.Get(ctx, input.ItemId)
	if err != nil {
		return false, err
	}

	// 如果商品项信息不为空
	if v != nil {
		// 初始化库存流水项
		stockBillItem := &do.StockBillItem{
			ProductId:             v.ProductId,
			ItemId:                v.ItemId,
			ItemName:              v.ItemName,
			BillItemQuantity:      input.ItemQuantity,
			WarehouseItemQuantity: v.ItemQuantity,
			BillItemUnitPrice:     v.ItemUnitPrice,
			BillItemSubtotal:      v.ItemUnitPrice * gconv.Float64(input.ItemQuantity),
		}

		// 如果是入库操作
		if input.BillTypeId == consts.BILL_TYPE_IN {
			stockBillItem.BillTypeId = consts.BILL_TYPE_IN
			stockBillItem.StockTransportTypeId = consts.STOCK_IN_OTHER
			v.ItemQuantity += input.ItemQuantity
		} else { // 如果是出库操作
			stockBillItem.BillTypeId = consts.BILL_TYPE_OUT
			stockBillItem.StockTransportTypeId = consts.STOCK_OUT_OTHER
			v.ItemQuantity -= input.ItemQuantity
		}

		// 添加库存流水项
		if _, err := dao.StockBillItem.Add(ctx, stockBillItem); err != nil {
			return false, err
		}

		// 更新商品项信息
		if _, err := dao.ProductItem.Edit(ctx, v.ItemId, &do.ProductItem{
			ItemQuantity: v.ItemQuantity,
		}); err != nil {
			return false, err
		}

		return true, nil
	}

	return false, fmt.Errorf("商品项不存在")
}

// ReleaseSkuStock 锁库存
func (s *sProductItem) ReleaseSkuStock(ctx context.Context, itemId uint64, releaseQuantity uint) (affected int64, err error) {
	res, err := dao.ProductItem.Ctx(ctx).WherePri(itemId).Where("item_quantity_frozen >= ?", releaseQuantity).Decrement(dao.ProductItem.Columns().ItemQuantityFrozen, releaseQuantity)

	if err != nil {
		return 0, err
	}

	affected, err = res.RowsAffected()
	return
}

// List 分页读取
func (d *sProductItem) ListItemKey(ctx context.Context, req *pt.ItemListReq) (out *do.ItemListKeyOutput, err error) {

	var (
		orm    = dao.ProductItem.Ctx(ctx).OmitEmpty()
		rTable = dao.ProductItem.Table()
		rCls   = dao.ProductItem.Columns()
		nTable = dao.ProductIndex.Table()
		nCls   = dao.ProductIndex.Columns()
	)

	where := do.ProductItem{
		ItemEnable: req.ItemEnable,
	}

	if !g.IsEmpty(req.ItemId) {
		where.ItemId = gconv.SliceUint64(gstr.Split(req.ItemId, ","))
	}

	orm = orm.LeftJoinOnField(nTable, rCls.ProductId)
	orm = orm.WherePrefix(rTable, where)

	orm = orm.WherePrefix(nTable, do.ProductIndex{
		StoreId:        req.StoreId,
		ProductId:      req.ProductId,
		CategoryId:     req.CategoryIds,
		StoreType:      req.StoreType,
		KindId:         req.KindId,
		BrandId:        req.BrandId,
		ProductStateId: req.ProductStateId,
	})

	// Fuzzy like querying.
	if req.ProductNameIndex != "" {
		var (
			keyLike = "%" + req.ProductNameIndex + "%"
		)
		whereFormat := fmt.Sprintf(
			"(`%s`.`%s` like ?)  ",
			nTable, nCls.ProductNameIndex,
		)

		orm = orm.Where(whereFormat, keyLike)
	}

	if req.ProductName != "" {
		var (
			keyLike = "%" + req.ProductName + "%"
		)
		whereFormat := fmt.Sprintf(
			"(`%s`.`%s` like ?)  ",
			nTable, nCls.ProductName,
		)

		orm = orm.Where(whereFormat, keyLike)
	}

	orm = orm.OmitNil()

	out = &do.ItemListKeyOutput{}
	out.Page = req.Page
	out.Size = req.Size

	// 查询记录总数
	count, err1 := orm.Count()
	if err1 != nil {
		return nil, err1
	}

	out.Records = count
	out.Total = int(math.Ceil(float64(count) / float64(out.Size)))

	orm = orm.Distinct().FieldsPrefix(rTable, rCls.ItemId)

	if !g.IsEmpty(req.Sidx) && !g.IsEmpty(req.Sort) {
		orm = orm.Order(req.Sidx, req.Sort)
	}

	// Resource items..
	idRes, err := orm.Page(req.Page, req.Size).All()

	if err != nil {
		return out, err
	}

	for _, record := range idRes {
		if !record[rCls.ItemId].IsEmpty() {
			out.Items = append(out.Items, record[rCls.ItemId].Uint64())
		}
	}

	return out, nil
}

// ifOnSale 函数判断商品是否在售
func (dao *sProductItem) IfOnSale(ctx context.Context, item *model.ProductItemVo) bool {
	// 计算可用库存
	availableQuantity := item.ItemQuantity - item.ItemQuantityFrozen

	// 判断是否在售
	return item.ProductStateId == consts.PRODUCT_STATE_NORMAL && item.ItemEnable == consts.PRODUCT_STATE_NORMAL &&
		availableQuantity > 0 && item.ItemQuantity > 0
}

// BatchEditStock 批量编辑库存
func (s *sProductItem) BatchEditStock(ctx context.Context, inputs []*model.ProductEditStockInput) (err error) {

	// 获取所有ItemId的集合
	itemIds := array.Column(inputs, "ItemId")

	// 根据itemIds获取商品SKU信息
	productItems, err := s.Gets(ctx, itemIds)
	if err != nil {
		return err
	}
	if len(productItems) == 0 {
		return gerror.New("商品SKU信息不存在！")
	}

	// 将输入的ProductEditStockInput转换为以ItemId为键的映射
	editStockInputMap := make(map[uint64]*model.ProductEditStockInput)
	for _, input := range inputs {
		editStockInputMap[input.ItemId] = input
	}

	// 初始化stockBillItems列表
	var stockBillItems []*do.StockBillItem

	// 遍历商品SKU信息，更新库存并生成StockBillItem
	for _, v := range productItems {
		input := editStockInputMap[v.ItemId]
		if input != nil {
			stockBillItem := &do.StockBillItem{
				ProductId:             v.ProductId,
				ItemId:                v.ItemId,
				ItemName:              v.ItemName,
				BillItemQuantity:      input.ItemQuantity,
				WarehouseItemQuantity: v.ItemQuantity,
			}

			// 根据BillTypeId设置入库或出库类型，并更新库存
			if input.BillTypeId == consts.BILL_TYPE_IN {
				stockBillItem.BillTypeId = consts.BILL_TYPE_IN
				stockBillItem.StockTransportTypeId = consts.STOCK_IN_OTHER

				// 增加库存
				v.ItemQuantity += input.ItemQuantity
			} else {
				stockBillItem.BillTypeId = consts.BILL_TYPE_OUT
				stockBillItem.StockTransportTypeId = consts.STOCK_OUT_OTHER

				// 减少库存
				if v.AvailableQuantity >= input.ItemQuantity {
					v.ItemQuantity -= input.ItemQuantity
				} else {
					return gerror.New("出库数量不能大于总库存！")
				}
			}

			// 设置单价和小计
			stockBillItem.BillItemUnitPrice = v.ItemUnitPrice
			stockBillItem.BillItemSubtotal = v.ItemUnitPrice * (gconv.Float64(stockBillItem.BillItemQuantity))

			stockBillItems = append(stockBillItems, stockBillItem)
		}
	}

	// 保存或更新StockBillItem
	if _, err := dao.StockBillItem.Saves(ctx, stockBillItems); err != nil {
		return gerror.New("保存出入库单据失败！")
	}

	// 保存或更新ProductItem
	input := []*do.ProductItem{}
	gconv.Structs(productItems, &input)
	if _, err := dao.ProductItem.Saves(ctx, input); err != nil {
		return gerror.New("修改商品SKU信息失败！")
	}

	return nil
}

// GetStockWarningItems 获取库存预警商品
func (s *sProductItem) GetStockWarningItems(ctx context.Context, input *model.ProductItemInput) (*model.ItemListOutput, error) {
	// 获取库存预警值
	stockWarning := service.ConfigBase().GetInt(ctx, "stock_warning", 5)
	input.StockWarning = stockWarning

	// 调用 DAO 层方法获取库存预警商品
	out, err := dao.ProductItem.GetStockWarningItems(ctx, input)
	if err != nil {
		return nil, err
	}

	return out, nil
}
