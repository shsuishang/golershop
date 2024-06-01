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
	"errors"
	"fmt"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/dao/global"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"sync"
)

var (
	userLevelMapMu sync.RWMutex
)

type sUserLevel struct{}

func init() {
	service.RegisterUserLevel(NewUserLevel())
}

func NewUserLevel() *sUserLevel {
	return &sUserLevel{}
}

// Find 查询数据
func (s *sUserLevel) Find(ctx context.Context, in *do.UserLevelListInput) (out []*entity.UserLevel, err error) {
	out, err = dao.UserLevel.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserLevel) List(ctx context.Context, in *do.UserLevelListInput) (out *do.UserLevelListOutput, err error) {
	out, err = dao.UserLevel.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserLevel) Add(ctx context.Context, in *do.UserLevel) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserLevel.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserLevel) Edit(ctx context.Context, in *do.UserLevel) (affected int64, err error) {
	_, err = dao.UserLevel.Edit(ctx, in.UserLevelId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserLevel) Remove(ctx context.Context, id any) (affected int64, err error) {
	//是否内置
	one, err := dao.UserLevel.Get(ctx, id)
	if one.UserLevelIsBuildin {
		return 0, errors.New("系统内置，不可删除")
	}

	//是否有子项
	count, err := dao.UserInfo.Ctx(ctx).Count(do.UserInfo{UserLevelId: id})

	if err != nil {
		return 0, err
	}

	if count > 0 {
		return 0, errors.New(fmt.Sprintf("有 %d 条记录使用，不可删除", count))
	}

	affected, err = dao.UserLevel.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

func (s *sUserLevel) initUserLevelData(ctx context.Context) {
	userLevels, _ := dao.UserLevel.Find(ctx, &do.UserLevelListInput{})

	//userLevelMapMu.Lock() // 加写锁
	//defer userLevelMapMu.Unlock()

	global.UserLevelMap = make(map[uint]string)
	global.UserLevelRateMap = make(map[uint]float64)

	for _, userLevel := range userLevels {
		global.UserLevelMap[userLevel.UserLevelId] = userLevel.UserLevelName
		global.UserLevelRateMap[userLevel.UserLevelId] = float64(userLevel.UserLevelRate)
	}
}

func (s *sUserLevel) GetUserLevelRateMap(ctx context.Context) map[uint]float64 {
	userLevelMapMu.RLock() // 加读锁
	defer userLevelMapMu.RUnlock()

	if len(global.UserLevelRateMap) == 0 {
		s.initUserLevelData(ctx)
	}

	return global.UserLevelRateMap
}
