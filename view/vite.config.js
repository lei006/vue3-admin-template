import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'




export default ({ command }) => {

  return defineConfig({
    plugins: [vue()],
    server:{
      open:true,
    },
    resolve:{   
      alias:{
        '@': resolve('./src')
      },
    },
    build:{
      //outDir: '../server/view',
    }    
  })
}
