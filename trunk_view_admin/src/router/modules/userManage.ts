import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/userManage',
    component: Layout,
    redirect: '/userManage',
    meta: { title: 'message.menu.systemManage.name', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'userManage',
        component: createNameComponent(() => import('@/views/usermanage/index.vue')),
        meta: { title: 'message.menu.systemManage.user', icon: 'sfont user', hideClose: true  }
      }
    ]
  }
]




export default route