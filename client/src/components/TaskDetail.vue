<template>
  <el-dialog width="960px" title="任务详情" top="5vh" v-bind="$attrs" @close="closeDialog" @open="loadData">
    <el-container class="content" v-loading="loading" direction="vertical">
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务归属项目
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            <el-select v-model="projectId" placeholder="请选择项目" class="fill" :disabled="showOnly">
              <el-option
                v-for="item in projectList"
                :key="item.id"
                :label="item.name"
                :value="item.id">
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
            <el-input
              v-model="taskGroupName"
              :disabled="showOnly">
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            创建人
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            <el-input
              v-model="creator"
              :disabled="showOnly">
            </el-input>
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
              @change="changeTitle"  
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
          <el-select v-model="executor" placeholder="请选择执行人" class="fill" @change="changeExecutor">
            <el-option
              v-for="item in userList"
              :key="item.uid"
              :label="item.username"
              :value="item.uid">
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
          <el-select v-model="priority" placeholder="请选择优先级" class="fill" @change="changePriority">
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
          <el-select v-model="status" placeholder="请选择任务状态" class="fill" @change="changeStatus">
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
            <el-select v-model="type" placeholder="请选择任务类型" class="fill" @change="changeType">
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
              @change="changeExpireDate"
            >
            </el-date-picker>
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务标签
          </div>
        </el-col>
        <el-col class="col" :span="18">
          <div class="grid-content bg-purple">
            <el-tag
              class="tag"
              size="small"
              :key="tag.id"
              :closable="true"
              v-for="tag in tags"
              :disable-transitions="false"
              @close="handleRemoveTag(tag)"
            >
              {{tag.text}}
            </el-tag>
            <el-input
              class="input-new-tag"
              v-if="inputVisible"
              v-model="inputValue"
              ref="saveTagInput"
              size="mini"
              @keyup.enter.native="handleInputConfirm"
            />
            <el-button v-else class="button-new-tag" size="mini" @click="showInput">添加标签</el-button>
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
              @change="changeDesc"  
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" class="row split-line" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            子任务
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" class="row" :gutter="20">
        <el-table
          :data="childTask"
          row-key="id"
          style="width: 100%"
          @row-click="changeRow"
        >
          <el-table-column
            prop="title"
            label="子任务名"
            width="180"
            fixed>
            <template scope="scope">
              <el-input v-if="scope.row.edit" size="mini" v-model="scope.row.title" placeholder="请输入任务名"></el-input>
              <span v-else>{{scope.row.title}}</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="status"
            label="子任务状态"
            width="100">
            <template scope="scope">
              <el-select v-if="scope.row.edit" size="mini" v-model="scope.row.status" placeholder="请选择任务状态">
                <el-option
                  v-for="item in statusList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
              <span v-else>{{scope.row.statusText}}</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="type"
            label="任务类型"
            width="100">
            <template scope="scope">
              <el-select v-if="scope.row.edit" size="mini" v-model="scope.row.type" placeholder="请选择任务类型">
                <el-option
                  v-for="item in taskTypeList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
              <span v-else>{{scope.row.typeText}}</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="priority"
            label="优先级"
            width="100">
            <template scope="scope">
              <el-select v-if="scope.row.edit" size="mini" v-model="scope.row.priority" placeholder="请选择优先级">
                <el-option
                  v-for="item in priorityList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
              <span v-else>{{scope.row.priorityText}}</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="executor"
            label="执行人"
            width="100">
            <template scope="scope">
              <el-select v-if="scope.row.edit" v-model="scope.row.executor" placeholder="请选择执行人">
                <el-option
                  size="mini"
                  v-for="item in userList"
                  :key="item.id"
                  :label="item.username"
                  :value="item.id">
                </el-option>
              </el-select>
              <span v-else>{{scope.row.executorText}}</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="expireDate"
            label="截止时间"
            width="240">
            <template scope="scope">
              <el-date-picker
                v-if="scope.row.edit"
                v-model="scope.row.expireDate"
                size="mini"
                type="datetime"
                placeholder="选择日期时间">
              </el-date-picker>
              <span v-else>{{scope.row.expireDate}}</span>
            </template>
          </el-table-column>
        </el-table>
      </el-row>
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col" :span="12">
          <el-button type="primary" size="mini" @click="showAppendTask">新增子任务</el-button>
          <el-button v-if="appendTaskVisible" type="success" size="mini" @click="saveChildTaskChange">保存</el-button>
        </el-col>
      </el-row>
      <el-row type="flex" class="row split-line comment-header" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务评论
          </div>
        </el-col>
        <el-col class="col col-title" :span="18">
          <div class="grid-content bg-purple">
            评论内容
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" v-for="(comment, index) in commentList" v-bind:key="index" class="row" :gutter="20">
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            {{comment.username}}
          </div>
        </el-col>
        <el-col class="col" :span="18">
          <div class="grid-content bg-purple">
            <el-input
              type="textarea"
              :rows="2"
              :disabled="true"
              v-model="comment.text">
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row v-if="!commentList || commentList.length === 0" type="flex" class="row empty" :gutter="20">
        <span class="empty-text">暂无数据</span>
      </el-row>
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col" :span="6">
          <el-button type="primary" size="mini" @click="showAddComment">添加评论</el-button>
        </el-col>
      </el-row>
      <el-row type="flex" class="row" v-if="addCommentVisible" :gutter="20">
        <el-col class="col" :span="18">
          <div class="grid-content bg-purple">
            <el-input
              type="textarea"
              :rows="2"
              v-model="addCommentText">
            </el-input>
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <el-button type="success" size="mini" @click="sendComment">发送</el-button>
        </el-col>
      </el-row>
    </el-container>
  </el-dialog>
