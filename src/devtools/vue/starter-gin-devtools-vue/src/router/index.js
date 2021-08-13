import Vue from 'vue'
import VueRouter from 'vue-router'

import Home from '../views/Home.vue'
import DebugContext from '../views/DebugContext.vue'
import DebugRequests from '../views/DebugRequests.vue'
import DebugModules from '../views/DebugModules.vue'
import DebugComponents from '../views/DebugComponents.vue'

Vue.use(VueRouter)

const routes = [{
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/context',
        name: 'debug-context',
        component: DebugContext
    },
    {
        path: '/requests',
        name: 'debug-requests',
        component: DebugRequests
    },
    {
        path: '/modules',
        name: 'debug-modules',
        component: DebugModules
    },
    {
        path: '/components',
        name: 'debug-components',
        component: DebugComponents
    },
    {
        path: '/about',
        name: 'About',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: function() {
            return import ( /* webpackChunkName: "about" */ '../views/About.vue')
        }
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router