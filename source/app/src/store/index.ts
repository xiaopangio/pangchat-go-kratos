import {atom, selector} from "recoil";
import {User} from "@/declare/type";
import {ProfileResponse} from "@/api/user/types";
import {Friend, FriendRequest, MessageCountType} from "@/store/db";

export const currentUserState = atom<User | null>({
    key: 'currentUserState',
    default: null
})
export const loginState = atom({
    key: 'loginState',
    default: false
})
export const SearchUserState = atom<ProfileResponse | null>({
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
export const UnreadFriendTextCount = atom<number>({
    key: 'UnreadFriendTextCount',
    default: 0
})
export const UnreadMessageCountMap = atom<Map<string, MessageCountType>>({
    key: 'UnreadMessageCountMap',
    default: new Map<string, MessageCountType>()
})
export const CurrentDealFriendRequest = atom<FriendRequest | null>(
    {
        key: "CurrentDealFriendRequest",
        default: null
    }
)
export const UnreadMessageCount = selector({
    key: "UnreadMessageCount",
    get: ({get}) => {
        let count = 0
        get(UnreadMessageCountMap).forEach((value) => {
            count += value.unread_count
        })
        return count
    }
})
export const RegisterAvatar = atom<string>({
    key: "RegisterAvatar",
    default: ""
})
export const currentDialogState = atom<string>({
    key: "currentDialogState",
    default: ""
})
export const AvatarMap = atom<Map<string, string>>({
    key: "AvatarMap",
    default: new Map<string, string>()
})
export const FriendsMap = atom<Map<string, Friend>>({
    key: "FriendsMap",
    default: new Map<string, Friend>()
})
export const DataSetFriend = atom<Friend | null>({
    key: "ProfileSettingInfo",
    default: null
})