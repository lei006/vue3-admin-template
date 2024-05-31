<template>
      <el-button>Default</el-button>
      <el-table mb-1 :data="[]" />
      <el-pagination :total="100" />
      <div class="flex flex-wrap items-center mb-4">
        <el-button size="large">Large</el-button>
        <el-button @click="onBtnClick">Default{{ count }}</el-button>
        <el-button size="small">Small</el-button>
        <el-button size="large" >
          <el-icon ><Edit /></el-icon>
        </el-button>
      </div>
</template>

<script setup>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'

import {ref } from 'vue'

let count = ref(0)


const onBtnClick = () => {
  count.value = count.value + 1
  console.log('click', count.value)
}

</script>
