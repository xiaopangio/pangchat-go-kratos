import "@/pages/register_finish/index.less"
import {useEffect, useRef} from "react";
import SvgIcon from "@/components/Icon";
import storage from "@/utils/storage";
import {dataURLtoBlob} from "@/utils/util";

type RegisterFinishProps = {
    name: string;
    phoneNumber: string;
    accountId: string;
    type: number;
}

function RegisterFinish({name, phoneNumber, accountId, type}: RegisterFinishProps) {
    // const [registerContext, setRegisterContext] = useRegisterContext();
    let avatar = useRef<HTMLImageElement>(null);
    useEffect(() => {
        let userAvatar = storage().get("user_avatar");
        if (userAvatar) {
            let blob = dataURLtoBlob(userAvatar.value);
            let url = (URL || webkitURL).createObjectURL(blob);
            if (avatar.current) {
                avatar.current.src = url
                avatar.current.onload = function () {
                    URL.revokeObjectURL(url)
                }
            }
        }

    }, [])
    return (
        <div className="register-finish">
            <div className="register-display">
                <img className="dis-avatar" ref={avatar} alt="">
                </img>
                <div className="dis-nickname">
                    <div className="dis-tip">NickName:</div>
                    <div className="dis-item">
                        <SvgIcon name="me"></SvgIcon>
                        <span>{name}</span>
                    </div>
                </div>
                {
                    type === 1 ?
                        <div className="dis-phone">
                            <div className="dis-tip">Phone:</div>
                            <div className="dis-item">
                                +86
                                <span>
                            {phoneNumber}
                        </span>
                            </div>
                        </div> :
                        <div className="dis-nickname">
                            <div className="dis-tip">AccountId:</div>
                            <div className="dis-item">
                                <SvgIcon name="account"></SvgIcon>
                                <span>{accountId}</span>
                            </div>
                        </div>
                }
            </div>

        </div>
    )
}

export default RegisterFinish;