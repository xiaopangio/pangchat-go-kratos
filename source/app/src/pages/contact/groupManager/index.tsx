import SettingHeader from "@/components/SettingHeader";
import SvgIcon from "@/components/Icon";
import {useRecoilState} from "recoil";
import {currentUserState, GroupName, GroupNames} from "@/store";
import {useNavigate} from "react-router";
import {isNull} from "lodash";
import {RefreshCurrentUser} from "@/utils/util";
import {
    DexieDeleteFriendGroup,
    DexieGetFriendGroup,
    DexieStoreFriendGroup,
    DexieStoreFriendGroups,
    DexieUpdateFriendGroup
} from "@/utils/store";
import {CreateFriendGroup, DeleteFriendGroup, GetFriendGroupList, UpdateFriendGroup} from "@/api/friend";
import {FriendGroup} from "@/store/db";
import {RefObject, useEffect, useRef, useState} from "react";
import "@/pages/contact/groupManager/index.less"
import message from "@/utils/message";

function GroupManager() {
    const [groupNames, setGroupNames] = useRecoilState(GroupNames);
    const [, setGroupName] = useRecoilState(GroupName);
    let [currentUser, setCurrentUser] = useRecoilState(currentUserState);
    const groupManageListRef = useRef<HTMLDivElement>(null);
    const addPopperRef = useRef<HTMLDialogElement>(null);
    const deletePopperRef = useRef<HTMLDialogElement>(null);
    const inputRef = useRef<HTMLInputElement>(null);
    const [isAdd, setIsAdd] = useState(true);
    let navigate = useNavigate();
    const [oldGroupName, setOldGroupName] = useState("");
    const [deleteGroupName, setDeleteGroupName] = useState("");
    // 初始化当前用户
    const initCurrentUser = () => {
        if (isNull(currentUser)) {
            setCurrentUser(RefreshCurrentUser())
        }
    }
    // 初始化分组
    const initGroupNames = async () => {
        if (isNull(currentUser)) {
            return
        }
        if (groupNames.length > 0) {
            return
        }

        const res = await DexieGetFriendGroup(currentUser.uid)
        if (res && res.length > 0) {
            setGroupNames(res)
            setGroupName(res[0])
            return
        } else {
            const res = await GetFriendGroupList({user_id: currentUser.uid})
            let groups = res.group_names.map((item: string) => {
                return {name: item, uid: currentUser?.uid} as FriendGroup
            })
            await DexieStoreFriendGroups(groups)
            setGroupNames(res.group_names)
            setGroupName(res.group_names[0])
        }
    }
    // 返回
    const back = () => {
        navigate(-1)
    }
    // 选择分组
    const chooseGroup = (item: string) => {
        setGroupName(item)
    }
    const showPopper = (ref: RefObject<HTMLDialogElement>) => {
        ref.current?.showModal()
        ref.current?.animate([
            {transform: "translateY(-240%)"},
            {transform: "translateY(0)"}
        ], {
            duration: 300,
            fill: "forwards"
        })
    }
    const hidePopper = (ref: RefObject<HTMLDialogElement>) => {
        if (isNull(ref.current)) {
            return
        }
        ref.current.animate([
            {transform: "translateY(0)"},
            {transform: "translateY(-240%)"}
        ], {
            duration: 300,
            fill: "forwards"
        })
        setTimeout(() => {
            ref.current?.close()
        }, 300)
    }
    const handleShowDeleteBtn = (index: number) => {
        if (isNull(groupManageListRef.current)) {
            return
        }
        let groupListDom = groupManageListRef.current;
        let deleteBtn = groupListDom.children[index].querySelector(".group-manager-item-remove");
        let activeBtn = document.querySelector(".group-manager-item-remove-active");
        if (!isNull(activeBtn)) {
            activeBtn.classList.remove("group-manager-item-remove-active")
            activeBtn.classList.add("group-manager-item-remove-inactive")
        }
        if (isNull(deleteBtn)) {
            return
        }
        if (activeBtn === deleteBtn) {
            return
        }
        deleteBtn.classList.remove("group-manager-item-remove-inactive")
        deleteBtn.classList.add("group-manager-item-remove-active")
    }

    const handleAdd = () => {
        showPopper(addPopperRef)
        setIsAdd(true)
    }
    const addCancel = () => {
        if (isNull(inputRef.current)) {
            return
        }
        inputRef.current.value = ""
    }
    const deleteCancel = () => {
        setDeleteGroupName("")
    }
    const handleCancel = (ref: RefObject<HTMLDialogElement>, cancel: Function) => {
        cancel()
        hidePopper(ref)
    }
    const handleClose = (e: any, ref: RefObject<HTMLDialogElement>) => {
        if (e.target === ref.current) {
            hidePopper(ref)
        }
    }
    const AddOrUpdate = async () => {
        if (isNull(inputRef.current)) {
            return
        }
        let newGroupName = inputRef.current.value;
        inputRef.current.value = ""
        if (newGroupName.length === 0) {
            message.warning({
                content: "分组名不能为空", duration: 2000
            })
            return
        }
        if (isNull(currentUser)) {
            return
        }
        if (isAdd) {
            if (groupNames.includes(newGroupName)) {
                message.warning({
                    content: "分组名已存在", duration: 2000
                })
                return
            }
            try {
                await CreateFriendGroup(currentUser.uid, newGroupName)
                await DexieStoreFriendGroup(currentUser.uid, newGroupName)
                setGroupNames([...groupNames, newGroupName])
                message.success({
                    content: "添加成功", duration: 2000
                })
            } catch (e) {
                console.log(e)
                message.error({
                    content: "添加失败", duration: 2000
                })
            }
        } else {
            if (oldGroupName === newGroupName) {
                message.warning({
                    content: "分组名未改变", duration: 2000
                })
                return
            }
            if (oldGroupName === "我的好友") {
                message.warning({
                    content: "默认分组不可修改", duration: 2000
                })
                return
            }
            try {
                await UpdateFriendGroup(currentUser.uid, oldGroupName, newGroupName)
                await DexieUpdateFriendGroup(currentUser.uid, oldGroupName, newGroupName)
                let newGroupNames = groupNames.map((item: string) => {
                    if (item === oldGroupName) {
                        return newGroupName
                    }
                    return item
                })
                setGroupNames(newGroupNames)
                message.success({
                    content: "修改成功", duration: 2000
                })
            } catch (e: any) {
                message.error({
                    content: e.message, duration: 2000
                })
            }
        }
        hidePopper(addPopperRef)
    }
    const handleUpdate = (groupName: string) => {
        showPopper(addPopperRef)
        setIsAdd(false)
        if (isNull(inputRef.current)) {
            return
        }
        inputRef.current.value = groupName
        setOldGroupName(groupName)
    }
    const handleDelete = (groupName: string) => {
        showPopper(deletePopperRef)
        setDeleteGroupName(groupName)
    }
    // Delete 删除分组
    const Delete = async (groupName: string) => {
        if (isNull(currentUser)) {
            return
        }
        if (groupName === "我的好友") {
            message.warning({
                content: "默认分组不能删除", duration: 2000
            })
            return
        }
        try {
            await DeleteFriendGroup(currentUser.uid, groupName)
            await DexieDeleteFriendGroup(currentUser.uid, groupName)
            let newGroupNames = groupNames.filter((item: string) => {
                return item !== groupName
            })
            setGroupNames(newGroupNames)
            message.success({
                content: "删除成功", duration: 2000
            })
        } catch (e: any) {
            console.log(e.message)
            message.error({
                content: "删除失败", duration: 2000
            })
        }
        hidePopper(deletePopperRef)
    }
    useEffect(() => {
        initCurrentUser()
    }, []);
    useEffect(() => {
        initGroupNames().then()
    }, [currentUser]);

    return (
        <>
            <div className="group-manager">
                <SettingHeader back={back} title="分组管理"/>
                <div className="group-manager-add" onClick={handleAdd}>
                    <div className="group-manager-add-icon">
                        +
                    </div>
                    <div className="group-manager-add-text">
                        添加分组
                    </div>
                </div>
                <div className="group-manager-list" ref={groupManageListRef}>
                    {
                        groupNames.map((item, index) => {
                            return (
                                <div className="group-manager-item" onClick={() => {
                                    chooseGroup(item)
                                }} key={item}>
                                    <div className="group-manager-remove" onClick={() => {
                                        handleShowDeleteBtn(index)
                                    }}>
                                        <SvgIcon name="remove" color="#FF4D4F"/>
                                    </div>
                                    <div className="group-manager-item-content" onClick={() => {
                                        handleUpdate(item)
                                    }}>
                                        {item}
                                    </div>
                                    <div className="group-manager-item-drag">
                                        <SvgIcon name="list" color="#57B77D"/>
                                        <div className="group-manager-item-remove" onClick={() => handleDelete(item)}>
                                            <span className="text">删除</span>
                                        </div>
                                    </div>
                                </div>
                            )
                        })
                    }
                </div>
                <dialog className="group-manager-popper" ref={addPopperRef}
                        onClick={event => handleClose(event, addPopperRef)}>
                    <div className="group-manager-popper-title">
                        {isAdd ? "添加分组" : "修改分组"}
                    </div>
                    <div className="group-manager-add-popper-input">
                        <label>请输入新的分组名称</label>
                        <input type="text" placeholder="分组名" ref={inputRef}/>
                    </div>
                    <div className="group-manager-popper-btn">
                        <div className="group-manager-popper-btn-cancel"
                             onClick={() => handleCancel(addPopperRef, addCancel)}>
                            取消
                        </div>
                        <div className="group-manager-popper-btn-ok" onClick={AddOrUpdate}>
                            确定
                        </div>
                    </div>
                </dialog>
                <dialog className="group-manager-popper" ref={deletePopperRef}
                        onClick={event => handleClose(event, deletePopperRef)}>
                    <div className="group-manager-popper-title">
                        你确定要删除该分组吗？
                    </div>
                    <div className="group-manager-popper-title">
                        该分组下的好友将会移动到默认分组
                    </div>
                    <div className="group-manager-popper-btn">
                        <div className="group-manager-popper-btn-cancel"
                             onClick={() => handleCancel(deletePopperRef, deleteCancel)}>
                            取消
                        </div>
                        <div className="group-manager-popper-btn-ok" onClick={() => Delete(deleteGroupName)}>
                            确定
                        </div>
                    </div>
                </dialog>
            </div>
        </>
    )
}

export default GroupManager