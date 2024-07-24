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

package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/dao/global"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility"
	"golershop.cn/utility/array"
	"golershop.cn/utility/log"
	"regexp"
	"sort"
)

var (
	Registry map[string]*entity.ConfigBase = make(map[string]*entity.ConfigBase)
)

type sConfigBase struct{}

func init() {
	service.RegisterConfigBase(New())
}

func New() *sConfigBase {
	return &sConfigBase{}
}

func (s *sConfigBase) GetBool(ctx context.Context, id any, defaultValue bool) (res bool) {
	config, err := s.Get(ctx, id)

	if err != nil || config == nil {
		return defaultValue
	}

	return gconv.Bool(config.ConfigValue)
}

func (s *sConfigBase) GetStr(ctx context.Context, id any, defaultValue string) (res string) {
	config, err := s.Get(ctx, id)

	if err != nil || config == nil {
		return defaultValue
	}

	return config.ConfigValue
}

func (s *sConfigBase) GetInt(ctx context.Context, id any, defaultValue int) (res int) {
	config, err := s.Get(ctx, id)

	if err != nil || config == nil {
		return defaultValue
	}

	return gconv.Int(config.ConfigValue)
}

func (s *sConfigBase) GetInt64(ctx context.Context, id any, defaultValue int64) (res int64) {
	config, err := s.Get(ctx, id)

	if err != nil || config == nil {
		return defaultValue
	}

	return gconv.Int64(config.ConfigValue)
}

func (s *sConfigBase) GetUint(ctx context.Context, id any, defaultValue uint) (res uint) {
	config, err := s.Get(ctx, id)

	if err != nil || config == nil {
		return defaultValue
	}

	return gconv.Uint(config.ConfigValue)
}

func (s *sConfigBase) GetFloat(ctx context.Context, id any, defaultValue float64) (res float64) {
	config, err := s.Get(ctx, id)

	if err != nil || config == nil {
		return defaultValue
	}

	return gconv.Float64(config.ConfigValue)
}

// Get 读取属性
func (s *sConfigBase) Get(ctx context.Context, id any) (config *entity.ConfigBase, err error) {
	if value, ok := Registry[id.(string)]; ok {
		config = value
	} else {
		var list []*entity.ConfigBase
		list, err = s.Gets(ctx, id)

		if err != nil {
			return nil, err
		}

		if len(list) > 0 {
			config = list[0]

			Registry[config.ConfigKey] = config
		}
	}

	return config, err
}

// Gets 读取多条属性
func (s *sConfigBase) Gets(ctx context.Context, id any) (list []*entity.ConfigBase, err error) {
	err = dao.ConfigBase.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sConfigBase) Find(ctx context.Context, in *do.ConfigBaseListInput) (out []*entity.ConfigBase, err error) {
	out, err = dao.ConfigBase.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sConfigBase) List(ctx context.Context, in *do.ConfigBaseListInput) (out *do.ConfigBaseListOutput, err error) {
	out, err = dao.ConfigBase.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sConfigBase) Add(ctx context.Context, in *do.ConfigBase) (lastInsertId interface{}, err error) {
	_, err = dao.ConfigBase.Add(ctx, in)
	if err != nil {
		return "", err
	}
	return in.ConfigKey, err
}

// Edit 编辑
func (s *sConfigBase) Edit(ctx context.Context, in *do.ConfigBase) (affected int64, err error) {
	_, err = dao.ConfigBase.Edit(ctx, in.ConfigKey, in)

	if err != nil {
		return 0, err
	}

	if in.ConfigKey.(string) == "sc_order_process" {
		global.StateIdRow = []uint{}
		global.StateIdAll = []uint{}
		global.StateIdSelect = []*model.SelectVo{}
	}

	if array.InArray([]string{"wechat_pay_enable", "alipay_enable", "offline_pay_enable", "money_pay_enable", "points_pay_enable"}, in.ConfigKey.(string)) {
		global.PaymentChannelSelect = []*model.SelectVo{}
		global.PaymentChannelMap = make(map[uint]string)
	}

	delete(Registry, in.ConfigKey.(string))

	return
}

