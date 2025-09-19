import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import './assets/fonts/iconfont.css';
import i18n from './utils/locales'
import pinia from './utils/pinia'

createApp(App).use(i18n).use(pinia).mount('#app')
