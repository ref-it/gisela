import './assets/main.css'

import { createApp, watch } from 'vue'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n'

import App from './App.vue'
import router from './router'
import messages from './i18n'

const pinia = createPinia()

const i18n = createI18n({
    locale: navigator.language,
    fallbackLocale: 'de',
    legacy: false,
    messages
})

const app = createApp(App)

app.use(pinia)
app.use(i18n)
app.use(router)

app.mount('#app')
