import "@/pages/chat/index.less"
import HomeHeader from "@/components/HomeHeader";
import {useRecoilState} from "recoil";
import {AvatarMap, currentDialogState, currentUserState, FriendsMap, UnreadMessageCountMap} from "@/store";
import {useEffect, useState} from "react";
import {
    DexieGetFriend,
    DexieGetImgList,
    DexieGetMessagesByMessageId,
    DexieGetUnreadMessagesCounts,
    DexieStoreUnreadMessagesCounts
} from "@/utils/store";
import {isNull, isUndefined} from "lodash";
import {RefreshCurrentUser} from "@/utils/util";
import {Friend, Message, MessageCountType} from "@/store/db";
import DefaultImg from "@img/default.jpg";
import {useNavigate} from "react-router";
import PubSub from "pubsub-js";
import {UnreadMessageResponse} from "@/service/service";

function Chat() {
    const [unreadMessageCountMap, setUnreadMessageCountMap] = useRecoilState(UnreadMessageCountMap)
    const [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    const [avatarMap, setAvatarMap] = useRecoilState(AvatarMap)
    const [friendsMap, setFriendsMap] = useRecoilState(FriendsMap)
    const [messageList, setMessageList] = useState<Message[]>([]);
    const [, setDialogState] = useRecoilState(currentDialogState)

    let navigate = useNavigate();
    // 初始化消息列表
    const initUnreadMessageCountMap = async () => {

        if (unreadMessageCountMap.size > 0) {
            return
        }
        try {
            let user = RefreshCurrentUser()
            if (isNull(user)) {
                return
            }
            let res = await DexieGetUnreadMessagesCounts(user.uid)
            if (res.size === 0) {
                return
            }
            setUnreadMessageCountMap(res)
        } catch (e) {
        }
    }
    // 初始化朋友信息，为了渲染消息列表时，能够渲染出朋友的信息
    const initFriend = async (uid: string, friend_id: string) => {
        let friend = friendsMap.get(friend_id);
        if (isUndefined(friend)) {
            try {
                friend = await DexieGetFriend(uid, friend_id)
                friendsMap.set(friend_id, friend)
            } catch (e) {
                console.log(e)
            }
        }
    }
    // 初始化消息列表中展示的消息
    const initMessageList = async () => {
        if (isNull(currentUser)) {
            return
        }
        // 消息列表要展示的消息id数组
        let messageIds: string[] = []
        unreadMessageCountMap.forEach((value) => {
            messageIds.push(value.message_id)
        })
        // 获取消息列表要展示的消息
        let MessageList = await DexieGetMessagesByMessageId(messageIds)
        for (let i = 0; i < MessageList.length; i++) {
            MessageList[i].friend_id = MessageList[i].sender_id === currentUser.uid ? MessageList[i].receiver_id : MessageList[i].sender_id
        }
        setMessageList(MessageList)
        // 获取消息列表要展示的好友
        let friendIs: string[] = []
        for (const value of MessageList) {
            if (value.sender_id === currentUser.uid) {
                // 如果是自己发的消息，那么好友就是接收者
                await initFriend(currentUser.uid, value.receiver_id)
                friendIs.push(value.receiver_id)
            } else {
                // 如果是别人发的消息，那么好友就是发送者
                await initFriend(currentUser.uid, value.sender_id)
                friendIs.push(value.sender_id)
            }
        }
        let tempFriendsMap = new Map<string, Friend>(friendsMap)
        setFriendsMap(tempFriendsMap)
        let avatarUrls: string[] = []
        for (const value of friendIs) {
            let friend = friendsMap.get(value)
            if (isUndefined(friend)) {
                continue
            }
            avatarUrls.push(friend.avatar)
        }
        let imgs = await DexieGetImgList(avatarUrls)
        for (const value of imgs) {
            avatarMap.set(value.name, value.blob)
        }
        setAvatarMap(new Map<string, string>(avatarMap))

    }
    // 获取朋友的名称信息
    const getName = (friend_id: string) => {
        let friend = friendsMap.get(friend_id);
        if (isUndefined(friend)) {
            return ""
        }
        return friend.nick_name === friend.note_name ? friend.nick_name : friend.note_name
    }
    // 获取朋友的头像信息
    const getAvatar = (friend_id: string) => {
        let friend = friendsMap.get(friend_id);
        if (isUndefined(friend)) {
            return DefaultImg
        }
        let avatar = avatarMap.get(friend.avatar);
        if (isUndefined(avatar)) {
            return DefaultImg
        }
        return avatar
    }
    // 获取消息的时间信息
    const getTime = (time: string) => {
        let date = new Date(time)
        let now = new Date()
        // 日期在今天的话，只显示时间
        if (date.getFullYear() === now.getFullYear() && date.getMonth() === now.getMonth() && date.getDate() === now.getDate()) {
            let hours = date.getHours();
            let minutes = date.getMinutes();
            if (hours < 10) {
                if (minutes < 10) {
                    return `0${hours}:0${minutes}`
                }
                return `0${hours}:${minutes}`
            }
            if (minutes < 10) {
                return `${hours}:0${minutes}`
            }
            return `${hours}:${minutes}`
        }
        // 日期在昨天的话，显示昨天
        if (date.getFullYear() === now.getFullYear() && date.getMonth() === now.getMonth() && date.getDate() === now.getDate() - 1) {
            return "昨天"
        }
        // 日期在前天的话，显示前天
        if (date.getFullYear() === now.getFullYear() && date.getMonth() === now.getMonth() && date.getDate() === now.getDate() - 2) {
            return "前天"
        }
        // 日期在其他的话，显示日期
        return `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`
    }
    // 获取朋友的未读消息数
    const getUnreadMessageCount = (friendId: string) => {
        let unreadMessageCount = unreadMessageCountMap.get(friendId)
        if (isUndefined(unreadMessageCount)) {
            return 0
        }
        return unreadMessageCount.unread_count
    }
    // 监听未读消息
    const ListenUnreadMessage = () => {
        // 负责监听在线期间的未读消息
        PubSub.unsubscribe("unreadMessage")
        PubSub.subscribe("unreadMessage", (msg: string, data: Message) => {
            // 拷贝一份未读消息map
            let tempMap = new Map<string, MessageCountType>(unreadMessageCountMap)
            let f = tempMap.get(data.sender_id);
            if (!isUndefined(f)) {
                tempMap.set(data.sender_id, {
                    friend_id: data.sender_id,
                    message_id: data.message_id,
                    uid: data.receiver_id,
                    unread_count: f.unread_count + 1
                })
            } else {
                tempMap.set(data.sender_id, {
                    friend_id: data.sender_id,
                    message_id: data.message_id,
                    uid: data.receiver_id,
                    unread_count: 1
                })
            }
            setUnreadMessageCountMap(tempMap)
        })
        // 负责监听离线期间的未读消息
        PubSub.unsubscribe("refreshMessageList")
        PubSub.subscribe("refreshMessageList", async (msg: string, data: UnreadMessageResponse) => {
            let tempMap = new Map<string, MessageCountType>(unreadMessageCountMap)
            for (let unreadMessageInfo of data.list) {
                let f = tempMap.get(unreadMessageInfo.latest_message.sender_id);
                if (isUndefined(f)) {
                    tempMap.set(unreadMessageInfo.latest_message.sender_id, {
                        friend_id: unreadMessageInfo.latest_message.sender_id,
                        message_id: unreadMessageInfo.latest_message.message_id,
                        uid: unreadMessageInfo.latest_message.receiver_id,
                        unread_count: unreadMessageInfo.unread_count
                    })
                } else {
                    if (f.message_id >= unreadMessageInfo.latest_message.message_id) {
                        continue
                    }
                    tempMap.set(unreadMessageInfo.latest_message.sender_id, {
                        friend_id: unreadMessageInfo.latest_message.sender_id,
                        message_id: unreadMessageInfo.latest_message.message_id,
                        uid: unreadMessageInfo.latest_message.receiver_id,
                        unread_count: f.unread_count + unreadMessageInfo.unread_count
                    })
                }
            }
            setUnreadMessageCountMap(tempMap)
            await DexieStoreUnreadMessagesCounts(currentUser?.uid as string, tempMap)
        })

    }
    // 跳转到聊天界面
    const toDialog = (friend_id: string | undefined) => {
        if (isUndefined(friend_id)) {
            return
        }
        setDialogState(friend_id)
        navigate("/dialogue")
    }
    useEffect(() => {
        if (isNull(currentUser)) {
            setCurrentUser(RefreshCurrentUser())
        }
        ListenUnreadMessage()
    }, [])
    useEffect(() => {
        initUnreadMessageCountMap().then()
        ListenUnreadMessage()
        return () => {
            DexieStoreUnreadMessagesCounts(currentUser?.uid as string, unreadMessageCountMap).then()
        }
    }, [currentUser]);
    useEffect(() => {
        initMessageList().then()
        ListenUnreadMessage()
    }, [unreadMessageCountMap]);
    return (
        <div className={"chat"}>
            <HomeHeader title="pangchat"/>
            <div className="chat__message__list">
                {
                    messageList.map((value) => {
                        return (
                            <div className="chat__message__item" key={value.message_id} onClick={() => {
                                toDialog(value.friend_id)
                            }}>
                                <div className="chat__message__item__avatar">
                                    <img
                                        src={getAvatar(value.friend_id as string)}
                                        alt=""/>
                                </div>
                                <div className="chat__message__item__mid">
                                    <div className="chat__message__item__name">
                                        {getName(value.friend_id as string)}
                                    </div>
                                    <div className="chat__message__item__content">
                                        {value.content}
                                    </div>
                                </div>
                                <div className="chat__message__item__right">
                                    <div className="chat__message__item__time">
                                        {getTime(value.send_at)}
                                    </div>
                                    {
                                        getUnreadMessageCount(value.friend_id as string) === 0 ? "" :
                                            <div className="chat__message__item__unread__count">
                                                {
                                                    getUnreadMessageCount(value.friend_id as string)
                                                }
                                            </div>
                                    }
                                </div>
                            </div>
                        )
                    })
                }

            </div>
        </div>
    )
}

export default Chat