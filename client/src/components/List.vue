<template>
  <div class="container">
    <h2>问卷列表</h2>
    <el-table
      border
      stripe
      :data="tableData"
    >
      <el-table-column
        prop="id"
        label="问卷ID"
        width="80">
      </el-table-column>
      <el-table-column
        prop="title"
        label="问卷标题"
        width="120">
      </el-table-column>
      <el-table-column
        prop="status"
        label="问卷状态">
      </el-table-column>
      <el-table-column
        prop="datetime"
        label="创建时间">
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button type="primary" icon="el-icon-search" circle size="mini" @click="detail(scope.row)"></el-button>
          <el-button type="danger" icon="el-icon-delete" circle size="mini" @click="remove(scope.row)"></el-button>
        </template>
      </el-table-column>
    </el-table>
      <el-row class="button-row">
        <el-col :span="8" :offset="16">
          <el-button type="success" round @click="add">新增问卷</el-button>
        </el-col>
      </el-row>
  </div>
</template>

<script>
import { getList, removeItem } from '../lib/data'

export default {
  name: 'List',
  methods: {
    remove(row) {
      const id = Number(row.id)
      this.$confirm('此操作将禁用问卷，后续用户将不能再进行填写, 是否确认?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        removeItem(id)
        this.$message({
          type: 'success',
          message: '删除成功!'
        });
        const index = this.$data.tableData.findIndex(item => item.id === id)

        this.$data.tableData.splice(index, 1)
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
  data() {
    return {
      tableData: getList()
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container {
  max-width: 720px;
  margin: 0 auto;
}
.button-row {
  margin-top: 40px;
}
</style>
