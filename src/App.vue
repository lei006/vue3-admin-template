<script setup>
import HelloWorld from './components/HelloWorld.vue'


import { computed, ref } from 'vue'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import en from 'element-plus/dist/locale/en.mjs'

const language = ref('zh-cn')
const locale = computed(() => (language.value === 'zh-cn' ? zhCn : en))
let size = ref('small')

const toggle = () => {
  language.value = language.value === 'zh-cn' ? 'en' : 'zh-cn'
}

</script>

<template>

  <el-config-provider :locale="locale" :size="size">
      <el-button mb-2 @click="toggle">Switch Language</el-button>
      <el-button @click="size='small'">small</el-button>
      <el-button @click="size='large'">large</el-button>
      <el-button @click="size='default'">default</el-button>
      <el-button>这是应用</el-button>

      <!--
      <el-table mb-1 :data="[]" />
      <el-pagination :total="100" />
      <div class="flex flex-wrap items-center mb-4">
        <el-button size="large">Large</el-button>
        <el-button>Default</el-button>
        <el-button size="small">Small</el-button>
        <el-button size="large" :icon="Search">Search</el-button>
        <el-button :icon="Search">Search</el-button>
        <el-button size="small" :icon="Search">Search</el-button>
      </div>
      -->
      <router-view></router-view>

  </el-config-provider>



</template>

<style scoped>

</style>
