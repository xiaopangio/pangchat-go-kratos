import "@/components/SwitchButton/index.less"
import {CSSProperties, useRef, useState} from "react";
import {isNull} from "lodash";

type SwitchButtonProps = {
    onChange: (checked: boolean) => void;
    size?: number;
}

function SwitchButton({onChange, size = 50}: SwitchButtonProps) {
    const labelRef = useRef<HTMLDivElement>(null);
    const containerRef = useRef<HTMLDivElement>(null);
    const [check, setCheck] = useState(false);
    const handleCheck = () => {
        if (isNull(labelRef.current)) return;
        if (isNull(containerRef.current)) return;
        if (!check) {
            containerRef.current.classList.remove("switch-container-inactive")
            containerRef.current.classList.add("switch-container-active")
            labelRef.current.classList.remove("switch-label-inactive")
            labelRef.current.classList.add("switch-label-active")
        } else {
            containerRef.current.classList.remove("switch-container-active")
            containerRef.current.classList.add("switch-container-inactive")
            labelRef.current.classList.remove("switch-label-active")
            labelRef.current.classList.add("switch-label-inactive")
        }
        setCheck(!check);
    }
    const generateStyle = () => {
        return {
            "--size": `${size}px`
        } as CSSProperties
    }
    return (
        <div className="switch-container" style={generateStyle()} ref={containerRef} onClick={handleCheck}>
            <div className="switch-label" ref={labelRef}></div>
        </div>
    )
}

export default SwitchButton