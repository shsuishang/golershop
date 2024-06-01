package model

type MenuKeysOutput struct {
	MenuId []uint `json:"menu_id"         ` // 菜单编号
}

type Meta struct {
	MenuTitle string `json:"title"     ` // 菜单名称

	MenuClose  bool   `json:"noClosable"     ` // 允许关闭(BOOL):0-禁止;1-允许
	MenuHidden bool   `json:"hidden"    `      // 是否隐藏(BOOL):0-展示；1-隐藏
	MenuIcon   string `json:"icon"      `      // 图标设置
	MenuDot    bool   `json:"dot"       `      // 是否红点(BOOL):0-隐藏；1-显示
	MenuBubble string `json:"badge"    `       // 菜单标签
}

// Menu is the golang structure for table sys_menu.
type Menu struct {
	MenuId       uint   `json:"menu_id"         ` // 菜单编号
	MenuParentId uint   `json:"menu_parent_id"  ` // 菜单父编号
	MenuTitle    string `json:"menu_title"      ` // 菜单名称
	MenuUrl      string `json:"menu_url"        ` // 页面网址

	MenuName      string                 `json:"name"      `   // VUE名称
	MenuPath      string                 `json:"path"      `   // VUE路径
	MenuParam     map[string]interface{} `json:"params"      ` // 参数
	MenuComponent string                 `json:"component" `   // 组件名称
	MenuRedirect  string                 `json:"redirect"  `   // 重定向

	MenuClose      bool   `json:"menu_close"      ` // 允许关闭(BOOL):0-禁止;1-允许
	MenuHidden     bool   `json:"menu_hidden"     ` // 是否隐藏(BOOL):0-展示；1-隐藏
	MenuIcon       string `json:"menu_icon"       ` // 图标设置
	MenuDot        bool   `json:"menu_dot"        ` // 是否红点(BOOL):0-隐藏；1-显示
	MenuBubble     string `json:"menu_bubble"     ` // 菜单标签
	MenuSort       int    `json:"menu_sort"      `  // 菜单排序
	MenuType       int    `json:"menu_type"       ` // 菜单类型(LIST):0-按钮1-菜单
	MenuFunc       string `json:"menu_func"       ` // 功能开启:设置config_key
	MenuPermission string `json:"menu_permission" ` // 权限标识:请求地址
	MenuBuildin    bool   `json:"menu_buildin"    ` // 系统内置(BOOL):1-是; 0-否

	Meta Meta `json:"meta"  `
}

// 菜单Vo
type TreeNode struct {
	Menu
	Children []*TreeNode `json:"children"` // 子菜单
}
