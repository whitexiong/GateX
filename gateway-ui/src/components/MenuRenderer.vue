<template>
  <div>
    <template v-for="menu in menus">
      <el-sub-menu v-if="menu.children && menu.children.length" :key="menu.id" :index="`${menu.id}`">
        <template #title>
          <component :is="getIconComponent(menu.icon)" class="icon-size" />
          <span v-if="!isCollapse">{{ menu.name }}</span>
        </template>
        <MenuRenderer :menus="menu.children" :isCollapse="isCollapse" />
      </el-sub-menu>
      <el-menu-item v-else :key="menu.id" :index="`${menu.id}`">
        <component :is="getIconComponent(menu.icon)" class="icon-size" />
        <router-link :to="menu.path">
          <span>{{ menu.name }}</span>
        </router-link>
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
.icon-size {
  width: 1em;
  height: 1em;
}
</style>
