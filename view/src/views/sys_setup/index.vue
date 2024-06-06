<template>
  <div class="layout-container">

    <div class="setup-items-box">
      <div class="item" v-for="item in items">
        <div class="title">{{ item.name }}</div>
        <div class="data">{{ item.data }}</div>
      </div>
      <hr>
      <div class="item" v-for="item in items">
        <div class="title">{{ item.name }}</div>
        <div class="data">{{ item.data }}</div>
      </div>
      <hr>
      <div class="item" v-for="item in items">
        <div class="title">{{ item.name }}</div>
        <div class="data">{{ item.data }}</div>
      </div>
      
    </div>

  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from "vue"

import apiSetup from "@/api/setup"


const items = ref([])


onMounted(async ()=>{

  const ret = await apiSetup.GetList().then((ret)=>{

    items.value = ret.data.items;
    console.log(items.value)

    return ret
  })

  
})



</script>

<style lang="scss" scoped>

.setup-items-box {
  padding: 80px;
}

.setup-items-box .item {
  width:  600px;
  display: flex;
  flex-direction: row;

  .title {
    width:  150px;
    height: 30px;
    padding-left: 20px;
    background-color: beige;
  }

  .data {
    width:  150px;
    height: 30px;
  }


}


</style>
