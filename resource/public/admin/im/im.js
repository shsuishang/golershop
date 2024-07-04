//扩展对象方法
$.extend({
  //为对象新增ajaxPost方法
  request: function (ajaxOpts) {
    var opts = {
      type: "POST",
      dataType: "json",
      timeout: 50000,
      loading: true,
      data: {
        typ: 'json'},
      success: function (data, status) {
      },

      error: function (err, status) {
        Public.tipMsg(__('操作无法成功，请稍后重试！'));
      },

      beforeSend: function (request) {
        // 规范写法 不可随意自定义
        if (localStorage.getItem("ukey")) {
          request.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("ukey"));
        }
      }
    };

    $.extend(true, opts, ajaxOpts);

    if (opts.loading) {    //loading
      //var $this = $(this);
      var loading;
      //var myTimer;
      //var preventTooFast = 'ui-btn-dis';

      $.extend(true, opts, {
        beforeSend: function (request) {
          // 规范写法 不可随意自定义
          if (localStorage.getItem("ukey")) {
            request.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("ukey"));
          }
        },
        complete: function () {
          //loading.close();
        }
      });

      /*
       if ($this.hasClass(preventTooFast))
       {
       return;
       }
       */
    }


    var successCallback = opts.success;
    var errorCallback = opts.error;

    opts.success = function (data, status) {
      /*if(data.status != 200){
       var defaultPage = Public.getDefaultPage();
       var msg = data.msg || '出错了=. =||| ,请点击这里拷贝错误信息 :)';
       var errorStr = msg;
       if(data.data.error){
       var errorStr = '<a id="myText" href="javascript:window.clipboardData.setData("Text",data.error);alert("详细信息已经复制到剪切板，请拷贝给管理员！");"'+msg+'</a>'
       }
       defaultPage.Public.tips({type:1, content:errorStr});
       return;
       }*/
      successCallback && successCallback(data, status);
    }

    opts.error = function (err, ms) {
      var content = __('服务端响应错误！')
      if (ms === 'timeout') {
        content = __('请求超时！');
      }

      Public.tipMsg(content);
      errorCallback && errorCallback(err);
    }

    $.ajax(opts);
  }
});


$.extend({
  //为对象新增ajaxPost方法
  send: function (url, data, callback, type) {
    // shift arguments if data argument was omitted
    if (jQuery.isFunction(data)) {
      type = type || callback;
      callback = data;
      data = undefined;
    }

    // The url can be an options object (which then must have .url)
    return $.request(jQuery.extend({
      url: url,
      type: 'GET',
      dataType: type,
      data: data,
      loading: false,
      success: callback
    }, jQuery.isPlainObject(url) && url));
  }
});

var socket_connect = 0;//连接状态
var socket_handle = null;

var vid = 10001; //固定用户调试使用

if (localStorage.getItem("ukey")) {

} else {
  //临时用户
  vid = parseInt(Math.random() * (99999999999 - 10000000000 + 1) + 10000000000, 10);
}

