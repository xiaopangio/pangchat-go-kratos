import React from 'react'
import ReactDOM from 'react-dom/client'
import "@/index.less"
import {RouterProvider} from 'react-router-dom'
import router from '@/routes/routes'
import 'virtual:svg-icons-register'
import {RecoilRoot} from "recoil";
import "./service/service"

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <RecoilRoot>
        <RouterProvider router={router}/>
    </RecoilRoot>
)

// PubSub.subscribe("test", (msg, data) => {
//     console.log(msg, data)
// })