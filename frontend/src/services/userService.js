import axios from 'axios';

const API_URL = 'http://localhost:8080';

export const createUser = async (user) => {
    const response = await axios.post(`${API_URL}/users`, user);
    return response.data;
};

export const getUserById = async (id) => {
    const response = await axios.get(`${API_URL}/users/${id}`);
    return response.data;
};

export const getAllUsers = async () => {
    const response = await axios.get(`${API_URL}/users`);
    return response.data;
};

export const deleteUser = async (id) => {
    await axios.delete(`${API_URL}/users/${id}`);
};
