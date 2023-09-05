import axios from 'axios';

const BASE_URL = "http://127.0.0.1:8051";

export const getList = () => {
    return axios.get(`${BASE_URL}/menu/list`);
}

export const create = (menuData) => {
    return axios.post(`${BASE_URL}/menu/add`, menuData);
}

export const getDetail = (id) => {
    return axios.get(`${BASE_URL}/menu/detail/${id}`);
}

export const update = (id, menuData) => {
    return axios.post(`${BASE_URL}/menu/update/${id}`, menuData);
}

export const deleted = (id) => {
    return axios.get(`${BASE_URL}/menu/delete/${id}`);
}
