import vite from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from "@vitejs/plugin-vue-jsx"
import legacy from '@vitejs/plugin-legacy'

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

  build: {
    minify: false,
    target: "es2015"
  },

  plugins: [
    vueJsx(),
    vue(),
    legacy({
      targets: ['Chrome >= 49', 'not IE 11'],
      polyfills: ['es/promise'],
      additionalLegacyPolyfills: ['regenerator-runtime/runtime']
    })

  ]
}

export default config