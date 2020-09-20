<template>
  <div class="container">
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
          <div class="input-field">手机号：</div>
        </el-col>
        <el-col :span="16">
          <div class="input-field">
            <el-input
              placeholder="请输入手机号"
              show-password
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row class="input-row">
        <el-col :span="8">
          <div class="input-field">验证码：</div>
        </el-col>
        <el-col :span="16">
          <div class="input-field">
            <el-input
              placeholder="请输入六位验证码"
              show-password
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row class="input-row">
        <el-col :span="8">
          <div class="input-field">新密码：</div>
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
      <el-row class="input-row">
        <el-col :span="8">
          <div class="input-field">确认密码：</div>
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
        <el-col :span="8" :offset="16">
          <el-button type="success" round>找回密码</el-button>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Reget',
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
