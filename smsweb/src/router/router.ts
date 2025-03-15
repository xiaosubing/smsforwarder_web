import {  createRouter, createWebHashHistory } from 'vue-router'

import Login from '../views/Login.vue'
import Index from '../views/Index.vue'
const routes = [
    {
        path: '/login',
        name: 'Login',
        meta: {title: "登台管理系统登录"},
        component:  Login
    },
    {
        path: '/',
        name: 'Home',
        component:  Login
    },
    {
        path: '/index',
        name: 'Index',
        meta: { requiresAuth: true },
        component:  Index
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})


router.beforeEach((to, from, next) => {
    const isAuthenticated = localStorage.getItem('token') || false
    if (to.meta.requiresAuth && !isAuthenticated) {
        next({ name: 'Login' })
    } else {
        next()
    }
})

export default router



















