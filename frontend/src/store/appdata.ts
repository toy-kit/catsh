import { defineStore } from 'pinia'
import i18n from '../utils/locales'
import { SaveConfig } from "../../bindings/catsh/service/appdataservice"
import {AppConfig, Config, AppConfigInfo, AppConfigRepository} from "../../bindings/catsh/types/models";

export const useAppDataStore = defineStore('AppData', {
    state: () => ({
        app_config: {
            info: {} as AppConfigInfo,
            repository: {} as AppConfigRepository,
            os: "",
        },
        config: {
            locale: "zh",
            theme: "light",
        },
        resize: {
            width: 0,
            height: 0,
        },
    }),
    getters: {
    },
    actions: {
        changeTheme(value: string, save: boolean = true) {
            this.config.theme = value;
            document.body.setAttribute('arco-theme', value)
            if (save) {
                this.saveConf();
            }
        },
        changeLocale(value: any, save: boolean = true) {
            this.config.locale = value;
            i18n.global.locale = value;
            if (save) {
                this.saveConf();
            }
        },
        setConfig(cfg: Config) {
            this.changeTheme(cfg.theme || "light", false);
            this.changeLocale(cfg.locale || "zh", false);
        },
        saveConf() {
            SaveConfig(this.config);
        },
        setAppConfig(value: AppConfig) {
            this.app_config = value;
        },
        setResize(value: any){
            this.resize = value;
        }
    },
})