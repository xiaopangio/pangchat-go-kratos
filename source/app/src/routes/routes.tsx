import {createHashRouter, RouteObject} from "react-router-dom";
import {lazy, Suspense} from "react";

const lazyLoadLayout = (path: string) => {
    let url = `../layouts/${path}/index.tsx`
    const viteModule = import.meta.glob('../**/**/**/**');
    const Comp = lazy(viteModule[`${url}`] as any)
    return (
        <Suspense>
            <Comp></Comp>
        </Suspense>
    )
}
const lazyLoadPage = (path: string) => {
    let url = `../pages/${path}/index.tsx`
    const viteModule = import.meta.glob('../**/**/**/**');
    const Comp = lazy(viteModule[`${url}`] as any)
    return (
        <Suspense>
            <Comp></Comp>
        </Suspense>
    )
}
const routes: RouteObject[] = [
    {
        path: "/home",
        element: lazyLoadLayout("Home"),
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
        path: "/setting",
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
        path: "/groupManager",
        element: lazyLoadPage("contact/groupManager"),
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
    },
    {
        path: "/dataSet",
        element: lazyLoadPage("Profile/dataSet"),
    }

]
const Router = createHashRouter(routes);
export default Router;