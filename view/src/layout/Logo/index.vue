<template>
  <div class="logo-container">
    <img class="img" src="@/assets/logo.png" alt="">
    <h1 style="margin-left:15px;" v-if="!layoutStore.isCollapse">我是系统</h1>
  </div>
</template>

<script setup>

import {ref} from 'vue'

import LayoutStore from "@/pinia/layout.js"
const layoutStore = LayoutStore();


const systemTitle = ref("AAAAA")





</script>

<style  scoped>
  .logo-container {
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    background-color: var(--system-logo-background);
    color:#f1f1f1;
  }

  .logo-container .img{
    height: 28px;
  }
</style>