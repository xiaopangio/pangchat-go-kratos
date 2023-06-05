import PubSub from "pubsub-js";
import {ListenFriend, ListenFriendRequest, ListenMessage, ListenMessageList} from "@/service/service";
import {DelOnlineDevice} from "@/api/online";
import {User} from "@/declare/type";
import {isNull} from "lodash";

type WebSocketStatus = 'connecting' | 'open' | 'closing' | 'closed'
type UniversalMessage = {
    type: string
    data: any
}
let ws: WebSocket | null = null
let status: WebSocketStatus = 'closed'
let IsUserExited: boolean = false
const connect = (url: string) => {
    try {
        ws = new WebSocket(url);
        ws.onopen = () => {
            status = 'open';
            start()
            service()
        }
        ws.onclose = () => {
            status = 'closed';
            console.log("连接被关闭")
            if (IsUserExited) {
                return
            }
            PubSub.publish("closeConn", "1")
        }
        ws.onerror = () => {
            status = 'closed';
            console.log("连接出错")
            // PubSub.publish("closeConn", "2")
        }
        ws.onmessage = publishMessage;
    } catch (e) {
        PubSub.publish("closeConn", "3")
        console.log(e);
    }
}
const exit = () => {
    IsUserExited = true
    reset()
}
//pubsub service
const service = () => {
    ListenFriend()
    ListenMessage()
    ListenFriendRequest()
    ListenMessageList()
}
// 引入pubsub，将接受到的消息发布出去
const publishMessage = (e: MessageEvent) => {
    const message: UniversalMessage = JSON.parse(e.data);
    PubSub.publish(message.type, message.data)
}
const send = (data: any) => {
    data = JSON.stringify(data);
    if (status === 'open') {
        ws?.send(data);
    }
}
const closeConn = (isExited: boolean = false, user: User | null = null) => {
    console.log("close conn")
    IsUserExited = isExited
    reset()
    ws?.close();
    ws = null;
    status = 'closed';
    if (isNull(user)) {
        return
    }
    DelOnlineDevice(user.uid).then(() => {
        console.log("del online device")
    }).catch(() => {
        console.log("del online device fail")
    })
}
const reconnect = (url: string) => {
    console.log("i am reconnecting")
    if (status === 'open' || status === 'connecting') {
        return;
    }
    if (status === 'closed') {
        try {
            connect(url);
        } catch (e) {
            // 重连失败
            console.log("重连失败")
        }
    }
}
// 心跳检测
const timeOut = 10000;
let timerObj: NodeJS.Timeout | null = null
let serverTimeoutObj: NodeJS.Timeout | null = null

const reset = () => {
    clearInterval(timerObj!);
    clearTimeout(serverTimeoutObj!);
    return this;
};

const start = () => {
    serverTimeoutObj = setTimeout(() => {
        console.log("未收到pong")
        closeConn()
        reset()
    }, timeOut + 1000)
    timerObj = setInterval(() => {
        let msg: UniversalMessage = {
            type: "ping",
            data: "ping"
        }
        send(msg)
    }, timeOut);
    PubSub.unsubscribe("pong")
    PubSub.subscribe("pong", (e) => {
        clearTimeout(serverTimeoutObj!);
        serverTimeoutObj = setTimeout(() => {
            console.log("未收到pong")
            closeConn()
            reset()
        }, timeOut + 1000)
    })
};
const websocketInit = (url: string) => {
    if (!ws) {
        connect(url);
    }
    return () => {
        closeConn();
    }
}
export {
    connect,
    send,
    closeConn,
    reconnect,
    status,
    websocketInit,
    UniversalMessage,
    exit
}