window.initIm = function () {
  //$.send(sprintf("%s/account.php?ctl=%s&met=%s&typ=json", SYS.CONFIG.base_url, 'Index', 'getConfig'), {uid:vid}, function (res)
  $.send(SYS.CONFIG.im_config, {uid: vid}, function (res) {
    if (res.status == 200 && res.data.im_chat) {

      var userInfo = res.data.user_info;
      var resourceSiteUrl = res.data.resource_site_url;
      var suid = res.data.suid;
      var puid = res.data.puid;

      resourceSiteUrl = window.admin_url

      function addCSS(url) {
        var link = document.createElement('link');
        link.type = 'text/css';
        link.rel = 'stylesheet';
        link.href = url;
        document.getElementsByTagName("head")[0].appendChild(link);
      }

      addCSS(resourceSiteUrl + '/im/css/layui.css?v=' + SYS.VER);

      $.getScript(resourceSiteUrl + '/im/layui.js?v=' + SYS.VER, function () {
        layui.config({
          base: resourceSiteUrl + '/im/', //你存放新模块的目录，注意，不是layui的模块目录
          dir: resourceSiteUrl + '/im/'
        }); //加载入口

        if (!/^http(s*):\/\//.test(location.href)) {
          alert('请部署到localhost上查看该演示');
        }


        layui.use('layim', function (layim) {

          //演示自动回复
          var autoReplay = [
            '您好，我现在有事不在，一会再和您联系。',
            '你没发错吧？face[微笑] ',
            '洗澡中，请勿打扰，偷窥请购票，个体四十，团体八折，订票电话：一般人我不告诉他！face[哈哈] ',
            '你好，我是主人的美女秘书，有什么事就跟我说吧，等他回来我会转告他的。face[心] face[心] face[心] ',
            'face[威武] face[威武] face[威武] face[威武] ',
            '<（@￣︶￣@）>',
            '你要和我说话？你真的要和我说话？你确定自己想说吗？你一定非说不可吗？那你说吧，这是自动回复。',
            'face[黑线]  你慢慢说，别急……',
            '(*^__^*) face[嘻嘻] ，是贤心吗？'
          ];

          var config_data = {

            //初始化接口
            init: {
              //url: sprintf("%s/account.php?mdu=sns&ctl=%s&met=%s&typ=json", SYS.CONFIG.base_url, 'User_Friend', 'getFriendsInfo'),
              //url: SYS.CONFIG.friend_info_lists,
              url: SYS.CONFIG.im_config,
              data: {
              }
            }

            //或采用以下方式初始化接口
            /*
             ,init: {
             mine: {
             "username": "LayIM体验者" //我的昵称
             ,"id": "100000123" //我的ID
             ,"status": "online" //在线状态 online：在线、hide：隐身
             ,"remark": "在深邃的编码世界，做一枚轻盈的纸飞机" //我的签名
             ,"avatar": "a.jpg" //我的头像
             }
             ,friend: []
             ,group: []
             }
             */


            //查看群员接口
            , members: {
              url: sprintf("%s/account.php?mdu=sns&ctl=%s&met=%s&typ=json", SYS.CONFIG.base_url, 'User_Zone', 'getMembers')
              , data: {}
            }

            //上传图片接口
            , uploadImage: {
              url: SYS.CONFIG.upload //（返回的数据格式见下文）
              , type: 'post' //默认post
            }

            //上传文件接口
            , uploadFile: {
              url: SYS.CONFIG.upload  //（返回的数据格式见下文）
              , type: 'post' //默认post
            }

            , isAudio: true //开启聊天工具栏音频
            , isVideo: true //开启聊天工具栏视频

            //扩展工具栏
            , tool: [{
              alias: 'code'
              , title: '代码'
              , icon: '&#xe64e;'
            }]

            //,brief: true //是否简约模式（若开启则不显示主面板）

            , title: 'WebIM' //自定义主面板最小化时的标题
            , right: '0px' //主面板相对浏览器右侧距离
            //,minRight: '90px' //聊天面板最小化时相对浏览器右侧距离
            , initSkin: '5.jpg' //1-5 设置初始背景
            //,skin: ['aaa.jpg'] //新增皮肤
            //,isfriend: false //是否开启好友
            , isgroup: true //是否开启群组
            , min: true //是否始终最小化主面板，默认false
            , notice: true //是否开启桌面消息提醒，默认false
            //,voice: false //声音提醒，默认开启，声音文件为：default.mp3

            , msgbox: layui.cache.dir + 'css/modules/layim/html/msgbox.html' //消息盒子页面地址，若不开启，剔除该项即可
            , find: layui.cache.dir + 'css/modules/layim/html/find.html' //发现页面地址，若不开启，剔除该项即可
            , chatLog: layui.cache.dir + 'css/modules/layim/html/chatlog.html' //聊天记录页面地址，若不开启，剔除该项即可

          };


          if (typeof ajaxCart != 'undefined') {
            config_data['right'] = '36px';
          }

          if (userInfo) {
          } else {
            $.cookie('vid', vid);

            config_data = {
              init: {
                //配置客户信息
                mine: {
                  "username": "访客" //我的昵称
                  , "id": puid || vid //我的ID
                  , "user_id": vid //我的ID
                  , "status": "online" //在线状态 online：在线、hide：隐身
                  , "remark": "在深邃的编码世界，做一枚轻盈的纸飞机" //我的签名
                  , "avatar": resourceSiteUrl + '/im/images/user_no_avatar.png' //我的头像
                }
              }
              , title: 'WebIM' //自定义主面板最小化时的标题

              //开启客服模式
              , brief: true
              , isfriend: false //是否开启好友
              , isgroup: false //是否开启群组
              , min: true //是否始终最小化主面板，默认false
              , notice: true //是否开启桌面消息提醒，默认false
            }
          }

          //基础配置
          layim.config(config_data);

          /*
           layim.chat({
           name: '在线客服-小苍'
           ,type: 'kefu'
           ,avatar: 'http://tva3.sinaimg.cn/crop.0.0.180.180.180/7f5f6861jw1e8qgp5bmzyj2050050aa8.jpg'
           ,id: -1
           });
           layim.chat({
           name: '在线客服-心心'
           ,type: 'kefu'
           ,avatar: 'http://tva1.sinaimg.cn/crop.219.144.555.555.180/0068iARejw8esk724mra6j30rs0rstap.jpg'
           ,id: -2
           });
           layim.setChatMin();*/

          //监听在线状态的切换事件
          layim.on('online', function (data) {
            //console.log(data);
            if (data == "online") {
              setonline(); //用户上线
            } else {
              sethide();//用户离线
            }
          });

          //监听签名修改
          layim.on('sign', function (value) {
            console.log(value);
            $.send(sprintf("%s/account.php?ctl=%s&met=%s&typ=json", SYS.CONFIG.base_url, 'User_Account', 'sign'), {'user_sign': value}, function (res) {
            });
          });

          //监听自定义工具栏点击，以添加代码为例
          layim.on('tool(code)', function (insert) {
            layer.prompt({
              title: __('插入代码')
              , formType: 2
              , shade: 0
            }, function (text, index) {
              layer.close(index);
              insert('[pre class=layui-code]' + text + '[/pre]'); //将内容插入到编辑器
            });
          });

          //监听layim建立就绪
          layim.on('ready', function (res) {

            //console.log(res.mine);
            /*
             layim.msgbox(5); //模拟消息盒子有新消息，实际使用时，一般是动态获得

             //添加好友（如果检测到该socket）

             layim.addList({
             type: 'group'
             ,avatar: "http://tva3.sinaimg.cn/crop.64.106.361.361.50/7181dbb3jw8evfbtem8edj20ci0dpq3a.jpg"
             ,groupname: 'Angular开发'
             ,id: "12333333"
             ,members: 0
             });
             layim.addList({
             type: 'friend'
             ,avatar: "http://tp2.sinaimg.cn/2386568184/180/40050524279/0"
             ,username: '冲田杏梨'
             ,groupid: 0
             ,id: "1233333312121212"
             ,remark: "本人冲田杏梨将结束AV女优的工作"
             });

             */

            setTimeout(function () {
              //接受消息（如果检测到该socket）
              /*
               layim.getMessage({
               username: "Hi"
               ,avatar: "http://qzapp.qlogo.cn/qzapp/100280987/56ADC83E78CEC046F8DF2C5D0DD63CDE/100"
               ,id: "10000111"
               ,type: "friend"
               ,content: "临时："+ new Date().getTime()
               });

               layim.getMessage({
               username: "贤心"
               ,avatar: "http://tp1.sinaimg.cn/1571889140/180/40030060651/1"
               ,id: "1"
               ,type: "friend"
               ,content: "嗨，你好！欢迎体验LayIM。演示标记："+ new Date().getTime()
               });
               */

            }, 3000);
          });

          //监听发送消息
          layim.on('sendMessage', function (data) {
            var To = data.to;//对方的信息
            //跨平台用户修正
            //data['to']['id'] = suid + '-' + data['to']['id'];
            //console.log(data);
            //发送消息
            if (socket_connect === 1) {
              if (userInfo) {
                if (To.type === 'group') {
                  var url = sprintf("%s/account.php?mdu=sns&ctl=%s&met=%s&typ=json", SYS.CONFIG.base_url, 'User_Zone', 'addMessage');

                  url = SYS.CONFIG.zonemsg_add_msg;
                } else {
                  var url = sprintf("%s/account.php?mdu=sns&ctl=%s&met=%s&typ=json", SYS.CONFIG.base_url, 'User_Message', 'add');
                  url = SYS.CONFIG.msg_add;
                }

                console.info(data);

                var params = {};
                params["user_other_id"] = data.to.friend_id
                params["message_content"] = data.mine.content
                params["item_id"] = 0
                params["length"] = 0
                params["w"] = 0
                params["h"] = 0
                //params["type"] = text



                $.request({
                  type: 'post',
                  //url: SYS.URL.user.msg_add,
                  url: url,
                  data: params,
                  dataType: 'json',
                  success: function (result) {
                    if (result.status == 200) {
                      var msgData = result.data;
                      data['mine']['message_id'] = msgData.message_other_id;

                      /*
                       var e = socket.createEvent('send_msg');
                       e.data = msgData;
                       socket.dispatchEvent(e);
                       */

                      //向服务器发送数据
                      var text = JSON.stringify(data);
                      socket_handle.send(text);
                    } else {
                      return false;
                    }
                  }
                });
              } else {
                //向服务器发送数据
                var text = JSON.stringify(data);
                socket_handle.send(text);

              }
            }

            /*
             if(To.type === 'friend'){
             layim.setChatStatus('<span style="color:#FF5722;">对方正在输入。。。</span>');
             }

             //演示自动回复
             setTimeout(function(){
             var obj = {};
             if(To.type === 'group'){
             obj = {
             username: '模拟群员'+(Math.random()*100|0)
             ,avatar: layui.cache.dir + 'images/face/'+ (Math.random()*72|0) + '.gif'
             ,id: To.id
             ,type: To.type
             ,content: autoReplay[Math.random()*9|0]
             }
             } else {
             obj = {
             username: To.name
             ,avatar: To.avatar
             ,id: To.id
             ,type: To.type
             ,content: autoReplay[Math.random()*9|0]
             }
             layim.setChatStatus('<span style="color:#FF5722;">在线</span>');
             }
             layim.getMessage(obj);
             }, 1000);
             */
          });

          //监听查看群员
          layim.on('members', function (data) {
            //console.log(data);
          });

          //监听聊天窗口的切换
          layim.on('chatChange', function (res) {
            var type = res.data.type;
            console.log(res.data.id)
            if (type === 'friend') {
              //模拟标注好友状态
              //layim.setChatStatus('<span style="color:#FF5722;">在线</span>');
            } else if (type === 'group') {
              //模拟系统消息
              /*
              layim.getMessage({
                  system: true
                  ,id: res.data.id
                  ,type: "group"
                  ,content: '模拟群员'+(Math.random()*100|0) + '加入群聊'
              });
              */
            }
          });


          var connentNode = function (data) {
            nodeSiteUrl = data.node_site_url;
            resourceSiteUrl = data.resource_site_url;


            var script = document.createElement("script");
            script.type = "text/javascript";
            script.src = resourceSiteUrl + '/js/reconnecting-websocket.js?v=' + SYS.VER;
            document.body.appendChild(script);
            checkIO();

            function checkIO() {
              setTimeout(function () {
                if (typeof ReconnectingWebSocket === "function") {
                  connect_node();
                } else {
                  checkIO();
                }
              }, 500);
            }

            function connect_node() {
              var connect_url = nodeSiteUrl;
              var member = {};


              socket_connect = 0;//连接状态
              socket_handle = new ReconnectingWebSocket(connect_url);
              socket_handle.onopen = function (event) {
                socket_connect = 1;
                console.info('open');
                console.info(event);
              };

              socket_handle.onconnecting = function (event) {
                console.info(event);
              };

              socket_handle.onmessage = function (event) {
                console.info(event);

                var msg_row = JSON.parse(event.data);


                if (typeof msg_row.msg_type !== "undefined") {
                  if (msg_row.msg_type == 'text') {

                  } else if (msg_row.msg_type == 'voice') {
                    msg_row.content = 'audio[' + msg_row.content + ']';

                  } else if (msg_row.msg_type == 'img') {
                    msg_row.content = 'img[' + msg_row.content + ']';
                  }
                }


                if (msg_row.message_id) {
                  //消息设置为已读
                  var url = sprintf("%s/account.php?mdu=sns&ctl=%s&met=%s&typ=json", SYS.CONFIG.base_url, 'User_Message', 'setRead');
                  url = SYS.CONFIG.msg_set_read;

                  $.request({
                    type: 'post',
                    url: url,
                    data: {message_id: msg_row.message_id},
                    dataType: 'json',
                    success: function (result) {
                    }
                  });

                }

                layim.getMessage(msg_row);
              }
              ;

              socket_handle.onerror = function (event) {
                console.info(event);
              };
              socket_handle.onclose = function (event) {
                console.info(event);
              };

              //开始
              //结束
            }

            // 表情
            function update_chat_msg(msg) {
              if (typeof smilies_array !== "undefined") {
                msg = '' + msg;
                for (var i in smilies_array[1]) {
                  var s = smilies_array[1][i];
                  var re = new RegExp("" + s[1], "g");
                  var smilieimg = '<img title="' + s[6] + '" alt="' + s[6] + '" src="' + resourceSiteUrl + '/images/smilies/' + s[2] + '">';
                  msg = msg.replace(re, smilieimg);
                }
              }
              return msg;
            }

            //发送消息的方法
            function send(mine, To) {
              socket_handle.send(currentUser_id + "_msg_" + To.id + "_msg_" + mine.content + "_msg_" + mine.avatar + "_msg_" + To.type + "_msg_" + currentName + "_msg_NAN");
            };

            //切换在线状态的方法
            function setonline() {
              socket_handle.send("_online_user_" + currentUser_id);
            };

            //切换离线状态的方法
            function sethide() {
              socket_handle.send("_leave_user_" + currentUser_id);
            };

            //更新在线用户信息
            function updateOnlineStatus(arra)//更新在线用户信息
            {
              $("div.layui-layim-main ul.layim-list-friend li ul.layui-layim-list li").each(function () {//状态还原
                if (this.className != 'layim-null') {
                  var span = $(this).find("span:first");
                  var name = span.html();
                  var loginName = this.className.replace("layim-friend", "").trim();
                  //alert(arra+"***"+loginName);
                  if ((',' + arra + ",").indexOf(',' + loginName + ',') >= 0) {
                    if (name.indexOf('(<font color="red">离线</font>)') >= 0) {
                      span.replace('(<font color="red">离线</font>)', '(<font color="green">在线</font>)')
                    } else if (name.indexOf('(<font color="green">在线</font>)') >= 0) {

                    } else {
                      span.html(name + '(<font color="green">在线</font>)');
                    }

                  } else {
                    if (name.indexOf('(<font color="red">离线</font>)') >= 0) {
                    } else if (name.indexOf('(<font color="green">在线</font>)') >= 0) {
                      span.replace('(<font color="green">在线</font>)', '(<font color="red">离线</font>)')
                    } else {
                      span.html(name + '(<font color="red">离线</font>)');
                    }

                  }
                }
              });
            }
          }

          connentNode(res.data);
        });


      }, true);
    }
  });

}

//登录后执行
//initIm();

