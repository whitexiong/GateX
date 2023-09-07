import { ref } from 'vue';
import { userMenu } from '@/services/userService'; // 假设你的服务函数在这个位置

export default function useMenu(userId) {
    const menuItems = ref([]);
    const error = ref(null);

    const fetchMenu = async () => {
        try {
            const response = await userMenu(userId);
            if (response.data) {
                menuItems.value = response.data;
            }
        } catch (err) {
            error.value = err.message || 'Error fetching menu';
        }
    };

    return {
        menuItems,
        error,
        fetchMenu
    };
}
