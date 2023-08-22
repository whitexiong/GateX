<template>
  <div>
    <el-table :data="menus" row-key="id" lazy :load="loadMenus" :expand="isExpand">
      <el-table-column type="expand" width="50">
        <template #default="scope">
          <el-table :data="scope.row.children" row-key="id" lazy :load="loadMenus">
            <el-table-column label="菜单名称">
              <template #default="scope">
                <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">{{ scope.row.name }}</span>
              </template>
            </el-table-column>
            <el-table-column label="操作" width="180">
              <template #default="scope">
                <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">
                  <button @click="editMenu(scope.row.id)">编辑</button>
                  <button @click="deleteMenu(scope.row.id)">删除</button>
                </span>
              </template>
            </el-table-column>
          </el-table>
        </template>
      </el-table-column>
      <el-table-column label="菜单名称">
        <template #default="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="scope">
          <button @click="editMenu(scope.row.id)">编辑</button>
          <button @click="deleteMenu(scope.row.id)">删除</button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>


<script>
import { getAllMenus, deleteMenu } from '@/services/menuService';

export default {
  data() {
    return {
      menus: []
    };
  },
  methods: {
    isExpand(row, index) {
      // 在这里根据需要判断是否展开当前行
      // 例如，展开第一行：return index === 0;
    },
    async loadMenus(row, treeNode, resolve) {
      if (row.children && row.children.length > 0) {
        return resolve(row.children);
      }

      try {
        const response = await getAllMenus({ parentId: row.id });
        const children = Array.isArray(response.data.menus) ? response.data.menus : [];
        children.forEach(child => {
          child.hasChildren = true;
          child._indent = (treeNode._indent || 0) + 1; // 记录缩进级别
        });
        resolve(children);
      } catch (error) {
        console.error("Error loading menus:", error);
      }
    },
    editMenu(id) {
      // 编辑逻辑
    },
    async deleteMenu(id) {
      try {
        await deleteMenu(id);
        await this.fetchMenus();
      } catch (error) {
        console.error("Error deleting menu:", error);
      }
    },
    async fetchMenus() {
      try {
        const response = await getAllMenus({ parentId: 0 });
        const menus = Array.isArray(response.data.menus) ? response.data.menus : [];
        menus.forEach(menu => {
          menu.hasChildren = true;
          menu._indent = 0; // 顶级菜单缩进为0
        });
        this.menus = menus;
      } catch (error) {
        console.error("Error fetching menus:", error);
      }
    }
  },
  mounted() {
    this.fetchMenus();
  }
};
</script>

<style scoped>
/* 在这里添加样式 */
</style>
