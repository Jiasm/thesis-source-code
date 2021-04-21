<template>
  <div class="container">
    <common-header></common-header>
    <div class="main">
      <common-tree-menu></common-tree-menu>
      <common-invite :visible.sync="inviteDialogVisible" :close="closeInvite" :is-group="false" />
      <div class="content">
        <el-row>
          <el-page-header @back="goBack" content="项目成员管理">
          </el-page-header>
        </el-row>
        <el-row type="flex" justify="end">
          <el-col class="invite" :span="4">
            <el-button type="success" @click="openInvite" size="small">邀请用户</el-button>
          </el-col>
        </el-row>
        <el-row>
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
                <el-button v-if="scope.row.status === 1" type="danger" @click="removeMember(scope.row)" size="small">冻结</el-button>
                <el-button v-else type="success" @click="removeMember(scope.row)" size="small">启用</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
import { getProjectMemberList } from '../lib/api'
import Header from './Header'
import TreeMenu from './TreeMenu'
import Invite from './Invite'

export default {
  components: { 'common-header': Header, 'common-tree-menu': TreeMenu, 'common-invite': Invite },
  name: 'ProjectMemberList',
  methods: {
    goBack () {
      this.$router.go(-1)
    },
    openInvite() {
      this.$data.inviteDialogVisible = true
    },
    closeInvite () {
      this.$data.inviteDialogVisible = false
    },
  },
  async mounted() {
    if (localStorage.getItem('project_id')) {
      this.$data.projectId = Number(localStorage.getItem('project_id'))
      this.$data.tableData = await getProjectMemberList(localStorage.getItem('project_id'))
      console.log(this.$data.projectId)
    }

    window.addEventListener('change-project-id', async (e) => {
      this.$data.projectId = Number(e.pid)
      this.$data.tableData = await getProjectMemberList(e.pid)  
      console.log(this.$data.projectId)
    })
  },
  data() {
    return {
      projectId: 0,
      inviteDialogVisible: false,
      tableData: [],
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.el-row {
  margin: 20px 0;
}

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

.invite {
  text-align: right;
}
</style>
