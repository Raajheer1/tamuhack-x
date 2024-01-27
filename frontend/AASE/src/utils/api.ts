import axios from "axios";

const apiUrl = "http://localhost:3000/"

export const API = axios.create({
    baseURL: apiUrl,
    // withCredentials: true,
});

export default apiUrl;