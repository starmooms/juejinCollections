import vite from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from "@vitejs/plugin-vue-jsx"

const config: vite.UserConfig = {
  base: "/",
  server: {
    proxy: {
      '/api': "http://localhost:8012"
    },
  },
  // optimizeDeps: {
  //   include: ['prismjs/plugins/line-numbers/prism-line-numbers.min']
  // },
  plugins: [vueJsx(), vue()]
}

export default config