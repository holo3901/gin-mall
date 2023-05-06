# gin-mall
**基于 gin+gorm+mysql读写分离 的一个电子商场**

# 项目的主要功能介绍

- 用户注册登录(JWT-Go鉴权)
- 用户基本信息修改，解绑定邮箱，修改密码
- 商品的发布，浏览等
- 购物车的加入，删除，浏览等
- 订单的创建，删除，支付等
- 地址的增加，删除，修改等
- 各个商品的浏览次数，以及部分种类商品的排行
- 将图片上传到对象存储


# 项目的主要依赖：
Golang V1.20
- gin
- gorm
- mysql
- redis
- jwt-go
- crypto
- logrus
- qiniu-go-sdk

# 项目结构
```
mall/
├── controller
├── dao
|   |——redis
|   |-mysql
|—— logger
|—— logic
├── conf
├── doc
├── middleware
├── model
├── pkg
│  ├── snowflake
│  └── JWT
├── routes
├── setting

```
- controller : 用于定义接口函数
- conf : 用于存储配置文件
- dao : 对持久层进行操作
- dao/redis：放置缓存
- doc : 存放接口文档
- loading : 需要加载的应用
- middleware : 应用中间件
- model : 应用数据库模型
- pkg/e : 封装错误码
- pkg/util : 工具函数
-
- routes : 路由逻辑处理
