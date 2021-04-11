<template>
  <div class="container">
    <common-header></common-header>
    <div class="main">
      <common-tree-menu></common-tree-menu>
      <div class="content">
        <el-table
          :data="tableData"
          style="width: 100%;margin-bottom: 20px;"
          row-key="id"
          border
          default-expand-all
          :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        >
          <el-table-column
            prop="title"
            label="任务名"
            sortable
            width="180">
          </el-table-column>
          <el-table-column
            prop="status"
            label="任务状态"
            sortable
            width="120">
          </el-table-column>
          <el-table-column
            prop="priority"
            sortable
            label="优先级">
          </el-table-column>
          <el-table-column
            prop="executor"
            sortable
            label="执行人">
          </el-table-column>
          <el-table-column
            prop="type"
            sortable
            label="任务类型">
          </el-table-column>
          <el-table-column
            prop="expireDate"
            sortable
            label="截止日期">
          </el-table-column>
          <el-table-column
            label="标签">
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
      </div>
    </div>
  </div>
</template>

<script>
import { getTaskList, removeQuestion } from '../lib/api'
import Header from './Header'
// import TaskDetail from './TaskDetail'
import TreeMenu from './TreeMenu'

export default {
  components: { 'common-header': Header, 'common-tree-menu': TreeMenu },
  name: 'List',
  methods: {
    remove(row) {
      const id = Number(row.id)
      this.$confirm('此操作将禁用问卷，后续用户将不能再进行填写, 是否确认?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        const res = await removeQuestion(id)

        if (res === 'success') {
          this.$message({
            type: 'success',
            message: '删除成功!'
          });
          location.reload()
        } else {
          this.$confirm(res.data, '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'error'
          })
        }
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });          
      });
    },
    add() {
      this.$router.push({ path: 'create' })
    },
    detail(row) {
      this.$router.push({ path: 'preview', query: { id: row.id } })
    },
  },
  async mounted() {
    this.$data.tableData = await getTaskList()
  },
  data() {
    return {
      dialogTableVisible: true,
      tableData: [],
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
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
