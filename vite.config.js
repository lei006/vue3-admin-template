import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'


import * as dotenv from 'dotenv'
import * as fs from 'fs'

export default ({mode}) => {

  const NODE_ENV = mode || 'development'
  const envFiles = [
    `.env.${NODE_ENV}`
  ]
  for (const file of envFiles) {
    const envConfig = dotenv.parse(fs.readFileSync(file))
    for (const k in envConfig) {
      process.env[k] = envConfig[k]
    }
  }

  return defineConfig({
    plugins: [vue()],
    //base:loadEnv(mode, process.cwd()).VITE_APP_NAME,
    server: {               
      open: true
    },  
    resolve:{   
      alias:{
        '@': resolve('./src')
      },
    },
    build:{
      outDir: './dist',
    }    
  })
}
