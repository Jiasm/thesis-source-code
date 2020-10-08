<template>
  <div class="container">
    <h2>创建问卷调查</h2>
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
            >
            </el-input>
          </div>
        </el-col>
        <el-col :span="2">
          <el-button type="primary" icon="el-icon-plus" circle size="mini"></el-button>
        </el-col>
      </el-row>
      <el-row class="input-row">
        <el-col :span="8">
        </el-col>
        <el-col>
          <el-checkbox label="选项 A" ></el-checkbox>
          <input v-model="optionsList[0]" /><br/>
          <el-checkbox label="选项 B"></el-checkbox>
          <input v-model="optionsList[1]" /><br/>
          <el-checkbox label="选项 C"></el-checkbox>
          <input v-model="optionsList[2]" /><br/>
          <el-checkbox label="选项 D"></el-checkbox>
          <input v-model="optionsList[3]" /><br/>
        </el-col>
      </el-row>
      <el-row class="button-row">
        <el-col :span="8" :offset="16">
          <el-button type="success" round @click="save">创建</el-button>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import { addQuestion } from '../lib/api'

export default {
  name: 'Create',
  data() {
    return {
      title: '',
      question: '',
      optionsList: ['', '', '', ''],
    }
  },
  methods: {
    async save() {
      console.log({
        title: this.$data.title,
        question: this.$data.question,
        optionsList: [this.$data.optionsList[0], this.$data.optionsList[1], this.$data.optionsList[2], this.$data.optionsList[3]]
      })
      const id = await addQuestion(this.$data.title, 1)

      this.$confirm('问卷创建成功', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'info'
      }).then(() => {
        this.$router.push({ path: '/preview', query: { id } })
      })
    }
  }
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
</style>
