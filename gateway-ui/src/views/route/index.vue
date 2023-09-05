<template>
  <div>
    <!-- 搜索和刷新区域 -->
    <div style="margin-bottom: 20px; display: flex; align-items: center; justify-content: space-between;">

      <!-- 搜索部分 -->
      <div style="display: flex; align-items: center;">
        <el-input
            v-model="searchText"
            placeholder="请输入节点名称"
            style="width: 200px;"
            @keyup.enter="listData">
          <template #append>
            <el-button @click="listData" style="margin-right: 5px;">
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
            title="新增节点"
            @confirm="saveRoute"
        >
          <el-form ref="RouteForm" :model="Route" label-width="80px" style="width: 100%;">
            <!-- 父节点 -->
            <el-form-item label="上级节点" style="width: 100%;">
              <el-cascader
                  v-model="Route.ParentID"
                  :options="RouteOptions"
                  :props="anyProps"
                  placeholder="节点名称"
                  @change="onRouteSelected"
                  style="width: 100%;"
              ></el-cascader>
            </el-form-item>

            <!-- 名称 -->
            <el-form-item label="节点">
              <el-input v-model="Route.Name" placeholder="请输入名称"></el-input>
            </el-form-item>

            <!-- 路径 -->
            <el-form-item label="路径">
              <el-input v-model="Route.Path" placeholder="请输入路径"></el-input>
            </el-form-item>

            <!-- 状态 -->
            <el-form-item label="状态">
              <el-switch
                  v-model="Route.Status"
                  :active-value="1"
                  :inactive-value="0"
              ></el-switch>

            </el-form-item>

          </el-form>
        </ADialog>
      </div>
    </div>

    <!-- 表格区域 -->
    <el-table :data="Routes" row-key="id" lazy :load="loadRoutes" style="width: 1980px; height: 1000px" border>

      <el-table-column label="节点">
        <template #default="scope">
          <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">{{ scope.row.path }}</span>
        </template>
      </el-table-column>

      <!-- 节点名称列 -->
      <el-table-column label="路由名称">
        <template #default="scope">
          <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">{{ scope.row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column label="操作" width="260">
        <template #default="{ row }">
          <div style="display: flex; align-items: center;">
            <el-button size="small" @click="toggleStatus(row)">
              {{ row.Status === 1 ? '禁用' : '开启' }}
            </el-button>
            <el-button type="primary" size="small" @click="getDetail(row.id)" style="color: black; margin-left: 5px;">编辑</el-button>
            <el-button type="danger" size="small" @click="deleteRoute(row.id)" style="color: black; margin-left: 10px;">删除</el-button>
          </div>
        </template>
      </el-table-column>

    </el-table>

  </div>
</template>

<script>
import {ref, onMounted, computed} from 'vue';
import { getList,add, deleted, update, detail } from '@/services/routeService';
import {Plus, Refresh, RefreshRight, Search} from "@element-plus/icons-vue";
import ADialog from '@/components/ADialog.vue';
import * as icons from '@element-plus/icons';
import {ElMessageBox} from "element-plus";
import { loadList } from '@/composables/useDataLoader';

export default {
  components: {Refresh, Search, Plus, RefreshRight,ADialog},
  setup() {
    const dialogVisible = ref(false);
    const allIcons = Object.keys(icons);  // 获取所有图标的名字
    const pageSize = ref(10);  // 每页显示的图标数量
    const currentPage = ref(1); // 当前页数
    const selectedIcon = ref("");
    const iconDialogVisible = ref(false);
    const RouteOptions = ref([]);
    const {Routes, searchText, listData} = loadList(getList);

    const anyProps = {
      checkStrictly: true,
    }

    const Route = ref({
      ID: null,
      Name: '',
      Status: null,
      ParentID: null,
      Path: null,
    });

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
      Route.value.ParentID = value[value.length - 1]; // 获取最后一个ID作为ParentID
    };

    onMounted(async () => {
      await listData();
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

    const toggleStatus = async (row) => {
      // 这里，你可以调用一个API来改变数据库中的状态或在前端暂时切换状态。
      row.Status = row.Status === 1 ? 0 : 1;
      // 如果你的API要求调用特定的函数来改变状态，你可以在这里添加。
      // 例如: await toggleRouteStatus(row.id, row.Status);
    };

    const selectIcon = (icon) => {
      displayIcon.value = icon;
    };

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

    const handleAddRoute = async () => {
      Route.value.Status = parseInt(Route.value.Status);
      const response = await add(Route.value);

      if (response.code === 200) {
        dialogVisible.value = false;
        await listData();
      } else {
        await ElMessageBox.alert(response.msg, '添加失败', {
          confirmButtonText: 'OK',
          type: 'error'
        });
      }
    };

    const loadRoutes = (row, treeNode, resolve) => {
      // 因为我们已经处理了所有节点，所以只需从现有的row中提取子节点
      if (row.children && row.children.length > 0) {
        return resolve(row.children);
      }
      return resolve([]);
    };

    const saveRoute = async () => {
      if (isEditing.value) {
        await updateRoute();
      } else {
        await handleAddRoute();
      }
    };

    const updateRoute = async () => {
      Route.value.Status = parseInt(Route.value.Status);
      console.log(Route.value)
      const response = await update(Route.value.ID,Route.value);

      if (response.code === 200) {
        dialogVisible.value = false;
        await listData();
      } else {
        await ElMessageBox.alert(response.msg, '更新失败', {
          confirmButtonText: 'OK',
          type: 'error'
        });
      }
    };

    const isEditing = ref(false);
    const getDetail = async (id) => {
      try {
        const routeDetail = await detail(id);

        if (routeDetail) {
          Object.assign(Route.value, routeDetail.data);
          Route.value.ParentID = routeDetail.data.ParentID;
          isEditing.value = true;
          dialogVisible.value = true;
        } else {
          console.error("Failed to fetch route details.");
        }
      } catch (error) {
        console.error("Error fetching route details:", error);
      }
    };

    const addNew = () => {
      dialogVisible.value = true;
    };

    const deleteRoute = async (id) => {
      try {
        await deleted(id);
        await listData();
      } catch (error) {
        console.error("Error deleting Route:", error);
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
      listData();           // 重新加载节点
    };

    onMounted(listData);

    return {
      Routes,
      searchText,
      loadRoutes,
      detail,
      deleteRoute,
      listData,
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
      RouteOptions,
      onRouteSelected,
      anyProps,
      saveRoute,
      getDetail
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
