import React, {useContext, useEffect, useState} from "react";
import {Link, useLocation} from 'react-router-dom';
import {Avatar} from "antd";
import {GlobalContext} from "../../App";
import {GetUserInfo} from '../../pages/account/service';

const Profile: React.FC<any> = (props) => {
    const globalContext = useContext(GlobalContext)
    let setLoginUser = globalContext.login
    let location = useLocation()
    console.log(location)
    let user = globalContext.user;
    let to = {
        pathname: "/login",
        state: {from: location.pathname}
    }
    useEffect(()=>{
        GetUserInfo().then(info => {
            console.log(info)
            if (info.ok) {
                setLoginUser(info.data)
            }
        }).catch(e=>{

        })
    },[])
    return (
        <>
            {
                user == null
                    ? <Link to={to}>登陆</Link>
                    : <Avatar src={globalContext.user.avatar}>  {globalContext.user.name} </Avatar>
            }
        </>

    )
}


export default React.memo(Profile)
