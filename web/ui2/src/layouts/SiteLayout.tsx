import React, {PropsWithChildren, useContext, useState} from 'react'
import {Redirect, Route, Switch} from "react-router-dom";
import NotFoundPage from "../components/404";
import {RouteWithSubRoutes} from '../pages/routes'
import {AutoComplete, Badge, Input, Layout, Row, Col, Divider, Button, Select} from "antd";
import LocaleSwitch from "../components/LocaleSwitch/LocaleSwitch";
import Profile from "../components/Profile/Profile";
import styled from "styled-components";
import {SelectProps} from "antd/es/select";
import logo from '~/assets/logo.png'
import {GlobalContext} from "../App";
import {SearchTopic} from '../pages/topic/service'
import {useRequest} from 'ahooks'
import {SearchOutlined} from '@ant-design/icons'
import {useHistory} from 'react-router-dom'

const {Header, Footer, Content, Sider} = Layout
const {Option} = Select;


const HeaderWrapper = styled(Header)`
background: #fff;
height:64px;
position: relative;
z-index: 10;
max-width: 100%;
box-shadow: 0 2px 8px #f0f1f2;
`


const Message = styled.a`
width: 42px;
height: 42px;
border-radius: 2px;
background: #eee;
display: inline-block;
vertical-align: middle;
`

const Logo = styled.a`
height: 64px;
overflow: hidden;
color: rgba(0,0,0,.85);
font-size: 18px;
line-height: 64px;
white-space: nowrap;
text-decoration: none;
& img{
    position: relative;
    top: -1.5px;
    height: 32px;
}
`


const FooterWrapper = styled(Footer)`
background-color: #F5F5F5;
padding-left: 176px;
`





const SiteLayout: React.FC<PropsWithChildren<any>> = (props) => {
    const globalContext = useContext(GlobalContext)
    let {routes} = props
    const [options, setOptions] = useState<Array<any>>([]);
    const {data, loading, run} = useRequest(SearchTopic, {
        manual: true, debounceInterval: 500,
        onSuccess: (res) => {
            let d = dualSearchResult(res)
            setOptions(d);
        }
    })
    const history = useHistory()

    const handleSearch = (value: string) => {
        if (value) {
            run(value).then(res => {
                //todo why null
                console.log(res)
            })
        } else {
            setOptions([])
        }
    };


    const dualSearchResult = (res: any) => {
        console.log(res)
        if (!res.ok || res.data == null || res.data.length === 0) {
            console.log('no data')
            return []
        }
        return res.data.map((item, idx) => {
            return <Option value={item.id} label={item.name} key={idx}>
                <div
                    style={{
                        display: 'flex',
                        justifyContent: 'space-between',
                    }}
                >
                        <span>
                            {item.name}
                        </span>
                    <span> </span>
                </div>
            </Option>
        });
    };


    const handleChange = (value: string, option: any) => {
        console.log('onSelect', value);
        console.log(option)
        history.push(`/t/${value}`)
    };
    return (

        <Layout>
            <HeaderWrapper>
                <Row justify="space-between">
                    <Col flex="127px">
                        <h1>
                            <Logo>
                                <img src={logo} alt="logo"/>
                            </Logo>

                        </h1>
                    </Col>
                    <Col flex="auto">
                        <Row justify="space-between" wrap={false}>
                            <Col flex="auto" xs={0} sm={12}>
                                <SearchOutlined style={{marginRight: ".25em"}}/>

                                <Select
                                    showSearch
                                    placeholder={"search topic here"}
                                    defaultActiveFirstOption={false}
                                    showArrow={false}
                                    filterOption={false}
                                    onSearch={handleSearch}
                                    onChange={handleChange}
                                    notFoundContent={"🔍"}
                                    style={{width: 200}}
                                >
                                    {options}
                                </Select>
                            </Col>
                            <Col flex="none" xs={0} sm={12}>
                                {/*<LocaleSwitch defaultLocale="zh-CN" onLocaleChange={(locale) => {*/}
                                {/*    globalContext.onLangChange(locale)*/}
                                {/*}}/>*/}


                                <Badge>
                                    <Profile/>
                                </Badge>


                            </Col>
                        </Row>


                    </Col>
                </Row>
            </HeaderWrapper>

            <Switch>
                {routes.map((route, i) => (
                    <RouteWithSubRoutes key={i} {...route} />
                ))}
                <Route path="/404" component={NotFoundPage}/>
                <Redirect to="/404"/>
            </Switch>

            <FooterWrapper>
                页脚
            </FooterWrapper>
        </Layout>

    )
}

export default React.memo(SiteLayout)
