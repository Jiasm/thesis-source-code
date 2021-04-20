<template>
  <div class="container">
    <common-header></common-header>
    <div class="main">
      <common-tree-menu></common-tree-menu>
      <div class="content">
        <el-table
          :data="tableData"
          style="width: 100%;margin-bottom: 20px;"
          row-key="id"
          border
          default-expand-all
          :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        >
          <el-table-column
            prop="username"
            label="用户名"
          >
          </el-table-column>
          <el-table-column
            prop="roleText"
            label="用户权限"
            width="150"
          >
          </el-table-column>
          <el-table-column
            prop="statusText"
            label="用户状态"
            width="150"
          >
          </el-table-column>
          <el-table-column
            prop="createdDate"
            label="添加时间"
            width="180"
          >
          </el-table-column>
          <el-table-column
            prop="desc"
            label="备注"
            width="200"
          >
          </el-table-column>
          <el-table-column
            label="操作"
            width="240"  
          >
            <template slot-scope="scope">
              <el-button type="warning" @click="changeRole(scope.row)" size="small">修改权限</el-button>
              <el-button type="error" @click="removeMember(scope.row)" size="small">移除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<script>
import { getProjectMemberList } from '../lib/api'
import Header from './Header'
import TreeMenu from './TreeMenu'

export default {
  components: { 'common-header': Header, 'common-tree-menu': TreeMenu },
  name: 'ProjectMemberList',
  methods: { },
  async mounted() {
    if (localStorage.getItem('project_id')) {
      this.$data.tableData = await getProjectMemberList(localStorage.getItem('project_id'))
    }

    window.addEventListener('change-project-id', async (e) => {
      this.$data.tableData = await getProjectMemberList(e.pid)  
    })
  },
  data() {
    return {
      tableData: [],
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.button-row {
  margin-top: 40px;
}

.main {
  display: flex;
}

.content {
  flex: 1;
  padding: 0 50px;
}

.tag {
  margin-right: 10px;
}
</style>
