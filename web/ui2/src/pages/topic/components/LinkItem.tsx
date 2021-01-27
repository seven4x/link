import React, {createElement, useContext, useState} from "react";
import styled from "styled-components";
import {DislikeOutlined, LikeOutlined, FireOutlined, DislikeFilled, LikeFilled} from '@ant-design/icons'
import {ReactComponent as CommentIcon} from "./comment.svg";
import CommentList from "./CommentList";
import {LinkItemData} from "../model";
import {GlobalContext} from "../../../App";
import Icon from '@ant-design/icons';
import {useRequest} from "ahooks";
import {Vote} from "../service";

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

    const [likes, setLikes] = useState(link.isLike === 1 ? 1 : 0);
    const [dislikes, setDislikes] = useState(link.isLike === 2 ? 1 : 0);
    const [action, setAction] = useState<number | null>(link.isLike);
    const {data, loading, run} = useRequest(Vote, {manual: true, debounceInterval: 500})

    const onLike = (item: LinkItemData) => {
        let t = (1 - likes)
        setLikes(t);
        setDislikes(0);
        let a = t > 0 ? 1 : 0
        setAction(a);
        if (loginContext == null || loginContext.user == null) {
            console.log('not login ,can not vote ')
            return
        }
        run('link', item.id, a).then(res => {
            console.log(res)
        })
    };
    const onDislike = (item: LinkItemData) => {
        let t = 1 - dislikes
        setLikes(0);
        setDislikes(t);
        let a = t > 0 ? 2 : 0
        setAction(a);
        if (loginContext == null || loginContext.user == null) {
            console.log('not login ,can not vote ')
            return
        }
        run('link', item.id, a).then(res => {
            console.log(res)
        })
    };
    return (
        <LinkItemWrapper>
            <Title>
                <Link href={link.url} target="_blank" rel="noreferrer noopener">{link.title}</Link>
            </Title>
            <FireOutlined/><Content>{link.hotComment && link.hotComment.context}</Content>
            <Controls>
                <Control onClick={() => {
                    onLike(link)
                }}>
                    {createElement((action === 1) ? LikeFilled : LikeOutlined)}
                    <span className="comment-action">{link.like + likes}</span>
                </Control>
                <Control onClick={() => {
                    onDislike(link)
                }}>
                    {createElement((action === 2) ? DislikeFilled : DislikeOutlined)}
                    <span className="comment-action">{link.dislike + dislikes}</span>
                </Control>

                <Control onClick={() => {
                    console.log(commentShow)
                    setCommentShow(!commentShow)
                }}>
                    <Icon component={CommentIcon}/><span>{link.comment}</span>
                </Control>
            </Controls>
            {
                commentShow && <CommentList linkId={link.id}/>
            }
        </LinkItemWrapper>

    );
};


export default React.memo(LinkItem)
