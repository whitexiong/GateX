<template>
  <div class="common-layout">
    <el-container>
        <SideBarMenu class="sidebar" :isCollapse="isCollapse" />
      <el-container>
        <el-header>
          <NavBar @toggle-sidebar="toggleSidebar"  class="navbar" />
        </el-header>
        <BreadcrumbNav />
        <el-main>
          <el-scrollbar>
            <Dashboard />
          </el-scrollbar>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Dashboard from '@/components/Dashboard.vue';
import SideBarMenu from "@/components/SideBarMenu.vue";
import NavBar from "@/components/NavBar.vue";
import BreadcrumbNav from '@/components/BreadcrumbNav.vue';

export default {
  name: 'App',
  components: {
    NavBar,
    Dashboard,
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
      console.log("传递到父组件");
      this.isCollapse = !this.isCollapse;
    }
  },
}

</script>

<style scoped>
</style>
