import { createApp, ref, provide } from 'vue';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import App from './App.vue';
import '@/assets/css/main.css';
import router from './router';

const app = createApp(App);

// 创建响应式引用
const isDarkMode = ref(false);

// 在应用级别提供 isDarkMode
app.provide('isDarkMode', isDarkMode);

app.use(router).use(ElementPlus).mount('#app');
