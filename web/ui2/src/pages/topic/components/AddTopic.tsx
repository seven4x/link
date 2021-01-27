import React, {useState} from 'react'
import {Button, Modal, Form, Input, Select, Radio, message} from "antd";
import {useFormatMessage} from "react-intl-hooks";
import {Topic} from "../model";
import {AllPosition} from "../../../utils/const";
import {SaveTopic} from '../service'
import {SaveMessages} from '../../../utils/message-util'
export interface AddTopicProps {
    topic: Topic,
}

const AddTopic: React.FC<AddTopicProps> = (props) => {
    let {topic} = props
    const t = useFormatMessage()
    const [form] = Form.useForm();

    const [visible, setVisible] = useState<boolean>(false)
    const [saving, setSaving] = useState<boolean>(false)
    const title = t({id: "topic.title.add-topic"}, {name: topic.name});
    const children = [];
    const layout = {
        labelCol: {span: 6},
        wrapperCol: {flex: "auto"},
    };
    const showModal = () => {
        setVisible(true);
    };

    function handleOk() {
        setSaving(true);
        form.validateFields()
            .then(values => {
                console.log(values)
                values.refId = topic.id
                SaveTopic(values).then(res => {
                    if (res.ok) {
                        setVisible(false)
                        form.resetFields();
                    }
                    SaveMessages(res, t)
                })
            })
            .catch(info => {
                console.log('Validate Failed:', info);
            })
            .finally(() => {
                setSaving(false);
            });

    }

    const handleCancel = () => {
        console.log('Clicked cancel button');
        setVisible(false)
    };

    function handleTagChange(value) {
        console.log(`selected ${value}`);
    }

    return (

        <>
            <Button onClick={showModal}>{t({id: "topic.button.add-topic"})}</Button>
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
                    initialValues={{position: "bottom"}}
                >
                    <Form.Item label={t({id: "topic.form.topic-name"})} name="name"
                               rules={[{required: true}]}>
                        <Input placeholder="new topic name"/>
                    </Form.Item>
                    <Form.Item label={t({id: "topic.form.topic-position"})} name="position">
                        <Radio.Group>
                            {
                                AllPosition.map((p, i) => {
                                    return <Radio.Button value={p} key={i}>{t({id: `topic.radio.${p}`})}</Radio.Button>
                                })
                            }
                        </Radio.Group>
                    </Form.Item>
                    <Form.Item label={t({id: "topic.form.topic-tags"})} name="tags">
                        <Select mode="tags" style={{width: '100%'}} onChange={handleTagChange} tokenSeparators={[',']}
                                maxTagCount={10}>
                            {children}
                        </Select>
                    </Form.Item>
                    <Form.Item label={t({id: "topic.form.topic-refDesc"})} name="refDesc"
                               rules={[{max: 140}]}>
                        <Input placeholder=""/>
                    </Form.Item>
                </Form>
            </Modal>
        </>

    )
}

export default  React.memo(AddTopic)