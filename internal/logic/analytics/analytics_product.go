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
package analytics

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/service"
)

type sAnalyticsProduct struct{}

func init() {
	service.RegisterAnalyticsProduct(NewAnalyticsProduct())
}

func NewAnalyticsProduct() *sAnalyticsProduct {
	return &sAnalyticsProduct{}
}

// 获取商品数量
func (s *sAnalyticsProduct) GetProductNum(ctx context.Context, input *model.AnalyticsProductInput) (*model.AnalyticsNumOutput, error) {
	topRes := &model.AnalyticsNumOutput{}

	// 获取当前周期内数据
	currentProductNum, err := dao.AnalyticsProduct.GetProductNum(ctx, input)
	if err != nil {
		return nil, err
	}
	if currentProductNum != nil {
		topRes.Current = currentProductNum
	}

	// 上个周期
	preInput := &model.AnalyticsProductInput{}
	gconv.Struct(input, preInput)

	if input.Stime > 0 && input.Etime > 0 {
		preInput.Stime = input.Stime - (input.Etime - input.Stime)
		preInput.Etime = input.Stime

		preProductNum, err := dao.AnalyticsProduct.GetProductNum(ctx, preInput)
		if err != nil {
			return nil, err
		}

		if !g.IsEmpty(preProductNum) {
			topRes.Pre = preProductNum

			// 计算日环比 日环比 = (当日数据 - 前一日数据) / 前一日数据 * 100%
			topRes.Daym2m = 0.0
			if gconv.Float64(preProductNum) != 0 {
				topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(currentProductNum)-gconv.Float64(preProductNum))/gconv.Float64(preProductNum)*100)
			}
		}
	}

	return topRes, nil
}
