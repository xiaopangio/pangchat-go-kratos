import HomeHeader from "@/components/HomeHeader";
import "@/pages/contact/index.less"
import newFriend from "@img/newFriend.png"
import Group from "@img/Group.png"
import {useNavigate} from "react-router";
import {useRecoilState} from "recoil";
import {UnreadFriendRequestCount} from "@/store";

function Contact() {
    let navigate = useNavigate();
    const [unreadFriendRequestCount, setUnreadFriendRequestCount] = useRecoilState(UnreadFriendRequestCount)
    const newFriends = () => {
        navigate("/newFriends")
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
            <div className="contact-list">
                <div className="contact-header">
                    家人
                </div>
                <div className="contact-item">
                    <div className="contact-item-icon">
                        <img src={newFriend} alt=""/>
                    </div>
                    <div className="contact-item-text">
                        新的朋友
                    </div>
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
        </div>
    )
}

export default Contact