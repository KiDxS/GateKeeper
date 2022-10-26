import axios from "axios";

export const fetchData = async (url, options) => {
    const response = await axios.get(url, options);
    const data = response.json();
    return data;
};
