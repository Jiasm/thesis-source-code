<template>
  <div class="container">
    <common-header :closecreate="headerChanged"></common-header>
    <div class="main">
      <common-tree-menu ref="treeMenu"></common-tree-menu>
      <div class="content">
        <el-row>
          <el-page-header @back="goBack" content="数据看板">
          </el-page-header>
        </el-row>
        <el-row>
          <el-col :span="12">
            <div class="grid-content">
              <div id="new-count-chart" :style="{width: '100%', height: '400px'}"></div>
            </div>
          </el-col>
          <el-col :span="12">
            <div class="grid-content">
              <div id="typed-count-chart" :style="{width: '100%', height: '400px'}"></div>
            </div>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <div class="grid-content">
              <div id="todo-count-chart" :style="{width: '100%', height: '400px'}"></div>
            </div>
          </el-col>
          <el-col :span="12">
            <div class="grid-content">
              <div id="done-count-chart" :style="{width: '100%', height: '400px'}"></div>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
import { Loading } from 'element-ui'
import { getNewCountList, getTypedTaskCountList, getTodoCountList, getDoneCountList, getUserMap } from '../lib/api'
import { role, getTaskType, getStatus } from '../util'
import Header from './Header'
import TreeMenu from './TreeMenu'

export default {
  components: { 'common-header': Header, 'common-tree-menu': TreeMenu },
  name: 'Dashboard',
  methods: {
    goBack () {
      this.$router.go(-1)
    },
    renderEcharts () {
      const { userMap, newCountList, typedCountList, doneCountList, todoCountList } = this.$data
      // 绘制图表
      this.newCountChart.setOption({
          title: { text: '新增任务数量' },
          tooltip: {},
          xAxis: {
              data: newCountList.map(row => row.date)
          },
          yAxis: {},
          series: [{
              name: '新增数量',
              type: 'line',
              data: newCountList.map(row => row.new_count)
          }]
      })

      this.typedCountChart.setOption({
          title: {
              text: '任务类型分类',
              left: 'center'
          },
          tooltip: {
              trigger: 'item'
          },
          legend: {
              orient: 'vertical',
              left: 'left',
          },
          series: [
              {
                  name: '任务类型',
                  type: 'pie',
                  radius: '50%',
                  data: typedCountList.map(row => ({ value: row.type_count, name: getTaskType(row.type) })),
                  emphasis: {
                      itemStyle: {
                          shadowBlur: 10,
                          shadowOffsetX: 0,
                          shadowColor: 'rgba(0, 0, 0, 0.5)'
                      }
                  }
              }
          ]
      })

      const todoSeries = [
        {
            name: getStatus(0),
            type: 'bar',
            stack: 'total',
            data: []
        },
        {
            name: getStatus(1),
            type: 'bar',
            stack: 'total',
            data: []
        },
      ]

      const todoUserMap = {}
      todoCountList.forEach(row => {
        if (!todoUserMap[row.executor]) {
          todoUserMap[row.executor] = {
            0: 0,
            1: 1,
          }
        }

        todoUserMap[row.executor][row.status] = row.todo_count
      })

      const todoXList = Object.keys(todoUserMap)

      todoXList.forEach(index => {
        todoSeries[0].data.push(todoUserMap[index][0])
        todoSeries[1].data.push(todoUserMap[index][1])
      })

      // 绘制图表
      this.todoCountChart.setOption({
          legend: {
              data: [getStatus(0), getStatus(1)]
          },
          title: { text: '待完成任务数量' },
          tooltip: {},
          xAxis: {
              data: todoXList.map(uid => userMap[uid].username)
          },
          yAxis: {},
          series: todoSeries
      })

      // 绘制图表
      this.doneCountChart.setOption({
          title: { text: '已完成任务数量' },
          tooltip: {},
          xAxis: {
              data: doneCountList.map(row => userMap[row.executor].username)
          },
          yAxis: {},
          series: [{
              name: '新增数量',
              type: 'bar',
              data: doneCountList.map(row => row.done_count)
          }]
      })
    },
    async getDataAndRender (projectId) {
      const loadingInstance = Loading.service({ fullscreen: true })
      const [newCountList, typedCountList, todoCountList, doneCountList] = await Promise.all([
        getNewCountList(projectId),
        getTypedTaskCountList(projectId),
        getTodoCountList(projectId),
        getDoneCountList(projectId),
      ])

      this.$data.userMap = await getUserMap([].concat(...todoCountList.map(row => row.executor), ...doneCountList.map(row => row.executor), ))
      this.$data.newCountList = newCountList
      this.$data.typedCountList = typedCountList
      this.$data.todoCountList = todoCountList
      this.$data.doneCountList = doneCountList
      this.renderEcharts()
      loadingInstance.close()
    },
    headerChanged () {
      this.$refs.treeMenu.loadData()
    }
  },
  async mounted() {
    this.newCountChart = this.$echarts.init(document.getElementById('new-count-chart'))
    this.typedCountChart = this.$echarts.init(document.getElementById('typed-count-chart'))
    this.doneCountChart = this.$echarts.init(document.getElementById('done-count-chart'))
    this.todoCountChart = this.$echarts.init(document.getElementById('todo-count-chart'))

    if (localStorage.getItem('project_id')) {
      this.getDataAndRender(localStorage.getItem('project_id'))
    }

    window.addEventListener('change-project-id', async (e) => {
      this.getDataAndRender(e.pid)
    })
  },
  newCountChart: null,
  typedCountChart: null,
  doneCountChart: null,
  todoCountChart: null,
  data() {
    return {
      groupId: 0,
      inviteDialogVisible: false,
      newCountList: [],
      typedCountList: [],
      todoCountList: [],
      doneCountList: [],
      userMap: {},
      ruleList: role.filter(row => row.value),
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.el-row {
  margin: 20px 0;
}

.button-row {
  margin-top: 40px;
}

.main {
  display: flex;
}

.content {
  flex: 1;
  padding: 0 50px;
}

.tag {
  margin-right: 10px;
}

.invite {
  text-align: right;
}
</style>
