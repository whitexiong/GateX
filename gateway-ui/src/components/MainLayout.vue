<template>
  <div class="common-layout">
    <el-container>
      <SideBarMenu class="sidebar" :isCollapse="isCollapse" />
      <el-container>
        <el-header>
          <NavBar @toggle-sidebar="toggleSidebar" :isCollapse="isCollapse" class="navbar" />
        </el-header>
        <BreadcrumbNav />
        <el-main>
          <el-scrollbar>
            <router-view />
          </el-scrollbar>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import SideBarMenu from "@/components/SideBarMenu.vue";
import NavBar from "@/components/NavBar.vue";
import BreadcrumbNav from '@/components/BreadcrumbNav.vue';

export default {
  name: 'App',
  components: {
    NavBar,
    SideBarMenu,
    BreadcrumbNav
  },
  data(){
    return {
      isCollapse: true
    }
  },
  methods: {
    toggleSidebar() {
      this.isCollapse = !this.isCollapse;
    }
  },
}

</script>

<style scoped>
</style>
