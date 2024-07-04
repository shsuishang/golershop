if ('undefined' == typeof window.SS)
{
    window.SS = {};
}

var SiteUrl = "https://test.suteshop.com";
var ApiUrl = "https://test.suteshop.com";
var pagesize = 10;
var WapSiteUrl = "https://test.suteshop.com/h5";
var IOSSiteUrl = "https://itunes.apple.com/us/app/b2b2c/id879996267?l=zh&ls=1&mt=8";
var AndroidSiteUrl = "http://www.shopsuite.cn/download/app/AndroidShopSuiteMoblie.apk";
var HallSiteUrl = "https://test.suteshop.com/hall";

var WapStaticUrl = "https://test.suteshop.com/h5";

var SiteLogo = "https://test.suteshop.com/image.php/shop/data/upload/media/user/10001/image/20210413/1618295584566944.png";
var SiteIcon = "https://test.suteshop.com/image.php/shop/data/upload/media/user/10001/image/20210413/1618295612543845.ico";
var WapSiteLogo = "https://test.suteshop.com/image.php/shop/data/upload/media/user/10001/image/20210413/1618295591886812.png";
var WechatStatus = 1;
var WechatBindMobile = 0;
var upload_type = "aliyun";

//扩展函数,需要放入lib
var __ = __ || function (str)
{
    return str;
};



