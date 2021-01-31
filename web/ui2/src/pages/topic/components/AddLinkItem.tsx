import React, {useState} from 'react'
import {Button, Divider, Form, Input, Modal} from "antd";
import {useFormatMessage} from 'react-intl-hooks'
import {useKeyPress} from 'ahooks';
import useClipboard from 'react-hook-clipboard'
import {AddLink, GetUrlPreviewView} from '../service'
import {SaveMessages} from '../../../utils/message-util'
import {LinkItemData} from "../model";

function isUrl(url: string) {
    try {
        let u = new URL(url)
        console.log(u)
        return u
    } catch (e) {
        return null
    }
}

interface AddLinkItemProps {
    topicId: number
    afterAdd: (link: LinkItemData) => void
}

const AddLinkItem: React.FC<AddLinkItemProps> = ({topicId, afterAdd}) => {
    const t = useFormatMessage()
    const [form] = Form.useForm();
    const [visible, setVisible] = useState<boolean>(false)
    const [saving, setSaving] = useState<boolean>(false)
    const [prefix, setPrefix] = useState(null)

    let tmp = t({id: "topic.button.add-link"}).toString();
    const [title, setTitle] = useState<string>(tmp)

    const handleClipboardReadError = error => {
        console.log(
            'There was an error reading from the clipboard:',
            error
        )
    }
    const [clipboard, copyToClipboard] = useClipboard({}, handleClipboardReadError)

    useKeyPress(['ctrl.v', "meta.v"], (e) => {
        if (!visible && isUrl(clipboard)) {
            showModal()
        }
    })

    const layout = {
        labelCol: {span: 6},
        wrapperCol: {flex: "auto"},
    };

    const showModal = () => {
        setVisible(true);
        console.log(clipboard)
        let u = isUrl(clipboard)
        if (u) {
            form.setFieldsValue({url: clipboard})
            setPrefix(<img width="14px" height="14px" src={`${u.origin}/favicon.ico`}/>)
            //  请求preview-link获取title
            GetUrlPreviewView(clipboard).then(res => {
                form.setFieldsValue({title: res && res.title})
                console.log(res)
            })
        } else {
            setPrefix(null)
        }

    };
    const handleOk = () => {
        setSaving(true);
        form.validateFields()
            .then(values => {
                console.log(values)
                values.topicId = topicId
                AddLink(values).then(res => {
                    if (res.ok) {
                        setVisible(false)
                        form.resetFields();
                        let link = {...values, id: res.data}
                        afterAdd(link as LinkItemData)
                    }
                    SaveMessages(res, t)
                    //todo add to list ,避免刷新查看
                })

            })
            .catch(info => {
                console.log('Validate Failed:', info);
            })
            .finally(() => {
                setSaving(false);
            });
    };
    const handleCancel = () => {
        console.log('Clicked cancel button');
        setVisible(false)
    };
    //todo 当URL input变化时，重新设置title


    return (
        <>
            <Button onClick={showModal}>{t({id: "topic.button.add-link"})}</Button>
            <Modal
                title={title}
                visible={visible}
                onOk={handleOk}
                confirmLoading={saving}
                onCancel={handleCancel}
            >
                <Form
                    {...layout}
                    form={form}
                    name="form_in_modal"
                >
                    <Form.Item label={t({id: "topic.form.link-url"})} name="url"
                               rules={[{required: true}, {type: "url"}]}
                    >
                        <Input placeholder="url" type="url" prefix={prefix}/>
                    </Form.Item>

                    <Form.Item label={t({id: "topic.form.link-title"})} name="title"
                               rules={[{required: true}]}>
                        <Input placeholder="title"/>
                    </Form.Item>

                    <Form.Item label={t({id: "topic.form.link-comment"})} name="comment"
                               rules={[{max: 140}]}>
                        <Input.TextArea autoSize={{minRows: 2}}/>
                    </Form.Item>
                    <Form.Item label={t({id: "topic.form.link-group"})} name="group"
                               rules={[{max: 24}]}>
                        <Input placeholder="group"/>
                    </Form.Item>
                    <Divider/>
                    <Form.Item label="tips">
                        <Form.Item noStyle>
                            {t({id: "topic.form.link-tip"})}
                        </Form.Item>
                    </Form.Item>
                </Form>
            </Modal>
        </>
    )
}

export default React.memo(AddLinkItem)