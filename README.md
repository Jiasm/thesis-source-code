# thesis-source-code
毕业论文的配套源码

## DB

driver： MySQL

## Server

GO

## Client

Browser、vue

## 环境依赖

- python >= 2.x
- go
- node.js >= 10.x
- mysql >= 5.6

## 如何启动

1. 数据库
   1. 创建用户：pm、密码：123456
   2. 设置账户具有所有数据库操作权限
   3. 执行 db/init.sql 文件创建数据库、表、插入初始数据
2. 后端工程
   1. 执行 sh run-server.sh 启动程序
3. 前端工程
   1. 进入 client 页面，执行 npm install 安装依赖
   2. 执行 npm run dev 进入开发预览模式
   3. 浏览器输入 http://localhost:8080/
   4. 输入账号 user1、密码 123456