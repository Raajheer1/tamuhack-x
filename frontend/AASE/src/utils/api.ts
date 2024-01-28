import axios from "axios";

const apiUrl = "https://tamuhack.raajpatel.dev/"

export const API = axios.create({
    baseURL: apiUrl,
    // withCredentials: true,
});

export default apiUrl;