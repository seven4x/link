import React, {useCallback, useEffect, useState} from "react";
import styled from "styled-components";
import {ListLinks} from "../service";
import {List, Radio, Button, Skeleton} from "antd";
import {LinkItem} from './LinkItem'
import {LinkItemData} from '../model'
import {useKeyPress} from "ahooks";


const Container = styled.div`
`

interface LinkListProps {
    topicId: number
    filter?: string
}

const count = 3;

const LinkList: React.FC<LinkListProps> = ({topicId, filter}) => {
    //初次加载 table的菊花
    const [initLoading, setInitLoading] = useState<boolean>(true)
    //加载更多时不显示加载更多按钮
    const [loading, setLoading] = useState<boolean>(false)
    //本地数据缓存
    const [data, setData] = useState<Map<number, Array<LinkItemData>>>(new Map())
    const [list, setList] = useState<Array<LinkItemData>>([])
    const [total, setTotal] = useState(0)
    const [page, setPage] = useState(0)

    const getData = (page: number, callBack: any) => {
        if (data.has(page)) {
            console.log('get from local ')
            console.log(data.get(page))
            callBack(data.get(page))
            return
        }
        ListLinks(topicId, page, filter).then((res) => {
            callBack(res)
        });
    }
    //初始化
    useEffect(() => {
        getData(1, (res: any) => {
            console.log(res);
            setInitLoading(false)
            data.set(1, res)
            setList(res.data)
            setTotal(res.page && res.page.total)
        })
    }, [topicId]);

    const loadData = (page: number) => {
        //搞个假数据为了现实骨架
        let newList = [...new Array(count)].map(() => ({
            isLike: 0,
            like: 0,
            dislike: 0,
            loading: true,
            title: "",
            url: "",
            comment: 0,
            hotComment: {avatar: "", context: "", uid: 0}
        }))
        setList(newList);
        setLoading(true)
        setPage(page)
        getData(page, (res: any) => {
            let newData = res.data
            console.log(res);
            setLoading(false)
            data.set(page, res)
            setList(newData)
        })


    };

    return (
        <Container>
            <List
                loading={initLoading}
                itemLayout="horizontal"
                dataSource={list}
                pagination={{
                    hideOnSinglePage: true,
                    showSizeChanger: false,
                    size: "small",
                    onChange: page => {
                        loadData(page)
                    },
                    total: total
                }}
                renderItem={(item, idx) => (
                    <List.Item key={"link_" + idx}>
                        <Skeleton title={false} loading={item.loading} active>
                            <LinkItem link={item}/>
                        </Skeleton>
                    </List.Item>
                )}
            />
        </Container>
    );
};

export default React.memo(LinkList);
