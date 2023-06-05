import "@/pages/chat/dialogue/index.less"
import {useRecoilState} from "recoil";
import {currentDialogState, currentUserState, UnreadMessageCountMap} from "@/store";
import SettingHeader from "@/components/SettingHeader";
import {useNavigate} from "react-router";
import React, {CSSProperties, useEffect, useRef, useState} from "react";
import {isNull, isUndefined} from "lodash";
import {RefreshCurrentUser} from "@/utils/util";
import {
    DexieGetFriend,
    DexieGetMessageByMessageId,
    DexieGetMessagesBefore,
    DexieGetMessagesLimit,
    DexieGetToolOptions,
    DexieStoreMessage,
    DexieStoreMessages,
    DexieStoreToolOptions,
    DexieUpdateMessage,
    GetImg
} from "@/utils/store";
import {Friend, Message, MessageCountType, ToolOption} from "@/store/db";
import SvgIcon from "@/components/Icon";
import Slider, {Settings} from "react-slick";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import {send, UniversalMessage} from "@/hooks/websocket";
import {GetToolOptions} from "@/api/logic";
import {isEmpty} from "fast-glob/out/utils/string";
import {MessageType} from "@/declare/const";
import message from "@/utils/message";
import ReactPullLoad, {STATS} from "react-pullload";
import {MessagePageSize} from "@/declare/dialog";
import {GetUnLoadMessageBefore} from "@/api/message";

