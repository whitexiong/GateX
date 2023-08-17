<template>
    <h2 class="text-3xl font-semibold mb-8 text-gray-700">数字营区网关</h2>

    <!-- 使用SCPTable组件呈现数据 -->
    <ATable :data="apis" v-if="apis && apis.length">
      <el-table-column prop="Path" label="路径" class="px-6 py-3"></el-table-column>
      <el-table-column prop="HitCount" label="请求次数" class="px-6 py-3"></el-table-column>
      <el-table-column prop="Type" label="请求类型" class="px-6 py-3"></el-table-column>
      <el-table-column prop="RequestMethod" label="方法" class="px-6 py-3"></el-table-column>
    </ATable>
</template>

<script>
import { getDashboardData } from '@/services/api';
import ATable from "@/components/ATable.vue"; // 根据你的文件结构调整路径

export default {
  name: 'UserDashboard',
  components: {
    ATable
  },
  data() {
    return {
      apis: []
    };
  },
  mounted() {
    // 当组件挂载时，从API获取数据
    getDashboardData()
        .then(response => {
          console.log("API Response", response.data)
          this.apis = response.data.apis;
        })
        .catch(error => {
          console.error("Error fetching dashboard data:", error);
        });
  }
}
</script>

<style scoped>
/* 你可以在这里添加样式 */
</style>
