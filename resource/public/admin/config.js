if ('undefined' == typeof window.SS)
{
    window.SS = {};
}

function getUrlParam(name) {
  var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
  var r = window.location.search.substr(1).match(reg);
  if (r != null) return unescape(r[2]); return null;
}

window.admin_url = typeof window.admin_url != "undefined" ? window.admin_url : getUrlParam('admin_url');
window.base_url = window.admin_url;
window.api_url = typeof window.api_url != "undefined" ? window.api_url : getUrlParam('api_url')


if ('undefined' == typeof window.SYS)
{
  window.SYS = {CONFIG:{}, URL:{}};
}

// 系统参数控制
window.SYSTEM = {
};

// 区分服务支持
window.SYSTEM.categoryInfo = {};
window.SYSTEM.category = {};

SYS.VER = typeof ver != "undefined" ? ver : Math.random();
SYS.DEBUG = 0;

SYS.CONFIG['index_url'] = base_url
SYS.CONFIG['base_url'] = base_url
SYS.CONFIG['im_config'] = api_url + '/front/account/userMessage/getImConfig'
SYS.CONFIG['kefu_config'] = api_url + '/front/account/userMessage/getKefuConfig'
SYS.CONFIG['friend_info_lists'] = api_url + '/mobile/sns/userFriend/getFriendsInfo'
SYS.CONFIG['zonemsg_lists'] = api_url + '/mobile/sns/userZoneMessage/list'
SYS.CONFIG['msg_chat_lists'] = api_url + '/front/account/userMessage/listChatMsg'
SYS.CONFIG['msg_chathistory_lists'] = api_url + '/front/account/userMessage/listChatMsg'
//SYS.CONFIG['msg_chathistory_lists'] = api_url + '/im/iim/chatHistory/list'
SYS.CONFIG['msg_set_read'] = api_url + '/front/account/userMessage/setRead'
SYS.CONFIG['zonemsg_add_msg'] = api_url + '/mobile/sns/zoneMessage/add'
SYS.CONFIG['msg_add'] = api_url + '/front/account/userMessage/add'
SYS.CONFIG['upload'] = api_url + '/front/sys/upload/index'

if ('undefined' == typeof window.SYS.CONFIG.static_url)
{
  SYS.CONFIG['static_url'] = '/static'
}

if ('undefined' == typeof window.SYS.URL.upload)
{
  SYS.URL['upload'] = api_url + '/front/sys/upload/index'
}

var authorization_token = '';

if (localStorage.getItem("ukey")) {
  authorization_token = localStorage.getItem("ukey");
}

SYS.URL['upload'] = SYS.URL['upload'] + '?material_type=image&perm_key=' + encodeURIComponent(authorization_token);

SYS.LD = {lang:'zh_CN', currency_id:86, symbol:'￥', symbol_right:'RMB'};



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
        else if (upload_type == 'tengxun')
        {
            image_url = sprintf('%s?imageMogr2/crop/%sx%s/gravity/center', image_url, w, h);
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
var User_BindConnectModel = {};

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


    SYS.VER   = 'undefined' != typeof SYS.VER ?  SYS.VER : '1.0.28';
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
    SYS.VER_SHOP   = 'undefined' != typeof SYS.VER_SHOP ?  SYS.VER :  '1.0.28';
    SYS.VER_WAP   = 'undefined' != typeof SYS.VER_WAP ?  SYS.VER_WAP :  '1.0.28';
    SYS.VER_ADMIN   = 'undefined' != typeof SYS.VER_ADMIN ?  SYS.VER_ADMIN : '1.0.28';
    SYS.SW_ENABLE = 'undefined' != typeof SYS.SW_ENABLE ?  SYS.SW_ENABLE : 0;
    SYS.HTTPS = 'undefined' != typeof SYS.HTTPS ?  SYS.HTTPS :  0;


    if (window.localStorage) {
        var version = localStorage.getItem("version");

        if (version)
        {
            SYS.VER   = version;
        }
    } else {
    }

    SYS.STATIC_IMAGE_PATH = 'https://static.shopsuite.cn/xcxfile/appicon/';
    SYS.AK_BROWSER = "Yi9XWlwa7sUGSuKGDiPBrS261bMeu6YF";
    SYS.AK_MINIAPP = "uWq8fmHbdvzOqLZlU8QZvbugoDyPFUg6";

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

