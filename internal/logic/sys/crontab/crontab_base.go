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

package crontab

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

type sCrontabBase struct{}

func init() {
	service.RegisterCrontabBase(New())
}

func New() *sCrontabBase {
	return &sCrontabBase{}
}

// Find 查询数据
func (s *sCrontabBase) Find(ctx context.Context, in *do.CrontabBaseListInput) (out []*entity.CrontabBase, err error) {
	out, err = dao.CrontabBase.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sCrontabBase) List(ctx context.Context, in *do.CrontabBaseListInput) (out *do.CrontabBaseListOutput, err error) {
	out, err = dao.CrontabBase.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sCrontabBase) Add(ctx context.Context, in *do.CrontabBase) (lastInsertId int64, err error) {
	lastInsertId, err = dao.CrontabBase.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sCrontabBase) Edit(ctx context.Context, in *do.CrontabBase) (affected int64, err error) {
	_, err = dao.CrontabBase.Edit(ctx, in.CrontabId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sCrontabBase) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.CrontabBase.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// EditState 编辑任务状态
func (c *sCrontabBase) EditState(ctx context.Context, req *sys.CrontabBaseEditStateReq) (res *sys.CrontabBaseEditStateRes, err error) {

	input := do.CrontabBase{}
	gconv.Scan(req, &input)

	var result, error = service.CrontabBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.CrontabBaseEditStateRes{
		CrontabId: result,
	}

	return
}
