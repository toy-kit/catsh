import { createWebHashHistory, createRouter } from 'vue-router'

const routes = [
    {
        path: '/',
        name: "Index",
        component: () => import('../views/index.vue'),
    },
    {
        path: '/setting',
        name: "Setting",
        component: () => import('../views/setting.vue'),
    },
    {
        path: '/about',
        name: "About",
        component: () => import('../views/about.vue'),
    },
]

const router = createRouter({
    scrollBehavior: () => ({ left: 0, top: 0 }),
    history: createWebHashHistory(),
    routes,
})

export default router