import {GetAvatar} from "@/api/user";

interface UserInfoHooks {
    getAvatar: typeof GetAvatar
}

export function useUserInfo(): UserInfoHooks {
    return {
        getAvatar: GetAvatar
    }
}