<template>
  <div class="container">
    <img class="logo" src="../assets/logo.png">
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
      <el-row class="button-row">
        <el-col :span="8" :offset="8">
          <el-button type="primary" round @click="login">登录</el-button>
        </el-col>
        <el-col :span="8">
          <el-button type="warning" round @click="forget">忘记密码</el-button>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
// import { login } from '../lib/api'

export default {
  name: 'Login',
  props: {
    title: String,
  },
  data() {
    return {
      username: '',
      password: '',
    }
  },
  methods: {
    async login() {
      // const { username, password } = this.$data
      
      // const data = await login(username, password)

      // if (/^success:/.test(data)) {
      //   const id = data.match(/^success:(\d+)/)[1]
        
      //   sessionStorage.setItem('uid', id)

      //   this.$router.push({path: '/list' })
      // } else {
      //   this.$confirm(data, '提示', {
      //     confirmButtonText: '确定',
      //     cancelButtonText: '取消',
      //     type: 'error'
      //   })
      // }
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
</style>
