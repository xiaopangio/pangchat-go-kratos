import {useRecoilState} from "recoil";
import {currentUserState} from "@/store";
import { useEffect } from "react";
import {RefreshCurrentUser} from "@/utils/util";

export function useRefreshUser(){
    let [currentUser,setCurrentUser]= useRecoilState(currentUserState)
    useEffect(() => {
        if (!currentUser) {
            setCurrentUser(RefreshCurrentUser())
        }
    }, []);
}