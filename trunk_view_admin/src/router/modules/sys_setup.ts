import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/setup',
    meta: { title: 'message.menu.config', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'setup',
        component: createNameComponent(() => import('@/views/sys_setup/index.vue')),
        meta: { title: 'message.menu.config', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route