import {Alert, Snackbar} from "@mui/material";
import {useState} from "react";

type MessageProps = {
    content: string;
    duration: number;
    type: "success" | "error" | "info" | "warning";
}

function Message(props: MessageProps) {
    const {content, duration, type} = {...props};
    // 开关控制：默认true,调用时会直接打开
    const [open, setOpen] = useState(true);
    // 关闭消息提示
    const handleClose = (event: any, reason: string) => {
        setOpen(false);
    };
    return <Snackbar
        open={open}
        autoHideDuration={duration}
        anchorOrigin={{vertical: 'top', horizontal: 'center'}}
        onClose={handleClose}>
        <Alert severity={type}>{content}</Alert>
    </Snackbar>
}

export default Message;