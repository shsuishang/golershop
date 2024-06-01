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

package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"time"
)

type sUserVoucher struct{}

func init() {
	service.RegisterUserVoucher(NewUserVoucher())
}

func NewUserVoucher() *sUserVoucher {
	return &sUserVoucher{}
}

// Get 读取用户优惠券
func (s *sUserVoucher) Get(ctx context.Context, id any) (out *entity.UserVoucher, err error) {
	var list []*entity.UserVoucher
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 读取多条用户优惠券
func (s *sUserVoucher) Gets(ctx context.Context, id any) (list []*entity.UserVoucher, err error) {
	err = dao.UserVoucher.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sUserVoucher) Find(ctx context.Context, in *do.UserVoucherListInput) (out []*entity.UserVoucher, err error) {
	out, err = dao.UserVoucher.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserVoucher) List(ctx context.Context, in *do.UserVoucherListInput) (out *do.UserVoucherListOutput, err error) {
	out, err = dao.UserVoucher.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserVoucher) Add(ctx context.Context, in *do.UserVoucher) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserVoucher.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserVoucher) Edit(ctx context.Context, in *do.UserVoucher) (affected int64, err error) {
	_, err = dao.UserVoucher.Edit(ctx, in.UserVoucherId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Edit 编辑
func (s *sUserVoucher) EditWhere(ctx context.Context, where *do.UserVoucherListInput, in *do.UserVoucher) (affected int64, err error) {
	affected, err = dao.UserVoucher.EditWhere(ctx, where, in)

	if err != nil {
		return 0, err
	}

	return
}

// Remove 删除多条记录模式
func (s *sUserVoucher) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.UserVoucher.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// Save 保存
func (s *sUserVoucher) Save(ctx context.Context, in *do.UserVoucher) (affected int64, err error) {
	return dao.UserVoucher.Save(ctx, in)
}

// List 分页读取
func (s *sUserVoucher) GetList(ctx context.Context, in *do.UserVoucherListInput) (output *model.UserVoucherListOutput, err error) {
	output = &model.UserVoucherListOutput{}

	voucherPage, err := s.List(ctx, in)
	if err != nil {
		return nil, err
	}

	gconv.Scan(voucherPage, output)

	if voucherPage != nil && len(voucherPage.Items) > 0 {
		output.Items = make([]*model.UserVoucherRes, 0, len(voucherPage.Items))
		currentTime := time.Now().Unix() * 1000

		for _, userVoucher := range voucherPage.Items {
			userVoucherRes := &model.UserVoucherRes{}

			// 判断优惠券状态是否过期
			if userVoucher.VoucherStateId == consts.VOUCHER_STATE_UNUSED && userVoucher.VoucherEndDate < currentTime {
				userVoucherRes.VoucherStateId = consts.VOUCHER_STATE_TIMEOUT

				// 更新数据
				voucher := &do.UserVoucher{UserVoucherId: userVoucher.UserVoucherId, VoucherStateId: consts.VOUCHER_STATE_TIMEOUT}
				s.Edit(ctx, voucher)
			}

			// 未生效标记
			if userVoucher.VoucherStateId == consts.VOUCHER_STATE_UNUSED && userVoucher.VoucherStartDate > currentTime {
				userVoucherRes.VoucherEffect = false
			}

			output.Items = append(output.Items, userVoucherRes)
		}
	}

	return output, nil
}
