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
