<script lang="ts" setup>
import { ref } from "vue";
import { Window } from "@wailsio/runtime"
import { useAppDataStore } from '../store/appdata';
const appDataStore = useAppDataStore();

const isMaximised = ref(false);
const handleToggleMaximise = () => {
    Window.ToggleMaximise();
    Window.IsMaximised().then((value) => {
        isMaximised.value = value
    });
}
</script>
<template>
    <div class="flex items-center titlebar-h z-50">
        <div class="ml-2 mr-1.5 drag">
            <img src="../assets/logo.svg" class="w-4.5 h-4.5" />
        </div>
        <div class="drag text-sm font-medium">{{ appDataStore.window_options.title }}</div>
        <div class="flex-1 h-full drag"></div>
        <slot></slot>
        <a-button @click="Window.Minimise" type='text' class="!px-5 titlebar-h" :title="$t('titlebar.minimise')">
            <template #icon>
                <i class="iconfont icon-windows-minimize text-xs"></i>
            </template>
        </a-button>
        <a-button v-if="appDataStore.window_options.resizable" @click="handleToggleMaximise" type='text' class="!px-5 titlebar-h" :title="$t('titlebar.maximize')">
            <template #icon>
                <i class="iconfont text-xs" :class="isMaximised ? 'icon-windows-maximize' : 'icon-windows-unmaximize'"></i>
            </template>
        </a-button>
        <a-button @click="Window.Close" type='text' class="!px-5 titlebar-h hover:!bg-red-500 hover:!text-white" :title="$t('titlebar.close')">
            <template #icon>
                <i class="iconfont icon-windows-close text-xs"></i>
            </template>
        </a-button>
    </div>
</template>