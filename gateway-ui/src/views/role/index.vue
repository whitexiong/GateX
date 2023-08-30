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
            title="新增角色"
            @confirm="handleAddRole">
          <el-form ref="roleForm" :model="role" label-width="100px">
            <!-- 名称 -->
            <el-form-item label="名称">
              <el-input v-model="role.Name" placeholder="请输入角色名称"></el-input>
            </el-form-item>
            <!-- 备注 -->
            <el-form-item label="备注">
              <el-input v-model="role.Remark" type="textarea" />
            </el-form-item>
            <!-- 状态 -->
            <el-form-item label="状态">
              <el-switch
                  v-model="role.Status"
                  active-value="1"
                  inactive-value="0"
              ></el-switch>
            </el-form-item>
          </el-form>
        </ADialog>
      </div>
    </div>

    <!-- 表格区域 -->
    <el-table :data="roles" row-key="ID" style="width: 1980px; height: 1000px">
      <el-table-column label="角色名称" prop="Name"></el-table-column>
      <el-table-column label="备注" prop="Remark"></el-table-column>
      <el-table-column label="状态">
        <template #default="{ row }">
          <el-tag v-if="row.Status === 1" type="success">开启</el-tag>
          <el-tag v-else type="info">禁用</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="{ row }">
          <el-button size="mini" @click="toggleStatus(row)">
            {{ row.Status === 1 ? '禁用' : '开启' }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {ref, onMounted} from 'vue';
import {Plus, Refresh, RefreshRight, Search} from "@element-plus/icons-vue";
import ADialog from '@/components/ADialog.vue';
import {getList,add} from "@/services/RoleService";  // Assume you have a rolesService

export default {
  components: {Refresh, Search, Plus, RefreshRight, ADialog},
  setup() {
    const roles = ref([]);
    const searchText = ref('');
    const dialogVisible = ref(false);


    const role = ref({
      Name: '',       // 角色名称
      Remark: '',     // 备注
      Status: null       // 默认状态为1（开启）
    });

    const handleAddRole = async () => {
      role.value = {
        ...role.value,
        Status: parseInt(role.value.Status, 10)
      };

      try {
        // 调用后端API来新增角色
        const response = await add(role.value);
        if (response && response.success) {
          // 成功后关闭弹窗并刷新角色列表
          dialogVisible.value = false;
          await fetchRoles();
        } else {
          console.error("Error adding role:", response);
        }
      } catch (error) {
        console.error("Error:", error);
      }
    };

    // 切换角色状态
    const toggleStatus = async (roleItem) => {
      // 这里你可以调用后端API来更改状态
      roleItem.Status = roleItem.Status === 1 ? 0 : 1;

      // 调用API更新状态
      // await yourApiFunctionToUpdateStatus(roleItem.ID, roleItem.Status);
      await fetchRoles();
    };

    const fetchRoles = async () => {
      const response = await getList({name: searchText.value});
      roles.value = response.data.roles;
      console.log(roles.value)
    };

    const refresh = () => {
      searchText.value = '';  // 清空搜索内容
      fetchRoles();           // 重新加载角色
    };

    const addNew = () => {
      dialogVisible.value = true;
    };

    const editRole = (id) => {
      // Edit logic
    };

    const deleteRole = async (id) => {
      // Delete logic
    };

    const resetSearch = () => {
      searchText.value = '';
      fetchRoles();
    };

    onMounted(fetchRoles);

    return {
      roles,
      searchText,
      fetchRoles,
      refresh,
      addNew,
      role,
      dialogVisible,
      handleAddRole,
      editRole,
      deleteRole,
      resetSearch,
      toggleStatus
    };
  },
};
</script>

<style scoped>
/* 你可以在这里添加一些样式 */
</style>
