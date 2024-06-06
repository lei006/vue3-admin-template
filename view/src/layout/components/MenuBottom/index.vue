<template> 
    <div class="menu-box box-shadow-top">
        <div class="btn-item" :class="activeMenu == '/sys_user/index'?'btn-item-active':'' " @click="jumpPath('/sys_user')">
            <div class="icon"><el-icon><User /></el-icon></div>
            <div class="text">用户</div>
        </div>
        <div class="btn-item" :class="activeMenu == '/setup/index'?'btn-item-active':'' " @click="jumpPath('/setup')">

            <div class="icon"><el-icon><Setting /></el-icon></div>
            <div class="text">设置</div>
        </div>
        <div class="btn-item" :class="activeMenu == '/about/index'?'btn-item-active':'' " @click="jumpPath('/about')">
            <div class="icon"><el-icon><Warning /></el-icon></div>
            <div class="text">关于</div>
        </div>
        {{ activeMenu }}
    </div>
</template>

<script setup>

import { defineComponent, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
const router = useRouter()


const route = useRoute()
const activeMenu = computed(() => {
    const { meta, path } = route;
    if (meta.activeMenu) {
      return meta.activeMenu;
    }
    return path;
});


const jumpPath = (path) => {
    router.push({path: path})
}



</script>


<style scoped>


.box-shadow-top {
    border-top-width: 1px;
    border-top-style: solid;
    border-top-color: #d0d7de20;
}



</style>