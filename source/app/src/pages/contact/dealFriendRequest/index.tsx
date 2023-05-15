import "@/pages/contact/dealFriendRequest/index.less"
import {useRecoilState} from "recoil";
import {CurrentDealFriendRequest} from "@/store";
import SettingHeader from "@/components/SettingHeader";
import {useNavigate} from "react-router";
import SvgIcon from "@/components/Icon";
import {useState} from "react";

function DealFriendRequest() {
    const [currentDealFriendRequest, setCurrentDealFriendRequest] = useRecoilState(CurrentDealFriendRequest)
    let navigate = useNavigate();
    const [friendsPermissions, setFriendsPermissions] = useState("all");
    const selectFriendsPermissions = (type: string) => {
        setFriendsPermissions(type)
    }
    const back = () => {
        navigate("/newFriends")
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
                        <input type="text" placeholder={currentDealFriendRequest?.nick_name}/>
                    </div>
                    <div className="text-box">
                        <span className="text1">{`"${currentDealFriendRequest?.desc}"`}</span>
                        <span className="text2">选词填入</span>
                    </div>
                </div>
                <div className="label">
                    添加标签与描述
                </div>
                <div className="list">
                    <div className="item border-bottom">
                        <div>标签</div>
                        <span><SvgIcon name="rightArr"/></span>
                    </div>
                    <div className="item">
                        <div>描述</div>
                        <span><SvgIcon name="rightArr"/></span>
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
        </div>
    )
}

export default DealFriendRequest