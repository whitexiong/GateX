import axios from 'axios';

const BASE_URL = "http://127.0.0.1:8051";

export const getList = () => {
    return axios.get(`${BASE_URL}/route/list`);
}

export const createMenu = (menuData) => {
    return axios.post(`${BASE_URL}/route`, menuData);
}

export const getMenu = (id) => {
    return axios.get(`${BASE_URL}/route/${id}`);
}

export const updateMenu = (id, menuData) => {
    return axios.put(`${BASE_URL}/route/${id}`, menuData);
}

export const deleteMenu = (id) => {
    return axios.delete(`${BASE_URL}/route/${id}`);
}
