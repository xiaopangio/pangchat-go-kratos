import Message from "@/components/Message";
import {createRoot} from "react-dom/client";

type messageProps = {
    content: string;
    duration: number;
}
const message = {
    dom: HTMLDivElement,
    success({content, duration}: messageProps) {
        // 创建一个dom
        this.dom = document.createElement('div');
        // 定义组件，
        const JSXdom = (<Message content={content} duration={duration} type='success'></Message>);
        let root = createRoot(this.dom);
        root.render(JSXdom);
        document.body.appendChild(this.dom);
    },
    error({content, duration}: messageProps) {
        this.dom = document.createElement('div');
        const JSXdom = (<Message content={content} duration={duration} type='error'></Message>);
        let root = createRoot(this.dom);
        root.render(JSXdom);
        document.body.appendChild(this.dom);
    },
    warning({content, duration}: messageProps) {
        this.dom = document.createElement('div');
        const JSXdom = (<Message content={content} duration={duration} type='warning'></Message>);
        let root = createRoot(this.dom);
        root.render(JSXdom);
        document.body.appendChild(this.dom);
    },
    info({content, duration}: messageProps) {
        this.dom = document.createElement('div');
        const JSXdom = (<Message content={content} duration={duration} type='warning'></Message>);
        let root = createRoot(this.dom);
        root.render(JSXdom);
        document.body.appendChild(this.dom);
    }
};

export default message;