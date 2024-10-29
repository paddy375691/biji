<template>

  <el-card class="box-card">
    <div class="table-header">
      <div>
         <el-row>
            <el-col :span="7">
                <el-input
                  v-model="urlParams.search"
                  placeholder="搜索名称，主机名，IP"
                  @keyup.enter="onSearch"
                  clearable
                  @clear="onSearch"
                  class="search"
                />
            </el-col>
             <!--选择IDC-->
            <el-col :span="7">
              <el-select v-model="urlParams.idc" class="m-2" @click="getIdc" clearable placeholder="IDC机房" @clear="onSearch">
                <el-option v-for="row in idc" :key="row.id" :label="row.name" :value="row.id">
                </el-option>
              </el-select>
            </el-col>
            <!--选择分组-->
            <el-col :span="7">
              <el-select v-model="urlParams.server_group" class="m-2" @click="getServerGroup" clearable placeholder="主机分组" @clear="onSearch">
                <el-option v-for="row in serverGroup" :key="row.id" :label="row.name" :value="row.id">
                </el-option>
              </el-select>
            </el-col>
            <el-button type="primary" @click="onSearch"><el-icon><search /></el-icon>&nbsp;查询</el-button>
        </el-row>
      </div>
      <div>
        <!--新建主机-->
        <el-dropdown style="margin-right: 12px">
          <el-button type="primary">
            <el-icon><plus /></el-icon>&nbsp;新建<el-icon class="el-icon--right"><arrow-down /></el-icon>
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="createDialogVisible = true"><el-icon color="#409EFC" :size="20"><Edit /></el-icon>新建主机</el-dropdown-item>
              <el-dropdown-item @click="excelDialogVisible = true"><el-icon color="#409EFC" :size="20"><Folder /></el-icon>Excel</el-dropdown-item>
              <el-dropdown-item @click="cloudDialogVisible = true"><el-icon color="#409EFC" :size="21"><MostlyCloudy /></el-icon>云主机</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>

        <!--展示列弹出框-->
        <el-popover placement="left" :width="100" v-model:visible="columnVisible">
          <template #reference>
            <el-button type="primary" @click="columnVisible = true"><el-icon><setting /></el-icon>&nbsp;展示列</el-button>
          </template>
              <el-checkbox v-model="showColumn.name" disabled>名称</el-checkbox>
              <el-checkbox v-model="showColumn.hostname" disabled>主机名</el-checkbox>
              <el-checkbox v-model="showColumn.idc">IDC机房</el-checkbox>
              <el-checkbox v-model="showColumn.server_group">主机分组</el-checkbox>
              <el-checkbox v-model="showColumn.machine_type">机器类型</el-checkbox>
              <el-checkbox v-model="showColumn.os_version">系统版本</el-checkbox>
              <el-checkbox v-model="showColumn.public_ip">公网IP</el-checkbox>
              <el-checkbox v-model="showColumn.private_ip">私有IP</el-checkbox>
              <el-checkbox v-model="showColumn.cpu_num">CPU数量</el-checkbox>
              <el-checkbox v-model="showColumn.cpu_model">CPU型号</el-checkbox>
              <el-checkbox v-model="showColumn.memory">内存</el-checkbox>
              <el-checkbox v-model="showColumn.disk">硬盘</el-checkbox>
              <el-checkbox v-model="showColumn.put_shelves_date">上架日期</el-checkbox>
              <el-checkbox v-model="showColumn.off_shelves_date">下架日期</el-checkbox>
              <el-checkbox v-model="showColumn.expire_datetime">租约过期时间</el-checkbox>
              <el-checkbox v-model="showColumn.is_verified">SSH验证状态</el-checkbox>
              <el-checkbox v-model="showColumn.note">备注</el-checkbox>
              <el-checkbox v-model="showColumn.update_time">更新时间</el-checkbox>
              <el-checkbox v-model="showColumn.create_time">创建时间</el-checkbox>
            <div style="text-align: right; margin: 0">
              <el-button size="small" type="text" @click="columnVisible = false">取消</el-button>
              <el-button size="small" type="primary" @click="saveColumn">确认</el-button>
            </div>
        </el-popover>

      </div>
    </div>

    <!--数据表格-->
    <el-table
            :data="tableData"
            style="width: 100%"
            border
            :header-cell-style="{background: '#F0F2F5'}"
    >
      <el-table-column prop="name" label="名称" width="130" fixed="left" sortable v-if="showColumn.name"/>
      <el-table-column prop="hostname" label="主机名" width="130" sortable v-if="showColumn.hostname"/>
      <el-table-column prop="idc" label="IDC机房" sortable v-if="showColumn.idc">
          <template #default="scope">
            <img src="../../assets/img/aliyun.png" style="width: 18px;height: 18px" v-if="scope.row.idc.provider == '阿里云'">
            <img src="../../assets/img/tencent.png" style="width: 18px;height: 18px" v-else-if="scope.row.idc.provider == '腾讯云'">
            <el-icon :size="18" color="#409EFC" v-else><office-building/></el-icon>
              {{scope.row.idc.city}}-{{scope.row.idc.name}}
          </template>
      </el-table-column>
      <el-table-column prop="server_group" label="主机分组" width="150" sortable v-if="showColumn.server_group">
          <template #default="scope">
              <el-tag type="info" v-for="group in scope.row.server_group" :key="group.id">{{group.name}}</el-tag>
          </template>
      </el-table-column>
      <el-table-column prop="machine_type" label="机器类型" width="110" sortable v-if="showColumn.machine_type">
          <template #default="scope">
              <span v-if="scope.row.machine_type == 'vm'">虚拟机</span>
              <span v-else-if="scope.row.machine_type == 'cloud_vm'">云主机</span>
              <span v-else>物理机</span>
          </template>
      </el-table-column>
      <el-table-column prop="os_version" label="系统版本" sortable :show-overflow-tooltip="true" v-if="showColumn.os_version"/>
      <el-table-column prop="public_ip" label="公网IP" width="180" sortable v-if="showColumn.public_ip">
          <template #default="scope">
              <el-tag type="info" v-for="(ip, index) in scope.row.public_ip" :key="index">{{ip}}</el-tag>
          </template>
      </el-table-column>
      <el-table-column prop="private_ip" label="私有IP" width="180" sortable v-if="showColumn.public_ip">
          <template #default="scope">
              <el-tag type="info" v-for="(ip, index) in scope.row.private_ip" :key="index">{{ip}}</el-tag>
          </template>
      </el-table-column>
      <el-table-column prop="cpu_num" label="CPU" width="80" sortable v-if="showColumn.cpu_num"/>
      <el-table-column prop="cpu_model" label="CPU型号" sortable v-if="showColumn.cpu_model"/>
      <el-table-column prop="memory" label="内存" width="80" sortable v-if="showColumn.memory"/>
      <el-table-column prop="disk" label="硬盘" width="200" sortable v-if="showColumn.disk">
          <template #default="scope">
            <table style="background: #ebeef5;width: 100%" v-if="scope.row.disk">  <!--表格背景设置灰色，表格内默认白色-->
              <thead>
                <tr>
                  <th>设备</th>
                  <th>类型</th>
                  <th>容量</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(disk, index) in scope.row.disk" :key="index">
                  <td>{{disk.device}}</td>
                  <td>{{disk.type}}</td>
                  <td>{{disk.size}}</td>
                </tr>
              </tbody>
            </table>
            <span v-else>{{scope.row.disk}}</span>
          </template>
      </el-table-column>
      <el-table-column prop="put_shelves_date" label="上架日期" sortable v-if="showColumn.put_shelves_date"/>
      <el-table-column prop="off_shelves_date" label="下架日期" sortable v-if="showColumn.off_shelves_date"/>
      <el-table-column prop="expire_datetime" label="租约过期时间" sortable v-if="showColumn.expire_datetime"/>
      <el-table-column prop="is_verified" label="SSH状态" width="120" sortable v-if="showColumn.is_verified">
          <template #default="scope">
              <el-tag type="success" v-if="scope.row.is_verified == 'verified'">已验证</el-tag>
              <el-tag type="warning" v-else-if="scope.row.is_verified == 'unverified'">未验证</el-tag>
          </template>
      </el-table-column>
      <el-table-column prop="note" label="备注"  v-if="showColumn.note"/>
      <el-table-column prop="update_time" label="更新时间" sortable v-if="showColumn.update_time"/>
      <el-table-column prop="create_time" label="创建时间" sortable v-if="showColumn.create_time"/>

      <!--操作栏-->
      <el-table-column fixed="right" label="操作" width="200">
        <template #default="scope">
          <el-button
                  type="primary"
                  size="small"
                  @click="handleEdit(scope.$index, scope.row)"
            >编辑</el-button>
          <el-button
                  type="danger"
                  size="small"
                  @click="handleDel(scope.$index, scope.row)"
          >删除</el-button>
          <el-button
                  type="warning"
                  size="small"
                  @click="handleSync(scope.$index, scope.row)"
          >同步</el-button>
        </template>
      </el-table-column>

    </el-table>

    <!--分页-->
    <div style="margin-top: 20px">
      <el-pagination
        v-model:currentPage="currentPage"
        :page-sizes="[10, 15, 20, 25, 30]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      >
      </el-pagination>
    </div>
  </el-card>

    <!--操作栏：点击同步选择凭据对话框-->
  <el-dialog
    v-model="credentialVisible"
    width="20%"
    title="请选择SSH连接凭据"
    >
    <el-col :span="2">
      <el-icon :size="23"><lock/></el-icon>
    </el-col>
    <el-col :span="22" :offset="1">
      <el-select class="m-2" v-model="credentialId" @click="getCredential" placeholder="请选择">
        <el-option v-for="row in credential" :key="row.id" :label="`${row.name}-${row.username}`" :value="row.id">
        </el-option>
      </el-select>
    </el-col>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="credentialVisible = false">取消</el-button>
        <el-button type="primary" @click="sync">确定</el-button>
      </span>
    </template>
  </el-dialog>

  <ServerEdit v-model:visible="editDialogVisible" :row="row"></ServerEdit>
  <ServerCreate v-model:visible="createDialogVisible"></ServerCreate>
  <ServerCreateExcel v-model:visible="excelDialogVisible"></ServerCreateExcel>
  <ServerCreateCloud v-model:visible="cloudDialogVisible"></ServerCreateCloud>

