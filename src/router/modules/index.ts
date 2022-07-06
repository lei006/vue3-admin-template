import { RouteRecordRaw } from "vue-router";

import Layout from '../../layout/index.vue'


const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    component: Layout,
    redirect: '/dashboard',
    children: [{
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('../../views/dashboard/index.vue'),
        meta: { title: 'dashboard', icon: 'dashboard' }
    }]
  },
  {
    path: "/report",
    component: Layout,
    redirect: '/report/index',
    children: [{
        path: 'index',
        name: 'Report',
        component: () => import('../../views/report/index.vue'),
        meta: { title: '报告', icon: 'dashboard' }
    }]
  },
  
  {
    path: "/about",
    name: "about",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ "../../views/system/about/index.vue"),
  },
];

export default routes;
