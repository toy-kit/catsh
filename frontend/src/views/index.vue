<script setup lang="ts">
import { ToSetting, ToAbout } from '../router/window';
import TitleBar from '../components/TitleBar.vue';
const handleSelect = (key: string) => {
    switch (key) {
        case 'setting':
            ToSetting();
            break;
        case 'about':
            ToAbout();
            break;
        case 'saveAs':
            break;
    }
};
</script>

<template>
    <div class="w-screen h-screen fixed">
        <TitleBar>
            <a-dropdown @select="handleSelect" position="bl">
                <a-button type='text' class="!px-5 titlebar-h" :title="$t('titlebar.minimise')">
                    <template #icon>
                        <i class="iconfont icon-menu text-xs"></i>
                    </template>
                </a-button>
                <template #content>
                    <a-doption value="light">{{ $t('titlebar.setting') }}</a-doption>
                    <a-divider class="!my-1" />
                    <a-doption value="setting">{{ $t('titlebar.setting') }}</a-doption>
                    <a-doption value="about">{{ $t('titlebar.about') }}</a-doption>
                </template>
            </a-dropdown>
        </TitleBar>
        <a-scrollbar class="overflow-y-auto body-h">
            <div v-for="item in 100" :key="item" class="m-4">
                <a-card>{{ item }}</a-card>
            </div>
        </a-scrollbar>
    </div>
</template>

<style scoped></style>
