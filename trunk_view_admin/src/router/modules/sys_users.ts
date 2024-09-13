import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/users',
    meta: { title: 'message.menu.systemManage.name', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'users',
        component: createNameComponent(() => import('@/views/sys_users/index.vue')),
        meta: { title: 'message.menu.systemManage.user', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route