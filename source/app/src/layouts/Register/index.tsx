import {useNavigate} from "react-router";
import React, {useEffect, useState} from "react";
import SvgIcon from "@/components/Icon";
import "@/layouts/Register/index.less";
import {RegisterPrefix} from "@/declare/const";
import {RegisterParam, User} from "@/declare/type";
import {useRegister} from "@/hooks/registerHooks";
import {checkAccountID, checkCode, checkNickname, checkPassword, checkPhone} from "@/utils/check";
import hashPassword from "@/utils/password";
import Phone from "@/pages/phone";
import VerifyCode from "@/pages/verifyCode";
import Name from "@/pages/name";
import Password from "@/pages/password";
import UploadAvatar from "@/pages/uploadAvatar";
import RegisterFinish from "@/pages/register_finish";
import {dataUrlToFile} from "@/utils/util";
import {ParseJwt} from "@/utils/jwt";
import {GenAvatarName} from "@/utils/gen";
import {GetImg, StoreImg} from "@/utils/store";
import {useRecoilState} from "recoil";
import {RegisterAvatar} from "@/store";
import message from "@/utils/message";


function Register() {
    // 拿到当前的路由
    const [back, setBack] = useState(false);
    const [next, setNext] = useState(false);
    const [step, setStep] = useState(1);
    const [avatarUrl, setAvatarUrl] = useRecoilState(RegisterAvatar);
    const [isUploadAvatar, setIsUploadAvatar] = useState(false);
    const [isRegister, setIsRegister] = useState(false);
    const [registerParam, setRegisterParam] = useState<RegisterParam>({
        type: 1,
        AccountId: "",
        ConfirmPassword: "",
        Code: "",
        Name: "",
        Password: "",
        phoneNumber: "",
        AvatarUrl: "",
        Avatar: ""
    });
    const {sendSmsCode, verifySmsCode, register, uploadAvatar} = useRegister()
    let navigate = useNavigate();
    const setUsername = (username: string) => {
        if (registerParam.type === 1) {
            setRegisterParam({
                ...registerParam,
                phoneNumber: username,
            })
        } else {
            setRegisterParam({
                ...registerParam,
                AccountId: username,
            })
        }
    }
    const setType = (type: number) => {
        setRegisterParam({
            ...registerParam,
            type: type,
        })
    }
    const setCode = (code: string) => {
        setRegisterParam({
            ...registerParam,
            Code: code,
        })
    }
    const setName = (name: string) => {
        setRegisterParam({
            ...registerParam,
            Name: name,
        })
    }
    const setPassword = (password: string) => {
        setRegisterParam({
            ...registerParam,
            Password: password,
        })
    }
    const setConfirmPassword = (confirmPassword: string) => {
        setRegisterParam({
            ...registerParam,
            ConfirmPassword: confirmPassword,
        })
    }
    const saveAvatarDone = async (file: File | Blob) => {
        let avatar = GenAvatarName();
        try {
            await StoreImg(avatar, file)
        } catch (e) {
            console.log(e)
            return
        }
        setAvatarUrl(avatar);
        setIsUploadAvatar(true);
    }
    useEffect(() => {
        if (step === 1) {
            setBack(false);
        } else {
            setBack(true);
        }
        if (isUploadAvatar) {
            setNext(true);
        }
        setNext(false)
        navigate(RegisterPrefix + step)
    }, [step]);
    useEffect(() => {
        if (checkPhone(registerParam.phoneNumber)) {
            setNext(true);
        } else {
            setNext(false);
        }
    }, [registerParam.phoneNumber])
    useEffect(() => {
        if (checkAccountID(registerParam.AccountId)) {
            setNext(true);
        } else {
            setNext(false);
        }
    }, [registerParam.AccountId]);
    useEffect(() => {
        if (checkCode(registerParam.Code)) {
            setNext(true);
        } else {
            setNext(false);
        }
    }, [registerParam.Code]);
    useEffect(() => {
        if (checkNickname(registerParam.Name)) {
            setNext(true);
        } else {
            setNext(false);
        }
    }, [registerParam.Name])
    useEffect(() => {
        if (registerParam.Password === registerParam.ConfirmPassword && registerParam.Password != "" && checkPassword(registerParam.Password)) {
            setNext(true);
        } else {
            setNext(false);
        }
    }, [registerParam.Password, registerParam.ConfirmPassword]);
    useEffect(() => {
        if (registerParam.AvatarUrl !== "") {
            setNext(true);
        } else {
            setNext(false);
        }
    }, [registerParam.AvatarUrl]);
    useEffect(() => {
        if (registerParam.type === 1) {
            if (checkPhone(registerParam.phoneNumber)) {
                setNext(true);
            } else {
                setNext(false);
            }
        } else {
            console.log(registerParam.AccountId)
            if (checkAccountID(registerParam.AccountId)) {
                console.log("checkAccountID")
                setNext(true);
            } else {
                setNext(false);
            }
        }
    }, [registerParam.type]);
    useEffect(() => {
        if (isUploadAvatar) {
            setNext(true);
        } else {
            setNext(false);
        }
    }, [isUploadAvatar]);

    useEffect(() => {
        setIsRegister(false);

        switch (step) {
            case 1:
                if (checkPhone(registerParam.phoneNumber) || checkAccountID(registerParam.AccountId)) {
                    setNext(true);
                } else {
                    setNext(false);
                }
                break
            case 2:
                break
            case 3:
                if (checkNickname(registerParam.Name)) {
                    setNext(true);
                } else {
                    setNext(false);
                }
                break
            case 4:
                if (registerParam.Password === registerParam.ConfirmPassword && registerParam.Password != "" && checkPassword(registerParam.Password)) {
                    setNext(true);
                } else {
                    setNext(false);
                }
                break
            case 5:
                if (isUploadAvatar) {
                    setNext(true);
                } else {
                    setNext(false);
                }
                break
            case 6:
                setIsRegister(true);
                setNext(true);
                break
        }
    }, [step]);
    const nextStep = async () => {
        if (!next) return;
        if (step === 7) {
            return;
        }
        if (step < 7) {
            if (step === 1 && registerParam.type === 2) {
                setStep(step + 2);
                navigate(RegisterPrefix + (step + 2));
                return;
            }
            switch (step) {
                case 1:
                    console.log(registerParam.type)
                    if (registerParam.type === 1) {
                        try {
                            await sendSmsCode({
                                phone: registerParam.phoneNumber,
                            });
                            setStep(step + 1);
                            navigate(RegisterPrefix + (step + 1));
                        } catch (e) {
                            console.log(e)
                        }
                    }
                    break
                case 2:
                    try {
                        await verifySmsCode({phone: registerParam.phoneNumber, code: registerParam.Code});
                        setStep(step + 1);
                        navigate(RegisterPrefix + (step + 1));
                    } catch (e) {
                        console.log(e)
                    }
                    break
                case 6:
                    try {
                        let hp = hashPassword(registerParam.Password);
                        let hpc = hashPassword(registerParam.ConfirmPassword);
                        let userAvatar = await GetImg(avatarUrl);
                        if (!userAvatar) {
                            console.log("no avatar")
                        }
                        let file = dataUrlToFile(userAvatar, avatarUrl)
                        let formData = new FormData();
                        formData.append("file", file);
                        let res = await uploadAvatar(formData);
                        let user = ParseJwt<User>(res.token);
                        console.log(user)
                        await register({
                            avatar_url: user.avatar,
                            nick_name: registerParam.Name,
                            password: hp as string,
                            password_confirm: hpc as string,
                            type: registerParam.type,
                            username: registerParam.type === 1 ? registerParam.phoneNumber : registerParam.AccountId,
                        });
                        message.success({
                            content: "注册成功", duration: 2000
                        });
                        navigate("/login");
                    } catch (e) {
                        console.log(e)
                    }
                    break
                default:
                    setStep(step + 1);
                    navigate(RegisterPrefix + (step + 1));
                    break
            }
        }

    }
    const backStep = () => {
        if (step > 1) {
            if (step === 3 && registerParam.type === 2) {
                setStep(step - 2);
                return;
            }
            setStep(step - 1);
            navigate(RegisterPrefix + (step - 1));
        }
    }

    return (

        <div className={"register"}>
            {
                back ?
                    <button className={"btn-back"} onClick={backStep}>
                        <SvgIcon name={"back"} width={24}></SvgIcon>
                    </button>
                    : null
            }
            <div className={"content"}>
                {
                    step === 1 ?
                        <Phone type={registerParam.type} setUsername={setUsername} phone={registerParam.phoneNumber}
                               accountId={registerParam.AccountId} setType={setType}/> : null
                }
                {
                    step === 2 ? <VerifyCode phoneNumber={registerParam.phoneNumber} setCode={setCode}/> : null
                }
                {
                    step === 3 ? <Name setName={setName} name={registerParam.Name}/> : null
                }
                {
                    step === 4 ? <Password setPassword={setPassword} setConfirmPassword={setConfirmPassword}
                                           confirmPassword={registerParam.ConfirmPassword}
                                           password={registerParam.Password}/> : null
                }
                {
                    step === 5 ? <UploadAvatar saveAvatarDo={saveAvatarDone} isUploadAvatar={isUploadAvatar}
                    /> : null
                }
                {
                    step === 6 ? <RegisterFinish type={registerParam.type} name={registerParam.Name}
                                                 accountId={registerParam.AccountId}
                                                 phoneNumber={registerParam.phoneNumber}/> : null
                }
            </div>


            <div className={`btn-next ${next ? 'btn-next-active' : ''}`}
                 onClick={nextStep}>{isRegister ? "注册" : "下一步"}</div>

        </div>
    )
}


export default Register;