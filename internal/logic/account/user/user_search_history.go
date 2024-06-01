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
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"math/rand"
	"time"
)

type sUserSearchHistory struct{}

func init() {
	service.RegisterUserSearchHistory(NewUserSearchHistory())
}

func NewUserSearchHistory() *sUserSearchHistory {
	return &sUserSearchHistory{}
}

// Find 查询数据
func (s *sUserSearchHistory) Find(ctx context.Context, in *do.UserSearchHistoryListInput) (out []*entity.UserSearchHistory, err error) {
	out, err = dao.UserSearchHistory.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserSearchHistory) List(ctx context.Context, in *do.UserSearchHistoryListInput) (out *do.UserSearchHistoryListOutput, err error) {
	out, err = dao.UserSearchHistory.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserSearchHistory) Add(ctx context.Context, in *do.UserSearchHistory) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserSearchHistory.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserSearchHistory) Edit(ctx context.Context, in *do.UserSearchHistory) (affected int64, err error) {
	_, err = dao.UserSearchHistory.Edit(ctx, in.SearchId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserSearchHistory) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.UserSearchHistory.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetSearchInfo 返回搜索关键词
func (s *sUserSearchHistory) GetSearchInfo(ctx context.Context) (*pt.SearchInfoRes, error) {
	infoRes := &pt.SearchInfoRes{}

	suggestSearchWords := service.ConfigBase().GetStr(ctx, "suggest_search_words", "")
	searchHotWords := service.ConfigBase().GetStr(ctx, "search_hot_words", "")

	if suggestSearchWords != "" {
		wordsArray := gstr.Split(suggestSearchWords, ",")
		if len(wordsArray) > 0 {
			// 设置种子，以确保每次运行都能生成不同的随机数序列
			rand.Seed(time.Now().UnixNano())

			// 生成一个介于 0 和 len(wordsArray)-1 之间的随机索引
			randomIndex := rand.Intn(len(wordsArray))

			wordStr := wordsArray[randomIndex]
			searchWordVo := model.SearchWordVo{
				DefaultSearchLabel: wordStr,
				DefaultSearchWords: wordStr,
			}
			infoRes.SuggestSearchWords = searchWordVo
		}

		infoRes.SearchHotWords = gstr.Split(searchHotWords, ",")
	}

	infoRes.SearchHistoryWords = []string{}
	user := service.BizCtx().GetUser(ctx)

	if user != nil {
		// 查询用户搜索记录
		searchHistories, _ := dao.UserSearchHistory.Find(ctx, &do.UserSearchHistoryListInput{
			Where: do.UserSearchHistory{UserId: user.UserId},
			BaseList: ml.BaseList{
				Sidx: dao.UserSearchHistory.Columns().SearchTime,
				Sort: ml.ORDER_BY_DESC,
			},
		})

		if len(searchHistories) > 0 {
			keyWords := make([]string, len(searchHistories))
			for i, history := range searchHistories {
				keyWords[i] = history.SearchKeyword
			}
			infoRes.SearchHistoryWords = keyWords
		} else {
			infoRes.SearchHistoryWords = []string{}
		}
	}

	return infoRes, nil
}
