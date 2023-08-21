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
      <el-menu-item index="2-2" @click="logout">退出登录</el-menu-item>
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
import { UserLogout } from '@/services/api';
import router from "@/router";
import { ElMessageBox } from 'element-plus';


export default {
  methods: {
    toggleSidebar() {
      this.$emit('toggle-sidebar');
    }
  }
}

const logout = () => {
  ElMessageBox.confirm('确定要退出吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await UserLogout();
      localStorage.removeItem('token');
      router.push('/login');
    } catch (error) {
      console.error("退出登录失败：", error);
    }
  }).catch(() => {
    // 这里处理用户点击“取消”按钮的情况，如果你不需要特殊处理，可以留空。
  });
};

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
