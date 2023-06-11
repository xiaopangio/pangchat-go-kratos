import "@/pages/addFriend/chooseGroup/index.less"
import SettingHeader from "@/components/SettingHeader";
import SvgIcon from "@/components/Icon";
import {useNavigate} from "react-router";
import {useRecoilState} from "recoil";
import {currentUserState, GroupName, GroupNames} from "@/store";
import {DexieGetFriendGroup, DexieStoreFriendGroups} from "@/utils/store";
import {isNull} from "lodash";
import {RefreshCurrentUser} from "@/utils/util";
import {GetFriendGroupList} from "@/api/friend";
import {FriendGroup} from "@/store/db";
import {useEffect} from "react";

function ChooseGroup() {
    const [groupNames, setGroupNames] = useRecoilState(GroupNames);
    const [groupName, setGroupName] = useRecoilState(GroupName);
    let [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    let navigate = useNavigate();
    const initCurrentUser = () => {
        if (isNull(currentUser)) {
            setCurrentUser(RefreshCurrentUser())
        }
    }
    const initGroupNames = async () => {
        if (isNull(currentUser)) {
            return
        }
        if (groupNames.length > 0) {
            return
        }

        const res = await DexieGetFriendGroup(currentUser.uid)
        if (res && res.length > 0) {
            setGroupNames(res)
            setGroupName(res[0])
            return
        } else {
            const res = await GetFriendGroupList({user_id: currentUser.uid})
            let groups = res.group_names.map((item: string) => {
                return {name: item, uid: currentUser?.uid} as FriendGroup
            })
            await DexieStoreFriendGroups(groups)
            setGroupNames(res.group_names)
            setGroupName(res.group_names[0])
        }
    }
    const back = () => {
        navigate(-1)
    }
    const chooseGroup = (item: string) => {
        setGroupName(item)
    }
    useEffect(() => {
        initCurrentUser()
    }, []);
    useEffect(() => {
        initGroupNames().then()
    }, [currentUser]);

    return (
        <div className="choose-group">
            <SettingHeader back={back} title="移至分组"/>
            <div className="choose-group-add">
                <div className="choose-group-add-icon">
                    +
                </div>
                <div className="choose-group-add-text">
                    添加分组
                </div>
            </div>
            {
                groupNames.map((item, index) => {
                    return (
                        <div className="choose-group-item" onClick={() => {
                            chooseGroup(item)
                        }} key={item}>
                            <div className="choose-group-item-content">
                                {item}
                            </div>
                            <div className="choose-group-item-checked">
                                {groupName === item ? (<SvgIcon name="gou" color="#57B77D"/>) : ""}
                            </div>
                        </div>
                    )
                })
            }

        </div>
    )
}

export default ChooseGroup