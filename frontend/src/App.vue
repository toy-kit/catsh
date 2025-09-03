<script lang="ts" setup>
import { ref } from 'vue'
import zhCN from '@arco-design/web-vue/es/locale/lang/zh-CN';
import enUS from '@arco-design/web-vue/es/locale/lang/en-us';
import { useMetadataStore } from './store/metadata'
import { Load } from "../wailsjs/go/main/App"
import i18n from './utils/locales'

// locale
const arcoLocales: any = {
  'zh': zhCN,
  'en': enUS,
};
const metadataStore = useMetadataStore()

const loading = ref(true)
const loadMetadata = () => {
  Load().then((res) => {
    metadataStore.setConfig(res.config || {})
    metadataStore.setWailsConfig(res.wails_config || {})
    const locales: any = res.locales || {}
    for (const key in locales) {
      i18n.global.setLocaleMessage(key, locales[key] || {})
    }
    loading.value = false
  })
}
loadMetadata()

</script>

<template>
  <div class="w-screen h-screen fixed" v-if="!loading">
    <a-config-provider :locale="arcoLocales[metadataStore.config.locale]" size="small" :scroll-to-close="true">
      <TitleBar />
      <div class="body-h">
        BODY
      </div>
    </a-config-provider>
  </div>
</template>
