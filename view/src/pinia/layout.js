import { defineStore } from 'pinia'
 
 
// 定义并导出容器，第一个参数是容器id，必须唯一，用来将所有的容器
// 挂载到根容器上
let LayoutStore = defineStore({
  id:'LayoutStore',
  // 定义state，用来存储状态的
  state: () => ({
    showLogo:true,
    isCollapse:false,
    showTabs:false,
    activeMenu:true,
    isBackMenu:true,
  }),

  actions:{
    changeCollapse(){
      
      if (this.isCollapse == true) {
        this.isCollapse = false;
      }else{
        this.isCollapse = true;
      }
      

      //this.isCollapse != this.isCollapse;

      console.log("isCollapse", this.isCollapse);
    },
  }
})


export default LayoutStore



