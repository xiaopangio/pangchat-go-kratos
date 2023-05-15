import {User} from "@/declare/type";
import {db, FriendRequest, image} from "@/store/db";
import storage from "@/utils/storage";

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
            console.log("存储成功")
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
                reject("")
            }
        }, () => {
            reject("")
        })
    })
}

export async function GetImgList(names: string[]) {
    return new Promise<image[]>((resolve, reject) => {
        db.images.where('name').anyOf(names).toArray().then((res) => {
            if (res) {
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
                res.sort((a, b) => {
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