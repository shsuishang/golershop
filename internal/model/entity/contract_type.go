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

// ContractType is the golang structure for table contract_type.
type ContractType struct {
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
