<template>


  <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
  >
    <el-menu-item index="0" @click="toggleSidebar">
<!--      <img src="@/assets/logo.png" alt="Logo" class="nav-logo"/>-->
      <el-icon class="toggle-sidebar-icon nav-logo" :name="isCollapse ? 'el-icon-arrow-right' : 'el-icon-arrow-left'" ><Fold /></el-icon>

    </el-menu-item>

    <div class="flex-grow" />
    <el-sub-menu index="2">
      <template #title>超级管理员</template>
      <el-menu-item index="2-1">修改密码</el-menu-item>
      <el-menu-item index="2-1">退出登录</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>
<script setup>
import { ref } from 'vue'
import {Fold} from "@element-plus/icons-vue";

const activeIndex = ref('1')

const handleSelect = (key, keyPath) => {
  console.log(key, keyPath)
}


</script>

<script>
export default {
  methods: {
    toggleSidebar() {
      this.$emit('toggle-sidebar');
    }
  }
}
</script>

<style>
.flex-grow {
  flex-grow: 1;
}

.nav-logo {
  height: 60px;  /* 调整为合适的大小 */
  vertical-align: middle;  /* 使图片在导航栏中垂直居中 */
}

</style>
