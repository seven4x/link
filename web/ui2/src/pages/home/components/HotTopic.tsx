import React, {useEffect, useState} from 'react'
import {Card, Space} from 'antd'
import {ListHotTopic} from "../service";
import {useHistory} from "react-router-dom";
import styles from './hotTopic.module.css'

const HostTopic: React.FC = () => {
    const [data, setData] = useState([])
    const history = useHistory()

    useEffect(() => {
        ListHotTopic().then(res => {
            console.log(res)
            if (res.ok) {
                setData(res.data)
            }
        })
    }, [])
    const toHot = function (item) {
        console.log(item)
        let id = item.shortCode == "" ? item.id : item.shortCode
        history.push(`/t/${id}`)
    }

    return (
        <Card title="热门主题" className={styles.title}>

            <Space className={styles.container}>
                {
                    data.map((item) =>
                        <Card key={"card" + item.id} className={styles.card}>
                            <span onClick={() => {
                                toHot(item)
                            }}>{item.name}</span>
                        </Card>
                    )
                }
            </Space>


        </Card>
    )
}


export default HostTopic
