const baseURL = process.env.REACT_APP_BASE_URL;
const apiURL = process.env.REACT_APP_API_URL;
export const api = {
    retrieve_all_ssh_keypairs: apiURL + "/api/v1/key",
    change_password: apiURL + "/api/v1/user/change-password",
    login: apiURL + "/api/v1/user/login",
    create_ssh_keypair: apiURL + "/api/v1/key",
    delete_ssh_keypair: apiURL + "/api/v1/key/",
    logout: apiURL + "/api/v1/user/logout",
};

export const app = {
    view_ssh_keypair: baseURL + "/view/"
}