import {useSendCodeReturn} from "@/declare/type";
import {useInterval} from "@/hooks/intervalHooks";
import {useEffect, useState} from "react";

export function useSendCode(delay = 60): useSendCodeReturn {
    const [reset, clear] = useInterval(() => {
        setTicker(ticker - 1);
    })
    const [isTickerValid, setIsTickerValid] = useState(false);
    const [ticker, setTicker] = useState(delay);
    const getCodeClick = () => {
        if (isTickerValid) {
            return;
        }
        setIsTickerValid(true);
        reset()
    }
    useEffect(() => {
        clear()
        setIsTickerValid(false);
    }, []);

    useEffect(() => {
        if (ticker <= 0) {
            clear();
            setTicker(delay)
            setIsTickerValid(false);
        }
    }, [ticker]);

    return {getCodeClick, isTickerValid, ticker}
}