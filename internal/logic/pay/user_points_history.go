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
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
)

type sUserPointsHistory struct{}

func init() {
	service.RegisterUserPointsHistory(NewUserPointsHistory())
}

func NewUserPointsHistory() *sUserPointsHistory {
	return &sUserPointsHistory{}
}

// Find 查询数据
func (s *sUserPointsHistory) Find(ctx context.Context, in *do.UserPointsHistoryListInput) (out []*entity.UserPointsHistory, err error) {
	out, err = dao.UserPointsHistory.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserPointsHistory) List(ctx context.Context, in *do.UserPointsHistoryListInput) (out *do.UserPointsHistoryListOutput, err error) {
	out, err = dao.UserPointsHistory.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserPointsHistory) Add(ctx context.Context, in *do.UserPointsHistory) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserPointsHistory.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserPointsHistory) Edit(ctx context.Context, in *do.UserPointsHistory) (affected int64, err error) {
	_, err = dao.UserPointsHistory.Edit(ctx, in.UserId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserPointsHistory) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.UserPointsHistory.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

func (s *sUserPointsHistory) GetList(ctx context.Context, in *do.UserPointsHistoryListInput) (out *model.UserPointsHistoryOutput, err error) {
	// 获取 UserPointsHistory 列表
	lists, err := s.List(ctx, in)
	if err != nil {
		return nil, err
	}

	if lists != nil && len(lists.Items) > 0 {

		// 提取所有 userId
		userIds := gconv.SliceUint(array.Column(lists.Items, "UserId"))

		// 获取用户信息映射
		userInfoMap, err := dao.UserInfo.GetUserInfoMap(ctx, userIds)
		if err != nil {
			return nil, err
		}

		out = &model.UserPointsHistoryOutput{}
		gconv.Scan(lists, out)
		for i := range out.Items {
			userPointsHistory := out.Items[i]
			if len(userInfoMap) > 0 {
				userInfo := userInfoMap[int(userPointsHistory.UserId)]
				if userInfo != nil {
					userPointsHistory.UserNickname = userInfo.UserNickname
				}
			}
		}
	}

	return out, nil
}
