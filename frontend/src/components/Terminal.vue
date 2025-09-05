<script setup>
import { onMounted, watch } from 'vue';
import { Terminal } from '@xterm/xterm';
import "@xterm/xterm/css/xterm.css"
import { FitAddon } from '@xterm/addon-fit';
import { WebLinksAddon } from '@xterm/addon-web-links';
import { useMetadataStore } from '../store/metadata'

const metadataStore = useMetadataStore()

let term = null;
let fitAddon = null;
const getTheme = () => {
    const style = getComputedStyle(document.documentElement)
    return {
        foreground: style.getPropertyValue("--color-text-1"),
        background: style.getPropertyValue("--color-bg-1"),
    }
}
const open = () => {
    term = new Terminal({
        convertEol: true,
        cursorBlink: true,
        disableStdin: true,
        theme: getTheme()
    });
    fitAddon = new FitAddon();
    term.loadAddon(fitAddon);
    term.loadAddon(new WebLinksAddon());
    term.open(document.getElementById('terminal'));
    fit();
}
const write = (value) => {
    term.write(value)
}
const fit = () => {
    fitAddon.fit();
}
watch(() => metadataStore.config.theme, () => {
    term.options.theme = getTheme()
})
onMounted(() => {
    window.addEventListener('resize', () => {
        fit()
    })
    open()
})
defineExpose({ open, write, fit })
</script>

<template>
    <div id="terminal" class="w-full h-full"></div>
</template>
