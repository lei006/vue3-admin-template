/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const systemRouter = {
  path: '/system',
  component: Layout,
  redirect: '/system/menu1/menu1-1',
  name: 'System',
  meta: {
    title: '系统信息',
    icon: 'nested'
  },
  children: [

    {
      path: 'user',
      name: 'Menu22',
      component: () => import('@/views/system/user/index'),
      meta: { title: '用户管理' , icon: 'el-icon-user'}
    },

    {
      path: 'setup',
      name: 'Menu21',
      component: () => import('@/views/system/setup/index'),
      meta: { title: '系统设置', icon: 'el-icon-setting' }
    },
    {
      path: 'about',
      name: 'Table4',
      component: () => import('@/views/system/about/index'),
      meta: { title: '关于软件', icon: 'el-icon-info'   }
    }
  ]
}

export default systemRouter
