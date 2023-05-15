import "@/pages/phone/index.less"
import {FormEvent, useEffect, useRef} from "react";
import {useNavigate} from "react-router";

type PhoneProps = {
    phone: string
    accountId: string
    type: number
    setUsername: (username: string) => void
    setType: (type: number) => void
}

function Phone({setUsername, phone, accountId, setType, type}: PhoneProps) {
    const phoneInput = useRef<HTMLInputElement>(null)
    const accountIDInput = useRef<HTMLInputElement>(null)
    let navigate = useNavigate();
    const onInput = (e: FormEvent) => {
        let value = (e.target as HTMLInputElement).value;
        if (type === 1) {
            setUsername(value)
            return
        } else {
            setUsername(value)
            return
        }
    }
    const switchRegisterType = () => {
        if (type === 1) {
            setType(2)
            return
        } else {
            setType(1)
            return
        }
    }
    useEffect(() => {
        phoneInput.current?.focus()
    }, [])
    return (
        <>
            {
                type === 1 ?
                    <div className={"phone"}>
                        <div className="title">
                            <div className="title1">你的电话号码是什么?</div>
                            <div className="title2">我们将给你发送验证码。</div>
                        </div>
                        <div className="phone-input">
                            <div className="label">手机号</div>
                            <input className="input" placeholder={"Phone Number"} ref={phoneInput}
                                   value={phone}
                                   onInput={onInput}></input>
                        </div>
                        <div className="phone-footer">
                            <div
                                className="phone-footer-left"
                                onClick={switchRegisterType}>{type === 1 ? '使用账户名注册?' : '使用手机号注册?'}
                            </div>
                            <div
                                className="phone-footer-right"
                                onClick={() => {
                                    navigate("/login")
                                }}>
                                立即登录
                            </div>
                        </div>
                    </div> :
                    <div className={"phone"}>
                        <div className="title">
                            <div className="title1">你的帐户ID是什么?</div>
                            <div className="title2">帐号ID由字母和数字组成,长度必须超过6位，但不超过16位
                            </div>
                        </div>
                        <div className="phone-input" style={{marginTop: "40px"}}>
                            <div className="label">帐户ID</div>
                            <input className="input" placeholder={"Account ID"} ref={accountIDInput}
                                   value={accountId}
                                   onInput={onInput}></input>
                        </div>
                        <div className="phone-footer">
                            <div
                                className="phone-footer-left"
                                onClick={switchRegisterType}>{type === 1 ? '使用账户名注册?' : '使用手机号注册?'}
                            </div>
                            <div
                                className="phone-footer-right"
                                onClick={() => {
                                    navigate("/login")
                                }}>
                                立即登录
                            </div>
                        </div>
                    </div>
            }
        </>
    );
}

export default Phone;