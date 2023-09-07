<template>
  <el-menu
      default-active="1"
      class="el-menu-vertical"
      :collapse="isCollapse">

    <MenuRenderer :menus="menuData" :isCollapse="isCollapse"/>

  </el-menu>
</template>

<script>
import {ref, onMounted} from 'vue';
import MenuRenderer from '@/components/MenuRenderer.vue';

export default {
  props: {
    isCollapse: {
      type: Boolean,
      required: true
    }
  },
  components: {
    MenuRenderer
  },

  setup() {

    const menuData = ref([]);

    onMounted(() => {
      const storedMenus = localStorage.getItem('userMenus');
      if (storedMenus) {
        menuData.value = JSON.parse(storedMenus);
      }
    });

    return {menuData};
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
