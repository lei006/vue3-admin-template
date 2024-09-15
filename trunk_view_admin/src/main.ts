/*
 * @Date: 2022-05-22 20:44:25
 * @Description: 
 */
import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import { baidu } from './utils/system/statistics'
import 'element-plus/theme-chalk/display.css' // 引入基于断点的隐藏类
import 'element-plus/dist/index.css'
import 'normalize.css' // css初始化
import './assets/style/common.scss' // 公共css
import './theme/modules/chinese/index.scss'
import App from './App.vue'
import store from './store'
import router from './router'
import { getAuthRoutes } from './router/permission'
import i18n from './locale'
if (import.meta.env.MODE !== 'development') { // 非开发环境调用百度统计
  baidu()
}

/** 权限路由处理主方法 */
getAuthRoutes().then(() => {
  const app = createApp(App)
  app.use(ElementPlus, { size: store.state.app.elementSize })
  app.use(store)
  app.use(router)
  app.use(i18n)
  // app.config.performance = true
  app.mount('#app')
})



Date.prototype.Format = function (fmt) { //author: meizz 
  var o = {
      "M+": this.getMonth() + 1, //月份 
      "d+": this.getDate(), //日 
      "h+": this.getHours(), //小时 
      "m+": this.getMinutes(), //分 
      "s+": this.getSeconds(), //秒 
      "q+": Math.floor((this.getMonth() + 3) / 3), //季度 
      "S": this.getMilliseconds() //毫秒 
  };
  if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
  for (var k in o)
  if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
  return fmt;
}



function getQueryVariable(variable){
    let hash = window.location.hash;
    var arr_l = hash.split("?");
    if (arr_l.length < 2) {
        return false;
    }
    var query = arr_l[1];
    var vars = query.split("&");
    for (var i=0;i<vars.length;i++) {
        var pair = vars[i].split("=");
        if(pair[0] == variable)
        {
            return pair[1];
        }
    }

    return(false);
}

