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
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

type sProductSpec struct{}

func init() {
	service.RegisterProductSpec(NewProductSpec())
}

func NewProductSpec() *sProductSpec {
	return &sProductSpec{}
}

// Get 读取规格
func (s *sProductSpec) Get(ctx context.Context, id any) (out *entity.ProductSpec, err error) {
	var list []*entity.ProductSpec
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 读取多条规格
func (s *sProductSpec) Gets(ctx context.Context, id any) (list []*entity.ProductSpec, err error) {
	err = dao.ProductSpec.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sProductSpec) Find(ctx context.Context, in *do.ProductSpecListInput) (out []*entity.ProductSpec, err error) {
	out, err = dao.ProductSpec.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sProductSpec) List(ctx context.Context, in *do.ProductSpecListInput) (out *do.ProductSpecListOutput, err error) {
	in.Sort = ml.ORDER_BY_ASC
	in.Sidx = dao.ProductSpec.Columns().SpecSort
	out, err = dao.ProductSpec.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sProductSpec) Add(ctx context.Context, in *do.ProductSpec) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ProductSpec.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sProductSpec) Edit(ctx context.Context, in *do.ProductSpec) (affected int64, err error) {
	_, err = dao.ProductSpec.Edit(ctx, in.SpecId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sProductSpec) Remove(ctx context.Context, id any) (affected int64, err error) {

	input := &do.ProductInfoListInput{}
	input.WhereExt = []*ml.WhereExt{{
		Column: dao.ProductInfo.Columns().SpecIds,
		Val:    id,
		Symbol: ml.IN,
	}}

	count, err := dao.ProductInfo.Count(ctx, input)
	if err != nil {
		return 0, err
	}

	if count > 0 {
		return 0, errors.New("规格已被商品使用，不可删除")
	}

	//是否内置
	one, err := dao.ProductSpec.Get(ctx, id)

	if one.SpecBuildin {
		return 0, errors.New("系统内置，不可删除")
	}

	affected, err = dao.ProductSpec.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}
