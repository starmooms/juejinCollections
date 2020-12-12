import { BuildConfig } from 'vite'
export default {
  base: "/",
  proxy: {
    '/api': "http://localhost:8012"
  }
} as Partial<BuildConfig>