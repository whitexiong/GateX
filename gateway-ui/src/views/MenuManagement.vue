<template>
  <div>
    <ATable :data="menus" :searchQuery="searchQuery">
      <el-table-column
          prop="name"
          label="菜单名称">
      </el-table-column>
      <el-table-column
          label="操作"
          width="180">
        <template #default="scope">
          <button @click="editMenu(scope.row.id)">编辑</button>
          <button @click="deleteMenu(scope.row.id)">删除</button>
        </template>
      </el-table-column>
    </ATable>
  </div>
</template>

<script>
import { getAllMenus, deleteMenu } from '@/services/menuService';
import ATable from "@/components/ATable.vue";

export default {
  components: {ATable},
  data() {
    return {
      menus: []
    };
  },
  methods: {
    async fetchMenus() {
      try {
        const response = await getAllMenus();
        this.menus = Array.isArray(response.data) ? response.data : [];
      } catch (error) {
        console.error("Error fetching menus:", error);
      }
    },
    editMenu(id) {
      // 对应的编辑逻辑
    },
    async deleteMenu(id) {
      try {
        await deleteMenu(id);
        // 删除成功后，重新获取菜单列表
        await this.fetchMenus();
      } catch (error) {
        console.error("Error deleting menu:", error);
      }
    }
  },
  mounted() {
    this.fetchMenus();
  }
}
</script>

<style scoped>
/* 你可以在这里添加样式 */
</style>
