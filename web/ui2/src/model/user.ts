export interface User {
    id: number;
    name: string;
    avatar?: string
    link?: string;
    uid?: string;
}

export interface LoginUser extends User {
    token?: string

}
