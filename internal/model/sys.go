package model

import "golershop.cn/internal/model/entity"

// ImConfigVo IM配置
type ImConfigVo struct {
	Puid     int  `json:"puid,omitempty"`       // IM平台编号
	ImUserId uint `json:"im_user_id,omitempty"` // 用户编号
	ImEnable bool `json:"im_enable,omitempty"`  // 是否启用IM
}

// PagePopUpVo 页面弹窗
type PagePopUpVo struct {
	PopUpEnable bool   `json:"pop_up_enable,omitempty"` // 是否展示弹窗图片
	PopUpImage  string `json:"pop_up_image,omitempty"`  // 弹窗图片
	PopUpUrl    string `json:"pop_up_url,omitempty"`    // 弹窗网址
}

type PageDetail struct {
	entity.PageBase                // 继承自PageBase
	Im              *ImConfigVo    `json:"im,omitempty"`          // IM配置
	PageLoaded      bool           `json:"page_loaded,omitempty"` // 页面是否加载完毕
	PopUps          []*PagePopUpVo `json:"pop_ups,omitempty"`     // 弹窗集合
}

// PageDataItemVo 页面数据项VO
type PageDataItemVo struct {
	Id            uint64  `json:"id"`              // 主键
	Path          string  `json:"path"`            // 图片地址
	PathBg        string  `json:"path_bg"`         // 图片地址 fix for swipe
	Name          string  `json:"name"`            // 标题
	MarketPrice   float64 `json:"market_price"`    // 市场价
	ItemSalePrice float64 `json:"item_sale_price"` // 销售价
	ProductTips   string  `json:"product_tips"`    // 广告语
	AppUrl        string  `json:"app_url"`         // 访问网址
	AppId         string  `json:"app_id"`          // 小程序AppId
	MinAppUrl     string  `json:"min_app_url"`     // 小程序跳转的页面
	StartTime     string  `json:"start_time"`      // 开始时间
	StartTimeStr  string  `json:"start_time_str"`  // 开始时间
	EndTime       string  `json:"end_time"`        // 结束时间
	EndTimeStr    string  `json:"end_time_str"`    // 结束时间
	UserLimit     uint    `json:"user_limit"`      // {{items.UserLimit}}人团
	OrderCount    uint    `json:"order_count"`     // 已有{{items.OrderCount}}人参加
	FlexNum       uint    `json:"flex_num"`        // 表格宽度
	SelectType    uint    `json:"select_type"`     // selectType
	Did           uint64  `json:"did"`             // 主键
	SpecImg       string  `json:"spec_img"`        // 图片规格
	KeyWord       string  `json:"key_word"`        // 搜索关键字
	Words         string  `json:"words"`           // 富文本
}