</template>

<script>
import { getTaskDetail, addComment, changeTask, newChildTask, addNewTag, removeTag, getProjectMemberList, getProjectList } from '../lib/api';
import { taskType, status, priority } from '../util'
export default {
  name: 'TaskDetail',
  props: ['viewState', 'taskId', 'close'],
  data() {
    return {
      id: 0,
      inputValue: '',
      taskGroupName: '',
      title: '',
      description: '',
      priority: '',
      status: '',
      type: '',
      creator: '',
      executor: '',
      expireDate: '',
      tags: [],
      commentList: [],
      childTask: [],
      inputVisible: false,
      dialogTableVisible: true,
      addCommentVisible: false,
      appendTaskVisible: false,
      taskGroupId: 0,
      projectId: 0,
      showOnly: this.$props.viewState,
      userList: [],
      addCommentText: '',
      incrNewChildTaskNum: 0,
      taskTypeList: taskType,
      statusList: status,
      priorityList: priority,
      projectList: [],
      loading: true,
    }
  },
  methods: {
    async handleRemoveTag(tag) {
      this.$data.loading = true
      await removeTag(this.$data.id, tag.id)
      await this.loadData()
    },
    showInput() {
      this.inputVisible = true;
      this.$nextTick(() => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    async handleInputConfirm () {
      this.$data.loading = true
      const newTag = this.$data.inputValue
      await addNewTag(this.$data.id, newTag)
      this.inputVisible = false
      await this.loadData()
    },
    closeDialog() {
      this.$props.close()
    },
    showAddComment () {
      this.$data.addCommentVisible = true
    },
    showAppendTask () {
      this.$data.appendTaskVisible = true
      this.$data.childTask.push({ id: `unkonwn-${this.$data.incrNewChildTaskNum}`, edit: true })
      this.$data.incrNewChildTaskNum += 1
    },
    changeRow (row) {
      const index = this.$data.childTask.findIndex(r => r.id === row.id)
      this.$data.appendTaskVisible = true
      this.$set(this.$data.childTask, index, { ...this.$data.childTask[index], edit: true })
      // this.$data.childTask[index].edit = true
    },
    async changeStatus () {
      this.$data.loading = true
      await changeTask({
        id: this.$data.id,
        status: this.$data.status
      })
      await this.loadData()
    },
    async changeType () {
      this.$data.loading = true
      await changeTask({
        id: this.$data.id,
        taskType: this.$data.type
      })
      await this.loadData()
    },
    async changePriority () {
      this.$data.loading = true
      await changeTask({
        id: this.$data.id,
        priority: this.$data.priority
      })
      await this.loadData()
    },
    async changeExpireDate () {
      this.$data.loading = true
      await changeTask({
        id: this.$data.id,
        expireDate: this.$data.expireDate
      })
      await this.loadData()
    },
    async changeTitle () {
      this.$data.loading = true
      await changeTask({
        id: this.$data.id,
        title: this.$data.title
      })
      await this.loadData()
    },
    async changeExecutor () {
      this.$data.loading = true
      await changeTask({
        id: this.$data.id,
        executor: this.$data.executor
      })
      await this.loadData()
    },
    async changeDesc () {
      this.$data.loading = true
      await changeTask({
        id: this.$data.id,
        description: this.$data.description
      })
      await this.loadData()
    },
    async saveChildTaskChange () {
      this.$data.loading = true
      const changedRows = []
      const newRows = []
      this.$data.childTask.filter(row => row.edit).forEach(row => {
        if (/^\d+$/.test(row.id)) {
          changedRows.push(row)
        } else {
          newRows.push(row)
        }
      })

      await Promise.all(changedRows.map(row => {
        return changeTask({
          id: row.id,
          title: row.title,
          executor: row.executor,
          status: row.status,
          expireDate: row.expireDate,
          type: row.type,
          priority: row.priority,
        })
      }))

      await Promise.all(newRows.map(row => {
        return newChildTask({
          title: row.title,
          desc: row.desc,
          executor: row.executor,
          status: row.status,
          expireDate: row.expireDate,
          taskProjectId: this.$data.taskGroupId,
          taskGroupId: this.$data.projectId,
          parentTaskId: this.$data.id,
          type: row.type,
          priority: row.priority,
        })
      }))

      this.$data.appendTaskVisible = false
      await this.loadData()
    },
    async sendComment () {
      this.$data.loading = true
      await addComment(this.$data.id, this.$data.addCommentText)
      this.$data.addCommentVisible = false
      await this.loadData()
    },
    async loadData() {
      this.$data.loading = true
      const [data, projInfo] = await Promise.all([
        getTaskDetail(this.$props.taskId),
        getProjectList(),
      ])
      const userList = await getProjectMemberList(data.projectId)

      this.$data.id = data.id
      this.$data.taskGroupName = data.taskGroupName
      this.$data.title = data.title
      this.$data.description = data.desc
      this.$data.creator = data.creator
      this.$data.executor = data.executor
      this.$data.priority = data.priority
      this.$data.status = data.status
      this.$data.type = data.type
      this.$data.expireDate = data.expireDate
      this.$data.childTask = data.childTask
      this.$data.tags = data.tags
      this.$data.commentList = data.commentList
      this.$data.userList = userList
      this.$data.taskGroupId = data.taskGroupId
      this.$data.projectId = data.projectId
      this.$data.projectList = projInfo.projectList
      this.$data.loading = false
    }
  },
  async mounted () {
    console.log('mounted')
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
