import type { Route } from '../index.type'
import Layout from '@/layout/index.vue'
import { createNameComponent } from '../createNode'
const route: Route[] = [
  {
    path: '/report',
    component: Layout,
    redirect: '/report/edit',
    meta: { title: 'message.menu.dashboard.name', icon: 'sfont system-home' },
    children: [
        {
            path: 'edit',
            component: createNameComponent(() => import('@/views/report_edit/index.vue')),
            meta: { title: 'message.menu.report_edit', icon: 'sfont system-home', hideClose: true }
        },
        {
            path: 'manage',
            component: createNameComponent(() => import('@/views/report_manage/index.vue')),
            meta: { title: 'message.menu.report_manage', icon: 'sfont system-home', hideClose: true }
        },
        {
            path: 'prefield',
            component: createNameComponent(() => import('@/views/report_prefield/index.vue')),
            meta: { title: 'message.menu.report_prefield', icon: 'sfont system-home', hideClose: true }
        },
        {
            path: 'template',
            component: createNameComponent(() => import('@/views/report_template/index.vue')),
            meta: { title: 'message.menu.report_template', icon: 'sfont system-home', hideClose: true }
        },
        {
            path: 'print',
            component: createNameComponent(() => import('@/views/report_print/index.vue')),
            meta: { title: 'message.menu.report_print', icon: 'sfont system-home', hideClose: true }
        },


    ]
  }
  
]

export default route
