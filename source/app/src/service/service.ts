import PubSub from "pubsub-js";
import {db, Friend, FriendRequest, Message, UnreadMessageInfo} from "@/store/db";
import storage from "@/utils/storage";
import {RefreshCurrentUser} from "@/utils/util";
import {DexieAddFriends, DexieStoreMessage, UpdateFriendRequestStatus} from "@/utils/store";
import {GetConnectionUrl} from "@/api/logic";
import {reconnect, status, websocketInit} from "@/hooks/websocket";
import {isNull} from "lodash";

export const Connect = async (f: Function) => {
    let user = RefreshCurrentUser();
    try {
        const url = await GetConnectionUrl()
        if (isNull(user)) {
            return
        }
        let cUrl = "ws://" + url.host + ":" + url.port + "/connect?uid=" + user.uid;
        f(cUrl)
    } catch (e) {
        console.log(e)
    }
}
export const InitConnection = () => {
    PubSub.unsubscribe("closeConn")
    PubSub.subscribe("closeConn", async (message, data) => {
        await Connect(function (cUrl: string) {
            setTimeout(() => {
                console.log(data)
                reconnect(cUrl)
            }, 10000)
        })
    })
    if (status === 'open') {
        return
    }
    Connect(function (cUrl: string) {
        websocketInit(cUrl)
    }).catch((e) => {
        console.log(e)
    })
}
// export const
export const ListenFriendRequest = () => {
    PubSub.unsubscribe("friendRequest")
    PubSub.subscribe("friendRequest", (e, data) => {
        console.log(data)
        let list = data as FriendRequest[];
        console.log("receive:", list)
        let user = RefreshCurrentUser();
        let requestsDealed: FriendRequest[] = []
        let requestsUnDealed: FriendRequest[] = []
        list.forEach((item) => {
            if (!user) {
                return
            }
            item.uid = user.uid
            if (item.status !== "0") {
                requestsDealed.push(item)
            } else {
                requestsUnDealed.push(item)
            }
        })
        let unreadFriendRequestsCount = 0;
        requestsDealed.forEach((item) => {
            UpdateFriendRequestStatus(item.request_id, item.status).then(() => {
                console.log("update friend request success")
            }).catch(() => {
                console.log("update friend request fail")
            })
        })
        db.friendRequests.bulkAdd(requestsUnDealed).then(() => {
            console.log("add friend request success")
            let Data = storage().get("unreadFriendRequestsCount")
            if (Data != null) {
                unreadFriendRequestsCount = Data.value as number + list.length
            }
            storage().set({key: "unreadFriendRequestsCount", value: unreadFriendRequestsCount + list.length})
        })
        PubSub.publish("unreadFriendRequestsCount", unreadFriendRequestsCount + list.length)
    })
}
export const ListenFriend = () => {
    PubSub.unsubscribe("friend")
    PubSub.subscribe("friend", (e, data) => {
        let list = data as Friend[]
        let user = RefreshCurrentUser()
        if (!user) {
            return
        }
        DexieAddFriends(user.uid, list).then(
            () => {
                console.log("add friends success")
            }
        )
    })
}
export const ListenMessage = () => {
    PubSub.unsubscribe("message")
    PubSub.subscribe("message", async (e, data: Message) => {
        await DexieStoreMessage(data)
        PubSub.publish("unreadMessage", data)
    })

}
export type UnreadMessageResponse = {
    list: UnreadMessageInfo[]
}
export const ListenMessageList = () => {
    PubSub.unsubscribe("unread_message_list")
    PubSub.subscribe("unread_message_list", async (e, data: UnreadMessageResponse) => {
        for (let unreadMessageInfo of data.list) {
            unreadMessageInfo.latest_message.is_need_load_before = true
            unreadMessageInfo.latest_message.need_load_before_count = unreadMessageInfo.unread_count - 1
            console.log(unreadMessageInfo.latest_message)
            await DexieStoreMessage(unreadMessageInfo.latest_message)
        }
        PubSub.publish("refreshMessageList", data)
    })
}