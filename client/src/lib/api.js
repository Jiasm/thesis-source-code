import axios from 'axios'

const host = 'http://127.0.0.1:8880'

export async function login (username, password) {
  const { data } = await axios.post(`${host}/login`, {
    username,
    password
  })

  return data
}

export async function getList () {
  const { data } = await axios.get(`${host}/list`)

  return data.map(row => ({
    ...row,
    status: row.status === 1 ? '启用' : '禁用'
  }))
}

export async function removeQuestion (id) {
  const { data } = await axios.post(`${host}/delete`, {
    id
  })

  return data
}

export async function getQuestion (id) {
  const { data } = await axios.get(`${host}/info?id=${id}`)

  return data
}

export async function addQuestion (title, creater) {
  const { data } = await axios.post(`${host}/add`, {
    datetime: new Date().toISOString(),
    creater,
    title
  })

  return data
}

export async function changeQuestion (id, title) {
  const { data } = await axios.post(`${host}/change`, {
    title,
    id
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