import "@/components/TakePhoto/index.less";
import {useEffect, useRef, useState} from "react";
import SvgIcon from "@/components/Icon";

type TakePhotoProps = {
    takeCancel: () => void;
    savePhoto: (img: string) => void;
}

function TakePhoto({takeCancel, savePhoto}: TakePhotoProps) {
    const cameraVideoRef = useRef<HTMLVideoElement>(null);
    const cameraCanvasRef = useRef<HTMLCanvasElement>(null);
    const [isTake, setIsTake] = useState(false);
    const [front, setFront] = useState(false);
    const [imageSrc, setImageSrc] = useState("");
    const takeClick = () => {
        setIsTake(true);
        let img = getImg();
        if (img) {
            setImageSrc(img);
        }
    }
    const cancelSaveClick = async () => {
        setIsTake(false);
        await onMedia()
    }
    const turnClick = async () => {
        setFront(!front);
        closeMedia()
    }
    const handleSavePhoto = () => {
        savePhoto(imageSrc);
        takeCancel();
    }
    const videoSuccess = (mediaStream: MediaStream) => {
        const video = cameraVideoRef.current;
        if (!video) return;
        video.srcObject = mediaStream;
        video.onloadedmetadata = () => {
            video.play();
        }
    }
    const videoError = (error: any) => {
        console.log(error);
    }
    const onMedia = async () => {
        const constraints: MediaStreamConstraints = {
            video: {
                facingMode: (front ? "user" : "environment"),
                width: 327,
                height: 450
            }
        };
        try {
            const mediaStream = await navigator.mediaDevices.getUserMedia(constraints);
            videoSuccess(mediaStream);
        } catch (e) {
            videoError(e);
        }
    }
    const closeMedia = () => {
        const video = cameraVideoRef.current;
        if (!video) return;
        const stream = video.srcObject as MediaStream;
        const tracks = stream?.getTracks();
        tracks?.forEach((track) => {
            track.stop();
        });
        video.srcObject = null;
    }
    const getImg = () => {
        const video = cameraVideoRef.current;
        const canvas = cameraCanvasRef.current;
        if (!video || !canvas) return;
        const ctx = canvas.getContext('2d');
        if (!ctx) return;
        const ratio = window.devicePixelRatio || 1;
        ctx.scale(ratio, ratio);
        canvas.width = video.offsetWidth * ratio;
        canvas.height = video.offsetHeight * ratio;
        ctx.drawImage(video, 0, 0, canvas.width, canvas.height); // 把视频中的一帧在canvas画布里面绘制出来
        const imgStr = canvas.toDataURL(); // 将图片资源转成字符串
        closeMedia(); // 获取到图片之后可以自动关闭摄像头
        return imgStr;
    };
    useEffect(() => {
        onMedia().then(() => {
        })
        return () => {
            closeMedia();
        }
    }, [front]);
    return (
        <div className="photo-take">
            <div className="photo">
                <canvas
                    id="cameraCanvas"
                    ref={cameraCanvasRef}
                    className="camera-canvas"
                />
                {
                    isTake ? <img className="display-img" src={imageSrc} alt={""}></img> : <video
                        id="cameraVideo"
                        ref={cameraVideoRef}
                        className="camera-video"
                    />
                }
            </div>
            <div className="button-group">
                {
                    isTake ? (
                            <div className="save-btn">
                                <button className="cancel" onClick={cancelSaveClick}>
                                    <SvgIcon name="close"></SvgIcon>
                                </button>
                                <button className="save" onClick={handleSavePhoto}>
                                    <SvgIcon name="gou"></SvgIcon>
                                </button>
                            </div>
                        ) :
                        (
                            <div className="take-btn">
                                <button className="cancel" onClick={takeCancel}>
                                    <SvgIcon name="close"></SvgIcon>
                                </button>
                                <button className="take" onClick={takeClick}></button>
                                <button className="turn" onClick={turnClick}>
                                    <SvgIcon name="flip"></SvgIcon>
                                </button>
                            </div>
                        )
                }


            </div>
        </div>
    )
}

export default TakePhoto;