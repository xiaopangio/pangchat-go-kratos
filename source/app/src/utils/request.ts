import type {AxiosError, AxiosInstance, AxiosRequestConfig, AxiosResponse, InternalAxiosRequestConfig} from 'axios'
import axios from 'axios'
import {checkToken} from "@/utils/check";
import message from './message';

const service: AxiosInstance = axios.create({
    baseURL: 'http://127.0.0.1:8081/api/v1',
    timeout: 30000
})

service.interceptors.request.use((config: InternalAxiosRequestConfig) => {
    if (config && config.headers) {
        const token = checkToken()
        if (token) {
            config.headers['Authorization'] = token.value
        }
        if ((config.url === '/user/avatar' || config.url === '/avatar') && config.method === 'post') {
            config.headers['Content-Type'] = 'multipart/form-data'
        } else {
            config.data = JSON.stringify(config.data)
            config.headers['Content-Type'] = 'application/json;charset=UTF-8'
        }
    }
    return config
}, (error: AxiosError) => {
    return Promise.reject(error)
})

service.interceptors.response.use((response: AxiosResponse) => {
    if (response.headers['content-type'] === 'application/octet-stream') {
        response.config.responseType = 'blob'
        return response.data
    }
    const {code, message, body} = response.data

    // 根据自定义错误码判断请求是否成功
    if (code === 0) {
        // 将组件用的数据返回
        return body
    } else {
        // 处理业务错误。
        message.error({content: message, duration: 2000})
        return Promise.reject(new Error(message))
    }
}, (error: AxiosError) => {
    // 处理 HTTP 网络错误
    let m: string
    // HTTP 状态码
    const status = error.response?.status
    switch (status) {
        case 401:
            m = 'token 失效，请重新登录'
            // 这里可以触发退出的 action
            break;
        case 403:
            m = '拒绝访问'
            break;
        case 404:
            m = '请求地址错误'
            break;
        case 500:
            m = '服务器故障'
            break;
        default:
            m = '网络连接故障'
    }
    message.error({content: error.message, duration: 2000})
    return Promise.reject(error)
})
/* 导出封装的请求方法 */
export const http = {
    get<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
        return service.get(url, config)
    },

    post<T = any>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
        return service.post(url, data, config)
    },

    put<T = any>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
        return service.put(url, data, config)
    },

    delete<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
        return service.delete(url, config)
    }
}
