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
Plus开发，在微信公众号、小程序、H5移动端都能使用，独立部署，二开很方便，能满足企业新零售、分销推广、拼团、砍价、秒杀等多种经营需求，自用、做二开项目都很合适。

ShopSuite开源商城Go版本：https://www.shopsuite.cn/golershop

系统支持独立部署、二开方便，方便扩展微服务版本，适用于企业新零售、分销、拼团、砍价，秒杀等各种业务需求。

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

![](https://docs.shopsuite.cn/modulithshop/demo_qrcode.png "ShopSuite 扫描体验")

账号：demo <br />
密码：shopsuite.cn

Java 版本演示站：
- 移动端： https://demo.modulithshop.cn/h5
- PC端： https://demo.modulithshop.cn/
- 后台：https://demo.modulithshop.cn/admin

Golang 版本演示站：
- 移动端： https://demo.golershop.cn/h5
- PC端： https://demo.golershop.cn/
- 后台：https://demo.golershop.cn/admin


PHP 版本演示站：
- 移动端： https://demo.kuteshop.cn/h5
- PC端： https://demo.kuteshop.cn/
- 后台：https://demo.kuteshop.cn/admin



Java SpringCloud 微服务版本演示站：
- 移动端： https://demo.mallsuite.cn/h5
- PC端： https://demo.mallsuite.cn/
- 平台后台：https://demo.mallsuite.cn/admin
- 商户后台：https://demo.mallsuite.cn/admin

~~~
    PC端（请在电脑端访问）： https://demo.mallsuite.cn
    H5端(微商城)：  https://demo.mallsuite.cn/h5
    小程序：Mallsuite（微信搜索小程序）
    
    
    (请在电脑端访问后台管理系统)
    平台后台： https://demo.mallsuite.cn/admin
    账号：demoadmin
    密码：111111
    
    商户后台： https://demo.mallsuite.cn/admin
    账号：seller
    密码：111111
    
    
    供应商后台： https://demo.mallsuite.cn/admin
    账号：gys2
    密码：111111
~~~



PHP 多商户版本演示站：
- 移动端： https://test.suteshop.com/h5
- PC端： https://test.suteshop.com
- 管理后台：https://test.suteshop.com/admin
- 商户后台：https://test.suteshop.com/admin


~~~
    演示站：
    PC端（请在电脑端访问）： https://test.suteshop.com
    H5端(微商城)：  https://test.suteshop.com/h5
    小程序：随商商城系统（微信搜索小程序）
    
    
    (请在电脑端访问后台管理系统)
    平台后台： https://test.suteshop.com/admin.php
    账号：demoadmin
    密码：111111
    
    商户后台： https://test.suteshop.com/admin.php
    账号：seller
    密码：111111
    
    
    供应商后台： https://test.suteshop.com/admin.php
    账号：供应商
    密码：111111
~~~

[想了解ShopSuite开源商城系统Go版整体框架，你可以戳这里快速掌握！](https://docs.shopsuite.cn/golershop/)

---

### 🎬 源码下载：
- 移动端/uniapp前端源码 [shopsuite-mobile](https://github.com/shsuishang/shopsuite-mobile)
- 管理端VUE3前端源码 [shopsuite-admin](https://github.com/shsuishang/shopsuite-admin)
- Java SpringBoot3 [Modulithshop-v3](https://github.com/shsuishang/modulithshop-v3-java)
- Java SpringBoot2 [Modulithshop](https://github.com/shsuishang/modulithshop)
- PHP [Kuteshop](https://github.com/shsuishang/kuteshop)
- Golang Gf框架 [Golershop-Goframe](https://github.com/shsuishang/golershop)
- Golang Gin框架 [Golershop-Gin](https://github.com/shsuishang/golershop-gin-mall)

---

### 💟 UI界面

#### 📱 移动端截图


![商城首页](https://docs.shopsuite.cn/modulithshop/intro/32398547-2363-48ca-a25c-818d28507df9.png "自定义装修商城首页")
![分类页](https://docs.shopsuite.cn/modulithshop/intro/e1f71dba-8a08-404b-b876-f635845d075e.png "三级分类页")
![分类商品页](https://docs.shopsuite.cn/modulithshop/intro/e2026e33-0e24-4d53-a818-fcebb4b9ab72.png "一二级分类展示商品")
![商品列表页](https://docs.shopsuite.cn/modulithshop/intro/daee2998-ae85-4849-970e-a111e45dfc2b.png "商品列表页")
![商品列表页](https://docs.shopsuite.cn/modulithshop/intro/b438933f-447c-41bf-97f9-43c8a10f1483.png "商品列表搜索过滤")
![商品搜索页](https://docs.shopsuite.cn/modulithshop/intro/3e5b3c3f-627c-485e-909e-b25fc3e87596.png "商品搜索页")
![商品详情页](https://docs.shopsuite.cn/modulithshop/intro/4b6ce8b4-2dc0-45ba-8c78-c6a1c5e39b4e.png "商品详情页")
![规格选择页](https://docs.shopsuite.cn/modulithshop/intro/3f4cf9ea-7564-449d-a029-fd66d536e1fc.png "商品规格选择-零售模式")
![规格选择页](https://docs.shopsuite.cn/modulithshop/intro/070c4e1a-ea6c-4c43-8453-cab0077f3eb1.png "商品规格选择-B2B批发模式")
![购物车页](https://docs.shopsuite.cn/modulithshop/intro/e3e4f9b7-3c01-4ed6-bcd2-2751865ea40b.png "购物车页")
![结算页](https://docs.shopsuite.cn/modulithshop/intro/37dada65-f291-4828-86bf-2d0892c06371.png "结算页")
![订单列表页](https://docs.shopsuite.cn/modulithshop/intro/48ab1bc2-7223-4833-acd8-4dd67fc99bf3.png "订单列表页")
![活动弹窗](https://docs.shopsuite.cn/modulithshop/intro/4d91d8e1-cff1-4b27-a243-48f279d6ee45.png "活动弹窗")
![秒杀](https://docs.shopsuite.cn/modulithshop/intro/53e718dd-b0b7-4677-b3cd-3b0b4efc6ae3.png "秒杀活动页")
![优惠券领取页](https://docs.shopsuite.cn/modulithshop/intro/a9c0e2d3-2d07-4f1d-be08-e8c0e5f74ef2.png "优惠券领取页")
![拼团页](https://docs.shopsuite.cn/modulithshop/intro/f8a7fc61-1f91-4449-9519-7a7ec2117ec3.png "拼团页")
![砍价](https://docs.shopsuite.cn/modulithshop/intro/28908e03-0e7e-417c-b207-e7b73e64b23c.png "砍价")
![组合套餐](https://docs.shopsuite.cn/modulithshop/intro/05cf565b-bc5b-42c9-a316-800c83fcf679.png "组合套餐")




#### 🔹 PC端截图
![PC首页](https://docs.shopsuite.cn/modulithshop/intro/pc/index.jpg "PC首页")
![列表页](https://docs.shopsuite.cn/modulithshop/intro/pc/list.jpg "列表页")
![详情页](https://docs.shopsuite.cn/modulithshop/intro/pc/detail.jpg "详情页")
![购物车](https://docs.shopsuite.cn/modulithshop/intro/pc/cart.jpg "购物车")
![结算页](https://docs.shopsuite.cn/modulithshop/intro/pc/checkout.jpg "结算页")
![用户中心](https://docs.shopsuite.cn/modulithshop/intro/pc/center.jpg "用户中心")

#### 🔹 管理端截图

![运营首页](https://docs.shopsuite.cn/modulithshop/intro/admin/analytics.png "运营首页")
![列表页](https://docs.shopsuite.cn/modulithshop/intro/admin/diy.png "首页自定义装修")
![商品管理](https://docs.shopsuite.cn/modulithshop/intro/admin/product.png "商品管理")
![活动管理](https://docs.shopsuite.cn/modulithshop/intro/admin/activity.png "活动管理")
![订单管理](https://docs.shopsuite.cn/modulithshop/intro/admin/order.png "订单管理")
![素材管理](https://docs.shopsuite.cn/modulithshop/intro/admin/media.png "素材管理")


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
