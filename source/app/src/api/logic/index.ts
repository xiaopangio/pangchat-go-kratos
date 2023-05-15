import {http} from "@/utils/request";
import {GetConnectionUrlResponse} from "@/api/logic/types";

export function GetConnectionUrl() {
    return http.get<GetConnectionUrlResponse>("/connectorUrl");
}
