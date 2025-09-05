<script lang="ts" setup>
import { ref } from 'vue'
import { useMetadataStore } from '../store/metadata'
import { OpenAbout, CheckUpgrade } from '../../wailsjs/go/main/App'
const metadataStore = useMetadataStore()
const runtime = (window as any).runtime
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
    <div class="flex items-center justify-center flex-col mb-2 gap-2 w-full">
        <a-dropdown @select="metadataStore.changeTheme" position="rb">
            <a-button type='text' class="!px-2">{{ $t('titlebar.theme') }}</a-button>
            <template #content>
                <a-doption value="light">{{ $t('titlebar.light') }}</a-doption>
                <a-doption value="dark">{{ $t('titlebar.dark') }}</a-doption>
            </template>
        </a-dropdown>
        <a-dropdown @select="metadataStore.changeLocale" position="rb">
            <a-button type='text' class="!px-2">{{ $t('titlebar.language') }}</a-button>
            <template #content>
                <a-doption value="zh">中文</a-doption>
                <a-doption value="en">English</a-doption>
            </template>
        </a-dropdown>
        <a-dropdown @select="handleHelpSelect" position="rb">
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
    </div>
</template>
