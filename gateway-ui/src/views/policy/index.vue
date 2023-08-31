<template>

  <!-- 搜索和操作部分 -->
  <div style="margin-bottom: 20px; display: flex; align-items: center; justify-content: space-between;">

    <!-- 搜索部分 -->
    <div style="display: flex; align-items: center;">
      <el-input
          v-model="searchText"
          placeholder="请输入策略类型"
          style="width: 200px;"
          @keyup.enter="fetchPolicies">
        <template #append>
          <el-button @click="fetchPolicies" style="margin-right: 5px;">
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

      <el-button @click="addNewPolicy" style="margin-left: 10px;">
        <el-icon>
          <Plus/>
        </el-icon>
        新增策略
      </el-button>

      <!-- 新增策略对话框 -->
      <ADialog
          v-model="dialogVisible"
          title="新增策略"
          @confirm="handleAddPolicy">
        <el-form ref="policyForm" :model="newPolicy" label-width="100px">
          <!-- 策略类型 -->
          <el-form-item label="策略类型">
            <el-input v-model="newPolicy.PType"></el-input>
          </el-form-item>
          <!-- ...其他字段输入... -->
        </el-form>
      </ADialog>
    </div>
  </div>

  <!-- 策略列表 -->
  <el-table :data="policy" row-key="ID" style="width: 1980px; height: 1000px">
    <el-table-column label="ID" prop="ID"></el-table-column>
    <el-table-column label="策略类型" prop="PType"></el-table-column>

    <!-- 使用v-for循环创建列 -->
    <el-table-column
        v-for="index in 6"
        :key="index"
        :label="['主体', '对象', '行为', '附加信息1', '附加信息2', '附加信息3'][index-1]"
        :prop="['V0', 'V1', 'V2', 'V3', 'V4', 'V5'][index-1]">
    </el-table-column>

    <el-table-column label="操作" width="180">
      <template #default="{ row }">
        <el-button size="mini" @click="editPolicy(row.ID)">编辑</el-button>
        <el-button size="mini" @click="deletePolicy(row.ID)" type="danger">删除</el-button>
      </template>
    </el-table-column>
  </el-table>

  <!-- 新增策略弹出框 -->
  <ADialog v-model="dialogVisible" title="新增策略" @confirm="handleAddPolicy">
    <el-form ref="policyForm" :model="newPolicy" label-width="100px">
      <!-- 常规策略字段输入 -->
      <el-form-item label="策略类型">
        <el-input v-model="newPolicy.PType"></el-input>
      </el-form-item>
      <el-form-item label="主体">
        <el-input v-model="newPolicy.V0"></el-input>
      </el-form-item>
      <el-form-item label="对象">
        <el-input v-model="newPolicy.V1"></el-input>
      </el-form-item>
      <el-form-item label="行为">
        <el-input v-model="newPolicy.V2"></el-input>
      </el-form-item>
      <!-- ...其他字段输入... -->
    </el-form>

    <!-- 添加关闭按钮 -->
    <template #footer>
      <el-button @click="dialogVisible = false">关闭</el-button>
    </template>
  </ADialog>

</template>


<script>
import { ref, onMounted } from 'vue';
import ADialog from '@/components/ADialog.vue';
import { getList, add, deletePolicy } from "@/services/policyService";

export default {
  components: { ADialog },
  setup() {
    const policy = ref([]);
    const dialogVisible = ref(false);
    const newPolicy = ref({
      PType: '', V0: '', V1: '', V2: '', V3: '', V4: '', V5: ''
    });

    const fetchPolicies = async () => {
      try {
        const response = await getList();
        policy.value = response.data.policy || [];
      } catch (error) {
        console.error("Error fetching policies:", error);
      }
    };

    const handleAddPolicy = async () => {
      try {
        const response = await add(newPolicy.value);
        if (response.data.success) {
          dialogVisible.value = false;
          await fetchPolicies();
        } else {
          console.error("Error adding policy:", response.data.message);
        }
      } catch (error) {
        console.error("Error adding policy:", error);
      }
    };

    const editPolicy = (id) => {
      const selectedPolicy = policy.value.find(p => p.ID === id);
      newPolicy.value = { ...selectedPolicy };
      dialogVisible.value = true;
    };

    const deletePolicy = async (id) => {
      try {
        const response = await deletePolicy(id);
        if (response.data.success) {
          await fetchPolicies();
        } else {
          console.error("Error deleting policy:", response.data.message);
        }
      } catch (error) {
        console.error("Error deleting policy:", error);
      }
    };

    onMounted(fetchPolicies);

    return {
      policy, dialogVisible, newPolicy, fetchPolicies, handleAddPolicy, editPolicy, deletePolicy
    };
  }
};

</script>
