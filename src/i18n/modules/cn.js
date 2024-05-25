import locale from 'element-plus/dist/locale/zh-cn.mjs'

const lang = {
  el: locale.el, // element-plus i18 setting
  language: '中文',
  main: {
    title: '加密磁盘工具',
  },
  state:{
    NoDesk : 'Desk设备未插入',
    ToManyDesk : '暂不支持多设备',
    RemoteDeskExist : '等待远端信息',
    HitConnectRemote : '点击查看远端',
    ConnecttingRemote : '正在连接中..',
    ConnectReq : '连接请求中..',
    ConnectNoMaster : '被连接,连接远端',
    ConnectFailed: '连接失败,点请重试',
    CancelDownload: '取消文件传输',
    FirewallRuleNotExist:'防火墙规则不存在',
    TaskDownloadFinsh:'文件传输结束',
  },
  prompt:{
    fullScreen:"退出全屏:F11"
  },
  
  //login:{password:'1111'}
}

export default lang
