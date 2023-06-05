import {User} from "@/declare/type";
import {db, Friend, FriendRequest, image, Message, MessageCountType, ToolOption} from "@/store/db";
import storage from "@/utils/storage";
import {isEmpty} from "fast-glob/out/utils/string";

function SetUser(user: User) {
    storage().set({key: 'user', value: user})
    localStorage.setItem('user', JSON.stringify(user));
}

function GetUser(): User | null {
    let user = storage().get('user') as { value: User };
    if (user) {
        return user.value
    }
    return null
}

export async function StoreImg(name: string, res: any) {
    let reader = new FileReader();
    if (res instanceof Blob || res instanceof File) {
        reader.readAsDataURL(res);
    } else {
        let blob = new Blob([res]);
        reader.readAsDataURL(blob);
    }
    return new Promise<string>((resolve, reject) => {
        reader.onload = function (e) {
            db.images.add({
                blob: e.target?.result as string, name: name
            })
            resolve(e.target?.result as string)
        }
    })
}

export async function GetImg(name: string) {
    return new Promise<string>((resolve, reject) => {
        db.images.get({name: name}).then((res) => {
            if (res) {
                resolve(res.blob)
            } else {
                resolve("")
            }
        }, () => {
            resolve("")
        })
    })
}

export async function DexieGetImgList(names: string[]) {
    return new Promise<image[]>((resolve, reject) => {
        db.images.where('name').anyOf(names).toArray().then((res) => {
            if (res) {
                if (res.length === names.length) {
                    resolve(res)
                } else {
                    let diff = names.filter((value) => {
                        return !res.some((value1) => {
                            return value === value1.name
                        })
                    })
                    reject(diff)
                }
                resolve(res)
            } else {
                reject("没有找到")
            }
        }, () => {
            reject("")
        })
    })
}

export async function GetFriendGroup(uid: string) {
    return new Promise<string[]>((resolve, reject) => {
        db.friendGroups.where('uid').equals(uid).toArray().then((res) => {
            if (res) {
                resolve(res.map((value) => {
                    return value.name
                }))
            } else {
                reject("没有找到")
            }
        }, () => {
            reject("")
        })
    })
}

export async function GetFriendRequests(uid: string) {
    return new Promise<FriendRequest[]>((resolve, reject) => {
        db.friendRequests.where('uid').equals(uid).toArray().then((res) => {
            if (res) {
                res = res.sort((a, b) => {
                    return b.request_id > a.request_id ? 1 : -1
                })
                resolve(res)
            } else {
                reject("没有找到")
            }
        }, () => {
            reject("")
        })
    })
}

export async function UpdateFriendRequestStatus(request_id: string, status: number | string) {
    if (typeof status === "number") {
        status = `${status}`
    }
    return new Promise<string>((resolve, reject) => {
        db.friendRequests.where("request_id").equals(request_id).modify({status: status}).then(() => {
            resolve("更新成功")
        }).catch(() => {
            reject("更新失败")
        })
    })
}

export async function UpdateFriendRequest(request_id: string, friend_request: FriendRequest) {
    return new Promise<string>((resolve, reject) => {
        db.friendRequests.where("request_id").equals(request_id).modify(friend_request).then(() => {
            resolve("更新成功")
        }).catch(() => {
            reject("更新失败")
        })
    })
}

export async function DexieGetFriends(uid: string) {
    return new Promise<Friend[]>((resolve, reject) => {
        db.friends.where('uid').equals(uid).toArray().then((res) => {
            if (res) {
                resolve(res)
            } else {
                resolve([])
            }
        }, () => {
            resolve([])
        })
    })
}

export async function DexieGetFriendsByFriendIds(uid: string, friend_ids: string[]) {
    return new Promise<Friend[]>((resolve, reject) => {
        db.friends.where("[uid+friend_id]").anyOf(friend_ids.map((value) => {
            return [uid, value]
        })).toArray().then((res) => {
            if (res) {
                resolve(res)
            } else {
                reject("没有找到")
            }
        }, () => {
            reject("没有找到")
        })
    })
}

