# API 设计文档

## 通用异常处理

Status: !200

```json
{
  "code": "number",
  "msg": "string"
}
```

## 登录

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{
  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 退出

URL: /logout  
METHOD: POST  

## 用户信息

URL: /user/:uid  
METHOD: GET  

```json
{
  "code": "number",
  "data": {
    "uid": "number",
    "username": "string",
    "roleId": "number",
    "status": "number"
  }
}
```

## 项目列表

URL: /project/list  
METHOD: GET  

```json
{
  "code": "number",
  "data": {
    "list": [
      {
        "id": "number",
        "creator": {
          "uid": "number",
          "username": "string"
        },
        "groupName": "string",
        "name": "string",
        "status": "number"
      }
    ]
  }
}
```

## 组织列表

URL: /group/list  
METHOD: GET  

```json
{
  "username": "string",
  "password": "string"
}

{
  "code": "number",
  "data": {
    "list": [
      {
        "id": "number",
        "creator": {
          "uid": "number",
          "username": "string"
        },
        "name": "string",
        "status": "number"
      }
    ]
  }
}
```

## 项目-任务列表

URL: /project/:pid/task/list  
METHOD: GET  

```json
{
  "code": "number",
  "data": {
    "list": [
      {
        "groupId": "number",
        "groupTitle": "string",
        "groupDesc": "string",
        "groupCreator": {
          "uid": "number",
          "username": "string"
        },
        "groupStatus": "number",
        "items": [{
          "id": "number",
          "title": "string",
          "desc": "string",
          "creator": {
            "uid": "number",
            "username": "string"
          },
          "expireDate": "number",
          "type": "number",
          "priority": "number",
          "status": "number",
          "childTasks": [{
            "id": "number",
            "title": "string",
            "desc": "string",
            "creator": {
              "uid": "number",
              "username": "string"
            },
            "expireDate": "number",
            "type": "number",
            "priority": "number",
            "status": "number",
          }]
        }]
      }
    ]
  }
}
```

1. 按照任务状态筛选
2. 按照任务等级筛选
3. 按照执行人筛选
4. 按照任务类型筛选
5. 按照标签筛选
6. 按照任务名筛选
7. 按照参与者筛选
8. 按照创建者筛选

## 我的待办

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

1. 筛选执行人为当前用户 & 状态为 doing 的任务

## 创建组织

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 创建项目

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 创建任务

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 获取消息列表

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 获取统计数据

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 任务详情

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 获取子任务列表

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 添加任务评论

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 新增子任务

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 修改任务状态

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 指定执行人

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 添加参与者

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 添加标签

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 邀请用户加入组织、项目

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 从组织、项目中移除用户

URL: /login  
METHOD: POST  

```json
{
  "username": "string",
  "password": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```