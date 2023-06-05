import {ToolOption} from "@/store/db";

export interface GetConnectionUrlResponse {
    host: string;
    port: string;
}

export interface GetToolOptionsResponse {
    options: ToolOption[];
}