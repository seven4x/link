import {message as antMessage} from "antd";


export function SaveMessages(response: any, t: any) {
    messages(response, t, "global.success.save", "global.error.save")
}

export function DeleteMessages(response: any, t: any) {
    messages(response, t, "global.success.delete", "global.error.delete")
}

/**
 * 前后端约定,失败返回数据结构：
 * {
 *     msg:"文案id"，
 *     data:{a:1,b:2}
 * }
 */
function messages(response: any, t: any, successId: string, errorId: string) {
    if (response.ok) {
        antMessage.success(t({id: successId}));
    } else {
        if (response.msgId) {
            antMessage.error(t({id: response.msgId}, response.errorData));
        }else if (response.msg){
            antMessage.error(response.msg)
        } else {
            antMessage.error(t({id: errorId}))
        }
    }
}