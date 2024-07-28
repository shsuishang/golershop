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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"golershop.cn/utility/mtime"
	"sort"
)

type sUserResource struct{}

func init() {
	service.RegisterUserResource(NewUserResource())
}

func NewUserResource() *sUserResource {
	return &sUserResource{}
}

// Get 读取信息
func (s *sUserResource) Get(ctx context.Context, id any) (out *entity.UserResource, err error) {
	var list []*entity.UserResource
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
func (s *sUserResource) Gets(ctx context.Context, id any) (list []*entity.UserResource, err error) {
	err = dao.UserResource.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sUserResource) Find(ctx context.Context, in *do.UserResourceListInput) (out []*entity.UserResource, err error) {
	out, err = dao.UserResource.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserResource) List(ctx context.Context, in *do.UserResourceListInput) (out *do.UserResourceListOutput, err error) {
	out, err = dao.UserResource.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserResource) Add(ctx context.Context, in *do.UserResource) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserResource.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserResource) Edit(ctx context.Context, in *do.UserResource) (affected int64, err error) {
	_, err = dao.UserResource.Edit(ctx, in.UserId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserResource) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.UserResource.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// InitUserPoints 初始化用户积分
func (s *sUserResource) InitUserPoints(ctx context.Context, userId uint) error {
	// 创建用户资源
	resource := &do.UserResource{
		UserId: userId,
	}
	if _, err := dao.UserResource.Save(ctx, resource); err != nil {
		return err
	}

	// 获取注册赠送积分
	strPointsReg := service.ConfigBase().GetInt(ctx, "points_reg", 0)
	pointsReg := gconv.Float64(strPointsReg)
	desc := fmt.Sprintf("注册赠送积分 %d", int(pointsReg))

	// 记录用户积分变动
	userPointsVo := &model.UserPointsVo{
		UserId:        userId,
		Points:        pointsReg,
		PointsTypeId:  consts.POINTS_TYPE_REG,
		PointsLogDesc: desc,
	}

	if _, err := s.Points(ctx, userPointsVo); err != nil {
		return err
	}

	return nil
}

// InitUserExperience 初始化用户经验等级
func (s *sUserResource) InitUserExperience(ctx context.Context, userId uint) error {
	// 获取注册赠送经验值
	expReg := service.ConfigBase().GetFloat(ctx, "exp_reg", 0.0)
	experienceVo := &model.ExperienceVo{
		UserId:    userId,
		Exp:       expReg,
		ExpTypeId: consts.EXP_TYPE_REG,
		Desc:      "用户注册",
	}
	if err := s.Experience(ctx, experienceVo); err != nil {
		return err
	}

	return nil
}

// Experience 操作用户经验
func (s *sUserResource) Experience(ctx context.Context, experienceVo *model.ExperienceVo) error {
	now := gtime.Now()
	userId := experienceVo.UserId
	expTypeId := experienceVo.ExpTypeId
	exp := experienceVo.Exp

	expKindId := 1
	if exp < 0 {
		expKindId = 2
	}

	// 处理用户等级
	if err := s.handleLevelCode(ctx, expTypeId, userId, now); err != nil {
		return err
	}

	// 写入用户经验记录
	expHistory := &do.UserExpHistory{
		ExpKindId:   expKindId,
		ExpTypeId:   expTypeId,
		UserId:      userId,
		ExpLogValue: int(exp),
		ExpLogDesc:  experienceVo.Desc,
		ExpLogTime:  now,
		ExpLogDate:  now,
	}
	if _, err := dao.UserExpHistory.Save(ctx, expHistory); err != nil {
		return err
	}

	// 更新用户经验值和等级
	if err := s.addExp(ctx, userId, uint(exp)); err != nil {
		return err
	}

	return nil
}

// addExp 增加用户经验值并更新用户等级
func (s *sUserResource) addExp(ctx context.Context, userId, exp uint) error {
	// 获取当前用户经验值
	userResource, err := s.Get(ctx, userId)
	userExp := uint(userResource.UserExp)

	// 更新用户经验记录
	expHistory := &do.UserExpHistory{
		UserExp:     userExp,
		ExpLogValue: exp,
	}
	if _, err := dao.UserExpHistory.Save(ctx, expHistory); err != nil {
		return err
	}

	// 更新用户经验值
	userExp += exp
	userResource.UserExp = uint64(userExp)

	data := &do.UserResource{
		UserId:  userId,
		UserExp: userExp,
	}
	if _, err := dao.UserResource.Save(ctx, data); err != nil {
		return err
	}

	// 更新用户等级
	input := &do.UserLevelListInput{}
	var likes = []*ml.WhereExt{{
		Column: dao.UserLevel.Columns().UserLevelExp,
		Val:    userExp,
		Symbol: ml.LE,
	}, {
		Column: dao.UserLevel.Columns().UserLevelExp,
		Val:    0,
		Symbol: ml.GT,
	}}

	input.WhereExt = likes

	userLevel, err := dao.UserLevel.FindOne(ctx, input)
	if err != nil {
		return err
	}

	userInfo, err := dao.UserInfo.Get(ctx, userId)

	// 判断是否需要更新用户等级
	userLevelId := userInfo.UserLevelId

	if userLevel != nil && userLevelId < userLevel.UserLevelId {
		userLevelId = userLevel.UserLevelId
	}

	// 更新用户等级
	if userLevelId > userInfo.UserLevelId {

	}

	return nil
}

func (s *sUserResource) handleLevelCode(ctx context.Context, expTypeId, userId uint, curDate *gtime.Time) error {
	switch expTypeId {
	case consts.EXP_TYPE_REG:
		// 注册只可以触发一次
		regWrapper := &do.UserExpHistoryListInput{
			Where: do.UserExpHistory{UserId: userId, ExpTypeId: expTypeId},
		}
		regExpLogs, err := dao.UserExpHistory.FindKey(ctx, regWrapper)
		if err != nil {
			return err
		}
		if len(regExpLogs) > 0 {
			return errors.New("注册经验日志已存在")
		}
	case consts.EXP_TYPE_LOGIN:
		// 登录，每天只可以触发一次
		loginWrapper := &do.UserExpHistoryListInput{
			Where: do.UserExpHistory{UserId: userId, ExpTypeId: expTypeId, ExpLogDate: curDate},
		}
		loginExpLogs, err := dao.UserExpHistory.FindKey(ctx, loginWrapper)
		if err != nil {
			return err
		}
		if len(loginExpLogs) > 0 {
			return errors.New("当天登录经验日志已存在")
		}
	case consts.EXP_TYPE_EVALUATE_PRODUCT:
	case consts.EXP_TYPE_EVALUATE_STORE:
	case consts.EXP_TYPE_CONSUME:
	case consts.EXP_TYPE_OTHER:
	case consts.EXP_TYPE_EXCHANGE_PRODUCT:
	case consts.EXP_TYPE_EXCHANGE_VOUCHER:
	default:
	}
	return nil
}

// Points 积分操作
func (s *sUserResource) Points(ctx context.Context, vo *model.UserPointsVo) (flag bool, err error) {
	var pointsKindId int
	date := gtime.Now()

	// 通用判断,注册和
	switch vo.PointsTypeId {
	case consts.POINTS_TYPE_REG:
		// 注册只可以触发一次
		where := &do.UserPointsHistoryListInput{Where: do.UserPointsHistory{UserId: vo.UserId, PointsTypeId: vo.PointsTypeId}}
		count, _ := dao.UserPointsHistory.Count(ctx, where)
		if count > 0 {
			return false, errors.New("已经赠送")
		}
		pointsKindId = consts.POINTS_ADD
	case consts.POINTS_TYPE_LOGIN:
		// 登录，每天只可以触发一次
		where := &do.UserPointsHistoryListInput{Where: do.UserPointsHistory{UserId: vo.UserId, PointsTypeId: vo.PointsTypeId, PointsLogDate: date}}
		num, _ := dao.UserPointsHistory.Count(ctx, where)
		if num > 0 {
			return false, errors.New("已经赠送")
		}
		pointsKindId = consts.POINTS_ADD
	case consts.POINTS_TYPE_EVALUATE_PRODUCT:
		pointsKindId = consts.POINTS_ADD
	case consts.POINTS_TYPE_EVALUATE_STORE:
		pointsKindId = consts.POINTS_ADD
	case consts.POINTS_TYPE_CONSUME:
		pointsKindId = consts.POINTS_ADD
	case consts.POINTS_TYPE_OTHER:
	case consts.POINTS_TYPE_EXCHANGE_PRODUCT:
		pointsKindId = consts.POINTS_MINUS
	case consts.POINTS_TYPE_EXCHANGE_VOUCHER:
		pointsKindId = consts.POINTS_MINUS
	case consts.POINTS_TYPE_EXCHANGE_SP:
		pointsKindId = consts.POINTS_MINUS
	case consts.POINTS_TYPE_TRANSFER_ADD:
		pointsKindId = consts.POINTS_ADD
	case consts.POINTS_TYPE_TRANSFER_MINUS:
		pointsKindId = consts.POINTS_MINUS
	case consts.POINTS_TYPE_CONSUME_RETRUN:
		pointsKindId = consts.POINTS_ADD
	case consts.POINTS_TYPE_FX_FANS:
		pointsKindId = consts.POINTS_ADD
	case consts.POINTS_TYPE_UP_SELLER:
		pointsKindId = consts.POINTS_ADD
	case consts.POINTS_TYPE_DEDUCTION:
		pointsKindId = consts.POINTS_MINUS
	default:
	}

	if pointsKindId == 0 {
		if vo.Points > 0 {
			pointsKindId = consts.POINTS_ADD
		} else {
			pointsKindId = consts.POINTS_MINUS
		}
	}

	// 取正数
	points := vo.Points
	if points < 0 {
		points = -points
	}

	data := &do.UserPointsHistory{
		PointsKindId:    pointsKindId,
		PointsTypeId:    vo.PointsTypeId,
		UserId:          vo.UserId,
		PointsLogPoints: points,
		PointsLogDesc:   vo.PointsLogDesc,
		PointsLogDate:   date,
		PointsLogTime:   date.UnixMilli(),
		StoreId:         vo.StoreId,
		UserIdOther:     vo.UserIdOther,
	}

	return s.addPoints(ctx, data)
}

// AddPoints 插入积分变动记录
func (s *sUserResource) addPoints(ctx context.Context, history *do.UserPointsHistory) (flag bool, err error) {
	userId := history.UserId
	pointsLogPoints := history.PointsLogPoints.(float64)

	// points_log_points均为正数，增减由points_kind_id 控制和判断
	changePoints := pointsLogPoints
	if history.PointsKindId == consts.POINTS_MINUS {
		changePoints = -pointsLogPoints
	}

	// 当前积分
	userResource, err := s.Get(ctx, userId)
	userPoints := userResource.UserPoints

	if changePoints < 0 && userPoints < changePoints {
		panic(errors.New("积分不足！"))
	}

	data := &do.UserResource{UserId: userId}
	data.UserPoints = userResource.UserPoints + changePoints

	if _, err = dao.UserResource.Save(ctx, data); err != nil {
		panic(errors.New("修改积分数据失败！"))
	}

	history.UserPoints = userResource.UserPoints
	if _, err = dao.UserPointsHistory.Save(ctx, history); err != nil {
		panic(errors.New("保存积分日志失败！"))
	}

	return true, err
}

// 获取签到基本信息
func (s *sUserResource) GetSignInfo(ctx context.Context, userId uint) (res *model.SignInfoOutput, err error) {
	// 获取本月开始与结束时间
	monthStartTime, monthEndTime := mtime.LastMonth()

	// 构建查询条件
	input := &do.UserPointsHistoryListInput{
		Where: do.UserPointsHistory{UserId: userId, PointsTypeId: consts.POINTS_TYPE_LOGIN},
	}

	var ext = []*ml.WhereExt{{
		Column: dao.UserPointsHistory.Columns().PointsLogTime,
		Val:    monthStartTime,
		Symbol: ml.GE,
	}, {
		Column: dao.UserPointsHistory.Columns().PointsLogTime,
		Val:    monthEndTime,
		Symbol: ml.LE,
	}}

	input.WhereExt = ext
	input.Sidx = dao.UserPointsHistory.Columns().PointsLogDate
	input.Sort = ml.ORDER_BY_DESC

	pointsHistoryList, err := dao.UserPointsHistory.Find(ctx, input)

	// 今日是否签到
	todayIsSign := 0
	// 已签到日期集合
	signDayArr := make([]string, 0)
	// 连续签到标识
	continueSignDays := 0
	// 判断连续签到情况
	countDaysEnd := true
	// 一天的秒数
	var oneDaySeconds int64 = 24 * 60 * 60
	// 获取当天日期 yyyy-mm-dd
	now := gtime.Now()
	todayDate := now.Format("Y-m-d")

	for _, userPointsHistory := range pointsHistoryList {
		pointsLogDate := userPointsHistory.PointsLogDate.Format("Y-m-d")
		signDayArr = append(signDayArr, pointsLogDate)

		// 判断今日是否已签到
		if todayDate == pointsLogDate {
			todayIsSign = 1
		}

		// 判断连续签到情况
		if countDaysEnd {
			if todayDate != pointsLogDate {
				dayDifference := int((gtime.New(todayDate).Unix() - gtime.New(pointsLogDate).Unix()) / oneDaySeconds)
				if dayDifference == 1 {
					continueSignDays++
				} else {
					countDaysEnd = false
				}
			} else {
				continueSignDays++
			}
			// 更新今天的日期为当前记录的日期
			todayDate = pointsLogDate
		}
	}

	infoRes := &model.SignInfoOutput{
		TodayIsSign:      todayIsSign,
		ContinueSignDays: continueSignDays,
		SignList:         make([]model.PointStepVo, 0),
		SignDayArr:       signDayArr,
	}

	signPointStep := service.ConfigBase().GetStr(ctx, "sign_point_step", "")
	var stepVos []model.PointStepVo
	if !g.IsEmpty(signPointStep) {

		err := json.Unmarshal([]byte(signPointStep), &stepVos)
		if err != nil {
			return nil, err
		}
	}

	infoRes.SignList = s.dealSignPointList(ctx, stepVos)

	return infoRes, err
}

// DealSignPointList 处理前端映射数据
func (s *sUserResource) dealSignPointList(ctx context.Context, stepVos []model.PointStepVo) []model.PointStepVo {
	if len(stepVos) > 0 {
		// 根据 Days 字段排序
		sort.Slice(stepVos, func(i, j int) bool {
			return stepVos[i].Days < stepVos[j].Days
		})

		// 获取配置值
		pointsLogin := service.ConfigBase().GetInt(ctx, "points_login", 0)

		// 是否存在天数为1的数据
		hasOne := false

		// 遍历处理数据
		for i := range stepVos {
			if stepVos[i].Days == 1 {
				hasOne = true
				stepVos[i].ValueStr = fmt.Sprintf("%d.积分", pointsLogin)
			} else {
				stepVos[i].ValueStr = fmt.Sprintf("%d.倍", stepVos[i].Multiples)
			}
		}

		// 如果不存在天数为1的数据，则添加一个
		if !hasOne {
			stepVo := model.PointStepVo{
				Days:     1,
				ValueStr: fmt.Sprintf("%d.积分", pointsLogin),
			}
			stepVos = append(stepVos, stepVo)
		}

		return stepVos
	}

	return nil
}

// 签到
func (s *sUserResource) SignIn(ctx context.Context, userId uint) (flag bool, err error) {
	pointsLogin := service.ConfigBase().GetFloat(ctx, "points_login", 0)
	//expLogin := service.ConfigBase().GetFloat(ctx, "exp_login", 0)

	desc := fmt.Sprintf("签到获取积分 %d", pointsLogin)
	userPointsVo := &model.UserPointsVo{
		UserId:        userId,
		Points:        pointsLogin,
		PointsTypeId:  consts.POINTS_TYPE_LOGIN,
		PointsLogDesc: desc,
	}

	flag, err = service.UserResource().Points(ctx, userPointsVo)

	/*
		experienceVo := &model.ExperienceVo{
			UserId:         userId,
			Experience:     expLogin,
			ExperienceType: consts.Login,
			Description:    "",
		}
		service.UserResource().Experience(ctx, experienceVo)
	*/
	return
}

// GetSignState 当天是否签到
func (s *sUserResource) GetSignState(ctx context.Context, userId uint) (flag bool, err error) {

	// 登录，每天只可以触发一次
	startTime, endTime := mtime.Today()
	input := &do.UserPointsHistoryListInput{}
	input.Where.UserId = userId
	input.Where.PointsTypeId = consts.POINTS_TYPE_LOGIN

	var ext = []*ml.WhereExt{
		{Column: dao.UserPointsHistory.Columns().PointsLogTime, Val: startTime, Symbol: ml.GE},
		{Column: dao.UserPointsHistory.Columns().PointsLogTime, Val: endTime, Symbol: ml.LE},
	}

	input.BaseList.WhereExt = ext

	userPointsHistoryNum, err := dao.UserPointsHistory.Count(ctx, input)

	// 已经签到过
	return userPointsHistoryNum > 0, err
}

func (s *sUserResource) GetList(ctx context.Context, in *do.UserResourceListInput) (out *model.UserResourceOutput, err error) {
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

		out = &model.UserResourceOutput{}
		gconv.Scan(lists, out)
		for i := range out.Items {
			userResource := out.Items[i]
			if len(userInfoMap) > 0 {
				userInfo := userInfoMap[int(userResource.UserId)]
				if userInfo != nil {
					userResource.UserNickname = userInfo.UserNickname
				}
			}
		}
	}

	return out, nil
}
