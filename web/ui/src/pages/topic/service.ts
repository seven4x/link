import request from '../../utils/request';

import {LocalLPVToken} from '../../utils/const'

export async function GetPreviewToken() {
    return request(`/link/preview-token`);
}

interface PreviewResult {
    title: string
    description: string
    image: string
    url: string
}

export async function GetUrlPreviewView(q: string) {
    let token = LocalLPVToken[0]

    return request(`/lpv?key=${token}&q=${q}`, {prefix: null})
}


export async function SaveTopic(topic: any) {
    if (topic.tags != null) {
        topic.tags = topic.tags.join(",")
    }
    return request("/topic", {method: "POST", data: topic})
}

export async function GetTopicDetail(id: number) {
    return request(`/topic/${id}`)
}

export async function ListRelationTopic(topicId: number, position: string, prev = 0) {
    let p
    switch (position) {
        case 'top':
            p = '1';
            break;
        case 'bottom':
            p = '2';
            break;
        case 'left':
            p = '3';
            break;
        case 'right':
            p = '4';
            break;
        default:
            p = '1';
    }
    return request(`/topic/${topicId}/related/${p}`, {method: "GET", params: {prev}})
}

export async function SearchTopic(q: string) {
    return request('/topic', {
        method: "GET",
        params: {
            q: q
        }
    })
}

export function AddLink(link: any) {
    link.from = 1
    if (link.group != null) {
        link.group = link.group.join(",")
    }
    return request("/link", {method: "POST", data: link})
}

export async function ListLinks(topicId: number, page: number, filter: string, group?: string) {
    if (topicId == null) {
        return Promise.resolve({data: [], ok: true})
    }
    let url
    switch (filter) {
        case 'hot':
            url = '/link/marks/hot'
            break
        case 'newest':
            url = '/link/marks/newest'
            break
        case 'mine':
            url = '/link/marks/mine'
            break
        default:
            url = '/link'
    }
    return request(url, {
        params: {
            tid: topicId,
            page,
            group
        }
    })

}

export async function AddComment(lid: number, comment: any) {
    let data = {
       content:comment
    }
    return request(`/link/${lid}/comment`, {method: "POST", data})
}

export function ListComment(lid: number, prev: number, sortBy = 'hot') {
    return request(`/link/${lid}/comment`, {method: "GET", params: {sortBy, prevNo: prev}})
}

/**
 * 删除评论
 * @param lid 链接ID
 * @param mid 评论ID
 */
export function DeleteComment(lid: number, mid: number) {
    return request(`/link/${lid}/comment/${mid}`, {method: "DELETE"})
}

export function ListMvpUser(topicId: number) {
    return request(`/user/marks/mvp`, {method: "GET", params: {topicId}})

}

export async function Vote(type: string, id: number, action: number) {
    return request("/vote", {method: "POST", data: {typeCode: type, id: id, isLike: action}})
}
