import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/admin',
    meta: { title: 'message.menu.systemManage.name', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'admin',
        component: createNameComponent(() => import('@/views/sys_admin/index.vue')),
        meta: { title: 'message.menu.admin', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route