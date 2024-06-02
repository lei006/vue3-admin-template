import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/systemManage',
    component: Layout,
    redirect: '/systemManage/menu',
    meta: { title: 'message.menu.systemManage.name', icon: 'sfont system-xitongzhuangtai' },
    alwayShow: true,
    children: [
      {
        path: 'menu',
        component: createNameComponent(() => import('@/views/main/systemManage/menu/index.vue')),
        meta: { title: 'message.menu.systemManage.menu' }
      },
      {
        path: 'role',
        component: createNameComponent(() => import('@/views/main/systemManage/role/index.vue')),
        meta: { title: 'message.menu.systemManage.role' }
      }
    ]
  },
  {
    path: '/users',
    component: Layout,
    redirect: '/users/index',
    meta: { title: 'message.menu.print.name', icon: 'sfont system-24gl-printer' },
    children: [
      {
        path: 'index',
        component: createNameComponent(() => import('@/views/users/index.vue')),
        meta: { title: 'message.menu.systemManage.user' }
      }
    ]
  }

]

export default route