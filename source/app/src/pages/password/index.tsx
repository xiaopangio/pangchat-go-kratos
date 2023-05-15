import SvgIcon from "@/components/Icon";
import "@/pages/password/index.less"
import {useEffect, useRef, useState} from "react";
import {checkPassword} from "@/utils/check";
import {NavigateToLogin} from "@/components/NavigateToLogin";

type PasswordProps = {
    setPassword: (password: string) => void
    setConfirmPassword: (confirmPassword: string) => void
    password: string
    confirmPassword: string
}

function Password({setPassword, setConfirmPassword, password, confirmPassword}: PasswordProps) {
    // const [registerContext, setRegisterContext] = useRegisterContext()
    const passwdInput = useRef<HTMLInputElement>(null)
    const confirmInput = useRef<HTMLInputElement>(null)
    const [check, setCheck] = useState(false);
    useEffect(() => {
        passwdInput.current?.focus()
    }, []);
    const onPasswdInput = (e: any) => {
        setPassword(e.target.value)
    }
    const onConfirmInput = (e: any) => {
        setConfirmPassword(e.target.value)
    }
    const onConfirmFocus = () => {
        if (confirmPassword.length === 0) {
            setCheck(false);
            return
        }
        setCheck(true);
    }
    useEffect(() => {
        if (confirmPassword.length === 0) {
            setCheck(false);
        } else {
            setCheck(true);
        }
    }, [confirmPassword]);

    return (
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
                           ref={passwdInput}
                           onInput={onPasswdInput}
                           value={password}></input>
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
                           ref={confirmInput}
                           value={confirmPassword}></input>
                </div>
                <div className="password-err">
                    {check && password !== confirmPassword ? "密码不一致" : ""}
                </div>
            </div>
            <NavigateToLogin/>
        </div>
    )
}

export default Password;