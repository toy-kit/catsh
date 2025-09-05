<script lang="ts" setup>
import { ref } from 'vue'
import { useMetadataStore } from '../store/metadata'
import { OpenAbout, CheckUpgrade } from '../../wailsjs/go/main/App'
const metadataStore = useMetadataStore()
const runtime = (window as any).runtime
const isMaximised = ref(false)
const getIsMaximised = () => {
    runtime.WindowIsMaximised().then((ok: boolean) => {
        isMaximised.value = ok
    })
}
getIsMaximised()
const handleToggleMaximise = () => {
    runtime.WindowToggleMaximise()
    isMaximised.value = !isMaximised.value
}
const handleSelect = (value: string) => {
    console.log(value)
}
const downloading = ref(false)
const handleHelpSelect = (value: string) => {
    if (value === 'docs') {
        runtime.BrowserOpenURL(metadataStore.wails_config.homepage)
    } else if (value === 'repository') {
        runtime.BrowserOpenURL(metadataStore.wails_config.repository)
    } else if (value === 'upgrade') {
        if (downloading.value) return
        downloading.value = true
        CheckUpgrade().finally(() => {
            downloading.value = false
        })
    } else if (value === 'about') {
        OpenAbout()
    }
}
</script>

<template>
    <div class="flex items-center titlebar-h z-50">
        <div class="ml-2 mr-1.5 drag">
            <img src="../assets/images/logo.svg" class="w-5.5 h-5.5" />
        </div>
        <div class="drag text-sm font-medium">{{ metadataStore.wails_config.name }}</div>
        <div class="flex-1 h-full drag"></div>
        <a-button @click="runtime.WindowMinimise()" type='text' class="!px-5 titlebar-h" :title="$t('titlebar.minimise')">
            <template #icon>
                <i class="iconfont icon-windows-minimize text-xs"></i>
            </template>
        </a-button>
        <a-button @click="handleToggleMaximise" type='text' class="!px-5 titlebar-h" :title="$t('titlebar.maximize')">
            <template #icon>
                <i class="iconfont text-xs" :class="isMaximised ? 'icon-windows-maximize' : 'icon-windows-unmaximize'"></i>
            </template>
        </a-button>
        <a-button @click="runtime.Quit()" type='text' class="!px-5 titlebar-h hover:!bg-red-500 hover:!text-white" :title="$t('titlebar.close')">
            <template #icon>
                <i class="iconfont icon-windows-close text-xs"></i>
            </template>
        </a-button>
    </div>
</template>