export async function DexieGetFriend(uid: string, friend_id: string) {
    return new Promise<Friend>((resolve, reject) => {
        db.friends.where('[uid+friend_id]').equals([uid, friend_id]).first().then((res) => {
            if (res) {
                resolve(res)
            } else {
                reject("不是好友")
            }
        })
    })
}

export async function UpdateFriend(uid: string, friend_id: string, friend: Friend) {
    return new Promise<string>((resolve, reject) => {
        db.friends.where(['uid', 'friend_id']).equals([uid, friend_id]).modify(friend).then(() => {
            resolve("更新成功")
        }).catch(() => {
            reject("更新失败")
        })
    })
}

export async function AddFriend(uid: string, friend: Friend) {
    friend.uid = uid
    return new Promise<string>((resolve, reject) => {
        db.friends.add(friend).then(() => {
            resolve("添加成功")
        }).catch(() => {
            reject("添加失败")
        })
    })
}

export async function DexieAddFriends(uid: string, friends: Friend[]) {

    return new Promise<string>((resolve, reject) => {
        friends.forEach((value) => {
            value.uid = uid
        })
        db.friends.bulkAdd(friends).then(() => {
            resolve("添加成功")
        }).catch(() => {
            reject("添加失败")
        })
    })
}

export async function DexieUpdateFriendInfo(friend: Friend) {
    return new Promise<string>((resolve, reject) => {
        db.friends.where('friend_id').equals(friend.friend_id).modify(friend).then(() => {
            resolve("更新成功")
        }).catch(() => {
            reject("更新失败")
        })
    })
}

export async function DexieStoreToolOptions(toolOptions: ToolOption[]) {
    return new Promise<string>((resolve, reject) => {
        db.toolOptions.bulkPut(toolOptions).then(() => {
            resolve("更新成功")
        }).catch(() => {
            reject("更新失败")
        })
    })
}

export async function DexieGetToolOptions() {
    return new Promise<ToolOption[]>((resolve, reject) => {
        db.toolOptions.toArray().then((res) => {
            if (res) {
                resolve(res)
            } else {
                reject("没有找到")
            }
        }, () => {
            reject("没有找到")
        })
    })
}

export async function DexieStoreMessages(messages: Message[]) {
    return new Promise<string>((resolve, reject) => {
        db.messages.bulkAdd(messages).then(() => {
            resolve("更新成功")
        }).catch(() => {
            reject("更新失败")
        })
    })
}

export async function DexieStoreMessage(message: Message) {
    return new Promise<string>((resolve, reject) => {
        db.messages.add(message).then(() => {
            resolve("更新成功")
        }).catch((reason) => {
            if (reason.name === "ConstraintError") {
                resolve("已经存在")
            }
        })
    })
}

export async function DexieUpdateMessage(message: Message) {
    return new Promise<string>((resolve, reject) => {
        db.messages.where('[sender_id+receiver_id+send_at]').equals([message.sender_id, message.receiver_id, message.send_at]).modify(message).then(() => {
            resolve("更新成功")
        }, () => {
            reject("更新失败")
        })
    })
}

export async function DexieGetMessagesBefore(uid: string, friend_id: string, message_id: string, limit: number = 20) {
    return new Promise<Message[]>((resolve, reject) => {
        if (isEmpty(message_id)) {
            db.messages.orderBy("message_id").reverse().filter((value) => {
                return value.sender_id === uid && value.receiver_id === friend_id
                    || value.sender_id === friend_id && value.receiver_id === uid
            }).limit(limit).toArray().then((res) => {
                if (res) {
                    resolve(res.reverse())
                } else {
                    reject("没有找到")
                }
            })
        }
        db.messages.orderBy("message_id").reverse().filter((value) => {
            return value.sender_id === uid && value.receiver_id === friend_id && value.message_id < message_id
                || value.sender_id === friend_id && value.receiver_id === uid && value.message_id < message_id
        }).limit(limit).toArray().then((res) => {
            if (res) {
                resolve(res.reverse())
            } else {
                reject("没有找到")
            }
        })
    })
}

