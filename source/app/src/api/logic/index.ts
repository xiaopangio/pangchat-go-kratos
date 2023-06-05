import {http} from "@/utils/request";
import {GetConnectionUrlResponse, GetToolOptionsResponse} from "@/api/logic/types";

export function GetConnectionUrl() {
    return http.get<GetConnectionUrlResponse>("/connectorUrl");
}

export function GetToolOptions() {
    return http.get<GetToolOptionsResponse>("/toolOptions");
}