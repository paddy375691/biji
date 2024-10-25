import {createRouter, createWebHashHistory} from 'vue-router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import Layout from "@/layout/Layout"

const routes = [
    {
        path: '/login',
        component: () => import('@/views/login/Login.vue'),
        icon: "odometer",
        meta: {title: "登录", requireAuth: false},
    },
    {
        path: '/',
        redirect: '/main'
    },
    {
        path: '/main',
        component: Layout,
        icon: "odometer",
        meta: {title: "概要", requireAuth: true},
        redirect: '/home',
        children: [
            {
                path: "/home",
                name: "概要",
                icon: "odometer",
                meta: {title: "概要", requireAuth: true},
                component: () => import('@/views/home/Home.vue'),
            }
        ]
    },
    {
        path: '/workflow',
        component: Layout,
        icon: "VideoPlay",
        meta: {title: "工作流", requireAuth: true},
        children: [
            {
                path: "/workflow",
                name: "工作流",
                icon: "VideoPlay",
                meta: {title: "工作流", requireAuth: true},
                component: () => import('@/views/workflow/Workflow.vue')
            }
        ]
    },
    {
        path: "/cluster",
        name: "集群",
        component: Layout,
        icon: "home-filled",
        meta: {title: "集群", requireAuth: true},
        children: [
            {
                path: "/node",
                name: "Node",
                icon: "el-icon-s-data",
                meta: {title: "Node", requireAuth: true},
                component: () => import("@/views/node/Node.vue")
            },
            {
                path: "/namespace",
                name: "Namespace",
                icon: "el-icon-document-add",
                meta: {title: "Namespace", requireAuth: true},
                component: () => import("@/views/namespace/Namespace.vue")
            },
            {
                path: "/persistentvolume",
                name: "PersistentVolume",
                icon: "el-icon-document-add",
                meta: {title: "PersistemtVolume", requireAuth: true},
                component: () => import("@/views/persistentvolume/PersistentVolume.vue")
            }
        ]
    },
    {
        path: "/workload",
        name: "工作负载",
        component: Layout,
        icon: "menu",
        meta: {title: "工作负载", requireAuth: true},
        children: [
            {
                path: "/deployment",
                name: "Deployment",
                icon: "el-icon-s-data",
                meta: {title: "Deployment", requireAuth: true},
                component: () => import("@/views/deployment/Deployment.vue")
            },
            {
                path: "/pod",
                name: "Pod",
                icon: "el-icon-document-add",
                meta: {title: "Pod", requireAuth: true},
                component: () => import("@/views/pod/Pod.vue")
            },
            {
                path: "/deamonset",
                name: "DaemonSet",
                icon: "el-icon-document-add",
                meta: {title: "DaemonSet", requireAuth: true},
                component: () => import("@/views/daemonset/DaemonSet.vue")
            },
            {
                path: "/statefulset",
                name: "StatefulSet",
                icon: "el-icon-document-add",
                meta: {title: "DaemonSets", requireAuth: true},
                component: () => import("@/views/statefulset/StatefulSet.vue")
            }
        ]
    },
    {
        path: "/loadbalance",
        name: "负载均衡",
        component: Layout,
        icon: "files",
        meta: {title: "负载均衡", requireAuth: true},
        children: [
            {
                path: "/service",
                name: "Service",
                icon: "el-icon-s-data",
                meta: {title: "Service", requireAuth: true},
                component: () => import("@/views/service/Service.vue")
            },
            {
                path: "/ingress",
                name: "Ingress",
                icon: "el-icon-document-add",
                meta: {title: "Ingress", requireAuth: true},
                component: () => import("@/views/ingress/Ingress.vue")
            }
        ]
    },
    {
        path: "/storage",
        name: "存储与配置",
        component: Layout,
        icon: "tickets",
        meta: {title: "存储与配置", requireAuth: true},
        children: [
            {
                path: "/configmap",
                name: "Configmap",
                icon: "el-icon-document-add",
                meta: {title: "Configmap", requireAuth: true},
                component: () => import("@/views/configmap/ConfigMap.vue")
            },
            {
                path: "/secret",
                name: "Secret",
                icon: "el-icon-document-add",
                meta: {title: "Secret", requireAuth: true},
                component: () => import("@/views/secret/Secret.vue")
            },
            {
                path: "/persistentvolumeclaim",
                name: "PersistentVolumeClaim",
                icon: "el-icon-s-data",
                meta: {title: "PersistentVolumeClaim", requireAuth: true},
                component: () => import("@/views/persistentvolumeclaim/PersistentVolumeClaim.vue")
            },
        ]
    },
]

// createRouter 创建路由实例
const router = createRouter({
    /**
     * hash模式：createWebHashHistory，
     * history模式：createWebHistory
     */
    history: createWebHashHistory(),
    routes
})

NProgress.inc(0.2)
NProgress.configure({ easing: 'ease', speed: 600, showSpinner: false })

// 设置title
router.beforeEach((to, from, next) => {
    // 启动进度条
    NProgress.start()

    // 设置头部
    if (to.meta.title) {
        document.title = to.meta.title
    } else {
        document.title = "默存"
    }
    next()
})

router.afterEach(() => {
    // 关闭进度条
    NProgress.done()
})

// 抛出路由实例, 在 main.js 中引用
export default router