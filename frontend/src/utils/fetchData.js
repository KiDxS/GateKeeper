import axios from "axios";

// fetchData is a function that is used throughout the codebase to send a GET request to the backend.
export const fetchData = async (url, options) => {
  const response = await axios.get(url, options);
  return response;
};
