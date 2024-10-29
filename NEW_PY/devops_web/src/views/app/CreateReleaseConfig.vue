<template>
  <el-dialog
          :model-value="visible"
          title="新建发布"
          @close="dialogClose"
          width="45%"
  >
     <el-steps :active="active" align-center style="margin-bottom: 10%">
      <el-step title="基本配置"></el-step>
      <el-step title="构建配置"></el-step>
      <el-step title="部署配置"></el-step>
     </el-steps>
     <el-form :model="form" ref="formRef" :rules="formRules" label-width="160px">
      <div v-show="active == 1">
          <el-form-item label="发布环境：" prop="env">
            <el-select class="m-2" v-model="form.env" @click="getEnv" placeholder="请选择" style="width: 100%">
              <el-option v-for="row in env" :key="row.id" :label="row.name" :value="row.id">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="目标主机：" prop="server_ids">
            <el-select class="m-2" v-model="form.server_ids" @click="getServer" multiple placeholder="请选择" style="width: 100%">
              <el-option v-for="row in server" :key="row.id" :label="row.name" :value="row.id">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="Git仓库：" prop="git_repo">
            <el-input v-model="form.git_repo"></el-input>
          </el-form-item>
          <el-form-item label="Git凭据：" prop="git_credential_id">
            <el-select class="m-2" v-model="form.git_credential_id" @click="getCredential"  placeholder="请选择" style="width: 100%" clearable>
              <el-option label="无" value="0"></el-option>
              <el-option v-for="row in credential" :key="row.id" :label="row.name" :value="row.id">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="消息通知：">
            <el-select class="m-2" v-model="form.notify_id" @click="getNotify"  placeholder="请选择" style="width: 100%" clearable>
              <el-option label="无" value="0"></el-option>
              <el-option v-for="row in notify" :key="row.id" :label="row.name" :value="row.id">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="备注：">
            <el-input v-model="form.note" type="textarea"></el-input>
          </el-form-item>
      </div>

      <div v-show="active == 2">
          <el-form-item label="源文件：">
            <el-input v-model="form.source_file"></el-input>
          </el-form-item>
          <el-form-item label="自定义全局变量：">
            <el-input v-model="form.global_variables" type="textarea"></el-input>
          </el-form-item>
          <el-form-item label="代码检出前执行脚本：">
            <el-input v-model="form.pre_checkout_script" type="textarea"></el-input>
          </el-form-item>
          <el-form-item label="代码检出后执行脚本：">
            <el-input v-model="form.post_checkout_script" type="textarea"></el-input>
          </el-form-item>
      </div>

      <div v-show="active == 3">
          <el-form-item label="服务器部署路径：" prop="dst_dir">
            <el-input v-model="form.dst_dir"></el-input>
          </el-form-item>
          <el-form-item label="历史版本路径：" prop="history_version_dir">
            <el-input v-model="form.history_version_dir"></el-input>
          </el-form-item>
          <el-form-item label="历史版本保留数：" prop="history_version_number">
            <el-input v-model="form.history_version_number" placeholder="默认7"></el-input>
          </el-form-item>
          <el-form-item label="部署前执行脚本：">
            <el-input v-model="form.pre_deploy_script" type="textarea"></el-input>
          </el-form-item>
          <el-form-item label="部署后执行脚本：">
            <el-input v-model="form.post_deploy_script" type="textarea"></el-input>
          </el-form-item>
      </div>
    </el-form>

      <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogClose" v-if="active == 1">取消</el-button>
            <el-button type="primary" @click="dialogToggle('pre')" v-if="active > 1">上一步</el-button>
            <el-button type="primary" @click="dialogToggle('next')" v-if="active < 3">下一步</el-button>
            <el-button type="primary" @click="submit" v-if="active == 3">确定</el-button>
          </span>
      </template>

  </el-dialog>
</template>

<script>
    export default {
        props: {
            visible: Boolean,
            rowId: '',  // 应用ID
        },
        data() {
            return {
                form: {
                  app: '',
                  env: '',
                  server_ids: '',
                  git_repo: '',
                  git_credential_id: '',
                  notify_id: '',
                  note: '',
                  source_file: '*',
                  global_variables: '',
                  pre_checkout_script: '',
                  post_checkout_script: '',
                  dst_dir: '',
                  history_version_dir: '',
                  history_version_number: '',
                  pre_deploy_script: '',
                  post_deploy_script: ''
                },
                formRules: {
                  env: [
                      {required: true, message: '请选择发布环境', trigger: 'blur'}
                  ],
                  server_ids: [
                      {required: true, message: '请选择目标主机', trigger: 'blur'},
                  ],
                  git_repo: [
                      {required: true, message: '请输入Git仓库地址', trigger: 'blur'},
                      {min: 16, message: 'Git仓库地址长度应不小于16个字符', trigger: 'blur'}
                  ],
                  dst_dir: [
                      {required: true, message: '请输入部署路径', trigger: 'blur'},
                      {min: 3, message: '请输入部署路径应不小于3个字符', trigger: 'blur'}
                  ],
                  history_version_dir: [
                      {required: true, message: '请输入历史版本路径', trigger: 'blur'},
                      {min: 3, message: '请输入历史版本路径应不小于3个字符', trigger: 'blur'}
                  ]
                },
                active: 1, // 默认显示第一个步骤
                env: '',
                server: '',
                credential: '',
                notify: ''
            }
        },
        methods: {
          dialogClose() {
            this.$emit('update:visible', false);  // 当对话框关闭，通过父组件更新为false
          },
          submit() {
            this.form.app = this.rowId;
            this.form.history_version_number = 7;
            this.$refs.formRef.validate((valid) => {
              if (valid) {
                this.$http.post('/app_release/config/', this.form)
                  .then(res => {
                    if (res.data.code == 200){
                      this.$message.success('创建成功');
                      this.dialogClose()  // 关闭窗口
                    }
                  })
              } else {
                this.$message.error('格式错误！')
              }
            })
          },
            dialogToggle(action) {
                if (action == 'pre') {
                    if (this.active-- < 2) {
                      this.active = 1;
                    }
                } else if (action == 'next') {
                    if (this.active++ > 3) {
                      this.active = 1;
                    }
                }
            },
            // 获取发布环境
            getEnv() {
              this.$http.get('/app_release/env/')
                .then(res => {
                    this.env = res.data.data;
                })
            },
            // 获取主机
            getServer() {
              this.$http.get('/cmdb/server/')
                .then(res => {
                    this.server = res.data.data;
                })
            },
            // 获取凭据
            getCredential() {
              this.$http.get('/config/credential/')
                .then(res => {
                    this.credential = res.data.data;
                })
            },
            // 获取通知
            getNotify() {
              this.$http.get('/config/notify/')
                .then(res => {
                    this.notify = res.data.data;
                })
            }
        },
    }
</script>

<style scoped>

</style>