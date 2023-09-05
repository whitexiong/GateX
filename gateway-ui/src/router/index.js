import { createRouter, createWebHistory } from 'vue-router';
import MainLayout from '@/components/MainLayout.vue';
import Dashboard from '@/components/Dashboard.vue';
import UserLogin from '@/components/UserLogin.vue';
import MenuManagement from '@/views/menu/index.vue';
import policy from '@/views/policy/index.vue';
import role from '@/views/role/index.vue';
import route from '@/views/route/index.vue'
import axios from "axios";

const routes = [
    {
        path: '/login',
        name: 'UserLogin',
        component: UserLogin
    },
    {
        path: '/',
        component: MainLayout,
        children: [
            {
                path: '',
                name: 'Dashboard',
                component: Dashboard,
                meta: {
                    requiresAuth: true
                }
            },
            {
                path: 'menu',
                component: MenuManagement,  // Placeholder, might not display anything or act as a layout
                children: [
                    {
                        path: 'index',
                        name: 'MenuIndex',
                        component: MenuManagement,  // Main list view
                        meta: {
                            requiresAuth: true
                        }
                    }
                ]
            },
            {
                path: 'policy',
                component: policy,
                children: [
                    {
                        path: 'index',
                        name: 'PolicyIndex',
                        component: policy,
                        meta: {
                            requiresAuth: true
                        }
                    },
                ]
            },
            {
                path: 'role',
                component: role,
                children: [
                    {
                        path: 'index',
                        name: 'RoleIndex',
                        component: role,
                        meta: {
                            requiresAuth: true
                        }
                    },
                ]
            },
            {
                path: 'route',
                component: route,
                children: [
                    {
                        path: 'index',
                        name: 'RouteIndex',
                        component: route,
                        meta: {
                            requiresAuth: true
                        }
                    }
                ]
            }
        ]
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

// 全局前置守卫
router.beforeEach((to, from, next) => {
    const requiresAuth = to.matched.some(record => record.meta.requiresAuth);
    const token = localStorage.getItem('token');

    if (requiresAuth && !token) {
        next({ name: 'UserLogin' });
    } else if (to.name === 'UserLogin' && token) {
        // 如果已经登录，从登录页跳转到 dashboard
        next({ name: 'Dashboard' });
    } else {
        next();
    }
});

// 全局请求拦截器
axios.interceptors.request.use(
    config => {
        // 从 localStorage 获取 token
        const token = localStorage.getItem('token');

        // 如果 token 存在，将其添加到所有请求的 Authorization header 中
        if (token) {
            config.headers.Authorization = 'Bearer ' + token;
        }

        return config;
    },
    error => {
        return Promise.reject(error);
    }
);


// 全局响应拦截器
axios.interceptors.response.use(
    response => {
        return response.data;
    },
    error => {
        return Promise.reject(error);
    }
);

export default router;
