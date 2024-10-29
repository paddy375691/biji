<template>
  <el-dialog
          :model-value="visible"
          title="修改通知"
          @close="dialogClose"
          width="30%"
  >
     <el-form :model="row" ref="formRef" :rules="formRules" label-width="100px">
          <el-form-item label="名称：" prop="name">
            <el-input v-model="row.name"></el-input>
          </el-form-item>
          <el-form-item label="邮箱收件人：" prop="email_recipient" v-if="row.notify_mode==1">
            <el-input v-model="row.email_recipient" ></el-input>
            <span style="color: darkgrey">多个邮件以系统逗号分割。例如zhangsan@163.com,lisi@163.com</span>
          </el-form-item>
          <el-form-item label="钉钉Webhook：" prop="dingding_webhook" v-if="row.notify_mode==2">
            <el-input v-model="row.dingding_webhook" ></el-input>
          </el-form-item>
          <el-form-item label="企业微信Webhook：" prop="weixin_webhook" v-if="row.notify_mode==3">
            <el-input v-model="row.weixin_webhook"></el-input>
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
                      {required: true, message: '请输入消息通知名称', trigger: 'blur'},
                      {min: 2, message: '消息通知长度应不小于2个字符', trigger: 'blur'}
                  ],
                  email_recipient: [
                      {required: true, message: '请输入邮箱收件人', trigger: 'blur'},
                      {min: 6, message: '邮箱收件人长度应不小于6个字符', trigger: 'blur'}
                  ],
                  dingding_webhook: [
                      {required: true, message: '请输入钉钉Webhook地址', trigger: 'blur'},
                      {min: 20, message: '地址长度应不小于20个字符', trigger: 'blur'}
                  ],
                  weixin_webhook: [
                      {required: true, message: '请输入企业微信Webhook地址', trigger: 'blur'},
                      {min: 20, message: '地址长度应不小于20个字符', trigger: 'blur'}
                  ]
                },
            }
        },
        methods: {
          dialogClose() {
            this.$emit('update:visible', false)  // 当对话框关闭，通过父组件更新为false
          },
          submit() {
            this.$refs.formRef.validate((valid) => {
              if (valid) {
                this.$http.put('/config/notify/' + this.row.id + '/', this.row)
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
        }
    }
</script>

<style scoped>

</style>