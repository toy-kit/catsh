<script setup lang="ts">
import TitleBar from '../components/TitleBar.vue';
import { useAppDataStore } from '../store/appdata';
import { CheckAndInstall } from '../../bindings/catsh/service/upgradeservice';
import { Browser } from "@wailsio/runtime"
const appDataStore = useAppDataStore();
</script>

<template>
    <div class="w-screen h-screen fixed">
        <TitleBar icon="about" title="关于" :maximize="false" />
        <div class="body-h flex flex-col drag">
            <div class="flex-1 h-0 flex items-center justify-center">
                <div class="text-center no-drag">
                    <img src="../assets/logo.svg" class="w-17 h-17 mx-auto" />
                    <div class="text-base font-bold my-2">Catsh</div>
                </div>
            </div>
            <div class="flex items-center justify-center mb-6 no-drag">
                <a-form auto-label-width class="!w-auto">
                    <a-form-item label="版本" class="!mb-1.5">
                        <div class="flex items-center gap-6">
                            <div class="flex-1">{{ appDataStore.app_config.info.version }}</div>
                            <a-button size="mini" title="检测更新" @click="CheckAndInstall">
                                检测更新
                            </a-button>
                        </div>
                    </a-form-item>
                    <a-form-item label="repository" class="!mb-1.5">
                        <a-link status="normal" @click="Browser.OpenURL(appDataStore.app_config.repository.url)">
                            {{ appDataStore.app_config.repository.url }}
                        </a-link>
                    </a-form-item>
                    <a-form-item label="homepage" class="!mb-1.5">
                        <a-link status="normal" @click="Browser.OpenURL(appDataStore.app_config.repository.homepage)">
                            {{ appDataStore.app_config.repository.homepage }}
                        </a-link>
                    </a-form-item>
                    <a-form-item label="author" class="!mb-1.5">{{ appDataStore.app_config.repository.author }}</a-form-item>
                    <a-form-item label="email" class="!mb-1.5">{{ appDataStore.app_config.repository.email }}</a-form-item>
                </a-form>
            </div>
            <div class="h-12 flex items-center justify-center no-drag">
                <a-typography-text type="secondary">
                    <div class="text-xs">{{ appDataStore.app_config.info.copyright }}</div>
                </a-typography-text>
            </div>
        </div>
    </div>
</template>