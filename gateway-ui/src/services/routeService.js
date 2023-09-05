import axios from 'axios';

const BASE_URL = "http://127.0.0.1:8051";

export const getList = () => {
    return axios.get(`${BASE_URL}/route/list`);
}

export const add = (data) => {
    return axios.post(`${BASE_URL}/route/add`, data);
}

export const detail = (id) => {
    return axios.get(`${BASE_URL}/route/detail/${id}`);
}

export const update = (id, data) => {
    return axios.post(`${BASE_URL}/route/update/${id}`, data);
}

export const deleted = (id) => {
    return axios.get(`${BASE_URL}/route/delete/${id}`);
}
