import moment from 'moment'

const data = [{
  id: 1,
  title: '问卷调查1',
  status: '启用',
  datetime: '2020-04-01 16:00',
  question: '问卷问题1',
  optionsList: ['选项 A 描述', '选项 B 描述', '选项 C 描述']
}, {
  id: 2,
  title: '问卷调查2',
  status: '启用',
  datetime: '2020-03-28 15:00',
  question: '问卷问题2',
  optionsList: ['选项 A 描述', '选项 B 描述']
}, {
  id: 3,
  title: '问卷调查3',
  status: '启用',
  datetime: '2020-03-20 12:00',
  question: '问卷问题3',
  optionsList: ['选项 A 描述', '选项 B 描述']
}, {
  id: 4,
  title: '问卷调查4',
  status: '禁用',
  datetime: '2020-03-09 22:00',
  question: '问卷问题4',
  optionsList: ['选项 A 描述', '选项 B 描述']
}]

let cursor = data[data.length - 1].id

export const getList = () => data.filter(item => item.status === '启用').map(item => {
  return {
    ...item,
    optionsList: item.optionsList.filter(item => item).map(item => ({
      checked: true,
      val: item
    })).concat(new Array(4).fill({ checked: false, val: '' })).slice(0, 4)
  }
})

export const getItem = id => getList().find(item => item.id === id)

const getOrigintItem = id => data.find(item => item.id === id)

export const newItem = ({
  title,
  question,
  optionsList,
}) => {
  const id = ++cursor

  data.push({
    id,
    title,
    question,
    status: '启用',
    datetime: moment().format('YYYY-MM-DD HH:mm'),
    optionsList,
  })

  return id
}

export const changeItem = (id, item) => {
  const originItem = getOrigintItem(id)

  for (const [key, val] of Object.entries(item)) {
    originItem[key] = val
  }
}

export const removeItem = id => {
  data.some(item => {
    if (item.id === id) {
      item.status = '禁用'

      return true
    }
  })
}