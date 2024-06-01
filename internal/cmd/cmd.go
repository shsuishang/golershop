package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gmode"
	"golershop.cn/internal/controller/account"
	"golershop.cn/internal/controller/admin"
	"golershop.cn/internal/controller/analytics"
	"golershop.cn/internal/controller/cms"
	"golershop.cn/internal/controller/marketing"
	"golershop.cn/internal/controller/pay"
	"golershop.cn/internal/controller/pt"
	"golershop.cn/internal/controller/shop"
	"golershop.cn/internal/controller/sys"
	"golershop.cn/internal/controller/trade"
	"golershop.cn/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start shopsuite http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// HOOK, 开发阶段禁止浏览器缓存,方便调试
			if gmode.IsDevelop() {
				s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					r.Response.Header().Set("Cache-Control", "no-store")
				})
			}

			// 跨域处理
			s.Use(service.Middleware().CORS)

			//需要传播给异步流程或者保持和之前逻辑兼容
			s.Use(service.Middleware().NeverDoneCtx)

			s.Use(service.Middleware().Ctx)

			s.Use(service.Middleware().MiddlewareErrorHandler)
			s.Use(service.Middleware().MiddlewareHandlerResponse)

			s.Use(service.Middleware().CheckLogin)

			s.Group("/", func(group *ghttp.RouterGroup) {
				//钩子记录异步日志
				group.Hook("/manage/*", ghttp.HookAfterOutput, service.LogAction().OperateLog)
				group.Hook("/front/*", ghttp.HookAfterOutput, service.AccessHistory().OperateAccess)

				group.Hook("/*", ghttp.HookAfterOutput, service.BizCtx().AfterOutput)

				/*
					// Group middlewares.
					group.Middleware(
						service.Middleware().Ctx,
						service.Middleware().CheckLogin,
					)

				*/

				group.Bind(
					admin.Menu,
					admin.UserRole,
					account.User,
					account.UserInvoice,
					account.UserDeliveryAddress,
					account.UserLevel,
					account.UserInfo,
					account.UserMessage,
					account.UserTagGroup,
					cms.ArticleBase,
					cms.ArticleTag,
					cms.ArticleCategory,
					cms.ArticleComment,
					cms.Article,
					account.Login,
					marketing.ActivityBase,
					pay.Resource,
					pay.UserResource,
					pay.ConsumeTrade,
					pay.ConsumeDeposit,
					pay.ConsumeRecord,
					pay.UserPointsHistory,
					pay.PaymentCallback,
					pay.PaymentIndex,
					pay.UserPay,
					pay.Record,
					pt.Product,
					pt.ProductCategory,
					pt.ProductTag,
					pt.ProductBrand,
					pt.ProductItem,
					pt.ProductType,
					pt.ProductBase,
					pt.ProductSpec,
					pt.ProductSpecItem,
					pt.ProductAssist,
					shop.Shop,
					shop.FavoritesItem,
					shop.StoreExpressLogistics,
					shop.StoreShippingAddress,
					shop.StoreTransportType,
					shop.StoreTransportItem,
					sys.Config,
					sys.Dict,
					sys.ExpressBase,
					sys.ContractType,
					sys.CrontabBase,
					sys.FeedbackBase,
					sys.Feedback,
					sys.FeedbackCategory,
					sys.FeedbackType,
					sys.LogAction,
					sys.Material,
					sys.MessageTemplate,
					sys.Captcha,
					sys.Upload,
					sys.Page,
					trade.Order,
					trade.Cart,
					trade.OrderBase,
					trade.OrderInvoice,
					trade.OrderLogistics,
					trade.OrderReturn,
					trade.Return,
					analytics.Analytics,
					analytics.AnalyticsReturn,
				)
			})

			//config init
			service.ConfigBase().Init(ctx)

			s.Run()
			return nil
		},
	}
)
