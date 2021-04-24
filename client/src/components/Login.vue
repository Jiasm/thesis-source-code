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
          <el-button type="primary" round @click="login">登录</el-button>
          <!-- <el-button type="text" round @click="forget">忘记密码</el-button> -->
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import { login } from '../lib/api'

export default {
  name: 'Login',
  data() {
    return {
      title: '软件开发项目管理系统',
      username: '',
      password: '',
    }
  },
  methods: {
    async login() {
      const { username, password } = this.$data
      
      try {
        const data = await login(username, password)

        if (data.code === 200) {
          // success
          sessionStorage.setItem('uid', data.data.id)
          sessionStorage.setItem('username', data.data.username)
          this.$router.push({ path: '/list' })
        } else {
          console.log(data.code)
          throw new Error('')
        }
      } catch (e) {
        this.$confirm('账号或密码错误', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        })
      }
    },
    forget() {
      this.$router.push({path:'/reget'})
    }
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
