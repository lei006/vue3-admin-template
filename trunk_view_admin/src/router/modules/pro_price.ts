import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/price',
    meta: { title: 'message.menu.price', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'price',
        component: createNameComponent(() => import('@/views/pro_price/index.vue')),
        meta: { title: 'message.menu.price', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route