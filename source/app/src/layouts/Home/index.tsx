import React, {useEffect} from "react";
import "@/layouts/Home/index.less"
import {Outlet, useNavigate} from "react-router";
import {HomePrefix} from "@/declare/const";
import {useRecoilState} from "recoil";
import {ConnectionUrlState, currentUserState} from "@/store";
import {GetConnectionUrl} from "@/api/logic";
import {checkToken} from "@/utils/check";
import {User} from "@/declare/type";
import {ParseJwt} from "@/utils/jwt";
import {reconnect, status, websocketInit} from "@/hooks/websocket";
import Box from "@mui/material/Box";
import TabBar from "@/components/TabBar";

function Home() {
    const navigate = useNavigate();
    const [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    const [connectionUrl, setConnectionUrl] = useRecoilState(ConnectionUrlState)
    const changePath = (path: string) => {
        navigate(HomePrefix + path)
    }

    useEffect(() => {
        let token = checkToken();
        let user: User | null = null;
        if (token) {
            user = ParseJwt<User>(token.value);
        }
        if (user) {
            setCurrentUser(user);
        }
        if (connectionUrl && connectionUrl !== "") {
            return
        }
        PubSub.unsubscribe("closeConn")
        PubSub.subscribe("closeConn", (message, data) => {
            if (!currentUser) {
                return
            }
            GetConnectionUrl().then((res => {
                let url = "ws://" + res.host + ":" + res.port + "/connect?uid=" + currentUser.uid;
                setConnectionUrl(url)
                setTimeout(() => {
                    reconnect(url)
                }, 5000)
            }), (err) => {
                console.log(err)
            })
        })
        GetConnectionUrl().then((res => {
            if (!user) {
                return
            }
            let url = "ws://" + res.host + ":" + res.port + "/connect?uid=" + user.uid;
            setConnectionUrl(url)
        }), (err) => {
            console.log(err)
        })
    }, []);


    useEffect(() => {
        if (connectionUrl === "") {
            return
        }
        if (status === 'open') {
            return;
        }
        websocketInit(connectionUrl)
    }, [connectionUrl]);
    return (
        <div className={"home"}>
            <Box className={"content"}>
                <Outlet/>
            </Box>
            <TabBar onchange={changePath} className={"tab-bar"}></TabBar>
        </div>
    )
}

export default Home