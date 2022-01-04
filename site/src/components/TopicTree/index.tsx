import React, {useState} from 'react'

import {Tree} from 'antd'
import styles from './topic_tree.module.css'
import {useRequest} from "ahooks";
import {ListRelationTopic} from "~/pages/topic/service";
import {useHistory} from "react-router-dom";

interface DataNode {
    title: string;
    key: string;
    isLeaf?: boolean;
    children?: DataNode[];
}

const initTreeDate: DataNode[] = [
    {title: '所有主题', key: '0'},
];

// It's just a simple demo. You can use tree map to optimize update perf.
function updateTreeData(list: DataNode[], key: React.Key, children: DataNode[]): DataNode[] {
    return list.map(node => {
        if (node.key === key) {
            return {
                ...node,
                children,
            };
        } else if (node.children) {
            return {
                ...node,
                children: updateTreeData(node.children, key, children),
            };
        }
        return node;
    });
}

const Index: React.FC = () => {
    let history = useHistory();

    const [treeData, setTreeData] = useState(initTreeDate);
    const {data, loading, run} = useRequest(ListRelationTopic, {
        manual: true,
        onSuccess: (result, params) => {
            if (result.ok) {
                let children = []
                result.data.forEach(t => {
                    children.push({
                        "title": t.name,
                        "key": t.id

                    })
                })
                setTreeData(origin =>
                    updateTreeData(origin, params[0], children),
                );

            }
        }
    })

    function onLoadData({key, children}: any) {
        return new Promise<void>(resolve => {
            if (children) {
                resolve();
                return;
            }
            run(key, 'bottom', 0).then(res => {

                resolve();
            })

        });
    }

    function onSelect(selectedKeys) {
        console.log(selectedKeys)
        let topicId = selectedKeys[0]
        history.replace({pathname: `/t/${topicId}`, state: null})
    }

    return (
        <div className={styles.topic_tree}>
            <Tree showLine={true} loadData={onLoadData} treeData={treeData} onSelect={onSelect}/>
        </div>
    )

}

export default React.memo(Index)
