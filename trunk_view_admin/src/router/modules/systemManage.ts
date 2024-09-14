import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/systemManage',
    component: Layout,
    redirect: '/systemManage/user',
    meta: { title: 'message.menu.systemManage.name', icon: 'sfont system-xitongzhuangtai' },
    alwayShow: true,
    children: [
      {
        path: 'option',
        component: createNameComponent(() => import('@/views/sys_option/index.vue')),
        meta: { title: 'message.menu.option', icon: 'sfont system-24gl-printer', hideClose: true  }
      },
      {
        path: 'admin',
        component: createNameComponent(() => import('@/views/sys_admin/index.vue')),
        meta: { title: 'message.menu.admin', icon: 'sfont system-24gl-printer', hideClose: true  }
      },            
      {
        path: 'limit_ip',
        component: createNameComponent(() => import('@/views/sys_limit_ip/index.vue')),
        meta: { title: 'message.menu.limit_ip', icon: 'sfont system-24gl-printer', hideClose: true  }
      },
      {
        path: 'setup',
        component: createNameComponent(() => import('@/views/sys_setup/index.vue')),
        meta: { title: 'message.menu.config', icon: 'sfont system-24gl-printer', hideClose: true  }
      }
    ]
  }
]

export default route