<template>
    <div class="workflow">
        <el-row>
            <el-col :span="24">
                <div>
                    <el-card class="workflow-head-card" shadow="never" :body-style="{padding:'10px'}">
                        <el-row>
                            <el-col :span="6">
                                <div>
                                    <span>命名空间: </span>
                                    <el-select v-model="namespaceValue" filterable placeholder="请选择">
                                        <el-option
                                        v-for="(item, index) in namespaceList"
                                        :key="index"
                                        :label="item.metadata.name"
                                        :value="item.metadata.name">
                                        </el-option>
                                    </el-select>
                                </div>
                            </el-col>
                            <el-col :span="2" :offset="16">
                                <div>
                                    <el-button style="border-radius:2px;" icon="Refresh" plain @click="getWorkflows()">刷新</el-button>
                                </div>
                            </el-col>
                        </el-row>
                    </el-card>
                </div>
            </el-col>
            <el-col :span="24">
                <div>
                    <el-card class="workflow-head-card" shadow="never" :body-style="{padding:'30px 10px 20px 10px'}">
                        <el-steps :active="active" align-center finish-status="success">
                            <el-step title="步骤1" description="选择工作流类型, ClusterIP NodePort Workflow"></el-step>
                            <el-step title="步骤2" description="填写Deployment Workflow Workflow表单"></el-step>
                            <el-step title="步骤3" description="创建Deployment Workflow Workflow"></el-step>
                        </el-steps>
                    </el-card>
                </div>
            </el-col>
            <el-col :span="24">
                <div>
                    <el-card class="workflow-head-card" shadow="never" :body-style="{padding:'10px'}">
                        <el-row>
                            <el-col :span="3">
                                <div>
                                    <el-button style="border-radius:2px;" icon="Edit" type="primary" @click="createWorkflowDrawerIndex1 = true" v-loading.fullscreen.lock="fullscreenLoading">创建工作流</el-button>
                                </div>
                            </el-col>
                            <el-col :span="6">
                                <div>
                                    <el-input class="workflow-head-search" clearable placeholder="请输入" v-model="searchInput"></el-input>
                                    <el-button style="border-radius:2px;" icon="Search" type="primary" plain @click="getWorkflows()">搜索</el-button>
                                </div>
                            </el-col>
                        </el-row>
                    </el-card>
                </div>
            </el-col>
            <el-col :span="24">
                <div>
                    <el-card class="workflow-body-card" shadow="never" :body-style="{padding:'5px'}">
                        <el-table
                        style="width:100%;font-size:12px;margin-bottom:10px;"
                        :data="workflowList"
                        v-loading="appLoading">
                            <el-table-column width="20"></el-table-column>
                            <el-table-column align=left label="Workflow名">
                                <template v-slot="scope">
                                    <a class="workflow-body-workflowname">{{ scope.row.metadata.name }}</a>
                                </template>
                            </el-table-column>
                            <el-table-column align=center label="标签">
                                <template v-slot="scope">
                                    <div v-for="(val, key) in scope.row.metadata.labels" :key="key">
                                        <el-popover
                                            placement="right"
                                            :width="200"
                                            trigger="hover"
                                            :content="key + ':' + val">
                                            <template #reference>
                                                <el-tag style="margin-bottom: 5px" type="warning">{{ ellipsis(key + ":" + val) }}</el-tag>
                                            </template>
                                        </el-popover>
                                    </div>
                                </template>
                            </el-table-column>
                            <el-table-column align=center label="容器组">
                                <template v-slot="scope">
                                    <span>{{ scope.row.status.availableReplicas>0?scope.row.status.availableReplicas:0  }} / {{ scope.row.spec.replicas>0?scope.row.spec.replicas:0 }} </span>
                                </template>
                            </el-table-column>
                            <el-table-column align=center min-width="100" label="创建时间">
                                <template v-slot="scope">
                                    <el-tag type="info">{{ timeTrans(scope.row.metadata.creationTimestamp) }} </el-tag>
                                </template>
                            </el-table-column>
                            <el-table-column align=center label="镜像">
                                <template v-slot="scope">
                                    <div v-for="(val, key) in scope.row.spec.template.spec.containers" :key="key">
                                        <el-popover
                                            placement="right"
                                            :width="200"
                                            trigger="hover"
                                            :content="val.image">
                                            <template #reference>
                                                <el-tag style="margin-bottom: 5px">{{ ellipsis(val.image.split('/')[2]==undefined?val.image:val.image.split('/')[2]) }}</el-tag>
                                            </template>
                                        </el-popover>
                                    </div>
                                </template>
                            </el-table-column>
                            <el-table-column align=center label="操作" width="400">
                                <template v-slot="scope">
                                    <el-button size="small" style="border-radius:2px;" icon="Edit" type="primary" plain @click="getWorkflowDetail(scope)">YAML</el-button>
                                    <el-button size="small" style="border-radius:2px;" icon="Plus" type="primary" @click="handleScale(scope)">扩缩</el-button>
                                    <el-button size="small" style="border-radius:2px;" icon="RefreshLeft" type="primary" @click="handleConfirm(scope, '重启', restartWorkflow)">重启</el-button>
                                    <el-button size="small" style="border-radius:2px;" icon="Delete" type="danger" @click="handleConfirm(scope, '删除', delWorkflow)">删除</el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                        <el-pagination
                        class="workflow-body-pagination"
                        background
                        @size-change="handleSizeChange"
                        @current-change="handleCurrentChange"
                        :current-page="currentPage"
                        :page-sizes="pagesizeList"
                        :page-size="pagesize"
                        layout="total, sizes, prev, pager, next, jumper"
                        :prev-click="getWorkflows"
                        :total="workflowTotal">
                        </el-pagination>
                    </el-card>
                </div>
            </el-col>
        </el-row>
    <el-drawer
        v-model="createWorkflowDrawerIndex1"
        :direction="direction"
        :before-close="handleClose">
        <template #title>
            <h4>创建Workflow-步骤1</h4>
        </template>
        <template #default>
            <el-row type="flex" justify="center">
                <el-col :span="20">
                    <el-form label-width="80px">
                        <el-form-item class="workflow-create-form" label="类型" prop="name">
                            <el-radio v-model="createWorkflow.type" label="ClusterIP">ClusterIP</el-radio>
                            <el-radio v-model="createWorkflow.type" label="NodePort">NodePort</el-radio>
                            <el-radio v-model="createWorkflow.type" label="Ingress">Ingress</el-radio>
                        </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
        </template>
        <template #footer>
            <el-button @click="drawerCancel('createWorkflowDrawerIndex1')">取消</el-button>
            <el-button type="primary" @click="workflowIndex1Next()">下一步</el-button>
        </template>
    </el-drawer>
    <el-drawer
        v-model="createWorkflowDrawerIndex2_1"
        :direction="direction"
        :before-close="handleClose">
        <template #title>
            <h4>创建Workflow-步骤2</h4>
        </template>
        <template #default>
            <el-row type="flex" justify="center">
                <el-col :span="20">
                    <el-form ref="createWorkflow" :rules="createWorkflowRules" :model="createWorkflow" label-width="80px">
                        <h4 style="margin-bottom:10px">Deployment</h4>
                        <el-form-item class="workflow-create-form" label="名称" prop="name">
                            <el-input v-model="createWorkflow.name"></el-input>
                        </el-form-item>
                        <el-form-item class="workflow-create-form" label="命名空间" prop="namespace">
                            <el-select v-model="createWorkflow.namespace" filterable placeholder="请选择">
                                <el-option
                                v-for="(item, index) in namespaceList"
                                :key="index"
                                :label="item.metadata.name"
                                :value="item.metadata.name">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item class="workflow-create-form" label="副本数" prop="replicas">
                            <el-input-number v-model="createWorkflow.replicas" :min="1" :max="10"></el-input-number>
                                <el-popover
                                    placement="top"
                                    :width="100"
                                    trigger="hover"
                                    content="申请副本数上限为10个">
                                    <template #reference>
                                        <el-icon style="width:2em;font-size:18px;color:#4795EE"><WarningFilled/></el-icon>
                                    </template>
                                </el-popover>
                        </el-form-item>
                        <el-form-item class="workflow-create-form" label="镜像" prop="image">
                            <el-input v-model="createWorkflow.image"></el-input>
                        </el-form-item>
                        <el-form-item class="workflow-create-form" label="标签" prop="label_str">
                            <el-input v-model="createWorkflow.label_str" placeholder="示例: project=ms,app=gateway"></el-input>
                        </el-form-item>
                        <el-form-item class="workflow-create-form" label="资源配额" prop="resource">
                            <el-select v-model="createWorkflow.resource" placeholder="请选择">
                                <el-option value="0.5/1" label="0.5C1G"></el-option>
                                <el-option value="1/2" label="1C2G"></el-option>
                                <el-option value="2/4" label="2C4G"></el-option>
                                <el-option value="4/8" label="4C8G"></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item class="workflow-create-form" label="容器端口" prop="container_port">
                            <el-input v-model="createWorkflow.container_port" placeholder="示例: 80"></el-input>
                        </el-form-item>
                        <el-form-item class="workflow-create-form" label="健康检查" prop="health">
                            <el-switch v-model="createWorkflow.health_check" />
                        </el-form-item>
                        <el-form-item class="workflow-create-form" label="检查路径" prop="healthPath">
                            <el-input v-model="createWorkflow.health_path" placeholder="示例: /health"></el-input>
                        </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
        </template>
        <template #footer>
            <el-button @click="drawerCancel('createWorkflowDrawerIndex2_1')">取消</el-button>
            <el-button type="primary" @click="submitForm('createWorkflow', workflowIndex2_1Next)">下一步</el-button>
        </template>
    </el-drawer>
    <el-drawer
        v-model="createWorkflowDrawerIndex2_2"
        :direction="direction"
        :before-close="handleClose">
        <template #title>
            <h4>创建Workflow-步骤2</h4>
        </template>
        <template #default>
            <el-row type="flex" justify="center">
                <el-col :span="20">
                    <el-form ref="createWorkflow" :rules="createWorkflowRules" :model="createWorkflow" label-width="80px">
                        <h4 style="margin-bottom:10px">Service</h4>
                        <el-form-item class="service-create-form" label="Service端口" prop="port">
                            <el-input v-model="createWorkflow.port" placeholder="示例: 80"></el-input>
                        </el-form-item>
                        <el-form-item v-if="createWorkflow.type == 'NodePort'" class="service-create-form" label="NodePort" prop="node_port">
                            <el-input v-model="createWorkflow.node_port" placeholder="示例: 30001"></el-input>
                        </el-form-item>
                        <el-divider v-if="createWorkflow.type == 'Ingress'"></el-divider>
                        <h4 v-if="createWorkflow.type == 'Ingress'" style="margin-bottom:10px">Ingress</h4>
                        <el-form-item v-if="createWorkflow.type == 'Ingress'" class="deploy-create-form" label="域名" prop="host">
                            <el-input v-model="createWorkflow.host" placeholder="示例: www.example.com"></el-input>
                        </el-form-item>
                        <el-form-item v-if="createWorkflow.type == 'Ingress'" class="ingress-create-form" label="Path" prop="path">
                            <el-input v-model="createWorkflow.path" placeholder="示例: /abc"></el-input>
                        </el-form-item>
                        <el-form-item v-if="createWorkflow.type == 'Ingress'" class="deploy-create-form" label="匹配类型" prop="path_type">
                            <el-select v-model="createWorkflow.path_type" placeholder="请选择">
                                <el-option value="Prefix" label="Prefix"></el-option>
                                <el-option value="Exact" label="Exact"></el-option>
                                <el-option value="ImplementationSpecific" label="ImplementationSpecific"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
        </template>
        <template #footer>
            <el-button @click="drawerCancel('createWorkflowDrawerIndex2_2')">取消</el-button>
            <el-button type="primary" @click="submitForm('createWorkflow', createWorkflowFunc)">立即创建</el-button>
        </template>
    </el-drawer>
    <el-dialog title="副本数调整" v-model="scaleDialog" width="25%">
        <div style="text-align:center">
            <span>实例数: </span>
            <el-input-number :step="1" v-model="scaleNum" :min="0" :max="30" label="描述文字"></el-input-number>
        </div>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="scaleDialog = false">取 消</el-button>
                <el-button type="primary" @click="scaleWorkflow()">更 新</el-button>
            </span>
        </template>
    </el-dialog>
    </div>
