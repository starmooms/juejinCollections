import { createApp } from 'vue'
import App from './App.vue'
import './index.less'
import Message from './components/Message'
import { router } from './routes'

const app = createApp(App)
app.use(router)
app.use(Message)
app.mount('#app')