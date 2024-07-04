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
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/pt"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"strings"
)

type sPageModule struct{}

func init() {
	service.RegisterPageModule(NewPageModule())
}

func NewPageModule() *sPageModule {
	return &sPageModule{}
}

// 读取商品分类
func (s *sPageModule) Get(ctx context.Context, id any) (out *entity.PageModule, err error) {
	var list []*entity.PageModule
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// 读取多条记录模式
func (s *sPageModule) Gets(ctx context.Context, id any) (list []*entity.PageModule, err error) {

	err = dao.PageModule.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sPageModule) Find(ctx context.Context, in *do.PageModuleListInput) (out []*entity.PageModule, err error) {
	out, err = dao.PageModule.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sPageModule) List(ctx context.Context, in *do.PageModuleListInput) (out *do.PageModuleListOutput, err error) {
	out, err = dao.PageModule.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sPageModule) Add(ctx context.Context, in *do.PageModule) (lastInsertId int64, err error) {
	lastInsertId, err = dao.PageModule.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sPageModule) Edit(ctx context.Context, in *do.PageModule) (affected int64, err error) {
	_, err = dao.PageModule.Edit(ctx, in.PmId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sPageModule) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.PageModule.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// FixPcPageModuleData 修复PC页面模块数据
func (s *sPageModule) FixPcPageModuleData(ctx context.Context, pageData []*entity.PageModule) ([]map[string]interface{}, error) {
	data := make([]map[string]interface{}, 0)
	if len(pageData) > 0 {
		for _, moduleRow := range pageData {
			// 将 moduleRow 转换为 map 并进行下划线命名
			moduleDefault := gconv.MapDeep(moduleRow)
			module := gconv.Map(moduleDefault)

			// 加入数据
			data = append(data, module)

			// 系统启用自动翻译功能
			moduleId := moduleRow.ModuleId
			var pmJson gjson.Json
			if err := gjson.DecodeTo(moduleRow.PmJson, &pmJson); err != nil {
				// json 解析错误可以忽略（脏数据）
				continue
			}

			module["pm_json"] = pmJson

			if gstr.InArray([]string{"1001", "1004", "1005", "1006"}, moduleId) {
				// 读取商品
				tabs := pmJson.GetJsons("tabs")
				if len(tabs) > 0 {
					itemIds := make([]uint64, 0)
					for _, tab := range tabs {
						items := tab.GetJsons("items")
						if len(items) > 0 {
							for _, item := range items {
								itemIds = append(itemIds, gconv.Uint64(item.Get("item_id")))
							}
						}
					}
					itemIds = gconv.Uint64s(garray.NewArrayFrom(gconv.Interfaces(itemIds)).Unique())

					// 将 itemIds 转换为以逗号分隔的字符串
					itemIdsStr := gconv.Strings(itemIds)
					itemIdsJoined := strings.Join(itemIdsStr, ",")

					// 注意：确保 pt.ItemListReq 结构体中的 ItemId 字段接受类型 string
					input := &pt.ItemListReq{
						ItemId: itemIdsJoined,
					}
					itemListRes, err := service.ProductIndex().ListItem(ctx, input)
					if err != nil {
						return nil, err
					}

					itemRows := itemListRes.Items

					for _, tab := range tabs {
						items := tab.GetJsons("items")
						if len(items) > 0 {
							for _, item := range items {
								itemId := gconv.Uint64(item.Get("item_id"))
								if len(itemRows) > 0 {
									for _, itemRow := range itemRows {
										if itemId == itemRow.ItemId {
											item.Set("item_unit_price", itemRow.ItemUnitPrice)
											item.Set("item_market_price", itemRow.ItemMarketPrice)
											break
										}
									}
								}

								item.Set("activity_type_id", 0)
								item.Set("activity_type_name", 0)
							}
						}
					}
				}
			} else if moduleId == "1104" {
				// 读取推荐品牌
				brandQueryWrapper := &do.ProductBrandListInput{
					Where: do.ProductBrand{
						BrandRecommend: 1,
						BrandEnable:    1,
					},
				}
				brandPage, err := service.ProductBrand().List(ctx, brandQueryWrapper)
				if err != nil {
					return nil, err
				}

				productBrands := brandPage.Records
				pmJson.Set("brand_rows", productBrands)
			}
		}
	}

	return data, nil
}

// GetModuleTpl 获取模块模板
func (s *sPageModule) GetModuleTpl(ctx context.Context) (map[string]interface{}, error) {
	// 获取配置 service_user_id
	configBaseUserId, err := service.ConfigBase().Get(ctx, "service_user_id")
	if err != nil {
		return nil, gerror.New("获取配置 service_user_id 失败: " + err.Error())
	}

	// 获取配置 service_app_key
	configBaseAppKey, err := service.ConfigBase().Get(ctx, "service_app_key")
	if err != nil {
		return nil, gerror.New("获取配置 service_app_key 失败: " + err.Error())
	}

	// 调用 cloundService 获取模块模板
	moduleTpl, err := service.Cloud().GetModuleTpl(gconv.Int(configBaseUserId.ConfigValue), configBaseAppKey.ConfigValue)
	if err != nil {
		return nil, gerror.New("获取模块模板失败: " + err.Error())
	}

	return moduleTpl, nil
}

// GetLists 获取页面模块列表
func (s *sPageModule) GetLists(ctx context.Context, req *do.PageModuleListInput) (pageModuleVo *model.PageModuleVoOutput, err error) {
	// 初始化返回的分页数据结构
	pageModuleVo = &model.PageModuleVoOutput{}

	// 查询条件
	pageModuleQueryWrapper := &do.PageModuleListInput{
		Where: do.PageModule{PageId: req.Where.PageId},
	}

	// 获取页面模块列表，分页查询
	modulePage, err := s.List(ctx, pageModuleQueryWrapper)
	if err != nil {
		return nil, err
	}

	// 将 modulePage 的分页数据复制到 pageModuleVo
	gconv.Scan(modulePage.Items, &pageModuleVo)

	// 初始化页面模块 VO 列表
	pageModuleVos := make([]*model.PageModuleVo, 0)

	// 判断 modulePage 是否为空
	if !g.IsEmpty(modulePage.Items) {
		for _, pageModule := range modulePage.Items {
			// 创建单个页面模块 VO
			moduleVo := &model.PageModuleVo{}
			gconv.Scan(pageModule, moduleVo)

			// 解析 JSON 数据并设置到 VO 中
			parse := gjson.New(pageModule.PmJson)
			moduleVo.PmJson = parse.Map()

			// 添加到页面模块 VO 列表
			pageModuleVos = append(pageModuleVos, moduleVo)
		}
		// 设置分页数据中的记录
		pageModuleVo.Items = pageModuleVos
	}

	return pageModuleVo, nil
}
