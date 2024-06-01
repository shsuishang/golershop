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
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/xlog"
	"golershop.cn/internal/dao/global"
	"golershop.cn/internal/service"
	"golershop.cn/utility"
	"os"
)

type sPaymentAlipay struct{}

func init() {
	service.RegisterPaymentAlipay(NewPaymentAlipay())
}

func NewPaymentAlipay() *sPaymentAlipay {
	return &sPaymentAlipay{}
}

// GetClient 初始化支付宝客户端并做配置
func (s *sPaymentAlipay) GetClient(ctx context.Context, tradeNo string) (client *alipay.Client, err error) {
	// 初始化支付宝客户端
	// appid：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。

	appId := service.ConfigBase().GetStr(ctx, "alipay_app_id", "")
	privateKey := service.ConfigBase().GetStr(ctx, "alipay_rsa_private_key", "")

	path := utility.UploadPath(ctx)

	//privateKeyPath = path + privateKeyPath
	//
	//privateKey, err := utility.FileGetContents(privateKeyPath)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}

	client, err = alipay.NewClient(appId, privateKey, true)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).                                                                // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).                                                               // 设置签名类型，不设置默认 RSA2
							SetReturnUrl(global.BaseUrl + "/front/pay/callback/returnUrl?out_trade_no=" + tradeNo). // 设置返回URL
							SetNotifyUrl(global.BaseUrl + "/front/pay/callback/alipayNotify")                       // 设置异步通知URL
	//.SetAppAuthToken() // 设置第三方应用授权

	// 设置biz_content加密KEY，设置此参数默认开启加密（目前不可用，设置后会报错）
	//client.SetAESKey("1234567890123456")

	// 公钥
	//alipayPublicCertPath := service.ConfigBase().GetStr(ctx, "alipay_rsa_public_key", "")
	//alipayPublicCertPath = path + alipayPublicCertPath
	//alipayPublicCert, err := os.ReadFile(alipayPublicCertPath)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}

	// 应用证书路径
	appPublicCertPath := service.ConfigBase().GetStr(ctx, "alipay_app_cert_path", "")
	appPublicCertPath = path + appPublicCertPath

	// 支付宝证书路径
	alipayPublicCertPath := service.ConfigBase().GetStr(ctx, "alipay_cert_path", "")
	alipayPublicCertPath = path + alipayPublicCertPath
	alipayPublicCert, err := os.ReadFile(alipayPublicCertPath)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 支付宝证书根路径
	alipayRootCertPath := service.ConfigBase().GetStr(ctx, "alipay_root_cert_path", "")
	alipayRootCertPath = path + alipayRootCertPath

	// 自动同步验签（只支持证书模式）
	// 传入 alipayPublicCert.crt 内容
	client.AutoVerifySign(alipayPublicCert)

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	//err = client.SetCertSnByPath("appPublicCert.crt", "alipayRootCert.crt", "alipayPublicCert.crt")
	err = client.SetCertSnByPath(appPublicCertPath, alipayRootCertPath, alipayPublicCertPath)
	// 证书内容
	//err := client.SetCertSnByContent("appPublicCert.crt bytes", "alipayRootCert bytes", "alipayPublicCert.crt bytes")

	if err != nil {
		xlog.Error(err)
		return
	}

	return
}
