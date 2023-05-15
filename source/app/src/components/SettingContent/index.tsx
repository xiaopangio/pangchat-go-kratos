import {SettingGroupList} from "@/store/data";
import SvgIcon from "@/components/Icon";
import "@/components/SettingContent/index.less"

type Props = {
    click: (title: string) => void
    settingGroupList: SettingGroupList
}

function SettingContent({click, settingGroupList}: Props) {

    return (
        <div className="setting__content">
            {
                settingGroupList.list.map((group, index) => {
                    return (
                        <div className="setting__group" key={index}>
                            {
                                group.title && (<div className="setting__group-title">{group?.title}</div>)
                            }
                            <div className="setting__group_inner">
                                {
                                    group.list.map((item, index) => {
                                        return (
                                            <div className="setting__item" key={item.title} onClick={() => {
                                                click(item.title)
                                            }}>
                                                {item.center ?
                                                    <div
                                                        className="setting__item-title-center">{item.title}</div> :
                                                    <div className="setting__item-title">{item.title}</div>}
                                                {
                                                    item.subTitle ? <div
                                                        className="setting__item-subTitle">{item.subTitle}</div> : null
                                                }
                                                {
                                                    item.center ? null :
                                                        <div className="setting__item-icon">
                                                            <SvgIcon name="rightArr" color="rgb(187,187,187)"/>
                                                        </div>
                                                }
                                            </div>
                                        )
                                    })
                                }
                            </div>
                        </div>
                    )
                })
            }
        </div>
    )
}

export default SettingContent
