package chain

import "github.com/gogf/gf/v2/frame/g"

type ChainListReq struct {
	g.Meta `path:"/front/o2o/chain/list" tags:"门店列表" method:"get" summary:"门店列表"`
}

type ChainListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records interface{} `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
