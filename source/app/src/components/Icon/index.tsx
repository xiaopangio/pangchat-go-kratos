import "@/components/Icon/index.less"

type props = {
    prefix?: string,
    name: string,
    color?: string,
    width?: number,
}

function SvgIcon({prefix = "icon", name, color = "#cccccc", width = 30}: props) {
    return (
        <svg className="svg-icon" aria-hidden="true" width={width + "px"} height={width + "px"}>
            <use xlinkHref={`#${prefix}-${name}`} fill={color}/>
        </svg>
    )
}

export default SvgIcon