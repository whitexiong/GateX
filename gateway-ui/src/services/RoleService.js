import axios from 'axios';

const BASE_URL = "http://127.0.0.1:8051";

export const getList = () => {
    return axios.get(`${BASE_URL}/role/list`);
}

export const add = (roleData) => {
    return axios.post(`${BASE_URL}/role/add`, roleData);
}

export const detail = (id) => {
    return axios.get(`${BASE_URL}/role/detail/${id}`);
}

export const update = (id, roleData) => {
    return axios.put(`${BASE_URL}/role/update/${id}`, roleData);
}

export const deleteRole = (id) => {
    return axios.delete(`${BASE_URL}/role/delete/${id}`);
}

export const getPermissionsFromAPI = () => {
    return axios.get(`${BASE_URL}/role/permissions`);
}

export const addRoleWithPermissions = (roleData, permissions) => {
    // 将权限数据添加到角色数据中
    const completeRoleData = {
        ...roleData,
        permissions: permissions
    };
    return axios.post(`${BASE_URL}/role/add`, completeRoleData);
}
