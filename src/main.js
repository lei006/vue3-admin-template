import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import {createPinia} from 'pinia'
const pinia = createPinia()

import i18n from './i18n/index.js'  // 引入配置好的文件

import '@/permission' // permission control
import '@/styles/index.css' // global css


console.log("meta.env", import.meta.env);

const app = createApp(App);

app.use(router)
app.use(pinia)
//app.use(ElementPlus)
app.use(ElementPlus, { size: 'small', zIndex: 3000 })

app.use(i18n)

app.mount('#app')
