export type SmsData = {
    phone: string
}

export interface VerifySmsData {
    phone: string
    code: string
}

export interface LoginData {
    type: number
    username: string
    password: string
}

export interface LoginResponse {
    token: string
}

export interface RegisterData {
    type: number
    username: string
    nick_name: string
    avatar_url: string
    password: string
    password_confirm: string
}

export interface RegisterResponse {
    token: string
}

export interface ResetPasswordData {
    uid: string
    password: string
    password_confirm: string
}

export interface ResetPasswordResponse {
    token: string
}

export interface UploadAvatarResp {
    token: string
}

export interface LoginData {
    type: number
    username: string
    password: string
}

export interface LoginResponse {
    token: string
}

export interface LogoutData {
    uid: string
}

export interface LogoutResponse {
}

export interface GetAvatarData {
    avatar_url: string
}

export interface ProfileData {
    account_id: string
}

export interface ProfileResponse {
    user_id: string
    account_id: string
    nick_name: string
    desc: string
    city: string
    province: string
    avatar_url: string
}