export async function DexieGetMessagesByPage(uid: string, friend_id: string, page: number, limit: number = 20) {
    return new Promise<Message[]>((resolve, reject) => {
        db.messages.orderBy("message_id").reverse().filter((message) => {
            return (
                    message.sender_id === uid && message.receiver_id === friend_id)
                || (
                    message.sender_id === friend_id && message.receiver_id === uid)
        }).offset((page - 1) * limit).limit(limit).toArray().then((res) => {
            if (res) {
                resolve(res.reverse())
            } else {
                reject("没有找到")
            }
        })
    })
}

export async function DexieGetMessages(uid: string, friend_id: string) {
    return new Promise<Message[]>((resolve, reject) => {
        db.messages.where('[sender_id+receiver_id]')
            .equals([uid, friend_id])
            .or('[sender_id+receiver_id]').equals([friend_id, uid])
            .toArray().then((res) => {
            if (res) {
                // 根据message_id排序
                res = res.sort((a, b) => {
                    return a.send_at > b.send_at ? 1 : -1
                })
                resolve(res)
            } else {
                reject("没有找到")
            }
        }, () => {
            reject("没有找到")
        })
    })
}

export async function DexieGetMessagesLimit(uid: string, friend_id: string, limit: number) {
    return new Promise<Message[]>((resolve, reject) => {
        db.messages.orderBy("message_id").reverse().or('[sender_id+receiver_id]')
            .equals([uid, friend_id])
            .or('[sender_id+receiver_id]').equals([friend_id, uid])
            .limit(limit).toArray().then((res) => {
            if (res) {
                resolve(res)
            } else {
                reject("没有找到")
            }
        }, () => {
            reject("没有找到")
        })
    })
}

export async function DexieGetMessagesByMessageId(message_ids: string[]) {
    return new Promise<Message[]>((resolve, reject) => {
        db.messages.where('message_id').anyOf(message_ids).toArray().then((res) => {
            if (res) {
                // 根据message_id排序
                res = res.sort((a, b) => {
                    return a.message_id > b.message_id ? 1 : -1
                })
                resolve(res)
            } else {
                reject("没有找到")
            }
        }, () => {
            reject("没有找到")
        })
    })
}

export async function DexieGetMessageByMessageId(message_id: string) {
    return new Promise<Message>((resolve, reject) => {
        db.messages.where('message_id').equals(message_id).first().then((res) => {
            if (res) {
                resolve(res)
            } else {
                reject("没有找到")
            }
        }, () => {
            reject("没有找到")
        })
    })
}

export async function DexieStoreUnreadMessagesCounts(uid: string, m: Map<string, MessageCountType>) {
    return new Promise<string>((resolve, reject) => {
        for (let [, value] of m) {
            //     先尝试添加，如果添加失败，则更新
            db.unreadMessageCounts.add(value).then(() => {
                resolve("添加成功")
            }, () => {
                db.unreadMessageCounts.where('[uid+friend_id]').equals([uid, value.friend_id]).modify(value).then(() => {
                    resolve("更新成功")
                }, () => {
                    reject("更新失败")
                })
            })
        }
    })
}

export async function DexieGetUnreadMessagesCounts(uid: string) {
    return new Promise<Map<string, MessageCountType>>((resolve, reject) => {
        db.unreadMessageCounts.where('uid').equals(uid).toArray().then((res) => {
            if (res) {
                let m = new Map<string, MessageCountType>()
                res.forEach((value) => {
                    m.set(value.friend_id, value)
                })
                resolve(m)
            } else {
                reject("没有找到")
            }
        }, () => {
            reject("没有找到")
        })
    })
}