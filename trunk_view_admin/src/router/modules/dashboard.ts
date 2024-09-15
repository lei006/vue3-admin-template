import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    meta: { title: 'message.menu.dashboard.name', icon: 'sfont system-home' },
    children: [
      {
        path: 'dashboard',
        component: createNameComponent(() => import('@/views/main/dashboard/index.vue')),
        meta: { title: 'message.menu.dashboard.index', icon: 'sfont system-home', hideClose: true }
      },
    ]
  },
  {
    path: '/',
    component: Layout,
    redirect: '/personal',
    meta: { title: 'message.menu.personal', icon: 'sfont system-home' },
    hideMenu: true,
    children: [
      {
        path: 'personal',
        component: createNameComponent(() => import('@/views/sys_personal/index.vue')),
        meta: { title: 'message.menu.personal', icon: 'sfont system-home', hideTabs: true, hideClose: true }
      },
    ]
  }



]

export default route