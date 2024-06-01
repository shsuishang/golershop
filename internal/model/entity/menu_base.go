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

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MenuBase is the golang structure for table menu_base.
type MenuBase struct {
	MenuId         uint        `json:"menu_id"         ` // 菜单编号
	MenuParentId   uint        `json:"menu_parent_id"  ` // 菜单父编号
	MenuTitle      string      `json:"menu_title"      ` // 菜单名称
	MenuUrl        string      `json:"menu_url"        ` // 页面网址
	MenuName       string      `json:"menu_name"       ` // 组件名称
	MenuPath       string      `json:"menu_path"       ` // 组件路由
	MenuComponent  string      `json:"menu_component"  ` // 组件路径
	MenuRedirect   string      `json:"menu_redirect"   ` // 重定向
	MenuClose      bool        `json:"menu_close"      ` // 允许关闭(BOOL):0-禁止;1-允许
	MenuHidden     bool        `json:"menu_hidden"     ` // 是否隐藏(BOOL):0-展示;1-隐藏
	MenuEnable     bool        `json:"menu_enable"     ` // 是否启用(BOOL):0-禁用;1-启用
	MenuClass      string      `json:"menu_class"      ` // 样式class
	MenuIcon       string      `json:"menu_icon"       ` // 图标设置
	MenuDot        bool        `json:"menu_dot"        ` // 是否红点(BOOL):0-隐藏;1-显示
	MenuBubble     string      `json:"menu_bubble"     ` // 菜单标签
	MenuSort       uint        `json:"menu_sort"       ` // 菜单排序
	MenuType       uint        `json:"menu_type"       ` // 菜单类型(LIST):0-按钮;1-菜单
	MenuNote       string      `json:"menu_note"       ` // 备注
	MenuFunc       string      `json:"menu_func"       ` // 功能开启:设置config_key
	MenuRole       uint        `json:"menu_role"       ` // 角色类型(LIST):1-平台;2-商户;3-门店
	MenuParam      string      `json:"menu_param"      ` // url参数
	MenuPermission string      `json:"menu_permission" ` // 权限标识:后端地址
	MenuBuildin    bool        `json:"menu_buildin"    ` // 系统内置(BOOL):1-是; 0-否
	MenuTime       *gtime.Time `json:"menu_time"       ` // 最后更新时间
}
