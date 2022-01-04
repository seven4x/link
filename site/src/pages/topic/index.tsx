import React, {useEffect, useState} from "react";
import {useParams} from "react-router-dom";
import {Avatar, Col, Layout, Row} from "antd";
import styled from "styled-components";
import {useFormatMessage} from "react-intl-hooks";
import Icon from '@ant-design/icons';
import {ReactComponent as TopicIcon} from "../../assets/icon/topic.svg";
import {ReactComponent as RelationIcon} from "../../assets/icon/relation.svg";
import AddTopic from "./components/AddTopic";
import OrderableLinkList from "./components/OrderableLinkList";
import TopicTree from "../../components/TopicTree/index";
import {Topic} from './model';
import {GetTopicDetail, ListMvpUser, ListRelationTopic} from './service';

const {Footer, Content, Sider} = Layout


const ContentWrapper = styled(Content)`
    background-color: #fff;
    min-height: calc(100vh - 130px);
    padding-top: 38px;
`

const FooterWrapper = styled(Footer)`
    @media (min-width: 1024px) {
      padding-left: 177px;
    }
`


function TopicHome() {
    let {topicId} = useParams<any>();
    let [isRealId, setIsRealId] = useState<boolean>(!isNaN(parseInt(topicId)))

    const [topic, setTopic] = useState<Topic>({})
    const [mvps, setMvps] = useState([])
    const [relTopic, setRelTopic] = useState([])
    const [topicIntId, setTopicIntId] = useState(topicId)
    const t = useFormatMessage();
    const relativeTopic = t({id: "topic.label.relative-topic"})

    useEffect(() => {
        GetTopicDetail(topicId).then(res => {
            setTopic(res.data || {})
            setIsRealId(true)
            setTopicIntId(res.data.id )
            document.title = (res.data?.name || "未知") + " 破茧♥"
            return res
        }).then(detail => {
            if (detail.data != null && detail.data.id != null) {
                ListMvpUser(detail.data.id).then(res => {
                    if (res.ok) {
                        setMvps(res.data)
                    }
                })
                ListRelationTopic(detail.data.id, "all", 0).then(res => {
                    res.data.forEach(d => d.name = "➡️️" + d.name)
                    setRelTopic(res.data)
                })
            }
        })
    }, [topicId])


    return (

        <Layout>
            <Layout>
                <Sider breakpoint="lg"
                       collapsedWidth="0"
                       trigger={null}
                       style={{backgroundColor: "#fff"}}
                >
                    <TopicTree/>
                </Sider>

                <ContentWrapper>
                    <Row align="middle" justify="space-between">
                        <Col flex="auto">
                            <Row>
                                <Col flex="none">
                                    <Icon component={TopicIcon}/> {t({id: "topic.label.current-topic"})}:&nbsp;&nbsp;
                                </Col>
                                <Col flex="auto">
                                    <b> {topic.name}</b> &nbsp;&nbsp;  {topic.shortCode}
                                </Col>
                            </Row>

                        </Col>
                        <Col flex="63px">
                            <AddTopic topic={topic}/>
                        </Col>
                    </Row>

                    {isRealId && <OrderableLinkList topicId={topicIntId}/>}
                </ContentWrapper>


                <Sider breakpoint="lg"
                       collapsedWidth="0"
                       trigger={null}
                       style={{backgroundColor: "#fff"}}
                >
                    <Icon component={RelationIcon}/> {relativeTopic}:&nbsp;&nbsp;

                    {
                        relTopic.map(rel => <div>
                            <a href={`${rel.id}`}> {rel.name}</a>
                        </div>)
                    }
                </Sider>


            </Layout>

            <FooterWrapper>
                <Row align="middle">

                    <Col flex="auto">
                        <Avatar.Group
                            maxCount={10}
                            size="large"
                            maxStyle={{color: '#f56a00', backgroundColor: '#fde3cf'}}
                        >
                            {
                                mvps.map((mvp, idx) => {
                                    if (mvp.avatar == null) {
                                        return <Avatar key={idx}
                                                       style={{backgroundColor: '#f56a00'}}>{mvp.name.substr(0, 1)}</Avatar>
                                    }
                                    return <Avatar key={idx} src={mvp.avatar}/>

                                })
                            }
                        </Avatar.Group>
                    </Col>
                </Row>

            </FooterWrapper>


        </Layout>


    )
}


export default TopicHome;
