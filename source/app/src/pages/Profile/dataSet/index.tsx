import "@/pages/Profile/dataSet/index.less"
import SettingHeader from "@/components/SettingHeader";
import {useEffect, useRef} from "react";
import {isNull} from "lodash";
import {defaultCallback, SlideIn, SlideOut} from "@/utils/animation";
import {useNavigate} from "react-router";
import SvgIcon from "@/components/Icon";
import SwitchButton from "@/components/SwitchButton";

function DataSet() {
    const dataSetRef = useRef<HTMLDivElement>(null);
    const navigate = useNavigate();
    useEffect(() => {
        if (isNull(dataSetRef.current)) return;
        SlideIn(dataSetRef, defaultCallback);
    }, [dataSetRef]);
    const back = () => {
        if (isNull(dataSetRef.current)) return;
        SlideOut(dataSetRef, () => {
            navigate(-1);
        });
    }
    const onChange = (checked: boolean) => {
        console.log(checked);
    }
    return (
        <div className="data-set" ref={dataSetRef}>
            <SettingHeader back={back} title="资料设置"/>
            <div className="data-set-container">
                <div className="data-set-item">
                    <div className="data-set-item-title">设置备注和标签</div>
                    <div className="data-set-item-option">
                        <div className="inner-option">
                            <div className="label">xxx</div>
                            <div className="arrow">
                                <SvgIcon name="rightArr"/>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="data-set-item">
                    <div className="data-set-item-title">朋友权限</div>
                    <div className="data-set-item-option">
                        <div className="inner-option">
                            <div className="arrow">
                                <SvgIcon name="rightArr"/>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="data-set-item">
                    <div className="data-set-item-title">把它推荐给朋友</div>
                    <div className="data-set-item-option">
                        <div className="inner-option">
                            <div className="arrow">
                                <SvgIcon name="rightArr"/>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="data-set-item">
                    <div className="data-set-item-title">添加到桌面</div>
                    <div className="data-set-item-option">
                        <div className="inner-option">
                            <div className="arrow">
                                <SvgIcon name="rightArr"/>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="data-set-item">
                    <div className="data-set-item-title">设为星标朋友</div>
                    <div className="data-set-item-option">
                        <div className="inner-option">
                            <SwitchButton onChange={onChange}/>
                        </div>
                    </div>
                </div>
                <div className="data-set-item">
                    <div className="data-set-item-title">加入黑名单</div>
                    <div className="data-set-item-option">
                        <div className="inner-option">
                            <div className="arrow">
                                <SvgIcon name="rightArr"/>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="data-set-item">
                    <div className="data-set-item-title">投诉</div>
                    <div className="data-set-item-option">
                        <div className="inner-option">
                            <div className="arrow">
                                <SvgIcon name="rightArr"/>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default DataSet