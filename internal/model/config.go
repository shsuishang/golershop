package model

// 配置项类型
var CONFIG_DATATYPE_LIST = map[string]string{
	"text":     "单行文本",
	"textarea": "多行文本",
	"ueditor":  "富文本编辑器",
	"date":     "日期",
	"datetime": "时间",
	"number":   "数字",
	"select":   "下拉框",
	"radio":    "单选框",
	"checkbox": "复选框",
	"image":    "单张图片",
	"images":   "多张图片",
	"password": "密码",
	"icon":     "字体图标",
	"file":     "单个文件",
	"files":    "多个文件",
	"hidden":   "隐藏",
	"readonly": "只读文本",
}

type SelectVo struct {
	Value  uint   `json:"value"`
	Label  string `json:"label"`
	Enable bool   `json:"enable"`
	Ext1   string `json:"ext1"`
	Ext2   string `json:"ext2"`
}
