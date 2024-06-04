import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [

  {
    path: '/about',
    component: Layout,
    redirect: '/about/index',
    hideMenu: true,
    meta: { title: 'message.menu.print.name', icon: 'sfont system-24gl-printer' },
    children: [
      {
        path: 'index',
        component: createNameComponent(() => import('@/views/sys_about/index.vue')),
        meta: { title: 'message.menu.systemManage.about' }
      }
    ]
  },
  {
    path: '/setup',
    component: Layout,
    redirect: '/setup/index',
    hideMenu: true,
    meta: { title: 'message.menu.print.name', icon: 'sfont system-24gl-printer' },
    children: [
      {
        path: 'index',
        component: createNameComponent(() => import('@/views/sys_setup/index.vue')),
        meta: { title: 'message.menu.systemManage.setup' }
      }
    ]
  },

]

export default route