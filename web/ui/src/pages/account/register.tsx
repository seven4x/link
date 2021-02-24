import React, {useEffect} from "react";

import styles from './register.module.css'
import {Button, Form, Input, message} from "antd";
import {useHistory, useLocation} from "react-router-dom";
import queryString from 'query-string'
import {RegisterRequest} from './service'

const layout = {
    labelCol: {span: 8},
    wrapperCol: {span: 16},
};
const tailLayout = {
    wrapperCol: {offset: 8, span: 16},
};
const Register: React.FC = () => {
    const [form] = Form.useForm();
    let location = useLocation()
    const history = useHistory()
    const onFinish = async (values: any) => {
        console.log(values)
        let {ok} = await RegisterRequest(values)
        if (ok) {
            message.success("注册成功,欢迎您")
            history.replace("/login")
        } else {
            message.warn("邀请码错误")
        }
    }
    useEffect(() => {
        console.log(location.search)
        const parsed = queryString.parse(location.search)
        form.setFieldsValue({code: parsed.code})
    }, [])

    return (
        <div className={styles.container}>
            <div className={styles.left}>

            </div>

            <div>
                <Form {...layout} name="base"
                      form={form}
                      onFinish={onFinish}
                >
                    <Form.Item label="邀请码" name="code" rules={[{required: true}]}>
                        <Input placeholder="seven4x"/>
                    </Form.Item>
                    <Form.Item label="登录账号" name="loginId" rules={[{required: true, min: 4,max:32}]}>
                        <Input/>
                    </Form.Item>

                    <Form.Item label="密码" name="password" rules={[{required: true, min: 6}]}>
                        <Input.Password/>
                    </Form.Item>
                    <Form.Item {...tailLayout}>
                        <Button type="primary" htmlType="submit">注册</Button>
                    </Form.Item>
                </Form>
            </div>
        </div>
    )

}

export default Register
