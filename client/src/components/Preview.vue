<template>
  <div class="container">
    <h2>预览问卷调查</h2>
    <div class="create-panel">
      <el-row class="input-row">
        <el-col :span="8">
          <div class="input-field">问卷标题：</div>
        </el-col>
        <el-col :span="14">
          <div class="input-field">
            <el-input
              placeholder="请输入问卷标题"
              v-model="title"
              disabled
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row class="input-row">
        <el-col :span="8">
          <div class="input-field">问题描述：</div>
        </el-col>
        <el-col :span="14">
          <div class="input-field">
            <el-input
              placeholder="请输入问卷问题"
              v-model="question"
              disabled
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row class="input-row">
        <el-col :span="8">
        </el-col>
        <el-col>
          <el-checkbox-group v-model="checkList" class="check-group">
            <el-checkbox label="选项 A" disabled ></el-checkbox>
            <input value="选项A的描述" disabled /><br/>
            <el-checkbox label="选项 B" disabled ></el-checkbox>
            <input value="选项B的描述" disabled /><br/>
            <el-checkbox label="选项 C" disabled ></el-checkbox>
            <input disabled /><br/>
            <el-checkbox label="选项 D" disabled ></el-checkbox>
            <input disabled />
          </el-checkbox-group>
        </el-col>
      </el-row>
      <el-row class="button-row">
        <el-col :span="8" :offset="16">
          <el-button type="success" round>修改</el-button>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Preview',
  data() {
    return {
      title: '这里是问卷的标题',
      question: '这里是问题描述',
      checkList: ['选项 A', '选项 B'],
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.container {
  width: 480px;
  margin: 0 auto;
}
.input-row {
  display: flex;
  align-items: center;
  margin: 20px 0;
}
.input-field {
  text-align: center;
}
.button-row {
  text-align: right;
}
.el-checkbox:last-of-type {
  margin-right: 30px;
}
.check-group input {
  background-color: #F5F7FA;
  color: #C0C4CC;
  cursor: not-allowed;
  border: 1px solid #E4E7ED;
}
</style>
