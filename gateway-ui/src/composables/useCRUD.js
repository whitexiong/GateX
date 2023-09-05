import {ref, reactive, toRefs} from 'vue';
import {ElMessageBox} from 'element-plus';

export const useCRUD = (apiMethods, initialData) => {
    const data = ref([]);
    const state = reactive({
        selected: {...initialData},
        isLoading: false,
        dialogVisible: false,
        isEditing: false,
        currentPage: 1,
        pageSize: 10,
        searchText: '',
        dialogTitle: '新增',
        error: null
    });

    const listData = async () => {
        state.isLoading = true;
        try {
            const response = await apiMethods.getList({
                name: state.searchText,
                pageSize: state.pageSize,
                currentPage: state.currentPage
            });
            data.value = response.data;
        } catch (error) {
            state.error = 'Failed to fetch data';
            console.error("Error:", error);
        } finally {
            state.isLoading = false;
        }
    };

    const processDataForRequest = (data) => {
        let processedData = {...data};
        if (Array.isArray(processedData.ParentID) && processedData.ParentID.length) {
            processedData.ParentID = processedData.ParentID[0];
        }
        return processedData;
    };

    const saveData = async () => {
        state.dialogVisible = true;
        let processedData = processDataForRequest(state.selected);
        try {
            if (state.isEditing) {
                await apiMethods.update(processedData.ID, processedData);
            } else {
                await apiMethods.add(processedData);
            }
            await listData();
            state.dialogVisible = false;
        } catch (error) {
            state.error = 'Failed to save data';
            console.error("Error:", error);
        } finally {
            state.dialogVisible = false;
        }
    };

    const getDetail = async (id) => {
        try {
            const detailResponse = await apiMethods.detail(id);

            if (detailResponse) {
                Object.keys(state.selected).forEach(key => {
                    if (key in detailResponse.data) {
                        state.selected[key] = detailResponse.data[key];
                    }
                });

                state.isEditing = true;
                state.dialogVisible = true;
                state.dialogTitle = '编辑节点';
            } else {
                console.error("Failed to fetch details.");
            }
        } catch (error) {
            state.error = "Error fetching details: " + error.message;
        }
    };

    //刷新
    const refresh = async () => {
        state.searchText = '';
        await listData();
    };

    //弹窗
    const addNew = () => {
        state.selected = {...initialData};
        state.dialogVisible = true;
        state.isEditing = false;
        state.dialogTitle = '新增';
    };

    const resetData = () => {
        state.selected = {...initialData};
    }


    const deleted = async (id) => {
        ElMessageBox.confirm('确定删除此项吗?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(async () => {
            try {
                await apiMethods.deletedById(id);
                await listData();
            } catch (error) {
                state.error = 'Failed to delete item';
            } finally {
                state.dialogVisible = false;
            }
        }).catch(() => {
            console.log('取消删除操作');
        });
    };

    const toggleStatus = async (row) => {
        row.Status = row.Status === 1 ? 0 : 1;
        state.selected = row;  // 确保 state.selected 是当前的 row 数据
        await saveData();
        state.dialogVisible = false;
    };

    const handlePageChange = async (newPage) => {
        state.currentPage = newPage;
        await listData();
    };

    const loadRoutes = async (row) => {
        if (row.children && row.children.length > 0) {
            return row.children;
        }
        return [];
    };

    return {
        data,
        ...toRefs(state),
        listData,
        saveData,
        refresh,
        addNew,
        getDetail,
        deleted,
        resetData,
        loadRoutes,
        handlePageChange,
        toggleStatus
    };
};
