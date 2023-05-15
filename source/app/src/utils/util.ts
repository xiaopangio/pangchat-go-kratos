import storage from "@/utils/storage";
import {RefObject} from "react";
import {checkToken} from "@/utils/check";
import {User} from "@/declare/type";
import {ParseJwt} from "@/utils/jwt";

export function dataURLtoBlob(dataurl: string, name?: string) {
    let arr = dataurl.split(',');
    let mime = arr[0].match(/:(.*?);/)![1];
    let bstr = atob(arr[1]);
    let n = bstr.length;
    let u8arr = new Uint8Array(n);
    while (n--) {
        u8arr[n] = bstr.charCodeAt(n);
    }
    return new Blob([u8arr], {type: mime});
}

export function fileToBlob(file: File): Promise<Blob> {
    return new Promise<Blob>((resolve, reject) => {
        let reader = new FileReader();
        reader.onload = function (e) {
            if (e.target) {
                resolve(dataURLtoBlob(e.target.result as string));
            }
        };
        reader.readAsDataURL(file);
    });
}

export function blobToFile(blob: Blob, name?: string): File {
    return new File([blob], name ? name : blob.name, {type: blob.type});
}

export const showImg = (blob: Blob, ref: RefObject<HTMLImageElement>) => {
    let url = (URL || webkitURL).createObjectURL(blob);
    if (ref.current) {
        ref.current.src = url
        ref.current.onload = function () {
            URL.revokeObjectURL(url)
        }
    }
}

export function dataUrlToFile(dataurl: string, filename: string) {
    let arr = dataurl.split(',');
    let mime = arr[0].match(/:(.*?);/)![1];
    let bstr = atob(arr[1]);
    let n = bstr.length;
    let u8arr = new Uint8Array(n);
    while (n--) {
        u8arr[n] = bstr.charCodeAt(n);
    }
    return new File([u8arr], filename, {type: mime});
}

export function storeBlob(key: string, blob: Blob) {
    let reader = new FileReader();
    reader.readAsDataURL(blob);
    return new Promise((resolve, reject) => {
        reader.onload = function (e) {
            storage().set({
                key: key,
                value: e.target?.result,
                expire: "30d"
            })
            resolve("ok");
        }
    })
}

export function RefreshCurrentUser() {
    let token = checkToken();
    let user: User | null = null;
    if (token) {
        user = ParseJwt<User>(token.value);
        return user;
    }
    return null;
}