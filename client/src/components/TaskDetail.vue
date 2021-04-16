<template>
  <el-dialog width="640px" title="任务详情" v-bind="$attrs" @close="closeDialog" @open="loadData">
    <div class="content">
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务归属项目
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            <el-input
              v-model="projectName"
              :disabled="viewState">
            </el-input>
          </div>
        </el-col>
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务归属组
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            <el-input
              v-model="taskGroupName"
              :disabled="viewState">
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
              :disabled="viewState">
            </el-input>
          </div>
        </el-col>
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            任务优先级
          </div>
        </el-col>
        <el-col class="col" :span="6">
            <el-input
              v-model="priority"
              :disabled="viewState">
            </el-input>
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
              :disabled="viewState">
            </el-input>
          </div>
        </el-col>
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            执行人
          </div>
        </el-col>
        <el-col class="col" :span="6">
            <el-input
              v-model="executor"
              :disabled="viewState">
            </el-input>
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
            <el-input
              v-model="type"
              :disabled="viewState">
            </el-input>
          </div>
        </el-col>
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            截止日期
          </div>
        </el-col>
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            <el-input
              v-model="expireDate"
              type="datetime"
              :disabled="viewState">
            </el-input>
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
              :closable="!viewState"
              v-for="tag in tags"
              :disable-transitions="false"
              @close="handleClose(tag)"
            >
              {{tag.text}}
            </el-tag>
            <el-input
              v-show="!viewState"
              class="input-new-tag"
              v-if="inputVisible"
              v-model="inputValue"
              ref="saveTagInput"
              size="mini"
              @keyup.enter.native="handleInputConfirm"
              @blur="handleInputConfirm"
            />
            <el-button v-show="!viewState" v-else class="button-new-tag" size="mini" @click="showInput">添加标签</el-button>
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
              :disabled="viewState">
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
              <el-input v-if="scope.row.edit" size="mini" v-model="scope.row.executor" placeholder="请选择执行人"></el-input>
              <span v-else>{{scope.row.executor}}</span>
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
      <el-row type="flex" class="row split-line" :gutter="20">
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
              :disabled="viewState"
              v-model="comment.text">
            </el-input>
          </div>
        </el-col>
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
      <el-row v-show="!viewState" type="flex" class="row" :gutter="20" justify="end">
        <el-col class="col" :span="4" >
          <el-button type="success" size="mini">创建任务</el-button>
        </el-col>
      </el-row>
    </div>
  </el-dialog>
</template>

<script>
import { getTaskDetail, addComment, changeTask, newChildTask } from '../lib/api';
import { taskType, status, priority } from '../util'
export default {
  name: 'TaskDetail',
  props: ['viewState', 'taskId', 'close'],
  data() {
    return {
      id: 0,
      inputValue: '',
      projectName: '',
      taskGroupName: '',
      title: '',
      description: '',
      priority: '',
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
      addCommentText: '',
      incrNewChildTaskNum: 0,
      taskTypeList: taskType,
      statusList: status,
      priorityList: priority,
    }
  },
  methods: {
    handleClose(tag) {
      this.dynamicTags.splice(this.dynamicTags.indexOf(tag), 1);
    },
    showInput() {
      this.inputVisible = true;
      this.$nextTick(() => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    handleInputConfirm() {
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
    async saveChildTaskChange () {
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
          taskProjectId: row.taskProjectId,
          taskGroupId: row.taskGroupId,
          parentTaskId: this.$data.id,
          type: row.type,
          priority: row.priority,
        })
      }))

      this.$data.appendTaskVisible = false
      await this.loadData()
    },
    async sendComment () {
      await addComment(this.$data.id, this.$data.addCommentText)
      this.$data.addCommentVisible = false
      await this.loadData()
    },
    async loadData() {
      const data = await getTaskDetail(this.$props.taskId)

      this.$data.id = data.id
      this.$data.projectName = data.projectName
      this.$data.taskGroupName = data.taskGroupName
      this.$data.title = data.title
      this.$data.description = data.desc
      this.$data.creator = data.creator
      this.$data.executor = data.executor
      this.$data.priority = data.priority
      this.$data.type = data.type
      this.$data.expireDate = data.expireDate
      this.$data.childTask = data.childTask
      this.$data.tags = data.tags
      this.$data.commentList = data.commentList

      // comments: [
      //   {
      //     nickName: 'stark',
      //     text: '这里是评论内容'
      //   },
      //   {
      //     nickName: 'stark',
      //     text: '这里是评论内容'
      //   },
      // ],
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
</style>
