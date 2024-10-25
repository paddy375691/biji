import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import * as ELIcons from '@element-plus/icons-vue'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from './router'
import jwt from 'jsonwebtoken'

//使用钩子函数对路由进行权限跳转
router.beforeEach((to, from, next) => {
    document.title = to.meta.title;

    jwt.verify(localStorage.getItem('token'), 'adoodevops', function (err) {
        if (to.path === '/login') {
            next()
        } else if (err) {
            next('/login');
        } else {
            next();
        }
    });

});

const app = createApp(App)
for (let iconName in ELIcons) {
	app.component(iconName, ELIcons[iconName])
}
app.use(ElementPlus)
app.use(router)
app.mount('#app')