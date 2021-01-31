import React, {PropsWithChildren, useContext, useState} from 'react'
import {Redirect, Route, Switch, useHistory} from "react-router-dom";
import NotFoundPage from "../components/404";
import {RouteWithSubRoutes} from '../pages/routes'
import {Col, Divider, Layout, Row, Select, Space} from "antd";
import Profile from "../components/Profile/Profile";
import styled from "styled-components";
import logo from '~/assets/logo.png'
import {GlobalContext} from "../App";
import {SearchTopic} from '../pages/topic/service'
import {useRequest} from 'ahooks'

const {Header, Footer, Content, Sider} = Layout
const {Option} = Select;


const HeaderWrapper = styled(Header)`
  background: #fff;
  height: 64px;
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
  color: rgba(0, 0, 0, .85);
  font-size: 18px;
  line-height: 64px;
  white-space: nowrap;
  text-decoration: none;

  & img {
    position: relative;
    top: -1.5px;
    height: 32px;
  }
`

const SelectOption = styled.div`
  display: flex;
  justify-content: space-between
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
            return [<Option value="" key="novalue">无结果</Option>]
        }
        return res.data.map((item, idx) => {
            return <Option value={item.id} label={item.name} key={idx}>
                <SelectOption>
                    <span>{item.name}</span>
                    <span>{item.shortCode}</span>
                </SelectOption>
            </Option>
        });
    };


    const handleChange = (value: string, option: any) => {
        if (value == "") {
            return
        }
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
                            <Logo onClick={() => {
                                history.push("/")
                            }}>
                                <img src={logo} alt="logo"/>
                            </Logo>

                        </h1>
                    </Col>
                    <Col flex="auto">
                        <Row justify="space-between" wrap={false}>
                            <Col flex="auto" xs={0} sm={12}>
                                <Select
                                    showSearch
                                    placeholder={"搜索主题"}
                                    defaultActiveFirstOption={false}
                                    showArrow={false}
                                    filterOption={false}
                                    onSearch={handleSearch}
                                    onChange={handleChange}
                                    notFoundContent={"🔍..."}
                                    style={{width: 400}}
                                >
                                    {options}
                                </Select>
                            </Col>
                            <Col flex="none" xs={0} sm={12}>
                                {/*<LocaleSwitch defaultLocale="zh-CN" onLocaleChange={(locale) => {*/}
                                {/*    globalContext.onLangChange(locale)*/}
                                {/*}}/>*/}
                                <Profile/>
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
                <Space split={<Divider type="vertical"/>}>
                    <a href="" target="_blank">FAQ</a>
                    <a href="" target="_blank">安装Chrome插件</a>
                    <a href="" target="_blank">联系站长</a>
                </Space>

            </FooterWrapper>
        </Layout>

    )
}

export default React.memo(SiteLayout)
