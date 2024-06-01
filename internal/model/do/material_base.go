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

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// MaterialBase is the golang structure of table sys_material_base for DAO operations like Where/Data.
type MaterialBase struct {
	g.Meta           `orm:"table:sys_material_base, do:true"`
	MaterialId       interface{} // 素材编号
	MaterialNumber   interface{} // 附件md5
	GalleryId        interface{} // 分类编号
	StoreId          interface{} // 店铺编号
	UserId           interface{} // 用户编号
	SubsiteId        interface{} // 所属分站:0-总站
	MaterialUrl      interface{} // 文件URL
	MaterialSource   interface{} // 素材来源
	MaterialSort     interface{} // 素材排序
	MaterialPath     interface{} // 素材path:本地存储
	MaterialType     interface{} // 素材类型
	MaterialImageH   interface{} // 素材高度
	MaterialImageW   interface{} // 素材宽度
	MaterialSize     interface{} // 素材大小
	MaterialMimeType interface{} // 素材类型
	MaterialAlt      interface{} // 素材alt
	MaterialName     interface{} // 素材标题
	MaterialDesc     interface{} // 素材描述
	MaterialDuration interface{} // 素材时长:（音频/视频）
	MaterialTime     *gtime.Time // 素材日期
}

type MaterialBaseListInput struct {
	ml.BaseList
	Where MaterialBase // 查询条件
}

type MaterialBaseListOutput struct {
	Items   []*entity.MaterialBase // 列表
	Page    int                    // 分页号码
	Total   int                    // 总页数
	Records int                    // 数据总数
	Size    int                    // 单页数量
}

type MaterialBaseListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
