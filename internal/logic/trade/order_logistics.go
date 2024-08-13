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

package trade

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"net/url"
)

type sOrderLogistics struct{}

func init() {
	service.RegisterOrderLogistics(NewOrderLogistics())
}

func NewOrderLogistics() *sOrderLogistics {
	return &sOrderLogistics{}
}

// Get 读取订单
func (s *sOrderLogistics) Get(ctx context.Context, id any) (out *entity.OrderLogistics, err error) {
	var list []*entity.OrderLogistics
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 读取多条订单
func (s *sOrderLogistics) Gets(ctx context.Context, id any) (list []*entity.OrderLogistics, err error) {
	err = dao.OrderLogistics.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sOrderLogistics) Find(ctx context.Context, in *do.OrderLogisticsListInput) (out []*entity.OrderLogistics, err error) {
	out, err = dao.OrderLogistics.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sOrderLogistics) List(ctx context.Context, in *do.OrderLogisticsListInput) (out *do.OrderLogisticsListOutput, err error) {
	out, err = dao.OrderLogistics.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sOrderLogistics) Add(ctx context.Context, in *do.OrderLogistics) (lastInsertId int64, err error) {
	lastInsertId, err = dao.OrderLogistics.Add(ctx, in)
	if err != nil {
		return 0, err
	}

	service.Order().CheckShippingComplete(ctx, in.OrderId.(string))

	return lastInsertId, err
}

// Edit 编辑
func (s *sOrderLogistics) Edit(ctx context.Context, in *do.OrderLogistics) (affected int64, err error) {
	_, err = dao.OrderLogistics.Edit(ctx, in.OrderLogisticsId, in)

	if err != nil {
		return 0, err
	}

	service.Order().CheckShippingComplete(ctx, in.OrderId.(string))

	return
}

// Remove 删除多条记录模式
func (s *sOrderLogistics) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.OrderLogistics.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

var stateMap map[string]string

func init() {
	stateMap = map[string]string{
		"0":   "没有记录",
		"1":   "已揽收",
		"2":   "运输途中",
		"201": "到达目的城市",
		"202": "派件中",
		"211": "已投放快递柜或驿站",
		"3":   "已签收",
		"301": "正常签收",
		"302": "派件异常后最终签收",
		"304": "代收签收",
		"311": "快递柜或驿站签收",
		"4":   "问题件",
		"401": "发货无信息",
		"402": "超时未签收",
		"403": "超时未更新",
		"404": "拒收(退件)",
		"405": "派件异常",
		"406": "退货签收",
		"407": "退货未签收",
		"412": "快递柜或驿站超时未取",
	}
}

// OrderOnlineByJson 组装快递查询请求参数，并发送请求
func (s *sOrderLogistics) OrderOnlineByJson(ctx context.Context, orderTrackingNumber, shipperCode, CustomerName string) (result string, err error) {
	// 组装应用级参数
	requestData := fmt.Sprintf(
		"{'OrderCode': '', 'shipperCode': '%s', 'CustomerName': '%s', 'logisticCode': '%s'}",
		shipperCode, CustomerName, orderTrackingNumber,
	)

	// 组装系统级参数
	params := g.Map{}
	appId := service.ConfigBase().GetStr(ctx, "kuaidiniao_e_business_id", "")
	appKey := service.ConfigBase().GetStr(ctx, "kuaidiniao_app_key", "")

	// URL编码函数
	urlEncoder := func(data, charset string) string {
		return url.QueryEscape(data)
	}

	params["RequestData"] = urlEncoder(requestData, "UTF-8")
	params["EBusinessID"] = appId
	params["RequestType"] = "8002" // 快递查询接口指令8002/地图版快递查询接口指令8004

	dataSign, err := Encrypt(requestData, appKey, "UTF-8")
	if err != nil {
		return "", err
	}
	params["DataSign"] = urlEncoder(dataSign, "UTF-8")
	params["DataType"] = "2"

	requestUrl := "https://api.kdniao.com/Ebusiness/EbusinessOrderHandle.aspx"
	// 发送POST请求
	response, err := g.Client().Post(ctx, requestUrl, params)
	if err != nil {
		return "", err
	}
	defer response.Close()

	// 获取响应内容
	result = string(response.ReadAll())
	return result, nil
}

func Encrypt(content, keyValue, charset string) (string, error) {
	if keyValue != "" {
		return base64Encode(md5Hash(content+keyValue, charset), charset)
	}
	return base64Encode(md5Hash(content, charset), charset)
}

func base64Encode(str, charset string) (string, error) {
	// 将字符串转换为字节数组并进行Base64编码
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return encoded, nil
}

func md5Hash(str, charset string) string {
	md5Ctx := md5.New()
	strBytes := []byte(str)
	md5Ctx.Write(strBytes)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// ReturnLogistics 查询物流信息
func (s *sOrderLogistics) ReturnLogistics(ctx context.Context, returnTrackingName, returnTrackingNumber string) (resultMap g.Map, err error) {
	// 创建查询条件
	expressBase, err := dao.ExpressBase.FindOne(ctx, &do.ExpressBaseListInput{
		Where: do.ExpressBase{
			ExpressName: returnTrackingName,
		},
	})

	if err != nil {
		return nil, err
	}

	if expressBase == nil {
		return nil, gerror.New("系统中未配置该物流信息，请检查发货信息是否正确！")
	}

	// 调用第三方物流查询接口，获取物流信息
	logisticsInfoStr, err := s.OrderOnlineByJson(ctx, returnTrackingNumber, expressBase.ExpressPinyin, "")
	if err != nil {
		return nil, err
	}

	// 解析物流信息
	logisticsInfo := gjson.New(logisticsInfoStr)
	state := logisticsInfo.Get("State").Int()
	if state == 0 {
		reason := logisticsInfo.Get("Reason").String()
		return nil, gerror.New(fmt.Sprintf("非系统错误，请联系管理员检查物流配置项，或检查发货信息是否真实有效！错误信息：{%s}", reason))
	}

	StateEx := logisticsInfo.Get("StateEx").String()
	if stateMap[StateEx] == "" {
		return nil, gerror.New("物流状态异常")
	}

	// 构建返回结果
	resultMap = g.Map{
		"shipperCode":   logisticsInfo.Get("ShipperCode").String(),
		"logisticCode":  logisticsInfo.Get("LogisticCode").String(),
		"state":         state,
		"stateEx":       StateEx,
		"express_state": stateMap[StateEx],
		"traces":        logisticsInfo.Get("Traces").String(),
	}

	return resultMap, nil
}
