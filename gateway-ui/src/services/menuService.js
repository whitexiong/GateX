import axios from 'axios';

const BASE_URL = "http://127.0.0.1:8051";

export const getAllMenus = () => {
    return axios.get(`${BASE_URL}/menu/list`);
}

export const createMenu = (menuData) => {
    return axios.post(`${BASE_URL}/menu`, menuData);
}

export const getMenu = (id) => {
    return axios.get(`${BASE_URL}/menu/${id}`);
}

export const updateMenu = (id, menuData) => {
    return axios.put(`${BASE_URL}/menu/${id}`, menuData);
}

export const deleteMenu = (id) => {
    return axios.delete(`${BASE_URL}/menu/${id}`);
}
