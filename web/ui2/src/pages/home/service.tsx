import request from '~/utils/request'


export async function ListHotTopic() {
    return request("/topic/marks/hot")
}