package sys

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
	"reflect"
	"regexp"
)

var (
	Config = cConfig{}
)

type cConfig struct{}

// =========================== 管理端使用 =============================

// ---------------------------- 配置分类 -------------------------------
// List 配置分类列表
func (c *cConfig) ListType(ctx context.Context, req *sys.ConfigTypeListReq) (res *sys.ConfigTypeListRes, err error) {
	var result, error = service.ConfigType().List(ctx, &do.ConfigTypeListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size: req.Size,
			Sidx: dao.ConfigType.Columns().ConfigTypeSort,
			Sort: "ASC"},
		Where: do.ConfigType{},
	})

	if error != nil {
		err = error
	}

	res = &sys.ConfigTypeListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 添加配置分类
func (c *cConfig) AddType(ctx context.Context, req *sys.ConfigTypeAddReq) (res *sys.ConfigTypeEditRes, err error) {
	input := do.ConfigType{}
	gconv.Scan(req, &input)

	var result, error = service.ConfigType().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.ConfigTypeEditRes{
		ConfigTypeId: uint(result),
	}

	return
}

// Edit 编辑配置分类
func (c *cConfig) EditType(ctx context.Context, req *sys.ConfigTypeEditReq) (res *sys.ConfigTypeEditRes, err error) {
	input := do.ConfigType{}
	gconv.Scan(req, &input)

	_, error := service.ConfigType().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.ConfigTypeEditRes{
		ConfigTypeId: req.ConfigTypeId,
	}

	return
}

// Remove 删除配置分类
func (c *cConfig) RemoveType(ctx context.Context, req *sys.ConfigTypeRemoveReq) (res *sys.ConfigTypeRemoveRes, err error) {
	var _, error = service.ConfigType().Remove(ctx, req.ConfigTypeId)

	if error != nil {
		err = error
	}

	res = &sys.ConfigTypeRemoveRes{}

	return
}

// List 配置项目列表
func (c *cConfig) ListBase(ctx context.Context, req *sys.ConfigBaseListReq) (res *sys.ConfigBaseListRes, err error) {
	dictId := req.ConfigTypeId
	item := do.ConfigBase{ConfigTypeId: req.ConfigTypeId}
	if dictId == 0 {
		item.ConfigTypeId = nil
	}

	var result, error = service.ConfigBase().List(ctx, &do.ConfigBaseListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size: req.Size,
			Sidx: dao.ConfigBase.Columns().ConfigSort,
			Sort: "ASC"},
		Where: item,
	})

	if error != nil {
		err = error
	}

	res = &sys.ConfigBaseListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 添加配置项目
func (c *cConfig) AddBase(ctx context.Context, req *sys.ConfigBaseAddReq) (res *sys.ConfigBaseEditRes, err error) {
	input := do.ConfigBase{}
	gconv.Scan(req, &input)

	var _, error = service.ConfigBase().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.ConfigBaseEditRes{
		ConfigKey: req.ConfigKey,
	}

	return
}

// Edit 编辑配置项目
func (c *cConfig) EditBase(ctx context.Context, req *sys.ConfigBaseEditReq) (res *sys.ConfigBaseEditRes, err error) {
	input := do.ConfigBase{}
	gconv.Scan(req, &input)

	var _, error = service.ConfigBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.ConfigBaseEditRes{
		ConfigKey: req.ConfigKey,
	}

	return
}

// Remove 删除配置项目
func (c *cConfig) RemoveBase(ctx context.Context, req *sys.ConfigBaseRemoveReq) (res *sys.ConfigBaseRemoveRes, err error) {
	var _, error = service.ConfigBase().Remove(ctx, req.ConfigKey)

	if error != nil {
		err = error
	}

	res = &sys.ConfigBaseRemoveRes{}

	return
}

// Tree 站点设置
func (c *cConfig) Tree(ctx context.Context, req *sys.ConfigTreeReq) (res sys.ConfigTreeRes, err error) {

	return
}

