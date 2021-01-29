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
    agree: number;
    disagree: number;
    commentCount: number;
    hotComment?: Poster

}

export interface Poster {
    avatar: string
    content: string
    uid: number
}


//评论
export interface CommentData {
    id?: number
    creator: User
    agree: number
    disagree: number
    isLike?: number
    createTime: number
    content: string
}
