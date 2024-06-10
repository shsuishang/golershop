package sys

import (
	"context"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/service"
	utility "golershop.cn/utility/rsa"
)

var (
	Release = cRelease{}
)

type cRelease struct{}

// =========================== 管理端使用 =============================

// Download
func (c *cRelease) Download(ctx context.Context, req *sys.DownloadReq) (res sys.DownloadRes, err error) {
	var result, error = service.Cloud().GetRelease(&ml.ReleaseDto{
		ServiceAppId:  100,
		ServiceUserId: service.ConfigBase().GetInt(ctx, "service_user_id", 0),
		ServiceAppKey: service.ConfigBase().GetStr(ctx, "service_app_key", ""),
		Version:       "1.0.2",
		Url:           req.UrlApi,
		AppId:         service.ConfigBase().GetStr(ctx, "wechat_xcx_app_id", ""),
		AppName:       service.ConfigBase().GetStr(ctx, "site_name", ""),
		PrimaryColor:  req.PrimaryColor,
	}, req.ReleaseType)

	if error != nil {
		err = error
	}

	res = sys.DownloadRes(result)

	return
}

// GetLicence
func (c *cRelease) GetLicence(ctx context.Context, req *sys.LicenceReq) (res *sys.LicenceRes, err error) {
	licence := service.ConfigBase().GetStr(ctx, "licence", "2+WJU0eC2hmoby4Fx3PHTu1dT1IRV5wMIp24otk29WtRThXgdzWPSyHdl0oeNkFNATwI8m+EXBJPAsd0Bo/WpyVkPzFOz7WAh1AMmYmFN1tI7nMCTWC2UmyP+cBuZukq0AalVfbMyzb9ll+t5LGwGD44DKS3CdsEIhqJEr1JUFCsj/D92lK0XB2pyZgyJCnQxDj2dapPbgsNXzsePheIaqQ3v2YHGViQ11ypUJDCwqO57HcNQLLXW0P0v2tCdGzkY9oi3G4Z1mC2ob15cim4baKgnkHBjg5RgvHqkNd1PC+ePeiwHDdDN7nToM3rOFFzt84+G4twZcBI335Gjyu/yg==")
	licencePublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5Fa6zAE5sJ9y1qzjhGAFWggKbMAS82xNG8wKglG0k6XojKBTw/7evSwC0aEgYU/BkIPzxNb7j6Oap5iZ43YgFgLI1dGalZJnvmLTmRK4+MqkKI6LlQQyCTZtSDkPNr62jYzSya+pPt0hgBHgk2x6YAns83SYmZf+7NT3qB+uxIgVIJTcO6m+SX3MyQU6cRKlt46+A9GwYiPx6davGxiX4TeeS5/sWiV1+yBb1xovNPjGK9d6N/3ObvSPtNXLnFn5jtwT1UanZPdZMR+oYIlltR9QGE3jnaTxlYTUhkY63GMek94dWbJTBQqpaA6t221iCwh8uawX4sbm4ZoRTy8SRwIDAQAB"

	decrypt, err := utility.PubkeyDecrypt(licence, licencePublicKey)

	if err != nil {
		return nil, err
	}

	res = &sys.LicenceRes{}
	res.LicenceStr = decrypt
	res.IsAuthorized = true

	return res, err
}
