import { NewWindow } from '../../bindings/catsh/service/windowservice';
import { WindowOptions } from "../../bindings/catsh/types/models";
export const ToSetting = () => {
    NewWindow(new WindowOptions({
        name: "setting",
        title: '设置',
        url: '/setting',
        width: 435,
        height: 475,
        resizable: false,
    }))
}