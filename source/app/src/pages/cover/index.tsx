import {useNavigate} from "react-router";
import "@/pages/cover/index.less"
import {useEffect} from "react";
import {useRecoilState} from "recoil";
import {currentUserState} from "@/store";
import {HomePrefix} from "@/declare/const";
import {checkToken} from "@/utils/check";
import {ParseJwt} from "@/utils/jwt";
import {User} from "@/declare/type";

function Cover() {
    let navigate = useNavigate();
    const [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    useEffect(() => {
        let token = checkToken();
        let user: User | null = null;
        if (token) {
            user = ParseJwt<User>(token.value);
        }
        setTimeout(() => {
            if (user) {
                setCurrentUser(user);
                navigate(HomePrefix + "/chat")
                return
            } else {
                navigate("/board")
            }
        }, 1500)
    }, [])
    return (
        <div className={"cover"}>
            <div className={"cover-content"}>
                <div className={"cover-logo"}></div>
                <div className={"cover-title"}>PangChat</div>
            </div>
        </div>
    )
}

export default Cover