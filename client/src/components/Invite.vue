<template>
  <el-dialog width="480px" height="360px" title="邀请用户" v-bind="$attrs" @close="closeDialog" @open="loadData">
    <div class="content">
      <el-row type="flex" class="content-row row" :gutter="20">
        <el-row type="flex" class="row" :gutter="20">
          <el-col class="col col-title" :span="6">
            <div class="grid-content bg-purple">
              用户名
            </div>
          </el-col>
          <el-col class="col" :span="18">
            <div class="grid-content bg-purple">
              <el-select v-model="uid" placeholder="请选择用户" class="fill" :filterable="true">
                <el-option
                  v-for="item in userList"
                  :key="item.id"
                  :label="item.username"
                  :value="item.id"
                  :disabled="item.exists"
                >
                </el-option>
              </el-select>
            </div>
          </el-col>
        </el-row>
        <el-row type="flex" class="row" :gutter="20">
          <el-col class="col col-title" :span="6">
            <div class="grid-content bg-purple">
              角色
            </div>
          </el-col>
          <el-col class="col" :span="18">
            <div class="grid-content bg-purple">
              <el-select v-model="roleId" placeholder="请选择角色" class="fill">
                <el-option
                  v-for="item in ruleList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </div>
          </el-col>
        </el-row>
      </el-row>
      <el-row type="flex" class="row" :gutter="20" justify="end">
        <el-col class="col col-confirm" :span="4" >
          <el-button type="success" size="mini" @click="add">确定</el-button>
        </el-col>
      </el-row>
    </div>
  </el-dialog>
</template>

<script>
import { getAllUserList, addMemberToProject, addMemberToGroup, getProjectMemberList, getGroupMemberList } from '../lib/api'
import { role } from '../util'
export default {
  name: 'CreateProject',
  props: ['close', 'isGroup'],
  data() {
    return {
      projectId: 0,
      groupId: 0,
      roleId: undefined,
      uid: undefined,
      ruleList: role.filter(row => row.value),
      userList: []
    }
  },
  mounted () {
    this.$data.projectId = Number(localStorage.getItem('project_id'))
    this.$data.groupId = Number(localStorage.getItem('group_id'))
  },
  methods: {
    closeDialog() {
      this.$props.close()
    },
    async loadData () {
      const userList = await getAllUserList()

      let existsUserList = []

      if (this.$props.isGroup) {
        const groupUser = await getGroupMemberList(this.$data.groupId)
        existsUserList = groupUser.map(row => row.uid)
      } else {
        const projectUser = await getProjectMemberList(this.$data.projectId)
        existsUserList = projectUser.filter(row => !row.desc).map(row => row.uid)
      }

      console.log(existsUserList)

      this.$data.userList = userList.filter(row => !existsUserList.includes(row.id))
    },
    async add () {
      if (this.$props.isGroup) {
        // do something
        await addMemberToGroup(this.$data.groupId, this.$data.uid, this.$data.roleId)
      } else {
        await addMemberToProject(this.$data.projectId, this.$data.uid, this.$data.roleId)
      }
      this.$props.close()
      // await createProject(this.$data.projectName, this.$data.groupId)
      // this.$props.close()
    },
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.content {
  padding: 8px 40px;
  height: 200px;
  margin-top: 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.content-row {
  flex-direction: column;
}

.content-row .row {
  width: 100%;
  margin-bottom: 40px;
}

.row {
  margin-bottom: 18px;
  text-align: left;
  align-items: center;
}

.col-title {
  font-weight: bold;
}

.col-confirm {
  text-align: right;
}
</style>
