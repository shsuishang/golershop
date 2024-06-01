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
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/xlog"
	"github.com/lihao1988/php2go/file"
	"golershop.cn/internal/service"
	utility "golershop.cn/utility"
)

type sPaymentWechat struct{}

func init() {
	service.RegisterPaymentWechat(NewPaymentWechat())
}

func NewPaymentWechat() *sPaymentWechat {
	return &sPaymentWechat{}
}

// GetClient 初始化微信v3客户端并做配置
func (s *sPaymentWechat) GetClient(ctx context.Context) (client *wechat.ClientV3, err error) {
	MchId := service.ConfigBase().GetStr(ctx, "wechat_pay_mchid", "")
	SerialNo := service.ConfigBase().GetStr(ctx, "wechat_pay_serial_no", "")
	APIv3Key := service.ConfigBase().GetStr(ctx, "wechat_pay_v3_key", "")
	privateKeyPath := service.ConfigBase().GetStr(ctx, "wechat_pay_apiclient_key", "")

	path := utility.UploadPath(ctx)
	privateKeyPath = path + privateKeyPath

	privateKey, err := file.FileGetContents(privateKeyPath)
	if err != nil {
		xlog.Error(err)
		return
	}
	//config, err := service.ConfigBase().Get(ctx, "wechat_pay_apiclient_cert")

	// NewClientV3 初始化微信客户端 V3
	//	mchid：商户ID
	// 	serialNo：商户证书的证书序列号
	//	apiV3Key：APIv3Key，商户平台获取
	//	privateKey：商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
	client, err = wechat.NewClientV3(MchId, SerialNo, APIv3Key, privateKey)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 设置微信平台证书和序列号，如开启自动验签，请忽略此步骤
	//client.SetPlatformCert([]byte(""), "")

	// 启用自动同步返回验签，并定时更新微信平台API证书
	err = client.AutoVerifySign()
	if err != nil {
		xlog.Error(err)
		return
	}

	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOn

	return
}
