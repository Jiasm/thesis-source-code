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
        <el-button type="primary" icon="el-icon-search" circle size="mini"></el-button>
        <el-button type="danger" icon="el-icon-delete" circle size="mini" @click="open" ></el-button>
      </el-table-column>
    </el-table>
      <el-row class="button-row">
        <el-col :span="8" :offset="16">
          <el-button type="success" round>新增问卷</el-button>
        </el-col>
      </el-row>
  </div>
</template>

<script>
export default {
  name: 'List',
  methods: {
    open() {
      this.$confirm('此操作将禁用问卷，后续用户将不能再进行填写, 是否确认?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$message({
          type: 'success',
          message: '删除成功!'
        });
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });          
      });
    }
  },
  data() {
    return {
      tableData: [{
        id: 1,
        title: '问卷调查1',
        status: '启用',
        datetime: '2020-04-01 16:00',
      }, {
        id: 2,
        title: '问卷调查2',
        status: '启用',
        datetime: '2020-03-28 15:00',
      }, {
        id: 3,
        title: '问卷调查3',
        status: '启用',
        datetime: '2020-03-20 12:00',
      }, {
        id: 4,
        title: '问卷调查4',
        status: '禁用',
        datetime: '2020-03-09 22:00',
      }]
    }
  }
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
