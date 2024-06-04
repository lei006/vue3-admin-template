<template>
    <div class="menu-box">

        <div class="btn-item" @click="jumpPath('/users')">
            <div class="icon"><el-icon><User /></el-icon></div>
            <div class="text">用户</div>
        </div>
        <div class="btn-item" @click="jumpPath('/setup')">
            <div class="icon"><el-icon><Setting /></el-icon></div>
            <div class="text">设置</div>
        </div>
        <div class="btn-item" @click="jumpPath('/about')">
            <div class="icon"><el-icon><Warning /></el-icon></div>
            <div class="text">关于</div>
        </div>
    </div>
</template>

<script setup>

import { useRouter } from 'vue-router'
const router = useRouter()


const jumpPath = (path) => {
    router.push({path: path})
}



</script>


<style scoped>


</style>