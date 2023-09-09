import axios from 'axios';

import { BASE_URL } from './constants';


export const getList = () => {
    return axios.get(`${BASE_URL}/menu/list`);
}

export const add = (menuData) => {
    return axios.post(`${BASE_URL}/menu/add`, menuData);
}

export const detail = (id) => {
    return axios.get(`${BASE_URL}/menu/detail/${id}`);
}

export const update = (id, menuData) => {
    return axios.post(`${BASE_URL}/menu/update/${id}`, menuData);
}

export const deletedById = (id) => {
    return axios.get(`${BASE_URL}/menu/delete/${id}`);
}
