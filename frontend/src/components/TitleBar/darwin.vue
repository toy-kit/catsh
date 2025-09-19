<script lang="ts" setup>

const props = defineProps(["title", "resize", "productName"]);

</script>
<template>
    <div>
        <div></div>
        <div>{{ props.title || props.productName }}</div>
    </div>
</template>