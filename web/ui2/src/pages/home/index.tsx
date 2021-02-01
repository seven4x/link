import React from "react";

import HotTopic from "~/pages/home/components/HotTopic";
import styles from './index.module.css'
function SiteHome(){
    return (
        <div className={styles.container}>
            <HotTopic/>
        </div>
    )
}


export default SiteHome;
