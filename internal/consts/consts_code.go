package consts

const (
	DELIVERY_TYPE_EXPRESS      uint = 1  //快递配送（运费 10 元）
	DELIVERY_TYPE_EMS          uint = 2  //EMS（邮政）
	DELIVERY_TYPE_MAIL         uint = 3  //平邮
	DELIVERY_TYPE_AIR_FREIGHT  uint = 4  //货运（空运、水运、铁路运输、公路运输）
	DELIVERY_TYPE_SELF_PICK_UP uint = 5  //自提（运费 0 元）
	DELIVERY_TYPE_EXP          uint = 10 //配送

	DELIVERY_TIME_NO_LIMIT    uint = 1 //不限送货时间：周一至周日
	DELIVERY_TIME_WORKING_DAY uint = 2 //工作日送货：周一至周五
	DELIVERY_TIME_WEEKEND     uint = 3 //双休日、假日送货：周六至周日

	USER_STATE_LOCKING    uint = 0 //用户状态:锁定
	USER_STATE_NOTACTIVE  uint = 1 //用户状态:未激活
	USER_STATE_ACTIVATION uint = 2 //用户状态:已激活

	PRODUCT_STATE_ILLEGAL       uint = 1000 //违规下架禁售
	PRODUCT_STATE_NORMAL        uint = 1001 //正常
	PRODUCT_STATE_OFF_THE_SHELF uint = 1002 //下架

	DEMAND_STATE_CONDUCT uint = 1000 //采购中
	DEMAND_STATE_REJECT  uint = 1030 //被驳回
	DEMAND_STATE_EXAMINE uint = 1040 //审核中

	//商品标签
	PRODUCT_TAG_NEW        uint = 1401 //新品上架
	PRODUCT_TAG_REC        uint = 1402 //热卖推荐
	PRODUCT_TAG_BARGAIN    uint = 1403 //清仓优惠
	PRODUCT_TAG_BARGAIN1   uint = 1404 //清仓优惠
	PRODUCT_TAG_CROSSBORDS uint = 1405 //清仓优惠

	//商品种类
	PRODUCT_KIND_ENTITY uint = 1201 //实体商品	实物商品 （网购物流发货）
	PRODUCT_KIND_FUWU   uint = 1202 //预购订单	服务类   （无需物流-门店或者上门服务，核销）
	PRODUCT_KIND_CARD   uint = 1203 //电子卡券	电子卡券 （无需物流，发送代码即代表订单完成-目前同步处理，未来可以异步处理）
	PRODUCT_KIND_WAIMAI uint = 1204 //外卖订单	外卖订单 （同城配送），类似PRODUCT_KIND_ENTITY， 根据类型增加配送时间选项。可以由O2O店铺类型决定是否需要配送时间等等
	PRODUCT_KIND_EDU    uint = 1205 //教育课程类订单

	PRODUCT_VERIFY_REFUSED uint = 3000 //审核未通过
	PRODUCT_VERIFY_PASSED  uint = 3001 //审核通过
	PRODUCT_VERIFY_WAITING uint = 3002 //审核中

	ORDER_STATE_WAIT_PAY            uint = 2010 //待付款 - 虚拟映射
	ORDER_STATE_WAIT_PAID           uint = 2016 //已经付款 - 虚拟映射
	ORDER_STATE_WAIT_REVIEW         uint = 2011 //待订单审核
	ORDER_STATE_WAIT_FINANCE_REVIEW uint = 2013 //待财务审核
	ORDER_STATE_PICKING             uint = 2020 //待配货
	ORDER_STATE_WAIT_SHIPPING       uint = 2030 //待发货
	ORDER_STATE_SHIPPED             uint = 2040 //已发货
	ORDER_STATE_RECEIVED            uint = 2050 //已签收
	ORDER_STATE_FINISH              uint = 2060 //已完成
	ORDER_STATE_CANCEL              uint = 2070 //已取消
	ORDER_STATE_SELF_PICKUP         uint = 2080 //自提 - 虚拟映射     交易关闭	         交易关闭

	ORDER_STATE_ERROR  uint = 2090 //异常订单
	ORDER_STATE_RETURN uint = 2091 //退回订单 - 虚拟映射

	//骑手端
	ORDER_STATE_PICKUP         uint = 2045 //骑手取货
	ORDER_STATE_RIDER_RECEIVED uint = 2046 //骑手取货

	ORDER_PAID_STATE_NO             uint = 3010 //未付款
	ORDER_PAID_STATE_FINANCE_REVIEW uint = 3011 //待付款审核
	ORDER_PAID_STATE_PART           uint = 3012 //部分付款
	ORDER_PAID_STATE_YES            uint = 3013 //已付款

	ORDER_PICKING_STATE_NO   uint = 3020 //未出库
	ORDER_PICKING_STATE_PART uint = 3021 //部分出库通过拆单解决这种问题
	ORDER_PICKING_STATE_YES  uint = 3022 //已出库

	ORDER_CARDKIND_STATE_CARD    uint = 1001 //次卡类订单
	ORDER_CARDKIND_STATE_VOUCHER uint = 1002 //优惠券类订单
	ORDER_CARDKIND_STATE_COUPON  uint = 1003 //券码类订单

	ORDER_SHIPPED_STATE_NO   uint = 3030 //未发货
	ORDER_SHIPPED_STATE_PART uint = 3031 //部分发货
	ORDER_SHIPPED_STATE_YES  uint = 3032 //已发货

	VIRTUAL_ORDER_USED    uint = 2101 //虚拟订单已使用
	VIRTUAL_ORDER_UNUSE   uint = 2100 //虚拟订单未使用
	VIRTUAL_ORDER_TIMEOUT uint = 2103 //虚拟订单过期

	ORDER_CANCEL_BY_BUYER  uint = 2201 //买家取消订单
	ORDER_CANCEL_BY_SELLER uint = 2202 //卖家取消订单
	ORDER_CANCEL_BY_ADMIN  uint = 2203 //平台取消

	SOURCE_TYPE_OTHER uint = 2310 //来源于其它
	SOURCE_TYPE_PC    uint = 2311 //来源于pc端
	SOURCE_TYPE_H5    uint = 2312 //来源于H5端
	SOURCE_TYPE_APP   uint = 2313 //来源于APP
	SOURCE_TYPE_MP    uint = 2314 //来源于小程序

	SOURCE_FROM_OTHER   uint = 2320 //来源于其它
	SOURCE_FROM_WECHAT  uint = 2321 //来源于微信平台，包含公众号，小程序等等
	SOURCE_FROM_BAIDU   uint = 2322 //来源于百度
	SOURCE_FROM_ALIPAY  uint = 2323 //来源于支付宝
	SOURCE_FROM_TOUTIAO uint = 2324 //来源于头条

	//订单来源
	ORDER_FROM_PC     uint = 2301 //来源于pc端
	ORDER_FROM_WAP    uint = 2302 //来源于WAP手机端
	ORDER_FROM_WEBPOS uint = 2303 //来源于WEBPOS线下下单

	//状态
	SETTLEMENT_STATE_WAIT_OPERATE       uint = 2401 //已出账
	SETTLEMENT_STATE_SELLER_COMFIRMED   uint = 2402 //商家已确认
	SETTLEMENT_STATE_PLATFORM_COMFIRMED uint = 2403 //平台已审核
	SETTLEMENT_STATE_FINISH             uint = 2404 //结算完成

	ORDER_RETURN_NO  uint = 2500 //无退货
	ORDER_RETURN_ING uint = 2501 //退货中
	ORDER_RETURN_END uint = 2502 //退货完成

	ORDER_REFUND_STATE_NO  uint = 2600 //无退款
	ORDER_REFUND_STATE_ING uint = 2601 //退款中
	ORDER_REFUND_STATE_END uint = 2602 //退款完成

	ORDER_TYPE_DD uint = 3061 //订单类型
	ORDER_TYPE_DC uint = 3063 //线下收银
	ORDER_TYPE_FX uint = 3062 //分销订单
	ORDER_TYPE_TH uint = 3066 //分销订单
	ORDER_TYPE_MD uint = 3068 //转单大厅订单
	ORDER_TYPE_PT uint = 3069 //跑腿订单

	ORDER_TYPE_XQ uint = 4034 //需求订单
	ORDER_TYPE_FW uint = 4035 //服务订单
	ORDER_TYPE_XX uint = 5000 //线下记账

	ACTIVITY_STATE_WAITING  uint = 0 //活动状态:0-未开启
	ACTIVITY_STATE_NORMAL   uint = 1 //活动状态:1-正常
	ACTIVITY_STATE_FINISHED uint = 2 //活动状态:2-已结束
	ACTIVITY_STATE_CLOSED   uint = 3 //活动状态:3-管理员关闭

	GET_VOUCHER_FREE        uint = 1 //活动状态:1-免费参与
	GET_VOUCHER_BY_POINT    uint = 2 //活动状态:2-积分参与
	GET_VOUCHER_BY_PURCHASE uint = 3 //活动状态:3-购买参与
	GET_VOUCHER_BY_SHARE    uint = 4 //活动状态:4-分享参与

	ACTIVITY_GROUP_BOOKING_FAIL     uint = 0 //拼团状态:0-失败
	ACTIVITY_GROUP_BOOKING_SUCCESS  uint = 1 //拼团状态:1-成功
	ACTIVITY_GROUP_BOOKING_UNDERWAY uint = 2 //拼团状态:2-进行中

	CART_GET_TYPE_BUY     uint = 1 //购买
	CART_GET_TYPE_POINT   uint = 2 //积分兑换
	CART_GET_TYPE_GIFT    uint = 3 //赠品
	CART_GET_TYPE_BARGAIN uint = 4 //活动促销

	/*
	   BILL_TYPE_PO   uint  = 4001   //购货单
	   BILL_TYPE_PORO uint  = 4002   //销货退货单
	   BILL_TYPE_OI   uint  = 4003   //其他入库单
	   BILL_TYPE_SO   uint  = 4031   //销货单
	   BILL_TYPE_SORO uint  = 4032   //购货退货单
	   BILL_TYPE_OO   uint  = 4033   //其他出库单
	*/

	STOCK_IN_PURCHASE    uint = 2701 //采购入库
	STOCK_IN_RETURN      uint = 2702 //退货入库
	STOCK_IN_ALLOCATE    uint = 2703 //调库入库
	STOCK_IN_INVENTORY_P uint = 2704 //盘盈入库
	STOCK_IN_INIT        uint = 2705 //期初入库
	STOCK_IN_OTHER       uint = 2706 //手工入库
	STOCK_OUT_SALE       uint = 2751 //销售出库
	STOCK_OUT_DAMAGED    uint = 2752 //损坏出库
	STOCK_OUT_ALLOCATE   uint = 2753 //调库出库
	STOCK_OUT_LOSSES     uint = 2754 //盘亏出库
	STOCK_OUT_OTHER      uint = 2755 //手工出库
	STOCK_OUT_PO_RETURN  uint = 2756 //损坏出库

	STOCK_OUT_ALL uint = 2700 //出库单
	STOCK_IN_ALL  uint = 2750 //入库单

	BILL_TYPE_OUT uint = 2700 //出库单
	BILL_TYPE_IN  uint = 2750 //入库单

	BILL_TYPE_SO uint = 2800 //销售订单
	BILL_TYPE_PO uint = 2850 //采购订单

	//修改掉，和订单状态对应。
	ORDER_PROCESS_SUBMIT uint = 3070 //【客户】提交订单1OrderOrder

	ORDER_PROCESS_PAY            uint = 2010 //待支付Order
	ORDER_PROCESS_CHECK          uint = 2011 //订单审核1OrderOrder
	ORDER_PROCESS_FINANCE_REVIEW uint = 2013 //财务审核0OrderOrder
	ORDER_PROCESS_OUT            uint = 2020 //出库审核商品库存在“出库审核”节点完成后扣减，如需进行库存管理或核算销售成本毛利，需开启此节点。0OrderOrder
	ORDER_PROCESS_SHIPPED        uint = 2030 //发货确认如需跟踪订单物流信息，需开启此节点0OrderOrder
	ORDER_PROCESS_RECEIVED       uint = 2040 //【客户】收货确认0OrderOrder

	ORDER_PROCESS_FINISH uint = 3098 //完成1OrderOrder

	RETURN_PROCESS_SUBMIT               uint = 3100 //【客户】提交退单1ReturnReturn
	RETURN_PROCESS_CHECK                uint = 3105 //退单审核1ReturnReturn
	RETURN_PROCESS_RECEIVED             uint = 3110 //收货确认0ReturnReturn
	RETURN_PROCESS_REFUND               uint = 3115 //退款确认0ReturnReturn
	RETURN_PROCESS_RECEIPT_CONFIRMATION uint = 3120 //客户】收款确认0ReturnReturn
	RETURN_PROCESS_FINISH               uint = 3125 //完成1ReturnReturn3130-商家拒绝退货
	RETURN_PROCESS_REFUSED              uint = 3130 //-商家拒绝退货
	RETURN_PROCESS_CANCEL               uint = 3135 //-买家取消

	PLANTFORM_RETURN_STATE_WAITING  uint = 3180 //申请状态平台(ENUM):3180-处理中
	PLANTFORM_RETURN_STATE_AGREE    uint = 3181 //为待管理员处理卖家同意或者收货后
	PLANTFORM_RETURN_PROCESS_FINISH uint = 3182 //-为已完成

	STORE_STATE_WAIT_PROFILE uint = 3210 //待完善资料
	STORE_STATE_WAIT_VERIFY  uint = 3220 //等待审核
	STORE_STATE_NO           uint = 3230 //审核资料没有通过
	STORE_STATE_YES          uint = 3240 //审核资料通过,待付款

	TRADE_TYPE_SHOPPING            uint = 1201 //购物
	TRADE_TYPE_TRANSFER            uint = 1202 //转账
	TRADE_TYPE_DEPOSIT             uint = 1203 //充值
	TRADE_TYPE_WITHDRAW            uint = 1204 //提现
	TRADE_TYPE_SALES               uint = 1205 //销售
	TRADE_TYPE_COMMISSION          uint = 1206 //佣金
	TRADE_TYPE_REFUND_PAY          uint = 1207 //退货付款
	TRADE_TYPE_REFUND_GATHERING    uint = 1208 //退货收款
	TRADE_TYPE_TRANSFER_GATHERING  uint = 1209 //转账收款
	TRADE_TYPE_COMMISSION_TRANSFER uint = 1210 //佣金付款
	TRADE_TYPE_BONUS               uint = 1211 //分红
	TRADE_TYPE_BUY_SP              uint = 1212 //购买SP
	TRADE_TYPE_SALE_SP             uint = 1213 //销售SP
	TRADE_TYPE_FAVORABLE           uint = 1214 //线下买单
	TRADE_TYPE_OTHER               uint = 1215 //线下买单
	TRADE_TYPE_BUY_SELLER          uint = 1216 //升级为商家
	TRADE_TYPE_SALE_SELLER         uint = 1217 //销售升级为商家
	TRADE_TYPE_WITHDRAW_CANCEL     uint = 1218 //提现驳回
	TRADE_TYPE_RETURN_GROUPBOOKING uint = 1219 //拼团失败退款
	TRADE_TYPE_TRANSFER_COMMISSION uint = 1220 //转单分佣
	TRADE_TYPE_PAOTUI_FEE          uint = 1221 //跑腿运费
	TRADE_TYPE_REBATE              uint = 1222 //购物返利

	TRADE_TYPE_HALL_FEE uint = 1221 //需求订单

	TRADE_TYPE_BUY_POINTS  uint = 1223 //购买积分
	TRADE_TYPE_SALE_POINTS uint = 1224 //销售积分
	TRADE_TYPE_STORE_RENEW uint = 1225 //店铺续费

	TRADE_TYPE_BUY_POSTER    uint = 1226 //购买广告
	TRADE_TYPE_SHOPPING_CARD uint = 1227 //充值卡消费
	TRADE_TYPE_DEPOSIT_CARD  uint = 1228 //充值卡充值

	TRADE_TYPE_OFFLINE_INCREASE uint = 1501 //线下记账收入
	TRADE_TYPE_OFFLINE_DECREASE uint = 1502 //线下记账支出

	TRADE_TYPE_DOCTOR_BUY     uint = 1503 //购物
	TRADE_TYPE_DOCTOR_SERVICE uint = 1504 //销售

	TRADE_TYPE_XQ_BUY     uint = 1505 //购物
	TRADE_TYPE_XQ_SERVICE uint = 1506 //销售

	PAYMENT_TYPE_DELIVER uint = 1301 //货到付款
	PAYMENT_TYPE_ONLINE  uint = 1302 //在线支付
	//PAYMENT_TYPE_CREDIT  uint  = 1303//白条支付
	//PAYMENT_TYPE_CASH    uint  = 1304//现金支付
	PAYMENT_TYPE_OFFLINE uint = 1305 //线下支付

	ORDER_ITEM_EVALUATION_NO      uint = 0 //未评价
	ORDER_ITEM_EVALUATION_YES     uint = 1 //已评价
	ORDER_ITEM_EVALUATION_TIMEOUT uint = 2 //失效评价

	ORDER_PICKUP_CODE_UNUSED  uint = 0 //未评价
	ORDER_PICKUP_CODE_USED    uint = 1 //已评价
	ORDER_PICKUP_CODE_TIMEOUT uint = 2 //失效评价

	ORDER_EVALUATION_NO      uint = 0 //未评价
	ORDER_EVALUATION_YES     uint = 1 //已评价
	ORDER_EVALUATION_TIMEOUT uint = 2 //失效评价

	ORDER_NOT_NEED_RETURN_GOODS uint = 0 //不用退货
	ORDER_NEED_RETURN_GOODS     uint = 1 //需要退货

	ORDER_REFUND         uint = 1 //1-退款申请 2-退货申请 3-虚拟退款
	ORDER_RETURN         uint = 2 //需要退货
	ORDER_VIRTUAL_REFUND uint = 3 //需要退货

	USER_CERTIFICATION_NO     uint = 0 //0-未实名认证 1-已实名认证 2-提交资料袋审核
	USER_CERTIFICATION_VERIFY uint = 1 // 待审核
	USER_CERTIFICATION_YES    uint = 2 // 认证通过
	USER_CERTIFICATION_FAILED uint = 3 //认证失败

	TO_STORE_SERVICE     uint = 1001 //到店取货
	DOOR_TO_DOOR_SERVICE uint = 1002 // 上门服务

	ORDER_AUTO_TRANSFER_HALL_TIME float64 = 0.5 //用户付款后多长时间没被接单就转到抢单大厅
	ORDER_AUTO_REFUND_TIME        float64 = 1   //用户付款后多长时间没被接单就自动退款

	SUPPLY_TASK_STATE_BIDDING    uint = 2000 //竞标中
	SUPPLY_TASK_STATE_OVER       uint = 2010 //已结束
	SUPPLY_TASK_STATE_ACCEPTANCE uint = 2020 //验收完成

	LIVEANCHOR_COMMIT uint = 1 // 提交
	LIVEANCHOR_PASS   uint = 2 // 已审核通过
	LIVEANCHOR_REFUSE uint = 3 // 审核拒绝

	CHECK_STATE_NO   uint = 1000 //不需处理
	CHECK_STATE_TODO uint = 1001 //待处理
	CHECK_STATE_OK   uint = 1002 //处理完成
	CHECK_STATE_ERR  uint = 1003 //异常

	CONTRACT_TYPE_7_RETURN    uint = 1001 //7天无理由退货
	CONTRACT_TYPE_DENY_RETURN uint = 1006 //不支持退货

	ADMIN_PLANTFORM_USERID uint = 10001
)
