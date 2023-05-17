import "@/pages/register_finish/index.less"
import {useEffect, useState} from "react";
import SvgIcon from "@/components/Icon";
import {useRecoilState} from "recoil";
import {RegisterAvatar} from "@/store";
import {GetImg} from "@/utils/store";

type RegisterFinishProps = {
    name: string;
    phoneNumber: string;
    accountId: string;
    type: number;
}

function RegisterFinish({name, phoneNumber, accountId, type}: RegisterFinishProps) {
    // const [registerContext, setRegisterContext] = useRegisterContext();
    const [avatarUrl] = useRecoilState(RegisterAvatar);
    const [avatarData, setAvatarData] = useState("");
    useEffect(() => {
        if (avatarUrl) {
            GetImg(avatarUrl).then((res) => {
                setAvatarData(res)
            })
        }
    }, [])
    return (
        <div className="register-finish">
            <div className="register-display">
                <img className="dis-avatar" src={avatarData} alt="">
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