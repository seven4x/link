import request from '~/utils/request'


export async function Login(data:any) {

    return request('/account/login', {method: 'POST', data})
}


export async function GetUserInfo() {
    return request('/account/info')
}

export async function Logout() {
    return request("/account/logout")
}

export async function RegisterRequest(data: any) {
    return request('/account/register', {method: "POST", data})
}
