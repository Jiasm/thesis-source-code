<template>
  <el-dialog width="960px" height="480px" title="创建任务组" v-bind="$attrs" @close="closeDialog">
    <div class="content">
      <el-row type="flex" class="content-row row" :gutter="20">
        <el-row type="flex" class="row" :gutter="20">
          <el-col class="col col-title" :span="6">
            <div class="grid-content bg-purple">
              任务组名
            </div>
          </el-col>
          <el-col class="col" :span="18">
            <div class="grid-content bg-purple">
              <el-input
                v-model="taskGroupName"
              >
              </el-input>
            </div>
          </el-col>
        </el-row>
        <el-row type="flex" class="row" :gutter="20">
          <el-col class="col col-title" :span="6">
            <div class="grid-content bg-purple">
              任务组描述
            </div>
          </el-col>
          <el-col class="col" :span="18">
            <div class="grid-content bg-purple">
              <el-input
                v-model="taskGroupDesc"
                type="textarea"
              >
              </el-input>
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
import { createTaskGroup } from '../lib/api'
export default {
  name: 'CreateTaskGroup',
  props: ['close'],
  data() {
    return {
      taskGroupName: '',
      taskGroupDesc: '',
    }
  },
  methods: {
    closeDialog() {
      this.$props.close()
    },
    async add () {
      await createTaskGroup(this.$data.taskGroupName, this.$data.taskGroupDesc)
      this.$props.close()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.content {
  padding: 8px 40px;
  height: 400px;
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
