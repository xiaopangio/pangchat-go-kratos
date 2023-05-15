import "@/pages/contact/showFriendRequestDetail/index.less"
import SettingHeader from "@/components/SettingHeader";
import {useNavigate} from "react-router";
import DefaultImg from "@img/default.jpg";
import {useRecoilState} from "recoil";
import {CurrentDealFriendRequest, currentUserState} from "@/store";
import {useEffect, useState} from "react";
import {GetImg} from "@/utils/store";
import {isNull} from "lodash";
import {isEmpty} from "fast-glob/out/utils/string";
import SvgIcon from "@/components/Icon";

function ShowFriendRequestDetail() {
    const [currentDealFriendRequest, setCurrentDealFriendRequest] = useRecoilState(CurrentDealFriendRequest)
    const [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    let navigate = useNavigate();
    const [avatarData, setAvatarData] = useState("");
    useEffect(() => {
        if (isNull(currentDealFriendRequest)) {
            return
        }
        GetImg(currentDealFriendRequest.avatar).then((res) => {
            setAvatarData(res)
        }, (reason) => {
            console.log(reason)
        })
    }, [])
    const GetAvatar = () => {
        if (!isEmpty(avatarData)) {
            return avatarData
        } else {
            return DefaultImg
        }
    }
    const back = () => {
        navigate("/newFriends")
    }
    const toDeal = () => {
        if (currentUser?.uid === currentDealFriendRequest?.requester_id) {
            return
        }
        navigate("/dealFriendRequest")
    }
    return (
        <div className="showFriendRequestDetail">
            <SettingHeader back={back} title=""/>
            <div className="header">
                <div className="avatar">
                    <img alt="" src={GetAvatar()}/>
                </div>
                <div className="content">
                    <div className="nickname">{currentDealFriendRequest?.nick_name}</div>
                    <div className="setting">设置备注和标签</div>
                </div>
            </div>
            <div className="desc-box">
                <div className="desc">
                    <div className="content">
                        {currentUser?.uid === currentDealFriendRequest?.requester_id ? "我" : currentDealFriendRequest?.nick_name}:{currentDealFriendRequest?.desc}
                    </div>
                    <div className="btn">
                        回复
                    </div>
                </div>
            </div>
            <div className="list">
                <div className="item">
                    <span>朋友圈</span>
                    <span className="right"> <SvgIcon name="rightArr"/></span>
                </div>
                <div className="item">
                    <span>来源</span>
                    <span className="center">对方通过搜索账号添加</span>
                </div>
            </div>
            <div className="goto-verify" onClick={toDeal}>
                {currentUser?.uid === currentDealFriendRequest?.requester_id ? "等待验证" : "前往验证"}
            </div>
            <div className="footer">
                <div className="item">
                    加入黑名单
                </div>
                <div className="item">
                    投诉
                </div>
            </div>
        </div>
    )
}

export default ShowFriendRequestDetail