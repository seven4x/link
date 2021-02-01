import React, {useContext} from 'react'
import {Form, Input, Button, Checkbox, message, Space} from 'antd';
import styled from 'styled-components'
import {
  useHistory, useLocation
} from "react-router-dom";
import styles from './login.module.css'
import {GlobalContext} from "../../App";
import {GetUserInfo, Login as LoginRequest} from "./service";

const Left = styled.div`

`
const Right = styled.div`
  margin-top: 20px;
  width: 50vw;
`


const layout = {
  labelCol: {span: 8},
  wrapperCol: {span: 10},
};
const tailLayout = {
  wrapperCol: {offset: 8, span: 10},
};

const Login = () => {
  const globalContext = useContext(GlobalContext)
  const setLoginUser = globalContext.login
  let location = useLocation();
  let history = useHistory();

  // @ts-ignore
  let { from } = location.state || {from: {pathname: "/"}};
  console.log(from)
  const onFinish = (values: any) => {
    LoginRequest(values).then(res => {
      console.log(res)
      if (res.ok) {
        GetUserInfo().then(info => {
          console.log(info)
          if (info.ok) {
            setLoginUser(info.data)
            history.replace(from)
          }
        })
      } else {
        message.error("登陆失败，账号或密码错误")
      }

    }).catch(e => {
      console.warn(e)
    })
    console.log('Success:', values);
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };

  return (
      <div className={styles.container}>
        <Left>

        </Left>
        <Right>
          <Form
              {...layout}
              name="basic"
              initialValues={{remember: true}}
              onFinish={onFinish}
              onFinishFailed={onFinishFailed}
          >
            <Form.Item
                label="账号"
                name="username"
                rules={[{required: true, }]}
            >
              <Input/>
            </Form.Item>

            <Form.Item
                label="密码"
                name="password"
                rules={[{required: true,  }]}
            >
              <Input.Password/>
            </Form.Item>


            <Form.Item {...tailLayout}>
              <Space>
                <Button type="primary" htmlType="submit">
                  登陆
                </Button>
                <Button onClick={()=>{
                  history.push("/register")
                }}>去注册</Button>
              </Space>


            </Form.Item>
          </Form>
        </Right>

      </div>

  );
};

export default Login;
