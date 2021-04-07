# API 设计文档

## 通用异常处理

Status: !200

```json
{
  "code": "number",
  "msg": "string"
}
```

## 登录 ✅

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

## 退出 ✅

URL: /logout  
METHOD: POST  

## 用户信息 ✅

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

## 项目列表 ✅

URL: /project/list  
METHOD: GET  

```json
{
  "code": "number",
  "data": {
    "created": [
      {
        "id": "number",
        "groupName": "string",
        "name": "string",
        "status": "number"
      }
    ],
    "participant": [
      {
        "id": "number",
        "groupName": "string",
        "name": "string",
        "status": "number"
      }
    ]
  }
}
```

## 组织列表 ✅

URL: /group/list  
METHOD: GET  

```json
{
  "code": "number",
  "data": {
    "created": [
      {
        "id": "number",
        "name": "string",
        "status": "number"
      }
    ],
    "participant": [
      {
        "id": "number",
        "name": "string",
        "status": "number"
      }
    ]
  }
}
```

## 任务列表 ✅

URL: /task/list  
METHOD: GET  

```json
{
  "code": "number",
  "data": {
    "list": [
      {
        "id": 2,
        "title": "测试任务 2",
        "desc": "任务描述",
        "status": 1,
        "creator": 2,
        "executor": 2,
        "created_date": 1616478832,
        "expire_date": 1614439750,
        "task_project_id": 2,
        "parent_task_id": 0,
        "type": 0,
        "priority": 0
      }
    ]
  }
}
```

- [x] 按照任务状态筛选
- [x] 按照任务等级筛选
- [x] 按照执行人筛选
- [x] 按照任务类型筛选
- [x] 按照标签筛选
- [ ] 按照任务名筛选
- [ ] 按照参与者筛选
- [x] 按照创建者筛选

## 我的待办

URL: /todo  
METHOD: GET  

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

## 创建组织 ✅

URL: /group  
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

## 创建项目 ✅

URL: /project  
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

## 创建任务组 ✅

URL: /task-group  
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

## 创建任务 ✅

URL: /task  
METHOD: POST  

```json
{
    "title": "string",
    "desc": "string",
    "executor": "number",
    "status": "number",
    "expire_date": "number",
    "task_project_id": "number",
    "task_group_id": "number",
    "type": "number",
    "priority": "number"
}

{

  "code": "number",
  "data": {
    "task_id": "number",
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

URL: /child-task/list  
METHOD: GET  

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

## 添加任务评论 ✅

URL: /task/add-comment  
METHOD: POST  

```json
{
  "task_id": "number",
  "text": "string"
}

{

  "code": "number",
  "data": {
    "comment_id": "number",
  }
}
```

## 新增子任务 ✅

URL: /child-task  
METHOD: POST  

```json
{
    "title": "string",
    "desc": "string",
    "executor": "number",
    "status": "number",
    "expire_date": "number",
    "task_project_id": "number",
    "task_group_id": "number",
    "parent_task_id": "number",
    "type": "number",
    "priority": "number"
}

{

  "code": "number",
  "data": {
    "task_id": "number",
  }
}
```

## 修改任务状态 ✅

URL: /task/update-status  
METHOD: POST  

```json
{
  "task_id": "number",
  "status": "number"
}

{

  "code": "number"
}
```

## 指定执行人 ✅

URL: /task/set-executor  
METHOD: POST  

```json
{
  "task_id": "number",
  "executor": "number"
}

{

  "code": "number"
}
```

## 添加参与者 ✅

URL: /task/append-member  
METHOD: POST  

```json
{
  "task_id": "number",
  "member": "number"
}

{

  "code": "number"
}
```

## 添加标签 ✅

URL: /task/add-tag  
METHOD: POST  

```json
{
  "task_id": "number",
  "tag": "string"
}

{

  "code": "number",
  "data": {
    "uid": "number",
  }
}
```

## 邀请用户加入组织、项目

URL: /group/invite  
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

URL: /group/remove  
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