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
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// ConfigBase is the golang structure of table sys_config_base for DAO operations like Where/Data.
type ConfigBase struct {
	g.Meta         `orm:"table:sys_config_base, do:true"`
	ConfigKey      interface{} // 配置编码
	ConfigTitle    interface{} // 配置标题
	ConfigDatatype interface{} // 配置类型(ENUM):readonly-只读文本;number-数字;text-单行文本;textarea-多行文本;array-数组;password-密码;radio-单选框;checkbox-复选框;select-下拉框;icon-字体图标;date-日期;datetime-时间;image-单张图片;images-多张图片;file-单个文件;files-多个文件;ueditor-富文本编辑器;area-地区选择
	ConfigOptions  interface{} // 配置项
	ConfigValue    interface{} // 配置值
	ConfigTypeId   interface{} // 所属分类
	ConfigNote     interface{} // 配置注释
	ConfigSort     interface{} // 配置排序:从小到大
	ConfigEnable   interface{} // 是否启用(BOOL):0-禁用;1-启用
	ConfigBuildin  interface{} // 系统内置(BOOL):1-是; 0-否
}

type ConfigBaseListInput struct {
	ml.BaseList
	Where ConfigBase // 查询条件
}

type ConfigBaseListOutput struct {
	Items   []*entity.ConfigBase // 列表
	Page    int                  // 分页号码
	Total   int                  // 总页数
	Records int                  // 数据总数
	Size    int                  // 单页数量
}

type ConfigBaseListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