// Save 保存
func (s *sConfigBase) Save(ctx context.Context, in *do.ConfigBase) (affected int64, err error) {
	affected, err = dao.ConfigBase.Save(ctx, in)

	if err != nil {
		return 0, err
	}

	if in.ConfigKey.(string) == "sc_order_process" {
		global.StateIdRow = []uint{}
		global.StateIdAll = []uint{}
		global.StateIdSelect = []*model.SelectVo{}
	}

	if array.InArray([]string{"wechat_pay_enable", "alipay_enable", "offline_pay_enable", "money_pay_enable", "points_pay_enable"}, in.ConfigKey.(string)) {
		global.PaymentChannelSelect = []*model.SelectVo{}
		global.PaymentChannelMap = make(map[uint]string)
	}

	delete(Registry, in.ConfigKey.(string))

	return
}

// Remove 删除多条记录模式
func (s *sConfigBase) Remove(ctx context.Context, id any) (affected int64, err error) {
	one, err := dao.ConfigBase.Get(ctx, id)

	if one.ConfigBuildin {
		return affected, errors.New("系统内置，不可删除！")
	}

	affected, err = dao.ConfigBase.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	delete(Registry, id.(string))

	return affected, err
}

// GetDefaultImage 系统默认图片
func (s *sConfigBase) GetDefaultImage(ctx context.Context) (defaultImage string) {
	config, _ := s.Get(ctx, "default_image")

	defaultImage = config.ConfigValue

	return
}

// getOrderStateList 读取订单状态选项
func (s *sConfigBase) GetOrderStateList(ctx context.Context) (optionsList []*model.SelectVo, err error) {
	if len(global.StateIdSelect) > 0 {
		optionsList = global.StateIdSelect
	} else {
		optionsList, err = s.initOrderStateList(ctx)

		if err != nil {
			return nil, err
		}
	}

	return
}

// getNextOrderStateId 读取配置，根据当前orderStateId获得下一状态 sc_order_process
func (s *sConfigBase) GetNextOrderStateId(ctx context.Context, orderStateId uint) (nextOrderStateId uint, err error) {
	if len(global.StateIdRow) > 0 {
	} else {
		process, err := s.initOrderProcess(ctx)
		if err != nil {
			return 0, err
		}

		global.StateIdRow = process
	}

	index := array.ArraySearch(global.StateIdRow, orderStateId)

	if index == -1 {
		return 0, errors.New("订单当前状态配置数据有误！")
	} else {
		//最后一个
		if len(global.StateIdRow) == index+1 {
			nextOrderStateId = consts.ORDER_STATE_FINISH

			/*
				//最后一个状态为已发货，强制下一个状态为 已完成
				if orderStateId == consts.ORDER_STATE_SHIPPED {
					nextOrderStateId = consts.ORDER_STATE_FINISH
				} else {
					return 0, errors.New("订单下一个状态配置数据有误！")
				}
			*/
		} else {
			nextOrderStateId = global.StateIdRow[index+1]
		}
	}

	return
}

// getAllNextOrderStateId 读取配置，根据当前orderStateId获得下一状态 sc_order_process
func (s *sConfigBase) GetAllNextOrderStateId(ctx context.Context, orderStateId uint) (nextOrderStateId uint, err error) {
	if len(global.StateIdAll) > 0 {
	} else {
		s.GetOrderStateList(ctx)
	}

	index := array.ArraySearch(global.StateIdAll, orderStateId)

	if index == -1 {
		return 0, errors.New("订单当前状态配置数据有误！")
	} else {
		//最后一个
		if len(global.StateIdAll) == index+1 {
			nextOrderStateId = consts.ORDER_STATE_FINISH
		} else {
			nextOrderStateId = global.StateIdAll[index+1]
		}
	}

	return
}

// GetNextReturnStateId 读取配置，根据当前returnStateId获得下一状态 sc_return_process
func (s *sConfigBase) GetNextReturnStateId(ctx context.Context, returnStateId uint) (uint, error) {
	var nextReturnStateId uint

	// 查找returnStateId在returnStateIdList中的索引
	index := -1
	for i, id := range global.ReturnStateIdList {
		if id == returnStateId {
			index = i
			break
		}
	}

	if index == -1 {
		nextReturnStateId = 0
		// return 0, gerror.New("订单当前状态配置数据有误！")
	} else {
		// 最后一个
		if len(global.ReturnStateIdList) == index+1 {
			nextReturnStateId = consts.RETURN_PROCESS_FINISH
		} else {
			nextReturnStateId = global.ReturnStateIdList[index+1]
		}
	}

	return nextReturnStateId, nil
}

