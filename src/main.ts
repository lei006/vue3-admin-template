import { createApp } from "vue";

import App from './App.vue'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import './styles/index.scss' // global css

import router from "./router";
import store from "./store";

let app = createApp(App)



// 导入 element plus 的全部图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.use(router)
app.use(store)
app.use(ElementPlus);

app.mount("#app");
