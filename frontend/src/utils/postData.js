import axios from "axios";

// postData is a function that is used throughout the codebase to send a POST request to the backend.
export const postData = async (url, data, options) => {
  const response = await axios.post(url, data, options);
  return response;
};
