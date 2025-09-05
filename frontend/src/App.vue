<script lang="ts" setup>
import { ref } from 'vue'
import zhCN from '@arco-design/web-vue/es/locale/lang/zh-CN';
import enUS from '@arco-design/web-vue/es/locale/lang/en-us';
import { useMetadataStore } from './store/metadata'
import { Load } from "../wailsjs/go/main/App"
import i18n from './utils/locales'
import Terminal from './components/Terminal.vue';
import Setting from './components/Setting.vue';
import Connect from './components/Connect.vue';
import Snippet from './components/Snippet.vue';
import Task from './components/Task.vue';
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
      <a-tabs class="body-h" default-active-key="connect" position="left" direction="vertical">
        <a-tab-pane key="terminal" title="终端">
          <Terminal></Terminal>
        </a-tab-pane>
        <a-tab-pane key="connect" title="连接">
          <Connect></Connect>
        </a-tab-pane>
        <a-tab-pane key="snippet" title="片段">
          <Snippet></Snippet>
        </a-tab-pane>
        <a-tab-pane key="task" title="任务">
          <Task></Task>
        </a-tab-pane>
        <template #extra>
          <Setting></Setting>
        </template>
      </a-tabs>
      <div class="border-t border-color h-7"></div>
    </a-config-provider>
  </div>
</template>
