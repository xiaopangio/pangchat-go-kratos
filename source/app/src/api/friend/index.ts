import {http} from "@/utils/request";
import {
    DealFriendRequestData,
    GetFriendGroupListRequest,
    GetFriendGroupListResponse, GetFriendInfoResponse, GetFriendListResponse,
    GetFriendRequestListRequest,
    GetFriendRequestListResponse,
    SendFriendRequestRequest,
    SendFriendRequestResponse, UpdateFriendInfoRequest
} from "@/api/friend/types";

const RelationshipPrefix = "/relationship";
const FriendPrefix = RelationshipPrefix + "/friend";
const FriendRequestPrefix = FriendPrefix + "/request";
const FriendGroupPrefix = FriendPrefix + "/group";
export function SendFriendRequest(data: SendFriendRequestRequest) {
    return http.post<SendFriendRequestResponse>(FriendRequestPrefix, data);
}

export function GetFriendRequestList(data: GetFriendRequestListRequest) {
    return http.get<GetFriendRequestListResponse>(FriendRequestPrefix + "/list", {
        params: data
    });
}

export function GetFriendGroupList(data: GetFriendGroupListRequest) {
    return http.get<GetFriendGroupListResponse>(FriendGroupPrefix+"/list", {
        params: data
    });
}
export function DealFriendRequestApi(data:DealFriendRequestData) {
    return http.put(FriendRequestPrefix ,data);
}
export function GetFriendList(id: string){
    return http.get<GetFriendListResponse>(FriendPrefix + "/list",{
        params: {
            user_id:id
        }
    });
}
export function DeleteFriend(id: string, friend_id: string){
    return http.delete(FriendPrefix ,{
        params: {
            user_id:id,
            friend_id: friend_id
        }
    });
}
export function GetFriendInfo(friend_id: string){
    return http.get<GetFriendInfoResponse>(FriendPrefix,{
        params: {
            friend_id: friend_id
        }
    });
}
export function UpdateFriendInfo(data: UpdateFriendInfoRequest){
    return http.put(FriendPrefix ,data);
}
export function CreateFriendGroup(user_id: string, group_name: string){
    return http.post(FriendGroupPrefix ,{
        user_id: user_id,
        group_name: group_name
    });
}
export function UpdateFriendGroup(user_id: string, group_name: string, new_group_name: string){
    return http.put(FriendGroupPrefix ,{
        user_id: user_id,
        group_name: group_name,
        new_group_name: new_group_name
    });
}
export function DeleteFriendGroup(user_id: string, group_name: string){
    return http.delete(FriendGroupPrefix ,{
        params: {
            user_id: user_id,
            group_name: group_name
        }
    });
}