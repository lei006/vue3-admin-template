import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import Permission from './permission'
import Layout from '@/layout/index.vue'

const routes = [
  /*
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  */
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    meta: { title: 'message.menu.dashboard.name', icon: 'House' },
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/dashboard/index.vue'),
        meta: { title: '仪表盘', icon: 'CircleCheckFilled', hideClose: true }
      },
    ]
  },
  {
    path: '/nested',
    component: Layout,
    redirect: '/nested/menu1',
    meta: { title: 'nested', icon: 'Search' },
    children: [
      {
        path: 'menu1',
        meta: { title: 'menu1', icon: 'Tickets', hideClose: true },
        children: [
          {
            path: 'menu1-1',
            component: () => import('@/views/nested/menu1/menu1-1/index.vue'),
            meta: { title: 'menu1-1', icon: 'Search', hideClose: true }
          },
          {
            path: 'menu1-2',
            meta: { title: 'menu1-2', icon: 'Edit', hideClose: true },
            children: [
              {
                path: 'menu1-2-1',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-1/index.vue'),
                meta: { title: 'menu1-2-1', icon: 'DishDot', hideClose: true }
              },
              {
                path: 'menu1-2-2',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-2/index.vue'),
                meta: { title: 'menu1-2-2', icon: 'Reading', hideClose: true }
              },
            ]            
          },
          {
            path: 'menu1-3',
            component: () => import('@/views/nested/menu1/menu1-3/index.vue'),
            meta: { title: 'menu1-3', icon: 'Mouse', hideClose: true }
          },
        ]
      },

      {
        path: 'menu2',
        component: () => import('@/views/nested/menu2/index.vue'),
        meta: { title: 'menu2', icon: 'Paperclip', hideClose: true }
      },

    ]
  },
  
  {
    path: '/about',
    redirect: '/dashboard',
    component: Layout,
    meta: { title: 'nested', icon: 'Coin' },
    children: [
      {
        path: 'About',
        component: () => import('@/views/AboutView.vue'),
        meta: { title: 'About', icon: 'sfont system-home', hideClose: true }
      }
    ]
  }
]



const router = createRouter({
  history: createWebHashHistory(),
  routes
})

// 检查权限
Permission(router)


export default router
