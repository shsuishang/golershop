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
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

type sProductAssist struct{}

func init() {
	service.RegisterProductAssist(NewAssist())
}

func NewAssist() *sProductAssist {
	return &sProductAssist{}
}

// Get 读取属性
func (s *sProductAssist) Get(ctx context.Context, id any) (out *entity.ProductAssist, err error) {
	var list []*entity.ProductAssist
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 读取多条属性
func (s *sProductAssist) Gets(ctx context.Context, id any) (list []*entity.ProductAssist, err error) {
	err = dao.ProductAssist.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sProductAssist) Find(ctx context.Context, in *do.ProductAssistListInput) (out []*entity.ProductAssist, err error) {
	out, err = dao.ProductAssist.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sProductAssist) List(ctx context.Context, in *do.ProductAssistListInput) (out *do.ProductAssistListOutput, err error) {
	out, err = dao.ProductAssist.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sProductAssist) Add(ctx context.Context, in *do.ProductAssist) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ProductAssist.Add(ctx, in)
	if err != nil {
		return 0, err
	}

	_, err = service.ProductType().UpdateAssistIds(ctx, in.TypeId)
	if err != nil {
		return 0, err
	}

	return lastInsertId, err
}

// Edit 编辑
func (s *sProductAssist) Edit(ctx context.Context, in *do.ProductAssist) (affected int64, err error) {
	affected, err = dao.ProductAssist.Edit(ctx, in.AssistId, in)

	if err != nil {
		return 0, err
	}

	return
}

// Remove 删除多条记录模式
func (s *sProductAssist) Remove(ctx context.Context, id any) (affected int64, err error) {
	one, err := dao.ProductAssist.Get(ctx, id)

	//todo 是否可以删除判断

	affected, err = dao.ProductAssist.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	_, err = service.ProductType().UpdateAssistIds(ctx, one.TypeId)
	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetAssists 获取商品辅助属性
func (s *sProductAssist) GetAssists(ctx context.Context, assistIds string) (out []*model.ProductAssistOutput, err error) {
	out = make([]*model.ProductAssistOutput, 0)

	if assistIds == "" {
		return
	}

	assistIdList := gconv.Strings(gstr.Split(assistIds, ","))

	if len(assistIdList) == 0 {
		return
	}

	assistList, err := dao.ProductAssist.Find(ctx, &do.ProductAssistListInput{Where: do.ProductAssist{
		AssistId: assistIdList,
	}})

	assistItemList, err := dao.ProductAssistItem.Find(ctx, &do.ProductAssistItemListInput{Where: do.ProductAssistItem{
		AssistId: assistIdList,
	}})

	for _, assist := range assistList {
		assistOutPut := &model.ProductAssistOutput{}
		gconv.Struct(assist, assistOutPut)

		out = append(out, assistOutPut)

		assistItems := make([]*entity.ProductAssistItem, 0)
		for _, assistItem := range assistItemList {
			if assistOutPut.AssistId == assistItem.AssistId {
				assistItems = append(assistItems, assistItem)
			}
		}

		assistOutPut.Items = assistItems
	}

	return out, err
}
