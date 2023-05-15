import "@/components/SettingHeader/index.less";
import SvgIcon from "@/components/Icon";

type Props = {
    back: () => void;
    title: string;
}
export default function SettingHeader({title, back}: Props) {
    return (

        <div className="setting-header">
            <div className="setting-header__back" onClick={back}>
                <SvgIcon name="back"/>
            </div>
            <div className="setting-header__title">{title}</div>
        </div>
    );
}
