import axios from "axios";

const baseURL = "http://172.17.224.127:3030/auth";

const api = axios.create({
    baseURL,
    withCredentials: true,
});

const registerUser = async (data) => {
    const response = await api.post(
        "/register",
        data
    );
    return response.data;
};

const loginUser = async (data) => {
    const response = await api.post(
        "/login",
        data
    );
    return response.data;
};

const logoutUser = async (data) => {
    const response = await api.post(
        "/logout",
        data
    );
    return response.data;
};

export {
    registerUser,
    loginUser,
    logoutUser,
};