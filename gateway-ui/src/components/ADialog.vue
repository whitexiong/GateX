<template>
  <el-dialog :title="title" :visible="visible" width="50%" @close="closeDialog">
    <slot></slot>
    <template #footer>
      <el-button @click="resetDialog">重置</el-button>
      <el-button @click="confirmDialog">确认</el-button>
    </template>
  </el-dialog>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      default: ''
    },
    visible: {
      type: Boolean,
      default: false
    }
  },
  emits: ['update:visible', 'close', 'confirm', 'reset'], // 显式声明发出的事件
  methods: {
    closeDialog() {
      this.$emit('update:visible', false);
      this.$emit('close');
    },
    confirmDialog() {
      this.$emit('confirm');
    },
    resetDialog() {
      this.$emit('reset'); // 新增的重置事件
    }
  }
};
</script>

<style scoped>
/* ... */
</style>
