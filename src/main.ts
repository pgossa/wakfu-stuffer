import './assets/main.css'

import { createApp, provide } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.provide('WURL', 'https://wakfu.cdn.ankama.com/gamedata/')

app.use(createPinia())
app.use(router)

app.mount('#app')
