import React from "react";
import {Button, Layout, List} from "antd";
import styled from "styled-components";
import styles from './index.module.css'

const {Content} = Layout;

const count = 5;
const fakeDataUrl = `/api1/link/recent`;
const Link = styled.a`
&:visited{
  color: #7a7a7a;
}
&:link{
}
`

class LoadMoreList extends React.Component {
    state = {
        initLoading: true,
        loading: false,
        data: [],
        list: [],
        prevId: 0,
        noMore: false
    };

    componentDidMount() {
        fetch(fakeDataUrl)
            .then(res => res.json())
            .then(res => {
                res.Data.map(d => d.loading = false)
                console.log(res)
                this.setState({
                    initLoading: false,
                    data: res.Data,
                    list: res.Data,
                    prevId: res.NextId
                });
            });
    }

    onLoadMore = () => {
        this.setState({
            loading: true,
            list: this.state.data.concat(
                [...new Array(count)].map(() => ({loading: true, Title: "", Link: ""})),
            ),
        });
        fetch("/api1/link/recent?prev=" + this.state.prevId)
            .then(res => res.json())
            .then(res => {
                res.Data.map(d => d.loading = false)
                const data = this.state.data.concat(res.Data);
                console.log(res)
                this.setState(
                    {
                        data,
                        list: data,
                        loading: false,
                        prevId: res.NextId,
                        noMore: res.NextId == 0
                    },
                    () => {
                        // Resetting window's offsetTop so as to display react-virtualized demo underfloor.
                        // In real scene, you can using public method of react-virtualized:
                        // https://stackoverflow.com/questions/46700726/how-to-use-public-method-updateposition-of-react-virtualized
                        window.dispatchEvent(new Event('resize'));
                    },
                );
            });
    };

    render() {
        const {initLoading, loading, list, noMore} = this.state;
        const loadMore =
            !initLoading && !loading && !noMore ? (
                <div
                    style={{
                        textAlign: 'center',
                        marginTop: 12,
                        height: 32,
                        lineHeight: '32px',
                    }}
                >
                    <Button onClick={this.onLoadMore}>查看更多 get more 🔥</Button>
                </div>
            ) : null;

        return (
            <Layout>
                <Content className={styles.content}>
                    <List
                        className="demo-loadmore-list"
                        loading={initLoading}
                        itemLayout="horizontal"
                        loadMore={loadMore}
                        dataSource={list}
                        renderItem={item => (
                            <List.Item>
                                <List.Item.Meta
                                    title={<Link href={item.Url} target="_blank"
                                                 rel="noreferrer noopener">{item.Title}</Link>}
                                    description={item.Tags}
                                />
                                <div>{item.CreateTime}</div>
                            </List.Item>
                        )}
                    />
                </Content>

            </Layout>

        );
    }
}

export default LoadMoreList;