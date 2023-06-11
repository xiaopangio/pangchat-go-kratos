import "@/pages/Profile/index.less"
import {useEffect, useState} from "react";
import {useNavigate} from "react-router";
import SettingHeader from "@/components/SettingHeader";
import SvgIcon from "@/components/Icon";
import {GetAvatar} from "@/api/user";
import {isNull, isUndefined} from "lodash";
import {useRecoilState, useRecoilValue} from "recoil";
import {currentDialogState, currentUserState, DataSetFriend, SearchUserState} from "@/store";
import {DexieGetFriend, GetImg, StoreImg} from "@/utils/store";
import {RefreshCurrentUser} from "@/utils/util";
import {Friend} from "@/store/db";

function Profile() {
    let navigate = useNavigate();
    const [imageData, setImageData] = useState("");
    let profile = useRecoilValue(SearchUserState);
    const [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    const [isFriend, setIsFriend] = useState(false);
    const [friend, setFriend] = useState<Friend>();
    const [dialogState, setDialogState] = useRecoilState(currentDialogState)
    const [, setDataSetFriend] = useRecoilState(DataSetFriend)
    const CheckIsFriend = async (uid: string, otherId: string) => {
        try {
            const res = await DexieGetFriend(uid, otherId)
            setFriend(res)
            return true
        } catch (e) {
            return false
        }
    }
    const ShowNoteName = () => {
        return !isUndefined(friend?.note_name) && friend?.note_name !== profile?.nick_name ? friend?.note_name : profile?.nick_name
    }
    const ShowNickName = () => {
        return profile?.nick_name
    }
    const IsShowNickName = () => {
        return friend?.note_name === profile?.nick_name
    }
    const IsShowAddress = () => {
        return profile?.city !== ""
    }
    const IsShowAccount = () => {
        return isFriend
    }
    const initData = async () => {
        if (isNull(currentUser)) {
            return;
        }
        if (isNull(profile)) {
            return
        }
        const isFriend = await CheckIsFriend(currentUser.uid, profile.user_id)
        setIsFriend(isFriend)
        try {
            let avatar = await GetImg(profile?.avatar_url as string)
            if (avatar !== "") {
                setImageData(avatar)
                return
            }
            const value = await GetAvatar({avatar_url: profile.avatar_url})
            avatar = await StoreImg(profile?.avatar_url as string, value)
            setImageData(avatar)
        } catch (e) {
            const value = await GetAvatar({avatar_url: profile.avatar_url})
            const avatar = await StoreImg(profile?.avatar_url as string, value)
            setImageData(avatar)
        }
    }
    useEffect(() => {
        if (isNull(currentUser)) {
            setCurrentUser(RefreshCurrentUser())
            return;
        }
    }, []);
    useEffect(() => {
        initData()
    }, [currentUser]);

    const back = () => {
        navigate(-1)
    }
    const addFriend = () => {
        navigate("/addDetail")
    }
    const chat = () => {
        setDialogState(profile?.user_id as string)
        navigate("/dialogue")
    }
    const goToDataSet = () => {
        if (isUndefined(friend)) {
            return
        }
        setDataSetFriend(friend)
        navigate("/dataSet")
    }
    return (
        <div className="search-target">
            <SettingHeader back={back} title="" menu={goToDataSet}/>
            <div className="target-profile">
                <img className="target-profile__avatar" alt="" src={imageData}/>
                <div className="target-profile__box">
                    <div className="target-profile__name">{ShowNoteName()}</div>
                    {
                        IsShowNickName() ?
                            <div className="target-profile-other">{`昵称: ${ShowNickName()}`}</div> : null
                    }
                    {
                        IsShowAddress() ?
                            <div
                                className="target-profile-other">地区: {profile?.province === profile?.city ? "中国大陆" : profile?.province} {profile?.city}</div> : null
                    }
                    {
                        IsShowAccount() ?
                            <div className="target-profile-other">微信号: {profile?.account_id}</div> : null
                    }
                </div>
            </div>
            <div className="target-note-set">
                <div className="target-note-set__title">设置备注和标签</div>
                <SvgIcon name="rightArr"/>
            </div>
            {
                profile?.desc === "" ? null :
                    <div className="target-desc">个性签名 <span>为简单生活而努力着</span></div>
            }
            {
                isFriend ? (<div>
                    <div className="target-btn" onClick={chat}>发消息</div>
                    <div className="target-btn" onClick={addFriend}>音视频通话</div>
                </div>) : <div className="target-btn" onClick={addFriend}>添加到通讯录</div>
            }
        </div>
    )
}

export default Profile