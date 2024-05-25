import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import Index from './pages/index.vue'
import "preline/preline"
import { createMemoryHistory, createRouter } from 'vue-router'
import { type IStaticMethods } from "preline/preline";

declare global {
  interface Window {
    HSStaticMethods: IStaticMethods;
  }
}


const routes = [
    {
        path: '/', component: Index,
    }
]

const router = createRouter({
    history: createMemoryHistory(),
    routes,
})

router.afterEach((_, _1, failure) => {
    if (!failure) {
        setTimeout(() => {
            window.HSStaticMethods.autoInit();
        }, 100)
    }
})

createApp(App).use(router).mount('#app')
