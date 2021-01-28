/*global chrome*/
import React from "react";
import {useState, useEffect} from "react";
import ReactDOM from "react-dom";
import "./App.css";
import "antd/dist/antd.css"; // or 'antd/dist/antd.less'
import {Input, Form, Button, Divider, message} from "antd";
import EditableTagGroup from "./comps/EditTagGroup";
import SearchInput from "./comps/SearchInput";
import {getInfo, saveLink, config} from "./service";

const {TextArea} = Input;
const LoginUrl = config.UrlPrefix + "/login?from=chrome"

function App() {
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
    chrome &&
      chrome.tabs &&
      chrome.tabs.query &&
      chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
        console.warn(tabs);
        let tab = tabs[0];
        form.setFieldsValue({
          title: tab.title,
        });
        setLink((link) => {
          return { ...link, url: tab.url, title: tab.title };
        });
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
    chrome.tabs.create({url: LoginUrl});
  };

  return (
    <div className="App">
      <Form
        onFinish={onFinish}
        name="control-hooks"
        onFinishFailed={onFinishFailed}
        form={form}
      >
        <Form.Item>
          <SearchInput
            placeholder="选择投稿的主题"
            onChange={(value) => {
              console.log(value);

              setLink((link) => {
                let newLink = { ...link, topicId: value };
                console.log(newLink);
                return newLink;
              });
            }}
          />
        </Form.Item>
        <Form.Item name="title" rules={[{ required: true, message: "标题" }]}>
          <TextArea placeholder="中国梦🇨🇳" autoSize />
        </Form.Item>
        <Form.Item
          name="comment"
          rules={[{ required: false, message: "输入评论" }]}
        >
          <TextArea rows={4} placeholder="说点什么吧" autoSize></TextArea>
        </Form.Item>
        <Form.Item name="tags">
          <EditableTagGroup
            onChange={(tags) => {
              console.log(tags.join(","));
              //⚠️：这里必须这样搞要不然会有并发异常
              setLink((link) => {
                let newLink = { ...link, tags: tags.join(",") };
                console.log(newLink);
                return newLink;
              });
            }}
          />
        </Form.Item>
        <Divider />

        <Form.Item>
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

      <a href="#" onClick={() => {
        chrome.tabs.create({url: config.UrlPrefix + "/"});
      }}>首页</a>
    </div>
  );
}

export default App;
