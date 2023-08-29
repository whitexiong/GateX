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
            <el-button @click="fetchMenus" style="margin-right: 5px;">
              <el-icon><Search /></el-icon>
            </el-button>
            <el-button @click="resetSearch">
              <el-icon><Refresh /></el-icon>
            </el-button>
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
        <ADialog
            v-model="dialogVisible"
            title="新增菜单"
            @confirm="handleAddMenu"
        >
          <el-form ref="menuForm" :model="menu" label-width="80px">
            <!-- 父菜单 -->
            <el-form-item label="父菜单">
              <el-cascader
                  v-model="menu.parent_id"
                  :options="menuOptions"
                  placeholder="请选择父菜单"
              ></el-cascader>
            </el-form-item>

            <!-- 名称 -->
            <el-form-item label="名称">
              <el-input v-model="menu.name" placeholder="请输入菜单名称"></el-input>
            </el-form-item>

            <!-- 路径 -->
            <el-form-item label="路径">
              <el-input v-model="menu.path" placeholder="请输入菜单路径"></el-input>
            </el-form-item>

            <!-- 图标选择 -->
            <el-form-item label="图标">
              <el-button @click="openIconSelector">选择图标</el-button>
              <el-input v-model="selectedIcon" readonly></el-input>
              <el-dialog v-model="iconDialogVisible" width="70%">
                <div class="icon-grid">
                  <div v-for="icon in paginatedIcons" :key="icon">
                    <div @click="selectIcon(icon)" class="icon-container">
                      <component :is="getIconComponent(icon)" class="selectable-icon" />
                      <span>{{ icon }}</span>
                    </div>
                  </div>
                </div>
                <el-pagination
                    @current-change="handlePageChange"
                    :current-page="currentPage"
                    :page-size="pageSize"
                    layout="prev, pager, next"
                    :total="allIcons.length"
                >
                </el-pagination>
              </el-dialog>
            </el-form-item>

            <!-- 权限标识 -->
            <el-form-item label="权限标识">
              <el-input v-model="menu.permission" placeholder="请输入权限标识"></el-input>
            </el-form-item>

            <!-- 排序 -->
            <el-form-item label="排序">
              <el-input-number v-model="menu.order" :min="0"></el-input-number>
            </el-form-item>
          </el-form>
        </ADialog>
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
import {ref, onMounted, computed} from 'vue';
import { getAllMenus, deleteMenu } from '@/services/menuService';
import {Plus, Refresh, RefreshRight, Search, Tools} from "@element-plus/icons-vue";
import ADialog from '@/components/ADialog.vue';
import * as icons from '@element-plus/icons';

export default {
  components: {Tools, Refresh, Search, Plus, RefreshRight,ADialog},
  setup() {
    const menus = ref([]);
    const searchText = ref('');
    const dialogVisible = ref(false);
    const allIcons = Object.keys(icons);  // 获取所有图标的名字


    const iconDialogVisible = ref(false);

    const getIconComponent = (icon) => {
      return icons[icon];
    };
    const openIconSelector = () => {
      iconDialogVisible.value = true;
    };
    const selectIcon = (icon) => {
      selectedIcon.value = icon;
      iconDialogVisible.value = false;
    };

    const pageSize = ref(10);  // 每页显示的图标数量
    const currentPage = ref(1); // 当前页数

    // 分页的图标列表
    const paginatedIcons = computed(() => {
      const start = (currentPage.value - 1) * pageSize.value;
      const end = start + pageSize.value;
      return allIcons.slice(start, end);
    });

    // 分页改变时的处理函数
    const handlePageChange = (newPage) => {
      currentPage.value = newPage;
    };


    const selectedIcon = ref("");
    const menu = ref({
      name: '',       // 菜单名称
      icon: '',       // 菜单图标
      path: '',       // 菜单链接
      parent_id: null, // 父菜单ID
      permission: '', // 权限标识
      order: 999      // 排序，默认为999
    });
    const handleAddMenu = () => {
      console.log('Menu data:', menu.value);
    };


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

    const addNew = () => {
      dialogVisible.value = true;
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
      refresh,
      addNew,
      menu,
      dialogVisible,
      handleAddMenu,
      iconDialogVisible,
      openIconSelector,
      icons,
      selectIcon,
      getIconComponent,
      allIcons,
      paginatedIcons,
      currentPage,
      handlePageChange,
      selectedIcon,
      pageSize
    };
  }
};
</script>

<style scoped>
.icon-grid {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 1rem;
}

.icon-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.2s;

  &:hover {
    background-color: rgba(0, 0, 0, 0.05);
  }
}

.selectable-icon {
  font-size: 20px;
  width: 20px;
  height: 20px;
}

.icon-container {
  display: inline-block;
  width: 20%;
  text-align: center;
  padding: 5px;
}


.icon-item component {
  font-size: 20px;
}

.icon-item > div {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.icon-item span {
  margin-top: 0.5rem;
  font-size: 0.8rem;
  color: rgba(0, 0, 0, 0.7);
  text-align: center;
}
</style>
