import PubSub from "pubsub-js";
import {ListenFriendRequest} from "@/service/service";

type WebSocketStatus = 'connecting' | 'open' | 'closing' | 'closed'
type UniversalMessage = {
    type: string
    data: any
}
let ws: WebSocket | null = null
let status: WebSocketStatus = 'closed'
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
            PubSub.publish("closeConn", "close")
        }
        ws.onerror = () => {
            status = 'closed';
            PubSub.publish("closeConn", "close")
        }
        ws.onmessage = publishMessage;
    } catch (e) {
        console.log(456)
        PubSub.publish("closeConn", "close")
        console.log(e);
    }
}
//pubsub service
const service = () => {
    ListenFriendRequest()
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
const closeConn = () => {
    ws?.close();
    console.log("close conn")
    ws = null;
    status = 'closed';
    reset()
}
const reconnect = (url: string) => {
    console.log("i am reconnecting")
    console.log(status)
    if (status === 'open' || status === 'connecting') {
        return;
    }
    if (status === 'closed') {
        try {
            closeConn()
            connect(url);
        } catch (e) {
            console.log(e);
        }
    }
}
// 心跳检测
const timeOut = 20000;
let timerObj: NodeJS.Timeout | null = null
let serverTimeoutObj: NodeJS.Timeout | null = null

const reset = () => {
    clearInterval(timerObj!);
    clearTimeout(serverTimeoutObj!);
    return this;
};

const start = () => {
    serverTimeoutObj = setTimeout(() => {
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
        close();
    }
}
export {
    connect,
    send,
    closeConn,
    reconnect,
    status,
    websocketInit
}

