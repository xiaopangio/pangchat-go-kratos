import SettingHeader from "@/components/SettingHeader";
import "@/pages/addFriend/addDetail/index.less"
import SvgIcon from "@/components/Icon";
import {useNavigate} from "react-router";
import {useRecoilState, useRecoilValue} from "recoil";
import {currentUserState, GroupName, GroupNames, SearchUserState} from "@/store";
import {useEffect, useState} from "react";
import {GetFriendGroupList, SendFriendRequest} from "@/api/friend";
import {db, FriendGroup} from "@/store/db";
import {DexieGetFriendGroup} from "@/utils/store";
import {useRefreshUser} from "@/hooks/refreshUser";

export function GetGroupNames(uid: string, onSuccess: (res: string[]) => void) {
    GetFriendGroupList({user_id: uid}).then(res => {
        let groups = res.group_names.map((item: string) => {
            return {name: item, uid: uid} as FriendGroup
        })
        db.friendGroups.bulkAdd(groups)
        onSuccess(res.group_names)
    })
}

function AddDetail() {
    let navigate = useNavigate();
    let profile = useRecoilValue(SearchUserState);
    let currentUser = useRecoilValue(currentUserState);
    const [groupNames, setGroupNames] = useRecoilState(GroupNames);
    const [groupName, setGroupName] = useRecoilState(GroupName);
    const [desc, setDesc] = useState("");
    const [noteName, setNoteName] = useState("");
    useRefreshUser()
    useEffect(() => {
        if (!currentUser) {
            return
        }
        if (groupNames.length > 0) {
            return
        }
        DexieGetFriendGroup(currentUser.uid).then((res) => {
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
    }, []);

    const back = () => {
        navigate("/Profile")
    }
    const chooseGroup = () => {
        navigate("/chooseGroup")
    }
    const handleInputDesc = (e: any) => {
        setDesc(e.target.value)
    }
    const handleInputNoteName = (e: any) => {
        setNoteName(e.target.value)
    }
    const handleAdd = async () => {
        if (!profile) {
            return
        }
        if (!currentUser) {
            return
        }
        try {
            let res = await SendFriendRequest({
                desc: desc === "" ? `我是${currentUser?.nick_name}，我想加您为好友` : desc,
                group_name: groupName,
                note_name: noteName === "" ? profile.nick_name : noteName,
                receiver_id: profile.user_id,
                requester_id: currentUser.uid
            })
            console.log(res)
            res.request.uid = currentUser.uid
            db.friendRequests.add(res.request).then(() => {
                console.log("添加成功")
            })
        } catch (e) {
            console.log(e)
        }
        navigate("/newFriends")
    }
    return (
        <div className="addDetail">
            <SettingHeader back={back} title="申请添加朋友"/>
            <div className="addDetail-content">
                <div className="addDetail-item">
                    <div className="addDetail-item-title">发送添加朋友申请</div>
                    <textarea className="addDetail-item-content textarea-item"
                              placeholder={`我是${currentUser?.nick_name}，我想加您为好友`} value={desc}
                              onInput={handleInputDesc}>
                    </textarea>
                </div>
                <div className="addDetail-item">
                    <div className="addDetail-item-title">设置备注</div>
                    <input className="addDetail-item-content" placeholder={profile?.nick_name}
                           onInput={handleInputNoteName}/>
                </div>
                <div className="addDetail-item">
                    <div className="addDetail-item-title">设置分组</div>
                    <div className="addDetail-item-content"
                         style={{
                             display: "flex",
                             justifyContent: "space-between",
                             padding: "10px",
                             alignItems: "center"
                         }}
                         onClick={chooseGroup}>
                        {groupName}
                        <SvgIcon name="rightArr"/>
                    </div>
                </div>
                <div className="add-submit" onClick={handleAdd}>
                    发送
                </div>
            </div>
        </div>
    )
}

export default AddDetail