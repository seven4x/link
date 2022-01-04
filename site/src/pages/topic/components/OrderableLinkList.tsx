import React, {useContext} from 'react'
import {Tabs} from 'antd';
import {Sticky, StickyContainer} from 'react-sticky';
import LinkList from "./LinkList";
import AddLinkItemV2 from "./AddLinkItemV2";

import {ReactComponent as NewIcon} from "../../../assets/icon/new.svg";
import Icon from '@ant-design/icons';
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

const OrderableLinkList: React.FC<any> = ({topicId}) => {
    const globalContext = useContext(GlobalContext)
    const user = globalContext.user

    const AddLinkItemButton = <AddLinkItemV2/>
    return (
        <StickyContainer>
            <Tabs defaultActiveKey="1" tabBarExtraContent={user != null ? AddLinkItemButton : <></>}
                  renderTabBar={renderTabBar}>

                <TabPane
                    tab={<span><Icon component={NewIcon}/>最新</span>}
                    key="3"
                >
                    <LinkList topicId={topicId} filter="newest"/>
                </TabPane>

            </Tabs>
        </StickyContainer>
    )
}


export default OrderableLinkList
