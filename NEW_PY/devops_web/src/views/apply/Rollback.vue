<template>
    <el-dialog
      :model-value="visible"
      width="40%"
      title="回滚"
      @close="dialogClose"
      >
      <el-form :model="form" ref="formRef" :rules="formRules" label-position=“right” label-width="100px" >
        <el-form-item label="选择版本：" prop="version_id">
          <el-select class="m-2" v-model="form.version_id" placeholder="请选择" style="width: 100%">
            <el-option v-for="version in rollback" :key="version.id" :label="`${version.version_id} (${version.title})`" :value="version.version_id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="目标主机：" prop="server_ids">
          <el-select class="m-2" v-model="form.server_ids" @click="getServer" placeholder="请选择" style="width: 100%" multiple>
            <el-option v-for="row in server" :key="row.id" :label="row.name" :value="row.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="代码回滚后执行脚本：">
          <el-input v-model="form.post_rollback_script" type="textarea"></el-input>
        </el-form-item>
        <el-form-item label="备注：">
          <el-input v-model="form.note" type="textarea"></el-input>
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
            row: '',
        },
        data() {
            return {
                form: {
                  'version_id': '',
                  'server_ids': '',
                  'post_rollback_script': '',
                  'note': ''
                },
                formRules: {
                    title: [
                        {required: true, message: '请输入标题', trigger: 'blur'},
                        {min: 2, message: '凭据名称长度应不小于2个字符', trigger: 'blur'}
                    ],
                    version_id: [
                        {required: true, message: '请选择回滚版本', trigger: 'blur'},
                    ],
                    server_ids: [
                        {required: true, message: '请选择目标主机', trigger: 'blur'},
                    ],
                },
                rollback: '',
                server: ''
            }
        },
        methods: {
            // 关闭对话框，修改父组件值
            dialogClose() {
                this.$emit('update:visible', false)
            },
            getServer() {
                this.$http.get('/cmdb/server/')
                  .then(res => {
                      this.server = res.data.data;
                  });
            },
            submit() {
              this.$refs.formRef.validate((valid) => {
                if (valid) {
                 const row = {
                   'dst_dir': this.row.release_config.dst_dir,
                   'history_version_dir': this.row.release_config.history_version_dir,
                   'apply_id': this.row.id,
                 };
                  const config = Object.assign(row, this.form);
                  this.$http.post('/app_release/rollback/', config)
                    .then(res => {
                      if (res.data.code == 200){
                        this.$message.success('回滚成功');
                        this.dialogClose()  // 关闭窗口
                      } else {
                        this.$message.error(res.data.data)
                      }
                    })
                } else {
                  this.$message.error('格式错误！')
                }
              })
            },
        },
        watch: {
            visible() {
               if(this.visible) {
                 const config = {
                   'project_id': this.row.release_config.app.project.id,
                   'env_id': this.row.release_config.env.id,
                   'app_id': this.row.release_config.app.id
                 };
                 // console.log(config)
                  this.$http.get('/app_release/rollback/', {params: config})
                      .then(res => {
                        if(res.data.code == 200) {
                            this.rollback = res.data.data;
                        } else {
                          this.dialogClose();  // 没有发布记录直接关闭窗口
                        }
                      });

               }
            },
          // 监听选择回滚版本后，自动渲染目标主机
          'form.version_id': {
                handler: function () {
                    this.getServer();  // 先渲染所有主机ID，再将已经选择的主机ID重新赋值
                    this.rollback.forEach(row =>{
                        if(row.version_id == this.form.version_id) {
                            this.form.server_ids = row.server_ids;
                        }
                    })
                }
            }
        }
    }
</script>

<style scoped>

</style>