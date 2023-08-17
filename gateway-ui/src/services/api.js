import axios from 'axios';

const BASE_URL = "http://127.0.0.1:8051";

export const getDashboardData = () => {
    return axios.get(`${BASE_URL}/dashboard`);
}
