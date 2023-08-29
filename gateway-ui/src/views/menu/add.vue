<template>
  <div>
    <!-- Triggering the dialog -->
    <el-button @click="showDialog = true">新增菜单</el-button>

    <!-- Using the abstracted ADialog component -->
    <ADialog
        v-model="showDialog"
        title="新增菜单"
        @submit="handleAddMenu"
    >
      <!-- Place form elements or any content you want to display inside the dialog here -->
      <el-input v-model="menu.name" placeholder="请输入菜单名称"></el-input>
      <el-input v-model="menu.icon" placeholder="请输入菜单图标"></el-input>
      <el-input v-model="menu.path" placeholder="请输入菜单链接"></el-input>
      <!-- Add more form elements as needed -->
    </ADialog>
  </div>
</template>

<script>
import ADialog from '@/components/ADialog.vue';

export default {
  components: {
    ADialog
  },
  data() {
    return {
      showDialog: false,  // This controls the visibility of the dialog
      menu: {             // This represents the new menu item data
        name: '',
        icon: '',
        path: ''
      }
    }
  },
  methods: {
    handleAddMenu() {
      // Call API to add new menu using this.menu data
      // Once done, reset the menu data and close the dialog
      this.menu = {
        name: '',
        icon: '',
        path: ''
      };
      this.showDialog = false;
    }
  }
}
</script>

<style scoped>
/* Any specific styles for add.vue */
</style>
