import vite from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from "@vitejs/plugin-vue-jsx"
import legacy from '@vitejs/plugin-legacy'
import UnoCSS from '@unocss/vite'
import { presetUno } from 'unocss'


const config: vite.UserConfig = {
  base: "/",
  server: {
    proxy: {
      '/api': "http://localhost:8012",
      "/echo": {
        target: "ws://localhost:8012",
        ws: true,
        secure: false,
      }
    },
  },
  // optimizeDeps: {
  //   include: ['prismjs/plugins/line-numbers/prism-line-numbers.min']
  // },

  build: {
    minify: false,
    // target: "es2015"
  },

  plugins: [
    vueJsx(),
    vue(),
    legacy({
      targets: ['Chrome >= 49', 'not IE 11'],
      polyfills: ['es/promise'],
      additionalLegacyPolyfills: ['regenerator-runtime/runtime']
    }),
    UnoCSS({
      presets: [
        presetUno(),
      ],
      variants: [
        (matcher) => {
          const result = matcher.match(/^(h-)(.*?):(.*)/)
          if (result && result.length === 4) {
            console.log(result[3], result[2])
            return {
              matcher: result[3],
              selector: s => `${s}.${result[2]}`,
            }
          }
          return matcher
        }
      ],
      theme: {
        colors: {
          mbg: '#ededed',
          mtext: '#222',
          primary: {
            DEFAULT: "#34d399",
            50: "#ecfdf5",
            100: "#d1fae5",
            200: "#a7f3d0",
            300: "#6ee7b7",
            400: "#34d399",
            500: "#10b981",
            600: "#059669",
            700: "#047857",
            800: "#065f46",
            900: "#064e3b"
          }
        },
      },
      preflights: [
        {
          getCSS: ({ theme }) => {
            const { mbg, mtext } = theme.colors
            return `
              body {
                background: ${mbg ?? '#333'};
                padding: 0;
                margin: 0;
                color: ${mtext ?? '#333'};
              }
            `
          }
        }
      ]
    }),
  ]
}

export default config