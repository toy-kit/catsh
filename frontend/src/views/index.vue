<script setup lang="ts">
import { ref } from 'vue';
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
const tabsActive = ref('home');
const data = ref([
    {
        key: '1',
        title: 'Tab 1',
        content: 'Content of Tab Panel 1'
    },
    {
        key: '2',
        title: 'Tab 2',
        content: 'Content of Tab Panel 2'
    },
    {
        key: '3',
        title: 'Tab 3',
        content: 'Content of Tab Panel 3'
    },
    {
        key: '4',
        title: 'Tab 4',
        content: 'Content of Tab Panel 4'
    }
]);
const handleAdd = () => {
    let number = data.value.length + 1;
    data.value = data.value.concat({
        key: `${number}`,
        title: `New Tab ${number}`,
        content: `Content of New Tab Panel ${number}`
    })
};
const handleDelete = (key: any) => {
    data.value = data.value.filter(item => item.key !== key)
};
</script>

<template>
    <div class="w-screen h-screen fixed">
        <TitleBar></TitleBar>
        <a-tabs v-model="tabsActive" class="body-tabs" type="card" size="medium" :editable="true" @add="handleAdd" @delete="handleDelete" show-add-button auto-switch>
            <template #extra>
                <a-dropdown @select="handleSelect" position="bl">
                    <a-button type='text' class="!px-5 titlebar-h" :title="$t('titlebar.minimise')">
                        <template #icon>
                            <i class="iconfont icon-menu text-xs"></i>
                        </template>
                    </a-button>
                    <template #content>
                        <a-doption value="setting">{{ $t('titlebar.setting') }}</a-doption>
                        <a-divider class="!my-1" />
                        <a-doption value="about">{{ $t('titlebar.about') }}</a-doption>
                    </template>
                </a-dropdown>
            </template>
            <a-tab-pane key="home" :closable="false">
                <template #title>
                    <i class="iconfont icon-menu mr-1"></i>首页
                </template>
                <a-scrollbar class="overflow-y-auto body-h">
                    <div v-for="item in 100" :key="item" class="m-4">
                        <a-card>{{ item }}</a-card>
                    </div>
                </a-scrollbar>
            </a-tab-pane>
            <a-tab-pane key="ssh" title="SSH">
                Terminal
            </a-tab-pane>
            <a-tab-pane key="ftp" title="FTP">
                FTP
            </a-tab-pane>
            <a-tab-pane key="sql" title="SQL">
                SQL
            </a-tab-pane>
        </a-tabs>
    </div>
</template>

<style scoped></style>
