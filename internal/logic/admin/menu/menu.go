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

package menu

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"time"
)

type sMenu struct{}

var (
	cachePreKey   = "menu_base"
	cacheDuration = time.Hour
	redisAdapter  = gcache.NewAdapterRedis(g.Redis())
	cache         = gcache.NewWithAdapter(redisAdapter)
)

func init() {
	service.RegisterMenu(New())
}

func New() *sMenu {
	return &sMenu{}
}

// 查看菜单
func (s *sMenu) GetMenu(ctx context.Context, id uint) (out *entity.MenuBase, err error) {
	//redisCache := gcache.NewAdapterRedis(g.Redis())
	//cache := gcache.NewWithAdapter(redisCache)

	// Create redis cache adapter and set it to cache object.
	//cache.SetAdapter(redisCache)

	//g.DB().GetCache().SetAdapter(redisCache)
	/*
		var (
			cacheKey   = "menu_base|" + "1001"
			cacheValue = `value`
		)

		// Set and Get using cache object.
		err = cache.Set(ctx, cacheKey, cacheValue, time.Hour)
		if err != nil {
			panic(err)
		}
		fmt.Println(cache.MustGet(ctx, cacheKey).String())
	*/

	// 缓存控制
	var (
		cacheKey  = "menu_base|" + "1001"
		cacheFunc = func(ctx context.Context) (interface{}, error) {
			var menu *entity.MenuBase
			var err error
			err = dao.MenuBase.Ctx(ctx).Where(do.MenuBase{
				MenuId: id,
			}).Scan(&menu)

			return menu, err
		}
	)

	v, err := cache.GetOrSetFuncLock(ctx, cacheKey, cacheFunc, time.Hour)
	if err != nil {
		return nil, err
	}

	err = v.Scan(&out)

	/*
		err = dao.MenuBase.Ctx(ctx).Where(do.MenuBase{
			MenuId: id,
		}).Scan(&out)

	*/

	if err != nil {
		return nil, err
	}

	return out, nil
}

