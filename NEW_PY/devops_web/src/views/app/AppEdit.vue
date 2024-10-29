<template>
  <el-dialog
          :model-value="visible"
          title="修改环境信息"
          @close="dialogClose"
          width="30%"
  >
     <el-form :model="row" ref="formRef" :rules="formRules" label-width="100px">
        <el-form-item label="环境名称：" prop="name">
          <el-input v-model="row.name"></el-input>
        </el-form-item>
        <el-form-item label="英文名称：" prop="english_name">
          <el-input v-model="row.english_name"></el-input>
          <span style="color: darkgrey">可以由字母、数字、横线和下划线组成。</span>
        </el-form-item>
        <el-form-item label="项目名称：" prop="project">
          <el-select class="m-2" v-model="row.project" @click="getProject" placeholder="请选择" style="width: 100%">
            <el-option v-for="row in project" :key="row.id" :label="row.name" :value="row.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="备注：">
          <el-input v-model="row.note" type="textarea"></el-input>
        </el-form-item>
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogClose">取消</el-button>
          <el-button type="primary" @click="submit">确定</el-button>
        </span>
      </template>

  </el-dialog>
</template>

<script>
    export default {
        props: {
            visible: Boolean,
            row: '',  // 当前行内容
        },
        data() {
            return {
                formRules: {
                    name: [
                        {required: true, message: '请输入凭据名称', trigger: 'blur'},
                        {min: 2, message: '凭据名称长度应不小于2个字符', trigger: 'blur'}
                    ],
                  english_name: [
                      {required: true, message: '请输入英文名称', trigger: 'blur'},
                      {min: 2, message: '英文名称长度应不小于2个字符', trigger: 'blur'},
                      {validator: function (rule, value, callback) {
                          // 效验英文
                          if(/[a-zA-Z0-9]$/.test(value) == false) {
                            callback(new Error('请输入英文名称'))
                          } else {
                            return callback();
                          }
                        }, trigger: 'blur'},
                  ],
                    project: [
                        {required: true, message: '请选择项目名称', trigger: 'blur'},
                    ],
                },
                project: ''
            }
        },
        methods: {
          dialogClose() {
            this.$emit('update:visible', false)  // 当对话框关闭，通过父组件更新为false
            this.$parent.getData();
          },
          submit() {
            this.$refs.formRef.validate((valid) => {
              if (valid) {
                this.$http.put('/app_release/app/' + this.row.id + '/', this.row)
                  .then(res => {
                    if (res.data.code == 200){
                      this.$message.success('修改成功');
                      this.dialogClose()  // 关闭窗口
                    }
                  })
              } else {
                this.$message.error('格式错误！')
              }
            })
          },
            // 获取项目
            getProject() {
                this.$http.get('/app_release/project/')
                    .then(res => {
                        this.project = res.data.data;
                    });
            },
        },
        // 监听打开编辑窗口
        watch: {
            visible() {
                this.row.project = this.row.project.id; // 重新赋值为ID
                this.getProject();
            }
        }
    }
</script>

<style scoped>

</style>