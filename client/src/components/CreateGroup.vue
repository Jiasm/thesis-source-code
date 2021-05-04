<template>
  <el-dialog width="480px" height="120px" title="创建组织" v-bind="$attrs" @close="closeDialog">
    <div class="content">
      <el-row type="flex" class="row" :gutter="20">
        <el-col class="col col-title" :span="6">
          <div class="grid-content bg-purple">
            组织名
          </div>
        </el-col>
        <el-col class="col" :span="18">
          <div class="grid-content bg-purple">
            <el-input
              v-model="groupName"
              placeholder="请输入组织名"
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" class="row row-end" :gutter="20" justify="end">
        <el-col class="col col-confirm" :span="4" >
          <el-button type="success" size="mini" @click="add">确定</el-button>
        </el-col>
      </el-row>
    </div>
  </el-dialog>
</template>

<script>
import { createGroup } from '../lib/api'
export default {
  name: 'CreateGroup',
  props: ['close'],
  data() {
    return {
      groupName: '',
    }
  },
  methods: {
    closeDialog() {
      this.$props.close()
    },
    async add () {
      await createGroup(this.$data.groupName)
      this.$props.close()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.content {
  padding: 8px 40px;
  height: 140px;
  margin-top: 30px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.row {
  margin-bottom: 18px;
  text-align: left;
  align-items: center;
}

.row-end {
  margin-bottom: 0;
}

.col-title {
  font-weight: bold;
}

.col-confirm {
  text-align: right;
}
</style>
