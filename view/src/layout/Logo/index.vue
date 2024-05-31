<template>
  <div class="logo-container">
    <img class="img" src="@/assets/logo.png" alt="">
    <h1 >{{ $t(systemTitle) }}</h1>
  </div>
</template>

<script setup>

import {ref} from 'vue'

let isCollapse  = ref(true)

const systemTitle = ref("AAAAA")


/*
import { defineComponent, computed } from 'vue'
import { useStore } from 'vuex'
import { systemTitle } from '@/config'
export default defineComponent({
  setup() {
    const store = useStore()
    const isCollapse = computed(() => store.state.app.isCollapse)
    return {
      isCollapse,
      systemTitle
    }
  }
})

*/

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
    height: 36px;
    margin-right: 10px;
  }
</style>