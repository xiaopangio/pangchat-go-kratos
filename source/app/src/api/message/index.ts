import {http} from "@/utils/request";
import {GetUnLoadMessageBeforeRequest, GetUnLoadMessageBeforeResponse} from "@/api/message/types";

const messagePrefix = "/message"

export function GetUnLoadMessageBefore(data: GetUnLoadMessageBeforeRequest) {
    return http.get<GetUnLoadMessageBeforeResponse>(messagePrefix + "/unload", {
        params: data
    })
}