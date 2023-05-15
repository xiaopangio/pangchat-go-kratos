import "@/pages/verifyCode/index.less"
import {FocusEvent, KeyboardEvent, RefObject, useEffect, useRef, useState} from "react";
import {useInterval} from "@/hooks/intervalHooks";
import {useRegister} from "@/hooks/registerHooks";

type VerifyCodeProps = {
    phoneNumber: string,
    setCode: (code: string) => void
}

function VerifyCode({setCode, phoneNumber}: VerifyCodeProps) {
    const [ticker, setTicker] = useState(60);
    const [currentInput, setCurrentInput] = useState(0);
    const input1 = useRef<HTMLInputElement>(null);
    const input2 = useRef<HTMLInputElement>(null);
    const input3 = useRef<HTMLInputElement>(null);
    const input4 = useRef<HTMLInputElement>(null);
    const {sendSmsCode} = useRegister();
    const inputMap = new Map<number, RefObject<HTMLInputElement>>([
        [1, input1],
        [2, input2],
        [3, input3],
        [4, input4]
    ])
    const [reset, clear] = useInterval(() => {
        setTicker(ticker - 1);
    })
    useEffect(() => {
        if (ticker <= 0) {
            clear();
        }
    }, [ticker]);
    const resendCode = async () => {
        if (ticker > 0) return;
        try {
            await sendSmsCode({phone: phoneNumber});
            setTicker(60);
            reset()
        } catch (e) {
            console.log(e)
        }
    }
    const nextInput = () => {
        switch (currentInput) {
            case 0:
                setCurrentInput(currentInput + 1);
                input1.current?.focus();
                break;
            case 1:
                setCurrentInput(currentInput + 1);
                input2.current?.focus();
                break;
            case 2:
                setCurrentInput(currentInput + 1);
                input3.current?.focus();
                break;
            case 3:
                setCurrentInput(currentInput + 1);
                input4.current?.focus();
                break;
        }
    }
    const handlerInput = (e: KeyboardEvent<HTMLInputElement>) => {
        switch (e.key) {
            case "0":
                nextInput()
                break
            case "1":
                nextInput()
                break
            case "2":
                nextInput()
                break
            case "3":
                nextInput()
                break
            case "4":
                nextInput()
                break
            case "5":
                nextInput()
                break
            case "6":
                nextInput()
                break
            case "7":
                nextInput()
                break
            case "8":
                nextInput()
                break
            case "9":
                nextInput()
                break
            case "Backspace":
                switch (currentInput) {
                    case 0:
                        if (input1.current?.value.length !== 0) {
                            changeCode()
                            return;
                        }
                        input1.current?.focus();
                        break;
                    case 1:
                        if (input2.current?.value.length !== 0) {
                            changeCode()
                            return;
                        }
                        setCurrentInput(currentInput - 1);
                        input1.current?.focus();
                        break;
                    case 2:
                        if (input3.current?.value.length !== 0) {
                            changeCode()
                            return;
                        }
                        setCurrentInput(currentInput - 1);
                        input2.current?.focus();
                        break;
                    case 3:
                        if (input4.current?.value.length !== 0) {
                            changeCode()
                            return;
                        }
                        setCurrentInput(currentInput - 1);
                        input3.current?.focus();
                        break;
                    case 4:
                        setCurrentInput(currentInput - 1);
                        input4.current?.focus();
                        break;
                    default:
                        break;
                }
                if (currentInput === 0) return;
                break;
            case "Tab":
                setCurrentInput(currentInput + 1);
                break;
        }
    }
    const handleFocus = (e: FocusEvent) => {
        //     拿到当前input的key值
        inputMap.forEach((value, key) => {
            if (value.current === e.target) {
                if (key - 1 === currentInput) return;
                setCurrentInput(key - 1);
            }
        })
    }
    const goNext = () => {
        console.log("go next")
    }
    useEffect(() => {
        input1.current?.focus();
    }, []);
    const changeCode = () => {
        const code = input1.current?.value === undefined ? "" : input1.current?.value + input2.current?.value + input3.current?.value + input4.current?.value;
        setCode(code);
    }
    useEffect(() => {
        changeCode()
    }, [currentInput, input1.current?.value, input2.current?.value, input3.current?.value, input4.current?.value])

    return (
        <div className={"verify-code"}>
            <div className="verify-title">
                <div className="verify-label">
                    验证码
                </div>
                <div className="verify-tip">
                    <span>输入我们发送的的验证码 <span style={{color: "rgb(8, 28, 44)"}}>+86 </span></span>
                    <span className="verify-phone">19944162167</span>
                </div>
            </div>
            <div className="verify-input">
                {Array.from({length: 4}, (v, k) => k + 1).map((item, index) => {
                    return <input type="text" maxLength={1} ref={inputMap.get(item)} key={item} onKeyDown={handlerInput}
                                  onFocus={handleFocus}/>
                })}
            </div>
            <div className="resend-code">
                <div className="resend-tip">如果你没有得到代码,重新发送它 <span
                    className={"resend-ticker"}>{ticker}</span> seconds.
                </div>
                <div className="resend-btn" onClick={resendCode}>重新发送</div>
            </div>
        </div>
    );
}

export default VerifyCode;