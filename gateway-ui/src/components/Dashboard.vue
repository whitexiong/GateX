<template>
  <el-container>
    <el-header>
      <h1>网关后台面板</h1>
    </el-header>

    <el-main>
      <!-- 概览统计部分 -->
      <el-row :gutter="20">
        <el-col :span="6">
          <el-card>
            <div class="big-number">100,000</div>
            <div>今日请求</div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card>
            <div class="big-number">0.5%</div>
            <div>错误率</div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card>
            <div class="big-number">150ms</div>
            <div>平均延迟</div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card>
            <div class="big-number">5</div>
            <div>被使用的服务</div>
          </el-card>
        </el-col>
      </el-row>

      <!-- 详细图表部分 -->
      <el-row :gutter="20" style="margin-top: 20px;">
        <el-col :span="12">
          <el-card>
            <template v-slot:header>
              <span>请求趋势</span>
            </template>
            <div id="request-trend" style="height: 300px;"></div>
          </el-card>
        </el-col>
        <el-col :span="12">
          <el-card>
            <template v-slot:header>
              <span>错误率趋势</span>
            </template>
            <div id="error-trend" style="height: 300px;"></div>
          </el-card>
        </el-col>
      </el-row>

      <el-row :gutter="20" style="margin-top: 20px;">
        <el-col :span="24">
          <el-card>
            <template v-slot:header>
              <span>服务使用统计</span>
            </template>
            <div id="service-usage" style="height: 400px;"></div>
          </el-card>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script>
import { onMounted } from 'vue';
import * as echarts from 'echarts';

export default {
  name: 'Dashboard',
  setup() {
    onMounted(() => {
      initCharts();
    });

    const initCharts = () => {
      // 请求趋势图
      const requestTrend = echarts.init(document.getElementById('request-trend'));
      requestTrend.setOption({
        tooltip: {},
        xAxis: {
          data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
        },
        yAxis: {},
        series: [
          {
            name: '请求',
            type: 'line',
            data: [820, 932, 901, 934, 1290, 1330, 1320],
          },
        ],
      });

      // 错误率趋势图
      const errorTrend = echarts.init(document.getElementById('error-trend'));
      errorTrend.setOption({
        tooltip: {},
        xAxis: {
          data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
        },
        yAxis: {
          min: 0,
          max: 100,
          axisLabel: {
            formatter: '{value} %'
          }
        },
        series: [
          {
            name: '错误率',
            type: 'line',
            data: [1.2, 0.8, 0.6, 1.4, 1.3, 0.9, 0.7],
          },
        ],
      });

      // 服务使用统计图
      const serviceUsage = echarts.init(document.getElementById('service-usage'));
      serviceUsage.setOption({
        tooltip: {},
        xAxis: {
          data: ['服务A', '服务B', '服务C', '服务D', '服务E'],
        },
        yAxis: {},
        series: [
          {
            name: '使用次数',
            type: 'bar',
            data: [500, 600, 700, 800, 900],
          },
        ],
      });
    };

    return {};
  },
};

</script>


<style scoped>
.el-header {
  background-color: #b3c0d1;
  line-height: 60px;
  padding-left: 20px;
  color: #333;
}
.big-number {
  font-size: 2rem;
  text-align: center;
  margin-bottom: 10px;
}
</style>
