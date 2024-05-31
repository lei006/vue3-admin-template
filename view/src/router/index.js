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
    meta: { title: 'message.menu.dashboard.name', icon: 'sfont system-home' },
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/HomeView.vue'),
        meta: { title: 'message.menu.dashboard.index', icon: 'sfont system-home', hideClose: true }
      }
    ]
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: function () {
      return import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
    }
  }
]



const router = createRouter({
  history: createWebHashHistory(),
  routes
})

// 检查权限
Permission(router)


export default router
