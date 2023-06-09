import "@/pages/musicPlayer/index.less";
import SvgIcon from "@/components/Icon";
import musicSource from "./qingtian.mp3"
import React, {useEffect, useRef, useState} from "react";
import {isNull} from "lodash";
import {lrcData} from "@/pages/musicPlayer/data";

function MusicPlayer() {
    const [isPlaying, setIsPlaying] = useState<boolean>(false);
    const [currentTime, setCurrentTime] = useState<string>("0:00");
    const [duration, setDuration] = useState<string>("0:00");
    const audioRef = useRef<HTMLAudioElement>(null);
    const progressBarRef = useRef<HTMLDivElement>(null);
    const progressRef = useRef<HTMLDivElement>(null);
    const currentTimeDotRef = useRef<HTMLDivElement>(null);
    const containerRef = useRef<HTMLDivElement>(null);
    const lrcListRef = useRef<HTMLUListElement>(null);
    const [progressBarWidth, setProgressBarWidth] = useState<number>(0);
    const [lrcItems, setLrcItems] = useState<LrcItem[]>([]);
    const [containerHeight, setContainerHeight] = useState<number>(0);
    const [lrcItemHeight, setLrcItemHeight] = useState<number>(0);
    /**
     * 播放和暂停音乐
     */
    const togglePlayPause = (): void => {
        const audio = audioRef.current;
        if (isPlaying) {
            audio?.pause();
            setIsPlaying(false);
        } else {
            if (audio) {
                audio.play().then(
                );
                setIsPlaying(true);
            }
        }
    };
    /**
     * 更新进度条和时间显示
     */
    const updateProgress = (): void => {
        if (progressBarWidth === 0) {
            const progressBar = progressBarRef.current;
            if (progressBar) {
                setProgressBarWidth(progressBar.clientWidth);
            }
        }
        if (containerHeight === 0) {
            const container = containerRef.current;
            if (container) {
                setContainerHeight(container.clientHeight);
            }
        }
        if (lrcItemHeight === 0) {
            const lrcList = lrcListRef.current;
            if (lrcList) {
                setLrcItemHeight(lrcList.children[0].clientHeight);
            }
        }
        const audio = audioRef.current;
        if (audio) {
            const percent = (audio.currentTime / audio.duration) * 100;
            const currentTimeFormatted = formatTime(audio.currentTime);
            setCurrentTime(currentTimeFormatted);
            setDuration(formatTime(audio.duration));
            const progress = progressRef.current;
            if (progress) {
                progress.style.width = percent + "%";
            }
            const currentTimeDot = currentTimeDotRef.current;
            if (currentTimeDot) {
                const transformX = progressBarWidth * (percent / 100) - 5;
                currentTimeDot.style.transform = "translateY(-50%) translateX(" + transformX + "px)";
            }
            const activeIndex = getActiveIndex(audio.currentTime);
            if (activeIndex !== -1) {
                const lrcList = lrcListRef.current;
                if (lrcList) {
                    let activeDom = document.querySelector(".active")
                    if (activeDom) {
                        activeDom.classList.remove("active");
                    }
                    activeDom = lrcList.children[activeIndex];
                    activeDom.classList.add("active");
                }
            }
            setListOffset(activeIndex);
        }
    };
    /**
     * 调整进度条,当点击进度条时调整进度条
     * @param e
     */
    const adjustProgress = (e: React.MouseEvent<HTMLDivElement, MouseEvent>): void => {
        // 拿到progressBar相对视口的位置
        const progressBar = progressBarRef.current;
        if (progressBar) {
            const progressBarRect = progressBar.getBoundingClientRect();
            const x = e.clientX - progressBarRect.left;
            const percent = (x / progressBar.clientWidth) * 100;
            const audio = audioRef.current;
            if (audio) {
                audio.currentTime = (audio.duration * percent) / 100;
            }
        }
    }
    /**
     * 移动小圆点，当拖动小圆点时调整进度条，使用移动端的touch事件
     * @param e
     */
    const moveDot = (e: React.TouchEvent<HTMLDivElement>): void => {
        const progressBar = progressBarRef.current;
        if (progressBar) {
            const progressBarRect = progressBar.getBoundingClientRect();
            const x = e.touches[0].clientX - progressBarRect.left;
            const percent = (x / progressBar.clientWidth) * 100;
            const audio = audioRef.current;
            if (audio) {
                audio.currentTime = (audio.duration * percent) / 100;
            }
        }
    }
    /**
     * 格式化时间
     * @param time
     */
    const formatTime = (time: number): string => {
        const minutes = Math.floor(time / 60);
        const seconds = Math.floor(time % 60);
        const secondsFormatted = seconds < 10 ? "0" + seconds : seconds;
        return minutes + ":" + secondsFormatted;
    };
    type LrcItem = {
        time: number,
        content: string
    }
    /**
     * 解析时间字符串为数字
     * @param timeStr
     */
    const parseTimeStrToNumber = (timeStr: string): number => {
        let formatTimeStr = timeStr.replace("[", "").replace("]", "");
        let parts = formatTimeStr.split(":");
        const minutes = parseInt(parts[0]);
        const seconds = parseFloat(parts[1]);
        return minutes * 60 + seconds;
    }
    /**
     * 解析歌词
     * @param lrcStr
     */
    const parseLrc = (lrcStr: string) => {
        let lrcItems: LrcItem[] = [];
        lrcStr.split("\n").filter((item) => {
            //[xx:xx.xx] xxx 是正确的格式,最后一个item可能出现 “QQ”字样，需要过滤掉
            return /^\[\d{2}:\d{2}.\d{2}]/.test(item) && item.indexOf("QQ") === -1;
        }).map((item) => {
            // item 格式 [xx:xx.xx][xx:xx.xx]xxx ,时间可能有多个，需要全部分开处理
            const timeStrArr = item.match(/\[\d{2}:\d{2}.\d{2}]/g);
            if (isNull(timeStrArr)) {
                return
            }
            //取出歌词
            const lrc = item.split(timeStrArr[timeStrArr.length - 1])[1];
            //取出时间
            let tempItems = timeStrArr.map((timeStr) => {
                return {
                    time: parseTimeStrToNumber(timeStr),
                    content: lrc
                } as unknown as LrcItem
            })
            lrcItems.push(...tempItems)
        })
        lrcItems.sort((a, b) => {
            return a.time - b.time
        })
        setLrcItems(lrcItems)
    }
    /**
     * 获取当前播放的歌词的索引
     * @param currentTime
     */
    const getActiveIndex = (currentTime: number): number => {
        let activeIndex = -1;
        for (let i = 0; i < lrcItems.length; i++) {
            if (currentTime < lrcItems[i].time) {
                return i - 1;
            }
        }
        if (activeIndex === -1) {
            return lrcItems.length - 1;
        }
        return activeIndex;
    }
    /**
     * 设置歌词的偏移量
     * @param index
     */
    const setListOffset = (index: number) => {
        const lrcList = lrcListRef.current;
        if (isNull(lrcList)) {
            return 0
        }
        let offset = index * lrcItemHeight + lrcItemHeight / 2 - containerHeight / 2;
        if (offset < 0) {
            offset = 0;
        }
        if (offset > lrcList.clientHeight - containerHeight) {
            offset = lrcList.clientHeight - containerHeight;
        }
        lrcList.style.transform = `translateY(-${offset}px)`;
    }
    useEffect(() => {
        const audio = audioRef.current;
        if (isNull(audio)) {
            return
        }
        audio.addEventListener("timeupdate", updateProgress);
        parseLrc(lrcData)
        return () => {
            audio.removeEventListener("timeupdate", updateProgress);
        };

    }, [progressBarWidth, lrcItemHeight, containerHeight]);

    return (
        <div className="music_play_box">
            <div className="header">
                <div className="header_left">
                    <SvgIcon name="w-back" color="#ffffff"/>
                </div>
                <div className="header_center">
                    歌曲
                </div>
            </div>
            <div className="container" ref={containerRef}>
                <ul className="list" ref={lrcListRef}>
                    {
                        lrcItems.map((item, index) => {
                            return (
                                <li
                                    key={item.time}>{item.content}</li>
                            )
                        })
                    }
                </ul>
            </div>
            <div className="music-player">
                <audio
                    ref={audioRef}
                    src={musicSource}
                    onEnded={() => setIsPlaying(false)}
                ></audio>
                <div className="progress-box">
                    <span className="current-time">{currentTime}</span>
                    <div className="progress-bar" onClick={adjustProgress}
                         ref={progressBarRef}>
                        <div className="progress" ref={progressRef}></div>
                        <div className="progress-dot" ref={currentTimeDotRef} onTouchMove={moveDot}></div>
                    </div>
                    <span className="duration">{duration}</span>
                </div>
                <div className="control">
                    <button className="play-pause" onClick={togglePlayPause}>
                        <SvgIcon name={!isPlaying ? "play" : "pause"} color="#ffffff" width={40}/>
                    </button>
                </div>
            </div>

        </div>
    )
}

export default MusicPlayer;