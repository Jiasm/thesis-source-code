<template>
  <div class="header">
    <div class="system-title">项目管理系统</div>
    <el-button class="btn" type="success" @mouseover.native=showCreateTools @mouseleave.native=hideCreateTools >
      创建
      <div class="tools" v-if="showTools">
        <div class="item" @click="openCreateTask">创建任务</div>
        <div class="item" @click="openCreateTaskGroup">创建任务组</div>
        <div class="item" @click="openCreateProject">创建项目</div>
        <div class="item" @click="openCreateGroup">创建组织</div>
      </div>
    </el-button>
    <!-- <el-button class="btn" type="primary" >搜索</el-button>
    <input class="search-input" placeholder="搜索" /> -->
    <div class="nick-name">用户：{{username}}</div>
    <div class="btn-wrap">
      <el-button type="text" @click="logout">注销</el-button>
    </div>
    <common-create-group :visible.sync="groupDialogVisible" :close="closeCreateGroup" />
    <common-create-project :visible.sync="projectDialogVisible" :close="closeCreateProject" />
    <common-create-task :visible.sync="taskDialogVisible" :close="closeCreateTask" />
    <common-create-task-group :visible.sync="taskGroupDialogVisible" :close="closeCreateTaskGroup" />
  </div>
</template>

<script>
import CreateProject from './CreateProject';
import CreateGroup from './CreateGroup';
import CreateTask from './CreateTask';
import CreateTaskGroup from './CreateTaskGroup';
import { logout } from '../lib/api'
export default {
  name: 'Header',
  components: { 'common-create-project': CreateProject, 'common-create-group': CreateGroup, 'common-create-task': CreateTask, 'common-create-task-group': CreateTaskGroup },
  data() {
    return {
      projectDialogVisible: false,
      groupDialogVisible: false,
      taskDialogVisible: false,
      taskGroupDialogVisible: false,
      showTools: false,
      username: '',
    };
  },
  methods: {
    showCreateTools: function(){
      this.showTools = true
    },
    hideCreateTools: function(){
      this.showTools = false
    },
    openCreateProject() {
      this.$data.projectDialogVisible = true
    },
    closeCreateProject () {
      this.$data.projectDialogVisible = false
    },
    openCreateGroup() {
      this.$data.groupDialogVisible = true
    },
    closeCreateGroup () {
      this.$data.groupDialogVisible = false
    },
    openCreateTask() {
      this.$data.taskDialogVisible = true
    },
    closeCreateTask () {
      this.$data.taskDialogVisible = false
    },
    openCreateTaskGroup() {
      this.$data.taskGroupDialogVisible = true
    },
    closeCreateTaskGroup () {
      this.$data.taskGroupDialogVisible = false
    },
    async logout () {
      await logout()
      this.$router.push({ path: '/' })
    }
  },
  mounted () {
    const username = sessionStorage.getItem('username')

    this.$data.username = username
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.header {
  display: flex;
  position: relative;
  height: 50px;
  background-color: #fff;
  margin-bottom: 10px;
  border-bottom: 1px solid #edeff2;
  color: rgba(0,0,0,.65);
  align-items: center;
  justify-content: flex-end;
}

.btn {
  position: relative;
  height: 32px;
  line-height: 32px;
  padding-top: 0;
  padding-bottom: 0;
  display: inline-flex;
  margin-right: 10px;
}

.tools {
  position: absolute;
  right: 0;
  width: 100px;
  padding: 6px;
  background: #ffffff;
  border: 1px solid #ebecf0;
  border-radius: 4px;
  box-shadow: 0 10px 14px 0 rgb(51 51 51 / 10%);
  display: flex;
  flex-direction: column;
  overflow: hidden;
  z-index: 2;
}

.tools .item {
  height: 36px;
  border-radius: 4px;
  padding: 0 6px;
  cursor: pointer;
  overflow: hidden;
  display: flex;
  align-items: center;
  color: #333333;
}

.tools .item:hover {
  background-color: #f3f3f3;
}

.search-input {
  display: inline-flex;
  width: 260px;
  max-width: 260px;
  justify-content: space-between;
  height: 32px;
  border-radius: 17px;
  background-color: #ffffff;
  align-items: center;
  box-sizing: border-box;
  border: 1px solid rgba(102,111,128, .2);
  padding: 0 2em;
}

.nick-name {
  color: #333333;
  padding: 0 1em;
}

.btn-wrap { 
  padding: 0 1em;
}

.system-title {
  flex: 1;
  text-align: left;
  padding-left: 1em;
  font-weight: bold;
  font-size: 24px;
}
</style>
