<template>
    <div class="content-box">
        
        <h3>关于我们</h3>
        <p>欢迎访问我们的网站！这里是我们项目的详细介绍。</p>
        <p>这是一个基于Vue.js和TypeScript构建的应用程序，旨在...</p>
        <p v-for="item in about_items">
            {{ item.title }}
            {{ item.data }}
            {{ item.desc }}
        </p>
        <p><el-button>导入授权文件</el-button></p>
        
    </div>
</template>

<script setup>
import { onMounted, ref, reactive } from "vue";
import apiAboutOption from "@/api/system/about";

const activeName = ref('first')

let about_items = ref([])


const handleClick = (tab, event) => {
console.log(tab, event)
}


onMounted(()=>{

    apiAboutOption.GetList().then((ret)=>{
        console.log(ret)
        about_items.value = ret.data.items
    })

})


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
