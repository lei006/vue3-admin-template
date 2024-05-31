import { createApp } from 'vue'
import './style.css'
import App from './App.vue'


import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import i18n from './i18n'

import pinia from './pinia'
import router from './router'


console.log(' import.meta.env.MODE: ',  import.meta.env.MODE);
console.log(' import.meta.env.BASE_URL: ',  import.meta.env.BASE_URL);
console.log(' import.meta.env.PROD: ',  import.meta.env.PROD);
console.log(' import.meta.env.DEV: ',  import.meta.env.DEV);
console.log(' import.meta.env.SSR: ',  import.meta.env.SSR);
console.log(' import.meta.env.VITE_BASE_URL: ',  import.meta.env.VITE_BASE_URL);
console.log(' import.meta.env.TEST_11: ',  import.meta.env.TEST_11);




const app = createApp(App)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}


app.use(ElementPlus)
app.use(pinia)
app.use(router)
app.use(i18n)
app.mount('#app')

