<template>
  <div class="navbar-box text-no-select">

    <div class="navbar-logo-title drag-win-bar">
      <div class="logo-box">
        <el-image style="width: 36px; height: 36px; margin-left:8px; margin-right:8px;" :src="logo_url" :fit="logo_fit"></el-image>报告工作站---云创电子
      </div>
    </div>

    <div class="navbar-system-control">
      <el-breadcrumb class="app-breadcrumb" separator="/">
        <BtnIconText text="设置" @click="onToRouter1('/page01/index')"><span class="icon iconfont icon-un-setup-o"></span></BtnIconText>
        <BtnIconText text="登录" @click="onToRouter1('/page02/index')"><span class="icon iconfont icon-un-setup-o"></span></BtnIconText>

        <SystemButton00 @click="onBtnMini">─</SystemButton00>
        <SystemButton00 @click="onBtnExit">✖</SystemButton00>
      </el-breadcrumb>
    </div>
  </div>
</template>

<script>
import Breadcrumb from '@/components/Breadcrumb'
import {SystemButton00, SystemButton01,BtnIconText} from '@/components/ButtonCustom'

const {ipcRenderer, shell} = require('electron')

export default {
  components: {
    Breadcrumb,SystemButton00,SystemButton01,BtnIconText
  },

  data() {
    return {
      userauth:false,
      userinfo: {},  //用户信息....
      userinfoShow:false,
      logo_url: require("@/assets/logo.png"),
      logo_fit:"cover",

      user_url: require("@/assets/img_navbar/user.png"),

    }
  },

  async mounted(){
    let _self = this;

  },

  methods: {

    onMenuClick(menu_cmd) {

    },

    onBtnLogin() {
      this.$emit("menuclick","login");
    },

    onBtnMini() {
      ipcRenderer.send('app-mini') // prints "pong"
    },
    onBtnExit() {
      ipcRenderer.send('app-exit') // prints "pong"
    },
    onToRouter1(path){
        this.$router.push({path})
    }

  }
}
</script>

<style lang="scss" scoped>


.logo-box {
  display: flex;
  justify-content: center;
  align-items: center;

  font-weight: 500;
  font-size: 20px;
}


.navbar-box {
  width: 100%;
  height: 100%;

  display: flex;
  justify-content: flex-end;
}

.navbar-logo-title {
  flex:1;
  height: 100%;

  display: flex;
  justify-content: flex-start;
}

.navbar-system-control {
  height: 100%;
}


.app-breadcrumb.el-breadcrumb {
  font-size: 14px;
  height: 100%;

  padding-left: 5px;
  padding-right: 5px;
  margin-left: 0px;
  margin-right: 0px;

  display: flex;
  justify-content: center;
  align-items: center;  
}



</style>
