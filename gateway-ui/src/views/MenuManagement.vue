<template>
  <div>
    <!-- 搜索和刷新区域 -->
    <div style="margin-bottom: 20px; display: flex; align-items: center; justify-content: space-between;">

      <!-- 搜索部分 -->
      <div style="display: flex; align-items: center;">
        <el-input
            v-model="searchText"
            placeholder="请输入菜单名称"
            style="width: 200px;"
            @keyup.enter="fetchMenus">
          <template #append>
            <el-button @click="fetchMenus"><el-icon ><Search /></el-icon></el-button>
            <el-button @click="resetSearch"><el-icon ><Refresh /></el-icon></el-button>
          </template>
        </el-input>
      </div>

      <!-- 操作部分 -->
      <div style="display: flex; align-items: center;">
        <el-button @click="refresh">
          <el-icon><RefreshRight /></el-icon>
        </el-button>

        <el-button @click="addNew" style="margin-left: 10px;">
          <el-icon><Plus /></el-icon>新增
        </el-button>
      </div>

    </div>

    <!-- 表格区域 -->
    <el-table :data="menus" row-key="id" lazy :load="loadMenus" style="width: 1980px; height: 1000px">
      <el-table-column label="菜单名称">
        <template #default="scope">
          <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="图标">
        <template #default="scope">
          <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">{{ scope.row.icon }}</span>
        </template>
      </el-table-column>
      <el-table-column label="菜单链接">
        <template #default="scope">
          <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">{{ scope.row.path }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="scope">
          <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">
            <el-button type="primary" size="small" @click="editMenu(scope.row.id)" style="color: black;">编辑</el-button>
            <el-button type="danger" size="small" @click="deleteMenu(scope.row.id)" style="margin-left: 10px; color: black;">删除</el-button>
          </span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { getAllMenus, deleteMenu } from '@/services/menuService';
import {Plus, Refresh, RefreshRight, Search} from "@element-plus/icons-vue";

export default {
  components: {Refresh, Search, Plus, RefreshRight},
  setup() {
    const menus = ref([]);
    const searchText = ref('');

    const loadMenus = async (row, treeNode, resolve) => {
      if (row.children && row.children.length > 0) {
        return resolve(row.children);
      }
      try {
        const response = await getAllMenus({ parentId: row.id });
        const children = Array.isArray(response.data.menus) ? response.data.menus : [];
        children.forEach(child => {
          child.hasChildren = child.children && child.children.length > 0;
          child._indent = (treeNode._indent || 0) + 1; // 记录缩进级别
        });
        resolve(children);
      } catch (error) {
        console.error("Error loading menus:", error);
      }
    };

    const editMenu = (id) => {
      // 编辑逻辑
    };

    const deleteMenu = async (id) => {
      try {
        await deleteMenu(id);
        await fetchMenus();
      } catch (error) {
        console.error("Error deleting menu:", error);
      }
    };

    const fetchMenus = async () => {
      try {
        const response = await getAllMenus({ parentId: 0, name: searchText.value });
        const fetchedMenus = Array.isArray(response.data.menus) ? response.data.menus : [];
        fetchedMenus.forEach(menu => {
          menu.hasChildren = menu.children && menu.children.length > 0;
          menu._indent = 0; // 顶级菜单缩进为0
        });
        menus.value = fetchedMenus;
      } catch (error) {
        console.error("Error fetching menus:", error);
      }
    };

    const refresh = () => {
      searchText.value = '';  // 清空搜索内容
      fetchMenus();           // 重新加载菜单
    };

    onMounted(fetchMenus);

    return {
      menus,
      searchText,
      loadMenus,
      editMenu,
      deleteMenu,
      fetchMenus,
      refresh
    };
  }
};
</script>

<style scoped>
/* 在这里添加你想要的样式，例如对搜索框和按钮的定制样式 */
</style>
