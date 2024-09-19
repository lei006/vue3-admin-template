<template>
    <div class="content-box">
        
        <h3>关于我们</h3>
        <p>欢迎访问我们的网站！这里是我们项目的详细介绍。</p>
        <p>这是一个基于Vue.js和TypeScript构建的应用程序，旨在...</p>
        <p v-for="item in about_items">
            {{ item.title }}:
            {{ item.data }}
            {{ item.desc }}
        </p>
        <p><el-button @click="handleClick">导入授权文件</el-button></p>

    </div>
</template>

<script setup>
import { onMounted, ref, reactive } from "vue";
import apiAboutOption from "@/api/system/about";

const activeName = ref('first')

let about_items = ref([])


const handleClick = (tab, event) => {
    readFile()
}


onMounted(()=>{


    refreshData()

})



const refreshData = () => {
    apiAboutOption.GetList().then((ret)=>{
        console.log(ret)
        about_items.value = ret.data.items
    })
}



// 优化后的代码
function readFile() {
    // 创建文件输入框
    const fileInput = document.createElement('input');
    fileInput.setAttribute('type', 'file');
    fileInput.setAttribute('style', 'display:none'); // 使用CSS隐藏
    document.body.appendChild(fileInput);
    fileInput.click();

    const reader = new FileReader();
    
    // 文件读取成功后的处理
    reader.onload = async function (event) {
        const fileContent = window.encodeURIComponent(event.target.result);
        const base64Content = btoa(fileContent);
        apiAboutOption.SetLicense(base64Content).then((res)=>{
            refreshData()
        })
    };

    // 文件读取失败后的处理
    reader.onerror = function (error) {
        console.error('文件读取失败：', error);
        document.body.removeChild(fileInput);
    };

    // 用户取消文件选择后的处理
    fileInput.onchange = function () {
        if (this.files && this.files.length > 0) {
            reader.readAsText(this.files[0]);
        } else {
            console.warn('用户取消了文件选择');
        }
        document.body.removeChild(fileInput);
    };
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
