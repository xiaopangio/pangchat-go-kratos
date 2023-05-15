import {useNavigate} from "react-router";

export function NavigateToLogin() {
    let navigate = useNavigate();
    return (
        <div
            style={{float: "right", marginTop: "20px", marginRight: "20px", color: "#4FA075"}}
            onClick={() => {
                navigate("/login")
            }}>
            立即登录
        </div>
    )

}