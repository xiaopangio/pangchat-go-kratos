import {http} from "@/utils/request";
import {
    GetAvatarData,
    LoginData,
    LoginResponse,
    LogoutData,
    LogoutResponse,
    ProfileData,
    ProfileResponse,
    RegisterData,
    RegisterResponse,
    SmsData,
    UploadAvatarResp,
    VerifySmsData
} from "@/api/user/types";

export function SendSmsCode(data: SmsData) {
    return http.post(`/sms`, data);
}

export function VerifySmsCode(data: VerifySmsData) {
    return http.post(`/verify`, data);
}

export function Login(data: LoginData) {
    return http.post<LoginResponse>(`/login`, data);
}

export function Register(data: RegisterData) {
    return http.post<RegisterResponse>(`/register`, data);
}

export function Logout(data: LogoutData) {
    return http.post<LogoutResponse>(`/logout`, data);
}

export function UploadAvatar(data: FormData) {
    return http.post<UploadAvatarResp>(`/avatar`, data);
}

export function UploadUserAvatar(data: FormData) {
    return http.post<UploadAvatarResp>(`/user/avatar`, data);
}

export function GetAvatar(data: GetAvatarData) {
    return http.get(`/user/avatar`, {
        responseType: 'blob',
        params: data
    });
}

export function Profile(data: ProfileData) {
    return http.get<ProfileResponse>(`/user/profile`, {
        params: data
    });
}