import axios from 'axios'
import moment from 'moment'
import { getStatus, getPriority, getTaskType, formatDate, filterTag } from '../util'

export async function login (username, password) {
  const { data } = await axios.post(`/login`, {
    username,
    password
  })

  return data
}

export async function getProjectList () {
  const { data: { data } } = await axios.get(`/project/all`)

  const { created, participant } = data

  const projectList = created.concat(participant.filter(item => !created.find(i => i.id === item.id))).map(row => ({
    id: row.id,
    groupName: row.group_name,
    name: row.name,
    status: row.status
  }))

  const groupMap = {}

  projectList.forEach(row => {
    if (!groupMap[row.groupName]) {
      groupMap[row.groupName] = []
    }

    groupMap[row.groupName].push(row)
  })

  const groupList = []

  Object.entries(groupMap).forEach(([groupName, pList]) => {
    groupList.push({
      groupName,
      list: pList
    })
  })

  return {
    projectList,
    groupList,
  }
}

async function getTaskGroupList (taskGroupIdList) {
  const { data: { data : { list: groupList }} } = await axios.get(`/task/group/list?group_ids=${[...new Set(taskGroupIdList)].join(',')}`)

  return groupList || []
}

async function getUserList (userIdList) {
  const { data: { data : { list: userList }} } = await axios.get(`/user/list?user_ids=${[...new Set(userIdList)].join(',')}`)

  return userList || []
}

async function getTaskTagList (taskIdList) {
  const { data: { data : { list: taskTagList }} } = await axios.get(`/task/tag/list?task_ids=${[...new Set(taskIdList)].join(',')}`)

  return taskTagList || []
}

async function getTagInfoList (tagIdList) {
  const { data: { data : { list: tagList }} } = await axios.get(`/tag/list?tag_ids=${[...new Set(tagIdList)].join(',')}`)

  return tagList || []
}

async function getProjectListById (projectIdList) {
  const { data: { data } } = await axios.get(`/project/list?project_ids=${[...new Set(projectIdList)].join(',')}`)

  return data || []
}

async function getTaskCommentListById (taskId) {
  const { data: { data: taskCommentList } } = await axios.get(`/task/comment/list?task_id=${taskId}`)

  return taskCommentList || []
}

export async function getTaskList (projectId = 1) {
  const { data: { data: { list } } } = await axios.get(`/task/list?task_project_id=${projectId}&size=50`)

  const taskGroupIdList = list.map(item => item.task_group_id)
  const userIdList = list.map(item => item.executor)
  const taskIdList = list.map(item => item.id)

  const groupList = await getTaskGroupList(taskGroupIdList)
  const userList = await getUserList(userIdList)
  const taskTagList = await getTaskTagList(taskIdList)

  groupList.push({ id: 0, title: '未分组' })

  const groupMap = {}
  const userMap = {}
  const taskTagMap = {}

  list.forEach(row => {
    const gId = row.task_group_id

    if (!groupMap[gId]) {
      groupMap[gId] = []
    }

    groupMap[gId].push(row)
  })

  userList.forEach(row => {
    userMap[row.id] = row
  })

  const tagIdList = []

  taskTagList.forEach(row => {
    if (!taskTagMap[row.task_id]) {
      taskTagMap[row.task_id] = []
    }

    taskTagMap[row.task_id].push(row.tag_id)
    
    tagIdList.push(row.tag_id)
  })

  const tagList = await getTagInfoList(tagIdList)

  const tagMap = {}

  tagList.forEach(row => {
    tagMap[row.id] = row
  })

  Object.entries(groupMap).forEach(([groupId, group]) => {
    const parentTaskMap = {}

    group.forEach(row => {
      const pId = row.parent_task_id

      if (!parentTaskMap[pId]) {
        parentTaskMap[pId] = []
      }

      parentTaskMap[pId].push(row)
    })

    groupMap[groupId] = parentTaskMap
  })

  // build list
  const tableList = []

  Object.entries(groupMap).forEach(([groupId, group]) => {
    const groupInfo = groupList.find(row => Number(row.id) === Number(groupId))

    const info = {
      id: `group-${groupInfo.id}`,
      title: groupInfo.title,
      children: []
    }

    // normal list
    group[0].forEach(task => {
      const item = {
        id: task.id,
        title: task.title,
        status: getStatus(task.status),
        priority: getPriority(task.priority),
        executor: userMap[task.executor].username,
        type: getTaskType(task.type),
        expireDate: formatDate(task.expire_date),
        tag: filterTag(taskTagMap, tagMap, task.id),
      }

      if (group[task.id]) {
        item.children = []

        group[task.id].forEach(childTask => {
          const childItem = {
            id: childTask.id,
            title: childTask.title,
            status:getStatus(childTask.status),
            priority: getPriority(childTask.priority),
            executor: userMap[childTask.executor].username,
            type: getTaskType(childTask.type),
            expireDate: formatDate(childTask.expire_date),
            tag: filterTag(taskTagMap, tagMap, childTask.id),
          }

          item.children.push(childItem)
        })
      }

      info.children.push(item)
    })

    tableList.push(info)
  })

  return tableList
}

