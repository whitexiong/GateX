<template>
  <div>
    <!-- 搜索框 -->
    <el-input
        v-model="internalSearchQuery"
        placeholder="搜索"
        @input="updateSearch"
    ></el-input>

    <!-- 表格 -->
    <el-table :data="filteredData" style="width: 1980px; height: 1000px">
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
import { ref, watch, toRefs } from "vue";

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
    const { data, searchQuery } = toRefs(props);

    const internalSearchQuery = ref(searchQuery.value);
    const currentPage = ref(1);
    const pageSize = ref(10);

    const filteredData = ref(data.value);

    watch(data, (newData) => {
      filteredData.value = newData;
    });

    watch(searchQuery, (newQuery) => {
      internalSearchQuery.value = newQuery;
      currentPage.value = 1;
    });

    const updateSearch = () => {
      filteredData.value = data.value.filter(item =>
          JSON.stringify(item).includes(internalSearchQuery.value)
      );
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