// GetPaymentChannelList 取得支付渠道
func (s *sConfigBase) GetPaymentChannelList(ctx context.Context) (optionsList []*model.SelectVo, err error) {
	if len(global.PaymentChannelSelect) > 0 {
		optionsList = global.PaymentChannelSelect
	} else {
		optionsList, err = s.initPaymentChannelList(ctx)

		if err != nil {
			return nil, err
		}
	}

	return
}

// GetPaymentChannelCode 读取配置，根据paymentChannelId读取paymentChannelCode
func (s *sConfigBase) GetPaymentChannelCode(ctx context.Context, paymentChannelId uint) (string, error) {
	// 获取支付渠道列表
	channelList := global.PaymentChannelSelect

	// 创建一个空的map来存储channelList中的数据
	channelMap := make(map[uint]*model.SelectVo)

	// 将channelList转换成map
	for _, selectVo := range channelList {
		channelMap[selectVo.Value] = selectVo
	}

	// 使用paymentChannelId查询SelectVo对象并获取Ext2字段
	selectedChannel, exists := channelMap[paymentChannelId]
	if !exists {
		return "", gerror.Newf("Payment channel with ID %d not found", paymentChannelId)
	}

	return selectedChannel.Ext2, nil
}

// GetReturnStateList 读取订单状态选项
func (s *sConfigBase) GetReturnStateList(ctx context.Context) (optionsList []*model.SelectVo, err error) {
	if len(global.ReturnStateSelectList) > 0 {
		optionsList = global.ReturnStateSelectList
	} else {
		optionsList, err = s.initReturnStateList(ctx)

		if err != nil {
			return nil, err
		}
	}

	return
}

// GetErrStateList 读取错误状态选项
func (s *sConfigBase) GetErrStateList(ctx context.Context) (optionsList []*model.SelectVo, err error) {
	returnStateSelectList := []*model.SelectVo{}

	labels := []string{
		"未分类异常",
		"微信JSPI异常",
		"微信小程序异常",
		"微信公众号异常",
		"支付宝支付异常",
		"消息推送异常",
		"阿里云服务异常",
		"腾讯云服务异常",
		"华为云服务异常",
	}

	values := []uint{
		consts.ERR_NOT_DEFINITION,
		consts.ERR_WX_JSPI,
		consts.ERR_WX_XCX,
		consts.ERR_WX_MP,
		consts.ERR_ALI_PAY,
		consts.ERR_PSUH_MSG,
		consts.ERR_ALI_SERVICE,
		consts.ERR_TENCENT_SERVICE,
		consts.ERR_HUAWEI_SERVICE,
	}

	for i := 0; i < len(labels); i++ {
		selectVo := &model.SelectVo{
			Label: labels[i],
			Value: values[i],
		}

		returnStateSelectList = append(returnStateSelectList, selectVo)
	}

	return returnStateSelectList, nil
}

// initOrderProcess 处理订单状态
func (s *sConfigBase) initOrderProcess(ctx context.Context) (res []uint, err error) {
	global.StateIdRow = []uint{}

	v, err := dao.ConfigBase.Get(ctx, "sc_order_process")

	if err != nil {
		return nil, err
	}

	// 选择值
	stateIdList := gconv.SliceUint(gstr.Split(v.ConfigValue, ","))

	//从小到大排序
	sort.Slice(stateIdList, func(i, j int) bool {
		return stateIdList[i] < stateIdList[j]
	})

	global.StateIdRow = stateIdList

	return stateIdList, err
}

// GetStateIdList 订单启用状态List
func (s *sConfigBase) GetStateIdList(ctx context.Context) (res []uint, err error) {
	return global.StateIdRow, err
}

