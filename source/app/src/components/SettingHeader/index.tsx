import "@/components/SettingHeader/index.less";
import SvgIcon from "@/components/Icon";
import {useEffect, useRef} from "react";

type Props = {
    back: () => void;
    title: string;
    color?: string;
    menu?: () => void;
}
export default function SettingHeader({title, back, menu, color = "#ededed"}: Props) {
    const headerRef = useRef<HTMLDivElement>(null);
    useEffect(() => {
        if (headerRef.current) {
            headerRef.current.style.setProperty("--header-color", color);
        }
    }, []);

    return (
        <div className="setting-header" ref={headerRef}>
            <div className="setting-header__back" onClick={back}>
                <SvgIcon name="back"/>
            </div>
            <div className="setting-header__title">{title}</div>
            {
                menu ? <div className="setting-header__menu" onClick={menu}>
                    <SvgIcon name="menu" color="#181818"/>
                </div> : null
            }
        </div>
    );
}
