<template>
  <div class="aside-box ">

    <div class="aside-button-group">
      <SystemButton02 text="影像" @click="onToRouter1('/')"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="页面1" @click="onToRouter1('/page01/index')"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="页面2" @click="onToRouter1('/page02/index')"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="扩展" @click="onToRouter1('/extensions/index')"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="example" @click="onToRouter1('/example')"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="登录" @click="onToRouter1('/login')"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>

    </div>

    <div class="aside-button-ex-group">
      <el-divider></el-divider>
      <SystemButton02 text="影像" @click="onToRouter1('/')"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="页面1" @click="onToRouter1('/page01/index')"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="页面2" @click="onToRouter1('/testApi/index')"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="扩展" @click="onToRouter1('/extensions/index')"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
    </div>
  </div>
</template>

<script>
import Breadcrumb from '@/components/Breadcrumb'
import {SystemButton02} from '@/components/ButtonCustom'

const {ipcRenderer, shell} = require('electron')

export default {
  components: {
    Breadcrumb,SystemButton02
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
        //this.$router.push({name:'home'})
        //this.$router.push({path:path})
        //this.$router.push({ path: '/page01' })

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


.aside-box{
  width: 100%;
  height: 100%;

  padding-top:10px;
  padding-bottom: 20px;

  display: flex;
  flex-direction: column;
  justify-content:space-between;
  align-items: center;
}

.aside-button-group {
  flex:1;

  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
}

.aside-button-ex-group {
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  align-items: center;
}



</style>
