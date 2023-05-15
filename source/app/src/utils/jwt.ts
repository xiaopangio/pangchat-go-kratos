import jwtDecode from "jwt-decode";

export function ParseJwt<T>(token: string) {
    return jwtDecode<T>(token)
}