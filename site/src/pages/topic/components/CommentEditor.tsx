import React, {useContext, useState} from "react";
import {GlobalContext} from "../../../App";
import {Avatar, Button, Comment, Form, Input} from "antd";
import styled from 'styled-components'
import {AddComment} from '../service'
import {SaveMessages} from '../../../utils/message-util'
import {useFormatMessage} from "react-intl-hooks";

const {TextArea} = Input;

const Editor = ({onChange, onSubmit, submitting, value}: any) => (
    <>
        <Form.Item>
            <TextArea rows={1} onChange={onChange} value={value} autoSize={{minRows: 2}}/>
        </Form.Item>
        <Form.Item>
            <Button htmlType="submit" loading={submitting} onClick={onSubmit} type="primary">
                留言
            </Button>
        </Form.Item>
    </>
);
const Wrapper = styled.div`
  position: sticky;
  top: 0;
  bottom: 0;
  z-index: 998;
  background-color: #ffffff;
`

interface CommentEditorProps {
    linkId: number
    afterAdd: (comment: any) => void
}

const EditorWrapper: React.FC<CommentEditorProps> = (props) => {
    const {linkId, afterAdd} = props
    const [submitting, setSubmitting] = useState<boolean>(false)
    const [newComment, setNewComment] = useState<string>("")
    const loginContext = useContext(GlobalContext)
    const t = useFormatMessage()
    const handleSubmit = () => {
        if (!newComment) {
            console.log("newcomment is null ")
            return;
        }

        setSubmitting(true);
        AddComment(linkId, newComment).then(res => {
            if (res.ok) {
                let comment = {
                    id: res.data,
                    content: newComment,
                    creator: loginContext.user,
                    agree: 0,
                    disagree: 0,
                    createTime: Date.now() / 1000,
                    isLike: 0
                }
                afterAdd(comment)

            }
            setSubmitting(false)
            SaveMessages(res, t)
        })
    };
    const handleChange = (e: any) => {
        setNewComment(e.target.value);
    };
    return (
        <>
            <GlobalContext.Consumer>
                {
                    loginInfo => {
                        if (loginInfo == null || loginInfo.user == null) {
                            return null
                        }
                        return <Wrapper>
                            <Comment
                                avatar={
                                    <Avatar
                                        src={loginInfo.user.avatar}
                                        alt={loginInfo.user.name}
                                    />
                                }
                                content={
                                    <Editor
                                        onChange={handleChange}
                                        onSubmit={handleSubmit}
                                        submitting={submitting}
                                        value={newComment}
                                    />
                                }
                            />
                        </Wrapper>
                    }
                }
            </GlobalContext.Consumer>

        </>
    )

}


export default React.memo(EditorWrapper)
