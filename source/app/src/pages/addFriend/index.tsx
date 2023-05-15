import "@/pages/addFriend/index.less"
import {useNavigate} from "react-router";
import {useRecoilValue} from "recoil";
import {currentUserState} from "@/store";
import SettingHeader from "@/components/SettingHeader";
import SvgIcon from "@/components/Icon";
import {AddSearchPrefix, HomePrefix} from "@/declare/const";

function AddFriend() {
    let navigate = useNavigate();
    let user = useRecoilValue(currentUserState);
    const toBack = () => {
        navigate(HomePrefix + "/chat")
    }
    const handleSearch = () => {
        navigate(AddSearchPrefix)
    }
    return (
        <div className="add-friend">
            <SettingHeader back={toBack} title="添加好友"/>
            <div className="add-box">
                <div className="add-click-box" onClick={handleSearch}>
                    <SvgIcon name="search" width={20} color="#b5b5b5"/> <span>账号/手机号</span>
                </div>
                <div className="add-content">
                    我的微信号: {user?.account_id}
                </div>
            </div>
            <div className="add-other-list">
                <div className="add-other-item">
                </div>
            </div>
        </div>
    )
}

export default AddFriend