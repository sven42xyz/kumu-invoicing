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

const logoutUser = async () => {
    const response = await api.post(
        "/logout",
    );
    return response.data;
};

const accessControll = async () => {
    const response = await api.post(
        "/access",
    );
    return response.data;
};

export {
    registerUser,
    loginUser,
    logoutUser,
    accessControll,
};