# thesis-source-code
毕业论文的配套源码

## DB

driver： MySQL

## Servier

GO

## Client

Browser、vue

## 如何启动

1. 数据库
   1. 创建用户：test、密码：123456
   2. 执行 db/init.sql 文件创建数据库、表、插入初始数据
2. 后端工程
   1. 执行 go run server/main.go 启动程序
3. 前端工程
   1. 进入 client 页面，执行 npm install 安装依赖
   2. 执行 npm run serve 进入开发预览模式
   3. 浏览器输入 http://localhost:8082/
   4. 输入账号 admin、密码 1234