import {http} from "@/utils/request";
import {
    GetFriendGroupListRequest,
    GetFriendGroupListResponse,
    GetFriendRequestListRequest,
    GetFriendRequestListResponse,
    SendFriendRequestRequest,
    SendFriendRequestResponse
} from "@/api/friend/types";

const RelationshipPrefix = "/relationship";
const FriendPrefix = RelationshipPrefix + "/friend";
const FriendRequestPrefix = FriendPrefix + "/request";

export function SendFriendRequest(data: SendFriendRequestRequest) {
    return http.post<SendFriendRequestResponse>(FriendRequestPrefix, data);
}

export function GetFriendRequestList(data: GetFriendRequestListRequest) {
    return http.get<GetFriendRequestListResponse>(FriendRequestPrefix + "/list", {
        params: data
    });
}

export function GetFriendGroupList(data: GetFriendGroupListRequest) {
    return http.get<GetFriendGroupListResponse>(FriendPrefix + "/group/list", {
        params: data
    });
}