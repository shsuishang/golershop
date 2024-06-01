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
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"sort"
)

type sProductType struct{}

func init() {
	service.RegisterProductType(NewProductType())
}

func NewProductType() *sProductType {
	return &sProductType{}
}

// 读取类型
func (s *sProductType) Get(ctx context.Context, id any) (out *entity.ProductType, err error) {
	var list []*entity.ProductType
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// 读取多条类型
func (s *sProductType) Gets(ctx context.Context, id any) (list []*entity.ProductType, err error) {
	err = dao.ProductType.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sProductType) Find(ctx context.Context, in *do.ProductTypeListInput) (out []*entity.ProductType, err error) {
	out, err = dao.ProductType.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sProductType) List(ctx context.Context, in *do.ProductTypeListInput) (out *do.ProductTypeListOutput, err error) {
	out, err = dao.ProductType.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sProductType) Add(ctx context.Context, in *do.ProductType) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ProductType.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sProductType) Edit(ctx context.Context, in *do.ProductType) (affected int64, err error) {
	//是否可以取消规格判断

	specIds := in.SpecIds

	//读取当前已经存在的type_id
	typeId := in.TypeId

	//修改
	if !g.IsEmpty(typeId) {
		productTypeOld, err := s.Get(ctx, typeId)

		if err != nil {
			return 0, err
		}

		if productTypeOld != nil {
			specIdsOld := productTypeOld.SpecIds

			typeSpecIdRowOld := gstr.Split(specIdsOld, ",")
			typeSpecIdRow := gstr.Split(specIds.(string), ",")

			for _, typeSpecIdOld := range typeSpecIdRowOld {
				//已经使用,本次更新不存在。
				if !array.InArray(typeSpecIdRow, typeSpecIdOld) {
					input := do.ProductInfoListInput{}

					var findInSetList []uint64

					input.WhereExt = append(input.WhereExt, &ml.WhereExt{
						Column: dao.ProductInfo.Columns().SpecIds,
						Val:    append(findInSetList, gconv.Uint64(typeSpecIdOld)),
						Symbol: ml.FIND_IN_SET,
					})

					count, err := dao.ProductInfo.Count(ctx, &input)

					if err != nil {
						return 0, err
					}

					if count > 0 {
						return 0, gerror.New(fmt.Sprintf("规格 %d 已经被 %d 个SPU商品使用，不可取消关联", typeSpecIdOld, count))
					}
				}
			}
		}
	}

	if g.IsEmpty(specIds) {
		in.SpecIds = ""
	} else {
	}

	if g.IsEmpty(in.BrandIds) {
		in.BrandIds = ""
	}

	_, err = dao.ProductType.Edit(ctx, in.TypeId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Update
func (s *sProductType) UpdateAssistIds(ctx context.Context, typeId interface{}) (affected int64, err error) {
	var in do.ProductAssistListInput
	in.Where.TypeId = typeId

	//特别标记， 一般不允许service 调用 service
	assistIds, err := dao.ProductAssist.FindKey(ctx, &in)

	var productType do.ProductType
	productType.AssistIds = gstr.JoinAny(assistIds, ",")

	_, err = dao.ProductType.Edit(ctx, typeId, &productType)

	if err != nil {
		return 0, err
	}

	return
}

// Remove 删除多条记录模式
func (s *sProductType) Remove(ctx context.Context, id any) (affected int64, err error) {
	//是否内置
	one, err := dao.ProductType.Get(ctx, id)

	if one.TypeBuildin {
		return 0, errors.New("系统内置，不可删除")
	}

	//是否有子项
	typeCount, err := dao.ProductCategory.Ctx(ctx).Count(do.ProductCategory{TypeId: id})

	if err != nil {
		return 0, err
	}

	if typeCount > 0 {
		return 0, errors.New(fmt.Sprintf("有 %d 条分类使用，不可删除", typeCount))
	}

	affected, err = dao.ProductType.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// Info 读取类型信息
func (s *sProductType) Info(ctx context.Context, id any) (out *model.ProductTypeInfoOutput, err error) {
	row, err := s.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	out = &model.ProductTypeInfoOutput{}

	brands, err := service.ProductBrand().Gets(ctx, gstr.Split(row.BrandIds, ","))
	out.Brands = brands

	//读取属性
	assistIds := gstr.Split(row.AssistIds, ",")
	assists, err := service.ProductAssist().Gets(ctx, assistIds)

	gconv.Scan(assists, &out.Assists)

	//读取属性值
	// 创建一个空 map 用于存储属性值
	assistMap := make(map[interface{}][]*entity.ProductAssistItem)

	assistItems, err := service.ProductAssistItem().Find(ctx, &do.ProductAssistItemListInput{
		Where: do.ProductAssistItem{AssistId: assistIds},
	})

	for _, item := range assistItems {
		assistMap[item.AssistId] = append(assistMap[item.AssistId], item)
	}

	for _, item := range out.Assists {
		item.Items = assistMap[item.AssistId]
	}

	specIds := gstr.Split(row.SpecIds, ",")
	specs, err := service.ProductSpec().Gets(ctx, specIds)
	gconv.Scan(specs, &out.Specs)

	//进行降序 i大于j 为降序
	sort.Slice(out.Specs, func(i, j int) bool {
		return out.Specs[i].SpecSort > out.Specs[j].SpecSort
	})

	//排序，image第一个显示
	sort.Slice(out.Specs, func(i, j int) bool {
		return "image" == out.Specs[i].SpecFormat
	})

	// 创建一个空 map 用于存储规格值
	specMap := make(map[interface{}][]*entity.ProductSpecItem)
	specItems, err := service.ProductSpecItem().Find(ctx, &do.ProductSpecItemListInput{Where: do.ProductSpecItem{SpecId: specIds}, BaseList: ml.BaseList{
		Sidx: dao.ProductSpecItem.Columns().SpecItemSort,
		Sort: "ASC"}})

	for _, item := range specItems {
		specMap[item.SpecId] = append(specMap[item.SpecId], item)
	}

	for _, item := range out.Specs {
		item.Items = specMap[item.SpecId]
	}

	return out, nil
}
