import {User} from "../../model/user";

export interface Topic {
    name?: string
    id?: number
    icon?: string
}

export interface NewTopic{
    name:string
    refId:number
    position:string
    refDesc:string
    tags:string
}

//链接
export interface LinkItemData {
    loading?: boolean;
    id?: number;
    isLike?: number;
    title: string;
    url: string;
    like: number;
    dislike: number;
    comment: number;
    hotComment?: Poster

}

export interface Poster {
    avatar: string
    context: string
    uid: number
}


//评论
export interface CommentData {
    id?: number
    user: User
    like: number
    dislike: number
    isLike?: number
    createTime?: string
    content: string
}