// initOrderStateList 读取订单状态选项
func (s *sConfigBase) initOrderStateList(ctx context.Context) (optionsList []*model.SelectVo, err error) {
	global.StateIdAll = []uint{}
	global.StateIdSelect = []*model.SelectVo{}

	v, err := dao.ConfigBase.Get(ctx, "sc_order_process")

	if err != nil {
		return nil, err
	}

	// 选择值
	stateIdRow := gstr.Split(v.ConfigValue, ",")

	if v.ConfigDatatype == consts.CHECKBOX {
		// 复选框
		re := regexp.MustCompile(`\r?\n`)
		list := gstr.Split(re.ReplaceAllString(v.ConfigOptions, "|"), "|")

		for _, val := range list {
			if !g.IsEmpty(val) {
				re2 := regexp.MustCompile(`:|：|\s+`)
				item := gstr.Split(re2.ReplaceAllString(val, "|"), "|")

				//判断是否开启
				if gstr.InArray(stateIdRow, item[0]) {
					optionsList = append(optionsList, &model.SelectVo{Value: gconv.Uint(item[0]), Label: item[1]})
				}

				global.StateIdAll = append(global.StateIdAll, gconv.Uint(item[0]))
			}
		}
	}

	optionsList = append(optionsList, &model.SelectVo{Value: consts.ORDER_STATE_CANCEL, Label: "交易取消"})

	global.StateIdSelect = optionsList

	return
}

// InitReturnProcess 初始化退货处理流程
func (s *sConfigBase) initReturnProcess(ctx context.Context) ([]uint, error) {
	global.ReturnStateIdList = []uint{}

	// 获取配置信息
	configBase, err := s.Get(ctx, "sc_return_process")
	if err != nil {
		return nil, err
	}

	// 选择值

	returnStateIdList := gconv.SliceUint(gstr.Split(configBase.ConfigValue, ","))

	// 从小到大排序
	sort.Slice(returnStateIdList, func(i, j int) bool {
		return returnStateIdList[i] < returnStateIdList[j]
	})

	global.ReturnStateIdList = returnStateIdList
	return returnStateIdList, nil
}

// initReturnStateList 读取订单状态选项
func (s *sConfigBase) initReturnStateList(ctx context.Context) (optionsList []*model.SelectVo, err error) {
	global.ReturnStateSelectList = []*model.SelectVo{}

	labels := []string{
		"提交退单",
		"退单审核",
		"收货确认",
		"退款确认",
		"收款确认",
		"完成",
		"拒绝退货",
		"买家取消",
	}

	values := []uint{
		consts.RETURN_PROCESS_SUBMIT,
		consts.RETURN_PROCESS_CHECK,
		consts.RETURN_PROCESS_RECEIVED,
		consts.RETURN_PROCESS_REFUND,
		consts.RETURN_PROCESS_RECEIPT_CONFIRMATION,
		consts.RETURN_PROCESS_FINISH,
		consts.RETURN_PROCESS_REFUSED,
		consts.RETURN_PROCESS_CANCEL,
	}

	for i := 0; i < len(labels); i++ {
		selectVo := &model.SelectVo{
			Label: labels[i],
			Value: values[i],
		}
		global.ReturnStateSelectList = append(global.ReturnStateSelectList, selectVo)
	}

	return global.ReturnStateSelectList, nil
}

