import React, {useState} from 'react'
import {Button, Divider, Form, Input, Modal, Select} from "antd";
import styles from './AddLinkItem.module.css'


const AddLinkItemV2: React.FC = () => {
    const [visible, setVisible] = useState<boolean>(false)


    const showModal = () => {
        setVisible(true);


    };

    const handleCancel = () => {
        console.log('Clicked cancel button');
        setVisible(false)
    };


    return (
        <>
            <Button onClick={showModal}>推荐链接 </Button>
            <Modal
                onOk={() => {
                    setVisible(false)
                }}

                visible={visible}
                onCancel={handleCancel}
            >
                <h1>添加“推荐到破茧”按钮到浏览器</h1>
                <h5>直接把下面的红色按钮拖拽到书签工具栏上，就可以快速添加链接到破茧 </h5>
                <h6>
                    如果你没有看到书签工具栏，点击浏览器右上方的菜单按钮，选择：书签 → 显示书签栏
                </h6>
                <div className={styles.share_button}>
                    <Button  type="primary" danger>
                        <a href={
                            `javascript: void (function (d, sc, e, w, h) {
    var r = "https://bitseatech.com/share?url=" + e(d.location.href) + "%26title=" + e(d.title),
      x = function () {
        if (
          !window.open(
            r,
            "bitsea",
            "toolbar=0,resizable=1,scrollbars=yes,status=1,width=" +
              w +
              ",height=" +
              h +
              ",left=" +
              (sc.width - w) / 2 +
              ",top=" +
              (sc.height - h) / 2
          )
        )
          location.href = r ;
      };
    setTimeout(x, 0);
  })(document, screen, encodeURIComponent, 720, 500);`
                        }
                           onClick={() => {
                               alert('请把这个按钮拖到你的浏览器书签栏');
                               return false;
                           }}>推荐到破茧</a>
                    </Button>
                </div>

            </Modal>
        </>
    )
}

export default React.memo(AddLinkItemV2)
