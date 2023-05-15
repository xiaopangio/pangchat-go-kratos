import "@/pages/setting/index.less"
import SettingHeader from "@/components/SettingHeader";
import {Outlet, useLocation, useNavigate} from "react-router";
import {LoginPrefix, MePrefix, SettingPrefix} from "@/declare/const";
import SettingList from "@/components/SettingContent";
import {settingGroupList} from "@/store/data";
import React, {useState} from "react";
import {useRecoilState, useRecoilValue} from "recoil";
import {ConnectionState, currentUserState} from "@/store";
import {Logout} from "@/api/user";
import storage from "@/utils/storage";
import {closeConn} from "@/hooks/websocket";


export default function Setting() {
    let navigate = useNavigate();
    let location = useLocation();
    const [maskClass, setMaskClass] = useState("mask");
    const [popperClass, setPopperClass] = useState("popper");
    const [connection, setConnection] = useRecoilState(ConnectionState)
    const [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    let user = useRecoilValue(currentUserState);
    const back = () => {
        navigate(MePrefix)
    }
    const closeMask = () => {
        setMaskClass("mask mask-hide")
        setPopperClass("popper popper-hide")
    }
    const handleExit = async () => {
        if (!user) {
            console.log("未登录")
            return
        }
        try {
            await Logout({uid: user.uid})
            storage().delete("token")
            storage().delete("user_avatar")
            setMaskClass("mask mask-hide")
            setPopperClass("popper popper-hide")
            setConnection(false)
            setCurrentUser(null)
            closeConn()
            navigate(LoginPrefix)
        } catch (e) {
            console.log(e)
        }
    }
    const click = (title: string) => {
        switch (title) {
            case "账号与安全":
                navigate(SettingPrefix + "/account_setting")
                break
            case "退出":
                setMaskClass("mask mask-active")
                setPopperClass("popper popper-active")
                break
        }
    }
    return (
        <div>
            {
                location.pathname === SettingPrefix ?
                    <div className="setting">
                        <SettingHeader title="设置" back={back}/>
                        <SettingList click={click} settingGroupList={settingGroupList}/>
                    </div> : <Outlet/>
            }
            {
                <div className="exit-popper">
                    <div className={maskClass} onClick={closeMask}></div>
                    <div className={popperClass}>
                        <div className="popper-item" onClick={handleExit}>退出登录</div>
                        <div className="popper-item">关闭微信</div>
                        <div className="popper-item">取消</div>
                    </div>
                </div>
            }
        </div>

    );
}