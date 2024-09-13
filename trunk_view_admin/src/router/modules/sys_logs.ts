import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/logs',
    meta: { title: 'message.menu.logs', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'logs',
        component: createNameComponent(() => import('@/views/sys_users/index.vue')),
        meta: { title: 'message.menu.logs', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route