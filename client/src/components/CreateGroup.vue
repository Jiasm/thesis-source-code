<template>
  <el-dialog width="640px" height="480px" title="创建组织" :visible.sync="dialogTableVisible">
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
            >
            </el-input>
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" class="row" :gutter="20" justify="end">
        <el-col class="col col-confirm" :span="4" >
          <el-button type="success" size="mini">确定</el-button>
        </el-col>
      </el-row>
    </div>
  </el-dialog>
</template>

<script>
export default {
  name: 'CreateGroup',
  props: ['visible'],
  data() {
    return {
      groupName: 'haha',
      dialogTableVisible: this.$props.visible,
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
  padding: 8px 40px;
  height: 400px;
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

.col-title {
  font-weight: bold;
}

.col-confirm {
  text-align: right;
}
</style>
