<template>
  <div class="user-login-container">
    <div class="login-box">
      <h2 class="mb-8 text-center">GATE</h2>
      <el-form :model="loginForm" @submit.prevent="handleLogin" class="space-y-4">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="loginForm.username"></el-input>
        </el-form-item>
        <el-form-item label="密码&nbsp&nbsp&nbsp" prop="password">
          <el-input type="password" v-model="loginForm.password"></el-input>
        </el-form-item>
        <el-form-item class="flex justify-between">
          <el-button @click="handleLogin">登录</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
        <p v-if="errorMsg" class="text-red-600 mt-2">{{ errorMsg }}</p>
      </el-form>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import { UserLogin, GetUserMenus } from "@/services/userService";
import {useRouter} from "vue-router";

export default {
  name: "UserLogin",

  setup() {
    const loginForm = ref({
      username: "",
      password: ""
    });

    const errorMsg = ref("");
    const router = useRouter();

    const fetchAndStoreMenus = async () => {
      try {
        const menuResponse = await GetUserMenus();
        if (menuResponse && menuResponse.data) {
          localStorage.setItem('userMenus', JSON.stringify(menuResponse.data));
        }
      } catch (error) {
        console.error("Failed to fetch menus:", error);
      }
    }

    const handleLogin = async () => {
      try {
        const response = await UserLogin(loginForm.value.username, loginForm.value.password);
        if (response && response.data.token) {
          // 存储token
          localStorage.setItem('token', response.data.token);

          // 请求并存储用户的菜单数据
          await fetchAndStoreMenus();

          // 导航到Dashboard
          await router.push({name: 'Dashboard'});
        } else {
          errorMsg.value = "登录失败: " + response.data.msg;
        }
      } catch (error) {
        errorMsg.value = "登录失败，请联系管理员!";
      }
    };

    const resetForm = () => {
      loginForm.value.username = '';
      loginForm.value.password = '';
      errorMsg.value = '';
    };

    return {
      loginForm,
      handleLogin,
      resetForm,
      errorMsg
    };
  }
};
</script>

<style scoped>
.user-login-container {
  background-color: #f3f4f6;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-image: url("~@/assets/img_1.png");
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
}

.login-box {
  background-color: rgba(255, 255, 255, 0.95);
  padding: 30px 50px;
  border-radius: 8px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  width: 320px;
}

</style>
