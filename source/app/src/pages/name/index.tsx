import SvgIcon from "@/components/Icon";
import "@/pages/name/index.less"
import {useEffect, useRef} from "react";
import {useNavigate} from "react-router";
import {NavigateToLogin} from "@/components/NavigateToLogin";

type NameProps = {
    name: string
    setName: (name: string) => void
}

function Name({name, setName}: NameProps) {
    // const [registerContext, setRegisterContext] = useRegisterContext()
    const nameInput = useRef<HTMLInputElement>(null);
    let navigate = useNavigate();
    const onInput = (e: any) => {
        setName(e.target.value);
    }
    useEffect(() => {
        nameInput.current?.focus();
    }, []);
    return (
        <div className={"nickname"}>
            <div className="nickname-label">
                你的昵称是什么?
            </div>
            <div className="nickname-tip">
                写下你的昵称。你可以改变它在设置中。
            </div>
            <div className="nickname-input">
                <div className="nickname-input-label">昵称</div>
                <div className="nickname-input-box">
                    <SvgIcon name={"me"} color={"#cccccc"}/>
                    <input className="nickname-input-input" ref={nameInput} value={name}
                           placeholder={"Nickname"}
                           onInput={onInput}></input>
                </div>
            </div>
            <NavigateToLogin/>
        </div>
    )
}

export default Name;