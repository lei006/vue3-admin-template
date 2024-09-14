import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/user',
    meta: { title: 'message.menu.systemManage.name', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'user',
        component: createNameComponent(() => import('@/views/sys_user/index.vue')),
        meta: { title: 'message.menu.systemManage.user', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route