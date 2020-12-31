import { UserConfig } from 'vite'

const config: UserConfig = {
  base: "/",
  proxy: {
    '/api': "http://localhost:8012"
  }
}

export default config