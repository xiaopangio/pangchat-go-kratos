import "@/pages/uploadAvatar/index.less"
import React, {ChangeEvent, useEffect, useRef, useState} from "react";
import avatar_example_1 from "@img/avatar-example-1.png"
import avatar_example_2 from "@img/avatar-example-2.png"
import avatar_example_3 from "@img/avatar-example-3.png"
import avatar_example_4 from "@img/avatar-example-4.png"
import successPng from "@img/success.png"
import loadingPng from "@img/loading.png"
import CameraPhoto from "@img/Camera.png"
import AvatarDefault from "@img/avatar_default.png"
import AvatarBtn from "@img/avatar_btn.png"
import SvgIcon from "@/components/Icon";
import TakePhoto from "@/components/TakePhoto";
import storage from "@/utils/storage";
import {dataURLtoBlob, fileToBlob} from "@/utils/util";

type UploadAvatarProps = {
    saveAvatarDo: (file: File | Blob) => void
    isUploadAvatar: boolean
}

function UploadAvatar({saveAvatarDo, isUploadAvatar}: UploadAvatarProps) {
    const [imgList, setImgList] = useState<string[]>([]);
    const [avatarTitle, setAvatarTitle] = useState("");
    const [uploadClassString, setUploadClassString] = useState("avatar-select avatar-select-none");
    const [coverClassString, setCoverClassString] = useState("cover");
    const [isTakePhoto, setIsTakePhoto] = useState(false);
    const uploadInputRef = useRef<HTMLInputElement>(null);
    const avatarImgRef = useRef<HTMLImageElement>(null);
    const avatarImg2Ref = useRef<HTMLImageElement>(null);
    const [isClose, setIsClose] = useState(false);
    const [avatarFile, setAvatarFile] = useState<File | Blob>();
    const initImgList = () => {
        let list: string[] = [];
        for (let i = 1; i < 5; i++) {
            list.push(String(i));
        }
        setImgList(list);
    }
    const handleUpload = () => {
        setUploadClassString("avatar-select  avatar-select-show");
        setCoverClassString("cover cover-active")
    }
    const closeMask = () => {
        setCoverClassString("cover cover-hide")
        setUploadClassString("avatar-select  avatar-select-none  avatar-select-hide")
    }
    const uploadInputClick = () => {
        if (uploadInputRef.current) {
            uploadInputRef.current?.click()
        }
    }
    const takePhotoClick = () => {
        setIsTakePhoto(true);
    }
    const saveAvatar = async (e: ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];
        if (file) {
            setAvatarTitle("Wait a second, your photo still uploading")
            if (avatarImgRef.current) {
                avatarImgRef.current.src = loadingPng
            }
            if (avatarImg2Ref.current) {
                avatarImg2Ref.current.src = ""
            }
            closeMask()
            setIsClose(false);
            try {
                const blob = await fileToBlob(file)
                showImg(blob)
                if (avatarImg2Ref.current) {
                    avatarImg2Ref.current.src = successPng
                }
                setAvatarTitle("Done! Your photo successfully uploaded")
            } catch (e) {
                console.log(e)
            }
            if (uploadInputRef.current) {
                uploadInputRef.current.value = ""
            }
            setAvatarFile(file)
            saveAvatarDo(file)
        }
    }
    const showImg = (blob: Blob) => {
        let url = (URL || webkitURL).createObjectURL(blob);
        if (avatarImgRef.current) {
            avatarImgRef.current.src = url
            avatarImgRef.current.onload = function () {
                URL.revokeObjectURL(url)
            }
        }
    }

    useEffect(() => {
        console.log("isUploadAvatar", isUploadAvatar)
        initImgList()
        if (isUploadAvatar) {
            let userAvatar = storage().get("user_avatar");
            if (userAvatar) {
                let blob = dataURLtoBlob(userAvatar.value);
                showImg(blob);
            }
        }
    }, []);
    const savePhoto = (img: string) => {
        setAvatarTitle("等一下,你的头像正在上传")
        if (avatarImgRef.current) {
            avatarImgRef.current.src = loadingPng
        }
        if (avatarImg2Ref.current) {
            avatarImg2Ref.current.src = ""
        }
        const blob = dataURLtoBlob(img, "avatar.png")
        setAvatarFile(blob)
        if (avatarImg2Ref.current) {
            avatarImg2Ref.current.src = successPng
        }
        setAvatarTitle("完成! 你的头像已经上传成功")
        setIsClose(false);
        saveAvatarDo(blob)
    }
    useEffect(() => {
        if (avatarFile && avatarFile instanceof Blob) {
            showImg(avatarFile)
        }
    }, [avatarFile]);

    const takeCancel = () => {
        setIsTakePhoto(false);
        closeMask()
    }
    return (
        <div>
            {
                isTakePhoto ? <div className="avatar"><TakePhoto savePhoto={savePhoto} takeCancel={takeCancel}/></div> :
                    <div className="avatar">
                        <div className="avatar-tip">上传你的头像</div>
                        <div className="avatar-upload">
                            <img className="avatar-img" ref={avatarImgRef} alt="" src={AvatarDefault}></img>
                            <img className="avatar-img2" ref={avatarImg2Ref} alt="" src={AvatarBtn}></img>
                        </div>
                        <div className="avatar-title">
                            {avatarTitle}
                        </div>
                        {isClose ? null :
                            <button className="avatar-upload-btn"
                                    onClick={handleUpload}>上传 </button>}
                        <div className={coverClassString} onClick={closeMask}></div>
                        <div className={uploadClassString}>
                            <div className="avatar-list">
                                <img className="avatar-item" src={CameraPhoto} alt={"picture"}
                                     onClick={takePhotoClick}></img>
                                <img className="avatar-item" alt={""} src={avatar_example_1}></img>
                                <img className="avatar-item" alt={""} src={avatar_example_2}></img>
                                <img className="avatar-item" alt={""} src={avatar_example_3}></img>
                                <img className="avatar-item" alt={""} src={avatar_example_4}></img>
                            </div>
                            <div className="take-photo">
                                <div className="icon">
                                    <SvgIcon name="take-photo" width={20}></SvgIcon>
                                </div>
                                <span className="tip" onClick={takePhotoClick}>Take Photo</span>
                            </div>
                            <div className="take-photo">
                                <div className="icon">
                                    <SvgIcon name="photo" width={20}></SvgIcon>
                                </div>
                                <input type="file" hidden accept="image/png,image/jpg" onChange={saveAvatar}
                                       ref={uploadInputRef}/>
                                <span className="tip" onClick={uploadInputClick}>Choose From Library</span>
                            </div>
                        </div>
                    </div>
            }
        </div>
    )
}

export default UploadAvatar