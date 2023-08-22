<template>
  <div>
    <!-- 搜索框 -->
    <el-input
        v-model="internalSearchQuery"
        placeholder="搜索"
        @input="updateSearch"
    ></el-input>

    <!-- 表格 -->
    <el-table :data="displayedData" style="width: 1980px; height: 1000px" tree-props="{children: 'children'}">
      <slot></slot>
    </el-table>


    <!-- 分页 -->
    <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="currentPage"
        :page-sizes="[5, 10, 20, 50]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="filteredData.length"
    ></el-pagination>
  </div>
</template>

<script>
import { ref, watch, toRefs, computed } from "vue";

export default {
  name: "ATable",
  props: {
    data: {
      type: Array,
      required: true
    },
    searchQuery: {
      type: String,
      default: ""
    }
  },
  setup(props) {
    const { data } = toRefs(props);

    const internalSearchQuery = ref(props.searchQuery);
    const currentPage = ref(1);
    const pageSize = ref(10);

    const filteredData = computed(() => {
      return data.value.filter(item =>
          JSON.stringify(item).toLowerCase().includes(internalSearchQuery.value.toLowerCase())
      );
    });

    const displayedData = computed(() => {
      const start = (currentPage.value - 1) * pageSize.value;
      const end = start + pageSize.value;
      return filteredData.value.slice(start, end);
    });

    const updateSearch = () => {
      currentPage.value = 1; // 重置当前页
    };

    const handleSizeChange = (newSize) => {
      pageSize.value = newSize;
    };

    const handleCurrentChange = (newPage) => {
      currentPage.value = newPage;
    };

    return {
      internalSearchQuery,
      currentPage,
      pageSize,
      filteredData,
      displayedData,
      updateSearch,
      handleSizeChange,
      handleCurrentChange
    };
  }
};
</script>

<style scoped>
/* 你可以在这里添加样式 */
</style>
