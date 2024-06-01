package cms

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

// start fo front

// start fo manage
type ArticleCategoryAdd struct {
	CategoryId       uint   `json:"category_id"        ` // 分类编号
	CategoryName     string `json:"category_name"      ` // 分类名称
	CategoryParentId int    `json:"category_parent_id" ` // 上级编号
	CategoryImageUrl string `json:"category_image_url" ` // 分类图标
	CategoryKeywords string `json:"category_keywords"  ` // 分类关键词
	CategoryDesc     string `json:"category_desc"      ` // 分类描述
	CategoryCount    uint   `json:"category_count"     ` // 内容数量
	CategoryTemplate string `json:"category_template"  ` // 分类模板
	CategoryAlias    string `json:"category_alias"     ` // 分类别名
	CategoryOrder    uint   `json:"category_order"     ` // 分类排序
	CategoryBuildin  bool   `json:"category_buildin"   ` // 系统内置(BOOL):0-非内置;1-内置;
	CategoryIsLeaf   bool   `json:"category_is_leaf"   ` // 叶节点(BOOL):0-否;1-是;
}
type ArticleCategoryEditReq struct {
	g.Meta `path:"/manage/cms/articleCategory/edit" tags:"文章分类" method:"post" summary:"分类编辑接口"`

	CategoryId uint `json:"category_id"              ` // 分类编号
	ArticleCategoryAdd
}

type ArticleCategoryEditRes struct {
	CategoryId interface{} `json:"category_id"   dc:"文章分类信息"`
}
type ArticleCategoryEditStateReq struct {
	g.Meta `path:"/manage/cms/articleCategory/editState" tags:"文章分类" method:"post" summary:"分类编辑接口"`

	CategoryId uint `json:"category_id"              ` // 分类编号
	ArticleCategoryAdd
}

type ArticleCategoryEditStateRes struct {
	CategoryId interface{} `json:"category_id"   dc:"文章分类信息"`
}

type ArticleCategoryAddReq struct {
	g.Meta `path:"/manage/cms/articleCategory/add" tags:"文章分类" method:"post" summary:"文章分类添加接口"`

	ArticleCategoryAdd
}

type ArticleCategoryRemoveReq struct {
	g.Meta     `path:"/manage/cms/articleCategory/remove" tags:"文章分类" method:"post" summary:"文章分类删除接口"`
	CategoryId []uint `json:"category_id" v:"required#请输入文章分类编号"   dc:"文章分类信息"`
}

type ArticleCategoryRemoveRes struct {
}

type ArticleCategoryListReq struct {
	g.Meta `path:"/manage/cms/articleCategory/list" tags:"文章分类" method:"get" summary:"文章分类列表接口"`
	Page   int `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size   int `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`
}

type ArticleCategoryListRes struct {
	Items   interface{} `json:"items"    dc:"分类列表"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ArticleCategoryTreeReq struct {
	g.Meta       `path:"/manage/cms/articleCategory/tree" tags:"文章分类" method:"get" summary:"后台文章分类Tree"`
	CategoryName string `json:"category_name"  d:""  dc:"搜索关键词"`
}

//type ArticleCategoryTreeRes model.TreeNode

//res []*v1.ArticleCategoryTreeRes,

/*
type ArticleCategoryTreeRes struct {
	model.ArticleCategory
	Children []*model.TreeNode `json:"children"` // 子文章分类
}
*/

type ArticleCategoryTreeRes []*model.ArticleCategoryTreeNode
