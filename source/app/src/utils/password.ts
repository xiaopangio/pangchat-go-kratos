import sha1 from "ts-sha1"
import {isEmpty} from "lodash";

const hashPassword = (password: string) => {
    if (isEmpty(password)) {
        return null
    }
    const salt = password.substring(0, password.length / 2)
    return sha1(sha1(password) + salt)
}
export default hashPassword