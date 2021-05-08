<template>
  <div class="container">
    <h1>{{ title }}</h1>
    <div class="login-panel">
      <el-row class="input-row">
        <el-col :span="8">
          <div class="input-field">账号：</div>
        </el-col>
        <el-col :span="16">
          <div class="input-field">
            <el-input
              placeholder="请输入账号"
              v-model="username"
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row class="input-row">
        <el-col :span="8">
          <div class="input-field">密码：</div>
        </el-col>
        <el-col :span="16">
          <div class="input-field">
            <el-input
              placeholder="请输入密码"
              show-password
              v-model="password"
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row class="button-row" type="flex" justify="end">
        <!-- <el-col :span="8" :offset="8">
        </el-col> -->
        <el-col class="btn-wrap" :span="8">
          <el-button type="primary" round @click="signin">注册</el-button>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import { signin } from '../lib/api'

export default {
  name: 'Signin',
  data() {
    return {
      title: '软件开发项目管理系统',
      username: '',
      password: '',
    }
  },
  methods: {
    async signin() {
      const { username, password } = this.$data
      
      try {
        const data = await signin(username, password)

        if (data.code === 200) {
          try {
            await this.$confirm('注册成功，跳转登录页面', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'error'
            })
          } finally {
            this.$router.push({path:'/login'})
          }
        } else {
          this.$confirm('用户名已经存在', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'error'
          })
        }
      } catch (e) {
        this.$confirm('用户名已经存在', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        })
      }
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
h1 {
  margin-bottom: 120px;
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
  margin: 10% auto 0;
}
.logo {
  width: 200px;
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
.btn-wrap {
  display: flex;
  flex: 1;
  justify-content: flex-end;
}
</style>
