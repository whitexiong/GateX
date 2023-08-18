import { createApp } from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css'
import App from './App.vue'
import '@/assets/css/main.css';
import router from './router'


createApp(App).use(router).use(ElementPlus).mount('#app')
