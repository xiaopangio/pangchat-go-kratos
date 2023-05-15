import "@/components/HomeHeader/index.less";
import SvgIcon from "@/components/Icon";
import {useEffect, useRef, useState} from "react";
import {useNavigate} from "react-router";

type Props = {
    title: string;
}

function HomeHeader({title}: Props) {
    const [addPopper, setAddPopper] = useState(false);
    const [addPopperClass, setAddPopperClass] = useState("add-popper");
    const addPopperRef = useRef<HTMLDivElement>(null);
    const addBtnRef = useRef<HTMLDivElement>(null);
    const searchBtnRef = useRef<HTMLDivElement>(null)
    let navigate = useNavigate();
    useEffect(() => {
        if (addPopper) {
            setAddPopperClass("add-popper add-popper-active");
        } else {
            setAddPopperClass("add-popper");
        }
    }, [addPopper]);
    const handleClick = (e: MouseEvent) => {
        //点击addBtnRef,将弹窗状态设置为true
        if (addBtnRef.current && addBtnRef.current.contains(e.target as Node)) {
            if (addPopperRef.current?.className.includes("add-popper-active")) {
                setAddPopper(false);
                return;
            }
            setAddPopper(true);
        } else if (addPopperRef.current && addPopperRef.current.contains(e.target as Node)) {
            //点击addPopperRef,不做任何操作
        } else if (searchBtnRef.current && searchBtnRef.current.contains(e.target as Node)) {
        } else {
            //点击其他地方,将弹窗状态设置为false
            setAddPopper(false);
        }
    }
    const handleAddFriend = () => {
        console.log("add friend");
        navigate("/addFriend")
    }
    useEffect(
        () => {
            document.addEventListener("click", handleClick);
        }, []
    );

    return (
        <div className="home-header">
            <div className="home-header-title">
                {title}
            </div>
            <div className="home-header-search" ref={searchBtnRef}>
                <SvgIcon name="search" width={25} color={"#000000"}/>
            </div>
            <div className="home-header-add" ref={addBtnRef}>
                <SvgIcon name="addto" width={32}/>
            </div>
            <div className={addPopperClass} ref={addPopperRef}>
                <div className="add-popper-item">
                    <SvgIcon name="chat" width={30} color="white"/>
                    <span>发起群聊</span>
                </div>
                <div className="add-popper-item" onClick={handleAddFriend}>
                    <SvgIcon name="add-people" width={30}/>
                    <span>添加好友</span>
                </div>
                <div className="add-popper-item">
                    <SvgIcon name="group" width={30}/>
                    <span>添加群聊</span>
                </div>
                <div className="add-popper-item">
                    <SvgIcon name="group" width={30}/>
                    <span>扫一扫</span>
                </div>
            </div>
        </div>
    )
}

export default HomeHeader;