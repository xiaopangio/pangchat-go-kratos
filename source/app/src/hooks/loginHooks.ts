import {Login} from "@/api/user";

interface LoginHooks {
    login: typeof Login
}

export function useLogin(): LoginHooks {
    return {
        login: Login
    }
}