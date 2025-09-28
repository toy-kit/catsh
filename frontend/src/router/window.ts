import { NewWindow } from '../../bindings/catsh/service/windowservice';
import { WindowOptions } from "../../bindings/catsh/types/models";
export const ToSetting = () => {
    NewWindow(new WindowOptions({
        name: "setting",
        title: '设置',
        url: '/#/setting',
        width: 435,
        height: 475,
        resizable: false,
    }))
}
export const ToAbout = () => {
    NewWindow(new WindowOptions({
        name: "about",
        title: '设置',
        url: '/#/about',
        width: 400,
        height: 450,
        resizable: false,
    }))
}