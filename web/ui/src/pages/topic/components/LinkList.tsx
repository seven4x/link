import React, { useContext, useEffect, useState} from "react";
import {ListLinks} from "../service";
import {List,   Skeleton, message} from "antd";
import {LinkItem} from './LinkItem'
import {LinkItemData} from '../model'
import {GlobalContext} from "../../../App";
import useUrlState from '@ahooksjs/use-url-state';


interface LinkListProps {
    topicId: number
    filter?: string
}

const count = 3;

const LinkList: React.FC<LinkListProps> = ({topicId, filter}) => {
    const loginContext = useContext(GlobalContext)
    let loginUser = loginContext.user
    //初次加载 table的菊花
    const [initLoading, setInitLoading] = useState<boolean>(true)
    //加载更多时不显示加载更多按钮
    const [loading, setLoading] = useState<boolean>(false)
    const [list, setList] = useState<Array<LinkItemData>>([])
    const [total, setTotal] = useState(0)
    const [page, setPage] = useUrlState({p: 1})

    //初始化
    useEffect(() => {
        loadData()
    }, [topicId, page?.p]);

    const loadData = () => {
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
        setList(newList);
        setLoading(true)
        ListLinks(topicId,Number(page?.p), filter).then((res) => {
            let newData = res.data
            console.log(res);
            setLoading(false)
            setInitLoading(false)
            setList(newData)
            setTotal(res.page && res.page.total)
        });

    };

    return (
        <div>
            <List
                loading={initLoading}
                itemLayout="horizontal"
                dataSource={list}
                pagination={{
                    hideOnSinglePage: true,
                    showSizeChanger: false,
                    size: "small",
                    current: Number(page?.p),
                    onChange: page => {
                        if (page > 3 && loginUser == null) {
                            message.info('请登录,查看更多')
                            return
                        }
                        setPage((s) => ({p: page}))
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
        </div>
    );
};

export default React.memo(LinkList);
