type SetFunc = (newValue: any) => void
type User = {
    uid: string,
    nick_name: string,
    account_id: string,
    personal_desc: string,
    avatar: string,
    avatar_url: string,
    phone: string,
}
type RegisterParam = {
    type: number;
    AccountId: string;
    phoneNumber: string;
    Code: string;
    Name: string;
    Password: string;
    ConfirmPassword: string;
    AvatarUrl: string;
    Avatar: string
}
type setRegisterFunc = (param: RegisterParam) => void
type RegisterContext = [
    register: RegisterParam,
    setRegister: setRegisterFunc,
]
type useSendCodeReturn = {
    getCodeClick?: () => void,
    isTickerValid?: boolean,
    ticker?: number,
}

interface Provicne {
    province_id: string,
    province_name: string,
    cities: City[]
}

interface City {
    city_id: string,
    city_name: string,
}

export {
    SetFunc,
    User,
    RegisterParam,
    RegisterContext,
    useSendCodeReturn,
    Provicne,
    City,
}