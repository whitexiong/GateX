<template>
  <div class="breadcrumb-container">
    <el-breadcrumb separator="/">
      <el-breadcrumb-item
          v-for="item in breadcrumbs"
          :key="item.path"
          :to="{ path: item.path }">
        {{ item.name }}
      </el-breadcrumb-item>
    </el-breadcrumb>
  </div>
</template>

<script>
export default {
  name: "BreadcrumbNav",

  computed: {
    breadcrumbs() {
      const routePath = this.$route.path;
      const allData = JSON.parse(localStorage.getItem('userMenus') || '[]');
      const breadcrumbs = [];

      this.findBreadcrumbPath(routePath, allData, breadcrumbs);

      return breadcrumbs.reverse();
    }
  },

  methods: {
    findBreadcrumbPath(targetPath, data, result) {
      for (let i = 0; i < data.length; i++) {
        if (targetPath.includes(data[i].path)) {
          result.push({ name: data[i].name, path: data[i].path });

          // 如果当前节点匹配，直接返回true
          return true;
        }
        if (data[i].children && data[i].children.length && this.findBreadcrumbPath(targetPath, data[i].children, result)) {
          // 如果子节点中找到了匹配项，将当前节点也添加到结果中
          result.push({ name: data[i].name, path: data[i].path });
          return true;
        }
      }
      return false;
    }
  }
};
</script>

<style scoped>
.breadcrumb-container {
  padding: 20px 30px;
  border-bottom: 1px solid #eaeaea;
}
</style>
