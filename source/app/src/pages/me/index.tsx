import "@/pages/me/index.less"
import {useRecoilState} from "recoil";
import {currentUserState} from "@/store";
import {useEffect, useState} from "react";
import {blobToFile, RefreshCurrentUser} from "@/utils/util";
import SvgIcon from "@/components/Icon";
import {GetAvatar, UploadUserAvatar} from "@/api/user";
import UploadAvatar from "../uploadAvatar";
import {useNavigate} from "react-router";
import {SettingPrefix} from "@/declare/const";
import {GetImg, StoreImg} from "@/utils/store";
import DefaultImg from "@img/default.jpg";
import {GenAvatarName} from "@/utils/gen";
import SettingHeader from "@/components/SettingHeader";
import {isNull} from "lodash";


function Me() {
    const [currentUser, setCurrentUser] = useRecoilState(currentUserState)
    const [avatarData, setAvatarData] = useState(DefaultImg);
    const [isTakePhoto, setIsTakePhoto] = useState(false);
    useEffect(() => {
        //检查当前用户信息
        if (currentUser) {
            return
        } else {
            // 重新从localstorage中获取用户信息，设置到recoil中
            let user = RefreshCurrentUser();
            setCurrentUser(user)
        }
    }, []);
    const initAvatar = async () => {
        if (isNull(currentUser)) {
            return
        }
        let res = await GetImg(currentUser.avatar)
        // 设置头像
        setAvatarData(res)
        // 如果头像已经缓存，就不再请求
        if (res !== "") {
            return
        }
        try {
            res = await GetAvatar({avatar_url: currentUser.avatar})
            res = await StoreImg(currentUser.avatar, res)
            // 设置头像
            setAvatarData(res)
        } catch (e) {
            console.log(e)
        }
    }
    useEffect(() => {
        //检查当前用户信息
        if (isNull(currentUser)) {
            return
        }
        initAvatar().then()

    }, [currentUser?.avatar]);
    const toTakePhoto = () => {
        setIsTakePhoto(true);
    }
    const uploadAvatarFile = async (file: File | Blob) => {
        if (!currentUser) {
            return
        }
        let avatar = GenAvatarName();
        try {
            await StoreImg(avatar, file)
        } catch (e) {
            console.log(e)
            return
        }
        if (file instanceof Blob) {
            try {
                const f = blobToFile(file, avatar)
                let formData = new FormData();
                formData.append("file", f);
                await UploadUserAvatar(formData);
                setIsTakePhoto(false)
            } catch (e) {
                console.log(e)
            }
        } else if (file as any instanceof File) {
            let formData = new FormData();
            formData.append("file", file);
            try {
                await UploadUserAvatar(formData);
                setIsTakePhoto(false)
            } catch (e) {
                console.log(e)
            }
        }
        setCurrentUser({...currentUser, avatar: avatar})
    }
    let navigate = useNavigate();
    const toSetting = () => {
        navigate(SettingPrefix)
    }
    const uploadBack = () => {
        setIsTakePhoto(false)
    }
    return (
        <>
            {
                isTakePhoto ? <div>
                        <SettingHeader title="修改头像" back={uploadBack}/>
                        <div style={{padding: "0 24px"}}>
                            <UploadAvatar isUploadAvatar={false}
                                          saveAvatarDo={uploadAvatarFile}/>
                        </div>
                    </div> :
                    (<div className="me">
                        <div className="me-top">
                            <button className="setting-btn" onClick={toSetting}>
                                <SvgIcon name="setting" width={35}/>
                            </button>
                        </div>
                        <img className="me-avatar" src={avatarData} alt=""></img>
                        <div className="me-avatar-take" onClick={toTakePhoto}>
                            <SvgIcon name="upload_photo" width={50}/>
                        </div>
                        <div className="me-content">
                            <div className="me-content-item">
                                <div className="me-content-item-label">nickname</div>
                                <div className="me-content-item-value"><SvgIcon
                                    name="me"/>{currentUser?.nick_name}
                                </div>
                            </div>
                            <div className="me-content-item">
                                <div className="me-content-item-label">account</div>
                                <div className="me-content-item-value">{currentUser?.account_id}</div>
                            </div>
                            <div className="me-content-item">
                                <div className="me-content-item-label">personal signature</div>
                                <div className="me-content-item-value">{currentUser?.personal_desc}</div>
                            </div>
                        </div>
                    </div>)
            }


        </>

    )
}

export default Me