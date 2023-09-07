import axios from 'axios';

const BASE_URL = "http://127.0.0.1:8051";

export const getList = () => {
    return axios.get(`${BASE_URL}/role/list`);
}

export const add = (roleData) => {
    return axios.post(`${BASE_URL}/role/add`, roleData);
}

export const detail = (id) => {
    console.log("详情id", id)
    return axios.get(`${BASE_URL}/role/detail/${id}`);
}

export const update = (id, roleData) => {
    return axios.post(`${BASE_URL}/role/update/${id}`, roleData);
}

export const deletedRole = (id) => {
    return axios.delete(`${BASE_URL}/role/delete/${id}`);
}

export const getPermissions = () => {
    return axios.get(`${BASE_URL}/role/permissions`);
}
