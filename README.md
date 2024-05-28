<div align="center" style="margin-top: 10px">
    <img src="https://www.shopsuite.cn/uploads/static/icon-s-default.png" />
</div>
<div align="center">

# ShopSuite开源商城系统Go版

</div>

<div align="center">

[官网](https://www.shopsuite.cn) | [在线体验](https://demo.golershop.cn)
| [帮助文档](https://docs.shopsuite.cn/golershop/)

</div>

---


随商信息技术（上海）有限公司是一家以电商系统开发为核心，为企业提供全面整合的电子商务解决方案和技术服务的技术型软件企业。随商专注电商的技术沉淀和行业积累，专业打造行业领先，功能强大，易用性强，扩展性强产品。助力企业信息化建设，帮助企业经营与互联网应用相结合。

公司团队在电子商务软件和互联网技术领域经验资深，历经多年市场实践已研发出具有自主知识产权的新零售智慧电商生态系统、
B2B2C多用户商城、B2B批发商城、O2O门店及收银系统、S2B2C供应链商城、跨境电商系统，骑手跑腿系统、短视频社交电商及直播系统等电子商务软件系统。并取得多项著作权及发明专利。

ShopSuite开源商城基于Go/PHP/Java + uniapp + Vue3 + ElementUI Plus框架开发的商城系统，Golershop 为 ShopSuite开源商城Go版本.

### 📖 简介：

ShopSuite 开源商城系统Go版，基于Go + Goframe + Vue + Uniapp + Element
Plus开发，在微信公众号、小程序、H5移动端都能使用，代码全开源无加密，独立部署，二开很方便，还支持免费商用，能满足企业新零售、分销推广、拼团、砍价、秒杀等多种经营需求，自用、做二开项目都很合适。

ShopSuite开源商城Go版本：https://gitee.com/suisung/golershop

系统代码全开源无加密，独立部署、二开方便，适用于企业新零售、分销、拼团、砍价，秒杀等各种业务需求。

---

### 💡 系统亮点：

> 1. Goframe 框架开发。  <br>
>2. 【前端】Web PC 管理端 Vue + Element UI + Element Plus。<br>
>3. 【前端】移动端使用 Uniapp 框架，前后端分离开发。<br>
>4. 标准RESTful 接口、标准数据传输，逻辑层次更明确，更多的提高api复用。<br>
>5. 支持Redis队列，降低流量高峰，解除耦合，高可用。<br>
>6. 数据导出，方便个性化分析。<br>
>7. 数据统计分析,使用ECharts图表统计，实现用户、产品、订单、资金等统计分析。<br>
>8. 后台多种角色，多重身份权限管理，权限可以控制到按钮级别的操作。<br>
>9. Vue表单生成控件，拖拽配置表单，减少前端重复表单工作量，提高前端开发效率。<br>

---

### 💻 运行环境及框架：

~~~
1.	移动端uniapp开发框架 可生成H5 公众号 微信小程序
2.	WEB Pc 管理后台使用Vue + Element UI 开发 兼容主流浏览器 ie11+
3.	后台服务 Go Goframe + Mysql + redis
4.	运行环境 linux和windows等都支持,只要有Go环境和对应的数据库 redis
5.	运行条件 Go 1.8
~~~

---

### 🔧 Go项目框架 和 WEB PC 项目运行环境

~~~
1. go 1.18
2. github.com/gogf/gf/v2 v2.2.1
3. github.com/gogf/gf/contrib/drivers/mysql/v2 v2.2.1
4. github.com/dgrijalva/jwt-go v3.2.0+incompatible
5. github.com/go-pay/gopay v1.5.94
6. node 16
7. vue 2.x & 3.x
8. element plus
~~~

---

### 🧭 项目代码包介绍

~~~
1. admin     WEB程序         PC端管理端 VUE3 + ElementUi + Element Plus
2. PC        PC商城         PC买家端 VUE3 + ElementUi + Element Plus
2. app       移动商城         UniApp标准开发(H5 + 微信小程序)
3. Go后端     Api            Go Goframe
4. 接口文档   Api对应的接口文档也可以部署项目后查看
~~~

---

### ⛅ 运行账号要求

- 公众号：服务号（已认证且开通支付功能）
- 小程序（已认证且开通支付功能）
- 微信支付
- 支付宝支付
- 微信开放平台（已认证）
  注：如果单独使用公众号或小程序，只需自备一个账号就可以，则不需要微信开放平台

### ⛅ 运行服务器相关

- 服务器
- 域名 （已完成备案）
- SSL证书
- OSS存储

### 🎬 系统演示：

<div class="pic-list" style="text-align: center;margin-top: 20px;margin-bottom: 20px;">
    <div class="img-div" style="display: inline-block;margin-right: 20px;">
        <img alt="H5微商城" src="https://docs.shopsuite.cn/modulithshop/overview/qrcode_h5.jpg" style="  width: 200px;height: 200px;box-shadow: 0px 0px 8px rgb(26 67 149 / 16%);border-radius: 12px;">
        <p style="color: #000000;font-size: 16px;font-weight: bold;text-align: center;margin-top: 12px;">H5微商城</p>
    </div>
    <div class="img-div" style="display: inline-block;margin-right: 20px;">
        <img alt="小程序" src="https://docs.shopsuite.cn/modulithshop/overview/qrcode_xcx.jpg" style="  width: 200px;height: 200px;box-shadow: 0px 0px 8px rgb(26 67 149 / 16%);border-radius: 12px;">
        <p style="color: #000000;font-size: 16px;font-weight: bold;text-align: center;margin-top: 12px;">微信小程序</p>
    </div>
    <div class="img-div" style="display: inline-block;margin-right: 20px;">
        <img alt="下载随商原生APP" src="https://docs.shopsuite.cn/modulithshop/overview/qrcode_app.jpg" style="  width: 200px;height: 200px;box-shadow: 0px 0px 8px rgb(26 67 149 / 16%);border-radius: 12px;">
        <p style="color: #000000;font-size: 16px;font-weight: bold;text-align: center;margin-top: 12px;">下载APP</p>
    </div>
</div>

移动端： https://demo.modulithshop.cn/h5

PC端： https://demo.modulithshop.cn/

后台演示地址：https://demo.modulithshop.cn/admin
账号：demoadmin  
密码：111111

[想了解ShopSuite开源商城系统Go版整体框架，你可以戳这里快速掌握！](https://docs.shopsuite.cn/golershop/)

---

### 📃 系统资料

需要系统文档的朋友看过来，安装文档、产品介绍、技术文档...你想要的我都有！
[https://docs.shopsuite.cn/golershop/](https://docs.shopsuite.cn/golershop/)


---

### 📞 技术交流

跟着官方，不迷路！欢迎扫码加入ShopSuite 开源项目群，一手消息及资源，尽在掌握！<br>

![](https://docs.shopsuite.cn/golershop/contact_golang_qr.png "ShopSuite Golang 技术交流微信群")


---

### 💌 特别鸣谢

核心开发团队

产品：海茵

技术：水塘，方方，小班，赛赛，贺龙

UI：天空

测试：小美丽

---

### 🔔 使用须知

1. 允许用于个人学习、毕业设计、教学案例、公益事业使用;<br>
2. 非商业授权必须保留版权信息，请自觉遵守;<br>
3. 禁止将本项目的代码和资源进行任何形式的出售，产生的一切任何后果责任由侵权者自负。<br>

---

### 👑 版权信息

本项目包含的第三方源码和二进制文件之版权信息另行标注。<br>
版权所有Copyright © 2018-2028 by ShopSuite (https://www.shopsuite.cn)<br>
All rights reserved。<br>
ShopSuite® 商标和著作权所有者为随商信息技术（上海）有限公司。<br>

---
