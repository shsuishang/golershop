package cms

import "github.com/gogf/gf/v2/frame/g"

type ArticleListCategoryReq struct {
	g.Meta `path:"/front/cms/articleBase/listCategory" tags:"文章分类" method:"get" summary:"文章分类分页查询"`

	CategoryName string `json:"category_name" dc:"分类名称" v:"max-length:100#分类名称长度不能超过100字符"`
}

type ArticleListCategoryRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ArticleListReq struct {
	g.Meta `path:"/front/cms/articleBase/list" tags:"文章列表" method:"get" summary:"文章列表"`

	ArticleTitle string `json:"article_title"  dc:"标题"`                   // 标题
	CategoryId   int    `json:"category_id"    dc:"所属分类"`                 // 所属分类
	ArticleType  int    `json:"article_type"   dc:"文章类型(ENUM):1-文章;2-公告"` // 文章类型(ENUM):1-文章;2-公告
	UserId       int    `json:"user_id"        dc:"文章作者"`                 // 文章作者
	TagId        int    `json:"tag_id"         dc:"文章标签"`                 // 文章标签
	Page         int    `json:"page"           dc:"当前页码"`                 // 当前页码
	PageSize     int    `json:"page_size"      dc:"每页条数"`                 // 每页条数
	OrderField   string `json:"order_field"    dc:"排序字段"`                 // 排序字段
	OrderType    string `json:"order_type"     dc:"排序方式"`                 // 排序方式
}
type ArticleListRes struct {
	ArticleBaseAdd
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records interface{} `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
type ArticleDetailReq struct {
	g.Meta `path:"/front/cms/articleBase/get" tags:"文章内容" method:"get" summary:"文章内容"`

	ArticleId int `json:"article_id"              ` // 编号
}

type ArticleDetailRes struct {
	ArticleBaseAdd
	ArticleTagList interface{} `json:"article_tag_list" dc:"文章标签集合"` // 文章标签集合
	UserNickname   string      `json:"user_nickname" dc:"用户昵称"`      // 用户昵称
}