</template>

<script>
    import ServerEdit from "./ServerEdit";
    import ServerCreate from "./ServerCreate";
    import ServerCreateExcel from "./ServerCreateExcel";
    import ServerCreateCloud from "./ServerCreateCloud";
    export default {
        data() {
          return {
            tableData: [],
            total: 0,  // 总数据条数
            pageSize: 10, // 默认每页的数据条数
            currentPage: 1,  // 默认开始页码
            urlParams: {  // URL查询参数
              page_num: 1,
              page_size: 10,
              search: '',
              idc: '',
              server_group: ''
            },
            editDialogVisible: false,
            row: '',
            createDialogVisible: false,
            showColumn: {
                name: true,
                hostname: true,
                idc: true,
                server_group: true,
                machine_type: false,
                os_version: true,
                public_ip: true,
                private_ip: true,
                cpu_num: true,
                cpu_model: false,
                memory: true,
                disk: true,
                put_shelves_date: false,
                off_shelves_date: false,
                expire_datetime: false,
                is_verified: false,
                update_time: false,
                create_time: false,
                note: false
            },
            columnVisible: false, // 展示与隐藏
            excelDialogVisible: false,
            cloudDialogVisible: false,
            credentialVisible: false,
            idc: '',
            serverGroup: '',
            credential: '',
            credentialId: '',
          }
        },
        mounted() {
          this.getData();
          // 从浏览器本地存储获取历史展示列
          const columnSet = localStorage.getItem(this.$route.path + '-columnSet');
          // 如果本地有存储则使用
          if(columnSet) {
            this.showColumn = JSON.parse(columnSet)
          }
        },
        methods: {
          getData() {
              // 处理URL查询参数
              if(!this.urlParams.search) {
                  delete this.urlParams.search;
              }
              if(!this.urlParams.idc) {
                  delete this.urlParams.idc;
              }
              if(!this.urlParams.server_group) {
                  delete this.urlParams.server_group;
              }
            this.$http.get('/cmdb/server/', {params: this.urlParams})
              .then(res => {
                // console.log(res.data);
                this.tableData = res.data.data;
                this.total = res.data.count;
              })
          },
           // 分页：监听【选择每页数量】的事件
          handleSizeChange(pageSize) {
              // console.log(pageSize)
              this.pageSize = pageSize; // 重新设置分页显示
              this.urlParams.page_size = pageSize;  // 将最新值设置到数据里，用于下面用新值重新获取数据
              this.getData()
          },
          // 分页：监听【点击页码】的事件
          handleCurrentChange(currentPage) {
              // console.log(currentPage)
              this.currentPage = currentPage; // 重新设置分页显示
              this.urlParams.page_num = currentPage;
              this.getData()
          },
          handleEdit(index, row) {
            // console.log(index,row);
            this.editDialogVisible = true;
            this.row = row;
          },
          handleDel(index, row) {
             this.$confirm("你确定要删除选中的吗？", "提示", {
                  confirmButtonText: "确定",
                  cancelButtonText: "取消",
                  type: "warning"
                  })
                  .then(() => {  // 点击确定
                    this.$http.delete('/cmdb/server/'+ row.id + '/')
                        .then(res => {
                          if(res.data.code == 200) {
                            this.$message.success('删除成功');
                            this.tableData.splice(index, 1); // 根据表格索引临时删除数据
                          } else {
                            this.$message.error(res.data.msg);
                          }
                        });
                  })
          },
          // 选择凭据后确认按钮
          sync() {
              if(this.credentialId) {
                  this.$http.get('/cmdb/host_collect/', {params: {
                        'hostname': this.currentRow.hostname,
                        'credential_id': this.credentialId
                    }
                  })
                    .then(res => {
                        if(res.data.code == 200) {
                            this.$message.success('同步成功');
                            this.getData();
                            this.credentialVisible = false;
                        }
                    })
              } else {
                  this.$message.warning('请选择凭据！')
              }
          },
          handleSync(index, row) {
              this.currentRow = row;
              // 如果主机有凭据直接同步
              if(row.credential) {
                  this.$http.get('/cmdb/host_collect/', {params: {'hostname': row.hostname}})
                    .then(res => {
                        if(res.data.code == 200) {
                            this.$message.success('同步成功');
                            this.getData();
                        }
                    })
              } else {
                  this.credentialVisible = true;
              }
          },
          onSearch() {
            this.getData();
          },
          saveColumn() {
            localStorage.setItem(this.$route.path + '-columnSet', JSON.stringify(this.showColumn));
            this.columnVisible = false;
          },
            getIdc() {
              console.log(this.$http.instance);
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
        components: {
          ServerEdit,
          ServerCreate,
          ServerCreateExcel,
          ServerCreateCloud
        }
    }
</script>

<style scoped>
  .table-header {
    display: flex;
    margin-bottom: 18px;
    justify-content: space-between;
  }
</style>