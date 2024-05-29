// 提示信息仅在开发环境生效
import { createI18n } from 'vue-i18n'

const messages = {}

import LangEn from './modules/en'
import LangZhCH from './modules/cn'



messages["en"] = LangEn;
messages["cn"] = LangZhCH;


//const lang = store.state.lang || navigator.language // 初次进入，采用浏览器当前设置的语言，默认采用中文
//const locale = lang.indexOf('en') !== -1 ? 'en' : 'zh-cn'
const locale = 'cn'

const i18n = createI18n({
  legacy: false, // VUE3 组合式API
  locale, // 默认cn语言环境
  fallbackLocale: 'cn', //备用语言环境
  messages
})

document.getElementsByTagName('html')[0].setAttribute('lang', locale)

export default i18n
