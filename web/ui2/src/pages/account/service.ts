import request from '~/utils/request'


export async function Login(user: string, password: string) {
    let form = {
        username: user,
        password
    }
    return request('/account/login', {method: 'POST', data: form})
}


