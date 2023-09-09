import axios from 'axios';

import { BASE_URL } from './constants';

export const getDashboardData = () => {
    return axios.get(`${BASE_URL}/dashboard`);
}

export const UserLogin = (username, password) => {
    return axios.post(`${BASE_URL}/user/login`, {
        username: username,
        password: password
    });
}

export const UserLogout = () => {
    return axios.post(`${BASE_URL}/user/logout`);
}

export const getList = () => {
    return axios.get(`${BASE_URL}/user/list`);
}

export const add = (userData) => {
    return axios.post(`${BASE_URL}/user/add`, userData);
}

export const detail = (userId) => {
    return axios.get(`${BASE_URL}/user/detail/${userId}`);
}

export const update = (userId, updatedData) => {
    return axios.post(`${BASE_URL}/user/update/${userId}`, updatedData);
}

export const deletedById = (userId) => {
    return axios.get(`${BASE_URL}/user/delete/${userId}`);
}

export const GetUserMenus = () => {
    return axios.get(`${BASE_URL}/user/menus`);
}
