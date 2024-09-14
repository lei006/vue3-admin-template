import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/limit_ip',
    meta: { title: 'message.menu.limit_ip', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'limit_ip',
        component: createNameComponent(() => import('@/views/sys_limit_ip/index.vue')),
        meta: { title: 'message.menu.limit_ip', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route