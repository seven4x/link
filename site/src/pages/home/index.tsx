import React from "react";

import HotTopic from "~/pages/home/components/HotTopic";
import styles from './index.module.css'
function SiteHome(){
    return (
        <div className={styles.container}>
            <div className={styles.content}>
                <HotTopic/>
            </div>
            <div className={styles.beian}>
                <a href="http://beian.miit.gov.cn/">浙ICP备2021004326号</a>
                <span className={styles.slogan}>给有价值的信息投票，获得有价值的信息</span>
            </div>
        </div>


    )
}


export default SiteHome;
