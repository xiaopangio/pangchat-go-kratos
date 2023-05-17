import {atom} from "recoil";
import {User} from "@/declare/type";
import {ProfileResponse} from "@/api/user/types";
import {FriendRequest} from "@/store/db";

export const currentUserState = atom<User | null>({
    key: 'currentUserState',
    default: null
})
export const loginState = atom({
    key: 'loginState',
    default: false
})
export const AddUserState = atom<ProfileResponse | null>({
    key: 'AddUserState',
    default: null
})
export const AddUserGroupState = atom<string>({
    key: 'AddUserGroupState',
    default: ""
})
export const ConnectionUrlState = atom<string>({
    key: 'ConnectionUrlState',
    default: ""
})
export const ConnectionState = atom<boolean>({
    key: 'ConnectionState',
    default: false
})
export const GroupNames = atom<string[]>({
    key: 'GroupNames',
    default: []
})
export const GroupName = atom<string>({
    key: 'GroupName',
    default: ""
})
export const UnreadFriendRequestCount = atom<number>({
    key: 'UnreadFriendRequestCount',
    default: 0
})
export const UnreadMessageCount = atom<number>({
    key: 'UnreadMessageCount',
    default: 0
})
export const UnreadFriendTextCount = atom<number>({
    key: 'UnreadFriendTextCount',
    default: 0
})
export const UnreadMessageCountMap = atom<Map<string, number>>({
    key: 'UnreadMessageCountMap',
    default: new Map<string, number>()
})
export const CurrentDealFriendRequest = atom<FriendRequest | null>(
    {
        key: "CurrentDealFriendRequest",
        default: null
    }
)
export const RegisterAvatar = atom<string>({
    key: "RegisterAvatar",
    default: ""
})