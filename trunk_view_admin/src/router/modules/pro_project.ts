import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/project',
    meta: { title: 'message.menu.project', icon: 'sfont system-xitongzhuangtai' },
    children: [
      {
        path: 'project',
        component: createNameComponent(() => import('@/views/pro_project/index.vue')),
        meta: { title: 'message.menu.project', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]




export default route