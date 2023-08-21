import axios from 'axios';

const BASE_URL = "http://127.0.0.1:8051";

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

