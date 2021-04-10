<template>
  <el-dialog width="640px" title="任务详情" :visible.sync="dialogTableVisible">
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
              :key="tag"
              :closable="!viewState"
              v-for="tag in tags"
              :disable-transitions="false"
              @close="handleClose(tag)"
            >
              {{tag.name}}
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
          :data="tableData"
          style="width: 100%">
          <el-table-column
            prop="taskName"
            label="子任务名"
            width="180"
            fixed>
          </el-table-column>
          <el-table-column
            prop="status"
            label="子任务状态"
            width="100">
          </el-table-column>
          <el-table-column
            prop="priority"
            label="优先级"
            width="100">
          </el-table-column>
          <el-table-column
            prop="executor"
            label="执行人"
            width="100">
          </el-table-column>
          <el-table-column
            prop="expireDate"
            label="截止时间"
            width="180">
          </el-table-column>
        </el-table>
      </el-row>
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col" :span="6">
          <el-button type="success" size="mini">新增子任务</el-button>
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
      <el-row type="flex" v-for="(comment, index) in comments" v-bind:key="index" class="row" :gutter="20">
        <el-col class="col" :span="6">
          <div class="grid-content bg-purple">
            {{comment.nickName}}
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
          <el-button type="primary" size="mini">添加评论</el-button>
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
export default {
  name: 'TaskDetail',
  props: ['visible', 'viewState'],
  data() {
    return {
      inputVisible: false,
      inputValue: '',
      projectName: '项目名',
      taskGroupName: '项目对应组名',
      title: '这里是项目的标题',
      description: '项目描述信息',
      priority: 'P0',
      type: '需求',
      expireDate: '2021-05-01',
      tags: [
        {
          name: '标签 1',
        }, {
          name: '标签 2',
        },
      ],
      comments: [
        {
          nickName: 'stark',
          text: '这里是评论内容'
        },
        {
          nickName: 'stark',
          text: '这里是评论内容'
        },
      ],
      dialogTableVisible: this.$props.visible,
      tableData: [{
        taskName: '子任务 1',
        status: '未开始',
        priority: 'P0',
        executor: 'Jarvis',
        expireDate: '2021-05-01',
      }, {
        taskName: '子任务 2',
        status: '未开始',
        priority: 'P0',
        executor: 'Jarvis',
        expireDate: '2021-05-01',
      }, {
        taskName: '子任务 3',
        status: '未开始',
        priority: 'P0',
        executor: 'Jarvis',
        expireDate: '2021-05-01',
      }, {
        taskName: '子任务 4',
        status: '未开始',
        priority: 'P0',
        executor: 'Jarvis',
        expireDate: '2021-05-01',
      }]
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
      let inputValue = this.inputValue;
      if (inputValue) {
        this.tags.push({ name: inputValue });
      }
      this.inputVisible = false;
      this.inputValue = '';
    }
  }
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
