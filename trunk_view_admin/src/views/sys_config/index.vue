<template>
    <div class="layout-box">
        系统配置
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref, reactive } from "vue";

onMounted(()=>{

})



</script>

<style lang="scss" scoped>


.layout-box {
    background-color: var(--system-container-main-background);
    width: calc(100% - 10px);
    height: calc(100% - 10px);
    margin: 15px;
    display: flex;
    flex-direction: column;
    overflow-y: auto;
}



</style>
