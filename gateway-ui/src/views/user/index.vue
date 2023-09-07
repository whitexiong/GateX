<template>
  <div>
    <div style="margin-bottom: 20px; display: flex; align-items: center; justify-content: space-between;">

      <div style="display: flex; align-items: center;">
        <el-input
            v-model="searchText"
            placeholder="请输入用户名"
            style="width: 200px;"
            @keyup.enter="listData">
          <template #append>
            <el-button @click="listData" style="margin-right: 5px;">
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
          新增用户
        </el-button>
        <ADialog
            v-model="dialogVisible"
            :title="dialogTitle"
            @confirm="saveData"
            @reset="resetData"
        >
          <el-form ref="UserForm" :model="User" label-width="80px" style="width: 100%;">
            <el-form-item label="用户名">
              <el-input v-model="User.Username" placeholder="请输入用户名"></el-input>
            </el-form-item>

            <el-form-item label="密码">
              <el-input
                  v-model="User.Password"
                  :type="showPassword ? 'text' : 'password'"
                  placeholder="请输入密码"
                  suffix-icon="el-icon-view"
                  @click:append="togglePasswordVisibility">
              </el-input>
            </el-form-item>

            <el-form-item label="邮箱">
              <el-input v-model="User.Email" placeholder="请输入邮箱"></el-input>
            </el-form-item>
            <el-form-item label="状态">
              <el-switch v-model="User.Status" :active-value="1" :inactive-value="0"></el-switch>
            </el-form-item>
            <el-form-item label="角色">
              <el-select v-model="User.Roles" multiple placeholder="请选择角色">
                <el-option
                    v-for="role in allRoles"
                    :key="role.ID"
                    :label="role.Name"
                    :value="role.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-form>
        </ADialog>
      </div>
    </div>

    <!-- 用户表格区域 -->
    <el-table :data="Users"  row-key="ID" lazy :load="loadTree" style="width: 1980px; height: 1000px" border>>
      <el-table-column label="用户名" prop="Username"></el-table-column>
      <el-table-column label="邮箱" prop="Email"></el-table-column>
      <el-table-column label="状态" width="100">
        <template #default="scope">
          <span>{{ scope.row.Status === 1 ? '启用' : '禁用' }}</span>
        </template>
      </el-table-column>
      <el-table-column label="角色" width="200">
        <template #default="scope">
          <span>{{ scope.row.Roles.map(role => role.Name).join(', ') }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="260">
        <template #default="{ row }">
          <div style="display: flex; align-items: center;">
            <el-button type="primary" size="small" @click="getDetail(row.ID)" style="color: black; margin-left: 5px;">
              编辑
            </el-button>
            <el-button type="danger" size="small" @click="deleteUser(row.ID)" style="color: black; margin-left: 10px;">
              删除
            </el-button>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {ref, onMounted} from 'vue';
import {getList as getRoleList} from '@/services/roleService';
import {getList as getUserList, add, update, detail, deletedById} from '@/services/userService';
import {Plus, Refresh, RefreshRight, Search} from "@element-plus/icons-vue";
import ADialog from '@/components/ADialog.vue';
import {useCRUD} from '@/composables/useCRUD';

export default {
  components: {Refresh, Search, Plus, RefreshRight, ADialog},
  setup() {
    const initialUser = {
      ID: null,
      Username: '',
      Email: '',
      Password: '',
      Status: null,
      Roles: []
    };

    const apiMethods = {
      getList: getUserList,
      add: add,
      update: update,
      detail: detail,
      deletedById: deletedById
    };

    const {
      data: Users,
      selected: User,
      dialogVisible,
      searchText,
      listData,
      saveData,
      refresh,
      addNew,
      getDetail,
      deleted: deleteUser,
      resetData,
      dialogTitle
    } = useCRUD(apiMethods, initialUser);

    const allRoles = ref([]);

    onMounted(async () => {
      await listData();
      const rolesResponse = await getRoleList();
      if (rolesResponse.data) {
        allRoles.value = rolesResponse.data;
      }
    });

    return {
      Users,
      User,
      allRoles,
      dialogVisible,
      searchText,
      listData,
      saveData,
      refresh,
      addNew,
      getDetail,
      deleteUser,
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
