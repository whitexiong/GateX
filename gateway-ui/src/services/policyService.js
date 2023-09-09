import axios from 'axios';

import { BASE_URL } from './constants';

export const getList = () => {
    return axios.get(`${BASE_URL}/policy/list`);
}

export const add = (PolicyData) => {
    return axios.post(`${BASE_URL}/policy/add`, PolicyData);
}

export const detail = (id) => {
    return axios.get(`${BASE_URL}/policy/detail/${id}`);
}

export const update = (id, PolicyData) => {
    return axios.put(`${BASE_URL}/policy/update/${id}`, PolicyData);
}

export const deletePolicy = (id) => {
    return axios.delete(`${BASE_URL}/policy/delete/${id}`);
}
