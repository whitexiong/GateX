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
            @keyup.enter="getList">
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
            :title="dialogTitle"
            @confirm="saveData"
            @reset="resetData"
        >
          <el-form ref="MenuForm" :model="Menu" label-width="80px" style="width: 100%;">
            <!-- 父菜单 -->
            <el-form-item label="上级菜单">
              <el-cascader
                  v-model="Menu.ParentID"
                  :options="options"
                  placeholder="请选择菜单"
                  :props="anyProps"
                  @change="onSelected"
                  clearable
              ></el-cascader>
            </el-form-item>

            <!-- 名称 -->
            <el-form-item label="名称">
              <el-input v-model="Menu.Name" placeholder="请输入菜单名称"></el-input>
            </el-form-item>

            <!-- 路径 -->
            <el-form-item label="路径">
              <el-input v-model="Menu.Path" placeholder="请输入菜单路径"></el-input>
            </el-form-item>

            <!-- 图标选择 -->
            <el-form-item label="图标">
              <div class="icon-input-container" @click="openIconSelector">
                <component :is="getIconComponent(displayIcon)" class="display-icon" v-if="displayIcon" />
                <span v-else class="icon-placeholder">选择图标</span>
              </div>
              <el-input v-model="displayIcon" readonly v-if="false"></el-input>

              <el-dialog v-model="iconDialogVisible" width="50%" hight="50%">
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

            <!-- 排序 -->
            <el-form-item label="排序">
              <el-input-number v-model="Menu.Order" :min="0"></el-input-number>
            </el-form-item>

            <el-form-item label="状态">
              <el-switch
                  v-model="Menu.Status"
                  :active-value="1"
                  :inactive-value="0"
              ></el-switch>
            </el-form-item>
          </el-form>
        </ADialog>
      </div>
    </div>

    <!-- 表格区域 -->
    <el-table :data="Menus" row-key="id" lazy :load="loadTree" style="width: 1980px; height: 1000px" border>
      <el-table-column label="菜单名称">
        <template #default="scope">
          <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="图标" align="center" header-align="center">
        <template #default="scope">
        <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">
            <component :is="getIconComponent(scope.row.icon)" class="display-icon table-icon" v-if="scope.row.icon" />
        </span>
        </template>
      </el-table-column>

      <el-table-column label="菜单链接">
        <template #default="scope">
          <span :style="{ paddingLeft: (scope.row._indent || 0) * 20 + 'px' }">{{ scope.row.path }}</span>
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
import {ref, onMounted, computed} from 'vue';
import { getList, detail, add, update, deletedById } from '@/services/menuService';
import {Plus, Refresh, RefreshRight, Search, Tools} from "@element-plus/icons-vue";
import ADialog from '@/components/ADialog.vue';
import * as icons from '@element-plus/icons';
import {useCRUD} from "@/composables/useCRUD";

export default {
  components: {Tools, Refresh, Search, Plus, RefreshRight,ADialog},
  setup() {
    const selectedIcon = ref("");
    const iconDialogVisible = ref(false);
    const options = ref([]);
    const allIcons = Object.keys(icons);  // 获取所有图标的名字
    const initFormData = {
        ID: null,
        Name: null,
        Icon: null,
        Path: null,
        ParentID: null,
        Order: 0,
        Status: null,
    }

    const apiMethods = {
      getList,
      add,
      update,
      detail,
      deletedById
    };

    const anyProps = {
      checkStrictly: true,
      value: 'value',
      label: 'label',
      children: 'children'
    }

    const {
      data: Menus,
      selected: Menu,
      dialogVisible,
      searchText,
      currentPage,
      loadTree,
      pageSize,
      listData,
      saveData,
      refresh,
      addNew,
      getDetail,
      deleted,
      resetData,
      dialogTitle,
      handlePageChange,
      toggleStatus,
    } = useCRUD(apiMethods, initFormData);

    const transCascader = (Menu) => {
      return {
        value: Menu.value,
        label: Menu.label,
        children: Menu.children && Menu.children.length
            ? Menu.children.map(child => transCascader(child))
            : null
      };
    };

    const onSelected = (value) => {
      if (value && Array.isArray(value) && value.length > 0) {
        Menu.ParentID = value[value.length - 1];
      }
    }

    const selectIcon = (icon) => {
      displayIcon.value = icon;
      Menu.value.Icon = icon;
      console.log(Menu.value.Icon)
    };

    onMounted(async () => {
      await listData();
      if (Menus.value) {
        options.value = Menus.value.map(Menu => transCascader(Menu));
      }
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

    const paginatedIcons = computed(() => {
      const start = (currentPage.value - 1) * pageSize.value;
      const end = start + pageSize.value;
      return allIcons.slice(start, end);
    });

    return {
      Menus,
      Menu,
      searchText,
      refresh,
      addNew,
      dialogVisible,
      iconDialogVisible,
      icons,
      allIcons,
      currentPage,
      handlePageChange,
      selectedIcon,
      pageSize,
      onSelected,
      getIconComponent,
      openIconSelector,
      paginatedIcons,
      displayIcon,
      anyProps,
      selectIcon,
      options,
      dialogTitle,
      saveData,
      resetData,
      deleted,
      getDetail,
      toggleStatus,
      loadTree,
    };
  }
};
</script>

<style scoped>
.icon-grid {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 2rem;  /* 调整间距 */
}

.selectable-icon {
  font-size: 32px;  /* 调整图标大小 */
  width: auto;  /* 使用图标的实际宽度 */
  height: auto;  /* 使用图标的实际高度 */
  margin-bottom: 10px;  /* 在图标和名字之间增加间距 */
}

.icon-container {
  display: flex;
  flex-direction: column;  /* 让内容垂直排列 */
  align-items: center;  /* 内容水平居中 */
  justify-content: center;  /* 内容垂直居中 */
  width: 100px;
  height: 100px;
  padding: 10px;
  text-align: center;
  border-radius: 5px;
  transition: background-color 0.3s;
  cursor: pointer;
}

.icon-container:hover {
  background-color: #f0f0f0;
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
}

.icon-input-container {
  display: inline-block;
  width: 100px;
  height: 40px;
  cursor: pointer;
  border: 1px solid #ccc;
  text-align: center;
  line-height: 40px;
  background-color: #f7f7f7;
  position: relative;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.icon-input-container:hover {
  background-color: #e5e5e5;
}

.display-icon {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  font-size: 32px;  /* 调整显示图标的大小 */
}

.icon-placeholder {
  color: #999;
  font-size: 14px;
}

.table-icon {
  font-size: 16px;   /* 如果是字体图标 */
  width: 16px;      /* 如果是SVG图标 */
  text-align: left;  /* 显式地设置文本对齐方式为左对齐 */
  padding: 0;
  margin: 0;
}

</style>
