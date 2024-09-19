import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/about',
    meta: { title: 'message.menu.about', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'about',
        component: createNameComponent(() => import('@/views/sys_about/index.vue')),
        meta: { title: 'message.menu.about', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route