<template>
  <el-dialog width="960px" title="新建任务" v-bind="$attrs" @close="closeDialog" @open="loadData">
    <div class="content">
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务归属项目
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            <el-select v-model="projectId" placeholder="请选择项目" class="fill" @change="changeProjectId">
              <el-option
                v-for="item in projectList"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              >
              </el-option>
            </el-select>
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务归属组
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            <el-select v-model="taskGroupId" placeholder="请选择任务组" class="fill">
              <el-option
                v-for="item in projectGroupList"
                :key="item.id"
                :label="item.title"
                :value="item.id">
              </el-option>
            </el-select>
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务标题
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            <el-input
              v-model="title"
            >
            </el-input>
          </div>
        </el-col>
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            执行人
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <el-select v-model="executor" placeholder="请选择执行人" class="fill">
            <el-option
              v-for="item in userList"
              :key="item.id"
              :label="item.username"
              :value="item.id">
            </el-option>
          </el-select>
        </el-col>
      </el-row>
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务优先级
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <el-select v-model="priority" placeholder="请选择优先级" class="fill">
            <el-option
              v-for="item in priorityList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-col>
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务状态
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <el-select v-model="status" placeholder="请选择任务状态" class="fill">
            <el-option
              v-for="item in statusList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-col>
      </el-row>
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务类型
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            <el-select v-model="type" placeholder="请选择任务类型" class="fill">
              <el-option
                v-for="item in taskTypeList"
                :key="item.value"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
          </div>
        </el-col>
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            截止日期
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            <el-date-picker
              class="fill"
              v-model="expireDate"
              type="datetime"
              placeholder="选择日期时间"
            >
            </el-date-picker>
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务描述
          </div>
        </el-col>
        <el-col class="col" :span="18">
          <div class="grid-content bg-purple">
            <el-input
              v-model="description"
              type="textarea"
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" class="row" :gutter="20" justify="end">
        <el-col class="col right-flow" :span="12" >
          <el-button v-show="!showOnly" type="success" size="mini" @click="submitCreateTask">创建任务</el-button>
        </el-col>
      </el-row>
    </div>
  </el-dialog>
</template>

<script>
import { getAllUserList, getProjectList, getProjectGroupList, createTask } from '../lib/api';
import { taskType, status, priority } from '../util'
export default {
  name: 'CreateTask',
  props: ['viewState', 'taskId', 'close'],
  data() {
    return {
      id: 0,
      inputValue: '',
      projectId: '',
      taskGroupName: '',
      title: '',
      description: '',
      priority: '',
      status: '',
      type: '',
      creator: '',
      executor: '',
      expireDate: '',
      taskGroupId: 0,
      showOnly: this.$props.viewState,
      userList: [],
      taskTypeList: taskType,
      statusList: status,
      priorityList: priority,
      projectList: [],
      projectGroupList: [{ id: 0, title: '未分组' }],
    }
  },
  methods: {
    closeDialog() {
      this.$props.close()
    },
    async changeProjectId () {
      const projectGroupList = await getProjectGroupList(this.$data.projectId)

      this.$data.projectGroupList = projectGroupList
    },
    async loadData() {
      const userList = await getAllUserList()
      const projInfo = await getProjectList()

      this.$data.userList = userList
      this.$data.projectList = projInfo.projectList
    },
    async submitCreateTask () {
      await createTask({
        title: this.$data.title,
        description: this.$data.description,
        executor: this.$data.executor,
        status: this.$data.status,
        expireDate: this.$data.expireDate,
        projectId: this.$data.projectId,
        taskGroupId: this.$data.taskGroupId,
        type: this.$data.type,
        priority: this.$data.priority,
      })
      this.$props.close()
    }
  },
  async mounted () {
    // console.log('mounted')
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.content {
  padding: 8px;
}

.row {
  margin-bottom: 18px;
  text-align: left;
  align-items: center;
}

.col-title {
  font-weight: bold;
}

.split-line {
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid rgba(235,236,240,0.6);
}

.tag {
  margin-right: 10px;
}

.input-new-tag {
  width: 90px;
  margin-left: 10px;
  vertical-align: bottom;
}

.right-flow {
  text-align: right;
}

.fill {
  width: 100%!important;
}

.empty {
  text-align: center;
  color: #909399;
}

.empty-text {
  width: 820px;
}

.comment-header {
  border-bottom: 1px solid #EBEEF5;
  padding-bottom: 12px;
  color: #909399;
}
</style>
