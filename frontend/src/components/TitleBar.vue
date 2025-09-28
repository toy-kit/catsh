<script lang="ts" setup>
import { ref } from "vue";
import { Window } from "@wailsio/runtime"
const props = defineProps({
    icon: { type: String, default: '' },
    title: { type: String, default: '' },
    minimise: { type: Boolean, default: true },
    maximize: { type: Boolean, default: true },
});

const isMaximised = ref(false);
const handleToggleMaximise = () => {
    Window.ToggleMaximise();
    Window.IsMaximised().then((value) => {
        isMaximised.value = value
    });
}
</script>
<template>
    <div class="flex items-center titlebar-h z-50 border-b border-color">
        <div v-if="props.icon" class="ml-2 mr-1.5 drag">
            <i class="iconfont text-sm" :class="props.icon"></i>
        </div>
        <div v-if="props.title" class="drag text-sm font-medium">{{ props.title }}</div>
        <div class="flex-1 h-full drag"></div>
        <slot></slot>
        <a-button v-if="minimise" @click="Window.Minimise" type='text' class="!px-5 titlebar-h" :title="$t('titlebar.minimise')">
            <template #icon>
                <i class="iconfont icon-windows-minimize text-xs"></i>
            </template>
        </a-button>
        <a-button v-if="maximize" @click="handleToggleMaximise" type='text' class="!px-5 titlebar-h" :title="$t('titlebar.maximize')">
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