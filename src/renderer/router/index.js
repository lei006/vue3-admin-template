import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/*
//解决，重复制时的冗余错误
const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}
*/

/* Layout */
import Layout from '@/layout'


/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  /*
  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },
  */
  {
    path: '/404',
    component: Layout,
    redirect: '/404/index',
    children: [{
      path: 'index',
      name: '404',
      component: () => import('@/views/404'),
      meta: { title: '404', icon: 'dashboard' }
    }]
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: 'Dashboard', icon: 'dashboard' }
    }]
  },

  {
    path: '/extensions',
    component: Layout,
    redirect: '/extensions/index',
    children: [{
      path: 'index',
      name: 'Extensions',
      component: () => import('@/views/extensions/index'),
      meta: { title: '扩展', icon: 'dashboard' }
    }]
  },

  {
    path: '/page01',
    component: Layout,
    redirect: '/page01/index',
    children: [{
      path: 'index',
      name: 'Page01',
      component: () => import('@/views/page01/index'),
      meta: { title: 'page01', icon: 'dashboard' }
    }]
  },
  {
    path: '/page02',
    component: Layout,
    redirect: '/page02/index',
    children: [{
      path: 'index',
      name: 'Page02',
      component: () => import('@/views/page02/index'),
      meta: { title: 'page02', icon: 'dashboard' }
    }]
  },
  {
    path: '/testApi',
    component: Layout,
    redirect: '/testApi/index',
    children: [{
      path: 'index',
      name: 'TestApi',
      component: () => import('@/views/testApi/index'),
      meta: { title: '测试API', icon: 'dashboard' }
    }]
  },
  
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
