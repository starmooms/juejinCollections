import vite from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'



const config: vite.UserConfig = {
  base: "/",
  server: {
    proxy: {
      '/api': "http://localhost:8012"
    },
  },
  plugins: [vueJsx(), vue()],
  // optimizeDeps: {
  //   exclude: ["vue-markdown"]
  // },
}

export default config