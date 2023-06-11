import React from 'react'
import {createRoot} from 'react-dom/client'
import "@/index.less"
import 'virtual:svg-icons-register'
import "./service/service"
import {InitConnection} from "@/service/service";
import {RouterProvider} from "react-router-dom";
import Router from "@/routes/routes";
import {RecoilRoot} from "recoil";

InitConnection()
let rootDom = document.getElementById('root') as HTMLElement
let root = createRoot(rootDom)
root.render(
    <RecoilRoot>
        <RouterProvider router={Router}/>
    </RecoilRoot>
)