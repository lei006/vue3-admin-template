import locale from 'element-plus/dist/locale/en.mjs'


const lang = {
  el: locale.el, // element-plus i18 setting
  language: 'English',
  main: {
    title: 'Encrypted Hard Disk Encloseure',
  },
  state:{
    NoDesk : 'Desk are not inserted',
    ToManyDesk : 'Too much Desk',
    RemoteDeskExist : 'Wait the remote',
    HitConnectRemote : 'connection',
    ConnecttingRemote : 'In the connection..',
    ConnectReq : 'Connect request..',
    ConnectNoMaster : 'Be connected',
    ConnectFailed: 'failed,try again',
    CancelDownload: 'Cancel Download',
    FirewallRuleNotExist:'Firewall rule does not exist',
    TaskDownloadFinsh:'File Transfer Ended',
  },
  prompt:{
    fullScreen:"Exit full screen:F11"
  },  
}


export default lang

