/**
 * 前端路由管理
 **/

/** 路由类型 */
import type { Route } from '../index.type'

/** 引入需要权限的Modules */
import Dashboard from '../modules/dashboard'
import Document from '../modules/document'
import Pages from '../modules/pages'
import Menu from '../modules/menu'
import Component from '../modules/component'
import Directive from '../modules/directive'
import SystemManage from '../modules/systemManage'
import SysUser from '../modules/sys_user'
import SysLicense from '../modules/sys_license'
import SysAbout from '../modules/sys_about'
import Chart from '../modules/chart'
import Print from '../modules/print'
import Community from '../modules/community'
import Tab from '../modules/tab'

/** 登录后需要动态加入的本地路由 */
const FrontRoutes: Route[] = [
  ...Dashboard,
  //...Document,
  //...Component,
  //...Pages,
  //...Menu,
  //...Directive,
  //...Chart,
  ...SysLicense,
  ...SysUser,
  ...SystemManage,
  ...SysAbout,
  //...Print,
  //...Community,
  //...Tab,
]

export default FrontRoutes