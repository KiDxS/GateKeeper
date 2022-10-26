import axios from "axios";



export const postData = async (url, data, options) => {
    const response = await axios.post(url, data, options);
    return response;
};
