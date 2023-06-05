import "@/pages/contact/newFriends/index.less"
import SettingHeader from "@/components/SettingHeader";
import {useNavigate} from "react-router";
import {AddSearchPrefix, HomePrefix} from "@/declare/const";
import SvgIcon from "@/components/Icon";
import DefaultImg from "@img/default.jpg";
import {useRecoilState} from "recoil";
import {CurrentDealFriendRequest, currentUserState, UnreadFriendRequestCount} from "@/store";
import {useEffect, useState} from "react";
import {FriendRequest} from "@/store/db";
import storage from "@/utils/storage";
import {DexieGetImgList, GetFriendRequests, StoreImg} from "@/utils/store";
import {isNull} from "lodash";
import {RefreshCurrentUser} from "@/utils/util";
import {GetAvatar} from "@/api/user";

function NewFriends() {
    let navigate = useNavigate();
    const [friendRequestsInThreeDays, setFriendRequestsInThreeDays] = useState<FriendRequest[]>([]);
    const [friendRequestsBeforeThreeDays, setFriendRequestsBeforeThreeDays] = useState<FriendRequest[]>([]);
    const [unreadFriendRequestCount, setUnreadFriendRequestCount] = useRecoilState(UnreadFriendRequestCount)
    const [currentDealFriendRequest, setCurrentDealFriendRequest] = useRecoilState(CurrentDealFriendRequest)
    const [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    const [avatarMap, setAvatarMap] = useState<Map<string, string>>();
    useEffect(() => {
        if (isNull(currentUser)) {
            let user = RefreshCurrentUser();
            setCurrentUser(user)
        } else {
            return
        }
    }, []);
    useEffect(() => {
        console.log("unreadFriendRequestCount", unreadFriendRequestCount)
    }, [unreadFriendRequestCount]);

    useEffect(() => {
        if (!currentUser) {
            return
        }
        storage().delete("unreadFriendRequestsCount")
        setUnreadFriendRequestCount(0)
        GetFriendRequests(currentUser.uid).then((res) => {
            let threeDaysAgo = new Date().getTime() - 3 * 24 * 60 * 60 * 1000
            let requestInThreeDays: FriendRequest[] = []
            let requestBeforeThreeDays: FriendRequest[] = []
            res.forEach((item: FriendRequest) => {
                new Date(item.create_time).getTime() > threeDaysAgo ? requestInThreeDays.push(item) : requestBeforeThreeDays.push(item)
            })
            setFriendRequestsInThreeDays(requestInThreeDays)
            setFriendRequestsBeforeThreeDays(requestBeforeThreeDays)
            let map = new Map<string, string>()
            let avatarUrls: string[] = []
            requestInThreeDays.forEach((item) => {
                avatarUrls.push(item.avatar)
            })
            requestBeforeThreeDays.forEach((item) => {
                avatarUrls.push(item.avatar)
            })
            DexieGetImgList(avatarUrls).then((res) => {
                res.forEach((item) => {
                    map.set(item.name, item.blob)
                })
                setAvatarMap(map)
            }).catch((diff) => {
                diff.forEach((item: any) => {
                    GetAvatar({avatar_url: item}).then((res) => {
                        StoreImg(item, res).then((res) => {
                            map.set(item, res)
                        }).catch((err) => {
                            console.log(err)
                        })
                    })
                })
            })
        }, () => {

        })
    }, [currentUser, unreadFriendRequestCount])
    const showAvatar = (avatar: string) => {
        return avatarMap?.get(avatar) as string || DefaultImg
    }
    const handleSearch = () => {
        navigate(AddSearchPrefix)
    }
    const back = () => {
        navigate(HomePrefix + "/contact")
    }
    const toDeal = (item: FriendRequest) => {
        setCurrentDealFriendRequest(item)
        navigate("/dealFriendRequest")
    }
    const toDetail = (item: FriendRequest) => {
        setCurrentDealFriendRequest(item)
        navigate("/showFriendRequestDetail")
    }
    return (
        <div className="new-friends">
            <SettingHeader back={back} title="新的朋友"/>
            <div className="add-click-box" onClick={handleSearch}>
                <SvgIcon name="search" width={20} color="#b5b5b5"/> <span>账号/手机号</span>
            </div>
            <div className="new-friends-item">
                <div className="new-friends-item-left">
                    <SvgIcon name="phone"/>
                </div>
                <div className="new-friends-item-center">
                    添加手机联系人
                </div>
                <div className="new-friends-item-right">
                    <span><SvgIcon name="rightArr"/></span>
                </div>
            </div>
            <div className="new-friends-label">
                近三天
            </div>
            {friendRequestsInThreeDays.map((item, index) => {
                return (
                    <div className="new-friends-item" key={item.request_id}>
                        <div className="new-friends-item-left">
                            <img src={showAvatar(item.avatar)} alt=""/>
                        </div>
                        <div className="new-friends-item-center" onClick={() => {
                            toDetail(item)
                        }}>
                            <div className="new-friends-item-center-top">
                                {item.nick_name}
                            </div>
                            <div className="new-friends-item-center-bottom">
                                {currentUser?.uid === item.requester_id ? `我: ${item.desc}` : `${item.desc}`}
                            </div>
                        </div>
                        <div className="new-friends-item-right">
                            {currentUser?.uid === item.requester_id ?
                                <SvgIcon name="ruArr" width={20} color="#737373"/> : null}
                            <span>
                                    {
                                        item.status === "0" ? currentUser?.uid === item.requester_id ? '等待验证' :
                                            <button
                                                className="accept-btn" onClick={() => {
                                                toDeal(item)
                                            }}>接受</button> : item.status === "1" ? '已添加' : '已拒绝'
                                    }
                                </span>
                        </div>
                    </div>
                )
            })}
            <div className="new-friends-label">
                三天前
            </div>
            {friendRequestsBeforeThreeDays.map((item, index) => {
                return (
                    <div className="new-friends-item" key={item.request_id}>
                        <div className="new-friends-item-left">
                            <img src={DefaultImg} alt=""/>
                        </div>
                        <div className="new-friends-item-center">
                            <div className="new-friends-item-center-top">
                                {currentUser?.uid === item.requester_id ? item.receiver_id : item.requester_id}
                            </div>
                            <div className="new-friends-item-center-bottom">
                                {currentUser?.uid === item.requester_id ? `我: ${item.desc}` : `${item.desc}`}
                            </div>
                        </div>
                        <div className="new-friends-item-right">
                            {currentUser?.uid === item.requester_id ?
                                <SvgIcon name="ruArr" width={20} color="#737373"/> : null}
                            <span>
                                    {
                                        item.status === "0" ? currentUser?.uid === item.requester_id ? '等待验证' :
                                            <button
                                                className="accept-btn">接受</button> : item.status === "1" ? '已添加' : '已拒绝'
                                    }
                                </span>
                        </div>
                    </div>
                )
            })}
        </div>
    )
}

export default NewFriends