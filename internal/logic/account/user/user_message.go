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

package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility"
	"golershop.cn/utility/array"
	"golershop.cn/utility/mtime"
	"time"
)

type sUserMessage struct{}

func init() {
	service.RegisterUserMessage(NewUserMessage())
}

func NewUserMessage() *sUserMessage {
	return &sUserMessage{}
}

// Find 查询数据
func (s *sUserMessage) Find(ctx context.Context, in *do.UserMessageListInput) (out []*entity.UserMessage, err error) {
	out, err = dao.UserMessage.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserMessage) List(ctx context.Context, in *do.UserMessageListInput) (out *do.UserMessageListOutput, err error) {
	out, err = dao.UserMessage.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserMessage) Add(ctx context.Context, in *do.UserMessage) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserMessage.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserMessage) Edit(ctx context.Context, in *do.UserMessage) (affected int64, err error) {
	_, err = dao.UserMessage.Edit(ctx, in.MessageId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserMessage) Remove(ctx context.Context, id any) (affected int64, err error) {
	//是否内置
	one, err := dao.UserMessage.Get(ctx, id)
	if one.MessageIsDelete {
		return 0, errors.New("已经删除，不可删除")
	}

	affected, err = dao.UserMessage.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetMsgCount 信消息数量
func (s *sUserMessage) GetMsgCount(ctx context.Context, in *model.MsgCountInput) (out *model.UserMessageVo, err error) {
	userMessageRes := &model.UserMessageVo{}

	// 创建查询条件
	messageQueryWrapper := &do.UserMessageListInput{
		Where: do.UserMessage{
			UserId:        in.UserId,
			MessageKind:   2,
			MessageIsRead: 0,
		},
	}

	// 统计未读消息数量
	count, err := dao.UserMessage.Count(ctx, messageQueryWrapper)
	if err != nil {
		return nil, err
	}
	userMessageRes.Num = int(count)

	// 构造最后一个用户聊天消息
	if count > 0 {
		// 创建查询条件
		queryWrapper := &do.UserMessageListInput{
			Where: do.UserMessage{
				UserId:        in.UserId,
				MessageKind:   2,
				MessageIsRead: 0,
			},
		}
		queryWrapper.Sidx = dao.UserMessage.Columns().MessageTime
		queryWrapper.Sort = ml.ORDER_BY_DESC

		// 如果recentlyFlag为true，查询最近5分钟内的消息
		if in.RecentlyFlag {
			time := time.Now().UnixMilli()
			var likes = []*ml.WhereExt{{
				Column: dao.UserMessage.Columns().MessageTime,
				Val:    time,
				Symbol: ml.GE,
			}}
			queryWrapper.WhereExt = likes
		}

		// 查询最后一条消息
		message, err := dao.UserMessage.FindOne(ctx, queryWrapper)
		if err != nil {
			return nil, err
		}

		if message != nil {
			userMessageRes.MsgRow = message
		}
	}

	return userMessageRes, nil
}

// GetList 获取用户消息列表
func (s *sUserMessage) GetList(ctx context.Context, input *do.UserMessageListInput) (output *model.UserMessageOutput, err error) {
	output = &model.UserMessageOutput{}
	userMessagePage, err := dao.UserMessage.List(ctx, input)
	if err != nil {
		return nil, err
	}

	if userMessagePage != nil && len(userMessagePage.Items) > 0 {
		output.Records = userMessagePage.Records
		output.Size = userMessagePage.Size
		output.Total = userMessagePage.Total
		output.Page = userMessagePage.Page
		gconv.Struct(userMessagePage.Items, &output.Items)

		userMessageList := output.Items

		// 用户头像
		userMap, err := s.getUserAvatarMap(ctx, userMessageList)
		if err != nil {
			return nil, err
		}

		// 相关用户头像
		otherUserMap, err := s.getOtherUserAvatarMap(ctx, userMessageList)
		if err != nil {
			return nil, err
		}

		for _, userMessage := range userMessageList {
			//messageRes := &model.UserMessageVo{}
			//gconv.Struct(userMessage, messageRes)

			if userMessage.MessageKind == 1 {
				if userMap != nil && len(userMap) > 0 {
					userMessage.UserAvatar = userMap[userMessage.UserId]
				}
			} else {
				if otherUserMap != nil && len(otherUserMap) > 0 {
					userMessage.UserOtherAvatar = otherUserMap[userMessage.UserOtherId]
				}
			}
		}
	}

	return output, nil
}

// getOtherUserAvatarMap 获取其他用户头像Map
func (s *sUserMessage) getOtherUserAvatarMap(ctx context.Context, userMessageList []*model.UserMessageVo) (map[uint]string, error) {
	otherUserIds := make([]uint, 0)
	for _, userMessage := range userMessageList {
		if userMessage.UserOtherId != 0 {
			otherUserIds = append(otherUserIds, userMessage.UserOtherId)
		}
	}
	otherUserInfos, err := service.UserInfo().Gets(ctx, otherUserIds)
	if err != nil {
		return nil, err
	}

	if len(otherUserInfos) == 0 {
		return nil, nil
	}

	otherUserMap := make(map[uint]string)
	for _, userInfo := range otherUserInfos {
		otherUserMap[userInfo.UserId] = userInfo.UserAvatar
	}

	return otherUserMap, nil
}

// getUserAvatarMap 获取用户头像Map
func (s *sUserMessage) getUserAvatarMap(ctx context.Context, userMessageList []*model.UserMessageVo) (map[uint]string, error) {
	userIds := make([]uint, 0)
	for _, userMessage := range userMessageList {
		userIds = append(userIds, userMessage.UserId)
	}
	userInfos, err := service.UserInfo().Gets(ctx, userIds)
	if err != nil {
		return nil, err
	}

	if len(userInfos) == 0 {
		return nil, nil
	}

	userMap := make(map[uint]string)
	for _, userInfo := range userInfos {
		userMap[userInfo.UserId] = userInfo.UserAvatar
	}

	return userMap, nil
}

// getById 获取用户消息详情
func (s *sUserMessage) GetById(ctx context.Context, messageId, userId uint) (*entity.UserMessage, error) {
	userMessage, err := dao.UserMessage.Get(ctx, messageId)
	if err != nil {
		return nil, err
	}

	if userMessage == nil {
		return nil, nil
	}

	if userId == userMessage.UserId {
		message := &do.UserMessage{
			MessageIsRead: true,
		}
		_, err = dao.UserMessage.Edit(ctx, messageId, message)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	return userMessage, nil
}

// setRead 设置消息为已读
func (s *sUserMessage) SetRead(ctx context.Context, messageId, userOtherId, userId uint) (bool, error) {
	// 创建查询条件
	messageQueryWrapper := &do.UserMessageListInput{
		Where: do.UserMessage{
			UserId: userId,
		},
	}

	if messageId != 0 {
		messageQueryWrapper.Where.MessageId = messageId
	}

	if userOtherId != 0 {
		messageQueryWrapper.Where.UserOtherId = userOtherId
	}

	userMessage := &do.UserMessage{
		MessageIsRead: true,
	}

	_, err := dao.UserMessage.EditWhere(ctx, messageQueryWrapper, userMessage)
	if err != nil {
		return false, err
	}

	return true, nil
}

// addMessage 添加消息
func (s *sUserMessage) AddMessage(ctx context.Context, messageAddInput *model.UserMessageAddInput, userId uint) (*model.UserMessageVo, error) {
	// 判断接收者是否存在
	var userInfo *entity.UserInfo
	userOtherNickname := messageAddInput.UserNickname

	if userOtherNickname != "" {
		userInfo, _ = dao.UserInfo.FindOne(ctx, &do.UserInfoListInput{
			Where: do.UserInfo{
				UserNickname: userOtherNickname,
			},
		})
	} else {
		userOtherId := messageAddInput.UserOtherId
		if userOtherId == 0 {
			to := messageAddInput.To
			if to != "" {
				var toObj map[string]interface{}
				err := json.Unmarshal([]byte(to), &toObj)
				if err != nil {
					return nil, err
				}
				userOtherId = gconv.Uint(toObj["friend_id"])
				if userOtherId == 0 {
					name := gconv.String(toObj["name"])
					userInfo, _ = dao.UserInfo.FindOne(ctx, &do.UserInfoListInput{
						Where: do.UserInfo{
							UserNickname: name,
						},
					})
				}
			}
		} else {
			userInfo, _ = dao.UserInfo.Get(ctx, userOtherId)
		}
	}

	messageContent := messageAddInput.MessageContent
	if messageContent == "" && messageAddInput.Mine != "" {
		var mineObj map[string]interface{}
		err := json.Unmarshal([]byte(messageAddInput.Mine), &mineObj)
		if err != nil {
			return nil, err
		}
		messageContent = gconv.String(mineObj["content"])
	}

	userMessageRes := &model.UserMessageVo{}

	if userInfo != nil {
		// 用户信息
		user, err := dao.UserInfo.Get(ctx, userId)
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, errors.New("用户详细信息不存在")
		}

		now := time.Now().UnixMilli()
		messageCat := messageAddInput.MessageCat
		if g.IsEmpty(messageCat) {
			messageCat = "text"
		}

		// 发件箱
		send := &do.UserMessage{
			UserId:            user.UserId,
			UserNickname:      user.UserNickname,
			UserOtherId:       userInfo.UserId,
			UserOtherNickname: userInfo.UserNickname,
			MessageTitle:      messageAddInput.MessageTitle,
			MessageContent:    messageContent,
			MessageLength:     messageAddInput.MessageLength,
			MessageW:          messageAddInput.MessageW,
			MessageH:          messageAddInput.MessageH,
			MessageIsRead:     true,
			MessageIsDelete:   false,
			MessageType:       2,
			MessageKind:       1,
			MessageCat:        messageCat,
			MessageTime:       now,
		}

		// 收件箱
		inbox := &do.UserMessage{
			UserId:            userInfo.UserId,
			UserNickname:      userInfo.UserNickname,
			UserOtherId:       user.UserId,
			UserOtherNickname: user.UserNickname,
			MessageTitle:      messageAddInput.MessageTitle,
			MessageContent:    messageContent,
			MessageLength:     messageAddInput.MessageLength,
			MessageW:          messageAddInput.MessageW,
			MessageH:          messageAddInput.MessageH,
			MessageIsRead:     false,
			MessageIsDelete:   false,
			MessageType:       2,
			MessageKind:       2,
			MessageCat:        messageCat,
			MessageTime:       now,
		}

		messageId, err := dao.UserMessage.Add(ctx, send)
		if err != nil {
			return nil, errors.New("保存发件箱失败")
		}
		userMessageRes.MessageId = gconv.Uint(messageId)

		messageOtherId, err := dao.UserMessage.Add(ctx, inbox)
		if err != nil {
			return nil, errors.New("保存收件箱失败")
		}

		userMessageRes.MessageOtherId = gconv.Uint(messageOtherId)
	}

	return userMessageRes, nil
}

// ListChatMsg 读取聊天消息
func (s *sUserMessage) ListChatMsg(ctx context.Context, req *do.UserMessageListInput) (messageResPage *model.MessageOutput, err error) {
	messageResPage = &model.MessageOutput{}

	userMessagePage, err := dao.UserMessage.List(ctx, req)
	if err != nil {
		return nil, err
	}

	if userMessagePage != nil && len(userMessagePage.Items) > 0 {
		messageResPage.Records = userMessagePage.Records
		messageResPage.Size = userMessagePage.Size
		messageResPage.Total = userMessagePage.Total
		messageResPage.Page = userMessagePage.Page

		userMessageList := make([]*model.UserMessageVo, 0)
		gconv.Structs(userMessagePage.Items, &userMessageList)

		// 用户头像
		userMap, _ := s.getUserAvatarMap(ctx, userMessageList)

		// 相关用户头像
		otherUserMap, _ := s.getOtherUserAvatarMap(ctx, userMessageList)

		messageResList := make([]*model.ChatMessageVo, 0)
		for _, item := range userMessageList {
			messageRes := &model.ChatMessageVo{MsgType: "user"}
			// 站内信数据
			userMessageRes := &model.UserMessageVo{}
			userMessageRes.MessageId = item.MessageId
			userMessageRes.MessageCat = item.MessageCat

			// 站内信 userInfo
			messageUserInfoRes := &model.MessageUserInfoVo{}

			if item.MessageKind == 1 {
				userId := item.UserId
				messageUserInfoRes.Username = item.UserNickname
				messageUserInfoRes.Uid = s.getPlatformUid(ctx, userId)

				if userMap != nil && len(userMap) > 0 {
					messageUserInfoRes.Face = userMap[userId]
				}
			} else {
				otherUserId := item.UserOtherId
				messageUserInfoRes.Username = item.UserOtherNickname
				messageUserInfoRes.Uid = s.getPlatformUid(ctx, otherUserId)

				if otherUserMap != nil && len(otherUserMap) > 0 {
					messageUserInfoRes.Face = otherUserMap[otherUserId]
				}
			}

			// 站内信 content
			messageContentRes := &model.MessageContentVo{}

			switch item.MessageCat {
			case "text", "blessing", "order":
				messageContentRes.MessageContent = item.MessageContent
			case "img":
				messageContentRes.MessageContent = item.MessageContent
				messageContentRes.MessageW = item.MessageW
				messageContentRes.MessageH = item.MessageH
			case "voice", "video":
				messageContentRes.MessageContent = item.MessageContent
				messageContentRes.MessageLength = item.MessageLength
			case "item":
				messageContent := item.MessageContent
				if !g.IsEmpty(messageContent) && gstr.IsNumeric(messageContent) {
					productItem, err := dao.ProductItem.Get(ctx, gconv.Uint64(messageContent))
					if err != nil {
						return nil, err
					}
					if productItem != nil {
						messageContentRes.ItemUnitPrice = productItem.ItemUnitPrice
						messageContentRes.ProductItemName = productItem.ItemName
						messageContentRes.ItemId = productItem.ItemId

						productImage, err := dao.ProductImage.FindOne(ctx, &do.ProductImageListInput{Where: do.ProductImage{ProductId: productItem.ProductId}})
						if err != nil {
							return nil, err
						}

						if productImage != nil {
							messageContentRes.ProductImage = productImage.ItemImageDefault
						}

						messageContentRes.MessageContent = item.MessageContent
					}
				}
			default:
				messageContentRes.MessageContent = item.MessageContent
			}

			userMessageRes.Userinfo = messageUserInfoRes
			userMessageRes.Content = messageContentRes

			messageRes.Msg = userMessageRes
			messageResList = append(messageResList, messageRes)
		}

		messageResPage.Items = messageResList
	}

	return messageResPage, nil
}

// SendSysNotice 发送系统通知
func (s *sUserMessage) SendSysNotice(ctx context.Context, adminUserId, userId int, msgContent string, msgType int) error {
	// 收件箱
	other := &do.UserMessage{
		UserId:          userId,
		UserOtherId:     adminUserId,
		MessageContent:  msgContent,
		MessageIsRead:   false,
		MessageIsDelete: false,
		MessageType:     msgType,
		MessageKind:     2,
		MessageTime:     gtime.Now().UnixMilli(),
	}

	if _, err := dao.UserMessage.Add(ctx, other); err != nil {
		return gerror.New("消息发送失败")
	}

	return nil
}

func (s *sUserMessage) getPlatformUid(ctx context.Context, userId uint) (uid int) {
	configBase, err := dao.ConfigBase.Ctx(ctx).Where(do.ConfigBase{
		ConfigKey: "service_user_id",
	}).One()
	if err != nil || configBase.IsEmpty() {
		return 0
	}
	serviceUserId := service.ConfigBase().GetUint(ctx, "service_user_id", 0)
	puid := fmt.Sprintf("%s-%d", serviceUserId, userId)

	return utility.BkdrHash(puid)
}

// 获取通知消息
func (s *sUserMessage) GetPlantFromNotice(ctx context.Context, userId uint) (*model.MessageNoticeVo, error) {
	noticeRes := &model.MessageNoticeVo{}

	// 构建查询条件
	messageQueryWrapper := &do.UserMessageListInput{
		Where: do.UserMessage{
			UserId:        userId,
			MessageIsRead: 0,
			MessageType:   1,
		},
	}
	messageQueryWrapper.Sidx = dao.UserMessage.Columns().MessageTime
	messageQueryWrapper.Sort = ml.ORDER_BY_DESC

	time, _ := mtime.LastNDays(10)
	var likes = []*ml.WhereExt{{
		Column: dao.UserMessage.Columns().MessageTime,
		Val:    time,
		Symbol: ml.GE,
	}}
	messageQueryWrapper.WhereExt = likes

	// 查询用户消息
	userMessages, err := dao.UserMessage.Find(ctx, messageQueryWrapper)
	if err != nil {
		return nil, err
	}

	userMessageResList := make([]*model.UserMessageVo, len(userMessages))
	gconv.Structs(userMessages, &userMessageResList)

	// 获取用户ID列表
	userIds := array.Column(userMessageResList, "UserId")
	userInfos, err := dao.UserInfo.Gets(ctx, userIds)
	if err != nil {
		return nil, err
	}

	for _, messageRes := range userMessageResList {
		for _, userInfo := range userInfos {
			if messageRes.UserId == userInfo.UserId {
				messageRes.UserAvatar = userInfo.UserAvatar
			}
		}
	}

	noticeRes.Items = userMessageResList
	noticeRes.Total = len(userMessages)

	return noticeRes, nil
}

// 清除通知消息
func (s *sUserMessage) ClearNotice(ctx context.Context, userId uint) (bool, error) {
	// 构建查询条件
	messageQueryWrapper := &do.UserMessageListInput{
		Where: do.UserMessage{
			UserId:        userId,
			MessageIsRead: 0,
			MessageType:   1,
		},
	}

	userMessage := &do.UserMessage{
		MessageIsRead: true,
	}

	if _, err := dao.UserMessage.EditWhere(ctx, messageQueryWrapper, userMessage); err != nil {
		return false, gerror.New("消息清空失败！")
	}

	return true, nil
}

// 获取消息数量
func (s *sUserMessage) GetMessageNum(ctx context.Context, input *do.UserMessageListInput) (*model.UserMessageVo, error) {
	userMessageRes := &model.UserMessageVo{}

	/*
		// 构建查询条件
		messageQueryWrapper := &do.UserMessageListInput{
			Where: do.UserMessage{
				UserId:      req.UserId,
				MessageKind: 2,
			},
		}

		// 查询未读消息数
		queryWrapperOne := messageQueryWrapper.Clone()
		queryWrapperOne.Where.MessageIsRead = 0
	*/
	unread, err := dao.UserMessage.Count(ctx, input)
	if err != nil {
		return nil, err
	}
	userMessageRes.UnreadNumber = unread

	// 查询已读消息数
	input.Where.MessageIsRead = 1
	read, err := dao.UserMessage.Count(ctx, input)
	if err != nil {
		return nil, err
	}

	userMessageRes.RedNumber = read

	return userMessageRes, nil
}
