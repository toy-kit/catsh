import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import './assets/fonts/iconfont.css';
import i18n from './utils/locales'
import pinia from './utils/pinia'
import router from './router'

createApp(App).use(router).use(i18n).use(pinia).mount('#app')
