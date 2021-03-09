import React, {useEffect, useState} from "react";
import styles from "./index.module.css";
import {Button, Divider, Form, Input, message} from "antd";
import EditableTagGroup from "./comps/EditTagGroup";
import SearchInput from "./comps/SearchInput";
import {getInfo, saveLink} from "./service";
import {useLocation, useHistory} from 'react-router-dom'
import queryString from 'query-string';

const {TextArea} = Input;
const layout = {
    labelCol: {span: 4},
    wrapperCol: {span: 16},
}
const tailLayout = {
    wrapperCol: {offset: 8, span: 16},
};

function Share() {
    const history = useHistory()
    const [form] = Form.useForm();
    const [link, setLink] = useState({
        title: "",
        url: "",
        comment: "",
        topicId: "",
        group: "",
        tags: "",
        from: 2,
    });
    const [isLogin, setIsLogin] = useState(false);
    const location = useLocation()
    const params = queryString.parse(location.search)
    const url = params.url as string
    let title = params.title as string

    useEffect(() => {
        getInfo()
            .then((res) => {
                if (res.ok) {
                    setIsLogin(true);
                } else {
                    setIsLogin(false);
                }
                console.log(res);
            })
            .catch((err) => {
                console.error(err);
            });

        form.setFieldsValue({
            title: title,
        });
        setLink((link) => {
            return {...link, url: url, title: title};
        });

        return () => {
            console.log("destory ....");
        };
    }, []);

    const onFinish = (values) => {
        console.log(values);

        //提交表单
        link.comment = values["comment"];
        link.title = values["title"];
        console.log(link);
        if (!isLogin) {
            setTimeout(() => {
                toLogin();
            }, 300);
            return;
        }
        saveLink(link).then((res) => {
            //close pop html
            if (!res.ok) {
                if (res.msgId == "link.repeat-in-same-topic") {
                    message.info("投稿成功❤️.");
                } else {
                    message.error("保存失败：" + res.msgId + "|" + res.msg);
                }
            } else {
                message.info("投稿成功❤️");
            }
        });
    };
    const onFinishFailed = (errorInfo) => {
        console.log("Failed:", errorInfo);
    };

    const toLogin = () => {
        history.push("/login", {from: location.pathname + location.search});
    };

    return (
        <div className={styles.container}>
            <h1>推荐到破茧</h1>
            <Form
                className={styles.form}
                {...layout}
                onFinish={onFinish}
                name="control-hooks"
                onFinishFailed={onFinishFailed}
                form={form}
            >
                <Form.Item
                    label="保存到主题"
                >
                    <SearchInput
                        placeholder="选择投稿的主题"
                        onChange={(value) => {
                            console.log(value);

                            setLink((link) => {
                                let newLink = {...link, topicId: value};
                                console.log(newLink);
                                return newLink;
                            });
                        }}
                    />
                </Form.Item>
                <Form.Item label="标题" name="title" rules={[{required: true, message: "标题"}]}>
                    <TextArea placeholder="🇨🇳" autoSize/>
                </Form.Item>
                <Form.Item
                    name="comment"
                    label="推荐理由"
                    rules={[{required: false, message: "输入推荐理由"}]}
                >
                    <TextArea placeholder="说点什么吧" autoSize={{minRows: 2}}/>
                </Form.Item>
                <Form.Item name="tags" label="标签" className={styles.tags}>
                    <EditableTagGroup
                        onChange={(tags) => {
                            console.log(tags.join(","));
                            //⚠️：这里必须这样搞要不然会有并发异常
                            setLink((link) => {
                                let newLink = {...link, tags: tags.join(",")};
                                console.log(newLink);
                                return newLink;
                            });
                        }}
                    />
                </Form.Item>

                <Form.Item {...tailLayout}>
                    {isLogin
                        ? <Button type="primary" htmlType="submit">
                            保存
                        </Button>
                        : <Button className="login-btn" onClick={toLogin}>
                            登陆
                        </Button>}
                </Form.Item>
            </Form>

            <Divider/>

        </div>
    );
}

export default Share;
