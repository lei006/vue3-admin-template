<template>
  <div class="side-bar">
    <el-tabs :tab-position="tabPosition" style="height: 100%" class="demo-tabs">
        <el-tab-pane label="Logo">
            <template v-slot:label>
                <div class="tabs-button-box">
                    <el-image style="width: 54px; height: 54px" :src="logo_url" fit="scale-down" />
                </div>
            </template>
        </el-tab-pane>
        <el-tab-pane v-for="panel in panelArray" :key="panel" label="User">
            <template v-slot:label>
                <div class="tabs-button-box">
                    <div class="tabs-button-item" :class="{'is-active':panel.active}" @click="onItemClick(panel)">
                        <el-icon size="20px"><Apple /></el-icon>
                        {{panel.title}}
                    </div>
                </div>
            </template>

        </el-tab-pane>
    </el-tabs>
</div>
</template>

<script setup>
import { ref } from 'vue'

let logo_url = "https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg";

const tabPosition = ref('left')

let panelArray = ref([{title:"首页", icon:"Check", active:true},{title:"Page0", icon:"Check"},{title:"Page1", icon:"Check"},{title:"Page2", icon:"Check"},])

const onItemClick = (item)=>{

    panelArray.value.forEach((panel) => {
        panel.active = false;
    });

    item.active = true;
}


</script>

<style lang="scss" scoped>

.tabs-button-box{
    width: 64px;
    height: 64px;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items:center;

    margin:1px;
}

.tabs-button-item{
    width: 54px;
    height: 54px;
    color: #eee;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items:center;    

    font-size: 14px;
    border-radius: 3px;
}

.tabs-button-item:hover{
    background-color: rgba(24, 144, 255, 0.2);
}

.is-active {
    background-color: rgb(24, 144, 255);
}




.side-bar{
    height: 100%;
    width: 170px;

    background-color: #282a34;


    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items:flex-start;
}

.list-box{
    width:100%;
    height: 100%;
}

.list-item{
    width:100%;
    height: 40px;
    color: aquamarine;

    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items:center;
}

.list-item:hover {
    background-color: #708196;
}

</style>