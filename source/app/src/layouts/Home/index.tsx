import React from "react";
import "@/layouts/Home/index.less"
import {Outlet, useNavigate} from "react-router";
import {HomePrefix} from "@/declare/const";
import Box from "@mui/material/Box";
import TabBar from "@/components/TabBar";

function Home() {
    const navigate = useNavigate();
    const changePath = (path: string) => {
        navigate(HomePrefix + path)
    }
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