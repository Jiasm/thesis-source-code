<template>
  <div class="container">
    <div class="nav">
      <div class="nav-item">个人中心</div>
      <div class="nav-item">数据统计</div>
    </div>
    <el-menu class="treemenu">
      <el-submenu index="1">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span>我的项目</span>
        </template>
        <el-menu-item v-for="(row, index) in projectList" v-bind:key="index" :index="`1-${row.id}`" @click="jumpProject">{{row.name}}</el-menu-item>
      </el-submenu>
      <el-submenu index="2">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span>我的组织</span>
        </template>
          <el-submenu v-for="(row, index) in groupList" v-bind:key="index" :index="`2-${index}`">
            <template slot="title">
              <i class="el-icon-location"></i>
              <span>{{row.groupName}}</span>
            </template>
            <el-menu-item v-for="(line, i) in row.list" v-bind:key="i" :index="`2-${index}-${line.id}`" @click="jumpProject">{{line.name}}</el-menu-item>
          </el-submenu>
      </el-submenu>
    </el-menu>
  </div>
</template>

<script>
import { getProjectList } from '../lib/api'

export default {
  name: 'TreeMenu',
  data() {
    return {
      showTools: false,
      projectList: [],
      groupList: [],
    };
  },
  methods: {
    showCreateTools: function(){
      console.log('trigger')
      this.showTools = true
    },
    hideCreateTools: function(){
      console.log('trigger close')
      this.showTools = false
    },
    jumpProject (item) {
      const index = item.index.split('-').pop()

      localStorage.setItem('project_id', index)

      const event = new Event('change-project-id')

      event.pid = index

      window.dispatchEvent(event)
    }
  },
  async mounted () {
    const { projectList, groupList } = await getProjectList()

    this.$data.projectList = projectList
    this.$data.groupList = groupList
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container {
  width: 250px;
  overflow: hidden;
  height: 100%;
  position: relative;
  left: 0;
  top: 0;
  bottom: 0;
  border-right: 1px solid rgba(235,236,240,0.6);
  padding: 0;
  padding-bottom: constant(safe-area-inset-bottom);
  padding-bottom: env(safe-area-inset-bottom);
  z-index: 200;
  transform: translate(0, 0);
}

.treemenu {
  text-align: left;
}

.nav {
  border-bottom: 1px solid rgba(235,236,240,0.6);
}

.nav-item {
  display: flex;
  height: 60px;
  border-radius: 4px;
  align-items: center;
  justify-content: space-between;
  padding-left: 18px;
  padding-right: 10px;
  font-size: 14px;
  font-weight: 400;
  color: #333333;
  cursor: pointer;
  font-weight: 500;
}

.nav-item:hover {
  background-color: #f6f7f9;
  color: #333333;
}
</style>
