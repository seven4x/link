import React, {createElement, useContext, useState} from 'react'
import {GlobalContext} from "../../../App";
import {CommentData} from "../model";
import {Comment, Tooltip} from "antd";
import {DislikeFilled, DislikeOutlined, LikeFilled, LikeOutlined} from "@ant-design/icons";
import {DeleteComment, Vote} from '../service'
import {DeleteMessages} from '../../../utils/message-util'
import {useFormatMessage} from "react-intl-hooks";
import {useRequest} from "ahooks";

export interface CommentItemProps {
    linkId: number
    item: CommentData
    afterDelete: (mid: number) => void
}

//0 不投票，1喜欢，2不喜欢
const CommentItem: React.FC<CommentItemProps> = (props) => {

    let {item, linkId, afterDelete} = props
    const [likes, setLikes] = useState(item.isLike === 1 ? 1 : 0);
    const [dislikes, setDislikes] = useState(item.isLike === 2 ? 1 : 0);
    const [action, setAction] = useState<number | null>(item.isLike);
    const loginContext = useContext(GlobalContext)
    const t = useFormatMessage()
    const {data, loading, run} = useRequest(Vote, {manual: true, debounceInterval: 500})

    const onLike = (item: CommentData) => {
        let t = 1 - likes
        setDislikes(0);
        setLikes(t);
        let a = t > 0 ? 1 : 0
        setAction(a)
        if (loginContext == null || loginContext.user == null) {
            console.log('not login ,can not vote ')
            return
        }
        if (item.id) {
            run('comment', item.id, a).then(res => {
                console.log(res)
            })
        }
    };
    const onDislike = (item: CommentData) => {
        setLikes(0);
        let t = 1 - dislikes;
        setDislikes(t);
        let a = t > 0 ? 2 : 0
        setAction(a)
        if (loginContext == null || loginContext.user == null) {
            console.log('not login ,can not vote ')
            return
        }
        if (item.id) {
            run('comment', item.id, a).then(res => {
                console.log(res)
            })
        }

    };
    const handleDelete = (item: CommentData) => {
        if (item.id) {
            console.log("remove:" + linkId + "," + item.id)
            DeleteComment(linkId, item.id).then(res => {
                DeleteMessages(res, t)
                if (res.ok) {
                    afterDelete(item.id)
                }
            })
        }
    }
    const getActions = (item: CommentData) => {
        return [
            <Tooltip key="comment-basic-like" title="Like">
                                      <span onClick={() => {
                                          onLike(item)
                                      }}>
                                        {createElement((action === 1) ? LikeFilled : LikeOutlined)}
                                          <span className="comment-action">{item.like + likes}</span>
                                      </span>
            </Tooltip>,
            <Tooltip key="comment-basic-dislike" title="Dislike">
                                  <span onClick={() => {
                                      onDislike(item)
                                  }}>
                                    {React.createElement((action === 2 && dislikes === 1) ? DislikeFilled : DislikeOutlined)}
                                      <span className="comment-action">{item.dislike + dislikes}</span>
                                  </span>
            </Tooltip>,
            <GlobalContext.Consumer>
                {loginInfo => {
                    if (loginInfo && loginInfo.user && loginInfo.user.id === item.user.id) {
                        return <span onClick={() => {
                            handleDelete(item)
                        }}>
                            <span className="comment-action">删除</span>
                            </span>
                    }
                }
                }

            </GlobalContext.Consumer>

        ]
    }

    return (
        <Comment
            actions={getActions(item)}
            author={item.user.name}
            avatar={item.user.avatar}
            content={<p>{item.content}</p>}
            datetime={item.createTime}
            key={item.id}
        />)
}

export default React.memo(CommentItem)