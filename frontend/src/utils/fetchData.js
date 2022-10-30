import axios from "axios";

export const fetchData = async (url, options) => {
  const response = await axios.get(url, options);
  return response;
};