</template>

<script>
import common from "../common/Config";
import httpClient from '../../utils/request';
export default {
    data() {
        return {
            //工作流
            active: 0,
            createWorkflowDrawerIndex1: false,
            createWorkflowDrawerIndex2_1: false,
            createWorkflowDrawerIndex2_2: false,
            //分页
            currentPage: 1,
            pagesize: 10,
            pagesizeList: [10, 20, 30],
            //
            searchInput: '',
            namespaceValue: 'default',
            namespaceList: [],
            namespaceListUrl: common.k8sNamespaceList,
            //列表
            appLoading: false,
            workflowList: [],
            workflowTotal: 0,
            getWorkflowsData: {
                url: common.k8sWorkflowList,
                params: {
                    filter_name: '',
                    namespace: '',
                    page: '',
                    limit: '',
                }
            },
            //创建
            fullscreenLoading: false,
            direction: 'rtl',
            createWorkflowDrawer: false,
            createWorkflow: {
                name: '',
                namespace: '',
                replicas: 1,
                image: '',
                resource: '',
                health_check: false,
                health_path: '',
                label_str: '',
                label: {},
                container_port: '',
                type: '',
                port: '',
                node_port: '',
                host: '',
                path: '',
                path_type: ''
            },
            createWorkflowData: {
                url: common.k8sWorkflowCreate,
                params: {}
            },
            createWorkflowRules: {
                name: [{
                    required: true,
                    message: '请填写名称',
                    trigger: 'change'
                }],
                image: [{
                    required: true,
                    message: '请填写镜像',
                    trigger: 'change'
                }],
                namespace: [{
                    required: true,
                    message: '请选择命名空间',
                    trigger: 'change'
                }],
                resource: [{
                    required: true,
                    message: '请选择配额',
                    trigger: 'change'
                }],
                label_str: [{
                    required: true,
                    message: '请填写标签',
                    trigger: 'change'
                }],
                container_port: [{
                    required: true,
                    message: '请填写容器端口',
                    trigger: 'change'
                }],
                type: [{
                    required: true,
                    message: '请填写工作流类型',
                    trigger: 'change'
                }],
                port: [{
                    required: true,
                    message: '请填写Workflow端口',
                    trigger: 'change'
                }],
                node_port: [{
                    required: true,
                    message: '请填写NodePort',
                    trigger: 'change'
                }],
                host: [{
                    required: true,
                    message: '请填写域名',
                    trigger: 'change'
                }],
                path: [{
                    required: true,
                    message: '请填写路径',
                    trigger: 'change'
                }],
                path_type: [{
                    required: true,
                    message: '你选择匹配类型',
                    trigger: 'change'
                }],
            },
            //扩缩容
            scaleNum: 0,
            scaleDialog: false,
            scaleWorkflowData: {
                url: common.k8sWorkflowScale,
                params: {
                    workflow_name: '',
                    namespace: '',
                    scale_num: ''
                }
            },
            //重启
            restartWorkflowData: {
                url: common.k8sWorkflowRestart,
                params: {
                    workflow_name: '',
                    namespace: '',
                }
            },
            //删除
            delWorkflowData: {
                url: common.k8sWorkflowDel,
                params: {
                    workflow_name: '',
                    namespace: '',
                }
            },
        }
    },
    methods: {
        handleSizeChange(size) {
            this.pagesize = size;
            this.getWorkflows()
        },
        handleCurrentChange(currentPage) {
            this.currentPage = currentPage;
            this.getWorkflows()
        },
        handleClose(done) {
            this.$confirm('确认关闭？')
            .then(() => {
                done();
            })
            .catch(() => {});
            this.active = 0
        },
        drawerCancel(drawerName) {
            switch (drawerName) {
                case 'createWorkflowDrawerIndex1':
                    this.createWorkflowDrawerIndex1 = false
                    break
                case 'createWorkflowDrawerIndex2_1':
                    this.createWorkflowDrawerIndex2_1 = false
                    break
                case 'createWorkflowDrawerIndex2_2':
                    this.createWorkflowDrawerIndex2_2 = false
            }
            this.active = 0
        },
        ellipsis(value) {
            return value.length>15?value.substring(0,15)+'...':value
        },
        timeTrans(timestamp) {
            let date = new Date(new Date(timestamp).getTime() + 8 * 3600 * 1000)
            date = date.toJSON();
            date = date.substring(0, 19).replace('T', ' ')
            return date 
        },
        getNamespaces() {
            httpClient.get(this.namespaceListUrl)
            .then(res => {
                this.namespaceList = res.data.items
            })
            .catch(res => {
                this.$message.error({
                message: res.msg
                })
            })
        },
        getWorkflows() {
            this.appLoading = true
            this.getWorkflowsData.params.filter_name = this.searchInput
            this.getWorkflowsData.params.namespace = this.namespaceValue
            this.getWorkflowsData.params.page = this.currentPage
            this.getWorkflowsData.params.limit = this.pagesize
            httpClient.get(this.getWorkflowsData.url, {params: this.getWorkflowsData.params})
            .then(res => {
                this.workflowList = res.data.items
                this.workflowTotal = res.data.total
            })
            .catch(res => {
                this.$message.error({
                message: res.msg
                })
            })
            this.appLoading = false
        },
        handleScale(e) {
            this.scaleDialog = true
            this.workflowDetail = e.row
            this.scaleNum = e.row.spec.replicas
        },
        scaleWorkflow() {
            this.scaleWorkflowData.params.workflow_name = this.workflowDetail.metadata.name
            this.scaleWorkflowData.params.namespace = this.namespaceValue
            this.scaleWorkflowData.params.scale_num = this.scaleNum
            httpClient.put(this.scaleWorkflowData.url, this.scaleWorkflowData.params)
            .then(res => {
                this.getWorkflows()
                this.$message.success({
                message: res.msg
                })
            })
            .catch(res => {
                this.$message.error({
                message: res.msg
                })
            })
            this.scaleDialog = false
        },
        restartWorkflow(e) {
            this.restartWorkflowData.params.workflow_name = e.row.metadata.name
            this.restartWorkflowData.params.namespace = this.namespaceValue
            httpClient.put(this.restartWorkflowData.url, this.restartWorkflowData.params)
            .then(res => {
                this.getWorkflows()
                this.$message.success({
                message: res.msg
                })
            })
            .catch(res => {
                this.$message.error({
                message: res.msg
                })
            })
        },
        delWorkflow(e) {
            this.delWorkflowData.params.workflow_name = e.row.metadata.name
            this.delWorkflowData.params.namespace = this.namespaceValue
            httpClient.delete(this.delWorkflowData.url, {data: this.delWorkflowData.params})
            .then(res => {
                this.getWorkflows()
                this.$message.success({
                message: res.msg
                })
            })
            .catch(res => {
                this.$message.error({
                message: res.msg
                })
            })
        },
        handleConfirm(obj, operateName, fn) {
            this.confirmContent = '确认继续 ' + operateName + ' 操作吗？'
            this.$confirm(this.confirmContent,'提示',{
                confirmButtonText: '确定',
                cancelButtonText: '取消',
            })
            .then(() => {
                fn(obj)
                })
            .catch(() => {
                this.$message.info({
                    message: '已取消操作'
                })          
            })
        },
        createWorkflowFunc() {
            //验证标签
            let reg = new RegExp("(^[A-Za-z]+=[A-Za-z0-9]+).*")
            if (!reg.test(this.createWorkflow.label_str)) {
                this.$message.warning({
                    message: "标签填写异常，请确认后重新填写"
                })
                return
            }
            this.fullscreenLoading = true
            //处理标签
            let label = new Map()
            let cpu, memory
            let a = (this.createWorkflow.label_str).split(",")
            a.forEach(item => {
                let b = item.split("=")
                label[b[0]] = b[1]
            })
            //处理配额
            let resourceList = this.createWorkflow.resource.split("/")
            cpu = resourceList[0]
            memory = resourceList[1] + "Gi"
            //
            this.createWorkflowData.params = this.createWorkflow
            this.createWorkflowData.params.label = label
            this.createWorkflowData.params.cpu = cpu
            this.createWorkflowData.params.memory = memory
            this.createWorkflowData.params.container_port = parseInt(this.createWorkflow.container_port)
            this.createWorkflowData.params.port = parseInt(this.createWorkflow.port)
            this.createWorkflowData.params.node_port = parseInt(this.createWorkflow.node_port)
            //处理Hosts
            if (this.createWorkflow.type == 'Ingress') {
                let hosts = new Map()
                let httpPaths = []
                let httpPath = {
                    path: this.createWorkflow.path,
                    path_type: this.createWorkflow.path_type,
                    service_name: this.createWorkflow.name,
                    service_port: parseInt(this.createWorkflow.port)
                }
                httpPaths.push(httpPath)
                hosts[this.createWorkflow.host] = httpPaths
                this.createWorkflowData.params.hosts = hosts
            }
            httpClient.post(this.createWorkflowData.url, this.createWorkflowData.params)
            .then(res => {
                this.$message.success({
                message: res.msg
                })
                //this.getWorkflows()
            })
            .catch(res => {
                this.$message.error({
                message: res.msg
                })
            })
            this.createWorkflow = {}
            this.fullscreenLoading = false
            this.createWorkflowDrawerIndex2_2 = false
            this.active = 3
        },
        submitForm(formName, fn) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    fn()
                } else {
                    return false;
                }
            })
        },
        workflowIndex1Next() {
            if (!this.createWorkflow.type) {
                this.$message.warning({
                    message: "请选择工作流类型"
                })
                return
            }
            this.createWorkflowDrawerIndex1 = false
            this.createWorkflowDrawerIndex2_1 = true
            this.active = 1
        },
        workflowIndex2_1Next() {
            this.createWorkflowDrawerIndex2_1 = false
            this.createWorkflowDrawerIndex2_2 = true
        }
    },
    watch: {
        namespaceValue: {
            handler() {
                localStorage.setItem('namespace', this.namespaceValue)
                this.currentPage = 1
                this.getWorkflows()
            }
        },
    },
    beforeMount() {
        this.namespaceValue = localStorage.getItem('namespace')
        this.getNamespaces()
        //this.getWorkflows()
    }
}
</script>


<style scoped>
    .workflow-head-card,.workflow-body-card {
        border-radius: 1px;
        margin-bottom: 5px;
    }
    .workflow-head-search {
        width:160px;
        margin-right:10px; 
    }
    .workflow-body-workflowname {
        color: #4795EE;
    }
    .workflow-body-workflowname:hover {
        color: rgb(84, 138, 238);
        cursor: pointer;
        font-weight: bold;
    }
    /deep/ .el-drawer__header {
        margin-bottom: 0px !important;
    }
    /deep/ .el-drawer__body {
        padding: 0px 0px 0px 0px;
    }
</style>