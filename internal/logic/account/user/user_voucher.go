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
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/shop"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
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

	if voucherPage != nil && len(voucherPage.Items) > 0 {
		output.Items = make([]*model.UserVoucherRes, 0, len(voucherPage.Items))
		currentTime := gtime.Now().TimestampMilli()

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

// GetList 获取用户优惠券列表
func (s *sUserVoucher) GetLists(ctx context.Context, voucherListReq *shop.UserVoucherListReq) (voucherResPage *model.UserVoucherListOutput, err error) {
	voucherResPage = &model.UserVoucherListOutput{}

	// 创建查询条件
	voucherQueryWrapper := do.UserVoucherListInput{}

	if !g.IsEmpty(voucherListReq.UserId) {
		voucherQueryWrapper.Where.UserId = voucherListReq.UserId
	}

	if !g.IsEmpty(voucherListReq.ActivityId) {
		voucherQueryWrapper.Where.ActivityId = voucherListReq.ActivityId
	}

	if !g.IsEmpty(voucherListReq.StoreId) {
		voucherQueryWrapper.Where.StoreId = voucherListReq.StoreId
	}

	//优惠券是否生效
	times := gtime.Now().TimestampMilli()
	if voucherListReq.VoucherEffect {
		var likes = []*ml.WhereExt{{
			Column: dao.UserVoucher.Columns().VoucherStartDate,
			Val:    times,
			Symbol: ml.LE,
		}, {
			Column: dao.UserVoucher.Columns().VoucherEndDate,
			Val:    times,
			Symbol: ml.GT,
		}}
		voucherQueryWrapper.WhereExt = likes
	}

	// 处理全部券1、线下券2、线上券3
	if !g.IsEmpty(voucherListReq) {
		switch voucherListReq.VoucherUserWay {
		case 2:
			var likes = []*ml.WhereExt{{
				Column: dao.UserVoucher.Columns().WriteoffCode,
				Val:    times,
				Symbol: ml.NE,
			}}
			voucherQueryWrapper.WhereExt = likes
		case 3:
			voucherQueryWrapper.Where.WriteoffCode = nil
		}
	}

	if !g.IsEmpty(voucherListReq.VoucherStateId) {
		voucherQueryWrapper.Where.VoucherStateId = voucherListReq.VoucherStateId
	}
	var order = []*ml.BaseOrder{{
		Sort: ml.ORDER_BY_DESC,
		Sidx: dao.UserVoucher.Columns().UserVoucherTime,
	}, {
		Sort: ml.ORDER_BY_ASC,
		Sidx: dao.UserVoucher.Columns().VoucherStateId,
	}}
	voucherQueryWrapper.Order = order

	// 查询分页数据
	voucherPage, err := dao.UserVoucher.Find(ctx, &voucherQueryWrapper)
	if err != nil {
		return nil, err
	}

	if !g.IsEmpty(voucherPage) {
		err := gconv.Struct(voucherPage, &voucherResPage.Items)
		if err != nil {
			return nil, err
		}
		var userVoucherList []model.UserVoucherRes
		gconv.Struct(voucherPage, &userVoucherList)
		userVoucherReList := make([]*model.UserVoucherRes, 0)
		currentTime := gtime.Now().TimestampMilli()

		for _, userVoucher := range userVoucherList {
			var userVoucherRes *model.UserVoucherRes
			gconv.Struct(userVoucher, &userVoucherRes)
			userVoucherRes.VoucherEffect = true
			userVoucherRes.Id = userVoucher.UserVoucherId
			voucherEndDate := userVoucher.VoucherEndDate

			if userVoucher.VoucherStateId == consts.VOUCHER_STATE_UNUSED {
				if voucherEndDate < currentTime {
					userVoucherRes.VoucherStateId = consts.VOUCHER_STATE_TIMEOUT
					// 更新数据
					updateData := &do.UserVoucher{
						UserVoucherId:  userVoucher.UserVoucherId,
						VoucherStateId: consts.VOUCHER_STATE_TIMEOUT,
					}
					dao.UserVoucher.Edit(ctx, userVoucher.UserVoucherId, updateData)
				}

				// 未生效标记
				voucherStartDate := userVoucher.VoucherStartDate
				if voucherStartDate > currentTime {
					userVoucherRes.VoucherEffect = false
				}
			}

			userVoucherReList = append(userVoucherReList, userVoucherRes)
		}

		voucherResPage.Items = userVoucherReList

	}

	return voucherResPage, nil
}

// GetEachVoucherNum 获取每种状态的优惠券数量
func (s *sUserVoucher) GetEachVoucherNum(ctx context.Context, voucherStateId, userId uint) (*shop.GetVoucherNumRes, error) {
	voucherCountRes := &shop.GetVoucherNumRes{}

	// 全部优惠券数量
	userVoucherQuery := &do.UserVoucherListInput{
		Where: do.UserVoucher{
			UserId: userId,
		},
	}
	if voucherStateId != 0 {
		userVoucherQuery.Where.VoucherStateId = voucherStateId
	}
	countAll, err := dao.UserVoucher.Count(ctx, userVoucherQuery)
	if err != nil {
		return nil, err
	}
	voucherCountRes.VoucherAllNum = countAll

	// 线下优惠券数量
	offlineQuery := &do.UserVoucherListInput{
		Where: do.UserVoucher{
			UserId: userId,
		},
	}
	if voucherStateId != 0 {
		offlineQuery.Where.VoucherStateId = voucherStateId
	}
	countOffline, err := dao.UserVoucher.Count(ctx, offlineQuery)
	if err != nil {
		return nil, err
	}
	voucherCountRes.VoucherOfflinedNum = countOffline

	// 线上优惠券数量
	onlineQuery := &do.UserVoucherListInput{
		Where: do.UserVoucher{
			UserId:       userId,
			WriteoffCode: "",
		},
	}
	if voucherStateId != 0 {
		onlineQuery.Where.VoucherStateId = voucherStateId
	}
	countOnline, err := dao.UserVoucher.Count(ctx, onlineQuery)
	if err != nil {
		return nil, err
	}
	voucherCountRes.VoucherOnlinedNum = countOnline

	// 未使用优惠券数量
	unusedQuery := &do.UserVoucherListInput{
		Where: do.UserVoucher{
			UserId:         userId,
			VoucherStateId: consts.VOUCHER_STATE_UNUSED,
		},
	}
	countUnused, err := dao.UserVoucher.Count(ctx, unusedQuery)
	if err != nil {
		return nil, err
	}
	voucherCountRes.VoucherUnusedNum = countUnused

	// 已使用优惠券数量
	usedQuery := &do.UserVoucherListInput{
		Where: do.UserVoucher{
			UserId:         userId,
			VoucherStateId: consts.VOUCHER_STATE_USED,
		},
	}
	countUsed, err := dao.UserVoucher.Count(ctx, usedQuery)
	if err != nil {
		return nil, err
	}
	voucherCountRes.VoucherUsedNum = countUsed

	// 已过期优惠券数量
	timeoutQuery := &do.UserVoucherListInput{
		Where: do.UserVoucher{
			UserId:         userId,
			VoucherStateId: consts.VOUCHER_STATE_TIMEOUT,
		},
	}
	countTimeout, err := dao.UserVoucher.Count(ctx, timeoutQuery)
	if err != nil {
		return nil, err
	}
	voucherCountRes.VoucherTimeoutNum = countTimeout

	return voucherCountRes, nil
}

// AddVoucher 领取代金券
func (s *sUserVoucher) AddVoucher(ctx context.Context, activityId, userId uint) (*entity.UserVoucher, error) {
	voucher := &entity.UserVoucher{
		UserVoucherTime: gtime.Now(),
		ActivityId:      activityId,
		VoucherStateId:  consts.VOUCHER_STATE_UNUSED,
		UserId:          userId,
	}

	// 获取活动信息
	activityBase, err := service.ActivityBase().Get(ctx, activityId)
	if err != nil {
		return nil, gerror.New("活动不存在！")
	}

	if activityBase == nil || activityBase.ActivityState != consts.ACTIVITY_STATE_NORMAL {
		return nil, gerror.New("活动未开启,领取失败！")
	}

	// 判断优惠券是否有等级限制
	activityUseLevel := activityBase.ActivityUseLevel
	if !g.IsEmpty(activityUseLevel) {
		userLevels := gconv.SliceUint(gstr.SplitAndTrim(activityUseLevel, ","))
		if len(userLevels) > 0 {
			userInfo, err := dao.UserInfo.Get(ctx, userId)
			if err != nil {
				return nil, gerror.New("用户信息不存在！")
			}

			if !array.InArray(userLevels, userInfo.UserLevelId) {
				return nil, gerror.New("不属于该优惠券指定的会员等级，领取失败！")
			}
		}
	}
	// 判断代金券数量是否足够
	activityRule := activityBase.ActivityRule
	var activityRuleVo model.ActivityRuleVo
	if err := gconv.Struct(activityRule, &activityRuleVo); err != nil {
		return nil, gerror.New("活动规则为空！")
	}

	voucherVo := activityRuleVo.Voucher
	if g.IsEmpty(voucherVo) {
		return nil, gerror.New("活动优惠券信息为空！")
	}

	voucherQuantity := voucherVo.VoucherQuantity
	voucherQuantityFree := voucherVo.VoucherQuantityFree
	if gconv.Int(voucherQuantity)+gconv.Int(voucherQuantityFree) <= 0 {
		return nil, gerror.New("代金券已经被抢完,领取失败！")
	}

	//优惠券限制判断
	voucherPreQuantity := voucherVo.VoucherPreQuantity
	voucherSize := 0
	userVoucherNum, err := dao.UserVoucherNum.FindOne(ctx, &do.UserVoucherNumListInput{
		Where: do.UserVoucherNum{
			ActivityId: activityId,
			UserId:     userId,
		},
	})
	if err != nil {
		return nil, gerror.New("查询用户优惠券数量失败！")
	}
	if userVoucherNum != nil {
		voucherSize = int(userVoucherNum.UvnNum)
	}

	if uint(voucherSize) < voucherPreQuantity {
		requirement := activityRuleVo.Requirement
		buy := requirement.Buy
		item := buy.Item
		subtotal := buy.Subtotal
		voucherPrice := voucherVo.VoucherPrice
		voucher.VoucherSubtotal = subtotal
		voucher.VoucherPrice = voucherPrice
		voucher.VoucherStartDate = voucherVo.VoucherStartDate
		voucher.VoucherEndDate = voucherVo.VoucherEndDate
		voucher.StoreId = activityBase.StoreId
		voucher.ActivityName = activityBase.ActivityName
		voucher.ActivityRule = activityBase.ActivityRule
		if len(item) > 0 {
			voucher.ItemId = gstr.Join(gconv.Strings(item), ",")
		}

		activityType := activityBase.ActivityType
		pointsNeeded := 0
		if activityType == consts.GET_VOUCHER_BY_POINT {
			points := requirement.Points
			pointsNeeded = points.Needed
			if pointsNeeded <= 0 {
				return nil, gerror.New("该活动信息错误，积分需大于0！")
			}
			// TODO: 积分操作
		}
		input := &do.UserVoucher{}
		gconv.Scan(voucher, input)
		if _, err := dao.UserVoucher.Save(ctx, input); err != nil {
			return nil, gerror.New("用户优惠券领取失败！")
		}

		newUserVoucherNum := do.UserVoucherNum{}
		gconv.Scan(userVoucherNum, &newUserVoucherNum)
		if !g.IsEmpty(newUserVoucherNum) {
			newUserVoucherNum.UvnNum = userVoucherNum.UvnNum + 1
			if _, err := dao.UserVoucherNum.Edit(ctx, newUserVoucherNum.UvnId, &newUserVoucherNum); err != nil {
				return nil, gerror.New("优惠券领取失败！")
			}
		} else {
			newUserVoucherNum = do.UserVoucherNum{
				ActivityId: activityId,
				UserId:     userId,
				UvnNum:     1,
			}
			if _, err := dao.UserVoucherNum.Add(ctx, &newUserVoucherNum); err != nil {
				return nil, gerror.New("优惠券领取失败！")
			}
		}

		// 更新活动数据
		voucherQuantityUse := gconv.Int(voucherVo.VoucherQuantityUse) + 1
		voucherVo.VoucherQuantityUse = gconv.Uint(voucherQuantityUse)
		voucherVo.VoucherQuantityFree = -voucherQuantityUse
		activityRuleVo.Voucher = voucherVo
		activity := &do.ActivityBase{
			ActivityId:   activityId,
			ActivityRule: gconv.String(activityRuleVo),
		}

		// 判断是否领完
		if gconv.Int(voucherQuantity) <= voucherQuantityUse {
			activity.ActivityState = consts.ACTIVITY_STATE_FINISHED
		}

		if _, err := dao.ActivityBase.Edit(ctx, activityId, activity); err != nil {
			return nil, gerror.New("更新优惠券信息失败！")
		}

		if activityType == consts.GET_VOUCHER_BY_POINT {
			// TODO: 积分操作
		}
	} else {
		return nil, gerror.New("领取数量超限！")
	}

	// 发送消息通知
	/*	messageId := "coupons-to-the-accounts"
		args := g.Map{
			"name":    activityBase.ActivityTitle,
			"endtime": gtime.New(activityBase.ActivityEndtime).Format("Y-m-d H:i:s"),
		}
		if err := service.Message().SendNoticeMsg(ctx, userId, messageId, args); err != nil {
			return nil, gerror.New("发送消息通知失败！")
		}*/

	return voucher, nil
}
