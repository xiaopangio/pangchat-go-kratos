import Dexie, {Table} from 'dexie';
import {MessageType} from "@/declare/const";

export interface FriendGroup {
    uid: string;
    name: string;
}

export interface FriendRequest {
    uid: string;
    request_id: string;
    requester_id: string;
    receiver_id: string;
    desc: string;
    status: string;
    create_time: string;
    update_time: string;
    nick_name: string;
    avatar: string;
}

export interface image {
    name: string
    blob: string
}

export interface Friend {
    uid: string
    friend_id: string
    account_id: string
    nick_name: string
    note_name: string
    avatar: string
    city_name: string
    province_name: string
    desc: string
    group_name: string
}

export interface ToolOption {
    name: string
    icon: string
}

// 定义type枚举
export interface Message {
    message_id: string
    type: MessageType
    content: string
    send_at: string
    sender_id: string
    receiver_id: string
    friend_id?: string
    is_need_load_before?: boolean
    need_load_before_count?: number
}

export interface MessageCountType {
    uid: string
    friend_id: string
    unread_count: number
    message_id: string
}

export interface UnreadMessageInfo {
    latest_message: Message
    unread_count: number
}

export class PangChatDatabase extends Dexie {
    // 'friends' is added by dexie when declaring the stores()
    // We just tell the typing system this is the case
    friendGroups!: Table<FriendGroup>;
    friendRequests!: Table<FriendRequest>;
    images!: Table<image>
    friends!: Table<Friend>;
    toolOptions!: Table<ToolOption>
    messages!: Table<Message>
    unreadMessageCounts!: Table<MessageCountType>

    constructor() {
        super('pangchat');
        this.version(3).stores({
            friendGroups: '++,name,uid,[uid+name]',
            friendRequests: '++,uid,&request_id',
            images: "&name",
            friends: '++,&[uid+friend_id],uid,friend_id,[uid+group_name]',
            toolOptions: '++',
            messages: '++,[sender_id+receiver_id],[sender_id+receiver_id+send_at],&message_id',
            unreadMessageCounts: '++,&[uid+friend_id],uid,friend_id'
        });
    }
}

export const db = new PangChatDatabase();