export interface User {
    id: number;
    name: string;
    avatar?: string;
    userName?:string;
    NickName?:string;
    link?: string;
    uid?: string;
}

export interface LoginUser extends User {
    token?: string

}