export async function getTaskDetail (taskId) {
  const { data: { data } } = await axios.get(`/task/detail?task_id=${taskId}`)
  const taskCommentList = await getTaskCommentListById(data.id)

  const projectIdList = [data.task_project_id]
  const taskGroupIdList = [data.task_group_id]
  const taskIdList = [data.id]
  const userIdList = [data.creator, data.executor]

  if (data.child_task && data.child_task.length) {
    data.child_task.forEach(item => {
      taskIdList.push(item.id)
      userIdList.push(item.creator)
      userIdList.push(item.executor)
    })
  }

  if (taskCommentList && taskCommentList.length) {
    taskCommentList.forEach(item => {
      userIdList.push(item.creator)
    })
  }

  const projectList = await getProjectListById(projectIdList)
  const groupList = await getTaskGroupList(taskGroupIdList)
  const userList = await getUserList(userIdList)
  const taskTagList = await getTaskTagList(taskIdList)

  const userMap = {}
  const taskTagMap = {}

  userList.forEach(row => {
    userMap[row.id] = row
  })

  const tagIdList = []

  taskTagList.forEach(row => {
    if (!taskTagMap[row.task_id]) {
      taskTagMap[row.task_id] = []
    }

    taskTagMap[row.task_id].push(row.tag_id)
    
    tagIdList.push(row.tag_id)
  })

  const tagList = await getTagInfoList(tagIdList)

  const tagMap = {}

  tagList.forEach(row => {
    tagMap[row.id] = row
  })

  const info = {
    createdDate: formatDate(data.created_date),
    creator: userMap[data.creator].username,
    desc: data.desc,
    executorText: userMap[data.executor].username,
    executor: data.executor,
    expireDate: formatDate(data.expire_date),
    id: data.id,
    parentTaskId: data.parent_task_id,
    priorityText: getPriority(data.priority),
    priority: data.priority,
    statusText: getStatus(data.status),
    status: data.status,
    taskGroupName: groupList.length ? groupList[0].title : '未分组',
    taskGroupId: groupList.length ? groupList[0].id : 0,
    projectName: projectList[0].name,
    projectId: projectList[0].id,
    title: data.title,
    typeText: getTaskType(data.type),
    type: data.type,
    tags: (taskTagMap[data.id] || []).map(tagId => tagMap[tagId]).filter(i => i),
    childTask: data.child_task ? data.child_task.map(row => {
      return {
        id: row.id,
        title: row.title,
        statusText: getStatus(row.status),
        status: row.status,
        typeText: getTaskType(row.type),
        type: row.type,
        priorityText: getPriority(row.priority),
        priority: row.priority,
        executorText: userMap[row.executor].username,
        executor: row.executor,
        expireDate: formatDate(row.expire_date),
      }
    }) : [],
    commentList: taskCommentList.map(item => ({
      id: item.id,
      text: item.text,
      username: userMap[item.creator].username
    }))
  }

  return info
}

export async function addComment(taskId, commentText) {
  const data = await axios.post('/task/add-comment', {
    task_id: taskId,
    text: commentText,
  })

  return data
}

export async function changeTask (taskItem) {
  await axios.post('/task/update', {
    task_id: String(taskItem.id),
    title: taskItem.title ? taskItem.title : '',
    desc: taskItem.desc ? taskItem.desc : '',
    executor: taskItem.executor ? String(taskItem.executor) : '',
    status: taskItem.status ? String(taskItem.status) : '',
    expire_date: taskItem.expireDate ? String(moment(taskItem.expireDate).unix()) : '',
    task_group_id: taskItem.taskGroupId ? String(taskItem.taskGroupId) : '',
    task_type: taskItem.type ? String(taskItem.type) : '',
    priority: taskItem.priority ? String(taskItem.priority) : '',
    task_project_id: taskItem.projectId
  })
}

export async function newChildTask (taskItem) {
  await axios.post('/child-task', {
    title: taskItem.title,
    desc: taskItem.desc,
    executor: taskItem.executor,
    status: taskItem.status,
    expire_date: moment(taskItem.expireDate).unix(),
    task_project_id: taskItem.taskProjectId,
    task_group_id: taskItem.taskGroupId,
    parent_task_id: taskItem.parentTaskId,
    task_type: taskItem.type,
    priority: taskItem.priority,
  })
}

export async function addNewTag (taskId, tagName) {
  await axios.post('/task/add-tag', {
    task_id: taskId,
    tag: tagName
  })
}

export async function removeTag (taskId, tagId) {
  await axios.post('/task/remove-tag', {
    task_id: taskId,
    tag_id: tagId
  })
}

export async function getGroupList () {
  const { data: { data: { created, participant } } } = await axios.get('/group/list')

  const groupList = created.concat(participant.filter(item => !created.find(i => i.id === item.id))).map(row => ({
    id: row.id,
    name: row.name,
  }))

  return [{ id: 0, name: '未分组' }].concat(groupList)
}

export async function createProject (title, groupId, status = 1) {
  await axios.post('/project', {
    group_id: groupId,
    name: title,
    status
  })
}

export async function createGroup (title, status = 1) {
  await axios.post('/group', {
    name: title,
    status
  })
}

export async function createTask (info) {
  await axios.post('/task', {
    title: info.title,
    desc: info.description,
    executor: info.executor,
    status: info.status,
    expire_date: moment(info.expireDate).unix(),
    task_project_id: info.projectId,
    task_group_id: info.taskGroupId,
    type: info.type,
    priority: info.priority,
  })
}

export async function getAllUserList () {
  const { data: { data: { list } }} = await axios.get('/user/list/all')

  return list
}

export async function getProjectGroupList (projectId) {
  const { data: { data } } = await axios.get(`/project/group/list?project_id=${projectId}`)

  return [{ id: 0, title: '未分组' }].concat(data || [])
}