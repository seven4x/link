import React, {useContext, useEffect, useState} from "react";
import {Link, useHistory, useLocation} from 'react-router-dom';
import {Avatar, Badge} from "antd";
import {GlobalContext} from "../../App";
import {GetUserInfo, Logout} from '~/pages/account/service';

const Profile: React.FC<any> = (props) => {
    const globalContext = useContext(GlobalContext)
    let setLoginUser = globalContext.login
    let location = useLocation()
    let history = useHistory();
    const [showPop, setShowPop] = useState<boolean>(false)
    console.log(location)
    let user = globalContext.user;
    let to = {
        pathname: "/login",
        state: {from: location.pathname}
    }
    useEffect(() => {
        GetUserInfo().then(info => {
            console.log(info)
            if (info.ok) {
                setLoginUser(info.data)
            }
        }).catch(e => {

        })
    }, [])

    const LoginOut = () => {
        Logout().then(res => {
            console.log(res)
        }).catch(e => {

        }).finally(() => {
            setLoginUser(null)
            history.replace("/")
        })
    }
    return (
        <>
            <Badge>
                {
                    user == null
                        ? <Link to={to}>登陆</Link>
                        : <div onClick={() => {
                            setShowPop(p => {
                                return !p;
                            })
                        }}><Avatar src={globalContext.user.avatar}>  {globalContext.user.name} </Avatar></div>
                }

            </Badge>

            {showPop && <div onClick={LoginOut}>退出</div>}
        </>


    )
}


export default React.memo(Profile)
