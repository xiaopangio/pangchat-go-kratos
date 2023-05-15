import {BottomNavigation, BottomNavigationAction} from "@mui/material";
import React, {useEffect, useState} from "react";
import SvgIcon from "@/components/Icon";
import "@/components/TabBar/index.less"
import {SetFunc} from "@/declare/type";
import {useLocation} from "react-router";
import {UnreadFriendRequestCount, UnreadFriendTextCount, UnreadMessageCount} from "@/store";
import {useRecoilState} from "recoil";

function TabBar({onchange, className}: { onchange: SetFunc, className?: string }) {
    let location = useLocation();
    const [unreadFriendRequestCount, setUnreadFriendRequestCount] = useRecoilState(UnreadFriendRequestCount)
    const [unreadMessageCount, setUnreadMessageCount] = useRecoilState(UnreadMessageCount)
    const [unreadFriendTextCount, setUnreadFriendTextCount] = useRecoilState(UnreadFriendTextCount)
    const name = location.pathname.split("/")[2]
    let [path, setPath] = useState("/" + name);
    //icon
    let [checked, setChecked] = useState("/chat");
    let btns = [
        {name: "/chat", icon: "chat", label: "Chat", color: ""},
        {name: "/contact", icon: "contact", label: "Contact", color: ""},
        {name: "/discovery", icon: "discovery", label: "Discovery", color: ""},
        {name: "/me", icon: "me", label: "me", color: ""},
    ]
    let [btnsState, setBtnsState] = useState(btns)
    let checkBtn = (btnName: string) => {
        if (btnName === checked) {
            return
        }
        for (let i = 0; i < btnsState.length; i++) {
            if (btnsState[i].name === btnName) {
                btnsState[i].color = "#4FA075"
            } else if (btnsState[i].name === checked) {
                btnsState[i].color = ""
            }
        }
        setBtnsState(btnsState)
        setChecked(btnName)
    }
    useEffect(() => {
        PubSub.unsubscribe("unreadFriendRequestsCount")
        PubSub.subscribe("unreadFriendRequestsCount", (message, data) => {
            setUnreadFriendRequestCount(data)
        })
    }, []);
    return (
        <div className={className}>
            <BottomNavigation value={path} onChange={(event, newValue) => {
                setPath(newValue)
                onchange(newValue);
                checkBtn(newValue)
            }} showLabels className={"bottomNavigation"}>
                {btns.map((btn, index) => {
                    return (
                        <BottomNavigationAction defaultChecked={btn.name === checked}
                                                key={btn.label}
                                                label={btn.label}
                                                value={btn.name}
                                                icon={<SvgIcon name={btn.icon} color={btn.color}
                                                />}
                        ></BottomNavigationAction>
                    )
                })}
            </BottomNavigation>
            <div className="count-list">
                <div className="count-item">
                    {unreadMessageCount === 0 ? "" : unreadMessageCount}
                </div>
                <div className="count-item">
                    {unreadFriendRequestCount === 0 ? "" : unreadFriendRequestCount}
                </div>
                <div className="count-item">
                    {unreadFriendTextCount === 0 ? "" : unreadFriendTextCount}
                </div>
            </div>
        </div>)
}

export default TabBar