import storage from "@/utils/storage";

export const checkPassword = (password: string) => {
    if (password.length == 0) {
        return true;
    }
    const reg = /^(?=.*\d)(?=.*[a-zA-Z])(?=.*[^\da-zA-Z\s]).{8,16}$/;
    return reg.test(password);
}
export const checkCode = (code: string) => {
    const reg = /^[0-9]{4}$/;
    return reg.test(code);
}
export const checkPhone = (phone: string) => {
    //中国大陆手机号正则
    const reg = /^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$/;
    return reg.test(phone);
}
export const checkNickname = (nickname: string) => {
    return nickname.length !== 0;
}
export const checkToken = () => {
    return storage().get("token");
}
export const checkAccountID = (accountId: string) => {
    // 正则 8-16位字母数字
    const reg = /^[a-zA-Z0-9]{8,16}$/;
    return reg.test(accountId);
}