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
            @keyup.enter="fetchRoutes">
          <template #append>
            <el-button @click="fetchRoutes" style="margin-right: 5px;">
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
            @confirm="handleAddRoute"
        >
          <el-form ref="RouteForm" :model="Route" label-width="80px" style="width: 100%;">
            <!-- 父菜单 -->
            <el-form-item label="上级节点" style="width: 100%;">
              <el-cascader
                  v-model="Route.parent_id"
                  :options="RouteOptions"
                  placeholder="节点名称"
                  @change="onRouteSelected"
                  style="width: 100%;"
              ></el-cascader>
            </el-form-item>

            <!-- 名称 -->
            <el-form-item label="节点">
              <el-input v-model="Route.name" placeholder="请输入名称"></el-input>
            </el-form-item>

            <!-- 路径 -->
            <el-form-item label="路径">
              <el-input v-model="Route.path" placeholder="请输入路径"></el-input>
            </el-form-item>

            <!-- 状态 -->
            <el-form-item label="状态">
              <el-switch
                  v-model="Route.status"
                  active-value="1"
                  inactive-value="0"
              ></el-switch>
            </el-form-item>

          </el-form>
        </ADialog>
      </div>
    </div>

    <!-- 表格区域 -->
    <el-table :data="Routes" row-key="id" lazy :load="loadRoutes" style="width: 1980px; height: 1000px">

      <el-table-column label="节点">
        <template #default="scope">
          <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">{{ scope.row.path }}</span>
        </template>
      </el-table-column>

      <!-- 菜单名称列 -->
      <el-table-column label="路由名称">
        <template #default="scope">
          <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">{{ scope.row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column label="操作" width="180">
        <template #default="{ row }">
          <el-button size="mini" @click="toggleStatus(row)">
            {{ row.Status === 1 ? '禁用' : '开启' }}
          </el-button>
        </template>
      </el-table-column>

      <!-- 操作列 -->
      <el-table-column label="操作" width="180">
        <template #default="scope">
      <span>
        <el-button type="primary" size="small" @click="editRoute(scope.row.id)" style="color: black;">编辑</el-button>
        <el-button type="danger" size="small" @click="deleteRoute(scope.row.id)" style="margin-left: 10px; color: black;">删除</el-button>
      </span>
        </template>
      </el-table-column>
    </el-table>

  </div>
</template>

<script>
import {ref, onMounted, computed} from 'vue';
import { getList } from '@/services/routeService';
import {Plus, Refresh, RefreshRight, Search} from "@element-plus/icons-vue";
import ADialog from '@/components/ADialog.vue';
import * as icons from '@element-plus/icons';

export default {
  components: {Refresh, Search, Plus, RefreshRight,ADialog},
  setup() {
    const Routes = ref([]);
    const searchText = ref('');
    const dialogVisible = ref(false);
    const allIcons = Object.keys(icons);  // 获取所有图标的名字


    const iconDialogVisible = ref(false);
    const RouteOptions = ref([]);

    const transformRouteToCascader = (Route) => {
      return {
        value: Route.id,
        label: Route.name,
        children: Route.children && Route.children.length
            ? Route.children.map(child => transformRouteToCascader(child))
            : null
      };
    };

    const onRouteSelected = (value) => {
      Route.value.parent_id = value[value.length - 1]; // 获取最后一个ID作为parent_id
    };


    onMounted(async () => {
      await fetchRoutes();
      RouteOptions.value = Routes.value.map(Route => transformRouteToCascader(Route));
    });

    const getIconComponent = (icon) => {
      return icons[icon];
    };
    const openIconSelector = () => {
      iconDialogVisible.value = true;
    };

    const displayIcon = computed({
      get: () => selectedIcon.value,
      set: (value) => {
        selectedIcon.value = value;
        iconDialogVisible.value = false;
      }
    });

    const selectIcon = (icon) => {
      displayIcon.value = icon;
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
    const Route = ref({
      name: '',       // 菜单名称
    });
    const handleAddRoute = () => {
      console.log('Route data:', Route.value);
    };

    const loadRoutes = (row, treeNode, resolve) => {
      // 因为我们已经处理了所有菜单，所以只需从现有的row中提取子菜单
      if (row.children && row.children.length > 0) {
        return resolve(row.children);
      }
      return resolve([]);
    };

    const editRoute = (id) => {
      // 编辑逻辑
    };

    const addNew = () => {
      dialogVisible.value = true;
    };

    const deleteRoute = async (id) => {
      try {
        await deleteRoute(id);
        await fetchRoutes();
      } catch (error) {
        console.error("Error deleting Route:", error);
      }
    };

    const fetchRoutes = async () => {
      try {
        const response = await getList({ name: searchText.value });
        const fetchedRoutes = Array.isArray(response.data.routes) ? response.data.routes : [];
        Routes.value = processRoutes(fetchedRoutes, 0);
      } catch (error) {
        console.error("Error fetching Routes:", error);
      }
    };

    const processRoutes = (Routes, indent) => {
      return Routes.map(Route => {
        Route.hasChildren = Route.children && Route.children.length > 0;
        Route._indent = 0; // 固定缩进为 0
        if (Route.hasChildren) {
          Route.children = processRoutes(Route.children, indent + 1);
        }
        return Route;
      });
    };

    const refresh = () => {
      searchText.value = '';  // 清空搜索内容
      fetchRoutes();           // 重新加载菜单
    };

    onMounted(fetchRoutes);

    return {
      Routes,
      searchText,
      loadRoutes,
      editRoute,
      deleteRoute,
      fetchRoutes,
      refresh,
      addNew,
      Route,
      dialogVisible,
      handleAddRoute,
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
      pageSize,
      displayIcon,
      RouteOptions,        // 新增属性
      onRouteSelected      // 新增方法
    };
  }
};
</script>

<style scoped>

.icon-item > div {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.icon-item span {
  margin-top: 0.5rem;
  font-size: 0.8rem;
  color: rgba(0, 0, 0, 0.7);
}

.wide-cascader .el-input__inner {
  width: 300px !important;
}


</style>
