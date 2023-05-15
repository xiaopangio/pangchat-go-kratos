import "@/pages/setting/account_setting/index.less"
import SettingHeader from "@/components/SettingHeader";
import {useNavigate} from "react-router";
import SettingContent from "@/components/SettingContent";
import {useRecoilValue} from "recoil";
import {currentUserState} from "@/store";
import {SettingGroupList} from "@/store/data";

export default function AccountSetting() {
    let navigate = useNavigate();
    let user = useRecoilValue(currentUserState);
    const settingGroupList: SettingGroupList = {
        list: [
            {
                list: [
                    {
                        title: "微信号",
                        subTitle: user?.account_id,
                    },
                    {
                        title: "手机号",
                        subTitle: user?.phone,
                    }
                ]
            },
            {
                list: [
                    {
                        title: "微信密码",
                    },
                    {
                        title: "声音锁",
                    },
                ]
            },
            {
                list: [
                    {
                        title: "应急联系人",
                    },
                    {
                        title: "登陆过的设备",
                    },
                    {
                        title: "更多安全设置",
                    }
                ]
            },
            {
                list: [
                    {
                        title: "微信安全中心",
                    },
                ]
            }
        ]
    }
    const toback = () => {
        navigate(-1)
    }
    const click = (title: string) => {
        console.log(title)
    }
    return (
        <div className="account_setting">
            <SettingHeader title="账号与安全" back={toback}/>
            <SettingContent click={click} settingGroupList={settingGroupList}/>
        </div>
    )
}