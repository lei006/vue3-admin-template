<template>
  <div class="app-wrapper" ref="tmp">
    <div class="sidebar-container">
        <Sidebar></Sidebar>
    </div>
    <div class="main-container">
        <div class="navbar-container">
            <Navbar></Navbar>
        </div>
        <div class="appmain-container">
            <AppMain></AppMain>
        </div>
    </div>
  </div>


</template>

<script setup>

import {AppMain,Navbar,Sidebar} from './components/index'



let obj = reactive({a:1,b:2,test:"awxxxxxxxxxxxx"});

let {test} = toRefs(obj);

let msg = ref("这是测试文本")

let onSet = ()=>{
    obj.a += 100;
    msg.value = "文本有改变";
    console.log(obj, msg)
    test.value += "-";
}
//return {obj, onSet};

onMounted(()=>{
    console.log("onMounted")
})

let msgChange = computed(()=>{
    return msg.value + "后坠" + test.value;
})

</script>


<style lang="scss" scoped>
  @import "../styles/variables.scss";

.app-wrapper{
    position: relative;
    height: 100%;
    width: 100%;

    display: flex;
    flex-direction: row;
}

.sidebar-container{

    height: 100%;
}

.main-container{
    flex:1;
    height: 100%;

    display: flex;
    flex-direction: column;
}

.navbar-container{
    width: 100%;
}

.appmain-container{
    width: 100%;
    flex:1;

    overflow: hidden;
}


</style>