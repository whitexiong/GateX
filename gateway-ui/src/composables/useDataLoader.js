import {ref} from 'vue';

export const loadList = (getListFunction) => {
    const Routes = ref([]);
    const searchText = ref('');

    const listData = async () => {
        try {
            const response = await getListFunction({ name: searchText.value });
            Routes.value = Array.isArray(response.data) ? response.data : [];
        } catch (error) {
            console.error("Error fetching Routes:", error);
        }
    };

    return {
        Routes,
        searchText,
        listData
    };
};
