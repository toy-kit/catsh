<script lang="ts" setup>
import Windows from './windows.vue';
import Linux from './linux.vue';
import Darwin from './darwin.vue';
import { useAppDataStore } from '../../store/appdata';
const appDataStore = useAppDataStore();
const props = defineProps(["title", "resize"]);
const productName = appDataStore.app_config.info.productName;
</script>
<template>
    <div>
        <Windows v-if="appDataStore.app_config.os == 'windows'" v-bind="props" :productName="productName"></Windows>
        <Linux v-if="appDataStore.app_config.os == 'linux'" v-bind="props" :productName="productName"></Linux>
        <Darwin v-if="appDataStore.app_config.os == 'darwin'" v-bind="props" :productName="productName"></Darwin>
    </div>
</template>