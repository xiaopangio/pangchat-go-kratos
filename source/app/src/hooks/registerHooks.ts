import {RegisterContext} from "@/declare/type";
import {useOutletContext} from "react-router";
import {Register, SendSmsCode, UploadAvatar, VerifySmsCode} from "@/api/user";

export function useRegisterContext() {
    return useOutletContext<RegisterContext>();
}

type RegisterHooks = {
    sendSmsCode: typeof SendSmsCode,
    verifySmsCode: typeof VerifySmsCode,
    uploadAvatar: typeof UploadAvatar,
    register: typeof Register
}

export function useRegister(): RegisterHooks {
    return {
        sendSmsCode: SendSmsCode,
        verifySmsCode: VerifySmsCode,
        uploadAvatar: UploadAvatar,
        register: Register
    };
}