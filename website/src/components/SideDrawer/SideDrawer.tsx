import React, {useState} from 'react'
import styled from "styled-components";
import {ReactComponent as UpIcon} from "./up.svg";
import {ReactComponent as DownIcon} from "./down.svg";

export interface SideDrawerProps {
    show?: boolean
    children: React.ReactNode[]
    direction?: string
    position?: string
    style?: object
}

const ChildrenWrapper = styled.div`
  margin: 5px;
`
const Wrapper = styled.div<any>`
  display: flex;
  flex-direction: ${prop => prop.direction === "row" ? "row" : "column"};
  background-color: #61dafb;
  height: 100%;
`
//top: 50%; right: 100%; 右侧
//top:50%;right:-16px;左侧
//left:50%;top:100%;上方
//left:50%;top:-16px;底部

const handlePositionTopValue = (position: string) => {
    switch (position) {
        case "left":
        case "right":
            return "80%";
        case "top":
            return "100%";
        case "bottom":
            return "-16px";
        default :
            return ""
    }
};

const handlePositionRightValue = (position: string) => {
    switch (position) {
        case "left":
            return "-16px";
        case "right":
            return "100%";
        default :
            return ""
    }
};

const handlePositionLeftValue = (position: string) => {
    console.log(position)
    switch (position) {
        case "top":
        case "bottom":
            return "80%";
        default :
            return ""
    }
};


const handlePositionRotateValue = (position: string) => {
    console.log(position)
    switch (position) {
        case "left":
            return "rotate(90deg)";
        case "right":
            return "rotate(-90deg)";
        case "top":
            return "rotate(180deg)";
        case "bottom":
            return "rotate(0deg)";
        default :
            return ""
    }
};

//https://stackoverflow.com/questions/56047659/multiple-props-options-for-styled-components
const Handler = styled.div<any>`
position:absolute;
top: ${props => handlePositionTopValue(props.position)};
right: ${props => handlePositionRightValue(props.position)};
left: ${props => handlePositionLeftValue(props.position)};
transform:${props => handlePositionRotateValue(props.position)};
`
//抽屉容器
const SideDrawer: React.FC<SideDrawerProps> = (props) => {
    let [show, setShow] = useState(props.show)

    return (
        <div style={{position: "relative", ...props.style}}>
            {show && <Wrapper direction={props.direction}>
                {
                    props.children.map((node, index) => {
                        return <ChildrenWrapper key={index}>
                            {node}
                        </ChildrenWrapper>
                    })

                }
            </Wrapper>
            }
            <Handler position={props.position} onClick={() => {
                setShow(!show)
            }}>
                {show ? <DownIcon width="16px" height="16px"/> : <UpIcon width="16px" height="16px"/>}
            </Handler>

        </div>
    )
}

export default React.memo(SideDrawer)