// initPaymentChannelList 取得支付渠道
func (s *sConfigBase) initPaymentChannelList(ctx context.Context) (optionsList []*model.SelectVo, err error) {
	global.PaymentChannelSelect = []*model.SelectVo{}
	global.PaymentChannelMap = make(map[uint]string)

	configTypes, err := dao.ConfigType.Find(ctx, &do.ConfigTypeListInput{Where: do.ConfigType{ConfigTypeModule: 1004}})

	if err != nil {
		return nil, err
	}

	for _, configType := range configTypes {

		//1403	微信支付 wechat_pay_enable
		//1401	支付宝支付 alipay_enable
		//1422	线下支付 offline_pay_enable
		//1406	余额支付 money_pay_enable
		//1413	积分支付 points_pay_enable
		var id string
		var ck string
		var img string
		if configType.ConfigTypeId == 1403 {
			id = "wxpay"
			ck = "wechat_pay_enable"
			img = "wechat_pay_logo"
		} else if configType.ConfigTypeId == 1401 {
			id = "alipay"
			ck = "alipay_enable"
			img = "alipay_logo"
		} else if configType.ConfigTypeId == 1422 {
			id = "offline"
			ck = "offline_pay_enable"
			img = "offline_pay_logo"
		} else if configType.ConfigTypeId == 1406 {
			id = "money"
			ck = "money_pay_enable"
			img = "money_pay_logo"
		} else if configType.ConfigTypeId == 1413 {
			id = "points"
			ck = "points_pay_enable"
			img = "points_pay_logo"
		} else {
			return nil, errors.New("支付类型有误！")
		}

		one, err := dao.ConfigBase.Get(ctx, ck)
		if err != nil {
			return nil, err
		}

		logo, err := dao.ConfigBase.Get(ctx, img)
		if err != nil {
			return nil, err
		}

		//启用的
		if true || one.ConfigValue == "1" {
			optionsList = append(optionsList, &model.SelectVo{Value: configType.ConfigTypeId, Label: configType.ConfigTypeName, Ext1: logo.ConfigValue, Ext2: id, Enable: gconv.Bool(one.ConfigValue)})
			global.PaymentChannelMap[configType.ConfigTypeId] = configType.ConfigTypeName
		}
	}

	global.PaymentChannelSelect = optionsList

	return
}

// Init 初始化
func (s *sConfigBase) Init(ctx context.Context) (res bool, err error) {
	fmt.Println("初始化Config....")
	global.Debug = utility.AppDebug(ctx)
	global.Cache = utility.CacheEnable(ctx)
	global.Namespace = utility.Namespace(ctx)
	global.BaseUrl = utility.UrlBase(ctx)
	global.UrlH5 = utility.UrlH5(ctx)
	global.UrlPc = utility.UrlPc(ctx)

	fmt.Println("初始化清理缓存....")
	s.CleanCache(ctx)

	//初始化  initOrderProcess
	_, err = s.initOrderProcess(ctx)
	if err != nil {
		return false, err
	}

	_, err = s.initOrderStateList(ctx)
	if err != nil {
		return false, err
	}

	_, err = s.initReturnProcess(ctx)
	if err != nil {
		return false, err
	}

	_, err = s.initReturnStateList(ctx)
	if err != nil {
		return false, err
	}

	_, err = s.initPaymentChannelList(ctx)

	if err != nil {
		return false, err
	}

	p := map[string]interface{}{
		"host":        global.BaseUrl,
		"licence_key": global.Lk,
	}

	cs := ml.CloudService{}
	cs.InitService(p)

	fmt.Println("初始化Config完成")

	return
}

// IfInvoicing 是否启用进销存管理
func (s *sConfigBase) IfInvoicing(ctx context.Context) (res bool) {
	return s.GetBool(ctx, "invoicing_enable", false)
}

// IfPlantformFx 是否启用分销
func (s *sConfigBase) IfPlantformFx(ctx context.Context) (res bool) {
	return s.GetBool(ctx, "plantform_fx_enable", false)
}

// IfSupplierMarket 是否启用供应商市场
func (s *sConfigBase) IfSupplierMarket(ctx context.Context) (res bool) {
	return s.GetBool(ctx, "supplier_market_enable", false)
}

