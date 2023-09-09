import axios from 'axios';

import { BASE_URL } from './constants';

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
    return axios.post(`${BASE_URL}/role/update/${id}`, roleData);
}

export const deletedRole = (id) => {
    return axios.delete(`${BASE_URL}/role/delete/${id}`);
}

export const getPermissions = () => {
    return axios.get(`${BASE_URL}/role/permissions`);
}
