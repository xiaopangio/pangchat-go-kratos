import PubSub from "pubsub-js";
import {db, FriendRequest} from "@/store/db";
import storage from "@/utils/storage";
import {RefreshCurrentUser} from "@/utils/util";

export const ListenFriendRequest = () => {
    PubSub.unsubscribe("friendRequest")
    PubSub.subscribe("friendRequest", (e, data) => {
        let list = data as FriendRequest[];
        console.log("receive:", list)
        let user = RefreshCurrentUser();
        list.forEach((item) => {
            if (!user) {
                return
            }
            item.uid = user.uid
        })
        let unreadFriendRequestsCount = 0;
        db.friendRequests.bulkAdd(list).then(() => {
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
