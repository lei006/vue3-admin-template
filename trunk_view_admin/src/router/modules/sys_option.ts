import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/option',
    meta: { title: 'message.menu.option', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'option',
        component: createNameComponent(() => import('@/views/sys_option/index.vue')),
        meta: { title: 'message.menu.option', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route