/*
 * @Author: luoxi
 * @Date: 2022-01-25 09:51:12
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2023-02-21 22:57:42
 * @FilePath: \vue-admin-box\vite.config.ts
 * @Description: 
 */
import { ConfigEnv, UserConfigExport } from 'vite'
import vue from '@vitejs/plugin-vue'
import { viteMockServe } from 'vite-plugin-mock'
import {vitePluginSvg} from "@webxrd/vite-plugin-svg"
import { resolve } from 'path'

const pathResolve = (dir: string): any => {
  return resolve(__dirname, ".", dir)
}

const alias: Record<string, string> = {
  '@': pathResolve("src"),
  'vue-i18n': 'vue-i18n/dist/vue-i18n.cjs.js',
}

/** 
 * @description-en vite document address
 * @description-cn vite官网
 * https://vitejs.cn/config/ */
export default ({ command }: ConfigEnv): UserConfigExport => {
  const prodMock = true;
  return {
    base: './',
    resolve: {
      alias
    },
    server: {
      open: true,
    },
    build: {
      rollupOptions: {
        output: {
          manualChunks: {
            'echarts': ['echarts']
          }
        }
      }
    },
    plugins: [
      vue(),
      viteMockServe({
        mockPath: 'mock',
        localEnabled: command === 'serve',
        prodEnabled: command !== 'serve' && prodMock,
        watchFiles: true,
        injectCode: `
          import { setupProdMockServer } from '../mockProdServer';
          setupProdMockServer();
        `,
        logger: true,
      }),
      vitePluginSvg({
        // 必要的。必须是绝对路径组成的数组。
        iconDirs: [
            resolve(__dirname, 'src/assets/svg'),
        ],
        // 必要的。入口script
        main: resolve(__dirname, 'src/main.js'),
        symbolIdFormat: 'icon-[name]'
      }),
    ],
    css: {
      postcss: {
        plugins: [
            {
              postcssPlugin: 'internal:charset-removal',
              AtRule: {
                charset: (atRule) => {
                  if (atRule.name === 'charset') {
                    atRule.remove();
                  }
                }
              }
            }
        ],
      },
    },
    define: {
      // enable hydration mismatch details in production build
      __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: 'true'
    },

  };
}
