import React, {useContext} from 'react'
import {Tabs} from 'antd';
import {StickyContainer,Sticky} from 'react-sticky';
import LinkList from "./LinkList";
import AddLinkItemV2 from "./AddLinkItemV2";
import {ReactComponent as AllIcon} from "../../../assets/icon/all.svg";
import {ReactComponent as GroupIcon} from "../../../assets/icon/group.svg";
import {ReactComponent as HotIcon} from "../../../assets/icon/hot.svg";
import {ReactComponent as MineIcon} from "../../../assets/icon/mine.svg";
import {ReactComponent as NewIcon} from "../../../assets/icon/new.svg";
import Icon from '@ant-design/icons';
import {useFormatMessage} from "react-intl-hooks";
import {GlobalContext} from "../../../App";

const {TabPane} = Tabs;

const renderTabBar = (props, DefaultTabBar) => (
    <Sticky  >
        {({ style }) => (
            <DefaultTabBar {...props} className="site-custom-tab-bar"
                           style={{...style, backgroundColor: "#fff", zIndex: 998}}/>
        )}
    </Sticky>
);

const OrderableLinkList: React.FC<any> = (props) => {
    let {topicId} = props
    const t = useFormatMessage()
    const globalContext = useContext(GlobalContext)
    const user = globalContext.user

    const AddLinkItemButton = <AddLinkItemV2/>
    return (
        <StickyContainer>
            <Tabs defaultActiveKey="1" tabBarExtraContent={user != null ? AddLinkItemButton : <></>}
                  renderTabBar={renderTabBar}>
                <TabPane
                    tab={<span><Icon component={AllIcon}/>{t({id: "topic.pane.all"})}</span>}
                    key="1"
                >
                    <LinkList topicId={topicId} filter="all"/>
                </TabPane>
                <TabPane
                    tab={<span><Icon component={HotIcon}/>最热</span>}
                    key="2"
                >
                    <LinkList topicId={topicId} filter="hot"/>
                </TabPane>
                <TabPane
                    tab={<span><Icon component={NewIcon}/>最新</span>}
                    key="3"
                >
                    <LinkList topicId={topicId} filter="newest"/>
                </TabPane>
                <TabPane
                    tab={<span><Icon component={MineIcon}/>我的</span>}
                    key="4"
                >
                    <LinkList topicId={topicId} filter="mine"/>
                </TabPane>
                <TabPane
                    tab={<span><Icon component={GroupIcon}/>标签</span>}
                    key="5"
                >
                    <LinkList topicId={topicId} filter="group"/>
                </TabPane>
            </Tabs>
        </StickyContainer>
    )
}


export default React.memo(OrderableLinkList)
