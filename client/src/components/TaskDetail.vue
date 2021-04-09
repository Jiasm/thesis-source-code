<template>
  <el-dialog title="任务详情" :visible.sync="dialogTableVisible">
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="grid-content bg-purple">
          任务归属项目
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple">
          任务归属项目的值
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple">
          任务归属组
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple">
          任务归属组的值
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="grid-content bg-purple">
          任务标题
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple">
          这里是任务标题
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple">
          任务优先级
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple">
          高优先级
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="grid-content bg-purple">
          任务类型
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple">
          需求
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple">
          截止日期
        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple">
          2021-03-24
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="grid-content bg-purple">
          任务标签
        </div>
      </el-col>
      <el-col :span="18">
        <div class="grid-content bg-purple">
          标签1、标签2
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="grid-content bg-purple">
          任务描述
        </div>
      </el-col>
      <el-col :span="18">
        <div class="grid-content bg-purple">
          这里是任务描述
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="grid-content bg-purple">
          子任务
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="20">
        <el-table
        :data="tableData"
        style="width: 100%">
        <el-table-column
          prop="date"
          label="日期"
          width="180">
        </el-table-column>
        <el-table-column
          prop="name"
          label="姓名"
          width="180">
        </el-table-column>
        <el-table-column
          prop="address"
          label="地址">
        </el-table-column>
      </el-table>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="6">
        <el-button type="success">新增子任务</el-button>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="grid-content bg-purple">
          任务评论
        </div>
      </el-col>
      <el-col :span="18">
        <div class="grid-content bg-purple">
          评论内容
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="grid-content bg-purple">
          NickName
        </div>
      </el-col>
      <el-col :span="4"></el-col>
      <el-col :span="14">
        <div class="grid-content bg-purple">
          这里是评论内容
        </div>
      </el-col>
    </el-row>
  </el-dialog>
</template>

<script>
export default {
  name: 'TaskDetail',
  props: ['visible'],
  data() {
    return {
      dialogTableVisible: this.$props.visible,
       gridData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }],
      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1517 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1519 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }]
    }
  },
  methods: {
    
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
