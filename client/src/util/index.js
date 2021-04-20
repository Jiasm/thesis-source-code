import moment from "moment"

const statusMap = {
  0: '未开始',
  1: '进行中',
  2: '已完成'
}

const priorityMap = {
  0: '最高优',
  1: '高优',
  2: '正常',
  3: '不紧急',
}

const typeMap = {
  0: '需求',
  1: '缺陷',
  2: '建议'
}

const roleMap = {
  0: '未知',
  1: '管理员',
  2: '参与者',
}

const userStatusMap = {
  0: '已冻结',
  1: '活跃中',
}

export const status = buildSelector(statusMap)
export const priority = buildSelector(priorityMap)
export const taskType = buildSelector(typeMap)
export const role = buildSelector(roleMap)
export const userStatus = buildSelector(userStatusMap)

export function getStatus (status) {
  return statusMap[status] || statusMap[0]
}

export function getPriority (priority) {
  return priorityMap[priority] || priorityMap[0]
}

export function getTaskType (type) {
  return typeMap[type] || typeMap[0]
}

export function getRole (role) {
  return roleMap[role] || roleMap[0]
}

export function getUserStatus (userStatus) {
  return userStatusMap[userStatus] || userStatusMap[0]
}

export function formatDate (date) {
  return moment(date, 'X').format('YYYY-MM-DD HH:mm:SS')
}

export function filterTag (taskTagMap, tagMap, taskId) {
  const tagInfo = taskTagMap[taskId]

  if (!tagInfo || !tagInfo.length) return []

  const tagList = tagInfo.map(tag => tagMap[tag] && tagMap[tag]).filter(i => i)

  return tagList
}

function buildSelector (obj) {
  const data = []
  Object.entries(obj).forEach(([value, label]) => {
    data.push({ value: Number(value), label })
  })

  return data
}