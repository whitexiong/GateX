<template>
  <div>
    <!-- 搜索和刷新区域 -->
    <div style="margin-bottom: 20px; display: flex; align-items: center; justify-content: space-between;">

      <!-- 搜索部分 -->
      <div style="display: flex; align-items: center;">
        <el-input
            v-model="searchText"
            placeholder="请输入策略名称"
            style="width: 200px;"
            @keyup.enter="fetchPolicies">
          <template #append>
            <el-button @click="fetchPolicies" style="margin-right: 5px;">
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
            title="新增策略"
            @confirm="handleAddPolicy"
        >
          <el-form ref="policyForm" :model="policy" label-width="80px">
            <!-- 名称 -->
            <el-form-item label="名称">
              <el-input v-model="policy.name" placeholder="请输入策略名称"></el-input>
            </el-form-item>

            <!-- 描述 -->
            <el-form-item label="描述">
              <el-input v-model="policy.description" placeholder="请输入策略描述"></el-input>
            </el-form-item>

          </el-form>
        </ADialog>
      </div>
    </div>

    <!-- 表格区域 -->
    <el-table :data="policies" row-key="id"  style="width: 1980px; height: 1000px">
      <el-table-column label="策略名称" prop="name"></el-table-column>
      <el-table-column label="描述" prop="description"></el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="scope">
          <el-button type="primary" size="small" @click="editPolicy(scope.row.id)" style="color: black;">编辑</el-button>
          <el-button type="danger" size="small" @click="deletePolicy(scope.row.id)" style="margin-left: 10px; color: black;">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { Plus, Refresh, RefreshRight, Search } from "@element-plus/icons-vue";
import ADialog from '@/components/ADialog.vue';
import {getList} from "@/services/policyService";

export default {
  components: { Refresh, Search, Plus, RefreshRight, ADialog },
  setup() {
    const policies = ref([]);
    const searchText = ref('');
    const dialogVisible = ref(false);

    const policy = ref({
      name: '',       // 策略名称
      description: '' // 策略描述
    });

    const handleAddPolicy = () => {
      console.log('Policy data:', policy.value);
    };

    const fetchPolicies = async () => {
      const response = await getList({ parentId: 0, name: searchText.value });
      // Fetch the policy using your service (example)
      // const response = await getPolicies({ name: searchText.value });
      // policy.value = response.data.policy;
    };

    const refresh = () => {
      searchText.value = '';  // 清空搜索内容
      fetchPolicies();        // 重新加载策略
    };

    const addNew = () => {
      dialogVisible.value = true;
    };

    const editPolicy = (id) => {
      // Edit logic
    };

    const deletePolicy = async (id) => {
      // Delete logic
    };

    const resetSearch = () => {
      searchText.value = '';
      fetchPolicies();
    };

    onMounted(fetchPolicies);

    return {
      policies,
      searchText,
      fetchPolicies,
      refresh,
      addNew,
      policy,
      dialogVisible,
      handleAddPolicy,
      editPolicy,
      deletePolicy,
      resetSearch,
    };
  },
};
</script>

<style scoped>
/* 你可以在这里添加一些样式 */
</style>
