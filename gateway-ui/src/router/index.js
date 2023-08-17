import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '@/components/Dashboard';

const routes = [
    {
        path: '/dashboard',
        name: 'Dashboard',
        component: Dashboard
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

export default router;
