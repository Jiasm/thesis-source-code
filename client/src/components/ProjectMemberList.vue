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
              <template scope="scope">
                <el-select v-if="!scope.row.desc" v-model="scope.row.role" placeholder="用户权限" @change="changeRole(scope.row.uid, scope.row.role)">
                  <el-option
                    size="mini"
                    v-for="item in ruleList"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                  </el-option>
                </el-select>
                <span v-else>{{scope.row.roleText}}</span>
              </template>
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
                <!-- <el-button type="warning" @click="changeRole(scope.row)" size="small">修改权限</el-button> -->
                <el-button v-if="scope.row.status === 1 && !scope.row.desc" type="danger" @click="removeMember(scope.row.uid)" size="small">冻结</el-button>
                <el-button v-else-if="!scope.row.desc" type="success" @click="activeMember(scope.row.uid)" size="small">启用</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
import { getProjectMemberList, changeMemberRoleToProject, removeMemberToProject, activeMemberToProject } from '../lib/api'
import { role } from '../util'
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
    async closeInvite () {
      this.$data.inviteDialogVisible = false
      if (localStorage.getItem('project_id')) {
        this.$data.projectId = Number(localStorage.getItem('project_id'))
        this.$data.tableData = await getProjectMemberList(localStorage.getItem('project_id'))
      }
    },
    async changeRole (uid, roleId) {
      await changeMemberRoleToProject(this.$data.projectId, uid, roleId)
    },
    async removeMember(uid) {
      await removeMemberToProject(this.$data.projectId, uid)
    },
    async activeMember(uid) {
      await activeMemberToProject(this.$data.projectId, uid)
    },
  },
  async mounted() {
    if (localStorage.getItem('project_id')) {
      this.$data.projectId = Number(localStorage.getItem('project_id'))
      this.$data.tableData = await getProjectMemberList(localStorage.getItem('project_id'))
    }

    window.addEventListener('change-project-id', async (e) => {
      this.$data.projectId = Number(e.pid)
      this.$data.tableData = await getProjectMemberList(e.pid)  
    })
  },
  data() {
    return {
      projectId: 0,
      inviteDialogVisible: false,
      tableData: [],
      ruleList: role.filter(row => row.value),
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
