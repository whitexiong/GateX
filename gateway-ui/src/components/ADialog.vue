<template>
  <el-dialog :title="title" :visible="visible" width="50%" @close="closeDialog">
    <slot></slot>
    <template #footer>
      <el-button @click="closeDialog">取消</el-button>
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
  emits: ['update:visible', 'close', 'confirm'], // 显式声明发出的事件
  methods: {
    closeDialog() {
      this.$emit('update:visible', false);
      this.$emit('close');
    },
    confirmDialog() {
      this.$emit('confirm');
    }
  }
};
</script>

<style scoped>
/* ... */
</style>
