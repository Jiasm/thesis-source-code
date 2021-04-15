import axios from 'axios'
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

  return groupList
}

async function getUserList (userIdList) {
  const { data: { data : { list: userList }} } = await axios.get(`/user/list?user_ids=${[...new Set(userIdList)].join(',')}`)

  return userList
}

async function getTaskTagList (taskIdList) {
  const { data: { data : { list: taskTagList }} } = await axios.get(`/task/tag/list?task_ids=${[...new Set(taskIdList)].join(',')}`)

  return taskTagList
}

async function getTagInfoList (tagIdList) {
  const { data: { data : { list: tagList }} } = await axios.get(`/tag/list?tag_ids=${[...new Set(tagIdList)].join(',')}`)

  return tagList
}

async function getProjectListById (projectIdList) {
  const { data: { data } } = await axios.get(`/project/list?project_ids=${[...new Set(projectIdList)].join(',')}`)

  return data
}

export async function getTaskList (projectId = 1) {
  const { data: { data: { list } } } = await axios.get(`/task/list?task_project_id=${projectId}&size=10`)

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
            status:getStatus(task.status),
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
    console.log({ row })
    tagMap[row.id] = row
  })

  const info = {
    createdDate: formatDate(data.created_date),
    creator: userMap[data.creator].username,
    desc: data.desc,
    executor: userMap[data.executor].username,
    expireDate: formatDate(data.expire_date),
    id: data.id,
    parentTaskId: data.parent_task_id,
    priority: getPriority(data.priority),
    status: getStatus(data.status),
    taskGroupName: groupList.length ? groupList[0].title : '未分组',
    projectName: projectList[0].name,
    title: data.title,
    type: getTaskType(data.type),
    tags: taskTagMap[data.id].map(tagId => tagMap[tagId]).filter(i => i),
    childTask: data.child_task ? data.child_task.map(row => {
      return {
        id: row.id,
        title: row.title,
        status: getStatus(row.status),
        priority: getPriority(row.priority),
        executor: userMap[row.executor].username,
        expireDate: formatDate(row.expire_date),
      }
    }) : []
  }

  return info
}

export async function getList () {
  const { data } = await axios.get(`/list`)

  return data.map(row => ({
    ...row,
    status: row.status === 1 ? '启用' : '禁用'
  }))
}

export async function removeQuestion (id) {
  const { data } = await axios.post(`/delete`, {
    id
  })

  return data
}

export async function getQuestion (id) {
  const { data } = await axios.get(`/info?id=${id}`)

  return data
}

export async function addQuestion (title, creater) {
  const { data } = await axios.post(`/add`, {
    datetime: new Date().toISOString(),
    creater: Number(creater),
    title
  })

  return data
}

export async function changeQuestion (id, title) {
  const { data } = await axios.post(`/change`, {
    title,
    id
  })

  return data
}

export async function getAnswer (id) {
  const { data } = await axios.get(`/answer-info?id=${id}`)

  const question = data.answer

  const optionsList = new Array(4)

  const fields = ['option_a', 'option_b', 'option_c', 'option_d']

  fields.forEach((field, index) => {
    if (data[field]) {
      optionsList[index] = { checked: true, val: data[field] }
    } else {
      optionsList[index] = {checked: false, val: ''}
    }
  })

  return {
    id: data.id,
    qid: data.qid,
    question,
    optionsList
  }
}

export async function addAnswer (qid, answer, options) {
  const { data } = await axios.post(`/add-answer`, {
    qid,
    answer,
    option_a: options[0],
    option_b: options[1],
    option_c: options[2],
    option_d: options[3]
  })

  return data
}

export async function changeAnswer (id, answer, options) {
  const { data } = await axios.post(`/change-answer`, {
    id,
    answer,
    option_a: options[0],
    option_b: options[1],
    option_c: options[2],
    option_d: options[3]
  })

  return data
}
// datetime: "2020-04-01 16:00"
// id: 1
// optionsList: Array(4)
// 0: {checked: true, val: "选项 A 描述"}
// 1: {checked: true, val: "选项 B 描述"}
// 2: {checked: true, val: "选项 C 描述"}
// 3: {checked: false, val: ""}
// length: 4
// __proto__: Array(0)
// question: "问卷问题1"
// status: "启用"
// title: "问卷调查1"