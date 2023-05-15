import "@/pages/addFriend/searchTarget/index.less"
import {useEffect, useState} from "react";
import {useNavigate} from "react-router";
import SettingHeader from "@/components/SettingHeader";
import SvgIcon from "@/components/Icon";
import {GetAvatar} from "@/api/user";
import {isNull} from "lodash";
import {useRecoilValue} from "recoil";
import {AddUserState} from "@/store";
import {GetImg, StoreImg} from "@/utils/store";

function SearchTarget() {
    let navigate = useNavigate();
    const [imageData, setImageData] = useState("");
    let profile = useRecoilValue(AddUserState);

    useEffect(() => {
        if (isNull(profile)) {
            return
        }
        GetImg(profile?.avatar_url as string).then((res) => {
            if (res !== "") {
                setImageData(res)
                return
            }
            if (isNull(profile)) {
                return;
            }
            GetAvatar({avatar_url: profile.avatar_url}).then(res => {
                StoreImg(profile?.avatar_url as string, res).then((value) => {
                    setImageData(value)
                })
            })
        }, () => {
            if (isNull(profile)) {
                return;
            }
            GetAvatar({avatar_url: profile.avatar_url}).then(res => {
                StoreImg(profile?.avatar_url as string, res).then((value) => {
                    setImageData(value)
                })
            })
        })
    }, []);
    const back = () => {
        navigate("/addSearch")
    }
    const addFriend = () => {
        navigate("/addDetail")
    }
    return (

        <div className="search-target">
            <SettingHeader back={back} title=""/>
            <div className="target-profile">
                <img className="target-profile__avatar" alt="" src={imageData}/>
                <div className="target-profile__box">
                    <div className="target-profile__name">{profile?.nick_name}</div>
                    {profile?.city === "" ? null : <div className="target-profile__address">
                        地区: {profile?.province === profile?.city ? "中国大陆" : profile?.province} {profile?.city}
                    </div>}
                </div>
            </div>
            <div className="target-note-set">
                <div className="target-note-set__title">设置备注和标签</div>
                <SvgIcon name="rightArr"/>
            </div>
            {
                profile?.desc === "" ? null :
                    <div className="target-desc">个性签名 <span>为简单生活而努力着</span></div>
            }
            <div className="target-btn" onClick={addFriend}>添加到通讯录</div>
        </div>
    )
}

export default SearchTarget