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

// UserMessage is the golang structure of table account_user_message for DAO operations like Where/Data.
type UserMessage struct {
	g.Meta            `orm:"table:account_user_message, do:true"`
	MessageId         interface{} // 消息编号
	MessageParentId   interface{} // 上级编号
	UserId            interface{} // 所属用户:发送者或者接收者，如果message_kind=1则为当前用户发送的消息。
	UserNickname      interface{} // 用户昵称
	MessageKind       interface{} // 消息种类(ENUM):1-发送消息;2-接收消息
	UserOtherId       interface{} // 相关用户:发送者或者接收者
	UserOtherNickname interface{} // 相关昵称:发送者或者接收者
	MessageTitle      interface{} // 消息标题
	MessageContent    interface{} // 消息内容
	MessageTime       interface{} // 发送时间
	MessageIsRead     interface{} // 是否读取(BOOL):0-未读;1-已读
	MessageIsDelete   interface{} // 是否删除(BOOL):0-正常状态;1-删除状态
	MessageType       interface{} // 消息类型(ENUM):1-系统消息;2-用户消息
	MessageCat        interface{} // 消息类型(ENUM):text-文本消息;img-图片消息;video-视频消息;file:文件;location:位置;redpack:红包
	MessageDataType   interface{} // 消息分类(ENUM):0-默认消息;1-公告消息;2-订单消息;3-商品消息;4-余额卡券;5-服务消息
	MessageDataId     interface{} // 消息数据:商品编号|订单编号
	MessageLength     interface{} // 消息长度
	MessageW          interface{} // 图片宽度
	MessageH          interface{} // 图片高度
}

type UserMessageListInput struct {
	ml.BaseList
	Where UserMessage // 查询条件
}

type UserMessageListOutput struct {
	Items   []*entity.UserMessage // 列表
	Page    int                   // 分页号码
	Total   int                   // 总页数
	Records int                   // 数据总数
	Size    int                   // 单页数量
}

type UserMessageListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
