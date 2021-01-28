import React, {useState, useEffect, useRef} from 'react'
import {List, Row, Col, Radio, Spin} from "antd";
import styled from "styled-components";
import {CommentData} from '../model'
import {useRequest} from 'ahooks'
import CommentItem from "./CommentItem";
import InfiniteScroll from 'react-infinite-scroller';
import {ListComment} from '../service'
import CommentEditor from './CommentEditor'


const CommentListWrapper = styled.div`
  border: 1px solid #e8e8e8;
  border-radius: 4px;
  overflow: auto;
  padding: 8px 24px;
  max-height: 98vh;
  width:56vw;
`




interface CommentListProps {
    linkId: number
}


const CommentList: React.FC<CommentListProps> = (props) => {
    let {linkId} = props
    const [allData, setAllData] = useState<Array<CommentData>>([])
    const [prev, setPrev] = useState(0)
    const [hasMore, setHasMore] = useState(true)
    const [sortBy, setSortBy] = useState("hot")
    const {data, loading, run} = useRequest(ListComment, {
        manual: true, debounceInterval: 500,
        onSuccess: (res) => {
            if (res.ok) {
                console.log(res.data)
                setAllData(allData.concat(res.data))
                if (res.data.length > 0) {
                    setPrev(res.data[res.data.length - 1].id)
                }
                setHasMore(res.page && res.page.hasMore)
            }
            return res
        }
    })
    const divEl = useRef(null);
    const [useWin, setUseWin] = useState(true)

    function handleInfiniteOnLoad() {
        if (divEl.current.scrollHeight > divEl.current.clientHeight) {
            setUseWin(false)
        }
        if (!hasMore) {
            console.log('no more,first load')
            return
        }
        console.info("useWin:" + useWin)
        run(linkId, prev, sortBy).then(r => {
            console.log(r)
        })
    }

    useEffect(() => {
        setAllData([])
        run(linkId, prev, sortBy).then(r => {
            console.log('init')
        })
    }, [linkId, sortBy])

    const header = <Row justify="space-between">
        <Col push={1}>
            {`${allData.length} replies`}
        </Col>
        <Col pull={1}>
            <Radio.Group defaultValue={sortBy} size="small" buttonStyle="solid">
                <Radio.Button value="hot" onClick={() => {
                    setSortBy("hot")
                }}>hot</Radio.Button>
                <Radio.Button value="newest" onClick={() => {
                    setSortBy("newest")
                }}>newest</Radio.Button>
            </Radio.Group>
        </Col>

    </Row>

    function afterAdd(comment: any) {
        console.log(comment)
        setAllData((pre) => {
            return [comment, ...pre]
        })
    }

    function afterDelete(mId: number) {
        setAllData((pre) => {
            return pre.filter(m => m.id !== mId)
        })
    }

    return (
        <CommentListWrapper ref={divEl}>

            <CommentEditor linkId={linkId} afterAdd={afterAdd}/>

            <InfiniteScroll
                pageStart={0}
                loadMore={handleInfiniteOnLoad}
                hasMore={!loading && hasMore}
                useWindow={useWin}
                initialLoad={false}
                threshold={600}
                style={{marginBottom: "70px"}}
            >
                <List
                    className="comment-list"
                    header={header}
                    itemLayout="horizontal"
                    dataSource={allData}
                    renderItem={item => (
                        <li>
                            <CommentItem data={item} linkId={linkId} afterDelete={afterDelete}/>
                        </li>
                    )}
                >
                    {loading && hasMore && (
                        <div className="demo-loading-container">
                            <Spin/>
                        </div>
                    )}
                    {
                        !hasMore && (<div style={{textAlign: "center"}}> 没有更多了 </div>)
                    }
                </List>
            </InfiniteScroll>


        </CommentListWrapper>
    )
}


export default React.memo(CommentList)
