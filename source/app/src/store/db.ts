import Dexie, {Table} from 'dexie';

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
    friend_id: string
    nick_name: string
    note_name: string
    avatar: string
    city_name: string
    province_name: string
    desc: string
    group_name: string
}
export class PangChatDatabase extends Dexie {
    // 'friends' is added by dexie when declaring the stores()
    // We just tell the typing system this is the case
    friendGroups!: Table<FriendGroup>;
    friendRequests!: Table<FriendRequest>;
    images!: Table<image>

    constructor() {
        super('pangchat');
        this.version(1).stores({
            friendGroups: '++,name,uid',
            friendRequests: '++,uid,&request_id',
            images: "&name"
        });
    }
}

export const db = new PangChatDatabase();