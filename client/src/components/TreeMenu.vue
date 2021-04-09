<template>
  <div class="container">
    <div class="nav">
      <div class="nav-item">个人中心</div>
      <div class="nav-item">数据统计</div>
    </div>
    <el-menu
      default-active="2"
      class="treemenu"
    >
      <el-submenu index="1">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span>我的项目</span>
        </template>
        <el-menu-item index="1-1">项目 1</el-menu-item>
        <el-menu-item index="1-2">项目 2</el-menu-item>
      </el-submenu>
      <el-submenu index="2">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span>我的组织</span>
        </template>
          <el-submenu index="2-1">
            <template slot="title">
              <i class="el-icon-location"></i>
              <span>组织 1</span>
            </template>
            <el-menu-item index="2-1-1">项目 1</el-menu-item>
            <el-menu-item index="2-1-2">项目 2</el-menu-item>
          </el-submenu>
          <el-submenu index="2-2">
            <template slot="title">
              <i class="el-icon-location"></i>
              <span>组织 2</span>
            </template>
            <el-menu-item index="2-2-1">项目 1</el-menu-item>
            <el-menu-item index="2-2-2">项目 2</el-menu-item>
          </el-submenu>
      </el-submenu>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: 'TreeMenu',
  data() {
    return {
      showTools: false,
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
    }
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
  height: 40px;
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
