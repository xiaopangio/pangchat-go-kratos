import {useEffect, useRef} from "react";

export function useInterval(callback: any, timeout = 1000) {
    const latestCallback = useRef(() => {
    });
    const timerRef = useRef<any>();
    useEffect(() => {
        latestCallback.current = callback;
    });
    useEffect(() => {
        timerRef.current = setInterval(() => latestCallback.current(), timeout);
        return () => clearInterval(timerRef.current);
    }, []);
    //重置定时器函数
    const reset = () => {
        timerRef.current = setInterval(() => latestCallback.current(), timeout);
    }
    //清除定时器函数
    const clear = () => {
        clearInterval(timerRef.current)
    }
    return [reset, clear]
}
