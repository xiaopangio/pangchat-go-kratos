import "@/pages/login/index.less"
import SvgIcon from "@/components/Icon";
import {FormEvent, useEffect, useState} from "react";
import {useNavigate} from "react-router";
import {useInterval} from "@/hooks/intervalHooks";
import {useLogin} from "@/hooks/loginHooks";
import {LoginData} from "@/api/user/types";
import {HomePrefix} from "@/declare/const";
import storage from "@/utils/storage";
import {useRecoilState} from "recoil";
import {currentUserState, loginState} from "@/store";
import hashPassword from "@/utils/password";
import {ParseJwt} from "@/utils/jwt";
import {User} from "@/declare/type";
import {checkPhone} from "@/utils/check";
import message from "@/utils/message";
import {useRegister} from "@/hooks/registerHooks";
import {websocketInit} from "@/hooks/websocket";
import {Connect, InitConnection} from "@/service/service";

function Login() {
    const [loginType, setLoginType] = useState("phone");
    const [isTickerValid, setIsTickerValid] = useState(false);
    const [ticker, setTicker] = useState(60);
    const [password, setPassword] = useState("");
    const [phone, setPhone] = useState("");
    const [code, setCode] = useState("");
    const [accountID, setAccountID] = useState("");
    const [loginS, setLoginS] = useRecoilState(loginState)
    const [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    const {sendSmsCode} = useRegister();
    let navigate = useNavigate();
    const {login} = useLogin();
    const [reset, clear] = useInterval(() => {
        setTicker(ticker - 1);
    })
    const getCodeClick = async (e: any) => {
        e.preventDefault();
        if (!checkPhone(phone)) {
            message.error({content: "请检查手机号是否正确", duration: 2000})
            return;
        }

        if (isTickerValid) {
            return;
        }
        setIsTickerValid(true);
        reset()

        try {
            await sendSmsCode({phone: phone});
        } catch (e) {
            console.log(e)
        }
    }
    useEffect(() => {
        clear()
        setIsTickerValid(false);
    }, []);

    useEffect(() => {
        if (ticker <= 0) {
            clear();
            setTicker(60)
            setIsTickerValid(false);
        }
    }, [ticker]);
    const loginSubmit = async (e: FormEvent<HTMLButtonElement>) => {
        e.preventDefault();
        if (loginType === "account" && (accountID === "" || password === "")) {
            return;
        } else if (loginType === "phone" && (phone === "" || code === "")) {
            return;
        }
        const loginData: LoginData = {
            type: 1,
            username: phone,
            password: code
        }
        if (loginType === "account") {
            loginData.type = 2;
            loginData.username = accountID;
            loginData.password = hashPassword(password) as string
        }
        try {
            const resp = await login(loginData);
            storage().set(
                {
                    key: "token",
                    value: resp.token,
                    expire: '30d'
                }
            )
            let user = ParseJwt<User>(resp.token);
            console.log(user)
            setCurrentUser(user)
            setLoginS(true)
            try {
                await Connect(function (cUrl: string) {
                    websocketInit(cUrl)
                })
            } catch (e) {
                console.log(e)
            }
            InitConnection()
            navigate(HomePrefix + "/chat")
        } catch (e) {
            console.log(e)
        }
    }
    const onAccountInput = (e: FormEvent<HTMLInputElement>) => {
        setAccountID(e.currentTarget.value)
    }
    const onPasswordInput = (e: FormEvent<HTMLInputElement>) => {
        setPassword(e.currentTarget.value)
    }
    const onPhoneInput = (e: FormEvent<HTMLInputElement>) => {
        setPhone(e.currentTarget.value)
    }
    const onCodeInput = (e: FormEvent<HTMLInputElement>) => {
        setCode(e.currentTarget.value)
    }
    return (
        <div className="login">
            <div className="login-logo">
                <SvgIcon name="login" width={80}></SvgIcon>
            </div>
            <form action="login" className="login-form">
                {loginType === "account" ? (
                    <div className="account-login">
                        <div className="login-form-item">
                            <div className="login-label">账号</div>
                            <input className="login-input" placeholder={"Account"} value={accountID}
                                   onInput={onAccountInput}></input>
                        </div>

                        <div className="login-form-item">
                            <div className="login-label">密码</div>
                            <input className="login-input" placeholder={"Password"} type="password" value={password}
                                   onInput={onPasswordInput}></input>
                        </div>
                    </div>
                ) : (
                    <div className="phone-login">
                        <div className="login-form-item">
                            <div className="login-label">手机号</div>
                            <input className="login-input" placeholder={"Phone Number"} value={phone}
                                   onInput={onPhoneInput}></input>
                        </div>
                        <div className="login-form-item">
                            <div className="login-label">验证码</div>
                            <div className="flex">
                                <input className="login-input" placeholder={"Verify Code"} value={code}
                                       onInput={onCodeInput}
                                       style={{width: "230px"}}></input>
                                <button onClick={getCodeClick} className={"send-code"}>
                                    {isTickerValid ? ticker : "发送"}
                                </button>
                            </div>
                        </div>
                    </div>
                )
                }
                <button type="submit" className="login-submit" onClick={loginSubmit}>登录</button>
            </form>
            <div className="login-footer">
                <div className="login-footer-left">
                    {
                        loginType === "account" ?
                            <span onClick={() => setLoginType("phone")}>手机登录</span> :
                            <span onClick={() => setLoginType("account")}>账号登录</span>
                    }
                </div>
                <div className="login-footer-right">
                    <span onClick={() => {
                        navigate("/forgetPassword")
                    }}>忘记密码</span>
                    <span onClick={() => {
                        navigate("/register")
                    }}>注册</span>
                </div>
            </div>
        </div>
    )
}

export default Login;