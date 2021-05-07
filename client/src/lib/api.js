import axios from 'axios'
import moment from 'moment'
import { getStatus, getPriority, getTaskType, formatDate, filterTag, getUserStatus, getRole } from '../util'

export async function login (username, password) {
  const { data } = await axios.post(`/login`, {
    username,
    password
  })

  return data
}

export async function logout () {
  await axios.post('/logout')
}

export async function getProjectList () {
  const { data: { data } } = await axios.get(`/project/all`)

  let { created, participant } = data

  if (!created) {
    created = []
  }

  if (!participant) {
    participant = []
  }

  const projectList = created.concat(participant.filter(item => !created.find(i => i.id === item.id))).map(row => ({
    id: row.id,
    groupName: row.group_name,
    groupId: row.group_id,
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
      groupId: pList[0].groupId,
      list: pList
    })
  })

  console.log({
    projectList,
    groupList,
  })

  return {
    projectList,
    groupList,
  }
}

function listToMapBuilder (func) {
  return async (...arg) => {
    const listData = await func(...arg)
  
    const mapData = {}

    listData.forEach(row => {
      mapData[row.id] = row
    })
  
    return mapData
  }
}

async function getTaskGroupList (taskGroupIdList) {
  const { data: { data : { list: groupList }} } = await axios.get(`/task/group/list?group_ids=${[...new Set(taskGroupIdList)].join(',')}`)

  return [{ id: 0, title: '未分组' }].concat(groupList || [])
}

// const getTaskGroupMap = listToMapBuilder(getTaskGroupList)

async function getUserList (userIdList) {
  const { data: { data : { list: userList }} } = await axios.get(`/user/list?user_ids=${[...new Set(userIdList)].join(',')}`)

  return userList || []
}

export const getUserMap = listToMapBuilder(getUserList)

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

const getProjectMap = listToMapBuilder(getProjectListById)

export async function getGroupListByGroupId (groupIdList) {
  let { data: { data: groupList } } = await axios.get(`/group/list?group_ids=${[...new Set(groupIdList)].join(',')}`)

  return groupList || []
}

const getGroupListMap = listToMapBuilder(getGroupListByGroupId)

async function getTaskCommentListById (taskId) {
  const { data: { data: taskCommentList } } = await axios.get(`/task/comment/list?task_id=${taskId}`)

  return taskCommentList || []
}

