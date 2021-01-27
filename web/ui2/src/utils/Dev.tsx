import React, {useContext, useState} from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link,
    useRouteMatch
} from "react-router-dom";
import LocaleSwitch from "../components/LocaleSwitch";
import SideDrawer from "../components/SideDrawer/SideDrawer";
import LinkItem from '../pages/topic/components/LinkItem'
import {useFormatMessage, useFormatRelativeTime} from 'react-intl-hooks';
import {GlobalContext} from "../App";
import useUrlState from '@ahooksjs/use-url-state';
import {Pagination} from "antd";

function Dev() {
    let {path} = useRouteMatch();
    console.log('dev render')
    console.log(path)
    const t = useFormatMessage();
    const globalContext = useContext(GlobalContext)
    const tr = useFormatRelativeTime();
    const [lang, setLang] = useUrlState<any>({lang: "zh"})

    return (
        <div>
            <div style={{display: "flex", flexDirection: "column", minHeight: "100vh"}}>
                <div>
                    <Link to="/dev/local-switch">locale-switch</Link>|
                    <Link to="/dev/side-drawer">side-drawer</Link>|
                    <Link to="/dev/link-item"> link-item</Link>|
                    <Link to="/dev/intl">intl</Link>|

                </div>

                <div style={{border: "solid 1px", flex: "1"}}>

                    <Switch>
                        <Route path="/dev/local-switch">
                            <div>
                                <LocaleSwitch defaultLocale={"zh_CN"} onLocaleChange={(e) => {
                                    console.log(e)
                                }}/>
                            </div>
                        </Route>
                        <Route path="/dev/side-drawer">
                            <div style={{display: "flex", flexDirection: "column", minHeight: "100vh"}}>
                                <SideDrawer show={false} position="top" direction="row"
                                            children={[<h1>上方</h1>, <h1>🚔</h1>]}/>

                                <div style={{display: "flex", flexDirection: "row", flex: "1"}}>
                                    <SideDrawer show={false} position="left" direction="column"
                                                children={[<h2>左侧</h2>, <h2>⬅️</h2>]}/>
                                    <div style={{flex: "1"}}>
                                        {/*<LinkItem data={{*/}
                                        {/*    like: "0",*/}
                                        {/*    title: "如何评价 AMD 在北京时间10月29日凌晨发布的 RX6000 系列显卡?",*/}
                                        {/*    url: "https://developer.mozilla.org/zh-CN/docs/Web/CS",*/}
                                        {/*    hot: "6800砍掉了一整个shader engine，所以是96rops，60CU。这是AMD GPU第一次出现这种阉割方式。类似的做法的产品有1070,2070s等。\n" +*/}
                                        {/*        "\n",*/}
                                        {/*    up: 1024, disup: 20*/}
                                        {/*}}/>*/}
                                        {/*<LinkItem data={{*/}
                                        {/*    like: "1",*/}
                                        {/*    title: "在gin框架中使用JWT | 李文周的博客 #95",*/}
                                        {/*    url: "https://github.com/Q1mi/BlogComments/issues/95",*/}
                                        {/*    hot: "必看！",*/}
                                        {/*    up: 10240, disup: 1*/}


                                        {/*}}/>*/}
                                        {/*<LinkItem*/}
                                        {/*    data={{*/}
                                        {/*        like: "2",*/}
                                        {/*        title: "连链link link",*/}
                                        {/*        url: "https://www.yuque.com/seven4x/rm0od8",*/}
                                        {/*        hot: `虽然我日日夜夜黑奋斗逼，但是我实际上对奋斗逼是充满同情的，归根结底的原因是：对的不一定是重要的.*/}
                                        {/*有两点比较关键的知识希望大家知晓1：“经书”是什么？经就是经线的意思，它类似北斗星在中国传统中的位置，它位于中天，且在古人来看是恒久不变的，因此它比太阳还要重要，太阳是变星，要东升西落，因此反而落了下乘。！`,*/}
                                        {/*        up: 1024,*/}
                                        {/*        disup: 1,*/}
                                        {/*        comment: 20*/}

                                        {/*    }}/>*/}
                                    </div>
                                    <SideDrawer show={false} position="right" direction="column"
                                                children={[<h1>右</h1>, <h1>🈶️</h1>]}/>
                                </div>

                                <SideDrawer show={false} position="bottom" direction="row"
                                            children={[<h2>下</h2>, <h2>⬇️</h2>]}/>
                            </div>


                        </Route>

                        <Route path="/dev/link-item">
                            {/*<LinkItem data={{*/}
                            {/*    isLike: 0,*/}
                            {/*    title: "如何评价 AMD 在北京时间10月29日凌晨发布的 RX6000 系列显卡?",*/}
                            {/*    url: "https://developer.mozilla.org/zh-CN/docs/Web/CS",*/}
                            {/*    hot: "6800砍掉了一整个shader engine，所以是96rops，60CU。这是AMD GPU第一次出现这种阉割方式。类似的做法的产品有1070,2070s等。\n" +*/}
                            {/*        "\n",*/}
                            {/*    up: 1024, dislike: 20, comment: 0*/}
                            {/*}}/>*/}
                            {/*<LinkItem data={{*/}
                            {/*    like: "1",*/}
                            {/*    title: "在gin框架中使用JWT | 李文周的博客 #95",*/}
                            {/*    url: "https://github.com/Q1mi/BlogComments/issues/95",*/}
                            {/*    hot: "必看！",*/}
                            {/*    up: 10240, disup: 12*/}


                            {/*}}/>*/}
                            {/*<LinkItem*/}
                            {/*    data={{*/}
                            {/*        like: "2", title: "连链link link", url: "https://www.yuque.com/seven4x/rm0od8",*/}
                            {/*        hot: `虽然我日日夜夜黑奋斗逼，但是我实际上对奋斗逼是充满同情的，归根结底的原因是：对的不一定是重要的.*/}
                            {/*            有两点比较关键的知识希望大家知晓1：“经书”是什么？经就是经线的意思，它类似北斗星在中国传统中的位置，它位于中天，且在古人来看是恒久不变的，因此它比太阳还要重要，太阳是变星，要东升西落，因此反而落了下乘。！`,*/}
                            {/*        up: 1024, disup: 1,*/}
                            {/*        comment: 20*/}

                            {/*    }}/>*/}
                        </Route>

                        <Route path="/dev/intl">
                            <div>
                                <LocaleSwitch defaultLocale={"zh_CN"} onLocaleChange={(locale) => {
                                    globalContext.onLangChange(locale)
                                }}/>
                            </div>
                            <div>{t({id: "page.localeProvider.unreadCount"}, {unreadCount: 233})}</div>
                            <div>{tr(new Date().getDate() - 600)}</div>
                            <div>{lang.lang}</div>
                            <Pagination defaultCurrent={1} total={50} showSizeChanger/>
                        </Route>
                    </Switch>
                </div>

            </div>
        </div>
    );
}

export default Dev