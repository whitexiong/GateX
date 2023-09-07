<template>
  <div>
    <!-- 搜索和刷新区域 -->
    <div style="margin-bottom: 20px; display: flex; align-items: center; justify-content: space-between;">

      <!-- 搜索部分 -->
      <div style="display: flex; align-items: center;">
        <el-input
            v-model="searchText"
            placeholder="请输入角色名称"
            style="width: 200px;"
            @keyup.enter="fetchRoles">
          <template #append>
            <el-button @click="fetchRoles" style="margin-right: 5px;">
              <el-icon>
                <Search/>
              </el-icon>
            </el-button>
            <el-button @click="resetSearch">
              <el-icon>
                <Refresh/>
              </el-icon>
            </el-button>
          </template>
        </el-input>
      </div>

      <!-- 操作部分 -->
      <div style="display: flex; align-items: center;">
        <el-button @click="refresh">
          <el-icon>
            <RefreshRight/>
          </el-icon>
        </el-button>

        <el-button @click="addNew" style="margin-left: 10px;">
          <el-icon>
            <Plus/>
          </el-icon>
          新增
        </el-button>

        <!-- 新增角色对话框 -->
        <ADialog
            v-model="dialogVisible"
            :title="dialogTitle"
            @confirm="saveData"
            @reset="resetData"
            @close="closeRole"
        >
          <el-form ref="roleForm" :model="Role" label-width="80px" style="width: 100%;">
            <!-- 名称 -->
            <el-form-item label="名称">
              <el-input v-model="Role.Name" placeholder="请输入角色名称"></el-input>
            </el-form-item>
            <!-- 备注 -->
            <el-form-item label="备注">
              <el-input v-model="Role.Remark" type="textarea"/>
            </el-form-item>
            <!-- 状态 -->
            <el-form-item label="状态">
              <el-switch
                  v-model="Role.Status"
                  :active-value="1"
                  :inactive-value="0"
              ></el-switch>
            </el-form-item>
            <el-form-item label="选择权限">
              <el-tree
                  :data="allPermissions"
                  ref="treeRef"
                  show-checkbox
                  default-expand-all
                  node-key="id"
                  highlight-current
                  :default-checked-keys="Role.Permissions"
                  :check-strictly="true"
                  @check="getSelectedPermissions"
              />
            </el-form-item>

          </el-form>
        </ADialog>
      </div>
    </div>

    <!-- 表格区域 -->
    <el-table :data="Roles" row-key="ID" style="width: 1980px; height: 1000px" border>
      <el-table-column label="角色名称" prop="Name"></el-table-column>
      <el-table-column label="备注" prop="Remark"></el-table-column>
      <el-table-column label="状态">
        <template #default="{ row }">
          <el-tag v-if="row.Status === 1" type="success">开启</el-tag>
          <el-tag v-else type="info">禁用</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="260">
        <template #default="{ row }">
          <div style="display: flex; align-items: center;">
            <el-button size="small" @click="toggleStatus(row)">
              {{ row.Status === 1 ? '禁用' : '开启' }}
            </el-button>
            <el-button type="primary" size="small" @click="getDetail(row.ID)" style="color: black; margin-left: 5px;">编辑</el-button>
            <el-button type="danger" size="small" @click="deletedRole(row.ID)" style="color: black; margin-left: 10px;">删除</el-button>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {ref, onMounted} from 'vue';
import {Plus, Refresh, RefreshRight, Search} from "@element-plus/icons-vue";
import ADialog from '@/components/ADialog.vue';
import {getList, detail, add, update, getPermissions, deletedRole} from '@/services/roleService';
import {useCRUD} from "@/composables/useCRUD";

export default {
  components: {Refresh, Search, Plus, RefreshRight, ADialog},
  setup() {
    const allPermissions = ref([]); // 所有的权限
    const selectedPermissions = ref([]); // 选中的权限
    const treeRef = ref(null);
    const initFormData = {
      ID: null,
      Name: '',
      Remark: '',
      Status: 1,
      Permissions:[],
    }

    const apiMethods = {
      getList,
      add,
      update,
      detail,
      deleted: deletedRole
    };

    const {
      data: Roles,
      selected: Role,
      dialogVisible,
      searchText,
      currentPage,
      pageSize,
      listData,
      saveData,
      refresh,
      addNew,
      getDetail,
      resetData,
      dialogTitle,
      handlePageChange,
      toggleStatus,
    } = useCRUD(apiMethods, initFormData);

    // 假设你有一个服务或API来获取所有的权限列表
    const fetchAllPermissions = async () => {
      const response = await getPermissions();
      if (response.data) {
        allPermissions.value = response.data;
      }
    };


    const getSelectedPermissions = (checkedNodes, { checkedKeys, halfCheckedKeys }) => {
      const allSelectedKeys = [...checkedKeys]; // 只考虑完全选中的节点
      Role.value.Permissions = allSelectedKeys;
      console.log(allSelectedKeys);
    }


    onMounted(async () => {
      await listData();
      await fetchAllPermissions();
    });

    const closeRole = () => {
      treeRef.value.setCheckedKeys([], false)
    }

    return {
      Roles,
      Role,
      searchText,
      refresh,
      addNew,
      dialogVisible,
      toggleStatus,
      allPermissions,
      selectedPermissions,
      treeRef,
      getSelectedPermissions,
      currentPage,
      pageSize,
      saveData,
      getDetail,
      resetData,
      dialogTitle,
      handlePageChange,
      closeRole,
      deletedRole
    };
  },
};
</script>

<style scoped>
/* 你可以在这里添加一些样式 */

.el-tree {
  width: 100%; /* 使用100%宽度，使其填满父容器 */
  max-height: 300px; /* 设置最大高度，你可以根据需要调整 */
  overflow-y: auto; /* 如果内容超出最大高度，则出现滚动条 */
}

.table-button {
  color: black;
  margin-left: 10px;
}

.table-button-danger {
  color: black; /* 你可以根据需要为特定的按钮设置特定的颜色 */
}

</style>
