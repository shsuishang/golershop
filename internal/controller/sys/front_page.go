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

/**
 * 装修页面
 * @author Xinze
 * @since 2021/11/18
 * @File : Page
 */
package sys

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/sys"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
)

// Page 移动导航
func (c *cPage) GetMobileIndexNav(ctx context.Context, req *sys.MobileIndexNavListReq) (out sys.MobileIndexNavListRes, err error) {

	return
}

// GetMobilePage 读取移动页面
func (c *cPage) GetMobilePage(ctx context.Context, req *sys.PageDetailReq) (out sys.PageDetailRes, err error) {
	user := service.BizCtx().GetUser(ctx)

	if req.PageId > 0 {
		// 根据页面编号处理
	} else if req.PageIndex != "" {
		// 根据类型读取pageId
		baseQueryWrapper := &do.PageBaseListInput{}

		//baseQueryWrapper.Eq("subsite_id", subsite_id).Eq("store_id", store_id).Eq("app_id", app_id).Eq("page_type", 3);
		baseQueryWrapper.Where.PageType = 3

		switch req.PageIndex {
		case "page_index":
			baseQueryWrapper.Where.PageIndex = true
		case "page_sns":
			baseQueryWrapper.Where.PageSns = true
		case "page_article":
			baseQueryWrapper.Where.PageArticle = true
		case "page_point":
			baseQueryWrapper.Where.PagePoint = true
		case "page_upgrade":
			baseQueryWrapper.Where.PageUpgrade = true
		case "page_zerobuy":
			baseQueryWrapper.Where.PageZerobuy = true
		case "page_higharea":
			baseQueryWrapper.Where.PageHigharea = true
		case "page_taday":
			baseQueryWrapper.Where.PageTaday = true
		case "page_everyday":
			baseQueryWrapper.Where.PageEveryday = true
		case "page_secondkill":
			baseQueryWrapper.Where.PageSecondkill = true
		case "page_secondday":
			baseQueryWrapper.Where.PageSecondday = true
		case "page_rura":
			baseQueryWrapper.Where.PageRura = true
		case "page_likeyou":
			baseQueryWrapper.Where.PageLikeyou = true
		case "page_exchange":
			baseQueryWrapper.Where.PageExchange = true
		case "page_new":
			baseQueryWrapper.Where.PageNew = true
		case "page_newperson":
			baseQueryWrapper.Where.PageNewperson = true
		}

		pageIds, _ := service.PageBase().FindKey(ctx, baseQueryWrapper)

		if len(pageIds) > 0 {
			req.PageId = gconv.Int64(pageIds[0])
		}
	} else if req.CategoryId > 0 {
		// 根据分类读取pageId
	} else {
		panic("请求数据有误！")
	}

	pageDetail, _ := service.PageBase().Detail(ctx, req.PageId)

	if req.PageIndex != "" && req.PageIndex == "page_index" {
		// 首页弹窗 新人优惠券
		activityBaseListReq := &do.ActivityBaseListInput{}
		activityBaseListReq.Where.ActivityState = consts.ACTIVITY_STATE_NORMAL
		activityBaseListReq.Where.ActivityTypeId = consts.ACTIVITY_TYPE_POP
		activityBaseResIPage, _ := service.ActivityBase().List(ctx, activityBaseListReq)

		if activityBaseResIPage != nil && len(activityBaseResIPage.Items) > 0 {
			activityList := activityBaseResIPage.Items

			// 未登录
			if user == nil {
				pageDetail.PopUps = dealWithPopUp(ctx, activityList, nil)
			} else {
				// 已登录
				userInfo, _ := service.UserInfo().Get(ctx, user.UserId)
				pageDetail.PopUps = dealWithPopUp(ctx, activityList, userInfo)
			}
		}
	}

	gconv.Scan(pageDetail, &out)

	return out, nil
}

func dealWithPopUp(ctx context.Context, activityList []*model.ActivityOutput, userInfo *entity.UserInfo) []*model.PagePopUpVo {
	pagePopUpVos := make([]*model.PagePopUpVo, 0)
	for _, activityBaseRes := range activityList {
		activityRuleJson := activityBaseRes.ActivityRuleJson

		if activityRuleJson != nil {
			popUp := activityRuleJson.Popup

			if !g.IsEmpty(popUp) {
				popUpType := popUp.PopUpType

				if userInfo != nil {
					// 如果用户不符合弹窗等级，过滤此弹窗
					activityUseLevel := activityBaseRes.ActivityUseLevel
					userLevelList := gconv.SliceUint(activityUseLevel)

					if len(userLevelList) > 0 {
						if !array.InArray(userLevelList, userInfo.UserLevelId) {
							continue
						}
					}

					// 如果不是新人，则不展示新人礼包弹窗
					if popUpType == 0 && !userInfo.UserNew {
						continue
					}
				}

				pagePopUpVo := &model.PagePopUpVo{
					PopUpEnable: true,
					PopUpImage:  popUp.PopUpImage,
					PopUpUrl:    popUp.PopUpUrl,
				}
				pagePopUpVos = append(pagePopUpVos, pagePopUpVo)
			}
		}
	}

	return pagePopUpVos
}