function Dialogue() {
    let [friendId, setFriendId] = useRecoilState(currentDialogState)
    const [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    const [unreadMessageCountMap, setUnreadMessageCountMap] = useRecoilState(UnreadMessageCountMap)
    const [profile, setProfile] = useState<Friend | null>(null);
    const [optionPages, setOptionPages] = useState<ToolOption[][]>()
    const [isMoreClose, setIsMoreClose] = useState(true)
    const [myAvatar, setMyAvatar] = useState("");
    const [friendAvatar, setFriendAvatar] = useState("");
    const [isReadySend, setIsReadySend] = useState(false);
    const [, setOptionList] = useState<ToolOption[]>([]);
    let navigate = useNavigate();
    const settings: Settings = {
        //详细的设置请查看官方API
        dots: true, //圆点显示（false隐藏）
        speed: 500, //自动播放速度（毫秒）
        slidesToShow: 1, //在一帧中显示3张卡片
        slidesToScroll: 1,//一次滚动3张卡片
        infinite: false, //无限循环
    };
    const moreToolsRef = useRef<HTMLDivElement>(null);
    const inputRef = useRef<HTMLDivElement>(null);
    const [messages, setMessages] = useState<Message[]>([]);
    const contentRef = useRef<HTMLDivElement>(null);
    const [pullLoadAction] = useState(STATS.init);
    const [hasMore, setHasMore] = useState(true);
    const [isReloadMessage, setIsReloadMessage] = useState(false);
    const [topMessageId, setTopMessageId] = useState("");
    const [isMeSendMsg, setIsMeSendMsg] = useState(false);
    const [firstIn, setFirstIn] = useState(true);
    // 单条消息的高度
    const [messageHeight, setMessageHeight] = useState(0);
    // scrollTop
    const [scrollTop, setScrollTop] = useState(1);

    enum LoadingMessage {
        Init = "",
        DropDown = "下拉加载更多",
        Release = "松开加载更多",
        Loading = "加载中...",
    }

    const [loadingMsg, setLoadingMsg] = useState<LoadingMessage>(LoadingMessage.Init);
    // 初始化聊天选项
    const InitOptions = async () => {
        try {
            let options = await DexieGetToolOptions()
            if (isUndefined(options) || options.length === 0) {
                const res = await GetToolOptions()
                options = res.options
                setOptionList(res.options)
                await DexieStoreToolOptions(res.options)
            } else {
                setOptionList(options)
            }
            let list = PagingOptionList(options, 8);
            setOptionPages(list)
        } catch (e: any) {
            message.error({
                content: e, duration: 1000
            })
        }
    }
    // 初始化用户头像
    const initAvatar = async () => {
        if (isNull(currentUser)) {
            return
        }
        try {
            let res = await GetImg(currentUser.avatar)
            setMyAvatar(res)
            if (isNull(profile)) {
                return
            }
            res = await GetImg(profile.avatar)
            setFriendAvatar(res)
        } catch (e: any) {
            message.error({
                content: e, duration: 1000
            })
        }
    }
    const initOfflineMessage = async (latestMessage: Message) => {
        if (!isUndefined(latestMessage.is_need_load_before) && !isUndefined(latestMessage.need_load_before_count)) {
            // 判断是否需要先加载
            if (latestMessage.is_need_load_before) {//需要先加载
                const loadCount = latestMessage.need_load_before_count - MessagePageSize > 0 ? MessagePageSize : latestMessage.need_load_before_count
                const response = await GetUnLoadMessageBefore({
                    message_id: latestMessage.message_id,
                    num: loadCount,
                    receiver_id: latestMessage.receiver_id,
                    sender_id: latestMessage.sender_id
                })
                let ms = response.messages
                latestMessage.is_need_load_before = false
                if (latestMessage.need_load_before_count - ms.length > 0) {
                    ms[ms.length - 1].is_need_load_before = true
                    ms[ms.length - 1].need_load_before_count = latestMessage.need_load_before_count - ms.length
                }
                await DexieUpdateMessage(latestMessage)
                await DexieStoreMessages(ms)
                return true
            }
        }
        return false
    }
    // 初始化消息列表
    const initMessageList = async () => {
        if (isNull(currentUser)) {
            return
        }
        if (isUndefined(unreadMessageCountMap)) {
            return
        }
        let f = unreadMessageCountMap.get(friendId)
        if (isUndefined(f)) { //map中没有该好友的未读消息
            const res = await DexieGetMessagesBefore(currentUser.uid, friendId, topMessageId, MessagePageSize)
            setMessages(res)
            setFirstIn(false)
            setScrollTop(0)
            return
        }
        if (firstIn) { //第一次进入此页面,此时进入页面时，有两种情况，一种是所有消息都在本地，另一种是有部分消息在本地，有部分消息在服务器需要先加载
            try {
                const latestMessage = await DexieGetMessageByMessageId(f.message_id)
                await initOfflineMessage(latestMessage)
            } catch (e: any) {
                message.error({
                    content: e, duration: 1000
                })
                return
            }
            const res = await DexieGetMessagesBefore(currentUser.uid, friendId, topMessageId, MessagePageSize)
            setMessages(res)
            setScrollTop(0)
            setFirstIn(false)
            RefreshUnreadCountMap();
            return
        }
        // 触发加载更多
        if (isReloadMessage) {
            setIsReloadMessage(false)
            await initOfflineMessage(messages[0])
            if (!hasMore) {
                return
            }
            let res = await DexieGetMessagesBefore(currentUser.uid, friendId, topMessageId, MessagePageSize)
            if (res.length === 0) {
                setHasMore(false)
                return
            }
            let isLoadBefore = false
            for (let m of res) {
                let ok = await initOfflineMessage(m)
                if (ok) {
                    isLoadBefore = true
                    break
                }
            }
            if (isLoadBefore) {
                res = await DexieGetMessagesBefore(currentUser.uid, friendId, topMessageId, MessagePageSize)
            }
            let num = messageHeight * res.length
            setScrollTop(num)
            setMessages(res.concat(messages))
            return
        }
        // 自己发送了消息
        if (isMeSendMsg) {
            setIsMeSendMsg(false)
            return
        }
        // 有未读消息
        const res = await DexieGetMessagesLimit(currentUser.uid, friendId, f.unread_count)
        if (res.length === 0) {
            return
        }
        let temp = messages.concat(res)
        setMessages(temp)
        RefreshUnreadCountMap();
        setScrollTop(0)
    }
    // 监听消息
    const ListenMsg = () => {
        PubSub.unsubscribe("single_message_reply")
        PubSub.subscribe("single_message_reply", async (msg: string, data: Message) => {
            await DexieUpdateMessage(data)
            unreadMessageCountMap.set(friendId, {
                friend_id: friendId,
                message_id: data.message_id,
                uid: currentUser?.uid,
                unread_count: 0
            } as MessageCountType)
            setUnreadMessageCountMap(new Map(unreadMessageCountMap))
        })
        PubSub.unsubscribe("single_message_reply_error")
        PubSub.subscribe("single_message_reply_error", () => {
            message.error(
                {
                    content: "消息发送失败",
                    duration: 2,
                }
            )
        })
    }
    // 对optionList分页
    const PagingOptionList = (optionList: ToolOption[], pageSize: number) => {
        // optionList分页，每页8个
        let pageList: ToolOption[][] = []
        let page: ToolOption[] = []
        for (let i = 0; i < optionList.length; i++) {
            page.push(optionList[i])
            if (page.length === pageSize || i === optionList.length - 1) {
                pageList.push(page)
                page = []
            }
        }
        return pageList
    }
    // 初始化好友信息
    const getFriendProfile = async () => {
        if (isNull(currentUser)) {
            return
        }
        try {
            const profile = await DexieGetFriend(currentUser.uid, friendId)
            setProfile(profile)
        } catch (e: any) {
            message.error({
                content: e, duration: 1000
            })
        }
    }
    // 获取好友名称
    const getName = () => {
        if (isNull(profile)) {
            return ""
        }
        return isUndefined(profile.note_name) || profile.note_name === profile.nick_name ? profile.nick_name : profile.note_name
    }
    // 输入框输入回调
    const onInput = (e: any) => {
        if (e.target.innerHTML === "") {
            setIsReadySend(false)
        } else {
            setIsReadySend(true)
        }
    }

    // 刷新未读消息数
    function RefreshUnreadCountMap() {
        let f = unreadMessageCountMap.get(friendId)
        if (isUndefined(f)) {
            return
        }
        if (f.unread_count === 0) {
            return
        }
        f.unread_count = 0
        let tempMap = new Map<string, MessageCountType>(unreadMessageCountMap)
        tempMap.set(friendId, f)
        setUnreadMessageCountMap(tempMap)
    }

    const scroll = (isTobottom: boolean = false) => {
        // 不要出现过渡动画
        if (isNull(contentRef.current)) {
            return
        }
        let number = contentRef.current.scrollHeight / messages.length;
        setMessageHeight(number)
        if (scrollTop === 0 || isTobottom) {
            setScrollTop(1)
            contentRef.current.scrollTop = contentRef.current.scrollHeight
            return
        }
        contentRef.current.scrollTop = scrollTop
    }
    const back = () => {
        setFriendId("")
        navigate(-1)
    }
    const menu = () => {
        console.log("menu")
    }
    const more = () => {
        if (isNull(moreToolsRef.current)) {
            return
        }
        if (!isMoreClose) {
            moreToolsRef.current.animate([
                {
                    height: "58vw"
                },
                {
                    height: "0"
                }
            ], {
                duration: 150,
                fill: "forwards",
            })
            setIsMoreClose(true)
            return
        }
        moreToolsRef.current.animate([
            {
                height: "0"
            },
            {
                height: "58vw"
            }
        ], {
            duration: 150,
            fill: "forwards",
        })
        setIsMoreClose(false)
    }
    const CloseMore = () => {
        if (isNull(moreToolsRef.current)) {
            return
        }
        if (!isMoreClose) {
            moreToolsRef.current.animate([
                {
                    height: "58vw"
                },
                {
                    height: "0vw"
                }
            ], {
                duration: 150,
                fill: "forwards",
            })
            setIsMoreClose(true)
            return
        }
    }
    // 获取时间的显示格式 2021-08-01 12:00:00.123 要求精确到毫秒
    const getTime = () => {
        const date = new Date()
        const year = date.getFullYear()
        const month = date.getMonth() + 1
        const day = date.getDate()
        const hour = date.getHours()
        const minute = date.getMinutes()
        const second = date.getSeconds()
        const millisecond = date.getMilliseconds()
        return `${year}-${month}-${day} ${hour}:${minute}:${second}.${millisecond}`
    }
    const sendMsg = async () => {
        if (isNull(inputRef.current)) {
            return
        }
        if (isNull(currentUser)) {
            return
        }
        if (isEmpty(inputRef.current.innerHTML)) {
            return
        }
        // 判断是否全为空格
        const reg = /^\s+$/g
        if (reg.test(inputRef.current.innerHTML)) {
            inputRef.current.innerHTML = ""
            return;
        }
        // 去除message的前后空格
        const content = inputRef.current.innerHTML.trim()
        let msg: Message = {
            message_id: "3",
            type: MessageType.Text,
            content: content,
            sender_id: currentUser.uid,
            receiver_id: friendId,
            send_at: getTime()
        }
        const newMessages = [...messages, msg]
        setMessages(newMessages)
        let tranMsg: UniversalMessage = {
            type: "single_message",
            data: msg
        }
        send(tranMsg)
        await DexieStoreMessage(msg)
        inputRef.current.innerHTML = ""
        setIsMeSendMsg(true)
        setScrollTop(0)
    }
    const getDisplayStyle = (sender: string) => {
        return {
            "--before-display": sender === friendId ? "block" : "none",
            "--after-display": sender === friendId ? "none" : "block"
        } as CSSProperties
    }
    // 下拉加载
    const handlePullLoadAction = (action: STATS) => {
        switch (action) {
            case STATS.pulling:
                setLoadingMsg(LoadingMessage.DropDown)
                break
            case STATS.enough:
                setLoadingMsg(LoadingMessage.Release)
                break
            case STATS.refreshing:
                setLoadingMsg(LoadingMessage.Loading)
                setTopMessageId(messages[0].message_id)
                setIsReloadMessage(true)
                break
        }

    }
    // 初始化current_user
    useEffect(() => {
        if (isNull(currentUser)) {
            setCurrentUser(RefreshCurrentUser())
        }
        InitOptions().then()
        ListenMsg()
    }, [])
    // 初始化好友信息
    useEffect(() => {
        getFriendProfile().then()
    }, [currentUser]);
    // 初始化消息列表
    useEffect(() => {
        initMessageList().then()
        ListenMsg()
    }, [currentUser, unreadMessageCountMap, topMessageId])
    // 消息列表滚动到底部
    useEffect(() => {
        if (scrollTop === 1) {
            return
        }
        scroll()
    }, [scrollTop, isMeSendMsg]);
    // 初始化头像
    useEffect(() => {
        initAvatar().then()
    }, [profile])
    return (
        <div className="dialogue">
            <SettingHeader title={getName()} back={back} color={"#ededed"} menu={menu}/>
            <div className="dialogue__content" onClick={CloseMore} ref={contentRef}>
                <ReactPullLoad
                    action={pullLoadAction}
                    handleAction={handlePullLoadAction}
                    hasMore={hasMore}
                    downEnough={150}
                >
                    <div>
                        <div className="pull-loading">
                            {
                                loadingMsg
                            }
                        </div>
                        {
                            messages.map((message, _) => {
                                return (
                                    <div
                                        className={`dial__message ${message.sender_id === friendId ? "" : "flex-reverse"}`}
                                        key={message.send_at}>
                                        <div className="dial__message__avatar">
                                            <img className="dial__avatar__img"
                                                 src={message.sender_id === friendId ? friendAvatar : myAvatar} alt=""/>
                                        </div>
                                        <div className="dial__message__content green"
                                             style={getDisplayStyle(message.sender_id)}>
                                            {
                                                message.content
                                            }
                                        </div>
                                    </div>
                                )
                            })
                        }
                    </div>
                </ReactPullLoad>
            </div>

            <div className="dialogue__input">
                <div className="dialogue__input__top__box">
                    <div className="dialogue__speak__btn">
                        <SvgIcon name="voice" color="#000000" width={20}/>
                    </div>
                    <div className="dialogue__text_input" contentEditable="true" onInput={onInput}
                         ref={inputRef}
                    >
                    </div>
                    <div className="dialogue__emoji__btn">
                        <SvgIcon name="smile" color="#000000" width={33}/>
                    </div>
                    {
                        <div className={`dialogue__more__btn ${isReadySend ? "more-to-send" : ""}`}>
                            {!isReadySend ?
                                <div onClick={more} style={{display: "flex", alignItems: "center"}}>
                                    <SvgIcon name="add" color="#000000" width={20}/>
                                </div>
                                : <span onClick={sendMsg}>发送</span>}
                        </div>
                    }
                </div>
                <div className="dialogue__input__bottom__box" ref={moreToolsRef}>
                    <div className="dialogue__more__box_out">
                        <Slider {...settings}>
                            {
                                optionPages?.map((page, index) => {
                                    return (
                                        <div key={index} className="dialogue__more__box">
                                            {
                                                page.map((option, _) => {
                                                    return (
                                                        <div className="dialogue__more__item" key={option.name}>
                                                            <div className="dialogue__more__item__icon">
                                                                <SvgIcon name={option.icon} color="#000000" width={30}/>
                                                            </div>
                                                            <div
                                                                className="dialogue__more__item__text">{option.name}</div>
                                                        </div>
                                                    )
                                                })
                                            }
                                        </div>
                                    )
                                })
                            }
                        </Slider>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Dialogue