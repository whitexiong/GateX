<template>
  <el-menu default-active="1" class="el-menu-vertical" :collapse="isCollapse">
    <!-- 系统概览 -->
    <el-menu-item index="1">
      <el-icon><HomeFilled /></el-icon>
      <template #title>系统概览</template>
    </el-menu-item>

    <!-- 路由管理 -->
    <el-sub-menu index="2">
      <template #title>
        <el-icon><Operation /></el-icon>
        <span>路由管理</span>
      </template>
      <el-menu-item index="2-1">添加路由</el-menu-item>
      <el-menu-item index="2-2">修改路由</el-menu-item>
      <el-menu-item index="2-3">删除路由</el-menu-item>
    </el-sub-menu>

    <!-- API 管理 -->
    <el-sub-menu index="3">
      <template #title>
        <el-icon><ScaleToOriginal /></el-icon>
        <span>API 管理</span>
      </template>
      <el-menu-item index="3-1">添加API</el-menu-item>
      <el-menu-item index="3-2">修改API</el-menu-item>
      <el-menu-item index="3-3">删除API</el-menu-item>
    </el-sub-menu>

    <!-- 访问控制 -->
    <el-sub-menu index="4">
      <template #title>
        <el-icon><Lock /></el-icon>
        <span>访问控制</span>
      </template>
      <el-menu-item index="4-1">IP 黑/白名单</el-menu-item>
      <el-menu-item index="4-2">API 访问权限</el-menu-item>
    </el-sub-menu>

    <!-- 日志查看 -->
    <el-menu-item index="5">
      <el-icon><Document /></el-icon>
      <template #title>日志查看</template>
    </el-menu-item>


    <!-- 监控&告警 -->
    <el-menu-item index="6">
      <el-icon><BellFilled /></el-icon>
      <template #title>监控&告警</template>
    </el-menu-item>

    <!-- 系统设置 -->
    <el-sub-menu index="7">
      <template #title>
        <el-icon><Setting /></el-icon>
        <span>系统设置</span>
      </template>

      <!-- 用户管理 -->
      <el-menu-item index="7-1">
        <router-link to="/users">用户管理</router-link>
      </el-menu-item>

      <!-- 角色管理 -->
      <el-menu-item index="7-2">
        <router-link to="/role">角色管理</router-link>
      </el-menu-item>

      <!-- 策略管理 -->
      <el-menu-item index="7-3">
        <router-link to="/policy">策略管理</router-link>
      </el-menu-item>

      <!-- 菜单管理 -->
      <el-menu-item index="7-4">
        <router-link to="/menu">菜单管理</router-link>
      </el-menu-item>

      <!-- 路由管理 -->
      <el-menu-item index="7-5">
        <router-link to="/routes">路由管理</router-link>
      </el-menu-item>

      <!-- 日志查看 -->
      <el-menu-item index="7-6">
        <router-link to="/logs">日志查看</router-link>
      </el-menu-item>

      <!-- 数据恢复 -->
      <el-menu-item index="7-7">
        <router-link to="/recovery">数据恢复</router-link>
      </el-menu-item>

    </el-sub-menu>


    <!-- 帮助&文档 -->
    <el-menu-item index="8">
      <el-icon><QuestionFilled /></el-icon>
      <template #title>帮助&文档</template>
    </el-menu-item>
  </el-menu>
</template>

<script setup>
import {
  Document,
  Setting, QuestionFilled, HomeFilled, Operation, ScaleToOriginal, Lock, BellFilled,
} from '@element-plus/icons-vue'
</script>

<script>
export default {
  props: {
    isCollapse: {
      type: Boolean,
      required: true
    }
  }
}
</script>

<style>
.el-menu-vertical:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}

</style>

<style scoped>
.el-menu-vertical {
  width: 200px; /* 默认宽度 */
}

.el-menu-vertical.el-menu--collapse {
  width: 80px; /* 折叠后的宽度 */
}
</style>
