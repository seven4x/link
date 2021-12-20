import React, {useContext, useEffect, useState} from "react";
import {useParams} from "react-router-dom";
import OrderableLinkList from "./components/OrderableLinkList";
import {Avatar, Col, Drawer, Layout, Row, Tag} from "antd";
import styled from "styled-components";
import {useFormatMessage} from "react-intl-hooks";
import Icon from '@ant-design/icons';
import {ReactComponent as TopicIcon} from "../../assets/icon/topic.svg";
import {ReactComponent as RelationIcon} from "../../assets/icon/relation.svg";
import AddTopic from "./components/AddTopic";
import {Topic} from './model'
import {AllPosition} from "~/utils/const";
import RelationTopic from './components/RelationTopic'
import {GetTopicDetail, ListMvpUser} from './service'
import {GlobalContext} from "../../App";
import TopicTree from "../../components/TopicTree/index";

const {Footer, Content, Sider} = Layout
const {CheckableTag} = Tag;


const ContentWrapper = styled(Content)`
    background-color: #fff;
    min-height: calc(100vh - 130px);
    padding-top: 38px;
    @media (min-width: 1024px) {
      padding-left: 177px;
    }
`

const FooterWrapper = styled(Footer)`
    @media (min-width: 1024px) {
      padding-left: 177px;
    }
`


function TopicHome() {
    let {topicId} = useParams<any>();
    let [isRealId, setIsRealId] = useState<boolean>(!isNaN(parseInt(topicId)))
    let [placement, setPlacement] = useState<any>("")
    let [visible, setVisible] = useState(false)
    let [selectedTag, setSelectedTag] = useState<string>("")
    const [topic, setTopic] = useState<Topic>({})
    const [mvps, setMvps] = useState([])
    const globalContext = useContext(GlobalContext)
    const user = globalContext.user

    const t = useFormatMessage();
    const allTags = AllPosition
    const relativeTopic = t({id: "topic.label.relative-topic"})

    useEffect(() => {
        GetTopicDetail(topicId).then(res => {
            setTopic(res.data || {})
            setIsRealId(true)
            document.title = (res.data?.name || "未知") +" 破茧♥"
            return res
        }).then(detail => {
            if (detail.data != null && detail.data.id != null) {
                ListMvpUser(detail.data.id).then(res => {
                    if (res.ok) {
                        setMvps(res.data)
                    }
                })
            }
        })
    }, [topicId])

    function handleChange(tag, checked) {

        setSelectedTag(tag);
        setVisible(true)
        setPlacement(tag)
    }

    return (

        <Layout>
            <Layout>
                <ContentWrapper>
                    <Row align="middle">
                        <Col flex="none">
                            <Icon component={TopicIcon}/> {t({id: "topic.label.current-topic"})}:&nbsp;&nbsp;
                        </Col>
                        <Col flex="auto">
                           <b> {topic.name}</b> &nbsp;&nbsp;  {topic.shortCode}
                        </Col>
                    </Row>

                    <Row align="middle" justify="space-between">
                        <Col flex="auto">
                            <Row>
                                <Col flex="none">
                                    <Icon component={RelationIcon}/> {relativeTopic}:&nbsp;&nbsp;
                                </Col>
                                <Col flex="auto">
                                    {allTags.map(tag => (
                                        <CheckableTag
                                            key={tag}
                                            checked={selectedTag.indexOf(tag) > -1}
                                            onChange={checked => handleChange(tag, checked)}
                                        >
                                            {t({id: `topic.radio.${tag}`})}
                                        </CheckableTag>
                                    ))}

                                </Col>
                            </Row>
                        </Col>
                        <Col flex="63px">
                            {user != null && <AddTopic topic={topic}/>}
                        </Col>
                    </Row>

                    {isRealId && <OrderableLinkList topicId={topic.id}/>}
                </ContentWrapper>

                <Sider breakpoint="lg"
                       collapsedWidth="0"
                       trigger={null}
                       style={{backgroundColor: "#fff"}}
                >
                    <TopicTree/>
                </Sider>


                <Drawer
                    title={<><Icon component={RelationIcon}/>{relativeTopic}</>}
                    placement={placement}
                    closable={true}
                    visible={visible}
                    key={placement}
                    onClose={() => {
                        setVisible(false)
                    }}
                >
                    {isRealId && <RelationTopic topicId={topic.id} position={placement} setVisible={setVisible}/>}
                </Drawer>
            </Layout>

            <FooterWrapper>
                <Row align="middle">
                    <Col flex="142px">
                        感谢:
                    </Col>
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
