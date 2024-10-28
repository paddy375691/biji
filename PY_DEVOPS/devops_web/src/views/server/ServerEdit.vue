<template>
  <el-dialog
          :model-value="visible"
          title="修改主机信息"
          @close="dialogClose"
          width="40%"
  >
     <el-form :model="row" ref="formRef" :rules="formRules" label-width="100px">
        <el-form-item label="机器名称：" prop="name">
          <el-input v-model="row.name" placeholder="例如：测试机"></el-input>
        </el-form-item>
        <el-form-item label="主机名称：" prop="hostname">
          <el-input v-model="row.hostname" disabled></el-input>
        </el-form-item>
        <el-form-item label="IDC机房：" prop="idc">
          <el-select v-model="row.idc" @click="getIdc">
              <el-option v-for="row in idc" :key="row.id" :label="`${row.city}-${row.name}`" :value="row.id"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="主机分组：" prop="server_group">
          <el-select v-model="row.server_group" @click="getServerGroup" multiple>
              <el-option v-for="row in serverGroup" :key="row.id" :label="row.name" :value="row.id"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="SSH连接：" required>
          <el-col :span="1.5">
            <el-tag size="large" type="info">IP</el-tag>
          </el-col>
          <el-col :span="10">
            <el-form-item prop="ssh_ip">
              <el-input v-model="row.ssh_ip"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="1.5">
            <el-tag size="large" type="info">端口</el-tag>
          </el-col>
          <el-col :span="5">
            <el-form-item prop="ssh_port">
            <el-input v-model="row.ssh_port"></el-input>
            </el-form-item>
          </el-col>
        </el-form-item>

         <el-form-item label="SSH凭据：" prop="credential">
          <el-select class="m-2" v-model="row.credential" @click="getCredential" placeholder="请选择">
            <el-option v-for="row in credential" :key="row.id" :label="`${row.name}-${row.username}`" :value="row.id">
            </el-option>
          </el-select>
        </el-form-item>

        <br>
        <el-divider content-position="left">
            <el-button @click="serverDetailVisible ? serverDetailVisible = false : serverDetailVisible = true">
                更多详情</el-button>
        </el-divider>

        <div v-show="serverDetailVisible">
            <el-form-item label="机器类型：" prop="machine_type">
              <el-select v-model="row.machine_type" class="m-2" placeholder="请选择">
                <el-option label="虚拟机" value="vm"></el-option>
                <el-option label="云主机" value="cloud_vm"></el-option>
                <el-option label="物理机" value="physical_machine"></el-option>
              </el-select>
            </el-form-item>

            <el-form-item label="系统版本：" prop="os_version">
              <el-input v-model="row.os_version"></el-input>
            </el-form-item>
            <el-form-item label="公网IP：" prop="public_ip">
              <el-input v-model="row.public_ip"></el-input>
            </el-form-item>
            <el-form-item label="私有IP：" prop="private_ip">
              <el-input v-model="row.private_ip"></el-input>
            </el-form-item>
            <el-form-item label="CPU数量：" prop="cpu_num">
              <el-input v-model="row.cpu_num"></el-input>
            </el-form-item>
            <el-form-item label="CPU型号：" prop="cpu_model">
              <el-input v-model="row.cpu_model"></el-input>
            </el-form-item>
            <el-form-item label="内存：" prop="memory">
              <el-input v-model="row.memory"></el-input>
            </el-form-item>
            <el-form-item label="硬盘：" prop="disk">
              <el-input v-model="row.disk"></el-input>
            </el-form-item>
            <el-form-item label="上架日期：" prop="put_shelves_date">
              <el-date-picker v-model="row.put_shelves_date" type="date" value-format="YYYY-MM-DD" placeholder="请选择日期">
              </el-date-picker>
            </el-form-item>
            <el-form-item label="下架日期：" prop="off_shelves_date">
              <el-date-picker v-model="row.off_shelves_date" type="date" value-format="YYYY-MM-DD" placeholder="请选择日期">
              </el-date-picker>
            </el-form-item>
            <el-form-item label="租约过期时间：" prop="expire_datetime">
              <el-date-picker v-model="row.expire_datetime" type="datetime" value-format="YYYY-MM-DD HH:mm:ss" placeholder="请选择时间">
              </el-date-picker>
            </el-form-item>
            <el-form-item label="备注：">
              <el-input v-model="row.note" type="textarea"></el-input>
            </el-form-item>
        </div>
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
                  idc: [
                      {required: true, message: '请选择IDC机房', trigger: 'clear'},
                  ],
                  server_group: [
                      {required: true, message: '请选择主机分组', trigger: 'clear'},
                  ],
                  name: [
                      {required: true, message: '请输入机器名称', trigger: 'c'},
                      {min: 2, message: '主机名长度应不小于2个字符', trigger: 'blur'}
                  ],
                  hostname: [
                      {required: true, message: '请输入主机名称', trigger: 'blur'},
                      {min: 4, message: '主机名长度不小于4个字符', trigger: 'blur'}
                  ],
                  ssh_ip: [
                      {required: true, message: '请输入SSH IP地址', trigger: 'blur'},
                      {min: 7, message: '主机名长度不小于8个字符', trigger: 'blur'}
                  ],
                  ssh_port: [
                      {required: true, message: '请输入SSH端口', trigger: 'blur'},
                  ],
                  credential_id: [
                      {required: true, message: '请选择SSH连接凭据', trigger: 'clear'},
                  ],
                },
                idc: '',
                serverGroup: '',
                credential: '',
                serverDetailVisible: false // 更多详情展示与隐藏
            }
        },
        methods: {
          dialogClose() {
            this.$emit('update:visible', false);  // 当对话框关闭，通过父组件更新为false
            this.$parent.getData();
          },
          submit() {
            this.$refs.formRef.validate((valid) => {
              if (valid) {
                this.$http.put('/cmdb/server/' + this.row.id + '/', this.row)
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
            getIdc() {
                this.$http.get('/cmdb/idc/?page_size=50')
                    .then(res => {
                        this.idc = res.data.data;
                    })
            },
            getServerGroup() {
                this.$http.get('/cmdb/server_group/?page_size=50')
                    .then(res => {
                        this.serverGroup = res.data.data;
                    });
            },
            getCredential() {
                this.$http.get('/config/credential/?page_size=50')
                    .then(res => {
                        this.credential = res.data.data;
                    });
            },

        },
        watch: {
            visible() {
                this.row.idc = this.row.idc.id;  // 从对象中提取ID字段重新赋值
                const group = this.row.server_group;
                this.row.server_group = []; // [1,5,3]
                for(let i in group) {
                    this.row.server_group.push(group[i].id)
                }

                this.getIdc();
                this.getServerGroup();
                this.getCredential();
            }
        }
    }
</script>

<style scoped>

</style>