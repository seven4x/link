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
    data: CommentData
    afterDelete: (mid: number) => void
}

//0 不投票，1喜欢，2不喜欢
const CommentItem: React.FC<CommentItemProps> = (props) => {

    let {data, linkId, afterDelete} = props
    const [isLike, setIsLike] = useState<number | null>(data.isLike);
    const loginContext = useContext(GlobalContext)
    const t = useFormatMessage()
    const {run} = useRequest(Vote, {manual: true, debounceInterval: 500})

    const onLike = (item: CommentData) => {
        if (isLike === 1) {
            item.agree = item.agree - 1
        } else if (isLike === 2) {
            item.agree = item.agree + 1
            item.disagree = item.disagree - 1
        } else if (isLike === 0) {
            item.agree = item.agree + 1
        }
        setIsLike(v => {
            return v === 1 ? 0 : 1
        })
        if (loginContext == null || loginContext.user == null) {
            console.log('not login ,can not vote ')
            return
        }
        if (item.id) {
            run('comment', item.id, isLike === 1 ? 0 : 1).then(res => {
                console.log(res)
            })
        }
    };
    const onDislike = (item: CommentData) => {
        if (isLike === 1) {
            item.disagree = item.disagree + 1
            item.agree = item.agree - 1
        } else if (isLike === 2) {
            item.disagree = item.disagree + 1
        } else if (isLike === 0) {
            item.disagree = item.disagree + 1
        }
        let newIsLike = isLike === 2 ? 0 : 2
        setIsLike((v) => {
            return v === 2 ? 0 : 2
        });
        if (loginContext == null || loginContext.user == null) {
            console.log('not login ,can not vote ')
            return
        }
        if (item.id) {
            run('comment', item.id, newIsLike).then(res => {
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
            <Tooltip title="Like">
                                      <span onClick={() => {
                                          onLike(item)
                                      }}>
                                        {createElement((isLike === 1) ? LikeFilled : LikeOutlined)}
                                          <span className="comment-action">{item.agree}</span>
                                      </span>
            </Tooltip>,
            <Tooltip title="Dislike">
                                  <span onClick={() => {
                                      onDislike(item)
                                  }}>
                                    {React.createElement((isLike === 2) ? DislikeFilled : DislikeOutlined)}
                                      <span className="comment-action">{item.disagree}</span>
                                  </span>
            </Tooltip>,
            <GlobalContext.Consumer>
                {loginInfo => {
                    if (loginInfo && loginInfo.user && loginInfo.user.id === item.creator.id) {
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
            actions={getActions(data)}
            author={data.creator.name}
            avatar={data.creator.avatar}
            content={<p>{data.content}</p>}
            datetime={data.createTime}
            key={data.id}
        />)
}

export default React.memo(CommentItem)
