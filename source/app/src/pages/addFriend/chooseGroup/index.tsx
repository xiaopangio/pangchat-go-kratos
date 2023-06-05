import "@/pages/addFriend/chooseGroup/index.less"
import SettingHeader from "@/components/SettingHeader";
import SvgIcon from "@/components/Icon";
import {useNavigate} from "react-router";
import {useRecoilState} from "recoil";
import {GroupName, GroupNames} from "@/store";

function ChooseGroup() {
    const [groupNames, setGroupNames] = useRecoilState(GroupNames);
    const [groupName, setGroupName] = useRecoilState(GroupName);
    let navigate = useNavigate();
    const back = () => {
        navigate(-1)
    }
    const chooseGroup = (item: string) => {
        setGroupName(item)
    }
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