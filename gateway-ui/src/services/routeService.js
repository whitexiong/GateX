import axios from 'axios';

import { BASE_URL } from './constants';

export const getList = (params) => {
    return axios.get(`${BASE_URL}/route/list`, { params });
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

export const deletedById = (id) => {
    return axios.get(`${BASE_URL}/route/delete/${id}`);
}
export const getPathList = (query) => {
    return axios.get(`${BASE_URL}/route/paths/${encodeURIComponent(query)}`);
}

