import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import Index from './pages/index.vue'
import Login from './pages/login.vue'
import { createWebHistory, createRouter } from 'vue-router'

const routes = [
    {
        path: '/', component: Index,
    },
    {
        
        path: '/login', component: Login
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

const app = createApp(App)
app.use(router)
app.mount('#app')
