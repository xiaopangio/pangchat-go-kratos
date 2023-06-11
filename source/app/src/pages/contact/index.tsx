import HomeHeader from "@/components/HomeHeader";
import "@/pages/contact/index.less"
import newFriend from "@img/newFriend.png"
import Group from "@img/Group.png"
import {useNavigate} from "react-router";
import {useRecoilState} from "recoil";
import {currentUserState, SearchUserState, UnreadFriendRequestCount} from "@/store";
import {useEffect, useRef, useState} from "react";
import {isNull, isUndefined} from "lodash";
import {RefreshCurrentUser} from "@/utils/util";
import {DexieAddFriends, DexieGetFriends, DexieGetImgList, DexieUpdateFriendInfo, StoreImg} from "@/utils/store";
import {GetFriendInfo, GetFriendList} from "@/api/friend";
import {Friend} from "@/store/db";
import {GetAvatar} from "@/api/user";
import DefaultImg from "@img/default.jpg";
import {isEmpty} from "fast-glob/out/utils/string";

interface FriendGroup {
    group_name: string
    friends: Friend[]
}

function Contact() {
    let navigate = useNavigate();
    const [unreadFriendRequestCount, setUnreadFriendRequestCount] = useRecoilState(UnreadFriendRequestCount)
    const [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    const [friendGroups, setFriendGroups] = useState<FriendGroup[]>([]);
    const [avatarMap, setAvatarMap] = useState<Map<string, string>>(new Map<string, string>());
    const [, setProfile] = useRecoilState(SearchUserState);
    const longPressInterval = 600;
    const longPressTimer = useRef<any>(null);
    const [isShowPopper, setIsShowPopper] = useState(false);
    const [pressGroupName, setPressGroupName] = useState("");
    const dialogRef = useRef<HTMLDialogElement>(null);
    let pressStartTime = 0;
    const InitData = async () => {
        if (isNull(currentUser)) {
            return
        }
        let res: Friend[] = [];
        try {
            res = await DexieGetFriends(currentUser.uid);
        } catch (e) {
            console.log(e)
        }

        if (res.length === 0) {
            const r = await GetFriendList(currentUser.uid)
            await DexieAddFriends(currentUser.uid, r.friends)
            res.push(...r.friends)
        }
        let friendGroups: FriendGroup[] = []
        //根据group_name分组
        res.forEach((item) => {
            let index = friendGroups.findIndex((group) => {
                return group.group_name === item.group_name
            })
            if (index === -1) {
                friendGroups.push({
                    group_name: item.group_name,
                    friends: [item]
                })
            } else {
                friendGroups[index].friends.push(item)
            }
        })
        const avatarMap = await InitAvatar(friendGroups)
        setFriendGroups(friendGroups)
        setAvatarMap(avatarMap)
    }
    const InitAvatar = async (friendGroups: FriendGroup[]) => {
        let map = new Map<string, string>()
        let avatarUrls: string[] = []
        friendGroups.forEach((group) => {
            group.friends.forEach((friend) => {
                avatarUrls.push(friend.avatar)
            })
        })
        try {
            const res = await DexieGetImgList(avatarUrls)
            res.forEach((item) => {
                map.set(item.name, item.blob)
            })
        } catch (diff: any) {
            for (const item of diff) {
                try {
                    const res = await GetAvatar({avatar_url: item})
                    try {
                        const r = await StoreImg(item, res)
                        map.set(item, r)
                    } catch (e) {
                        console.log(e)
                    }
                } catch (e) {
                    console.log(e)
                }
            }
            await InitAvatar(friendGroups)
        }
        return map
    }
    const getAvatar = (avatar: string) => {
        if (isNull(avatarMap)) {
            return
        }
        if (avatarMap.has(avatar)) {
            return avatarMap.get(avatar)
        } else {
            return DefaultImg
        }
    }
    useEffect(() => {
        if (isNull(currentUser)) {
            let user = RefreshCurrentUser();
            setCurrentUser(user)
        }
    }, []);
    useEffect(() => {
        InitData().then(r => r)
    }, [currentUser]);
    const newFriends = () => {
        navigate("/newFriends")
    }
    const showProfile = async (friend: Friend) => {
        if (isUndefined(friend.account_id) || isNull(friend.account_id) || isEmpty(friend.account_id)) {
            try {
                const res = await GetFriendInfo(friend.friend_id)
                console.log(res)
                friend.account_id = res.account_id
                friend.city_name = res.city_name
                friend.desc = res.desc
                friend.province_name = res.province_name
                await DexieUpdateFriendInfo(friend)
            } catch (e) {
                console.log(e)
            }
        }
        setProfile(
            {
                account_id: friend.account_id,
                avatar_url: friend.avatar,
                city: friend.city_name,
                desc: friend.desc,
                nick_name: friend.nick_name,
                province: friend.province_name,
                user_id: friend.friend_id
            }
        )
        navigate("/profile")
    }
    //长按事件，弹出一个选项框，可以跳转到分组管理
    const longPress = (group_name: string) => {
        setPressGroupName(group_name)
        dialogRef.current?.showModal()
        dialogRef.current?.animate([
            {transform: "translateY(100%)"},
            {transform: "translateY(0)"}
        ], {
            duration: 200,
            fill: "forwards"
        })
    }
    const itemTouchStart = (group_name: string) => {
        pressStartTime = new Date().getTime()
        longPressTimer.current = setTimeout(() => {
            longPress(group_name)
        }, longPressInterval)
    }
    const itemTouchEnd = () => {
        let endTime = new Date().getTime()
        if (endTime - pressStartTime < longPressInterval) {
            clearTimeout(longPressTimer.current)
        }
    }
    const gotoFriendGroupManage = () => {
        dialogRef.current?.close()
        navigate("/groupManager")
    }
    const closePopper = () => {
        if (isNull(dialogRef.current)) {
            return
        }
        dialogRef.current.animate([
            {transform: "translateY(0)"},
            {transform: "translateY(100%)"}
        ], {
            duration: 200,
            fill: "forwards"
        })
        setTimeout(() => {
            dialogRef.current?.close()
        }, 200)
    }
    const handleClose = (event: any) => {
        event.preventDefault();
        if (event.target === dialogRef.current) {
            closePopper()
        }
    }
    return (
        <div className="contact">
            <HomeHeader title="通讯录"/>
            <div className="contact-list">
                <div className="contact-item" onClick={newFriends}>
                    <div className="contact-item-icon">
                        <img src={newFriend} alt=""/>
                    </div>
                    <div className="contact-item-text">
                        新的朋友
                    </div>
                    {
                        unreadFriendRequestCount === 0 ? null :
                            <div className="contact-unread-count">{unreadFriendRequestCount}</div>
                    }
                </div>
                <div className="contact-item">
                    <div className="contact-item-icon">
                        <img src={Group} alt=""/>
                    </div>
                    <div className="contact-item-text">
                        群聊
                    </div>
                </div>
            </div>
            {
                friendGroups.map((group) => {
                    return (
                        <div className="contact-list" key={group.group_name}>
                            <div className="contact-header" onTouchStart={() => {
                                itemTouchStart(group.group_name)
                            }} onTouchEnd={itemTouchEnd}>
                                {group.group_name}
                            </div>
                            {
                                group.friends.map((friend) => {
                                    return (
                                        <div className="contact-item" key={friend.friend_id} onClick={() => {
                                            showProfile(friend)
                                        }}>
                                            <div className="contact-item-icon">
                                                <img src={getAvatar(friend.avatar)} alt=""/>
                                            </div>
                                            <div className="contact-item-text">
                                                {friend.nick_name}
                                            </div>
                                        </div>
                                    )
                                })
                            }
                        </div>
                    )
                })
            }
            <dialog ref={dialogRef} className="contact-popper" onClick={handleClose}>
                <div className="contact-popper-item gray">
                    设置好友权限"{pressGroupName}"
                </div>
                <div className="contact-popper-item">设置好友权限</div>

                <div className="contact-popper-item bottom-shadow" onClick={gotoFriendGroupManage}>
                    分组管理
                </div>
                <div className="contact-popper-item mt" onClick={closePopper}>取消</div>
            </dialog>
        </div>
    )
}

export default Contact