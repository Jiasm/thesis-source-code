<template>
  <div class="container">
    <common-header></common-header>
    <div class="main">
      <common-tree-menu></common-tree-menu>
      <common-task-detail :visible.sync="dialogTableVisible" :task-id="taskId" :close="close" :view-state=true />
      <div class="content">
        <el-row>
          <el-page-header @back="goBack" content="任务列表">
          </el-page-header>
        </el-row>
        <el-row>
          <el-table
            :data="tableData"
            style="width: 100%;margin-bottom: 20px;"
            row-key="id"
            border
            default-expand-all
            @row-click="openTaskDetail"
            :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
          >
            <el-table-column
              prop="title"
              label="任务名"
            >
            </el-table-column>
            <el-table-column
              prop="status"
              label="任务状态"
              width="150"
            >
            </el-table-column>
            <el-table-column
              prop="priority"
              label="优先级"
              width="150"
            >
            </el-table-column>
            <el-table-column
              prop="executor"
              label="执行人"
              width="150"
            >
            </el-table-column>
            <el-table-column
              prop="type"
              label="任务类型"
              width="150"
            >
            </el-table-column>
            <el-table-column
              prop="expireDate"
              label="截止日期"
              width="200"  
            >
            </el-table-column>
            <el-table-column
              label="标签"
              width="300"
            >
              <template slot-scope="scope">
                <el-tag
                  class="tag"
                  size="small"
                  v-for="tag in scope.row.tag"
                  :key="tag.id"
                >
                  {{tag.text}}
                </el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
import { getTaskList } from '../lib/api'
import Header from './Header'
import TaskDetail from './TaskDetail'
import TreeMenu from './TreeMenu'

export default {
  components: { 'common-header': Header, 'common-tree-menu': TreeMenu, 'common-task-detail': TaskDetail },
  name: 'List',
  methods: {
    openTaskDetail(row) {
      if (/^\d+$/.test(row.id)) {
        this.$data.taskId = row.id
        this.$data.dialogTableVisible = true
        this.$forceUpdate()
      }
    },
    async close () {
      this.$data.dialogTableVisible = false
      if (localStorage.getItem('project_id')) {
        this.$data.tableData = await getTaskList(localStorage.getItem('project_id'))
      }
    },
    goBack () {
      this.$router.go(-1)
    }
  },
  async mounted() {
    if (localStorage.getItem('project_id')) {
      this.$data.tableData = await getTaskList(localStorage.getItem('project_id'))
    }

    window.addEventListener('change-project-id', async (e) => {
      this.$data.tableData = await getTaskList(e.pid)  
    })
  },
  data() {
    return {
      dialogTableVisible: false,
      taskId: 0,
      tableData: [],
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
</style>