// GetSiteInfo 初始化
func (s *sConfigBase) GetSiteInfo(ctx context.Context, sourceUccCode string) (res map[string]interface{}, err error) {

	keyStr := "site_name,site_meta_keyword,site_meta_description,site_version,copyright,icp_number,site_company_name,site_address,site_tel,account_login_bg,site_admin_logo,site_mobile_logo,site_pc_logo,date_format,time_format,cache_enable,cache_expire,site_status,advertisement_open,wechat_connect_auto,wechat_app_id,product_spec_edit,default_image,product_salenum_flag,b2b_flag,hall_b2b_enable,product_ziti_flag,plantform_fx_enable,plantform_fx_gift_point,plantform_fx_withdraw_min_amount,plantform_poster_bg,plantform_commission_withdraw_mode,product_poster_bg,live_mode_xcx,kefu_type_id,withdraw_received_day,withdraw_monthday,default_shipping_district,points_enable,voucher_enable,b2b_enable,chain_enable,edu_enable,hall_enable,multilang_enable,sns_enable,subsite_enable,supplier_enable,im_enable,chat_global,service_qrcode,wechat_mp_qrcode,chain_enable,baidu_client_ak,prodcut_addcart_flag"
	keyIds := gstr.Split(keyStr, ",")

	var list, error = service.ConfigBase().Find(ctx, &do.ConfigBaseListInput{Where: do.ConfigBase{ConfigKey: keyIds}})

	if error != nil {
		err = error
	}

	item := make(map[string]interface{})

	// 过滤删除敏感配置
	urlItem := []string{"ss"}

	for _, v := range list {
		// 过滤删除敏感配置
		if !gstr.InArray(urlItem, v.ConfigKey) {
			if v.ConfigDatatype == consts.RADIO || v.ConfigDatatype == consts.NUMBER || v.ConfigDatatype == consts.SELECT {
				item[v.ConfigKey] = gconv.Int(v.ConfigValue)
			} else {
				item[v.ConfigKey] = v.ConfigValue
			}
		}
	}

	//订单状态
	optionsList, err := service.ConfigBase().GetOrderStateList(ctx)
	item["order_state_list"] = optionsList

	//支付渠道
	channelList, err := service.ConfigBase().GetPaymentChannelList(ctx)
	item["payment_channel_list"] = channelList

	//退款退货 卖家处理状态
	returnStateList, err := service.ConfigBase().GetReturnStateList(ctx)
	item["return_state_list"] = returnStateList

	// user_level_name
	userLevelMap := make(map[uint]string)
	userLevelRateMap := make(map[uint]uint)

	// Simulating userLevels data
	userLevels, err := service.UserLevel().Find(ctx, &do.UserLevelListInput{})

	for _, userLevel := range userLevels {
		userLevelMap[userLevel.UserLevelId] = userLevel.UserLevelName
		userLevelRateMap[userLevel.UserLevelId] = userLevel.UserLevelRate
	}

	item["user_level_map"] = userLevelMap
	item["user_level_rate_map"] = userLevelRateMap

	// Error log types
	errStateList, err := service.ConfigBase().GetErrStateList(ctx)
	item["error_type_list"] = errStateList

	// User center menu
	userCenterMenu, err := service.PageBase().GetUserCenterMenu(ctx)
	item["user_center_menu"] = userCenterMenu

	// Project version
	version := utility.GetVersion(ctx)
	item["version"] = version
	item["site_version"] = version

	res = item

	return
}

// CleanCache 清理缓存
func (s *sConfigBase) CleanCache(ctx context.Context) (res bool, err error) {
	var keys []string

	keys, err = g.Redis().Keys(ctx, global.Namespace+"*")
	if err != nil {
		log.Error(ctx, err)
	}

	if !g.IsEmpty(keys) {
		num, err := g.Redis().Del(ctx, keys...)
		if err != nil {
			log.Error(ctx, err)
		}

		fmt.Println(num)
	}

	//err = g.Redis().FlushDB(ctx)
	//if err != nil {
	//	log.Error(ctx, err)
	//}

	return
}

// ifIm
func (s *sConfigBase) IfIm(ctx context.Context) (res bool) {
	return s.GetBool(ctx, "im_enable", false)
}

func (s *sConfigBase) IfB2B(ctx context.Context) (res bool) {
	return s.GetBool(ctx, "b2b_enable", false)
}

func (s *sConfigBase) IfChain(ctx context.Context) (res bool) {
	return s.GetBool(ctx, "chain_enable", false)
}

func (s *sConfigBase) IfEdu(ctx context.Context) (res bool) {
	return s.GetBool(ctx, "edu_enable", false)
}

func (s *sConfigBase) IfB2bHall(ctx context.Context) (res bool) {
	return s.GetBool(ctx, "hall_enable", false)
}

func (s *sConfigBase) IfMultilang(ctx context.Context) (res bool) {
	return s.GetBool(ctx, "multilang_enable", false)
}

func (s *sConfigBase) IfSns(ctx context.Context) (res bool) {
	return s.GetBool(ctx, "sns_enable", false)
}

func (s *sConfigBase) IfSubsite(ctx context.Context) (res bool) {
	return s.GetBool(ctx, "subsite_enable", false)
}
