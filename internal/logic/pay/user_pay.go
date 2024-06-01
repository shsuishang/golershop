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

package pay

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

type sUserPay struct{}

func init() {
	service.RegisterUserPay(NewUserPay())
}

func NewUserPay() *sUserPay {
	return &sUserPay{}
}

// Get 读取信息
func (s *sUserPay) Get(ctx context.Context, id any) (out *entity.UserPay, err error) {
	var list []*entity.UserPay
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 读取多条信息
func (s *sUserPay) Gets(ctx context.Context, id any) (list []*entity.UserPay, err error) {
	err = dao.UserPay.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sUserPay) Find(ctx context.Context, in *do.UserPayListInput) (out []*entity.UserPay, err error) {
	out, err = dao.UserPay.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserPay) List(ctx context.Context, in *do.UserPayListInput) (out *do.UserPayListOutput, err error) {
	out, err = dao.UserPay.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserPay) Add(ctx context.Context, in *do.UserPay) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserPay.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserPay) Edit(ctx context.Context, in *do.UserPay) (affected int64, err error) {
	_, err = dao.UserPay.Edit(ctx, in.UserId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserPay) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.UserPay.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetPayPasswd 获取支付密码
func (s *sUserPay) GetPayPasswd(ctx context.Context, userId uint) (*entity.UserPay, error) {
	userPay, err := dao.UserPay.Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	if userPay == nil {
		return nil, errors.New("支付密码获取失败")
	}
	return userPay, nil
}

// ChangePayPassword 修改支付密码
func (s *sUserPay) ChangePayPassword(ctx context.Context, oldPayPassword, newPayPassword, payPassword string, userId uint) (bool, error) {
	userPay, err := s.GetPayPasswd(ctx, userId)
	if err != nil {
		return false, err
	}

	if userPay == nil {
		// 新建支付密码
		if g.IsEmpty(newPayPassword) || g.IsEmpty(payPassword) {
			return false, errors.New("密码不能为空！")
		}
		if newPayPassword != payPassword {
			return false, errors.New("两次输入密码不一致！")
		}

		userPaySalt := grand.S(10)
		encryptedPassword := gmd5.MustEncryptString(userPaySalt + gmd5.MustEncryptString(newPayPassword))

		pay := &do.UserPay{
			UserId:        userId,
			UserPayPasswd: encryptedPassword,
			UserPaySalt:   userPaySalt,
		}

		if _, err := dao.UserPay.Save(ctx, pay); err != nil {
			return false, errors.New("添加支付密码失败！")
		}
		return true, nil
	}

	// 修改密码
	oldSalt := userPay.UserPaySalt
	oldEncryptedPassword := gmd5.MustEncryptString(oldSalt + gmd5.MustEncryptString(oldPayPassword))

	if userPay.UserPayPasswd != oldEncryptedPassword {
		return false, errors.New("原支付密码不正确！")
	}

	newSalt := grand.S(10)
	newEncryptedPassword := gmd5.MustEncryptString(newSalt + gmd5.MustEncryptString(newPayPassword))

	if newEncryptedPassword == oldEncryptedPassword {
		return false, errors.New("新密码不能与原密码相同！")
	}

	pay := &do.UserPay{
		UserPayPasswd: newEncryptedPassword,
		UserPaySalt:   newSalt,
	}
	if _, err := dao.UserPay.Edit(ctx, userId, pay); err != nil {
		return false, errors.New("修改支付密码失败！")
	}

	return true, nil
}
