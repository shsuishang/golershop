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

package page

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/pt"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility"
	"golershop.cn/utility/array"
	"math"
)

type sPageBase struct{}

func init() {
	service.RegisterPageBase(New())
}

func New() *sPageBase {
	return &sPageBase{}
}

// Find 查询数据
func (s *sPageBase) Find(ctx context.Context, in *do.PageBaseListInput) (out []*entity.PageBase, err error) {
	out, err = dao.PageBase.Find(ctx, in)

	return out, err
}

// 查询主键
func (s *sPageBase) FindKey(ctx context.Context, in *do.PageBaseListInput) (out []interface{}, err error) {
	return dao.PageBase.FindKey(ctx, in)
}

// List 分页读取
func (s *sPageBase) List(ctx context.Context, in *do.PageBaseListInput) (out *do.PageBaseListOutput, err error) {
	out, err = dao.PageBase.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sPageBase) Add(ctx context.Context, in *do.PageBase) (lastInsertId int64, err error) {
	lastInsertId, err = dao.PageBase.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sPageBase) Edit(ctx context.Context, in *do.PageBase) (affected int64, err error) {
	_, err = dao.PageBase.Edit(ctx, in.PageId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sPageBase) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.PageBase.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetUserCenterMenu
func (s *sPageBase) GetUserCenterMenu(ctx context.Context) (out map[string]interface{}, err error) {
	menus := make(map[string]interface{})

	// 获取配置的appMemberCenter
	appMemberCenter := service.ConfigBase().GetStr(ctx, "app_member_center", "")

	// 如果appMemberCenter不为空
	if appMemberCenter != "" {
		// 解析JSON字符串为结构体
		var jsonObject model.PageMobileVo
		if err := json.Unmarshal([]byte(appMemberCenter), &jsonObject); err != nil {
			return nil, err
		}

		// 获取页面代码
		pageCode := jsonObject.PageCode
		if pageCode != "" {
			// 解析页面代码为Map
			if err := json.Unmarshal([]byte(pageCode), &menus); err != nil {
				return nil, err
			}

			// 过滤不存在的menu
			//pageMenuVoList := menus["list"].([]model.PageMenuVo)
		}
	} else {
		// 如果appMemberCenter为空，则获取所有中心菜单
		menus = s.getAllCenterMenu(ctx)
	}

	// 设置默认类型为2
	if _, ok := menus["type"]; !ok {
		menus["type"] = 2
	}

	return menus, nil

}

func (s *sPageBase) getAllCenterMenu(ctx context.Context) map[string]interface{} {
	menu := make(map[string]interface{})
	list := make([]interface{}, 0)

	// 读取类型
	menu["type"] = 2

	menu["list"] = list

	plantformFxEnable := service.ConfigBase().GetBool(ctx, "plantform_fx_enable", false)

	menu1 := map[string]interface{}{
		"id":         1,
		"name":       "我的拼团",
		"is_show":    true,
		"cat":        1,
		"color":      "#DB384C",
		"icon":       "icon-gouwu",
		"featureKey": "FightGrp",
		"url":        "/activity/fightgroup/order",
	}
	list = append(list, menu1)

	menu4 := map[string]interface{}{
		"id":         36,
		"name":       "售后服务",
		"is_show":    true,
		"cat":        1,
		"color":      "#44afa4",
		"icon":       "zc zc-tuihuanhuo",
		"featureKey": "service",
		"url":        "/member/member/returnlist",
	}
	list = append(list, menu4)

	menu5 := map[string]interface{}{
		"id":         4,
		"name":       "我的砍价",
		"is_show":    true,
		"cat":        1,
		"color":      "#ffc333",
		"icon":       "icon-kanjia",
		"featureKey": "CutPrice",
		"url":        "/activity/cutprice/userlist",
	}
	list = append(list, menu5)

	menu44 := map[string]interface{}{
		"id":         44,
		"name":       "签到",
		"is_show":    true,
		"cat":        1,
		"color":      "#ffc333",
		"icon":       "icon-edit",
		"featureKey": "MemSign",
		"url":        "/member/member/sign",
	}
	list = append(list, menu44)

	menu8 := map[string]interface{}{
		"id":         6,
		"name":       "会员中心",
		"is_show":    true,
		"cat":        1,
		"color":      "#ffc333",
		"icon":       "icon-zuanshi",
		"featureKey": "MemGrade",
		"url":        "/member/member/task",
	}
	list = append(list, menu8)

	menu10 := map[string]interface{}{
		"id":         107,
		"name":       "商品收藏",
		"is_show":    true,
		"cat":        1,
		"color":      "#56ABE4",
		"icon":       "icon-liwu",
		"featureKey": "FavProd",
		"url":        "/member/member/favorites",
	}
	list = append(list, menu10)

	menu11 := map[string]interface{}{
		"id":         108,
		"name":       "我的足迹",
		"is_show":    true,
		"cat":        1,
		"color":      "#56ABE4",
		"icon":       "zc zc-zuji",
		"featureKey": "FavProd",
		"url":        "/member/member/browse",
	}
	list = append(list, menu11)

	menu12 := map[string]interface{}{
		"id":         8,
		"name":       "收货地址",
		"is_show":    true,
		"cat":        1,
		"color":      "#1BC2A6",
		"icon":       "icon-shouhuodizhi",
		"featureKey": "UserAddress",
		"url":        "/member/address/list",
	}
	list = append(list, menu12)

	menu120 := map[string]interface{}{
		"id":         120,
		"name":       "开票信息",
		"is_show":    true,
		"cat":        1,
		"color":      "#1BC2A6",
		"icon":       "zc-caiwukaipiao",
		"featureKey": "UserInvoice",
		"url":        "/member/invoice/list",
	}
	list = append(list, menu120)

	menu121 := map[string]interface{}{
		"id":         120,
		"name":       "我的发票",
		"is_show":    true,
		"cat":        1,
		"color":      "#1BC2A6",
		"icon":       "zc-kaipiao",
		"featureKey": "OrderInvoice",
		"url":        "/member/invoice/order",
	}
	list = append(list, menu121)

	menu14 := map[string]interface{}{
		"id":         21,
		"name":       "推广中心",
		"is_show":    plantformFxEnable,
		"cat":        1,
		"color":      "#327eac",
		"icon":       "zc zc-fenxiao",
		"featureKey": "fenxiao",
		"url":        "/member/fans/index",
	}
	list = append(list, menu14)

	menu17 := map[string]interface{}{
		"id":         32,
		"name":       "帮助",
		"is_show":    true,
		"cat":        1,
		"color":      "#ac8dd5",
		"icon":       "zc zc-bangzhu",
		"featureKey": "Help",
		"url":        "/pagesub/article/list",
	}
	list = append(list, menu17)

	menu21 := map[string]interface{}{
		"id":         11,
		"name":       "清除缓存",
		"is_show":    true,
		"cat":        1,
		"color":      "#DB384C",
		"icon":       "zc zc-qingchuhuancun",
		"featureKey": "CleanCacheKey",
		"url":        "",
	}
	list = append(list, menu21)

	request := g.RequestFromCtx(ctx)

	sourceType := gconv.Uint(request.Get("source_type"))

	liveModeAliyun := service.ConfigBase().GetInt(ctx, "live_mode_aliyun", 0)

	if sourceType != consts.SOURCE_TYPE_H5 && liveModeAliyun == 1 {
		menu22 := map[string]interface{}{
			"id":         33,
			"name":       "我的直播",
			"is_show":    true,
			"cat":        1,
			"color":      "#ac8dd5",
			"icon":       "zc zc-zhibo",
			"featureKey": "Live",
			"url":        "/pagesub/livepush/add",
		}
		list = append(list, menu22)
	}

	pluginPaotui := service.ConfigBase().GetBool(ctx, "Plugin_Paotui", false)
	if pluginPaotui {
		menu23 := map[string]interface{}{
			"id":         109,
			"name":       "骑手大厅",
			"is_show":    false,
			"cat":        2,
			"color":      "#56ABE4",
			"icon":       "zc zc-zuji",
			"featureKey": "FavProd",
			"url":        "/paotui/index/index",
		}
		list = append(list, menu23)
	}

	makeLangPackageEnable := service.ConfigBase().GetBool(ctx, "make_lang_package_enable", false)
	if makeLangPackageEnable {
		menu23 := map[string]interface{}{
			"id":         35,
			"name":       "翻译制作",
			"is_show":    true,
			"cat":        2,
			"color":      "#ac8dd5",
			"icon":       "zc zc-zhibo",
			"featureKey": "ReloadLang",
			"url":        "",
		}
		list = append(list, menu23)
	}

	liveModeXcx := service.ConfigBase().GetInt(ctx, "live_mode_xcx", 0)
	if liveModeXcx == 1 {
		menu25 := map[string]interface{}{
			"id":         109,
			"name":       "申请主播",
			"is_show":    true,
			"cat":        2,
			"color":      "#56ABE4",
			"icon":       "zc zc-15",
			"featureKey": "FavProd",
			"url":        "/xcxlive/anchor/apply",
		}
		list = append(list, menu25)

		menu26 := map[string]interface{}{
			"id":         109,
			"name":       "创建房间",
			"is_show":    true,
			"cat":        2,
			"color":      "#56ABE4",
			"icon":       "zc zc-fangjian",
			"featureKey": "FavProd",
			"url":        "/xcxlive/room/add",
		}
		list = append(list, menu26)

		menu27 := map[string]interface{}{
			"id":         109,
			"name":       "房间列表",
			"is_show":    true,
			"cat":        2,
			"color":      "#56ABE4",
			"icon":       "zc zc-fenlei1",
			"featureKey": "FavProd",
			"url":        "/xcxlive/room/list",
		}
		list = append(list, menu27)
	}

	menu["list"] = list

	return menu
}

// Detail 获取页面详情
func (s *sPageBase) Detail(ctx context.Context, pageId any) (out *model.PageDetail, err error) {

	pageBaseRes := &model.PageDetail{}
	pageBase, _ := dao.PageBase.Get(ctx, pageId)

	if pageBase != nil {
		gconv.Scan(pageBase, pageBaseRes)
	}

	// 修复数据
	fixData(ctx, pageBaseRes)

	return pageBaseRes, nil
}

// fixData 修复数据
func fixData(ctx context.Context, pageBaseRes *model.PageDetail) {
	// 获取IM配置
	imEnable := service.ConfigBase().GetBool(ctx, "im_enable", false)
	imUserId := service.ConfigBase().GetUint(ctx, "site_im", 10001)

	// 获取当前登录用户
	user := service.BizCtx().GetUser(ctx)

	// 构建IM配置对象
	imConfigVo := &model.ImConfigVo{
		Puid:     0,
		ImEnable: imEnable,
		ImUserId: imUserId,
	}
	pageBaseRes.Im = imConfigVo

	if user != nil {
		serviceUserId := service.ConfigBase().GetStr(ctx, "service_user_id", "")
		imConfigVo.Puid = utility.GetPlantformUid(serviceUserId, gconv.String(user.UserId))
	}

	// 处理页面编码
	pageCode := pageBaseRes.PageCode

	var pageCodeRows []interface{}

	/*
		if err := json.Unmarshal([]byte(pageCode), &pageCodeRows); err != nil {
			fmt.Println("解析 JSON 失败:", err)
			return
		}
	*/

	content, _ := gjson.LoadContent(pageCode)
	pageCodeRows = content.Array()

	itemIdRow := make([]uint64, 0)

	for _, pageCodeRow := range pageCodeRows {
		pageCodeMap, ok := pageCodeRow.(map[string]interface{})

		if !ok {
			fmt.Println("Data is not a map[string]interface{}")
			return
		}

		eltmType, _ := pageCodeMap["eltmType"].(int)

		if eltmType == 4 {
			eltm4 := pageCodeMap["eltm4"].(map[string]interface{})
			if eltm4 == nil {
				continue
			}

			data := eltm4["data"].([]map[string]interface{})
			for _, datum := range data {
				did := datum["did"].(uint64)
				if !g.IsEmpty(did) {
					itemIdRow = append(itemIdRow, did)
				}
			}
		} else if eltmType == 16 {
			eltm16 := pageCodeMap["eltm16"].(map[string]interface{})
			if eltm16 == nil {
				continue
			}

			data := eltm16["data"].([]map[string]interface{})
			for _, item := range data {
				did := item["did"].(uint64)
				if !g.IsEmpty(did) {
					itemIdRow = append(itemIdRow, did)
				}
			}
		}
	}

	currencyExchangeRate := 1.0

	if len(itemIdRow) > 0 {
		// 获取产品列表
		productItemInput := &pt.ItemListReq{
			ItemId:         gstr.JoinAny(itemIdRow, ","),
			ItemEnable:     consts.PRODUCT_STATE_NORMAL,
			ProductStateId: consts.PRODUCT_STATE_NORMAL,
		}

		itemOutputList, _ := service.ProductIndex().ListItem(ctx, productItemInput)
		itemIdss := make([]uint64, len(itemOutputList.Items))
		for i, v := range itemOutputList.Items {
			itemIdss[i] = v.ItemId
		}

		// 获取活动数据
		activityItemInput := &do.ActivityItemListInput{
			Where: do.ActivityItem{
				ItemId:            itemIdRow,
				ActivityItemState: consts.ACTIVITY_STATE_NORMAL,
			},
		}

		activityItemList, _ := dao.ActivityItem.Find(ctx, activityItemInput)

		// 遍历页面代码行
		for _, pageCodeRow := range pageCodeRows {
			pageCodeMap, ok := pageCodeRow.(map[string]interface{})

			if !ok {
				fmt.Println("Data is not a map[string]interface{}")
				return
			}

			eltmType, _ := pageCodeMap["eltmType"].(int)

			if eltmType == 4 {
				eltm4 := gconv.Map(pageCodeMap["eltm4"])
				if eltm4 == nil {
					continue
				}
				//data := gconv.SliceAny(eltm4["data"])
				data := eltm4["data"].([]map[string]interface{})

				// 将 JSON 字符串解析为 PageDataItemVo 列表
				as := make([]*model.PageDataItemVo, 0)
				//gconv.Scan(data, &as)
				err := gjson.DecodeTo(data, &as)
				if err != nil {
					// 解析失败，处理错误
				}

				filtered := make([]*model.PageDataItemVo, 0)
				for _, item := range as {
					if array.InArray(itemIdss, item.Did) {
						filtered = append(filtered, item)
					}
				}

				data = gconv.SliceMap(filtered)

				//todo 判断对象是否跟随更改！！！！
				//eltm4["data"] = data
				// 设置数据到 eltm4 中
				//eltm4.Set("data", data)
				//pageCodeRow.Set("eltm4", eltm4)

				for _, item := range data {
					did := item["did"]

					for _, productItemRow := range itemOutputList.Items {
						if productItemRow.ItemId == did {
							itemSalePrice := gconv.Float64(item["ItemSalePrice"])
							item_unit_price := productItemRow.ItemUnitPrice
							item_unit_points := productItemRow.ItemUnitPoints
							item_unit_sp := 0.0

							var activity_type_id uint
							for _, activity_item_row := range activityItemList {
								if activity_item_row.ItemId == did {
									activity_type_id = activity_item_row.ActivityTypeId
									break
								}
							}

							//秒杀
							if consts.ACTIVITY_TYPE_LIMITED_DISCOUNT == activity_type_id {

							}

							int_item_unit_price := 0.0
							if item_unit_price > 0 {
								int_item_unit_price = item_unit_price * currencyExchangeRate
							}
							round_ItemSalePrice := math.Round(itemSalePrice * currencyExchangeRate * 100 / 100)
							int_item_unit_points := 0
							if item_unit_points > 0 {
								int_item_unit_points = int(item_unit_points * currencyExchangeRate)
							}
							int_item_unit_sp := 0
							if item_unit_sp > 0 {
								int_item_unit_sp = int(item_unit_sp * currencyExchangeRate)
							}

							item["item_unit_price"] = int_item_unit_price
							item["ItemSalePrice"] = round_ItemSalePrice
							item["item_unit_points"] = int_item_unit_points
							item["item_unit_sp"] = int_item_unit_sp
						}
					}
				}
			} else if eltmType == 13 {
				eltm13 := gconv.Map(pageCodeMap["eltm13"])
				if eltm13 == nil {
					continue
				}
				data := gconv.SliceMap(eltm13["data"])
				for _, o := range data {
					jsonObject := gconv.Map(o)
					_type := gconv.String(jsonObject["type"])
					if _type == "1" {
						//placeholderText := gconv.String(jsonObject["placeholderText"])
						// TODO: 处理业务逻辑
					}
				}
			} else if eltmType == 16 {
				eltm16 := gconv.Map(pageCodeMap["eltm16"])
				if eltm16 == nil {
					continue
				}
				data := gconv.SliceMap(eltm16["data"])
				for _, item := range data {
					did := gconv.Uint64(item["did"])
					for _, activity_item_row := range activityItemList {
						if activity_item_row.ItemId == did {
							//itemSalePrice := gconv.Float64(item["ItemSalePrice"])
							activity_item_price := activity_item_row.ActivityItemPrice
							selectType := gconv.Int(item["selectType"])
							if selectType == 14 {
								item["ItemSalePrice"] = activity_item_price
							}
						}
					}
				}
			} else if eltmType == 104 {
				eltm104 := gconv.Map(pageCodeMap["eltm104"])
				if eltm104 == nil {
					continue
				}
				data := gconv.SliceMap(eltm104["data"])

				var item_id_104 []uint64
				for _, datum := range data {
					jsonObject := gconv.Map(datum)
					ids := gconv.SliceUint64(jsonObject["ids"])
					item_id_104 = append(item_id_104, ids...)
				}

				itemIds, _ := dao.ProductItem.FindKey(ctx, &do.ProductItemListInput{Where: do.ProductItem{
					ItemId:     item_id_104,
					ItemEnable: consts.PRODUCT_STATE_NORMAL,
				}})

				data1 := garray.New()
				for _, datum := range data {
					jsonObject := gconv.Map(datum)
					ids := gconv.SliceUint64(jsonObject["ids"])
					var id_s []uint64
					for _, id := range ids {
						for _, itemId := range itemIds {
							if id == itemId.(uint64) {
								id_s = append(id_s, id)
							}
						}
					}

					jsonObject["ids"] = gconv.String(id_s)
					data1.Append(jsonObject)
				}

				eltm104["data"] = data1.Slice()
				pageCodeMap["eltm104"] = eltm104

				/*
					for _, item := range data {
						did := gconv.Int64(item["did"])
						// TODO: 处理业务逻辑
					}
				*/
			}
		}
	}

	// 设置页面加载状态
	pageBaseRes.PageLoaded = len(pageCodeRows) > 0

	// 将修复后的页面编码转换为字符串
	pageBaseRes.PageCode, _ = gjson.EncodeString(pageCodeRows)
}
