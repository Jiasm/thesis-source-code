<template>
  <div class="container">
    <h1>问卷调查系统</h1>
    <h2>用户注册</h2>
    <div class="signin-panel">
      <el-row class="input-row">
        <el-col :span="8">
          <div class="input-field">手机号：</div>
        </el-col>
        <el-col :span="16">
          <div class="input-field">
            <el-input
              placeholder="请输入手机号"
              v-model="phone"
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row class="input-row">
        <el-col :span="8">
          <div class="input-field">手机验证码：</div>
        </el-col>
        <el-col :span="16">
          <div class="input-field">
            <el-input
              placeholder="请输入六位数手机验证码"
              v-model="verify"
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row class="input-row">
        <el-col :span="8">
          <div class="input-field">昵称：</div>
        </el-col>
        <el-col :span="16">
          <div class="input-field">
            <el-input
              placeholder="请输入昵称"
              v-model="nickname"
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
        <el-col :span="4" :offset="20">
          <el-button type="success" round>注册</el-button>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: 'UserSignin',
  props: {
    title: String,
  },
  data() {
    return {
      username: '',
      password: '',
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
