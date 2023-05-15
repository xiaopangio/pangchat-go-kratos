import "@/pages/forgetPassword/index.less"
import "@/pages/password/index.less"
import {useInterval} from "@/hooks/intervalHooks";
import {FormEvent, MouseEvent, useEffect, useState} from "react";
import SvgIcon from "@/components/Icon";
import {checkPassword} from "@/utils/check";

function ForgetPassword() {
    const [reset, clear] = useInterval(() => {
        setTicker(ticker - 1);
    })
    const [isTickerValid, setIsTickerValid] = useState(false);
    const [ticker, setTicker] = useState(60);
    const [isPassVerify, setIsPassVerify] = useState(false);
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [check, setCheck] = useState(false);
    const getCodeClick = () => {
        if (isTickerValid) {
            return;
        }
        setIsTickerValid(true);
        reset()
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
    const verifyCode = (e: MouseEvent) => {
        e.preventDefault();
        setIsPassVerify(true);
    }
    const onPasswdInput = (e: FormEvent) => {

    }
    const onConfirmInput = (e: FormEvent) => {

    }
    const onConfirmFocus = () => {

    }
    return (
        <div className="forget-password">
            {
                isPassVerify ?
                    <div className={"password"}>
                        <div className="password-label">
                            你的密码不被别人知道是很重要的
                        </div>
                        <div className="password-tip">
                            所以请设置一个你自己的密码。你可以改变它在设置中。
                        </div>
                        <div className="password-input">
                            <div className="password-input-label">密码</div>
                            <div className="password-input-box">
                                <SvgIcon name={"password"} color={"#cccccc"}/>
                                <input className="password-input-input" type={"password"} placeholder={"Password"}
                                       value={password}
                                       onInput={onPasswdInput}
                                ></input>
                            </div>
                            <div className="password-err">
                                {checkPassword(password) ? "" : "密码长度必须在8-16位，由字母、数字、特殊字符组成"}
                            </div>
                        </div>
                        <div className="password-input">
                            <div className="password-input-label">确认</div>
                            <div className="password-input-box">
                                <SvgIcon name={"password-b"} color={"#cccccc"}/>
                                <input className="password-input-input" type={"password"} placeholder={"Confirm"}
                                       onInput={onConfirmInput}
                                       onFocus={onConfirmFocus}
                                       value={confirmPassword}></input>
                            </div>
                            <div className="password-err">
                                {check && password !== confirmPassword ? "密码不一致" : ""}
                            </div>
                        </div>
                        <button className="password-submit">提交</button>
                    </div> : (
                        <div>
                            <div className="forget-logo">
                                <SvgIcon name="forget-password" width={80}></SvgIcon>
                            </div>
                            <form className="forget-verify">
                                <div className="forget-verify-item">
                                    <div className="forget-label">手机号</div>
                                    <input className="forget-input" placeholder={"Phone Number"}></input>
                                </div>
                                <div className="forget-verify-item">
                                    <div className="forget-label">验证码</div>
                                    <div className="flex">
                                        <input className="forget-input" placeholder={"Verify Code"}
                                               style={{width: "230px"}}></input>
                                        <button onClick={getCodeClick} className={"send-code"}>
                                            {isTickerValid ? ticker : "发送"}
                                        </button>
                                    </div>
                                </div>
                                <button type="submit" className="forget-submit" onClick={verifyCode}>验证</button>
                            </form>
                        </div>
                    )
            }

        </div>
    )
}

export default ForgetPassword;