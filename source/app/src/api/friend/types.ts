import {Friend, FriendRequest} from "@/store/db";

export interface SendFriendRequestRequest {
    requester_id: string;
    receiver_id: string;
    note_name: string;
    desc: string;
    group_name: string;
}

export interface SendFriendRequestResponse {
    request: FriendRequest;
}

export interface GetFriendRequestListRequest {
    user_id: string;
    page_number: number;
    page_size: number;
}

export interface GetFriendRequestListResponse {
    total: number;
    list: FriendRequest[];
}


export interface GetFriendGroupListRequest {
    user_id: string;
}

export interface GetFriendGroupListResponse {
    group_names: string[];
}

export interface GetFriendListResponse {
    friends: Friend[]
}

export interface GetFriendInfoResponse {
    city_name: string
    province_name: string
    desc: string
    account_id: string
}

export interface UpdateFriendInfoRequest {
    user_id: string
    friend_id: string
    note_name: string
    group_name: string
}

export interface DealFriendRequestData {
    request_id: string,
    status: number,
    note_name: string,
    group_name: string
}