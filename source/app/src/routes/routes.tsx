import {createBrowserRouter, RouteObject} from "react-router-dom";
import React, {lazy, Suspense} from "react";

const lazyLoadLayout = (path: string) => {
    let url = `../layouts/${path}/index.tsx`
    const Comp = lazy(() => import(url))
    return (
        <Suspense>
            <Comp></Comp>
        </Suspense>
    )
}
const lazyLoadPage = (path: string) => {
    let url = `../pages/${path}/index.tsx`
    const Comp = lazy(() => import(url))
    return (
        <Suspense>
            <Comp></Comp>
        </Suspense>
    )
}
const routes: RouteObject[] = [
    {
        path: "/home",
        element: lazyLoadLayout("home"),
        children: [
            {
                path: "chat",
                element: lazyLoadPage("chat"),
            },
            {
                path: "contact",
                element: lazyLoadPage("contact")
            },
            {
                path: "discovery",
                element: lazyLoadPage("discovery")
            },
            {
                path: "me",
                element: lazyLoadPage("me"),
            }
        ]
    },
    {
        path: "setting",
        element: lazyLoadPage("setting"),
        children: [
            {
                path: "account_setting",
                element: lazyLoadPage("setting/account_setting")
            }
        ]
    },
    {
        path: "/",
        element: lazyLoadPage("cover"),
    },
    {
        path: "/board",
        element: lazyLoadPage("onboarding"),
    },
    {
        path: "/register/*",
        element: lazyLoadLayout("Register"),
    },
    {
        path: "/login",
        element: lazyLoadPage("login"),
    },
    {
        path: "/forgetPassword",
        element: lazyLoadPage("forgetPassword"),
    },
    {
        path: "/addFriend",
        element: lazyLoadPage("addFriend"),

    },
    {
        path: "/addSearch",
        element: lazyLoadPage("addFriend/addSearch")
    },
    {
        path: "/Profile",
        element: lazyLoadPage("Profile"),
    },
    {
        path: "/addDetail",
        element: lazyLoadPage("addFriend/addDetail"),
    },
    {
        path: "/chooseGroup",
        element: lazyLoadPage("addFriend/chooseGroup"),
    },
    {
        path: "/newFriends",
        element: lazyLoadPage("contact/newFriends"),
    },
    {
        path: "/dealFriendRequest",
        element: lazyLoadPage("contact/dealFriendRequest"),
    },
    {
        path: "/showFriendRequestDetail",
        element: lazyLoadPage("contact/showFriendRequestDetail")
    },
    {
        path: "/dialogue",
        element: lazyLoadPage("chat/dialogue"),
    },
    {
        path: "/musicPlayer",
        element: lazyLoadPage("musicPlayer"),
    }
]
const router = createBrowserRouter(routes);
export default router;