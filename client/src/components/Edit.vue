<template>
  <div class="container">
    <h2>修改问卷调查</h2>
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
      </el-row>
      <el-row class="input-row">
        <el-col :span="8">
        </el-col>
        <el-col>
          <el-checkbox label="选项 A" v-model="optionsList[0].checked" ></el-checkbox>
          <input v-model="optionsList[0].val" /><br/>
          <el-checkbox label="选项 B" v-model="optionsList[1].checked" ></el-checkbox>
          <input v-model="optionsList[1].val" /><br/>
          <el-checkbox label="选项 C" v-model="optionsList[2].checked" ></el-checkbox>
          <input v-model="optionsList[2].val" /><br/>
          <el-checkbox label="选项 D" v-model="optionsList[3].checked" ></el-checkbox>
          <input v-model="optionsList[3].val" /><br/>
        </el-col>
      </el-row>
      <el-row class="button-row">
        <el-col :span="8" :offset="16">
          <el-button type="success" round @click="save">保存</el-button>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import { getQuestion, changeQuestion } from '../lib/api'

export default {
  name: 'Preview',
  data() {
    // const item = getQuestion(id)
    return {
      optionsList: [{}, {}, {}, {}],
      question: '',
      title: ''
    }
  },
  async mounted() {
    const id = Number(this.$route.query.id)
    const item = await getQuestion(id)
    this.$data.title = item.title
  },
  methods: {
    async save() {
      const id = Number(this.$route.query.id)
      await changeQuestion(id, this.$data.title)
        // question: this.$data.question,
      //   : this.$data.title
      // })
      this.$confirm('问卷修改成功', '提示', {
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
