<template>
  <el-dialog width="480px" height="160px" title="创建项目" v-bind="$attrs" @close="closeDialog" @open="loadData">
    <el-container class="content" v-loading="loading" direction="vertical">
      <el-row type="flex" class="content-row row" :gutter="20">
        <el-row type="flex" class="row" :gutter="20">
          <el-col class="col col-title" :span="6">
            <div class="grid-content bg-purple">
              项目名
            </div>
          </el-col>
          <el-col class="col" :span="18">
            <div class="grid-content bg-purple">
              <el-input
                v-model="projectName"
                placeholder="请输入项目名"
              >
              </el-input>
            </div>
          </el-col>
        </el-row>
        <el-row type="flex" class="row" :gutter="20">
          <el-col class="col col-title" :span="6">
            <div class="grid-content bg-purple">
              归属组织
            </div>
          </el-col>
          <el-col class="col" :span="18">
            <div class="grid-content bg-purple">
              <el-select v-model="groupId" placeholder="请选择组织" class="fill select">
                <el-option
                  v-for="item in groupList"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id">
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
    </el-container>
  </el-dialog>
</template>

<script>
import { getGroupList, createProject } from '../lib/api'
export default {
  name: 'CreateProject',
  props: ['close'],
  data() {
    return {
      groupId: 0,
      projectName: '',
      groupList: [],
      loading: true,
    }
  },
  methods: {
    closeDialog() {
      this.$props.close()
    },
    async loadData () {
      this.$data.loading = true
      const groupList = await getGroupList()
      this.$data.loading = false

      this.$data.groupList = groupList
    },
    async add () {
      this.$data.loading = true
      await createProject(this.$data.projectName, this.$data.groupId)
      this.$data.loading = false
      this.$props.close()
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.content {
  padding: 8px 40px;
  height: 220px;
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

.select {
  width: 100%;
}
</style>
