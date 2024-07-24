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

package marketing

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

type sActivityBase struct{}

func init() {
	service.RegisterActivityBase(NewActivityBase())
}

func NewActivityBase() *sActivityBase {
	return &sActivityBase{}
}

// Get 根据编号读取活动信息
func (s *sActivityBase) Get(ctx context.Context, id any) (out *entity.ActivityBase, err error) {
	var list []*entity.ActivityBase
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 根据编号读取读取多条活动信息
func (s *sActivityBase) Gets(ctx context.Context, id any) (list []*entity.ActivityBase, err error) {
	err = dao.ActivityBase.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询活动数据
func (s *sActivityBase) Find(ctx context.Context, in *do.ActivityBaseListInput) (out []*entity.ActivityBase, err error) {
	out, err = dao.ActivityBase.Find(ctx, in)

	return out, err
}

// List 分页读取活动
func (s *sActivityBase) List(ctx context.Context, in *do.ActivityBaseListInput) (out *model.ActivityListOutput, err error) {
	activityBaseList, err := dao.ActivityBase.List(ctx, in)

	gconv.Scan(activityBaseList, &out)

	return out, err
}

// Add 新增活动
func (s *sActivityBase) Add(ctx context.Context, in *do.ActivityBase) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ActivityBase.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑活动
func (s *sActivityBase) Edit(ctx context.Context, in *do.ActivityBase) (affected int64, err error) {
	_, err = dao.ActivityBase.Edit(ctx, in.ActivityId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除活动记录
func (s *sActivityBase) Remove(ctx context.Context, id any) (affected int64, err error) {
	row, err := s.Get(ctx, id)

	if err != nil {
		return 0, err
	}

	if row.ActivityState == consts.ACTIVITY_STATE_NORMAL {
		return 0, errors.New(fmt.Sprintf("活动:%d 生效使用中，不可删除", row.ActivityId))
	}

	affected, err = dao.ActivityBase.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}
