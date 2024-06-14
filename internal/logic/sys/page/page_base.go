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
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/api/sys"
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

// GetDataInfo 获取数据信息
func (s *sPageBase) GetDataInfo(ctx context.Context, pageDataReq *sys.PageBaseGetDataInfoReq) (*sys.PageBaseGetDataInfoRes, error) {
	baseListRes := &sys.PageBaseGetDataInfoRes{}
	pageDataItemVoList := make([]*model.PageDataItemVo, 0)

	// 检查类型是否为空
	if pageDataReq.Type == 0 {
		return nil, gerror.New("请求参数不能为空")
	}

	switch pageDataReq.Type {
	case 1, 104:
		// 商品
		input := &pt.ItemListReq{}

		if !g.IsEmpty(pageDataReq.Name) {
			input.ProductNameIndex = pageDataReq.Name
		}

		input.Page = pageDataReq.Page
		input.Size = pageDataReq.Size

		productItemRes, err := service.ProductIndex().ListItem(ctx, input)
		if err != nil {
			return nil, err
		}

		baseListRes.Total = productItemRes.Total
		baseListRes.Size = productItemRes.Size
		baseListRes.Records = productItemRes.Records
		baseListRes.Page = productItemRes.Page

		gconv.Structs(productItemRes, baseListRes)

		productItemList := productItemRes.Items

		if len(productItemList) > 0 {
			for _, productIndex := range productItemList {
				pageDataItemVo := &model.PageDataItemVo{
					ProductTips:   productIndex.ProductTips,
					Id:            productIndex.ItemId,
					Name:          fmt.Sprintf("%s%s", productIndex.ProductName, productIndex.ItemName),
					MarketPrice:   gconv.Float64(productIndex.ItemMarketPrice),
					ItemSalePrice: gconv.Float64(productIndex.ItemUnitPrice),
					Path:          productIndex.ProductImage,
				}

				pageDataItemVoList = append(pageDataItemVoList, pageDataItemVo)
			}

			baseListRes.Items = pageDataItemVoList
		}

	case 2:
		// 店铺分类
		productCategoryQueryWrapper := &do.ProductCategoryListInput{}

		if pageDataReq.Name != "" {
			productCategoryQueryWrapper.Where.CategoryName = "%" + pageDataReq.Name + "%"
		}
		productCategoryQueryWrapper.Page = pageDataReq.Page
		productCategoryQueryWrapper.Size = pageDataReq.Size
		productCategoryPage, err := dao.ProductCategory.List(ctx, productCategoryQueryWrapper)
		if err != nil {
			return nil, err
		}

		gconv.Struct(productCategoryPage, baseListRes)

		if len(productCategoryPage.Items) > 0 {
			for _, item := range productCategoryPage.Items {
				pageDataItemVo := &model.PageDataItemVo{}
				pageDataItemVo.Id = uint64(item.CategoryId)
				pageDataItemVo.Path = item.CategoryImage
				pageDataItemVo.Name = item.CategoryName
				pageDataItemVoList = append(pageDataItemVoList, pageDataItemVo)
			}
			baseListRes.Items = pageDataItemVoList
		}
	case 3:
		// APP
		strJsonData := "[{\"id\":872,\"name\":\"<section  ><section style=\\\"margin:5px 0;box-sizing: border-box;\\\"><section  style=\\\"width:96%;clear:both;overflow:hidden;margin:0 auto;background-color:#8787B5;\\\"><section style=\\\"width:55%;float:left;overflow:hidden;\\\"><img   style=\\\"max-width:100%;float:left;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/0d119208-64c3-4a75-988f-745c7b0241fe1.jpg\\\"></section><section style=\\\"display:inline-block;width:45%;float:right;\\\"><section style=\\\"margin-right:0.3em;margin-top:30px;padding:0.3em 0.5em;color:#FFFFFF;font-size:0.9em;font-family:inherit;font-weight:inherit;text-align:center;text-decoration:inherit;\\\"><p>燕飞蝉寒秋叶黄，</p><p>雀叫枣红荷叶清。</p><p>碧云蓝天气色浓，</p><p>转身回头又风景。</p></section></section></section><section style=\\\"width: 0px; height: 0px; clear: both;\\\"></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null,\"RowNumber\":1,\"Total\":190},{\"RowNumber\":2,\"Total\":190,\"id\":871,\"name\":\"<section  ><section style=\\\"margin:5px 0;box-sizing: border-box;\\\"><section style=\\\"margin:0;padding:0;box-sizing:border-box;padding-bottom:40px;\\\"><section style=\\\"text-align:center;color:inherit;\\\"><section style=\\\"color:inherit;width:320px;display:inline-block;\\\"><img   style=\\\"color:inherit;width:100%;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/be76aed9-a53a-455e-ae21-8464a2951d702.jpg\\\"></section></section><section style=\\\"text-align:center;color:inherit;margin-top:-110px;\\\"><section style=\\\"display:inline-block;color:inherit;\\\"><section  style=\\\"height:0px;border-style:solid;border-width:0px 0 70px 90px;border-color:transparent transparent #292929;color:inherit;float:left;margin-left:20px;\\\"></section><section  style=\\\"float:left;height:70px;background-color:#292929;width:190px;color:#ffffff;font-size:18px;text-align:center;padding-top:6px;\\\">离时髦忽远又忽近？</section></section></section><section style=\\\"margin-top:-40px;text-align:center;margin-bottom:12px;color:inherit;\\\"><section style=\\\"width:320px;display:inline-block;padding-left:20px;color:inherit;\\\"><section style=\\\"color:#F7F9F6;border-color:#80B135;text-align:left;margin-left:40px;\\\"><span style=\\\"border-color:#FDFDFD;color:inherit;font-size:18px;\\\">“文青风”不是你想穿就能穿！</span></section></section></section></section><section style=\\\"width: 0px; height: 0px; clear: both;\\\"></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":3,\"Total\":190,\"id\":870,\"name\":\"<section  ><section style=\\\"border: 20px solid #f96e57;-webkit-border-image: url(http://files.qiluzhaoshang.com//fck007/2017042015/c74022df-eb0c-481c-a5d5-7c8d2eda39a132.png) 20 20;-moz-border-image: url(http://files.qiluzhaoshang.com//fck007/2017042015/c74022df-eb0c-481c-a5d5-7c8d2eda39a132.png) 20 20;-ms-border-image: url(http://files.qiluzhaoshang.com//fck007/2017042015/c74022df-eb0c-481c-a5d5-7c8d2eda39a132.png) 20 20;padding: 0;margin: 0;\\\"><p style=\\\"text-align:center;white-space:normal\\\"><img   style=\\\"width:100%;margin:0;height:auto!important\\\" height=\\\"auto\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/aba52f28-0c27-44dd-bc6c-6dbeeea869cd3.jpg\\\"></p></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":4,\"Total\":190,\"id\":869,\"name\":\"<section  ><section style=\\\"box-sizing:border-box;margin:5px;\\\"><img  style=\\\"width:200px;float:left;margin-right:10px;margin-bottom:5px\\\"  src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/8f61b9f8-3e2b-45e6-9dc0-f1b39822d93a4.jpg\\\"><section style=\\\"font-size:14px;font-family:inherit;line-height:30px;text-decoration:inherit\\\"><section>人没安全感，总会不确定。试着去接受一个爱你的人，也只有在与之相处的过程里你才会体会到，这到底是怎样的爱。过度的怕和试探忐忑的矫情，也许会让你看的更清透，但也同样会让你错过感情最美的时候，情感的递增需要一个过程，而它的开始是心动。当有一天你懂了，那时你失去的也许是你的全世界</section></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":5,\"Total\":190,\"id\":868,\"name\":\"<section  ><section style=\\\"background:#fff;text-align:center;border-style:none;clear:both;overflow:hidden;margin:5px 0;\\\"><span  style=\\\"padding: 5px; margin-left: 6px; border: 1px solid rgb(95, 170, 255); float: right;\\\"><img   style=\\\"display:block;width:200px;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/7838ecdd-96a4-4012-b9f0-00b5445de3bd5.jpg\\\"></span><section style=\\\"line-height:1.5;text-align:left;font-size:14px\\\"><p style=\\\"display:inline\\\">智慧本身就是好的。有一天我们都会死去，追求智慧的道路还会有人在走着。死掉以后的事我看不到。但在我活着的时候，想到这件事，心里就很高兴。</p></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":6,\"Total\":190,\"id\":867,\"name\":\"<section  ><section style=\\\"background:#fff;text-align:center;border-style:none;clear:both;overflow:hidden;margin:5px 0;\\\"><span  style=\\\"margin-right: 6px; padding: 5px; border: 1px solid rgb(95, 170, 255); float: left;\\\"><img   style=\\\"display:block;width:200px\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/67863ba5-d73b-48b2-85cd-7634ac11502b6.jpg\\\"></span><section style=\\\"line-height:30px;text-align:left;font-size:14px\\\"><p style=\\\"display:inline\\\">&nbsp;智慧本身就是好的。有一天我们都会死去，追求智慧的道路还会有人在走着。死掉以后的事我看不到。但在我活着的时候，想到这件事，心里就很高兴。</p></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":7,\"Total\":190,\"id\":866,\"name\":\"<section  ><section style=\\\"margin-top:.5em;margin-bottom:.5em;box-sizing:border-box\\\"><section style=\\\"overflow:hidden\\\"><section  style=\\\"width: 10em; height: 2em; line-height: 2em; margin-top: 1em; margin-bottom: -4em; margin-right: -3em; -webkit-transform: rotate(45deg); font-size: 1.5em; font-family: inherit; font-weight: inherit; text-align: center; text-decoration: inherit; color: rgb(255, 255, 255); border-color: rgb(255, 255, 255); box-sizing: border-box; float: right; background-color: rgb(95, 170, 255);\\\"><section style=\\\"box-sizing:border-box\\\">咖啡物语</section></section><img  style=\\\"box-sizing:border-box;width:100%\\\"  src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/1ec3b5ed-9640-475e-bd31-72ddb95b8f3c7.jpg\\\"></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":8,\"Total\":190,\"id\":865,\"name\":\"<section  ><section style=\\\"background:#fff;text-align:center;overflow:hidden;margin: 10px auto;display: -webkit-box;display: -ms-flexbox;display: -webkit-flex;display: flex;-webkit-flex-wrap: nowarp;-ms-flex-wrap: nowarp;flex-wrap: nowarp;\\\"><span style=\\\"display: block;padding:3px;border:solid 1px #bfbfbf;margin-right:10px;-webkit-box-flex: 1;-webkit-flex: auto;-ms-flex: auto;flex: auto;\\\"><img   style=\\\"display:block;width: 100%;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/b6707719-3a67-4b52-b2ee-5662d6e802ea9.jpg\\\"></span><span style=\\\"padding:3px;display: block;border:solid 1px #bfbfbf;-webkit-box-flex: 1;-webkit-flex: auto;-ms-flex: auto;flex: auto;\\\"><img   style=\\\"display:block;width: 100%;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/054a991d-6b84-41c3-8aa8-4f4a8fedc81d8.jpg\\\"></span></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":9,\"Total\":190,\"id\":864,\"name\":\"<section  ><section style=\\\"margin-top: 0.5em; margin-bottom: 0.5em;box-sizing: border-box;\\\"><section style=\\\"margin: 3px; box-sizing: border-box; padding: 0px;\\\"><p style=\\\"text-align: center; box-sizing: border-box; color: inherit;\\\"><img   style=\\\"box-sizing: border-box; margin: 0px; padding: 0px; width: 100%; color: inherit;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/4458d6f2-64b2-4a3a-abe1-7ba9894a4d9c10.jpg\\\"></p><section style=\\\"padding: 2px 0px; box-sizing: border-box; margin: 0px; color: inherit;\\\"><section style=\\\"float: left; margin-right: 20px; margin-left: 5px; box-sizing: border-box; padding: 0px; color: inherit;\\\"><span style=\\\"box-sizing: border-box; color: rgb(216, 40, 33); font-size: 30px; margin: 0px; padding: 0px; border-color: rgb(216, 40, 33);\\\"><em  style=\\\"box-sizing: border-box; padding: 0px; margin: 0px; border-color: rgb(216, 40, 33); color: inherit;\\\">1</em></span><span style=\\\"box-sizing: border-box; font-size: 14px; margin: 0px; padding: 0px; color: inherit;\\\"><em style=\\\"box-sizing: border-box; padding: 0px; margin: 0px;color:rgb(153,153,153)\\\">/6</em></span></section><section style=\\\"padding: 5px 0px; box-sizing: border-box; margin-top: 5px; color: inherit;\\\"><p style=\\\"clear: none; font-size: 12px; line-height: 17px; box-sizing: border-box; padding: 0px; margin: 0px; color: inherit;\\\"><span style=\\\"box-sizing:border-box; color:rgb(165, 165, 165); margin:0px; padding:0px\\\">我要你知道,在这个世界上总有一个人是等着你的,不管在什么时候,不管在什么地方,反正你知道,总有这么个人。—— 张爱玲</span></p></section></section></section><section style=\\\"display: block; width: 0; height: 0; clear: both;\\\"></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":10,\"Total\":190,\"id\":863,\"name\":\"<section  ><section style=\\\"margin: 5px 0;box-sizing: border-box;\\\"><section style=\\\"border: 0px rgb(145, 168, 252);box-sizing: border-box;width: 100%;clear: both;padding: 0px 0.5em 0.5em 0px;text-align: center;\\\"><img   style=\\\"border-radius: 50%;box-sizing: border-box;vertical-align: baseline;width: 222px;height: 222px !important;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/59683df0-1688-4877-b3b4-a2624ffa78cd11.jpg\\\"><section style=\\\"max-width: 100%;margin: -3.2em 0 0 0;box-sizing: border-box;\\\"><section  style=\\\"display: inline-block;height: 45px;vertical-align: top;border-right-width: 21px;border-right-style: solid;border-right-color: rgb(95, 170, 255);box-sizing: border-box !important;border-top-width: 21px !important;border-top-style: solid !important;border-top-color: transparent !important;border-bottom-width: 21px !important;border-bottom-style: solid !important;border-bottom-color: transparent !important;\\\"></section><section  style=\\\"height: 45px;width: 192px;display: inline-block;color: rgb(255, 255, 255);font-size: 16px;font-weight: bold;padding: 4px 10px;line-height: 36px;vertical-align: middle;border-color: rgb(245, 248, 254);box-sizing: border-box !important;background-color: rgb(95, 170, 255);\\\"><span style=\\\"border-color: rgb(145, 168, 252);box-sizing: border-box;font-size: 16px;\\\">我是樱桃小丸子</span></section><section  style=\\\"display: inline-block;height: 45px;vertical-align: top;border-left-width: 22px;border-left-style: solid;border-left-color: rgb(95, 170, 255);box-sizing: border-box !important;border-top-width: 22px !important;border-top-style: solid !important;border-top-color: transparent !important;border-bottom-width: 22px !important;border-bottom-style: solid !important;border-bottom-color: transparent !important;\\\"></section></section></section><section style=\\\"display: block; width: 0; height: 0; clear: both;\\\"></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":11,\"Total\":190,\"id\":862,\"name\":\"<section  ><section style=\\\"border:none;border-style:none;margin: 5px 0;text-align:center;\\\"><span  style=\\\"width: 0px; height: 0px; border-style: solid; border-width: 1.5em 1em 1em; border-color: rgb(58, 188, 255) transparent transparent; display: inline-block;\\\"></span><span  style=\\\"width: 0px; height: 0px; border-style: solid; border-width: 1.3em 0.8em 0.8em; border-color: rgb(58, 188, 255) transparent transparent; margin: 0px auto; display: block;\\\"></span><span  style=\\\"padding: 0.5em; border: 1px solid rgb(58, 188, 255); border-radius: 50%; display: inline-block;\\\"><img   style=\\\"width: 14em;height: 14em; border-radius: 50%; display: block;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/fa44985f-9f16-4221-a32b-ae4ffcc3d59712.jpg\\\"></span><span style=\\\"padding: 0.5em 0;font-size: 1.25em;display: block;\\\"><section style=\\\"padding: 0 0.5em;display: inline-block;\\\"><p style=\\\"margin: 0\\\">面朝大海&nbsp;|&nbsp;春暖花开</p></section></span></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":12,\"Total\":190,\"id\":861,\"name\":\"<section  ><section style=\\\"margin: 5px 0;text-align:center;\\\"><section  style=\\\"width: 0px; height: 0px; margin: 0px auto; border-style: solid; border-width: 1.2em 1.2em 1.8em; border-color: transparent transparent rgb(58, 188, 255); display: block;\\\"><section style=\\\"width: 0;height: 0;border: solid 0.5em transparent;border-bottom: solid 1em #ffffff;margin-left: -0.5em;display: block;\\\"></section></section><section  style=\\\"width: 10em; margin: -0.6em auto 0px; border-top-style: solid; border-top-width: 1px; border-color: rgb(58, 188, 255); display: block; color: rgb(58, 188, 255);\\\"></section><section style=\\\"display: block;margin-top: 1.5em;\\\"><img   style=\\\"width: 8.5em;margin-right: 1em;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/cccb4867-d8ee-4ed4-99a5-50dbb573182014.jpg\\\"><img   style=\\\"width:8.5em\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/fcd696ed-5a10-41ab-a27a-58bbcbf4b8ca13.jpg\\\"></section><section  style=\\\"border-top-style: solid; border-top-width: 1px; border-color: rgb(58, 188, 255); width: 5em; color: rgb(58, 188, 255); margin-top: 0.5em; font-size: 2em; height: 0.5em; line-height: 0.5em; display: inline-block;\\\">....</section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":13,\"Total\":190,\"id\":860,\"name\":\"<section  ><section style=\\\"border:none;border-style:none;margin: 1em auto;text-align:center;\\\"><section  style=\\\"width: 16em;border-top: 2px solid rgb(58, 188, 255);display: inline-block;\\\"><section style=\\\"height: 1.4em;line-height: 1.4em;margin-top: -0.9em;display: block;\\\"><section  style=\\\"font-size: 1.25em; color: rgb(58, 188, 255); min-width: 6em; display: inline-block;\\\"><section style=\\\"font-size:14px;margin-right:-1px;display:inline-block;\\\">●</section><p style=\\\"margin: 0;display:inline-block;padding:0 10px;background-color: #ffffff;\\\">动静之间</p><section style=\\\"font-size: 14px;margin-left:-2px;display:inline-block;\\\">●</section></section></section></section><section style=\\\"margin: 1em auto;\\\"><img   style=\\\"display: inline-block;vertical-align: top;width: 20em;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/2f4ee319-bb38-4de7-b6ed-215b3652899316.jpg\\\"></section><section style=\\\"\\\"><img   style=\\\"width: 20em;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/74c3e9b2-f770-4c55-8954-a973e7a3f16515.png\\\"></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":14,\"Total\":190,\"id\":859,\"name\":\"<section  ><section style=\\\"border:none;border-style:none;margin: 1em auto 1em;text-align:center;width: 20em;color: #000000;\\\"><span style=\\\"display:block;\\\"><img   style=\\\"width: 20em;vertical-align: middle;float: left;display:inline-block;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/4432d79c-ca72-4e6b-8d2a-2f9ed4d4fa7417.jpg\\\"><span style=\\\"background-color: #ffffff;opacity: 0.5;width: 8.5em;height: 8.5em;line-height: 8.5em;overflow: hidden;border-radius: 50%;vertical-align: middle;margin-top: -13em;display: inline-block;\\\"><section style=\\\"font-weight: bold;margin-right: 1em;display: inline-block;vertical-align: middle;font-size: 1.25em;\\\"><p style=\\\"margin: 0;\\\">慢</p></section><section style=\\\"width: 1em;margin: 0.5em auto;display: inline-block;vertical-align: middle;font-size: 1em;\\\"><p style=\\\"margin: 0;line-height:1.1;\\\">是一种生活态度</p></section></span></span></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":15,\"Total\":190,\"id\":858,\"name\":\"<section  ><section style=\\\"width: 100%;clear: both;overflow: hidden;margin: 10px auto;\\\"><section style=\\\"width: 50%; float: left;\\\"><img style=\\\"width: 100%  !important;height: 201px !important;\\\"   src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/b2521319-647d-4641-a4e5-881e7f2b53b518.jpg\\\"></section><section style=\\\"display: inline-block; width: 50%;height: 201px; float: right;background: #68744e;\\\"><section style=\\\"margin-right: 0.3em;margin-top: 30px; padding: 0.3em 0.5em; color: rgb(255, 255, 255); font-size: 0.9em; font-family: inherit; font-weight: inherit; text-align: center; text-decoration: inherit;\\\"><p>孤独与否</p><p>岁月依旧</p><p>红了樱桃</p><p>绿了芭蕉</p></section></section><p></p></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":16,\"Total\":190,\"id\":857,\"name\":\"<section  ><section style=\\\"border:none;margin:.5em 0;box-sizing:border-box;padding:0;font-family:微软雅黑;font-size:14px\\\"><section  style=\\\"border-radius: 0.8em; width: 100%; border: 2px solid rgb(95, 170, 255); box-sizing: border-box;\\\"><section  style=\\\"border-radius: 0.8em; width: 100%; text-align: center; display: table; padding: 10px; box-sizing: border-box; background-color: rgb(95, 170, 255);\\\"><section style=\\\"display:table-cell;vertical-align:middle;min-height:4em;width:100%;height:100%;padding:10px;line-height:1.2;border:2px dotted #fff;font-family:inherit;font-weight:inherit;text-decoration:inherit;color:#fff;box-sizing:border-box;background-color:transparent\\\"><section style=\\\"width:7em;height:7em;border:5px solid #ffcf2d;border-radius:100%;margin:20px auto;box-sizing:border-box\\\"><img  style=\\\"box-sizing:border-box;width:100%;height:100%;border-radius:100%;background-image:url(http://files.qiluzhaoshang.com//fck007/2017042015/fbc90ab8-7e07-4088-9415-880e425a670b19.png);background-size:cover;background-position:50% 50%;background-repeat:no-repeat\\\" ></section><section style=\\\"box-sizing:border-box\\\"><section style=\\\"box-sizing:border-box\\\">为保证效果，强烈建议更换图片为正方形</section></section></section></section></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":17,\"Total\":190,\"id\":856,\"name\":\"<section  ><section style=\\\"border:0 none;padding:0;box-sizing:border-box;margin:0;font-family:微软雅黑\\\"><section style=\\\"margin:.5em 0 1em;padding:0;box-sizing:border-box;min-width:0;color:#3e3e3e;font-size:15px;word-wrap:break-word!important\\\"><section style=\\\"text-align:right;box-sizing:border-box;padding:0;margin:0;color:inherit;font-size:14px\\\"><section style=\\\"margin-right:15px;padding:0;box-sizing:border-box;display:inline-block;vertical-align:top;height:6em;width:6em;border-top-left-radius:50%;border-top-right-radius:50%;border-bottom-right-radius:50%;border-bottom-left-radius:50%;border:5px solid rgba(0,0,0,.2);font-family:inherit;font-weight:inherit;text-decoration:inherit;font-size:1.6em;color:inherit;word-wrap:break-word!important\\\"><img   style=\\\"border-radius:50%;box-sizing:border-box;color:inherit;display:inline-block\\\" width=\\\"100%\\\" height=\\\"100%\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/649b0ef3-b833-4e9f-9f78-8729a0936db320.jpg\\\"></section></section><section  style=\\\"margin: -8.5em 0px 1em; padding: 10px 50% 10px 15px; border: 0px solid rgb(73, 158, 243); font-size: 14px; font-weight: inherit; text-decoration: inherit; color: rgb(255, 255, 255); box-sizing: border-box; overflow: hidden; min-height: 105px; background-color: rgb(95, 170, 255);\\\"><p style=\\\"color:inherit;white-space:normal;line-height:2em\\\"><span style=\\\"color:#fff\\\">输入标题</span></p><p style=\\\"color:inherit;white-space:normal;line-height:2em\\\"><span style=\\\"color:#fff\\\">输入内容正文</span></p><p style=\\\"color:inherit;white-space:normal;line-height:2em\\\"><span style=\\\"color:#fff\\\">为保证效果，强烈建议更换图片为正方形。</span></p></section></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":18,\"Total\":190,\"id\":855,\"name\":\"<section  ><section style=\\\"font-size:14px;font-family:'Microsoft YaHei';margin: 5px auto;white-space: normal;\\\"><section style=\\\"margin:20px auto;padding:0;width:80%;text-align:center\\\"><section style=\\\"margin:0;padding:0;border:1px solid #a9a9a9;text-align:center;box-shadow:0 0 8px #787878\\\"><section style=\\\"margin-top:-5px;margin-left:-5px;margin-bottom:4px;border:1px solid #a9a9a9;padding:0;background-color:#fff;box-shadow:0 0 8px #c6c6c6\\\"><section style=\\\"margin-top:-5px;margin-left:-5px;margin-bottom:4px;border:1px solid #a9a9a9;padding:10px;background-color:#fefefe;box-shadow:#c6c6c6 0 0 8px\\\"><section style=\\\"clear:both;overflow:hidden;border:0;margin:0;padding:0;display:inline-block;width:100%\\\"><section style=\\\"width:100%;margin:0;padding:0;border-color:#757576;color:inherit\\\"><img   style=\\\"width:100%;display:block;padding:0;margin:0\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/c37004c7-d925-42d0-afdf-40f18bdb64ce21.jpg\\\"></section></section></section></section></section></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":19,\"Total\":190,\"id\":854,\"name\":\"<section  ><section style=\\\"font-size:14px;font-family:'Microsoft YaHei';margin: 5px auto;white-space: normal;\\\"><section style=\\\"margin:20px 0;padding:0\\\"><section style=\\\"margin:0;padding:0;box-sizing:border-box;text-align:center\\\"><section style=\\\"margin:0;padding:0;box-sizing:border-box;display:inline-block\\\"><section style=\\\"margin:0;padding:0;width:1.8em;height:1.8em;border-radius:50%;box-shadow:0 2px 3px #999;border:1px solid #ccc;display:table-cell;vertical-align:middle\\\"><section style=\\\"margin:0 auto;padding:0;width:1.4em;height:1.4em;border-radius:50%;background-color:#ccc\\\"></section></section></section><section style=\\\"margin:-10px auto 0;padding:0;width:4em;height:4em;border-right:1px solid #ccc;border-top:1px solid #ccc;box-sizing:border-box;transform:rotate(-45deg) translateZ(0);-moz-transform:rotate(-45deg) translateZ(0);-ms-transform:rotate(-45deg) translateZ(0);-o-transform:rotate(-45deg) translateZ(0);-webkit-transform:rotate(-45deg) translateZ(0)\\\"></section></section><section style=\\\"margin-top:-2em;padding:15px;box-sizing:border-box;border:1px solid #ccc;border-radius:10px\\\"><img   style=\\\"width:100%;display:block\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/d2ea9da3-a853-41d2-886d-8ce7e65fe94d22.jpg\\\"></section></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null},{\"RowNumber\":20,\"Total\":190,\"id\":853,\"name\":\"<section  ><section style=\\\"margin: 5px 0;box-sizing: border-box;display: table;width: 100%;\\\"><section style=\\\"line-height: 0; box-sizing: border-box; color: inherit;\\\"><img   style=\\\"border: 0px; box-sizing: border-box; display: inline-block; width: 100%; max-width: 100%; height: auto !important; color: inherit;\\\" src=\\\"http://files.qiluzhaoshang.com//fck007/2017042015/1e48f04b-3030-4e12-998e-bf85cf394a7723.png\\\"></section><section style=\\\"width: 30%; display: inline-block; float: left; text-align: right; margin: 15px 0px 0px; padding: 0px; box-sizing: border-box; color: inherit;\\\"><section style=\\\"float: right; text-align: center; margin-top: -15px; box-sizing: border-box; color: inherit;\\\"><section style=\\\"width: 1px; height: 1.2em; margin-left: 13px; background-color: rgb(102, 102, 102); box-sizing: border-box; color: inherit;\\\"></section><section style=\\\"width: 2em; height: 2em; border: 1px solid rgb(102, 102, 102); border-top-left-radius: 50%; border-top-right-radius: 50%; border-bottom-right-radius: 50%; border-bottom-left-radius: 50%; line-height: 2em; font-size: 1em; font-weight: inherit; text-decoration: inherit; box-sizing: border-box; color: inherit;\\\"><section style=\\\"box-sizing: border-box; color: inherit;\\\">巴</section></section><section style=\\\"width: 2em; height: 2em; border: 1px solid rgb(102, 102, 102); margin-top: 2px; border-top-left-radius: 50%; border-top-right-radius: 50%; border-bottom-right-radius: 50%; border-bottom-left-radius: 50%; line-height: 2em; font-size: 1em; font-weight: inherit; text-decoration: inherit; box-sizing: border-box; color: inherit;\\\"><section style=\\\"box-sizing: border-box; color: inherit;\\\">瑶</section></section></section></section><section  style=\\\"width: 65%; float: left; margin-top: 20px; line-height: 1.5em; padding-left: 20px; font-size: 1em; font-weight: inherit; text-decoration: inherit; color: rgb(58, 188, 255); box-sizing: border-box;\\\"><section style=\\\"box-sizing: border-box; border-color: rgb(58, 188, 255); color: inherit;\\\"><section style=\\\"box-sizing: border-box; font-size: 175%; margin: 5px 0px; border-color: rgb(58, 188, 255); color: inherit;\\\">海上人家</section><section style=\\\"box-sizing: border-box; font-size: 16px; border-color: rgb(58, 188, 255); color: inherit;\\\">巴瑶族，唯一的海上民族</section></section></section></section></section>\",\"path\":null,\"ItemSalePrice\":0,\"AppUrl\":null,\"ProductForm\":0,\"ProductTips\":null}]"
		err := gjson.DecodeTo(strJsonData, &pageDataItemVoList)
		if err != nil {
			return nil, err
		}

		if !g.IsEmpty(pageDataItemVoList) {
			baseListRes.Items = pageDataItemVoList
		}
	case 4:
		// 快捷入口
		strJson := "[\n" +
			"  {\n" +
			"    \"id\": 23,\n" +
			"    \"name\": \"扫码点餐\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon23.png\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"AppUrl\": \"/chain/chain/scan\",\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null,\n" +
			"    \"RowNumber\": 1,\n" +
			"    \"Total\": 19\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 2,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 22,\n" +
			"    \"name\": \"好友砍价\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon22.png\",\n" +
			"    \"AppUrl\": \"/activity/cutprice/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 5,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 17,\n" +
			"    \"name\": \"餐饮外卖\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon17.png\",\n" +
			"    \"AppUrl\": \"/pagesub/index/store-list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 3,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 16,\n" +
			"    \"name\": \"新闻资讯\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon16.png\",\n" +
			"    \"AppUrl\": \"/pagesub/article/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 4,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 15,\n" +
			"    \"name\": \"优惠买单\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon15.png\",\n" +
			"    \"AppUrl\": \"/chain/chain/favorable\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 6,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 14,\n" +
			"    \"name\": \"服务预约\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon14.png\",\n" +
			"    \"AppUrl\": \"/pagesub/product/list?kind_id=1202\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 7,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 13,\n" +
			"    \"name\": \"拼团活动\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon13.png\",\n" +
			"    \"AppUrl\": \"/activity/fightgroup/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 8,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 12,\n" +
			"    \"name\": \"粉丝榜\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon12.png\",\n" +
			"    \"AppUrl\": \"/member/fans/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 9,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 11,\n" +
			"    \"name\": \"砸金蛋\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon11.png\",\n" +
			"    \"AppUrl\": \"/activity/smashgoldeneggs/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 10,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 10,\n" +
			"    \"name\": \"幸运抽奖\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon10.png\",\n" +
			"    \"AppUrl\": \"/member/smashgoldeneggs/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 11,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 9,\n" +
			"    \"name\": \"领券中心\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon9.png\",\n" +
			"    \"AppUrl\": \"/activity/coupon/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 12,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 8,\n" +
			"    \"name\": \"附近门店\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon8.png\",\n" +
			"    \"AppUrl\": \"/chain/chain/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 14,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 7,\n" +
			"    \"name\": \"物流查询\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon7.png\",\n" +
			"    \"AppUrl\": \"/member/order/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 13,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 6,\n" +
			"    \"name\": \"活动中心\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon6.png\",\n" +
			"    \"AppUrl\": \"/activity/market/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 15,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 5,\n" +
			"    \"name\": \"我的粉丝\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon5.png\",\n" +
			"    \"AppUrl\": \"/member/fans/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 16,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 4,\n" +
			"    \"name\": \"分享赚钱\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon4.png\",\n" +
			"    \"AppUrl\": \"/member/fans/endorsement\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 17,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 3,\n" +
			"    \"name\": \"我的收藏\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon3.png\",\n" +
			"    \"AppUrl\": \"/member/member/favorites\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 18,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 2,\n" +
			"    \"name\": \"我的拼团\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon2.png\",\n" +
			"    \"AppUrl\": \"/activity/fightgroup/order\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 19,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 1,\n" +
			"    \"name\": \"我的金库\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon1.png\",\n" +
			"    \"AppUrl\": \"/member/cash/predeposit\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 20,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 30,\n" +
			"    \"name\": \"店铺街\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon30.png\",\n" +
			"    \"AppUrl\": \"/pagesub/index/store-list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 20,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 31,\n" +
			"    \"name\": \"抢购活动\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon31.png\",\n" +
			"    \"AppUrl\": \"/pagesub/product/list?tag_id=1404&cname=抢购活动\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 20,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 32,\n" +
			"    \"name\": \"众宝区\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon32.png\",\n" +
			"    \"AppUrl\": \"/pagesub/product/list?sp_from=1&sp_to=1000000&cname=众宝区\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 20,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 33,\n" +
			"    \"name\": \"积分区\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon33.png\",\n" +
			"    \"AppUrl\": \"/pagesub/product/list?points_from=1&points_to=1000000&cname=积分区\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 20,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 34,\n" +
			"    \"name\": \"积分商城\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon34.png\",\n" +
			"    \"AppUrl\": \"/integral/integral/integral\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 20,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 35,\n" +
			"    \"name\": \"跨境商品\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon35.png\",\n" +
			"    \"AppUrl\": \"/pagesub/product/list?tag_id=1405&cname=跨境商品\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 20,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 36,\n" +
			"    \"name\": \"限时折扣\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon36.png\",\n" +
			"    \"AppUrl\": \"/pagesub/product/list?activity_type_id=1103&cname=限时折扣\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 20,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 37,\n" +
			"    \"name\": \"满减\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon37.png\",\n" +
			"    \"AppUrl\": \"/pagesub/product/list?activity_type_id=1107&cname=满减\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 20,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 38,\n" +
			"    \"name\": \"平台秒杀\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon38.png\",\n" +
			"    \"AppUrl\": \"/activity/plantform/secondlist?cname=限时秒杀\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 20,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 39,\n" +
			"    \"name\": \"直播\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon39.png\",\n" +
			"    \"AppUrl\": \"/pagesub/uLive/index\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  },\n" +
			"  {\n" +
			"    \"RowNumber\": 20,\n" +
			"    \"Total\": 19,\n" +
			"    \"id\": 40,\n" +
			"    \"name\": \"组合套餐\",\n" +
			"    \"path\": \"https://static.shopsuite.cn/xcxfile/appicon/icon39.png\",\n" +
			"    \"AppUrl\": \"/activity/giftbag/list\",\n" +
			"    \"ItemSalePrice\": 0,\n" +
			"    \"ProductForm\": 0,\n" +
			"    \"ProductTips\": null\n" +
			"  }\n" +
			"]"
		var dataItemVos []*model.PageDataItemVo
		json.Unmarshal([]byte(strJson), &dataItemVos)

		idList := []uint64{1, 2, 6, 7, 8, 10, 11, 12, 14, 15, 17, 23, 35, 37, 38, 30, 32, 31}

		filteredDataItemVos := []*model.PageDataItemVo{}
		for _, item := range dataItemVos {
			contains := false
			for _, id := range idList {
				if item.Id == id {
					contains = true
					break
				}
			}
			if !contains {
				filteredDataItemVos = append(filteredDataItemVos, item)
			}
		}

		if len(filteredDataItemVos) > 0 {
			baseListRes.Items = filteredDataItemVos
		}
	case 5:
		articleCategoryQueryWrapper := &do.ArticleCategoryListInput{
			BaseList: ml.BaseList{
				Page: pageDataReq.Page,
				Size: pageDataReq.Size,
			},
		}

		if !g.IsEmpty(pageDataReq.Name) {
			var likes = []*ml.WhereExt{{
				Column: dao.ArticleCategory.Columns().CategoryName,
				Val:    "%" + pageDataReq.Name + "%",
				Symbol: ml.LIKE,
			}}
			articleCategoryQueryWrapper.WhereExt = likes
		}

		articleCategoryPage, err := dao.ArticleCategory.List(ctx, articleCategoryQueryWrapper)

		if err != nil {
			return nil, err
		}

		gconv.Struct(articleCategoryPage, baseListRes)

		if !g.IsEmpty(articleCategoryPage.Items) {
			for _, item := range articleCategoryPage.Items {
				pageDataItemVo := &model.PageDataItemVo{
					Id:   uint64(item.CategoryId),
					Path: item.CategoryImageUrl,
					Name: item.CategoryName,
				}
				pageDataItemVoList = append(pageDataItemVoList, pageDataItemVo)
			}
			baseListRes.Items = pageDataItemVoList
		}
	case 6:
		// 资讯
		articleBaseQueryWrapper := &do.ArticleBaseListInput{
			BaseList: ml.BaseList{
				Page: pageDataReq.Page,
				Size: pageDataReq.Size,
			},
		}

		if !g.IsEmpty(pageDataReq.Name) {
			var likes = []*ml.WhereExt{{
				Column: dao.ArticleBase.Columns().ArticleName,
				Val:    "%" + pageDataReq.Name + "%",
				Symbol: ml.LIKE,
			}}
			articleBaseQueryWrapper.WhereExt = likes
		}
		articleBasePage, err := dao.ArticleBase.List(ctx, articleBaseQueryWrapper)
		if err != nil {
			return nil, err
		}

		gconv.Scan(articleBasePage, baseListRes)

		if len(articleBasePage.Items) > 0 {
			for _, item := range articleBasePage.Items {
				pageDataItemVo := &model.PageDataItemVo{
					Id:   uint64(item.ArticleId),
					Path: item.ArticleImage,
					Name: item.ArticleTitle,
				}
				pageDataItemVoList = append(pageDataItemVoList, pageDataItemVo)
			}
			baseListRes.Items = pageDataItemVoList
		}
	case 8:
		// 自定义页面
		pageBaseQueryWrapper := &do.PageBaseListInput{
			BaseList: ml.BaseList{
				Page: pageDataReq.Page,
				Size: pageDataReq.Size,
			},
		}

		if !g.IsEmpty(pageDataReq.Name) {
			var likes = []*ml.WhereExt{{
				Column: dao.PageBase.Columns().PageName,
				Val:    "%" + pageDataReq.Name + "%",
				Symbol: ml.LIKE,
			}}
			pageBaseQueryWrapper.WhereExt = likes
		}
		pageBasePage, err := dao.PageBase.List(ctx, pageBaseQueryWrapper)
		if err != nil {
			return nil, err
		}

		gconv.Scan(pageBasePage, baseListRes)

		for _, item := range pageBasePage.Items {
			pageDataItemVo := &model.PageDataItemVo{}
			pageDataItemVo.Id = item.PageId
			pageDataItemVo.Path = item.PageShareImage
			pageDataItemVo.Name = item.PageName
			pageDataItemVoList = append(pageDataItemVoList, pageDataItemVo)
		}
		baseListRes.Items = pageDataItemVoList

	default:
		break
	}
	return baseListRes, nil
}
