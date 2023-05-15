import "@/pages/onboarding/index.less"
import {useNavigate} from "react-router";
import {useRecoilValue} from "recoil";
import {loginState} from "@/store";

function Onboarding() {
    let navigate = useNavigate();
    const isLogin = useRecoilValue(loginState)
    const gotoNext = () => {
        navigate("/login")
    }
    return (
        <div className={"boarding"}>
            <div className={"boarding-top"}>
                <div className={"boarding-top-content"}>
                    <div className={"boarding-top-content-logo"}></div>
                    <div className={"boarding-top-content-title"}>PangChat</div>
                </div>
                <div className={"boarding-top-imgs"}>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                </div>
            </div>
            <div className={"boarding-bottom"}>
                <div className={"boarding-bottom-content1"}>与你的朋友和家人保持联系</div>
                <div className={"boarding-bottom-content2"}>PangChat是聊天通信程序,这将有助于你与每一个人。
                </div>
                <button className={"boarding-bottom-content3"} onClick={
                    gotoNext
                }>
                    开始使用
                </button>
            </div>
        </div>
    );
}

export default Onboarding;