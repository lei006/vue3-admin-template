import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import {createPinia} from 'pinia'
const pinia = createPinia()


console.log("meta.env", import.meta.env);

const app = createApp(App);

app.use(router)
app.use(pinia)
app.use(ElementPlus)

app.mount('#app')
