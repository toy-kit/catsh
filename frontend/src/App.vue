<script setup lang="ts">
import { ref } from 'vue';
import zhCN from '@arco-design/web-vue/es/locale/lang/zh-CN';
import enUS from '@arco-design/web-vue/es/locale/lang/en-us';
import { useAppDataStore } from './store/appdata';
import { Load } from '../bindings/catsh/service/appdataservice';
import i18n from './utils/locales'

const arcoLocales: any = {
  'zh': zhCN,
  'en': enUS,
};
const appDataStore = useAppDataStore();
const loading = ref(true)

Load().then((data) => {
  appDataStore.setConfig(data.config);
  appDataStore.setAppConfig(data.app_config);
  appDataStore.setWindowOptions(data.window_options);
  const locales: any = data.locales || {}
    for (const key in locales) {
      i18n.global.setLocaleMessage(key, locales[key] || {})
    }
    loading.value = false
});

</script>

<template>
  <div v-if="!loading">
    <a-config-provider :locale="arcoLocales[appDataStore.config.locale]" size="small" :scroll-to-close="true">
      <router-view></router-view>
    </a-config-provider>
  </div>
</template>

<style scoped>
</style>
