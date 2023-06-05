import "@/pages/contact/dealFriendRequest/index.less"
import {useRecoilState, useRecoilValue} from "recoil";
import {CurrentDealFriendRequest, currentUserState, GroupName, GroupNames} from "@/store";
import SettingHeader from "@/components/SettingHeader";
import {useNavigate} from "react-router";
import SvgIcon from "@/components/Icon";
import {useEffect, useState} from "react";
import {GetFriendGroup, UpdateFriendRequestStatus} from "@/utils/store";
import {GetGroupNames} from "@/pages/addFriend/addDetail";
import {useRefreshUser} from "@/hooks/refreshUser";
import {DealFriendRequestData} from "@/api/friend/types";
import {DealFriendRequestApi} from "@/api/friend";
import { db } from "@/store/db";
import {isNull} from "lodash";

function DealFriendRequest() {
    const [currentDealFriendRequest, ] = useRecoilState(CurrentDealFriendRequest)
    let navigate = useNavigate();
    const [friendsPermissions, setFriendsPermissions] = useState("all");
    const [noteName, setNoteName] = useState("");
    const [, setGroupNames] = useRecoilState(GroupNames);
    const [groupName, setGroupName] = useRecoilState(GroupName);
    let currentUser = useRecoilValue(currentUserState);
    useRefreshUser()
    useEffect(() => {
        if (!currentUser) {
            return
        }
        if (groupName) {
            return
        }
        GetFriendGroup(currentUser.uid).then((res) => {
            if (res && res.length > 0) {
                setGroupNames(res)
                setGroupName(res[0])
                return
            } else {
                GetGroupNames(currentUser?.uid as string, (res) => {
                    setGroupNames(res)
                    setGroupName(res[0])
                })
            }
        })
    }, [currentUser]);
    const selectFriendsPermissions = (type: string) => {
        setFriendsPermissions(type)
    }
    const back = () => {
        navigate("/newFriends")
    }
    const chooseGroup = () => {
        navigate("/chooseGroup")
    }
    const handleDealFriendRequest = async (status:number) => {
        // 1:同意 2:拒绝
        let data:DealFriendRequestData = {
            request_id: currentDealFriendRequest?.request_id as string,
            status: 0,
            note_name: "",
            group_name: ""
        }
        if (status === 1) {
            data.status = 1
            if (noteName){
                data.note_name = noteName
            }else {
                data.note_name = currentDealFriendRequest?.nick_name as string
            }
            data.group_name = groupName
        } else {
            data.status = 2
        }
        try {
            await DealFriendRequestApi(data)
            if (isNull(currentDealFriendRequest)){
                return
            }
            const res=await UpdateFriendRequestStatus(currentDealFriendRequest.request_id,status)
            console.log(res)
            navigate("/newFriends")
        }catch (e){
            console.log(e)
        }
    }
    const inputNoteName = (e: any) => {
        setNoteName(e.target.value)
    }
    return (
        <div className="deal-friend-request">
            <SettingHeader back={back} title="处理好友请求"/>
            <div className="deal-content">
                <div className="label">
                    设置备注
                </div>
                <div className="list margin-bottom">
                    <div className="item">
                        <input type="text" placeholder={currentDealFriendRequest?.nick_name} value={noteName} onInput={inputNoteName}/>
                    </div>
                    <div className="text-box">
                        <span className="text1">{`"${currentDealFriendRequest?.desc}"`}</span>
                        <span className="text2">选词填入</span>
                    </div>
                </div>
                <div className="label">
                    选择分组
                </div>
                <div className="list">
                    <div className="item border-bottom">
                        <div>{groupName}</div>
                        <span onClick={chooseGroup}><SvgIcon name="rightArr"/></span>
                    </div>
                </div>
                <div className="label">
                    设置朋友权限
                </div>
                <div className="list">
                    <div className="item border-bottom" onClick={() => {
                        selectFriendsPermissions("all")
                    }}>
                        <div>聊天、朋友圈、微信运动等</div>
                        {
                            friendsPermissions === "all" ?
                                <span><SvgIcon name="gou" color="#05c160" width={24}/></span> : ""
                        }
                    </div>
                    <div className="item" onClick={() => {
                        selectFriendsPermissions("only")
                    }}>
                        <div>仅聊天</div>
                        {
                            friendsPermissions === "only" ?
                                <span><SvgIcon name="gou" color="#05c160" width={24}/></span> : ""
                        }
                    </div>
                </div>
            </div>
            <div className="footer">
                <div className="footer-btn red" onClick={async ()=>{await handleDealFriendRequest(2)}}>
                    拒绝
                </div>
                <div className="footer-btn" onClick={async ()=>{await handleDealFriendRequest(1)}}>
                    完成
                </div>
            </div>
        </div>
    )
}

export default DealFriendRequest