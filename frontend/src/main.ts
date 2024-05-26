import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import Index from './pages/index.vue'
import { createMemoryHistory, createRouter } from 'vue-router'

const routes = [
    {
        path: '/', component: Index,
    }
]

const router = createRouter({
    history: createMemoryHistory(),
    routes,
})

createApp(App).use(router).mount('#app')
