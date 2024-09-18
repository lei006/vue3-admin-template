import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/license',
    meta: { title: 'message.menu.license', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'license',
        component: createNameComponent(() => import('@/views/sys_license/index.vue')),
        meta: { title: 'message.menu.license', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route