import vite from 'vite'
import vue from '@vitejs/plugin-vue'

const config: vite.UserConfig = {
  base: "/",
  server: {
    proxy: {
      '/api': "http://localhost:8012"
    },
  },
  plugins: [vue()]
}

export default config