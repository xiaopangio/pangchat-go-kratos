import { http } from "@/utils/request";

export function DelOnlineDevice(uid :string){
   return http.delete("/online",{
        params:{
            uid:uid
        }
   })
}