<template>
  <div class="breadcrumb-container">
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/' }">homepage</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: '/route1' }">route 1</el-breadcrumb-item>
      <el-breadcrumb-item>route 2</el-breadcrumb-item>
    </el-breadcrumb>
  </div>
</template>

<script>
export default {
  name: "BreadcrumbNav"
};
</script>

<style scoped>
/* Add your styles here if needed */
.breadcrumb-container {
  padding: 20px 30px;  /* 根据需要调整 */
  border-bottom: 1px solid #eaeaea;
}
</style>
