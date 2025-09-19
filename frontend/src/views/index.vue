<script setup lang="ts">
import { ToSetting } from '../router/window';
import TitleBar from '../components/TitleBar.vue';
</script>

<template>
    <div class="w-screen h-screen fixed">
        <TitleBar>
            <a-button @click="ToSetting" type='text' class="!px-5 titlebar-h" :title="$t('titlebar.minimise')">
                <template #icon>
                    <i class="iconfont icon-menu text-xs"></i>
                </template>
            </a-button>
        </TitleBar>
        <a-scrollbar class="overflow-y-auto body-h">
            <div v-for="item in 100" :key="item" class="m-4">
                <a-card>{{ item }}</a-card>
            </div>
        </a-scrollbar>
    </div>
</template>

<style scoped></style>
