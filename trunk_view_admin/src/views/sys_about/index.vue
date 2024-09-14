<template>
    <div class="content-box">
        
        <h3>关于我们</h3>
        <p>欢迎访问我们的网站！这里是我们项目的详细介绍。</p>
        <p>这是一个基于Vue.js和TypeScript构建的应用程序，旨在...</p>


    </div>
</template>

<script setup>
import { onMounted, ref, reactive } from "vue";
import apiSystemOption from "@/api/system/option";

const activeName = ref('first')

const handleClick = (tab, event) => {
console.log(tab, event)
}


</script>

<style lang="scss" scoped>

.content-box {
    margin-top: 25px;
    margin-left: auto;
    margin-right: auto;
    width: 800px;
    height: 640px;
    padding: 25px;
    background-color: #fff;

    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}


</style>
