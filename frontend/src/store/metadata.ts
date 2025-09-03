import { defineStore } from 'pinia'
import i18n from '../utils/locales'
import { SaveConf } from "../../wailsjs/go/main/App"

export const useMetadataStore = defineStore('Metadata', {
    state: () => ({
        wails_config: {
            homepage: "",
            repository: "",
        },
        config: {
            locale: "zh",
            theme: "light",
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
        setConfig(cfg: any) {
            this.changeTheme(cfg.theme || "light", false);
            this.changeLocale(cfg.locale || "zh", false);
        },
        saveConf() {
            SaveConf(this.config);
        },
        setWailsConfig(value: any) {
            this.wails_config = value;
        }
    },
})