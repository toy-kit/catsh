<script setup lang="ts">
import TitleBar from '../components/TitleBar.vue';
</script>

<template>
  <div class="w-screen h-screen fixed">
    <TitleBar />
    <a-scrollbar class="overflow-y-auto body-h py-3">
      <a-list class="mx-4 mb-5">
        <a-list-item>
          <a-list-item-meta :title="$t('titlebar.theme')">
            <template #avatar>
              <i class="iconfont icon-keyboard text-xl"></i>
            </template>
          </a-list-item-meta>
          <template #actions>
            <a-select size="mini">
              <a-option value="light">{{ $t('titlebar.light') }}</a-option>
              <a-option value="dark">{{ $t('titlebar.dark') }}</a-option>
            </a-select>
          </template>
        </a-list-item>
        <a-list-item>
          <a-list-item-meta :title="$t('titlebar.language')">
            <template #avatar>
              <i class="iconfont icon-keyboard text-xl"></i>
            </template>
          </a-list-item-meta>
          <template #actions>
            <a-select size="mini">
              <a-option value="zh">中文</a-option>
              <a-option value="en">English</a-option>
            </a-select>
          </template>
        </a-list-item>
      </a-list>
      <a-list class="mx-4 mb-5">
        <a-list-item>
          <a-list-item-meta title="开机时自动启动">
            <template #avatar>
              <i class="iconfont icon-keyboard text-xl"></i>
            </template>
          </a-list-item-meta>
          <template #actions>
            <a-switch type="round" />
          </template>
        </a-list-item>
        <a-list-item>
          <a-list-item-meta title="主窗口总在最前">
            <template #avatar>
              <i class="iconfont icon-keyboard text-xl"></i>
            </template>
          </a-list-item-meta>
          <template #actions>
            <a-switch type="round" />
          </template>
        </a-list-item>
      </a-list>
      <a-list class="mx-4 mb-5">
        <a-list-item>
          <a-list-item-meta title="更新源">
            <template #avatar>
              <i class="iconfont icon-keyboard text-xl"></i>
            </template>
          </a-list-item-meta>
          <template #actions>
            <a-select size="mini">
              <a-option value="github">GitHub</a-option>
              <a-option value="gitee">Gitee</a-option>
            </a-select>
          </template>
        </a-list-item>
        <a-list-item>
          <a-list-item-meta title="自动更新">
            <template #avatar>
              <i class="iconfont icon-keyboard text-xl"></i>
            </template>
          </a-list-item-meta>
          <template #actions>
            <a-switch type="round" />
          </template>
        </a-list-item>
      </a-list>
    </a-scrollbar>
  </div>
</template>

<style scoped></style>
