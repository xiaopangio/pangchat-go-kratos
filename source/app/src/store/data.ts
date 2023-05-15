export type SettingItem = {
    title: string;
    subTitle?: string;
    center?: boolean;
}
export type SettingGroup = {
    title?: string;
    list: SettingItem[];
}
export type SettingGroupList = {
    list: SettingGroup[];
}
export const settingGroupList: SettingGroupList = {
    list: [
        {
            list: [
                {
                    title: "账号与安全",

                },
            ]
        },
        {
            list: [
                {
                    title: "青少年模式",

                },
                {
                    title: "关怀模式",

                },
            ],
        },
        {
            list: [
                {
                    title: "新消息通知",

                },
                {
                    title: "聊天",

                },
                {
                    title: "通用",

                }
            ],
        },
        {
            title: "隐私",
            list: [
                {
                    title: "朋友权限",

                },
                {
                    title: "个人信息与权限",

                },
                {
                    title: "个人信息收集清单",

                },
                {
                    title: "第三方信息共享清单",

                }
            ]
        },
        {
            list: [
                {
                    title: "插件",

                },
            ]
        },
        {
            list: [
                {
                    title: "关于pangchat",

                },
                {
                    title: "帮助与反馈",

                },
            ]
        },
        {
            list: [
                {
                    title: "切换账号",
                    center: true
                },
            ]
        }, {
            list: [
                {
                    title: "退出",
                    center: true
                },
            ],
        }
    ]
}
