import React, {useState, useEffect} from 'react'
import {List, Card, Spin} from 'antd';
import {ListRelationTopic} from '../service'
import {useResponsive, useRequest} from 'ahooks';
import InfiniteScroll from 'react-infinite-scroller';
import {useHistory} from "react-router-dom";

export interface RelationTopicProps {
    topicId: number
    position: string
    setVisible: React.Dispatch<React.SetStateAction<boolean>>
}

function getColumn(position: string, responsive: any) {
    let column
    if (position === 'left' || position === 'right') {
        column = 1
    } else {
        if (responsive['xl']) {
            column = 12
        } else if (responsive['lg']) {
            column = 8
        } else if (responsive['md']) {
            column = 6
        } else if (responsive['sm']) {
            column = 4
        } else {
            column = 1
        }
    }
    return column
}

const RelationTopic: React.FC<RelationTopicProps> = (props) => {
    let {topicId, position, setVisible} = props;
    const [allData, setAllData] = useState([])
    const [prev, setPrev] = useState(0)
    const [hasMore, setHasMore] = useState(true)
    let history = useHistory();

    const {data, loading, run} = useRequest(ListRelationTopic, {
        manual: true,
        onSuccess: (result, params) => {
            if (result.ok) {
                setAllData(allData.concat(result.data))
                if (result.data.length > 0) {
                    setPrev(result.data[result.data.length - 1].id)
                }
                setHasMore(result.page && result.page.hasMore)
            }
        }
    })

    useEffect(() => {
        if (position === "") {
            return
        }
        run(topicId, position, prev).then(r => {
            console.log(r)
        })
    }, [topicId, position])

    const responsive = useResponsive()
    let column = getColumn(position, responsive)
    // console.log(column)

    function handleInfiniteOnLoad() {
        run(topicId, position, prev).then(r => {
            console.log(r)
        })
    }

    return (
        <>
            <InfiniteScroll
                initialLoad={false}
                pageStart={0}
                loadMore={handleInfiniteOnLoad}
                hasMore={!loading && hasMore}
                useWindow={false}
            >
                <List
                    grid={{column: column, gutter: 4}}
                    dataSource={allData}
                    renderItem={item => (
                        <List.Item>
                            <Card hoverable={true}
                                  onClick={() => {
                                      history.push(`/t/${item.id}`)
                                      console.log(item.name)
                                      setVisible(false)
                                  }}>{item.name}</Card>
                        </List.Item>
                    )}
                >
                    {loading && hasMore && (
                        <div className="demo-loading-container">
                            <Spin/>
                        </div>
                    )}
                </List>
            </InfiniteScroll>
        </>
    )
}
export default React.memo(RelationTopic)
