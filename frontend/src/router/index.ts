import { createWebHistory, createRouter } from 'vue-router'

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
]

const router = createRouter({
    scrollBehavior: () => ({ left: 0, top: 0 }),
    history: createWebHistory(),
    routes,
})

export default router