// 读取菜单
func (s *sMenu) Get(ctx context.Context, menuId any) (out *entity.MenuBase, err error) {
	/*
		err = dao.MenuBase.Ctx(ctx).WherePri(menuId).Scan(&out)

		if err != nil {
			return nil, err
		}

		return out, nil
	*/

	var list []*entity.MenuBase
	list, err = s.Gets(ctx, menuId)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// 读取多条记录模式
func (s *sMenu) Gets(ctx context.Context, menuId any) (list []*entity.MenuBase, err error) {

	err = dao.MenuBase.Ctx(ctx).WherePri(menuId).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// 查询主键
func (s *sMenu) FindKey(ctx context.Context, in *do.MenuBaseListInput) (out *model.MenuKeysOutput, err error) {
	var (
		m = dao.MenuBase.Ctx(ctx)
	)

	m = m.Fields(dao.MenuBase.Columns().MenuId)

	query := m.Where(in.Where)

	// 排序
	//query = query.Order("sort asc")
	//query = query.OrderDesc(in.Sidx)

	// 对象转换
	var ids []struct {
		MenuId int
	}

	if err := query.Scan(&ids); err != nil {
		return out, err
	}

	fmt.Println("====================\n")
	fmt.Println(ids)
	fmt.Println("====================\n")

	return out, nil
}

// 查询数据
func (s *sMenu) Find(ctx context.Context, in *do.MenuBaseListInput) (out []*entity.MenuBase, err error) {
	var (
		m = dao.MenuBase.Ctx(ctx)
	)

	query := m.Where(in.Where)

	// 排序
	//query = query.Order("sort asc")
	//query = query.OrderDesc(in.Sidx)

	// 对象转换
	//var list []*entity.MenuBase
	if err := query.Scan(&out); err != nil {
		return out, err
	}

	return out, nil
}

// 分页读取
func (s *sMenu) List(ctx context.Context, in *do.MenuBaseListInput) (out *do.MenuBaseListOutput, err error) {
	list, err := dao.MenuBase.List(ctx, in)

	gconv.Scan(list, &out)

	return out, nil
}

// 业务封装

// 查询数据
func (s *sMenu) GetTree(ctx context.Context, in *do.MenuBaseListInput) (out []*model.TreeNode, err error) {
	in.Where.MenuRole = 1
	in.Where.MenuEnable = 1
	in.Sidx = dao.MenuBase.Columns().MenuSort
	in.Sort = "ASC"

	res, err := dao.MenuBase.List(ctx, in)

	// 数据转换
	var list []*entity.MenuBase
	gconv.Scan(res.Items, &list)

	// begin 为了菜单显示隐藏判断
	var menuIds []string
	loginUser := service.BizCtx().GetUser(ctx)
	userAdmin, _ := dao.UserAdmin.Get(ctx, loginUser.UserId)

	if !g.IsEmpty(userAdmin) {
		userRole, _ := dao.UserRole.Get(ctx, userAdmin.UserRoleId)
		menuIds = gstr.Split(userRole.MenuIds, ",")
	}

	all, err := dao.MenuBase.Find(ctx, &do.MenuBaseListInput{})
	// end 为了菜单显示隐藏判断

	// 数据处理
	var menuNode model.TreeNode
	makeTree(list, &menuNode, menuIds, all)

	//或者无上级数据，加入列表 -- 用户树形搜索展示
	columnIds := array.Column(list, dao.MenuBase.Columns().MenuId)

	for _, c := range list {
		if c.MenuParentId != 0 && !array.InArray(columnIds, c.MenuParentId) {
			child := &model.TreeNode{}
			gconv.Scan(*c, &child.Menu)
			gconv.Scan(*c, &child.Menu.Meta)
			child.Menu.Meta.MenuClose = !c.MenuClose

			menuNode.Children = append(menuNode.Children, child)
		}
	}

	return menuNode.Children, nil
}

// 递归生成分类列表
func makeTree(menu []*entity.MenuBase, tn *model.TreeNode, menuIds []string, all []*entity.MenuBase) {
	for _, c := range menu {
		// 判断权限隐藏
		c.MenuHidden = !checkChildShow(all, c, menuIds)

		if c.MenuParentId == tn.MenuId {
			child := &model.TreeNode{}
			//child.Menu = *c
			gconv.Scan(*c, &child.Menu)
			gconv.Scan(*c, &child.Menu.Meta)
			child.Menu.Meta.MenuClose = !c.MenuClose

			/*
				child.Menu.MenuId = c.MenuId
				child.Menu.MenuName = c.MenuName
				child.Menu.MenuPath = c.MenuPath
				child.Menu.MenuComponent = c.MenuComponent
				child.Menu.MenuRedirect = c.MenuRedirect

				child.Menu.Meta.MenuTitle = c.MenuTitle
				child.Menu.Meta.MenuHidden = c.MenuHidden
				child.Menu.Meta.MenuIcon = c.MenuIcon
				child.Menu.Meta.MenuBubble = c.MenuBubble
				child.Menu.Meta.MenuClose = !c.MenuClose

			*/

			tn.Children = append(tn.Children, child)
			makeTree(menu, child, menuIds, all)
		}
	}
}

// 递归判断是否存在有权限的子菜单
func checkChildShow(menu []*entity.MenuBase, tn *entity.MenuBase, menuIds []string) bool {
	flag := false

	//默认设置隐藏
	if tn.MenuHidden {
		flag = false
	} else {
		if gstr.InArray(menuIds, gconv.String(tn.MenuId)) {
			flag = true
		} else {
			//当前菜单无权限，判断所有子节点都不存在权限中，则隐藏菜单, 使用递归判断实现
			for _, c := range menu {
				if c.MenuParentId == tn.MenuId {
					flag = checkChildShow(menu, c, menuIds)

					if flag {
						break
					}
				}
			}
		}
	}

	return flag
}

// 新增
func (s *sMenu) Add(ctx context.Context, in *do.MenuBase) (out int64, err error) {
	// 不允许HTML代码
	//if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
	//	return out, err
	//}

	lastInsertId, err := dao.MenuBase.Add(ctx, in)
	if err != nil {
		return out, err
	}
	return lastInsertId, err
}

// 编辑
func (s *sMenu) Edit(ctx context.Context, in *do.MenuBase) (affected int64, err error) {
	// 不允许HTML代码
	//if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
	//	return out, err
	//}

	affected, err = dao.MenuBase.Edit(ctx, in.MenuId, in)

	return affected, err
}

// 删除多条记录模式
func (s *sMenu) Remove(ctx context.Context, menuId any) (affected int64, err error) {
	//是否内置
	one, err := dao.MenuBase.Get(ctx, menuId)
	if one.MenuBuildin {
		return 0, errors.New("系统内置，不可删除")
	}

	count, err := dao.MenuBase.Ctx(ctx).Count(do.MenuBase{MenuParentId: menuId})

	if err != nil {
		return 0, err
	}

	if count > 0 {
		return 0, errors.New(fmt.Sprintf("有 %d 个子级菜单，不可删除", count))
	}

	affected, err = dao.MenuBase.Remove(ctx, menuId)

	if err != nil {
		return 0, err
	}

	return affected, err
}
