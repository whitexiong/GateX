<template>
  <div>
    <div style="margin-bottom: 20px; display: flex; align-items: center; justify-content: space-between;">

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

      <div style="display: flex; align-items: center;">
        <el-button @click="refresh">
          <el-icon><RefreshRight /></el-icon>
        </el-button>

        <el-button @click="addNew" style="margin-left: 10px;">
          <el-icon><Plus /></el-icon>新增
        </el-button>
        <ADialog
            v-model="dialogVisible"
            :title="dialogTitle"
            @confirm="saveData"
            @reset="resetData"
        >
          <el-form ref="RouteForm" :model="Route" label-width="80px" style="width: 100%;">
            <el-form-item label="上级节点" style="width: 100%;">
              <el-cascader
                  v-model="Route.ParentID"
                  :options="RouteOptions"
                  :props="anyProps"
                  placeholder="节点名称"
                  @change="onRouteSelected"
                  clearable
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
    <el-table :data="Routes" row-key="id" lazy :load="loadTree" style="width: 1980px; height: 1000px" border>

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
            <el-button type="danger" size="small" @click="deleted(row.id)" style="color: black; margin-left: 10px;">删除</el-button>
          </div>
        </template>
      </el-table-column>
    </el-table>

  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { getList, add, deletedById, update, detail } from '@/services/routeService';
import { Plus, Refresh, RefreshRight, Search } from "@element-plus/icons-vue";
import ADialog from '@/components/ADialog.vue';
import * as icons from '@element-plus/icons';
import { useCRUD } from '@/composables/useCRUD';

export default {
  components: { Refresh, Search, Plus, RefreshRight, ADialog },
  setup() {
    const initialRoute = {
      ID: null,
      Name: '',
      Status: null,
      ParentID: null,
      Path: null,
    };

    const apiMethods = {
      getList,
      add,
      update,
      detail,
      deletedById
    };

    const {
      data: Routes,
      selected: Route,
      dialogVisible,
      searchText,
      currentPage,
      pageSize,
      listData,
      saveData,
      refresh,
      addNew,
      getDetail,
      deleted,
      resetData,
      dialogTitle,
      loadTree,
      handlePageChange,
      toggleStatus
    } = useCRUD(apiMethods, initialRoute);

    const allIcons = Object.keys(icons);
    const selectedIcon = ref("");
    const iconDialogVisible = ref(false);
    const RouteOptions = ref([]);
    const anyProps = {
      checkStrictly: true,
      value: 'value',
      label: 'label',
      children: 'children'
    }

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
      if (value && Array.isArray(value) && value.length > 0) {
        Route.ParentID = value[value.length - 1];
      }
    };

    onMounted(async () => {
      await listData();
      if (Routes.value) {
        RouteOptions.value = Routes.value.map(Route => transformRouteToCascader(Route));
      }
    });

    return {
      Routes,
      Route,
      dialogVisible,
      searchText,
      currentPage,
      pageSize,
      listData,
      saveData,
      refresh,
      allIcons,
      selectedIcon,
      iconDialogVisible,
      handlePageChange,
      RouteOptions,
      onRouteSelected,
      anyProps,
      loadTree,
      toggleStatus,
      addNew,
      getDetail,
      deleted,
      resetData,
      dialogTitle
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