function sprintf () {
    var regex = /%%|%(\d+$)?([\-+'#0 ]*)(\*\d+$|\*|\d+)?(?:\.(\*\d+$|\*|\d+))?([scboxXuideEfFgG])/g
    var a = arguments
    var i = 0
    var format = a[i++]

    var _pad = function (str, len, chr, leftJustify) {
        if (!chr) {
            chr = ' '
        }
        var padding = (str.length >= len) ? '' : new Array(1 + len - str.length >>> 0).join(chr)
        return leftJustify ? str + padding : padding + str
    }

    var justify = function (value, prefix, leftJustify, minWidth, zeroPad, customPadChar) {
        var diff = minWidth - value.length
        if (diff > 0) {
            if (leftJustify || !zeroPad) {
                value = _pad(value, minWidth, customPadChar, leftJustify)
            } else {
                value = [
                    value.slice(0, prefix.length),
                    _pad('', diff, '0', true),
                    value.slice(prefix.length)
                ].join('')
            }
        }
        return value
    }

    var _formatBaseX = function (value, base, prefix, leftJustify, minWidth, precision, zeroPad) {
        // Note: casts negative numbers to positive ones
        var number = value >>> 0
        prefix = (prefix && number && {
                '2': '0b',
                '8': '0',
                '16': '0x'
            }[base]) || ''
        value = prefix + _pad(number.toString(base), precision || 0, '0', false)
        return justify(value, prefix, leftJustify, minWidth, zeroPad)
    }

    // _formatString()
    var _formatString = function (value, leftJustify, minWidth, precision, zeroPad, customPadChar) {
        if (precision !== null && precision !== undefined) {
            value = value.slice(0, precision)
        }
        return justify(value, '', leftJustify, minWidth, zeroPad, customPadChar)
    }

    // doFormat()
    var doFormat = function (substring, valueIndex, flags, minWidth, precision, type) {
        var number, prefix, method, textTransform, value

        if (substring === '%%') {
            return '%'
        }

        // parse flags
        var leftJustify = false
        var positivePrefix = ''
        var zeroPad = false
        var prefixBaseX = false
        var customPadChar = ' '
        var flagsl = flags.length
        var j
        for (j = 0; j < flagsl; j++) {
            switch (flags.charAt(j)) {
                case ' ':
                    positivePrefix = ' '
                    break
                case '+':
                    positivePrefix = '+'
                    break
                case '-':
                    leftJustify = true
                    break
                case "'":
                    customPadChar = flags.charAt(j + 1)
                    break
                case '0':
                    zeroPad = true
                    customPadChar = '0'
                    break
                case '#':
                    prefixBaseX = true
                    break
            }
        }

        // parameters may be null, undefined, empty-string or real valued
        // we want to ignore null, undefined and empty-string values
        if (!minWidth) {
            minWidth = 0
        } else if (minWidth === '*') {
            minWidth = +a[i++]
        } else if (minWidth.charAt(0) === '*') {
            minWidth = +a[minWidth.slice(1, -1)]
        } else {
            minWidth = +minWidth
        }

        // Note: undocumented perl feature:
        if (minWidth < 0) {
            minWidth = -minWidth
            leftJustify = true
        }

        if (!isFinite(minWidth)) {
            throw new Error('sprintf: (minimum-)width must be finite')
        }

        if (!precision) {
            precision = 'fFeE'.indexOf(type) > -1 ? 6 : (type === 'd') ? 0 : undefined
        } else if (precision === '*') {
            precision = +a[i++]
        } else if (precision.charAt(0) === '*') {
            precision = +a[precision.slice(1, -1)]
        } else {
            precision = +precision
        }

        // grab value using valueIndex if required?
        value = valueIndex ? a[valueIndex.slice(0, -1)] : a[i++]

        switch (type) {
            case 's':
                return _formatString(value + '', leftJustify, minWidth, precision, zeroPad, customPadChar)
            case 'c':
                return _formatString(String.fromCharCode(+value), leftJustify, minWidth, precision, zeroPad)
            case 'b':
                return _formatBaseX(value, 2, prefixBaseX, leftJustify, minWidth, precision, zeroPad)
            case 'o':
                return _formatBaseX(value, 8, prefixBaseX, leftJustify, minWidth, precision, zeroPad)
            case 'x':
                return _formatBaseX(value, 16, prefixBaseX, leftJustify, minWidth, precision, zeroPad)
            case 'X':
                return _formatBaseX(value, 16, prefixBaseX, leftJustify, minWidth, precision, zeroPad)
                    .toUpperCase()
            case 'u':
                return _formatBaseX(value, 10, prefixBaseX, leftJustify, minWidth, precision, zeroPad)
            case 'i':
            case 'd':
                number = +value || 0
                // Plain Math.round doesn't just truncate
                number = Math.round(number - number % 1)
                prefix = number < 0 ? '-' : positivePrefix
                value = prefix + _pad(String(Math.abs(number)), precision, '0', false)
                return justify(value, prefix, leftJustify, minWidth, zeroPad)
            case 'e':
            case 'E':
            case 'f': // @todo: Should handle locales (as per setlocale)
            case 'F':
            case 'g':
            case 'G':
                number = +value
                prefix = number < 0 ? '-' : positivePrefix
                method = ['toExponential', 'toFixed', 'toPrecision']['efg'.indexOf(type.toLowerCase())]
                textTransform = ['toString', 'toUpperCase']['eEfFgG'.indexOf(type) % 2]
                value = prefix + Math.abs(number)[method](precision)
                return justify(value, prefix, leftJustify, minWidth, zeroPad)[textTransform]()
            default:
                return substring
        }
    }

    return format.replace(regex, doFormat)
}

 function buildUlr(url, param) {

    var LG = (function(lg){
        var objURL=function(url){
            this.ourl=url||window.location.href;
            this.href="";//?前面部分
            this.params={};//url参数对象
            this.jing="";//#及后面部分
            this.init();
        }
        //分析url,得到?前面存入this.href,参数解析为this.params对象，#号及后面存入this.jing
        objURL.prototype.init=function(){
            var str=this.ourl;
            var index=str.indexOf("#");
            if(index>0){
                this.jing=str.substr(index);
                str=str.substring(0,index);
            }
            index=str.indexOf("?");
            if(index>0){
                this.href=str.substring(0,index);
                str=str.substr(index+1);
                var parts=str.split("&");
                for(var i=0;i<parts.length;i++){
                    var kv=parts[i].split("=");
                    this.params[kv[0]]=kv[1];
                }
            }
            else{
                this.href=this.ourl;
                this.params={};
            }
        }
        //只是修改this.params
        objURL.prototype.set=function(key,val){
            this.params[key]=val;
        }
        //只是设置this.params
        objURL.prototype.remove=function(key){
            this.params[key]=undefined;
        }
        //根据三部分组成操作后的url
        objURL.prototype.url=function(){
            var strurl=this.href;
            var objps=[];//这里用数组组织,再做join操作
            for(var k in this.params){
                if(this.params[k]){
                    objps.push(k+"="+this.params[k]);
                }
            }
            if(objps.length>0){
                strurl+="?"+objps.join("&");
            }
            if(this.jing.length>0){
                strurl+=this.jing;
            }
            return strurl;
        }
        //得到参数值
        objURL.prototype.get=function(key){
            return this.params[key];
        }
        lg.URL=objURL;
        return lg;
    }(LG||{}));

    var obj =  new LG.URL(url);

    for(var o in param){
        obj.set(o, param[o]);
    }

    return obj.url();
}

function get_ext(filename){
    var postf = '';
    if (filename)
    {
        var index1=filename.lastIndexOf(".");

        var index2=filename.length;
        var postf=filename.substring(index1,index2);//后缀名
    }
    else
    {

    }

    return postf;
}

function image_thumb(image_url, w, h) {
    if ('undefined' == typeof w) {
        w = 60;
    }

    if ('undefined' == typeof h) {
        h = w;
    }


    var ext = get_ext(image_url);


    if (image_url.indexOf('image.php') !== -1)
    {
        image_url = sprintf('%s!%sx%s%s', image_url, w, h, ext);
    }
    else
    {
        if (upload_type == 'default')
        {
            image_url = sprintf('%s!%sx%s%s', image_url, w, h, ext);
        }
        else if (upload_type == 'aliyun')
        {
            //将图自动裁剪成宽度为 100px，高度为 100px 的效果图。 ?x-oss-process=image/resize,m_fill,h_100,w_100

            //http://image-demo.oss-cn-hangzhou.aliyuncs.com/example.jpg?x-oss-process=image/resize,m_mfit,h_100,w_100
            //将图缩略成宽度为 100px，高度为 100px，按长边优先 image/resize,m_lfit,h_100,w_100
            image_url = sprintf('%s?x-oss-process=image/resize,m_fill,h_%s,w_%s', image_url, w, h);
        }
        else
        {
            image_url = sprintf('%s!%sx%s%s', image_url, w, h, ext);
        }
    }

    return image_url;
}

img = image_thumb;



function number_format(number, decimals, dec_point, thousands_sep) {
    //  discuss at: http://phpjs.org/functions/number_format/
    // original by: Jonas Raoni Soares Silva (http://www.jsfromhell.com)
    // improved by: Kevin van Zonneveld (http://kevin.vanzonneveld.net)
    // improved by: davook
    // improved by: Brett Zamir (http://brett-zamir.me)
    // improved by: Brett Zamir (http://brett-zamir.me)
    // improved by: Theriault
    // improved by: Kevin van Zonneveld (http://kevin.vanzonneveld.net)
    // bugfixed by: Michael White (http://getsprink.com)
    // bugfixed by: Benjamin Lupton
    // bugfixed by: Allan Jensen (http://www.winternet.no)
    // bugfixed by: Howard Yeend
    // bugfixed by: Diogo Resende
    // bugfixed by: Rival
    // bugfixed by: Brett Zamir (http://brett-zamir.me)
    //  revised by: Jonas Raoni Soares Silva (http://www.jsfromhell.com)
    //  revised by: Luke Smith (http://lucassmith.name)
    //    input by: Kheang Hok Chin (http://www.distantia.ca/)
    //    input by: Jay Klehr
    //    input by: Amir Habibi (http://www.residence-mixte.com/)
    //    input by: Amirouche
    //   example 1: number_format(1234.56);
    //   returns 1: '1,235'
    //   example 2: number_format(1234.56, 2, ',', ' ');
    //   returns 2: '1 234,56'
    //   example 3: number_format(1234.5678, 2, '.', '');
    //   returns 3: '1234.57'
    //   example 4: number_format(67, 2, ',', '.');
    //   returns 4: '67,00'
    //   example 5: number_format(1000);
    //   returns 5: '1,000'
    //   example 6: number_format(67.311, 2);
    //   returns 6: '67.31'
    //   example 7: number_format(1000.55, 1);
    //   returns 7: '1,000.6'
    //   example 8: number_format(67000, 5, ',', '.');
    //   returns 8: '67.000,00000'
    //   example 9: number_format(0.9, 0);
    //   returns 9: '1'
    //  example 10: number_format('1.20', 2);
    //  returns 10: '1.20'
    //  example 11: number_format('1.20', 4);
    //  returns 11: '1.2000'
    //  example 12: number_format('1.2000', 3);
    //  returns 12: '1.200'
    //  example 13: number_format('1 000,50', 2, '.', ' ');
    //  returns 13: '100 050.00'
    //  example 14: number_format(1e-8, 8, '.', '');
    //  returns 14: '0.00000001'

    number = (number + '')
        .replace(/[^0-9+\-Ee.]/g, '');
    var n = !isFinite(+number) ? 0 : +number,
        prec = !isFinite(+decimals) ? 0 : Math.abs(decimals),
        sep = (typeof thousands_sep === 'undefined') ? '' : thousands_sep,
        dec = (typeof dec_point === 'undefined') ? '.' : dec_point,
        s = '',
        toFixedFix = function(n, prec) {
            var k = Math.pow(10, prec);
            return '' + (Math.round(n * k) / k)
                .toFixed(prec);
        };
    // Fix for IE parseFloat(0.55).toFixed(0) = 0;
    s = (prec ? toFixedFix(n, prec) : '' + Math.round(n))
        .split('.');
    if (s[0].length > 3) {
        s[0] = s[0].replace(/\B(?=(?:\d{3})+(?!\d))/g, sep);
    }
    if ((s[1] || '')
        .length < prec) {
        s[1] = s[1] || '';
        s[1] += new Array(prec - s[1].length + 1)
            .join('0');
    }
    return s.join(dec);
}

function mf(number, decimals, dec_point, thousands_sep) {
    //判断语言货币，修正 decimals
    return number_format(number, decimals, dec_point, thousands_sep);
}

//商品状态
var StateCode = {};

StateCode.DELIVERY_TIME_NO_LIMIT      = 1;    //不限送货时间：周一至周日
StateCode.DELIVERY_TIME_WORKING_DAY   = 2;    //工作日送货：周一至周五
StateCode.DELIVERY_TIME_WEEKEND       = 3;    //双休日、假日送货：周六至周日


StateCode.PRODUCT_STATE_ILLEGAL       = 1000; //违规下架禁售
StateCode.PRODUCT_STATE_NORMAL        = 1001; //正常
StateCode.PRODUCT_STATE_OFF_THE_SHELF = 1002; //下架


StateCode.ACTIVITY_TYPE_BARGAIN          = 1101; //加价购
StateCode.ACTIVITY_TYPE_GIFT             = 1102; //店铺满赠-小礼品
StateCode.ACTIVITY_TYPE_LIMITED_DISCOUNT = 1103; //限时折扣
StateCode.ACTIVITY_TYPE_DISCOUNT_PACKAGE = 1104; //优惠套装
StateCode.ACTIVITY_TYPE_VOUCHER          = 1105; //店铺优惠券  coupon 优惠券
StateCode.ACTIVITY_TYPE_DIY_PACKAGE      = 1106; //拼团
StateCode.ACTIVITY_TYPE_REDUCTION        = 1107; //满减

StateCode.ACTIVITY_TYPE_POINT_SHOPPING   = 1109; //积分换购

StateCode.ACTIVITY_TYPE_MARKETING      = 1131; //市场活动
StateCode.ACTIVITY_TYPE_LOTTERY        = 1121; //幸运大抽奖
StateCode.ACTIVITY_TYPE_FLASHSALE      = 1122; //秒杀
StateCode.ACTIVITY_TYPE_GROUPBOOKING   = 1123; //拼团
StateCode.ACTIVITY_TYPE_CUTPRICE       = 1124; //砍价
StateCode.ACTIVITY_TYPE_YIYUAN         = 1125; //一元购
StateCode.ACTIVITY_TYPE_GROUPBUY_STORE = 1126; //团购
StateCode.ACTIVITY_TYPE_PF_GROUPBUY_STORE = 1127; //批发团购


StateCode.ACTIVITY_GROUPBOOKING_SALE_PRICE     = 1; //以固定折扣购买
StateCode.ACTIVITY_GROUPBOOKING_FIXED_AMOUNT   = 2; //以固定价格购买
StateCode.ACTIVITY_GROUPBOOKING_FIXED_DISCOUNT = 3; //优惠固定金额

StateCode.MARKRTING_ACTIVITY_JOIN           = 1;//参加活动
StateCode.MARKRTING_ACTIVITY_JOIN_BY_QRCODE = 2;//通过二维码参加


StateCode.VOUCHER_STATE_UNUSED  = 1501; //未用
StateCode.VOUCHER_STATE_USED    = 1502; //已用
StateCode.VOUCHER_STATE_TIMEOUT = 1503; //过期
StateCode.VOUCHER_STATE_DEL     = 1504; //收回

    //商品标签
StateCode.PRODUCT_TAG_NEW     = 1401; //新品上架
StateCode.PRODUCT_TAG_REC     = 1402; //热卖推荐
StateCode.PRODUCT_TAG_BARGAIN = 1403; //清仓优惠
StateCode.PRODUCT_TAG_BARGAIN1 = 1404; //清仓优惠
StateCode.PRODUCT_TAG_CROSSBORDS = 1405; //清仓优惠

    //商品种类
StateCode.PRODUCT_KIND_ENTITY  = 1201; //实体商品	实物商品 （物流发货）
StateCode.PRODUCT_KIND_FUWU = 1202; //虚拟商品	虚拟商品 （无需物流）
StateCode.PRODUCT_KIND_CARD    = 1203; //电子卡券	电子卡券 （无需物流）
StateCode.PRODUCT_KIND_WAIMAI    = 1204; //外卖订单	外卖订单 （物流发货）
StateCode.PRODUCT_KIND_EDU    = 1205; //视频类 （无需物流发货）


StateCode.PRODUCT_VERIFY_REFUSED = 3000; //审核未通过
StateCode.PRODUCT_VERIFY_PASSED  = 3001; //审核通过
StateCode.PRODUCT_VERIFY_WAITING = 3002; //审核中

StateCode.ORDER_STATE_WAIT_PAY            = 2010; //待付款 - 虚拟映射
StateCode.ORDER_STATE_WAIT_PAID           = 2016; //已经付款 - 虚拟映射
StateCode.ORDER_STATE_WAIT_REVIEW         = 2011; //待订单审核
StateCode.ORDER_STATE_WAIT_FINANCE_REVIEW = 2013; //待财务审核
StateCode.ORDER_STATE_PICKING             = 2020; //待配货
StateCode.ORDER_STATE_WAIT_SHIPPING       = 2030; //待发货
StateCode.ORDER_STATE_SHIPPED             = 2040; //已发货
StateCode.ORDER_STATE_RECEIVED            = 2050; //已签收
StateCode.ORDER_STATE_FINISH              = 2060; //已完成
StateCode.ORDER_STATE_CANCEL              = 2070; //已取消
StateCode.ORDER_STATE_SELF_PICKUP         = 2080; //自提     交易关闭	         交易关闭
StateCode.ORDER_STATE_ERROR               = 2090; //异常订单
StateCode.ORDER_STATE_RETURN              = 2091; //退回订单 - 虚拟映射


StateCode.ORDER_PAID_STATE_NO             = 3010; //未付款
StateCode.ORDER_PAID_STATE_FINANCE_REVIEW = 3011; //待付款审核
StateCode.ORDER_PAID_STATE_PART           = 3012; //部分付款
StateCode.ORDER_PAID_STATE_YES            = 3013; //已付款

StateCode.ORDER_PICKING_STATE_NO             = 3020; //未出库
StateCode.ORDER_PICKING_STATE_PART           = 3021; //部分出库通过拆单解决这种问题
StateCode.ORDER_PICKING_STATE_YES            = 3022; //已出库

StateCode.ORDER_SHIPPED_STATE_NO             = 3030; //未发货
StateCode.ORDER_SHIPPED_STATE_PART           = 3031; //部分发货
StateCode.ORDER_SHIPPED_STATE_YES            = 3032; //已发货

StateCode.VIRTUAL_ORDER_USED    = 2101; //虚拟订单已使用
StateCode.VIRTUAL_ORDER_UNUSE   = 2100; //虚拟订单未使用
StateCode.VIRTUAL_ORDER_TIMEOUT = 2103; //虚拟订单过期

StateCode.ORDER_CANCEL_BY_BUYER  = 2201; //买家取消订单
StateCode.ORDER_CANCEL_BY_SELLER = 2202; //卖家取消订单
StateCode.ORDER_CANCEL_BY_ADMIN  = 2203; //平台取消


    //订单来源
StateCode.ORDER_FROM_PC     = 2301; //来源于pc端
StateCode.ORDER_FROM_WAP    = 2302; //来源于WAP手机端
StateCode.ORDER_FROM_WEBPOS = 2303; //来源于WEBPOS线下下单

    //状态
StateCode.SETTLEMENT_STATE_WAIT_OPERATE       = 2401; //已出账
StateCode.SETTLEMENT_STATE_SELLER_COMFIRMED   = 2402; //商家已确认
StateCode.SETTLEMENT_STATE_PLATFORM_COMFIRMED = 2403; //平台已审核
StateCode.SETTLEMENT_STATE_FINISH             = 2404; //结算完成

StateCode.ORDER_RETURN_NO  = 2500; //无退货
StateCode.ORDER_RETURN_ING = 2501; //退货中
StateCode.ORDER_RETURN_END = 2502; //退货完成

StateCode.ORDER_REFUND_STATE_NO  = 2600; //无退款
StateCode.ORDER_REFUND_STATE_ING = 2601; //退款中
StateCode.ORDER_REFUND_STATE_END = 2602; //退款完成


StateCode.ORDER_TYPE_DD = 3061; //订单类型
StateCode.ORDER_TYPE_FX = 3062; //分销订单
StateCode.ORDER_TYPE_TH = 3066; //分销订单


StateCode.ACTIVITY_STATE_WAITING  = 0; //活动状态:0-未开启
StateCode.ACTIVITY_STATE_NORMAL   = 1; //活动状态:1-正常
StateCode.ACTIVITY_STATE_FINISHED = 2; //活动状态:2-已结束
StateCode.ACTIVITY_STATE_CLOSED   = 3; //活动状态:3-管理员关闭

StateCode.GET_VOUCHER_FREE  = 1; //活动状态:1-免费参与;
StateCode.GET_VOUCHER_BY_POINT  = 2; //活动状态:2-积分参与;
StateCode.GET_VOUCHER_BY_PURCHASE = 3; //活动状态:3-购买参与


StateCode.CART_GET_TYPE_BUY     = 1; //购买
StateCode.CART_GET_TYPE_POINT   = 2; //积分兑换
StateCode.CART_GET_TYPE_GIFT    = 3; //赠品
StateCode.CART_GET_TYPE_BARGAIN = 4; //活动促销

    /*
StateCode.   BILL_TYPE_PO   = 4001;   //购货单
StateCode.   BILL_TYPE_PORO = 4002;   //销货退货单
StateCode.   BILL_TYPE_OI   = 4003;   //其他入库单
StateCode.   BILL_TYPE_SO   = 4031;   //销货单
StateCode.   BILL_TYPE_SORO = 4032;   //购货退货单
StateCode.   BILL_TYPE_OO   = 4033;   //其他出库单
    */

StateCode.   STOCK_IN_PURCHASE    = 2701;   //采购入库
StateCode.   STOCK_IN_RETURN      = 2702;   //退货入库
StateCode.   STOCK_IN_ALLOCATE    = 2703;   //调库入库
StateCode.   STOCK_IN_INVENTORY_P = 2704;   //盘盈入库
StateCode.   STOCK_IN_INIT        = 2705;   //期初入库
StateCode.   STOCK_IN_OTHER       = 2706;   //手工入库
StateCode.   STOCK_OUT_SALE       = 2751;   //销售出库
StateCode.   STOCK_OUT_DAMAGED    = 2752;   //损坏出库
StateCode.   STOCK_OUT_ALLOCATE   = 2753;   //调库出库
StateCode.   STOCK_OUT_LOSSES     = 2754;   //盘亏出库
StateCode.   STOCK_OUT_OTHER      = 2755;   //手工出库
StateCode.   STOCK_OUT_PO_RETURN  = 2756;   //损坏出库


StateCode.   STOCK_OUT_ALL = 2700;   //出库单
StateCode.   STOCK_IN_ALL  = 2750;   //入库单

StateCode.   BILL_TYPE_OUT = 2700;   //出库单
StateCode.   BILL_TYPE_IN  = 2750;   //入库单


StateCode.   BILL_TYPE_SO = 2800;   //销售订单
StateCode.   BILL_TYPE_PO = 2850;   //采购订单


    //修改掉，和订单状态对应。
StateCode.ORDER_PROCESS_SUBMIT         = 3070; //【客户】提交订单1OrderOrder

StateCode.ORDER_PROCESS_PAY            = 2010; //待支付Order
StateCode.ORDER_PROCESS_CHECK          = 2011; //订单审核1OrderOrder
StateCode.ORDER_PROCESS_FINANCE_REVIEW = 2013; //财务审核0OrderOrder
StateCode.ORDER_PROCESS_OUT            = 2020; //出库审核商品库存在“出库审核”节点完成后扣减，如需进行库存管理或核算销售成本毛利，需开启此节点。0OrderOrder
StateCode.ORDER_PROCESS_SHIPPED        = 2030; //发货确认如需跟踪订单物流信息，需开启此节点0OrderOrder
StateCode.ORDER_PROCESS_RECEIVED       = 2040; //【客户】收货确认0OrderOrder

StateCode.ORDER_PROCESS_FINISH         = 3098; //完成1OrderOrder

StateCode.RETURN_PROCESS_SUBMIT               = 3100; //【客户】提交退单1ReturnReturn
StateCode.RETURN_PROCESS_CHECK                = 3105; //退单审核1ReturnReturn
StateCode.RETURN_PROCESS_RECEIVED             = 3110; //收货确认0ReturnReturn
StateCode.RETURN_PROCESS_REFUND               = 3115; //退款确认0ReturnReturn
StateCode.RETURN_PROCESS_RECEIPT_CONFIRMATION = 3120; //客户】收款确认0ReturnReturn
StateCode.RETURN_PROCESS_FINISH               = 3125; //完成1ReturnReturn3130-商家拒绝退货
StateCode.RETURN_PROCESS_REFUSED              = 3130; //-商家拒绝退货
StateCode.RETURN_PROCESS_CANCEL               = 3135; //-买家取消


StateCode.PLANTFORM_RETURN_STATE_WAITING               = 3180; //申请状态平台(ENUM):3180-处理中;
StateCode.PLANTFORM_RETURN_STATE_AGREE               = 3181; //为待管理员处理卖家同意或者收货后;
StateCode.PLANTFORM_RETURN_PROCESS_FINISH               = 3182; //-为已完成


StateCode.STORE_STATE_WAIT_PROFILE       = 3210; //待完善资料
StateCode.STORE_STATE_WAIT_VERIFY        = 3220; //等待审核
StateCode.STORE_STATE_NO                 = 3230; //审核资料没有通过
StateCode.STORE_STATE_YES                = 3240; //审核资料通过,待付款

StateCode.TRADE_TYPE_SHOPPING   = 1201;//购物
StateCode.TRADE_TYPE_TRANSFER   = 1202;//转账
StateCode.TRADE_TYPE_DEPOSIT    = 1203;//充值
StateCode.TRADE_TYPE_WITHDRAW   = 1204;//提现
StateCode.TRADE_TYPE_SALES      = 1205;//销售
StateCode.TRADE_TYPE_COMMISSION = 1206;//佣金
StateCode.TRADE_TYPE_REFUND_PAY = 1207;//退货付款
StateCode.TRADE_TYPE_REFUND_GATHERING = 1208;//退货收款
StateCode.TRADE_TYPE_TRANSFER_GATHERING = 1209;//转账收款
StateCode.TRADE_TYPE_COMMISSION_TRANSFER = 1210;//佣金付款
StateCode.TRADE_TYPE_BONUS = 1211;//退货收款


StateCode.PAYMENT_TYPE_DELIVER = 1301;//货到付款
StateCode.PAYMENT_TYPE_ONLINE  = 1302;//在线支付
    //StateCode.PAYMENT_TYPE_CREDIT  = 1303;//白条支付
    //StateCode.PAYMENT_TYPE_CASH    = 1304;//现金支付
StateCode.PAYMENT_TYPE_OFFLINE = 1305;//线下支付

StateCode.ORDER_ITEM_EVALUATION_NO      = 0;    //未评价
StateCode.ORDER_ITEM_EVALUATION_YES   = 1;    //已评价
StateCode.ORDER_ITEM_EVALUATION_TIMEOUT       = 2;    //失效评价

StateCode.ORDER_EVALUATION_NO      = 0;    //未评价
StateCode.ORDER_EVALUATION_YES   = 1;    //已评价
StateCode.ORDER_EVALUATION_TIMEOUT       = 2;    //失效评价

StateCode.ORDER_NOT_NEED_RETURN_GOODS       = 0;    //不用退货
StateCode.ORDER_NEED_RETURN_GOODS           = 1;    //需要退货

StateCode.ORDER_REFUND           = 1;    //1-退款申请; 2-退货申请; 3-虚拟退款
StateCode.ORDER_RETURN           = 2;    //需要退货
StateCode.ORDER_VIRTUAL_REFUND           = 3;    //需要退货

StateCode.TO_STORE_SERVICE           = 1001;    //到店取货
StateCode.DOOR_TO_DOOR_SERVICE          = 1002;    // 上门服务




var User_BindConnectModel = {};

User_BindConnectModel.MOBILE     = 1;
User_BindConnectModel.EMAIL      = 2;

User_BindConnectModel.SINA_WEIBO = 11;
User_BindConnectModel.QQ         = 12;
User_BindConnectModel.WEIXIN     = 13;

User_BindConnectModel.ALIPAY     = 18;
User_BindConnectModel.FB     = 19;
User_BindConnectModel.GOOGLE     = 20;

;(function (factory) {
    if (typeof define === "function" && define.amd) {
        // AMD模式
        define(factory);
    } else {
        // 全局模式
        factory();
    }
}(function () {
    if ('undefined' == typeof window.SYS)
    {
        window.SYS = {};
    }


    SYS.VER   = 'undefined' != typeof SYS.VER ?  SYS.VER : '2.0.2445';
    SYS.DEBUG = 'undefined' != typeof SYS.DEBUG ?  SYS.DEBUG : 0;

    //尚未启用
    if (window.localStorage)
    {
        var cache = localStorage.getItem("cache");
        var cache_expire = localStorage.getItem("cache_expire");
    }
    else
    {
    }

    SYS.CACHE = !SYS.DEBUG && ('undefined' != typeof SYS.CACHE ? SYS.CACHE : 0);
    SYS.CACHE_EXPIRE = 3600 * 1;

    if ('undefined' == typeof SYS.CONFIG)
    {
        SYS.CONFIG = {
            base_url : "https://test.suteshop.com",
            index_url : "https://test.suteshop.com/index.php",
            admin_url : "https://test.suteshop.com/admin.php",
            account_url : "https://test.suteshop.com/account.php",
            index_page : "index.php",
            static_url : document.location.protocol + "https://test.suteshop.com/shop/static/src/default",
            static_lib_url : "https://test.suteshop.com/shop/static/src/common",
            kefu_atuo: "0"
        };
    }


    SYS.VER_SHOP   = 'undefined' != typeof SYS.VER_SHOP ?  SYS.VER :  '2.0.2445';
    SYS.VER_WAP   = 'undefined' != typeof SYS.VER_WAP ?  SYS.VER_WAP :  '2.0.2445';
    SYS.VER_ADMIN   = 'undefined' != typeof SYS.VER_ADMIN ?  SYS.VER_ADMIN : '2.0.2445';
    SYS.SW_ENABLE = 'undefined' != typeof SYS.SW_ENABLE ?  SYS.SW_ENABLE : 0;
    SYS.HTTPS = 'undefined' != typeof SYS.HTTPS ?  SYS.HTTPS :  1;


    if (window.localStorage) {
        var version = localStorage.getItem("version");

        if (version)
        {
            SYS.VER   = version;
        }
    } else {
    }

    SYS.STATIC_IMAGE_PATH = 'https://static.shopsuite.cn/xcxfile/appicon/';
    SYS.SYS_TYPE = 'multi';
    SYS.MULTISHOP_ENABLE = 1;
    SYS.STORE_ID = 1;

    SYS.EVALUATION_ENABLE = 1;
    SYS.SAAS_STATUS = 0;
    SYS.VIRTUAL_ENABLE = 1;
    SYS.O2O_ENABLE = 1;
    SYS.CHAIN_ENABLE = 0;
    SYS.PLANTFORM_FX_ENABLE = 1;
    SYS.STORE_FX_ENABLE = 1;
    SYS.PLANTFORM_SP_PRIZE_ENABLE = 1;

    SYS.PLANTFORM_FX_AGENT_ENABLE = 1;
    SYS.PLANTFORM_FX_PT_ENABLE = 1;
    SYS.PLANTFORM_FX_WESTORE_ENABLE = 0;
    SYS.PLANTFORM_DELIVERY_TIME_ENABLE = 0;
    SYS.PLANTFORM_USER_LEVEL_RATE_ENABLE = 0;

    SYS.STORE_SUPPLIER_ENABLE = 1;
    SYS.SUPPLIER_MARKET_ENABLE = 1;
    SYS.PLANTFORM_REBATE_ENABLE = 0;

    SYS.PAOTUI_ENABLE = 1;
    SYS.USERSTOCK_ENABLE = 0;
    SYS.ORDERCONFIRM_ENABLE = 0;
    SYS.MARKETORDER_ENABLE = 0;


    SYS.REDPACKET_ENABLE = 0;
    SYS.CREDIT_ENABLE = 0;
    SYS.POINT_ENABLE = 1;
    SYS.SP_ENABLE = 0;
    SYS.BP_ENABLE = 0;

    SYS.EXCHANGECARD_ENABLE = 0;
    SYS.MONEY_TRANSFER_ENABLE = 0;

    SYS.SNS_ENABLE = 1;
    SYS.IM_ENABLE = 1;
    SYS.SUBSITE_ENABLE = 1;
    SYS.SCAN_ENABLE = 0;


    SYS.AK_BROWSER = "Yi9XWlwa7sUGSuKGDiPBrS261bMeu6YF";
    SYS.AK_MINIAPP = "uWq8fmHbdvzOqLZlU8QZvbugoDyPFUg6";

    SYS.URL = {"index":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=mobile&typ=json","index_mobile":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=mobile&typ=json","index_app":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=app&typ=json","index_guide":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=guide&typ=json","center_menu":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=userCenterMenu&typ=json","subsite_list":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=getSubsite&typ=json","upload":"https:\/\/demo.mallsuite.cn\/front/sys/upload/index","upload_file":"https:\/\/test.suteshop.com\/index.php?ctl=Media&met=uploadFile&typ=json","upload_scrawl":"https:\/\/test.suteshop.com\/index.php?ctl=Media&met=uploadScrawl&typ=json","upload_config":"https:\/\/test.suteshop.com\/index.php?ctl=Media&met=config&typ=json","uptoken":"https:\/\/test.suteshop.com\/index.php?ctl=Media&met=uptoken&typ=json","info":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=info&typ=json","listTranslateLang":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=listTranslateLang&typ=json","page":"https:\/\/test.suteshop.com\/index.php?ctl=Page&met=get&typ=json","minipage":"https:\/\/test.suteshop.com\/index.php?ctl=Page&met=getMiniPage&typ=json","survey":"https:\/\/test.suteshop.com\/index.php?ctl=Page&met=doSurvey&typ=json","update_app":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=updateApp&typ=json","category_mobile_nav":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=getCategoryMobileNav&typ=json","product":{"item":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=item&typ=json","quick":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=quick&typ=json","info":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=info&typ=json","lists":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=lists&typ=json","auto_complete":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=listName&typ=json","shipping_district":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=shippingDistrict&typ=json","category":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=category&typ=json","brand":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=brand&typ=json","list_brands":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=listsBrands&typ=json","list_brands_products":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=listsBrandsAndProducts&typ=json","product_comment":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=comment&typ=json","add_comment_reply":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=addCommentReply&typ=json","comment_helpful":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=commentHelpful&typ=json","faq":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=faq&typ=json","popular":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=productPopular&typ=json","spec":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=getSpec&typ=json"},"loginInfo":"https:\/\/test.suteshop.com\/index.php?ctl=Index&met=getLoginInfo&typ=json","search_filter":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=getSearchFilter&typ=json","district":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=district&typ=json","district_all":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=getAllDistrict&typ=json","district_id":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=getDistrictByName&typ=json","search_hot_info":"https:\/\/test.suteshop.com\/index.php?ctl=Mobile&met=getSearchInfo&typ=json","connect":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=getConnectInfo&typ=json","login_box":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=index&typ=e&login_box=1","set_pwd":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=setNewPassword&typ=json","login":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=doLogin&typ=json","doSmsLogin":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=doSmsLogin&typ=json","register":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=doRegister&typ=json","logout":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=logout&typ=json","check_mobile_or_email":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=checkMobileOrEmail&typ=json","protocol":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=protocol&typ=json","check_login":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=checkLogin&typ=json","check_islogin":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=ifLogin&typ=json","app_login":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=doAppConnectLogin&typ=json&flag=app","do_app_login":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=doAppConnectLogin&typ=json&flag=app","check_app_login":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=checkAppConnectLogin&typ=json&flag=app","get_miniapp_open_id":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=getOpenIdByCode&typ=json&flag=app","register_wechat_account":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=doRegisterWechatAccount&typ=json","find_pwd_s2":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=findpwdStepTwo&typ=e&step=2","find_pwd_s3":"https:\/\/test.suteshop.com\/account.php?ctl=Login&met=findpwdStepThree&typ=e&step=3","delivery_info":"https:\/\/test.suteshop.com\/shop\/api\/delivery.php","sf":"https:\/\/test.suteshop.com\/shop\/api\/sf.php","download_proxy":"https:\/\/test.suteshop.com\/shop\/api\/download.php","qrcode":"https:\/\/test.suteshop.com\/shop\/api\/qrcode.php","account":{"certificate":"https:\/\/test.suteshop.com\/account.php?ctl=User_Security&met=certification&typ=e","get_mobile_info":"https:\/\/test.suteshop.com\/account.php?ctl=User_Security&met=index&typ=json","check_mobile_code":"https:\/\/test.suteshop.com\/account.php?ctl=User_Security&met=checkMobile&typ=json","get_mobile_checkcode":"https:\/\/test.suteshop.com\/account.php?ctl=VerifyCode&met=mobile&typ=json","get_email_checkcode":"https:\/\/test.suteshop.com\/account.php?ctl=VerifyCode&met=email&typ=json","check_security_change":"https:\/\/test.suteshop.com\/account.php?ctl=User_Security&met=checkSecurityChange&typ=json","reset_password":"https:\/\/test.suteshop.com\/account.php?ctl=User_Security&met=resetPassword&typ=json","bind_mobile":"https:\/\/test.suteshop.com\/account.php?ctl=User_Security&met=bindMobile&typ=json","commit_certificate":"https:\/\/test.suteshop.com\/account.php?ctl=User_Security&met=saveCertificate&typ=json","change_password":"https:\/\/test.suteshop.com\/account.php?ctl=User_Security&met=changePassword&typ=json","edit_user_info":"https:\/\/test.suteshop.com\/account.php?ctl=User_Account&met=edit&typ=json","edit_user_sign":"https:\/\/test.suteshop.com\/account.php?ctl=User_Account&met=sign&typ=json"},"wx":{"share":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=share&typ=e","config":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=wxConfig&typ=e&body_class_none=1","pay_config":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=wxPayConfig&typ=json","mplogin":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=login&typ=e&flag=mp","openlogin":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=login&typ=e&flag=open","applogin":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=jscode2session&typ=json&flag=app","checkAppLogin":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=checkAppLogin&typ=json&flag=app","getQRCode":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=getQRCode&typ=json&flag=app","getMiniAppQRCode":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=getMiniAppQRCode&typ=json&flag=app","getMiniAppQRCodeUnlimit":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=getMiniAppQRCodeUnlimit&typ=json&flag=app","getMiniAppQRCodeUnlimitPoster":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=getMiniAppQRCodeUnlimitPoster&typ=json&flag=app","pay":"https:\/\/test.suteshop.com\/account\/modules\/pay\/api\/payment\/wx\/pay.php","get_tpl_msg_config":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=getTmpMsgConfig&typ=json&flag=app","send_tpl_msg":"https:\/\/test.suteshop.com\/account.php?ctl=Connect_Weixin&met=sendTplMsg&typ=json&flag=app"},"store":{"get":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=get&typ=json","add":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=add&typ=json","credit":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=credit&typ=json","info":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=info&typ=json","contents":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=contents&typ=json","index_diy":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=indexDiy&typ=json","profile":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=profile&typ=json","menu":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=menu&typ=json","lists":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=lists&typ=json","listsChain":"https:\/\/test.suteshop.com\/index.php?ctl=Chain&met=lists&typ=json","getChain":"https:\/\/test.suteshop.com\/index.php?ctl=Chain&met=getChainInfo&typ=json","listsChainItems":"https:\/\/test.suteshop.com\/index.php?ctl=Chain&met=listsChainItems&typ=json","listsChainProduct":"https:\/\/test.suteshop.com\/index.php?ctl=Chain&met=listsChainProduct&typ=json","listChainByItem":"https:\/\/test.suteshop.com\/index.php?ctl=Chain&met=listChainByItem&typ=json","getNearChain":"https:\/\/test.suteshop.com\/index.php?ctl=Chain&met=getNearChain&typ=json","isDefautChain":"https:\/\/test.suteshop.com\/index.php?ctl=Chain&met=isDefautChain&typ=json","setChain":"https:\/\/test.suteshop.com\/index.php?ctl=Chain&met=setChain&typ=json","near":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=lists&typ=json","category":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=category&typ=json","activity":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=activity&typ=json","product":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=lists&typ=json","item":"https:\/\/test.suteshop.com\/index.php?ctl=Product&met=lists&typ=json","product_category":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=productCategory&typ=json","product_popular":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=productPopular&typ=json","lists_store_grade":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=listsStoreGrade&typ=json","lists_store_category":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=listsStoreCategory&typ=json","commit_payment_voucher":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=paymentVoucher&typ=json","survey":"https:\/\/test.suteshop.com\/index.php?ctl=Store&met=doSurvey&typ=json","lists_groupbuy_store":"https:\/\/test.suteshop.com\/index.php?ctl=Activity&met=listsGroupBuyStore&typ=json"},"point":{"voucher":"https:\/\/test.suteshop.com\/index.php?ctl=Point&met=voucher&typ=json","index":"https:\/\/test.suteshop.com\/index.php?ctl=Point&met=index&typ=json","product_detail":"https:\/\/test.suteshop.com\/index.php?ctl=Point&met=productDetail&typ=json","product":"https:\/\/test.suteshop.com\/index.php?ctl=Point&met=product&typ=json"},"verify":{"image":"https:\/\/test.suteshop.com\/account.php?ctl=VerifyCode&met=image&typ=json","mobile":"https:\/\/test.suteshop.com\/account.php?ctl=VerifyCode&met=mobile&typ=json","email":"https:\/\/test.suteshop.com\/account.php?ctl=VerifyCode&met=email&typ=json"},"diy":{"fonts":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getFonts&typ=json","getFonts":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getFonts&typ=json","getCartFileUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getCartFileUrl&typ=json","getModelShowUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getModelShow&typ=json","getItemSkuColorShowUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getItemSkuColorShowUrl&typ=json","getTypeBasesByParentUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getMaterialCategory&typ=json","queryOriginalByCategoryUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getMaterial&typ=json","removeBgColor":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=removeBgColor&typ=json","getStyleUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getStyleUrl&typ=json","imgstyUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getImageStyleUrl&typ=json","getUserDiyGoodsSkuListUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getUserDiyGoodsSkuListUrl&typ=json","saveDiyGoodsAndAddShopCartUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=saveDiyGoodsAndAddShopCartUrl&typ=json","upload":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=uploadImage&typ=json","upload_file":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=uploadFile&typ=json","upload_scrawl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=uploadScrawl&typ=json","upload_scrawl_diy":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=uploadScrawlDiy&typ=json","upload_config":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=config&typ=json","uptoken":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=uptoken&typ=json","getDespLists":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getDespLists&typ=json","buyGoodsUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=buyGoodsUrl&typ=json","getDesignList":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getDesignList&typ=json","getDiyMaskListUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getDiyMaskList&typ=json","getAllColorUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getAllColorUrl&typ=json","getTechnicsListUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getTechnicsListUrl&typ=json","getTechnicsPriceUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=getTechnicsPriceUrl&typ=json","searchMaterialsUrl":"https:\/\/test.suteshop.com\/index.php?mdu=diy&ctl=Diy&met=searchMaterialsUrl&typ=json"},"cart":{"add":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=add&typ=json","remove":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=remove&typ=json","edit":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=edit&typ=json","lists":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=lists&typ=json","listsMini":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=listsMini&typ=json","quantity":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=editQuantity&typ=json","cookie":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=getCookieCart&typ=json","index":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=index&typ=json","sel":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=sel&typ=json","checkout":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=checkout&typ=json","checkDelivery":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=checkDelivery&typ=json","order":"https:\/\/test.suteshop.com\/index.php?ctl=Cart&met=order&typ=e"},"cms":{"lists":"https:\/\/test.suteshop.com\/index.php?mdu=cms&ctl=Article&met=lists&typ=json","category":"https:\/\/test.suteshop.com\/index.php?mdu=cms&ctl=Article&met=category&typ=json","get":"https:\/\/test.suteshop.com\/index.php?mdu=cms&ctl=Article&met=get&typ=json","add_article_comment":"https:\/\/test.suteshop.com\/index.php?mdu=cms&ctl=Article&met=addComment&typ=json","add_article_comment_reply":"https:\/\/test.suteshop.com\/index.php?mdu=cms&ctl=Article&met=addCommentReply&typ=json","get_related_article":"https:\/\/test.suteshop.com\/index.php?mdu=cms&ctl=Article&met=getRelatedArticle&typ=json","comment_helpful":"https:\/\/test.suteshop.com\/index.php?mdu=cms&ctl=Article&met=commentHelpful&typ=json","remove_comment_helpful":"https:\/\/test.suteshop.com\/index.php?mdu=cms&ctl=Article&met=removeCommentHelpful&typ=json","comment_reply_helpful":"https:\/\/test.suteshop.com\/index.php?mdu=cms&ctl=Article&met=commentReplyHelpful&typ=json","remove_comment_reply_helpful":"https:\/\/test.suteshop.com\/index.php?mdu=cms&ctl=Article&met=removeCommentReplyHelpful&typ=json"},"user":{"overview":"https:\/\/test.suteshop.com\/index.php?ctl=User_Account&met=overview&typ=json","cart_count":"https:\/\/test.suteshop.com\/index.php?ctl=User_Account&met=getCartNum&typ=json","msg_count":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Message&met=getMsgCount&typ=json","msg_config":"https:\/\/test.suteshop.com\/account.php?ctl=Index&met=getConfig&typ=json","kefu_config":"https:\/\/test.suteshop.com\/account.php?ctl=Index&met=getKefuConfig&typ=json","msg_lists":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Message&met=lists&typ=json","msg_chat_lists":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Message&met=listChatMsg&typ=json","msg_get":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Message&met=get&typ=json","msg_set_read":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Message&met=setRead&typ=json","zonemsg_lists":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_ZoneMessage&met=lists&typ=json","msg_user_lists":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Message&met=getMsgUser&typ=json","msg_add":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Message&met=add&typ=json","msg_remove":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Message&met=remove&typ=json","msg_remove_user":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Message&met=removeUserMsg&typ=json","msg_notice_lists":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Message&met=listNotice&typ=json","lists_base_level":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Index&met=listBaseUserLevel&typ=json","friend_info_lists":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Friend&met=getFriendsInfo&typ=json","friend_agree":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Friend&met=agree&typ=json","friend_refuse":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Friend&met=refuse&typ=json","friend_chat":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Friend&met=chat&typ=json","more_comment":"https:\/\/test.suteshop.com\/index.php?ctl=User_Comment&met=loadMoreComment&typ=json","add_product_comment":"https:\/\/test.suteshop.com\/index.php?ctl=User_Comment&met=add&typ=json","add_order_comment":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=addOrderComment&typ=json","ask_helpful":"https:\/\/test.suteshop.com\/index.php?ctl=User_Ask&met=hasHelpful&typ=json","ask_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_Ask&met=add&typ=json","edit":"https:\/\/test.suteshop.com\/index.php?ctl=User_Account&met=edit&typ=json","lists":"https:\/\/test.suteshop.com\/index.php?ctl=User_Account&met=lists&typ=json","points":"https:\/\/test.suteshop.com\/index.php?ctl=User_Resource&met=pointsHistory&typ=json","money2points":"https:\/\/test.suteshop.com\/index.php?ctl=User_Resource&met=money2Points&typ=json","points2money":"https:\/\/test.suteshop.com\/index.php?ctl=User_Resource&met=points2Money&typ=json","listsExp":"https:\/\/test.suteshop.com\/index.php?ctl=User_Resource&met=listsExp&typ=json","listsExpRule":"https:\/\/test.suteshop.com\/index.php?ctl=User_Resource&met=listsExpRule&typ=json","signState":"https:\/\/test.suteshop.com\/index.php?ctl=User_Resource&met=signState&typ=json","resource":"https:\/\/test.suteshop.com\/index.php?ctl=User_Resource&met=resource&typ=json","signIn":"https:\/\/test.suteshop.com\/index.php?ctl=User_Resource&met=signIn&typ=json","lists_chain_code":"https:\/\/test.suteshop.com\/index.php?ctl=User_Resource&met=listsChainCode&typ=json","lists_voucher_product":"https:\/\/test.suteshop.com\/index.php?ctl=User_Voucher&met=listsVoucherProduct&typ=json","voucherList":"https:\/\/test.suteshop.com\/index.php?ctl=User_Voucher&met=lists&typ=json","voucher_get":"https:\/\/test.suteshop.com\/index.php?ctl=User_Voucher&met=getVoucher&typ=json","voucherNum":"https:\/\/test.suteshop.com\/index.php?ctl=User_Voucher&met=getVoucherNum&typ=json","eachVoucherNum":"https:\/\/test.suteshop.com\/index.php?ctl=User_Voucher&met=getEachVoucherNum&typ=json","voucher_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_Voucher&met=add&typ=json","voucher_used":"https:\/\/test.suteshop.com\/index.php?ctl=User_Voucher&met=checkVoucher&typ=json","voucher_offline":"https:\/\/test.suteshop.com\/index.php?ctl=User_Voucher&met=addOffline&typ=json","receive_new_gift":"https:\/\/test.suteshop.com\/index.php?ctl=User_Voucher&met=receiveGift&typ=json","return_lists":"https:\/\/test.suteshop.com\/index.php?ctl=User_Return&met=lists&typ=json","return_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_Return&met=add&typ=json","return_add_one":"https:\/\/test.suteshop.com\/index.php?ctl=User_Return&met=addItem&typ=json","return_item":"https:\/\/test.suteshop.com\/index.php?ctl=User_Return&met=returnItem&typ=json","return_get":"https:\/\/test.suteshop.com\/index.php?ctl=User_Return&met=get&typ=json","return_cancel":"https:\/\/test.suteshop.com\/index.php?ctl=User_Return&met=cancel&typ=json","return_confirm_refund":"https:\/\/test.suteshop.com\/index.php?ctl=User_Return&met=confirmRefund&typ=json","invoice_type":"https:\/\/test.suteshop.com\/index.php?ctl=User_Invoice&met=type&typ=json","invoice_lists":"https:\/\/test.suteshop.com\/index.php?ctl=User_Invoice&met=lists&typ=json","invoice_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_Invoice&met=add&typ=json","invoice_get":"https:\/\/test.suteshop.com\/index.php?ctl=Activity&met=getVoucher&typ=json","invoice_remove":"https:\/\/test.suteshop.com\/index.php?ctl=User_Invoice&met=remove&typ=json","add_point_shopping_order":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=addPointShoppingOrder&typ=json","order_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=add&typ=json","order_lists":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=lists&typ=json","order_cancel":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=cancel&typ=json","order_receive":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=receive&typ=json","order_remove":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=remove&typ=json","order_detail":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=detail&typ=json","order_delivery":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=delivery&typ=json","order_index":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=index&typ=e","order_comment_manage":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=storeCommentManage&typ=json","order_comment_with_content":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=storeEvaluationWithContent&typ=json","order_comment_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=addOrderComment&typ=json","wish_store_lists":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=lists&typ=json&action=store","wish_item_lists":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=lists&typ=json&action=item","wish_brand_lists":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=lists&typ=json&action=brand","wish_store_get":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=get&typ=json&action=store","wish_item_get":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=get&typ=json&action=item","wish_brand_get":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=get&typ=json&action=brand","wish_store_remove":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=remove&typ=json&action=store","wish_item_remove":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=remove&typ=json&action=item","wish_brand_remove":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=remove&typ=json&action=brand","wish_store_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=add&typ=json&action=store","wish_item_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=add&typ=json&action=item","wish_brand_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=add&typ=json&action=brand","browser_lists":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=browser&typ=json&action=item","browser_remove":"https:\/\/test.suteshop.com\/index.php?ctl=User_Favorites&met=removeBrowser&typ=json&action=item","feedback_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_Feedback&met=add&typ=json","check_wechat_address":"https:\/\/test.suteshop.com\/index.php?ctl=User_DeliveryAddress&met=checkWeChatAddress&typ=json","address_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_DeliveryAddress&met=save&typ=json","address_get":"https:\/\/test.suteshop.com\/index.php?ctl=User_DeliveryAddress&met=get&typ=json","address_edit":"https:\/\/test.suteshop.com\/index.php?ctl=User_DeliveryAddress&met=save&typ=json","address_lists":"https:\/\/test.suteshop.com\/index.php?ctl=User_DeliveryAddress&met=lists&typ=json","address_remove":"https:\/\/test.suteshop.com\/index.php?ctl=User_DeliveryAddress&met=remove&typ=json","address_manage":"https:\/\/test.suteshop.com\/index.php?ctl=User_DeliveryAddress&met=manage&typ=e","joinIn":"https:\/\/test.suteshop.com\/index.php?ctl=Join&met=step&typ=e","listsLottery":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=getLottery&typ=json","doLottery":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=doLottery&typ=json","listsLotteryHistory":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=listsLotteryHistory&typ=json","getLotteryHistory":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=getLotteryHistory&typ=json","updateLotteryAddress":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=updateLotteryAddress&typ=json","listsMarketing":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=listsMarketing&typ=json","listsUserMarketing":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=listsUserMarketing&typ=json","getMarketing":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=getMarketing&typ=json","doMarketing":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=doMarketing&typ=json","doCutPrice":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=doCutPrice&typ=json","listsUserCutPriceHistory":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=listsUserCutPriceHistory&typ=json","listsCutPriceActivity":"https:\/\/test.suteshop.com\/index.php?ctl=Activity&met=listsCutPriceActivity&typ=json","getCutPriceActivity":"https:\/\/test.suteshop.com\/index.php?ctl=Activity&met=getCutPriceActivity&typ=json","listsCutPriceHistory":"https:\/\/test.suteshop.com\/index.php?ctl=Activity&met=listsCutPriceHistory&typ=json","listsGroupbookingActivity":"https:\/\/test.suteshop.com\/index.php?ctl=Activity&met=listsGroupbookingActivity&typ=json","listsGroupbooking":"https:\/\/test.suteshop.com\/index.php?ctl=Activity&met=listsGroupbooking&typ=json","listsUserGroupbooking":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=listsUserGroupbooking&typ=json","getUserGroupbooking":"https:\/\/test.suteshop.com\/index.php?ctl=User_Activity&met=getUserGroupbooking&typ=json","direct_store_add":"https:\/\/test.suteshop.com\/index.php?ctl=User_DistributionStoreDirectsellerProduct&met=add&typ=json","direct_store_lists":"https:\/\/test.suteshop.com\/index.php?ctl=User_DistributionStoreDirectsellerProduct&met=index&typ=json","direct_store_remove":"https:\/\/test.suteshop.com\/index.php?ctl=User_DistributionStoreDirectsellerProduct&met=remove&typ=json","direct_store_index":"https:\/\/test.suteshop.com\/index.php?mdu=distribution&ctl=WeStore&met=index&typ=json","listsUserVideo":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=listsUserVideo&typ=json","confirmOrderItemStatus":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=confirmOrderItemStatus&typ=json","getOrderLog":"https:\/\/test.suteshop.com\/index.php?ctl=User_Order&met=getOrderLog&typ=json","add_message":"https:\/\/test.suteshop.com\/index.php?ctl=User_Info&met=add&typ=json"},"pay":{"type":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=payType&typ=json","lists":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=payLists&typ=json","recharge":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=recharge&typ=json","recharge_level":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=rechargeByLevel&typ=json","recharge_list":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=listRechargeLevel&typ=json","favorable":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=favorable&typ=json","other":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=other&typ=json","pay":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=pay&typ=e","check_pay_passwd":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=checkPayPasswd&typ=json","get_pay_passwd":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=getPayPasswd&typ=json","asset":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=resourceIndex&typ=json","consume_record":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=consumeRecord&typ=json","get_consume_record":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=getConsumeRecord&typ=json","consume_deposit":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=consumeDeposit&typ=json","consume_trade":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=consumeTrade&typ=json","consume_trade_detail":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=getConsumeTradeDetail&typ=json","consume_withdraw":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=consumeWithdraw&typ=json","consume_withdraw_info":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=withdrawInfo&typ=json","consume_withdraw_add":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=addWithdraw&typ=json","addCard":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Card&met=addCard&typ=json","cardHistory":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Card&met=cardHistory&typ=json","reset_paypasswd":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=resetPayPassword&typ=json","change_paypasswd":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=changePayPassword&typ=json","receive_qrcode":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=receiveQrcode&typ=json","transfer":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=transfer&typ=json","transfer_user_row":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=transferUserRow&typ=json","addOfflineConsumeRecord":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=addOfflineConsumeRecord&typ=json","consumeRecordByStoreBook":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=consumeRecordByStoreBook&typ=json","list_user_bank":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=listUserBank&typ=json","get_user_bank":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=getUserBank&typ=json","add_user_bank":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=addUserBank&typ=json","remove_user_bank":"https:\/\/test.suteshop.com\/account.php?mdu=pay&ctl=Index&met=removeUserBank&typ=json"},"exchange":{"sp_exchange":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=exchange&typ=json","sp_transfer":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=transfer&typ=json","point_transfer":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=transferPoint&typ=json","sp_sale":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=sale&typ=json","sp_buy":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=buy&typ=json","bp_exchange":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Bp&met=exchange&typ=json","bp_market":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Bp&met=market&typ=json","sp_history":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=listSpHistory&typ=json","sp_trade":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=listSpMarket&typ=json","sp_trade_cancel":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=cancelSpMarket&typ=json","sp_trade_buy":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=buySpMarket&typ=json","do_sp_exchange":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=doSpExchange&typ=json","do_sp_transfer":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=doSpTransfer&typ=json","do_point_transfer":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=doPointTransfer&typ=json","do_sp_sale":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Sp&met=doSpSale&typ=json","bp_history":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Bp&met=listBpHistory&typ=json","do_bp_exchange":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Bp&met=doBpExchange&typ=json","do_bp_market":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Bp&met=doBpMarket&typ=json","funds_lists":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Funds&met=listFunds&typ=json","funds_profit_lists":"https:\/\/test.suteshop.com\/account.php?mdu=exchange&ctl=Funds&met=listFundsProfit&typ=json"},"fx":{"invite":"https:\/\/test.suteshop.com\/index.php?ctl=Distribution_User&met=invite&typ=json","poster":"https:\/\/test.suteshop.com\/index.php?ctl=Distribution_User&met=poster&typ=json","withdraw":"https:\/\/test.suteshop.com\/index.php?ctl=Distribution_User&met=withdraw&typ=json","doWithdraw":"https:\/\/test.suteshop.com\/index.php?ctl=Distribution_User&met=doWithdraw&typ=json","lists":"https:\/\/test.suteshop.com\/index.php?ctl=Distribution_User&met=lists&typ=json","lists_commission":"https:\/\/test.suteshop.com\/index.php?ctl=Distribution_User&met=listsCommission&typ=json","lists_team":"https:\/\/test.suteshop.com\/index.php?ctl=Distribution_User&met=listsTeam&typ=json","lists_order":"https:\/\/test.suteshop.com\/index.php?ctl=Distribution_User&met=listsOrder&typ=json","index":"https:\/\/test.suteshop.com\/index.php?ctl=Distribution_User&met=index&typ=json"},"admin":{"pay":{"orderRecord":"https:\/\/test.suteshop.com\/admin.php?mdu=pay&ctl=Consume_Record&met=orderRecord&typ=json","offline":"https:\/\/test.suteshop.com\/admin.php?mdu=pay&ctl=Consume_Trade&met=offline&typ=json"},"lists":{"payment_channel":"https:\/\/test.suteshop.com\/admin.php?ctl=Category&met=lists&typ=json&isDelete=2&type_number=payment_channel_id","express":"https:\/\/test.suteshop.com\/admin.php?ctl=Category&met=lists&typ=json&isDelete=2&type_number=express_id"}},"sns":{"story_timeline_lists":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Timeline&met=lists&typ=json","story_lists":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story&met=lists&typ=json","story_get":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story&met=get&typ=json","story_add":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Base&met=add&typ=json","story_remove":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Base&met=remove&typ=json","story_edit":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Base&met=edit&typ=json","story_comment_add":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Comment&met=add&typ=json","story_comment_remove":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Comment&met=remove&typ=json","story_comment_lists":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story&met=listComment&typ=json","story_rel_data":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story&met=getRelData&typ=json","category_lists":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story&met=listCategory&typ=json","story_collection_add":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Collection&met=add&typ=json","story_collection_remove":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Collection&met=remove&typ=json","story_like_add":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Like&met=add&typ=json","story_like_remove":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Like&met=remove&typ=json","story_comment_like_add":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Comment&met=like&typ=json","story_comment_like_remove":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Comment&met=unlike&typ=json","story_comment_reply_like_add":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Comment&met=likeReply&typ=json","story_comment_reply_like_remove":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=Story_Comment&met=unlikeReply&typ=json","user_space":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User&met=space&typ=json","user_story":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User&met=story&typ=json","user_comment_story":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User&met=listUserCommentStory&typ=json","user_collect_story":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User&met=listUserCollectStory&typ=json"},"seller":{"store":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_Base&met=get&typ=json","product":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Product&met=lists&typ=json","product_add":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Product&met=add&typ=json","get_store_info":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_Base&met=get&typ=json","edit_store_info":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_Base&met=edit&typ=json","lists_product":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Product_Base&met=lists&typ=json","edit_state":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Product_Base&met=editState&typ=json","remove_product":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Product_Base&met=remove&typ=json","dashboard":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Index&met=dashboard&typ=json","order_lists":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=listsOrderAndsProduct&typ=json","edit_fee":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=editFee&typ=json","review":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=review&typ=json","review_finance":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=reviewFinance&typ=json","review_picking":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=reviewPicking&typ=json","cancel_order":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=cancel&typ=json","getOrderStock":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=getOrderStock&typ=json","saveOrderLogistics":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=saveOrderLogistics&typ=json","lists_shipping_address":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ShippingAddress&met=lists&typ=json","del_shipping_address":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ShippingAddress&met=remove&typ=json","add_shipping_address":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ShippingAddress&met=add&typ=json","edit_shipping_address":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ShippingAddress&met=edit&typ=json","get_shipping_address":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ShippingAddress&met=get&typ=json","get_express_logistics":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ExpressLogistics&met=get&typ=json","select_express_logistics":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ExpressLogistics&met=selectDefault&typ=json","enabled_express_logistics":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ExpressLogistics&met=selectEnabled&typ=json","lists_express_logistics":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ExpressLogistics&met=lists&typ=json","del_express_logistics":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ExpressLogistics&met=remove&typ=json","add_express_logistics":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ExpressLogistics&met=add&typ=json","edit_express_logistics":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ExpressLogistics&met=edit&typ=json","order_cancel":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=cancel&typ=json","order_receive":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=receive&typ=json","order_remove":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=remove&typ=json","order_detail":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=detail&typ=json","order_get_by_pickupcode":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=getOrderByPickUpCode&typ=json","do_pickup":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=doPickUp&typ=json","order_delivery":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=delivery&typ=json","market_lists":"https:\/\/test.suteshop.com\/admin.php?mdu=city&ctl=Market_OrderBase&met=lists&typ=json","waybill_tpl_lists":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_WaybillTpl&met=lists&typ=json","waybill_tpl_add":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_WaybillTpl&met=add&typ=json","express_logistics_lists":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ExpressLogistics&met=lists&typ=json","express_logistics_add":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ExpressLogistics&met=add&typ=json","shipping_address_lists":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ShippingAddress&met=lists&typ=json","shipping_address_add":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ShippingAddress&met=add&typ=json","add_store_book":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_Book&met=add&typ=json","lists_store_book":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_Book&met=lists&typ=json","del_store_book":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_Book&met=remove&typ=json","edit_store_book":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_Book&met=edit&typ=json","transport_type_lists":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_TransportType&met=lists&typ=json","productCategory":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Store_ProductCategory&met=treeview&typ=json"},"city":{"paotui_info":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=User_Paotuier&met=info&typ=json","paotui_active":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=User_Paotuier&met=setActive&typ=json","paotui_apply":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=User_Paotuier&met=add&typ=json","paotui_order_list":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=User_Paotuier&met=getOrderList&typ=json","paotui_finish_order":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=User_Paotuier&met=getFinishOrderList&typ=json","paotui_certificate":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=User_Paotuier&met=saveCertificate&typ=json","team_add":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=Paotui_Team&met=add&typ=json","order_add":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=Paotui_Order&met=add&typ=json","order_lists":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=Paotui_Order&met=lists&typ=json","order_apply":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=Paotui_Order&met=apply&typ=json","order_pickup":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=Paotui_Order&met=pickup&typ=json","order_signin":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=Paotui_Order&met=signin&typ=json","order_get":"https:\/\/test.suteshop.com\/index.php?mdu=city&ctl=Paotui_Order&met=get&typ=json","receive_order":"https:\/\/test.suteshop.com\/admin.php?mdu=city&ctl=Market_OrderBase&met=getMarketOrder&typ=json","confirm_receive":"https:\/\/test.suteshop.com\/admin.php?mdu=city&ctl=Market_OrderBase&met=enable&typ=json","store_Base_list":"https:\/\/test.suteshop.com\/admin.php?ctl=Store_Base&met=lists&typ=json","transfer_market_order":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Base&met=transferMarketOrder&typ=json","upload_images":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Item&met=updateItemFile&typ=json","multipload_images":"https:\/\/test.suteshop.com\/admin.php?mdu=seller&ctl=Order_Item&met=multiUpload&typ=json","riderLists":"https:\/\/test.suteshop.com\/admin.php?mdu=city&ctl=Store_PaotuierRel&met=lists&typ=json","send":"https:\/\/test.suteshop.com\/admin.php?mdu=city&ctl=Paotui_OrderBase&met=send&typ=json"},"hall":{"index":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=index&typ=e","orderForIndex":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=orderForIndex&typ=e","member":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=member&typ=e","workerHome":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=home&typ=e","postAddress":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=postAddress&typ=e","companyIndex":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=index&typ=e","companyOrder":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=order&typ=e","companyPayment":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Service_OrderBase&met=payment&typ=e","companyPaymentSuccess":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Service_OrderBase&met=paymentSuccess&typ=e","companyInfo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=info&typ=e","workerIndex":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=index&typ=e","workerOrder":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=order&typ=e","workerInfo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=info&typ=e","orderManage":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Service_OrderBase&met=manage&typ=e","workerComment":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=comment&typ=e","companyComment":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=comment&typ=e","orderDetail":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Service_OrderBase&met=orderDetail&typ=e","editBankCard":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_BankCard&met=editBankCard&typ=e","delBankCard":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_BankCard&met=delBankCard&typ=json","myBankCard":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_BankCard&met=myBankCard&typ=e","getBankList":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_BankCard&met=getBankList&typ=json","getUserBankList":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_BankCard&met=getUserBankList&typ=json","saveBankCardInfo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_BankCard&met=saveBankCardInfo&typ=json","getTheBankCardInfo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_BankCard&met=getTheBankCardInfo&typ=json","isLogin":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=isHallLogin&typ=json","getCity":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=getCity&typ=json","scrollPic":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=scrollPic&typ=json","getStars":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=getStars&typ=json","getCompanyInfo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=getCompanyInfo&typ=json","getCodeList":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=getCodeList&typ=json","saveCompanyValidate":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=saveCompanyValidate&typ=json","saveCompanyInfo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=saveCompanyInfo&typ=json","isValidateIdentity":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=isValidateIdentity&typ=json","savePersonValidate":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=savePersonValidate&typ=json","getAuthorizeSign":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=getAuthorizeSign&typ=json","getSkill":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=getSkill&typ=json","completePersonInfo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=completePersonInfo&typ=json","saveServiceOrderInfo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Service_OrderBase&met=saveServiceOrderInfo&typ=json","orderListPartial":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Service_OrderBase&met=orderListPartial&typ=json","orderListPartialForCompany":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=orderListPartial&typ=json","orderListPartialForWorker":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=orderListPartial&typ=e","getOrder":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Service_OrderBase&met=getOrder&typ=json","getUserInfo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=getUserInfo&typ=json","addService":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Service&met=add&typ=json","editService":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Service&met=edit&typ=json","getService":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Service&met=get&typ=json","delService":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Service&met=remove&typ=json","myServiceListPartial":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Service&met=lists&typ=json","checkSignUp":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=checkSingUp&typ=json","getListByCate":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=getListByCate&typ=json","getHomeData":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=getHomeData&typ=json","serviceListPartial":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=serviceListPartial&typ=json","getOrderList":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=orderForIndex&typ=json","getOrderDetail":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Service_OrderBase&met=order&typ=json","getWorkerStudio":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Worker_Studio&met=lists&typ=json","checkTax":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=checkTax&typ=json","signUpListPartial":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=signUpListPartial&typ=e","saveSignUp":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=saveSignUp&typ=json","checkChoose":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Index&met=checkChoose&typ=json","cancelSign":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=cancelSign&typ=json","newBalancePay":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=newBalancePay&typ=json","countTime":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=countTime&typ=json","requestedChange":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=requestedChange&typ=json","submitCheck":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=submitCheck&typ=json","getSignUpInfo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=getSignUpInfo&typ=json","getSignWorker":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=getSignWorker&typ=json","getUserBank":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=getUserBank&typ=json","savePostAddress":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=savePostAddress&typ=json","getPostAddress":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=getPostAddress&typ=json","delPostAddress":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=delPostAddress&typ=json","saveEmail":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=saveEmail&typ=json","getEmail":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=getEmail&typ=json","delEmail":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=delEmail&typ=json","getAddress":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=getAddress&typ=json","getAddressByOne":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=getAddressByOne&typ=json","getEmails":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=getEmails&typ=json","newSubmitQualified":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=newSubmitQualified&typ=json","newSubmitQuaItem":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=newSubmitQuaItem&typ=json","saveJudgeLevelByCompany":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Company&met=saveJudgeLevel&typ=json","saveJudgeLevelByWorker":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Worker&met=saveJudgeLevel&typ=json","getUserBlanceByNo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Resource&met=getUserBalanceByNo&typ=json","systemMessage":"https:\/\/test.suteshop.com\/account.php?mdu=sns&ctl=User_Message&met=index&typ=e","getBillsList":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Bill&met=index&typ=json","setDefaultPostAddress":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=setDefaultPostAddress&typ=json","saveUserBank":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=saveUserBank&typ=json","getServiceListForSale":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Invoice&met=getServiceListForSale&typ=json","getServiceInfo":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Service&met=getServiceInfo&typ=json","addServiceOrder":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=User_Service&met=addServiceOrder&typ=json","register":"https:\/\/test.suteshop.com?mdu=account&ctl=Login&met=register&typ=e","studioLookup":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Worker_Studio&met=studioLookup&typ=json","studioNum":"https:\/\/test.suteshop.com\/index.php?mdu=hall&ctl=Worker_Studio&met=studioNum&typ=json"},"live":{"checkPerm":"https:\/\/test.suteshop.com\/index.php?mdu=live&ctl=Live&met=checkPerm&typ=e","saveRoom":"https:\/\/test.suteshop.com\/index.php?mdu=live&ctl=Live&met=saveRoom&typ=e","index":"https:\/\/test.suteshop.com\/index.php?mdu=live&ctl=Live&met=index&typ=e"}}

    SYS.CONFIG.URL = SYS.URL;

    return SYS.CONFIG;
}));


// utils.js
if ('undefined' == typeof window.verifyUtils)
{
    window.verifyUtils = {};
}

window.verifyUtils = {
    smsTimer: function (that)
    {
        var self = this;

        var wait = $(that).data('wait');

        if (wait == 0)
        {
            $(that).removeAttr("disabled").val('重新获取验证码');
            $(that).removeClass("disabled");

            $(that).data('status', true);
            $(that).data('wait', $(that).data('times'));
        }
        else
        {
            $(that).attr("disabled", true).val(wait + '秒后点此重发');
            $(that).addClass("disabled");

            $(that).data('wait', --wait);

            setTimeout(function ()
            {
                self.smsTimer(that);
            }, 1000)
        }
    },

    countDown:function (that, times)
    {
        var self = this;

        $(that).data('times', times);
        $(that).data('wait', times);

        if (typeof($(that).data('status')) === 'undefined' || $(that).data('status'))
        {
            $(that).data('status', false);
            self.smsTimer(that);
        }
    },



    imageVerifyCode : function (verify_code, key_obj)
    {
        $(verify_code).on('click', function()
        {
            var rand_key = Math.random();
            var url = itemUtil.getUrl(SYS.URL.verify.image, {rand_key: rand_key});

            $(verify_code).css("backgroundImage","url(" + url + ")");
            $(key_obj).val(rand_key);
        });

        //$(this).css("backgroundImage","url(" + url + ")");
        var rand_key = Math.random();
        var url = itemUtil.getUrl(SYS.URL.verify.image, {rand_key: rand_key});

        $(verify_code).css("background","url(" + url + ") no-repeat center");
        $(verify_code).css("cursor", "pointer");
        $(verify_code).css("backgroundSize", "cover");
        $(key_obj).val(rand_key);
    },


    emailVerifyCode : function (email, that, captcha)
    {
        var self = this;
        $(that).on('click', function()
        {
            let mv = '';
            if (typeof(email) === 'object')
            {
                mv = $(email).val();
            }
            else
            {
                mv = email;
            }

            var url = itemUtil.getUrl(SYS.URL.verify.email, {rand_key: mv})

            if (typeof(captcha) !== 'undefined')
            {
                url = url + '&captcha=' + captcha;
            }

            $.ajax({
                type: "GET",
                url: url,
                data: {},
                dataType: "jsonp",
                jsonp: "jsonp_callback",
                success: function(res){

                    if (200 == res.status)
                    {
                        //服务端返回times
                        var times = 60;
                        self.countDown(that, times);
                    }
                    else
                    {
                        if (typeof $.fancybox != 'undefined')
                        {
                            Public.tipMsg(res.msg);
                        }
                        else
                        {
                            alert(res.msg);
                        }
                    }
                }
            });
        });
    },


    smsVerifyCode : function (mobile, that, image_key, image_value)
    {
        var self = this;
        $(that).on('click', function()
        {
            let mv = '';
            if (typeof(mobile) === 'object')
            {
                mv =   '' + $('#user_intl').val() + $(mobile).val();

                if (!$(mobile).val())
                {
                     alert(__('请输入手机号码'));
                     return ;
                }
            }
            else
            {
                mv = mobile;

                if (!mv)
                {
                     alert(__('请输入手机号码'));
                     return ;
                }
            }

            var url = itemUtil.getUrl(SYS.URL.verify.mobile, {mobile: encodeURIComponent(mv)});

            if (typeof(image_value) !== 'undefined')
            {
                url = url + '&image_value=' + $(image_value).val();;
            }

            if (typeof(image_key) !== 'undefined')
            {
                url = url + '&image_key=' + $(image_key).val();;
            }

            $.ajax({
                type: "post",
                url: url,
                data: {},
                dataType: "jsonp",
                jsonp: "jsonp_callback",
                success: function(res){

                    if (200 == res.status)
                    {
                        //服务端返回times
                        var times = 60;
                        self.countDown(that, times);
                    }
                    else
                    {
                        if (typeof $.fancybox != 'undefined')
                        {
                            Public.tipMsg(res.msg);
                        }
                        else
                        {
                            alert(res.msg);
                        }
                    }
                }
            });

        });
    }
};


function formatMoney(value, prefix, endfix) {
    var num = new Number(value);

    if(typeof prefix != 'undefined')
    {
    }
    else
    {
        prefix = __("￥")
    }

    if(typeof endfix != 'undefined')
    {
    }
    else
    {
        endfix = ""
    }

    return prefix + num.toFixed(2) +  endfix
}


function payment_met_id(val, opt, row) {
    var r = {
    "1": "余额支付",
    "2": "充值卡支付",
    "3": "积分支付",
    "4": "信用支付",
    "5": "红包支付"
    };
    return r[val];
}

function trade_type_id(val, opt, row) {
    var r = {
    "1201": "购物",
    "1202": "转账",
    "1203": "充值",
    "1204": "提现",
    "1205": "销售",
    "1206": "佣金"
    };
    return r[val];
}


function payment_type_id(val, opt, row) {
    var r = {
    "1301": "货到付款",
    "1302": "在线支付",
    "1303": "白条支付",
    "1304": "现金支付",
    "1305": "线下支付",
    "": null
    };
    return r[val];
}

