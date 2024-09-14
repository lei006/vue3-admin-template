<template>
    <div class="content-box">
      <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
        <el-tab-pane label="User" name="first">User</el-tab-pane>
        <el-tab-pane label="Config" name="second">Config</el-tab-pane>
        <el-tab-pane label="Role" name="third">Role</el-tab-pane>
        <el-tab-pane label="Task" name="fourth">Task</el-tab-pane>
      </el-tabs>
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
}


.statusName {
  margin-right: 10px;
}
</style>
