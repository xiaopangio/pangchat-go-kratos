import {Message} from "@/store/db"

export interface GetUnLoadMessageBeforeRequest {
    sender_id: string
    receiver_id: string
    message_id: string
    num: number
}

export interface GetUnLoadMessageBeforeResponse {
    messages: Message[]
}