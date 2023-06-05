import "@/components/PullLoadHeaderNode/index.less"
import {useEffect} from "react";
import {STATS} from "react-pullload";

type PullLoadHeaderNodeProps = {
    loaderState: STATS
}

function PullLoadHeaderNode({loaderState = STATS.init}: PullLoadHeaderNodeProps) {
    console.log(loaderState);
    useEffect(() => {
        console.log(loaderState);
    }, [loaderState]);

    return (
        <div>123</div>
    )
}

export default PullLoadHeaderNode
