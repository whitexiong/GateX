<template>
  <div>
    <template v-for="menu in menus">
      <el-sub-menu v-if="menu.children && menu.children.length" :key="menu.id" :index="`${menu.id}`">
        <template #title>
<!--            <component :is="getIconComponent(menu.icon)" />-->
          <span>{{ menu.name }}</span>
        </template>
        <MenuRenderer :menus="menu.children" />
      </el-sub-menu>
      <el-menu-item v-else :key="menu.id" :index="`${menu.id}`">
<!--          <component :is="getIconComponent(menu.icon)" />-->
        <router-link :to="menu.path">{{ menu.name }}</router-link>
      </el-menu-item>
    </template>
  </div>
</template>


<script>

import * as icons from '@element-plus/icons-vue';

export default {
  name: 'MenuRenderer',
  props: {
    menus: {
      type: Array,
      default: () => []
    },
    isCollapse: {
      type: Boolean,
      default: false
    }
  },

  setup(){
    const getIconComponent = (icon) => {
      return icons[icon];
    };

    return {
      getIconComponent
    }
  }
}
</script>

<style scoped>
</style>
