<template>
  <header>
      <div class="left-box">
        <!-- 收缩按钮 -->
        <div class="menu-icon" @click="opendStateChange">
          <el-icon v-if="layoutStore.isCollapse"><Menu /></el-icon>
          <el-icon v-if="!layoutStore.isCollapse"><Position /></el-icon>      
        </div>
        <Breadcrumb />
      </div>
      
      <div class="right-box">
        <!-- 快捷功能按钮 -->
        <div class="function-list">
          <div class="function-list-item hidden-sm-and-down"><el-icon><FullScreen /></el-icon></div>
          <div class="function-list-item"><el-icon><Suitcase /></el-icon></div>
          <div class="function-list-item"><el-icon><Umbrella /></el-icon></div>
          <div class="function-list-item hidden-sm-and-down"><el-icon><Magnet /></el-icon></div>
          <div class="function-list-item hidden-sm-and-down"><el-icon><Help /></el-icon></div>
        </div>
        <!-- 用户信息 -->
        <div class="user-info">
          <el-dropdown>
            <span class="el-dropdown-link">
              {{ $t('message.system.user') }}
              <i class="sfont system-xiala"></i>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="showPasswordLayer">{{ $t('message.system.changePassword') }}</el-dropdown-item>
                <el-dropdown-item @click="loginOut">{{ $t('message.system.loginOut') }}</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
        <password-layer :layer="layer" v-if="layer.show" />
      </div>      

  </header>
</template>

<script setup>
import { useRouter, useRoute } from 'vue-router'

import Breadcrumb from './Breadcrumb.vue'
import PasswordLayer from './passwordLayer.vue'
import { defineComponent, computed, reactive } from 'vue'

import LayoutStore from "@/pinia/layout.js"
const layoutStore = LayoutStore();


const layer = reactive({
      show: false,
      showButton: true
  });

    const router = useRouter()
    const route = useRoute()


    const opendStateChange = () => {
      console.log("route", route)
      layoutStore.changeCollapse();

    }
    
    // login out the system
    const loginOut = () => {
      //store.dispatch('user/loginOut')
    }
    
    const showPasswordLayer = () => {
      layer.show = true
    }

</script>

<style scoped>



  .header-box {
    width: 100%;
    height: 100%;
  }


  header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 60px;
    background-color: var(--system-header-background);
    padding-right: 22px;
    border-bottom: 1px solid #dcdfe6;    
  }
  .left-box {
    height: 100%;
    display: flex;
    align-items: center;

  }

  .menu-icon {
    width: 60px;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 25px;
    font-weight: 100;
    cursor: pointer;
    margin-right: 10px;
    &:hover {
      background-color: var(--system-header-item-hover-color);
    }
    i {
      color: var(--system-header-text-color);
    }
  }






  .right-box {
    display: flex;
    justify-content: center;
    align-items: center;


  }

  .user-info {
    margin-left: 20px;

  }
  .el-dropdown-link {
    color: var(--system-header-breadcrumb-text-color);
  }
  .function-list{
    display: flex;

  }

  .function-list-item {
    width: 30px;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .function-list-item :hover {
      color: #0FF;
      font-size:16px;
  }

  .head-fold {
    font-size: 20px;
  }
</style>