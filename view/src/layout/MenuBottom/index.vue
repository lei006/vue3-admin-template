<template>
    
      <el-menu
        class="el-menu-vertical-demo"
        text-color="#fff"
        :collapse="layoutStore.isCollapse"
        active-text-color="#ffd04b"
        background-color="#545c64"        
        @open="handleOpen"
        @close="handleClose"
      >
        <el-menu-item index="2">
          <el-icon><icon-menu /></el-icon>
          <span>Navigator Two</span>
        </el-menu-item>
        <el-menu-item index="3" disabled>
          <el-icon><document /></el-icon>
          <span>Navigator Three</span>
        </el-menu-item>
        <el-menu-item index="4">
          <el-icon><setting /></el-icon>
          <span>Navigator Four</span>
        </el-menu-item>
      </el-menu>

</template>


<script setup>

import { defineComponent, computed, onMounted,ref } from 'vue'

import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'

import LayoutStore from "@/pinia/layout.js"
const layoutStore = LayoutStore();


const handleOpen = (key, keyPath) => {
  console.log(key, keyPath)
}
const handleClose = (key, keyPath) => {
  console.log(key, keyPath)
}



</script>