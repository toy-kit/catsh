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
        <div class="ml-2 mr-1.5">
            <img src="../assets/images/logo.svg" class="w-4.5 h-4.5" />
        </div>
        <a-dropdown @select="handleSelect" position="bl">
            <a-button type='text' class="!px-2">{{ $t('titlebar.file') }}</a-button>
            <template #content>
                <a-doption>{{ $t('titlebar.newFile') }}</a-doption>
                <a-divider class="!my-1" />
                <a-doption>{{ $t('titlebar.save') }}</a-doption>
                <a-doption>{{ $t('titlebar.saveAs') }}</a-doption>
            </template>
        </a-dropdown>
        <a-dropdown @select="metadataStore.changeTheme" position="bl">
            <a-button type='text' class="!px-2">{{ $t('titlebar.theme') }}</a-button>
            <template #content>
                <a-doption value="light">{{ $t('titlebar.light') }}</a-doption>
                <a-doption value="dark">{{ $t('titlebar.dark') }}</a-doption>
            </template>
        </a-dropdown>
        <a-dropdown @select="metadataStore.changeLocale" position="bl">
            <a-button type='text' class="!px-2">{{ $t('titlebar.language') }}</a-button>
            <template #content>
                <a-doption value="zh">中文</a-doption>
                <a-doption value="en">English</a-doption>
            </template>
        </a-dropdown>
        <a-dropdown @select="handleHelpSelect" position="bl">
            <a-button type='text' class="!px-2">{{ $t('titlebar.help') }}</a-button>
            <template #content>
                <a-doption value="docs">{{ $t('titlebar.docs') }}</a-doption>
                <a-doption value="repository">{{ $t('titlebar.repository') }}</a-doption>
                <a-divider class="!my-1" />
                <a-doption value="upgrade">{{ downloading ? $t('titlebar.downloading') : $t('titlebar.upgrade') }}</a-doption>
                <a-divider class="!my-1" />
                <a-doption value="about">{{ $t('titlebar.about') }}</a-doption>
            </template>
        </a-dropdown>
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
