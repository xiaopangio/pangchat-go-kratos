import {RefObject} from "react";

function SlideInAnimation() {
    let keyframes: Keyframe[] = [
        {
            transform: 'translateX(100%)'
        },
        {
            transform: 'translateX(0)'
        }
    ]
    let options: KeyframeAnimationOptions = {
        duration: 200,
        easing: 'ease-in-out',
        fill: 'forwards'
    }
    return {
        keyframes,
        options
    }
}

function SlideOutAnimation() {
    let keyframes: Keyframe[] = [
        {
            transform: 'translateX(0)'
        },
        {
            transform: 'translateX(100%)'
        }
    ]
    let options: KeyframeAnimationOptions = {
        duration: 200,
        easing: 'ease-in-out',
        fill: 'forwards'
    }
    return {
        keyframes,
        options
    }
}

const defaultCallback = () => {
}

function SlideIn(ref: RefObject<any>, callback: () => void) {
    let animation = SlideInAnimation()
    ref.current.animate(animation.keyframes, animation.options)
    if (callback) {
        setTimeout(callback, animation.options.duration as number)
    }
}

function SlideOut(ref: RefObject<any>, callback: () => void) {
    let animation = SlideOutAnimation()
    ref.current.animate(animation.keyframes, animation.options)
    if (callback) {
        setTimeout(callback, animation.options.duration as number)
    }
}

export {SlideInAnimation, SlideOutAnimation, SlideIn, SlideOut, defaultCallback}