import { createApp } from 'vue'
import App from './App.vue'
import './index.less'
import Message from './components/Message'

const app = createApp(App)
app.use(Message)
app.mount('#app')