export async function getTaskList (projectId = 1) {
  let { data: { data: { list } } } = await axios.get(`/task/list?task_project_id=${projectId}&size=500`)

  if (!list) {
    return []
  }

  const taskGroupIdList = list.map(item => item.task_group_id)
  const userIdList = list.map(item => item.executor)
  const taskIdList = list.map(item => item.id)

  const [groupList, userList, taskTagList] = await Promise.all([
    getTaskGroupList(taskGroupIdList),
    getUserList(userIdList),
    getTaskTagList(taskIdList),
  ])

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

  const [projectList, groupList, userList, taskTagList] = await Promise.all([
    getProjectListById(projectIdList),
    getTaskGroupList(taskGroupIdList),
    getUserList(userIdList),
    getTaskTagList(taskIdList),
  ])

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
    taskGroupId: data.task_group_id,
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
  let { data: { data: { created, participant } } } = await axios.get('/group/list/all')

  if (!created) {
    created = []
  }

  if (!participant) {
    participant = []
  }

  const groupList = created.concat(participant.filter(item => !created.find(i => i.id === item.id))).map(row => ({
    id: row.id,
    name: row.name,
  }))

  return groupList
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

export async function getAllTaskGroup () {
  const { data: { data: { list } } } = await axios.get(`/task/group/list/all`)

  return [{ id: 0, title: '未分组' }].concat(list || [])
}

export async function createTaskGroup (title, desc) {
  await axios.post('/task-group', {
    title,
    desc,
    status: 1
  })
}

export async function getProjectMemberList (projectId) {
  let { data: { data: { project, group } } } = await axios.get(`/project-member/list?project_id=${projectId}`)

  if (!project) {
    project = []
  }

  if (!group) {
    group = []
  }

  let uidList = []
  let projectList = [projectId]
  let groupList = []

  project.forEach(row => {
    uidList.push(row.uid)
    projectList.push(row.project_id)
  })

  group.forEach(row => {
    uidList.push(row.uid)
    groupList.push(row.group_id)
  })

  const [userMap, projectMap, groupMap] = await Promise.all([
    getUserMap(uidList),
    getProjectMap(projectList),
    getGroupListMap(groupList),
  ])

  group = group.filter(row => !project.find(i => i.uid === row.uid))
  // getUserList
  // getProjectListById
  // getGroupListByGroupId

  const list = []

  project.forEach(row => {
    list.push({
      projectId,
      role: row.role_id,
      roleText: getRole(row.role_id),
      status: row.status,
      statusText: getUserStatus(row.status),
      uid: row.uid,
      username: userMap[row.uid].username,
      createdDate: formatDate(row.created_date),
    })
  })

  group.forEach(row => {
    list.push({
      projectId,
      role: row.role_id,
      roleText: getRole(row.role_id),
      status: row.status,
      statusText: getUserStatus(row.status),
      uid: row.uid,
      username: userMap[row.uid].username,
      createdDate: formatDate(row.created_date),
      desc: `继承自 ${groupMap[projectMap[projectId].group_id].name}`,
    })
  })

  return list
}

export async function getGroupMemberList (groupId) {
  let { data: { data: listData} } = await axios.get(`/group-member/list?group_id=${groupId}`)

  if (!listData) {
    listData = []
  }

  let uidList = []
  let groupList = []
  
  listData.forEach(row => {
    uidList.push(row.uid)
    groupList.push(row.group_id)
  })

  const userMap = await getUserMap(uidList)

  const list = []

  listData.forEach(row => {
    list.push({
      groupId,
      role: row.role_id,
      roleText: getRole(row.role_id),
      status: row.status,
      statusText: getUserStatus(row.status),
      uid: row.uid,
      username: userMap[row.uid].username,
      createdDate: formatDate(row.created_date),
    })
  })

  return list
}

export async function addMemberToProject (projectId, uid, roleId, status = 1) {
  await axios.post('/project-member/add', {
    project_id: projectId,
    uid,
    role_id: roleId,
    status,
  })
}

export async function changeMemberRoleToProject (projectId, uid, roleId) {
  await axios.post('/project-member/change-role', {
    project_id: Number(projectId),
    uid,
    role_id: roleId,
  })
}

export async function removeMemberToProject (projectId, uid) {
  await axios.post('/project-member/remove', {
    project_id: Number(projectId),
    uid,
  })
}

export async function activeMemberToProject (projectId, uid) {
  await axios.post('/group-member/active', {
    project_id: Number(projectId),
    uid,
  })
}

export async function addMemberToGroup (groupId, uid, roleId, status = 1) {
  await axios.post('/group-member/add', {
    group_id: groupId,
    uid,
    role_id: roleId,
    status,
  })
}

export async function changeMemberRoleToGroup (groupId, uid, roleId) {
  await axios.post('/group-member/change-role', {
    group_id: Number(groupId),
    uid,
    role_id: roleId,
  })
}

export async function removeMemberToGroup (groupId, uid) {
  await axios.post('/group-member/remove', {
    group_id: Number(groupId),
    uid,
  })
}

export async function activeMemberToGroup (groupId, uid) {
  await axios.post('/group-member/active', {
    group_id: Number(groupId),
    uid,
  })
}

export async function getNewCountList (projectId) {
  const { data: { data } } = await axios.get(`/info/new-count-list?project_id=${projectId}`)

  return data || []
}

export async function getTypedTaskCountList (projectId) {
  const { data: { data } } = await axios.get(`/info/typed-count-list?project_id=${projectId}`)

  return data || []
}

export async function getTodoCountList (projectId) {
  const { data: { data } } = await axios.get(`/info/todo-count-list?project_id=${projectId}`)

  return data || []
}

export async function getDoneCountList (projectId) {
  const { data: { data } } = await axios.get(`/info/done-count-list?project_id=${projectId}`)

  return data || []
}