// Index 配置项目
func (c *cConfig) Index(ctx context.Context, req *sys.ConfigBaseIndexReq) (res sys.ConfigBaseIndexRes, err error) {
	input := do.ConfigTypeListInput{}
	gconv.Scan(req, &input)
	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	//默认排序
	input.Sidx = dao.ConfigType.Columns().ConfigTypeSort
	input.Sort = ml.ORDER_BY_ASC

	// 获取配置列表
	var configType, errorType = service.ConfigType().Find(ctx, &input)
	if errorType != nil {
		err = errorType
	}

	var configData, error = service.ConfigBase().Find(ctx, &do.ConfigBaseListInput{
		BaseList: ml.BaseList{
			Sidx: dao.ConfigBase.Columns().ConfigSort,
			Sort: "ASC"},
	})

	if error != nil {
		err = error
	}

	for _, v := range configType {
		it := sys.ConfigBaseIndexVo{}
		it.ConfigTypeId = v.ConfigTypeId
		it.ConfigTypeName = v.ConfigTypeName

		// 查询配置项列表
		itemList := make([]map[string]interface{}, 0)
		for _, v := range configData {
			if it.ConfigTypeId == v.ConfigTypeId {

				item := make(map[string]interface{})
				item["config_key"] = v.ConfigKey
				item["config_title"] = v.ConfigTitle
				//it["config_code"] = v.ConfigValue
				item["config_value"] = v.ConfigValue
				item["config_datatype"] = v.ConfigDatatype
				item["config_note"] = v.ConfigNote

				if v.ConfigDatatype == consts.CHECKBOX {
					// 复选框
					re := regexp.MustCompile(`\r?\n`)
					list := gstr.Split(re.ReplaceAllString(v.ConfigOptions, "|"), "|")
					optionsList := make(map[string]string)
					for _, val := range list {
						if !g.IsEmpty(val) {
							re2 := regexp.MustCompile(`:|：|\s+`)
							item := gstr.Split(re2.ReplaceAllString(val, "|"), "|")
							optionsList[item[0]] = item[1]
						}
					}
					// 选择项
					item["optionsList"] = optionsList
					// 选择值
					item["config_value"] = gstr.Split(v.ConfigValue, ",")

				} else if v.ConfigDatatype == consts.RADIO {
					// 单选框
					re := regexp.MustCompile(`\r?\n`)
					list := gstr.Split(re.ReplaceAllString(v.ConfigOptions, "|"), "|")
					optionsList := make(map[string]string)
					for _, v := range list {
						if !g.IsEmpty(v) {
							re2 := regexp.MustCompile(`:|：|\s+`)
							item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
							optionsList[item[0]] = item[1]
						}
					}

					item["optionsList"] = optionsList

				} else if v.ConfigDatatype == consts.SELECT {
					// 下拉选择框
					re := regexp.MustCompile(`\r?\n`)
					list := gstr.Split(re.ReplaceAllString(v.ConfigOptions, "|"), "|")
					optionsList := make(map[string]string)
					for _, v := range list {
						if !g.IsEmpty(v) {
							re2 := regexp.MustCompile(`:|：|\s+`)
							item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
							optionsList[item[0]] = item[1]
						}
					}
					item["optionsList"] = optionsList
				} else if v.ConfigDatatype == consts.IMAGE {
					// 单图片
					item["config_value"] = v.ConfigValue
				} else if v.ConfigDatatype == consts.IMAGES {
					// 多图片
					list := gstr.Split(v.ConfigValue, ",")
					itemList := make([]string, 0)
					for _, v := range list {
						// 图片地址
						item := v
						itemList = append(itemList, item)
					}
					item["config_value"] = itemList
				}

				itemList = append(itemList, item)
			}
		}

		it.Items = itemList

		// 加入数组
		res.Items = append(res.Items, it)
	}

	return
}

// EditSite 编辑站点设置
func (c *cConfig) EditSite(ctx context.Context, req *sys.ConfigBaseEditSiteReq) (res sys.ConfigBaseEditSiteRes, err error) {

	// key：string类型，value：interface{}  类型能存任何数据类型
	var jsonObj map[string]interface{}
	//request := g.RequestFromCtx(ctx)
	//json.Unmarshal(request.GetBody(), &jsonObj)
	jsonObj = req.Configs

	for key, val := range jsonObj {
		// 数组处理
		if reflect.ValueOf(val).Kind() == reflect.Slice {
			if reflect.TypeOf(val).String() == "[]interface {}" {
				// 初始化URL数组
				item := make([]string, 0)
				for _, v := range val.([]interface{}) {
					value := gconv.String(v)
					// 判断是否http(s)开头
					if gstr.SubStr(value, 0, 4) == "http" ||
						gstr.SubStr(value, 0, 5) == "https" {
						// 图片地址
						item = append(item, value)
					} else {
						// 复选框处理
						item = append(item, value)
					}
				}
				// 逗号拼接
				val = gstr.Join(item, ",")
			}
		} else {
			// 图片处理
			val = gstr.Trim(gconv.String(val), ",")
		}

		// 查询记录
		info, err := service.ConfigBase().Get(ctx, key)
		if err != nil || info == nil {
			continue
		}

		// 更新记录
		rows, err := service.ConfigBase().Edit(ctx, &do.ConfigBase{
			ConfigKey:   key,
			ConfigValue: val,
		})

		if err != nil {
			continue
		}

		// 获取受影响行数
		if rows == 0 {
			continue
		}
	}

	return
}

func (c *cConfig) CleanCache(ctx context.Context, req *sys.CleanCacheReq) (res sys.CleanCacheRes, err error) {
	return
}
