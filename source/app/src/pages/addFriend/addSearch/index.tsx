import "@/pages/addFriend/addSearch/index.less"
import SvgIcon from "@/components/Icon";
import {FormEvent, useState} from "react";
import {debounce, isUndefined} from "lodash";
import AddFriendImg from "@img/add-friend.png"
import {useNavigate} from "react-router";
import {useRecoilState} from "recoil";
import {AddUserState} from "@/store";
import {Profile} from "@/api/user";

function AddSearch() {
    const [value, setValue] = useState("");
    const [profile, setProfile] = useRecoilState(AddUserState)
    const [isUserNull, setIsUserNull] = useState(false);
    let navigate = useNavigate();
    const inputDebounce = debounce((e: FormEvent<HTMLInputElement>) => {
        setIsUserNull(false)
        let target = e.target as HTMLInputElement;
        setValue(target.value)
    }, 500)
    const handleInput = (e: FormEvent<HTMLInputElement>) => {
        inputDebounce(e)
    }
    const handleSearch = async () => {
        try {
            const res = await Profile({account_id: value})
            if (isUndefined(res)) {
                setIsUserNull(true)
                return
            }
            setProfile(res)
            navigate(`/searchTarget`)
        } catch (e) {
        }
    }
    const handleCancel = () => {
        navigate("/addFriend")
    }
    return (
        <div className="add-search">
            <div className="add-search-header">
                <div className="add-click-box">
                    <SvgIcon name="search" width={20} color="#b5b5b5"/>
                    <input placeholder="账号/手机号"
                           onInput={handleInput}></input>
                </div>
                <div className="add-search-cancel" onClick={handleCancel}>取消</div>
            </div>
            {isUserNull ? <div className="add-null">该用户不存在</div> :
                <div>
                    {
                        value === "" ? null : <div className="add-search-content" onClick={handleSearch}>
                            <div className="add-search-item">
                                <img src={AddFriendImg} alt="" className="add-search-item-img"/>
                                <div className="add-search-item-content">
                                    搜索:<span>{value}</span>
                                </div>
                            </div>
                        </div>
                    }
                </div>
            }
        </div>
    )
}

export default AddSearch