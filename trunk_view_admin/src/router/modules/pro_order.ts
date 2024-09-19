import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/order',
    meta: { title: 'message.menu.order', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'order',
        component: createNameComponent(() => import('@/views/pro_order/index.vue')),
        meta: { title: 'message.menu.order', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route