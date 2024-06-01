package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ContractTypeAdd struct {
	ContractTypeId      uint    `json:"contract_type_id"      ` // 保障编号
	ContractTypeName    string  `json:"contract_type_name"    ` // 保障名称
	ContractTypeDesc    string  `json:"contract_type_desc"    ` // 保障简写
	ContractTypeText    string  `json:"contract_type_text"    ` // 保障描述
	ContractTypeDeposit float64 `json:"contract_type_deposit" ` // 保证金
	ContractTypeIcon    string  `json:"contract_type_icon"    ` // 项目图标
	ContractTypeUrl     string  `json:"contract_type_url"     ` // 说明网址
	ContractTypeOrder   uint    `json:"contract_type_order"   ` // 保障排序
	ContractTypeEnable  bool    `json:"contract_type_enable"  ` // 是否开启(BOOL):0-关闭;1-开启
	ContractTypeBuildin bool    `json:"contract_type_buildin" ` // 系统内置(BOOL): 0-非内置;1-系统内置
}
type ContractTypeEditReq struct {
	g.Meta `path:"/manage/sys/contractType/edit" tags:"保障服务" method:"post" summary:"保障服务编辑接口"`

	ContractTypeId uint `json:"contract_type_id"   ` // 保障服务编号`
	ContractTypeAdd
}

type ContractTypeEditRes struct {
	ContractTypeId interface{} `json:"contract_type_id"   dc:"保障服务信息"`
}

type ContractTypeAddReq struct {
	g.Meta `path:"/manage/sys/contractType/add" tags:"保障服务" method:"post" summary:"保障服务编辑接口"`

	ContractTypeAdd
}

type ContractTypeRemoveReq struct {
	g.Meta         `path:"/manage/sys/contractType/remove" tags:"保障服务" method:"post" summary:"保障服务删除接口"`
	ContractTypeId string `json:"contract_type_id"   ` // 保障服务编号
}

type ContractTypeRemoveRes struct {
}

type ContractTypeListReq struct {
	g.Meta `path:"/manage/sys/contractType/list" tags:"保障服务" method:"get" summary:"保障服务列表接口"`
	ml.BaseList

	ContractTypeName string `json:"contract_type_name"    ` // 保障名称
}

type ContractTypeListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
