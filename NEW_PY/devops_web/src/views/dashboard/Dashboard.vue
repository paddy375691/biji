<template>
  <el-row :gutter="16">
    <el-col :span="6">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>机房数量</span>
          </div>
        </template>
        <div>
          <el-progress type="circle" :percentage="100" status="success" width="50"/>
          <span class="number">{{ idc_number }}</span>
        </div>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>服务器数量</span>
          </div>
        </template>
        <div>
          <el-progress type="circle" :percentage="100" status="success" width="50"/>
          <span class="number">{{ server_number }}</span>
        </div>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>项目数量</span>
          </div>
        </template>
        <div>
          <el-progress type="circle" :percentage="100" status="success" width="50"/>
          <span class="number">{{ project_number }}</span>
        </div>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>应用数量</span>
          </div>
        </template>
        <div>
          <el-progress type="circle" :percentage="100" status="success" width="50"/>
          <span class="number">{{ app_number }}</span>
        </div>
      </el-card>
    </el-col>
  </el-row>
  <br>
  <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>发布状态</span>
          </div>
        </template>
        <div id="myChart" style="width: 100%;height:400px;"></div>
  </el-card>

</template>

<script>
    import * as echarts from 'echarts'
    export default {
        name: "Dashboard",
        data() {
            return {
                idc_number: 0,
                server_number: 0,
                project_number: 0,
                app_number: 0,
                xData: '',
                yFailData: '',
                ySuccessData: '',
            }
        },
        mounted() {
            this.getNumber();
            this.releaseEchat();
            this.getReleaseData();
        },
        methods: {
            getNumber() {
              this.$http.get('/cmdb/idc/')
                  .then(res => {
                      this.idc_number = res.data.count;
                  });
              this.$http.get('/cmdb/server/')
                  .then(res => {
                      this.server_number = res.data.count;
                  });
              this.$http.get('/app_release/project/')
                  .then(res => {
                      this.project_number = res.data.count;
                  });
              this.$http.get('/app_release/app/')
                  .then(res => {
                      this.app_number = res.data.count;
                  });
            },
            releaseEchat(){
              // 基于准备好的dom，初始化echarts实例
              var myChart = echarts.init(document.getElementById('myChart'));
              // 绘制图表
              var option = {
                // title: {
                //   text: '趋势图'
                // },
                tooltip: {
                  trigger: 'axis'
                },
                legend: {
                  data: ['发布失败', '发布成功']
                },
                grid: {
                  left: '3%',
                  right: '4%',
                  bottom: '3%',
                  containLabel: true
                },
                // toolbox: {
                //   feature: {
                //     saveAsImage: {}
                //   }
                // },
                xAxis: {
                  type: 'category',
                  boundaryGap: false,
                  data: this.xData
                },
                yAxis: {
                  type: 'value'
                },
                series: [
                  {
                    name: '发布失败',
                    type: 'line',
                    stack: 'Total',
                    lineStyle: {color: 'red'},
                    data: this.yFailData
                  },
                  {
                    name: '发布成功',
                    type: 'line',
                    stack: 'Total',
                    lineStyle: {color: 'green'},
                    data: this.ySuccessData
                  }
                ]
              };
              myChart.setOption(option); // 使用刚指定的配置项和数据显示图表。
            },
            getReleaseData() {
                this.$http.get('/app_release/apply_echart/')
                    .then(res => {
                        if(res.data.code == 200) {
                            this.xData = res.data.data.x_data;
                            this.yFailData = res.data.data.y_fail_data;
                            this.ySuccessData = res.data.data.y_success_data;
                            this.releaseEchat();
                        }
                    });
            }
        }
    }
</script>

<style scoped>
    .number {
        margin-left: 40%;
        font-size: 40px;
    }
</style>