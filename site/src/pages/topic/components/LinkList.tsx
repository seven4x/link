import React, {useEffect, useState} from "react";
import {ListLinks} from "../service";
import {Button, List, Skeleton} from "antd";
import {LinkItem} from './LinkItem'
import {LinkItemData} from '../model'


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
    const [more, setMore] = useState<boolean>(true)
    const [list, setList] = useState<Array<LinkItemData>>([])
    const [data, setData] = useState<Array<LinkItemData>>([])
    const [prev, setPrev] = useState(0)
    //初始化
    useEffect(() => {
        console.log("LinkList useEffect")
        setData([])
        setList([])
        setPrev(0)
        loadData()
        return () => {
            console.log("LinkList cleanup")
            setData([])
            setList([])
            setPrev(0)
        }
    }, [topicId]);


    const loadData = async () => {
        //搞个假数据为了现实骨架
        let newList = [...new Array(count)].map(() => ({
            isLike: 0,
            agree: 0,
            disagree: 0,
            loading: true,
            title: "",
            link: "",
            commentCount: 0,
            hotComment: {avatar: "", content: "", uid: 0}
        }))
        setList(old => {
            return old.concat(newList)
        });
        setLoading(true)
        const res = await ListLinks(topicId, prev, filter)

        console.log(res)
        setLoading(false)
        setInitLoading(false)
        setData(old => {
            return data.concat(res.data)
        })
        setList(old => {
            return data.concat(res.data)
        })
        if (!res?.page?.hasMore) {
            setMore(false)
        } else {
            setMore(true)
        }
        setPrev(res?.page && res?.page.nextId)


    };

    const loadMore =
        !initLoading && !loading && more ? (
            <div
                style={{
                    textAlign: 'center',
                    marginTop: 12,
                    height: 32,
                    lineHeight: '32px',
                }}
            >
                <Button onClick={loadData}>查看更多</Button>
            </div>
        ) : null;

    return (
        <div>
            <List
                loading={initLoading}
                itemLayout="horizontal"
                dataSource={list}
                loadMore={loadMore}
                renderItem={(item, idx) => (
                    <List.Item key={"link_" + idx}>
                        <Skeleton title={false} loading={item.loading} active>
                            <LinkItem link={item}/>
                        </Skeleton>
                    </List.Item>
                )}
            />
        </div>
    );
};

export default LinkList;
