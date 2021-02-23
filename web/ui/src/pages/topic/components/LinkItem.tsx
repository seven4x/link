import React, {createElement, useContext, useState} from "react";
import styled from "styled-components";
import {DislikeOutlined, LikeOutlined, FireOutlined, DislikeFilled, LikeFilled} from '@ant-design/icons'
import {ReactComponent as CommentIcon} from "./assets/comment.svg";
import CommentList from "./CommentList";
import {LinkItemData} from "../model";
import {GlobalContext} from "../../../App";
import Icon from '@ant-design/icons';
import {useRequest} from "ahooks";
import {Vote} from "../service";
import {Space} from "antd";

const LinkItemWrapper = styled.div`
padding-top: 0.2rem;
padding-right: 0.2rem;
padding-left: 0.2rem;

`

const Title = styled.div`
line-height: 1.6;
font-weight: 500;
&:hover{
  font-size: 15px;
}
`

const Link = styled.a`
&:visited{
  color: #7a7a7a;
}
&:link{
}
`

const Content = styled.span`
  margin: .2rem;
  font-weight: 350;
`

const Controls = styled.ul`
margin-bottom: inherit;
padding-left: 0;
`

const Control = styled.li`
display: inline-block;
color: rgba(0, 0, 0, 0.45);
margin-right: 10px;
`


export interface LinkItemProps {
    link: LinkItemData
}

export const LinkItem: React.FC<LinkItemProps> = (props: LinkItemProps) => {
    let {link} = props
    let [commentShow, setCommentShow] = useState(false)
    const loginContext = useContext(GlobalContext)

    const [isLike, setIsLike] = useState<number | null>(link.isLike);
    const {data, loading, run} = useRequest(Vote, {manual: true, debounceInterval: 500})

    const onLike = (item: LinkItemData) => {
        if (isLike === 1) {
            link.agree = link.agree - 1
        } else if (isLike === 2) {
            link.agree = link.agree + 1
            link.disagree = link.disagree - 1
        } else if (isLike === 0) {
            link.agree = link.agree + 1
        }
        setIsLike((v) => {
            return v === 1 ? 0 : 1
        });
        if (loginContext == null || loginContext.user == null) {
            console.log('not login ,can not vote ')
            return
        }
        run('link', item.id, isLike === 1 ? 0 : 1).then(res => {
            console.log(res)
        })
    };
    const onDislike = (item: LinkItemData) => {
        if (isLike === 1) {
            link.disagree = link.disagree + 1
            link.agree = link.agree - 1
        } else if (isLike === 2) {
            link.disagree = link.disagree - 1
        } else if (isLike === 0) {
            link.disagree = link.disagree + 1
        }
        let newIsLike = isLike === 2 ? 0 : 2
        setIsLike((v) => {
            return v === 2 ? 0 : 2
        });

        if (loginContext == null || loginContext.user == null) {
            console.log('not login ,can not vote ')
            return
        }
        run('link', item.id, newIsLike).then(res => {
            console.log(res)
        })
    };
    const getIcon = (l) => {
        let u = new URL(l.link)
        return u.origin + "/favicon.ico"
    }
    return (
        <LinkItemWrapper>
            <Title>
                <Link href={link.link} target="_blank" rel="noreferrer noopener">{link.title}</Link>
            </Title>
            {link.hotComment && link.hotComment.content && <> <FireOutlined/><Content>{link.hotComment.content}</Content></>}
            <Controls>
                <Control onClick={() => {
                    onLike(link)
                }}>
                    {createElement((isLike === 1) ? LikeFilled : LikeOutlined)}
                    <span className="comment-action">{link.agree}</span>
                </Control>
                <Control onClick={() => {
                    onDislike(link)
                }}>
                    {createElement((isLike === 2) ? DislikeFilled : DislikeOutlined)}
                    <span className="comment-action">{link.disagree}</span>
                </Control>

                <Control onClick={() => {
                    console.log(commentShow)
                    setCommentShow(!commentShow)
                }}>
                    <Icon component={CommentIcon}/><span>{link.commentCount}</span>
                </Control>
            </Controls>
            {
                commentShow && <CommentList linkId={link.id}/>
            }
        </LinkItemWrapper>

    );
};


export default React.memo(LinkItem)
