
interface fetchOpts extends RequestInit {
  params?: Record<any, any>
  parseJson?: boolean
}
type middleFun = (config: fetchOpts, next: () => Promise<ReturnType<FetchRequest['request']>>) => unknown

interface Opts {
  baseUrl: string
}

export default class FetchRequest {
  opts: Opts = {
    baseUrl: ""
  }
  middleware: middleFun[] = []

  constructor(opts: Partial<Opts> = {}) {
    this.opts = { ...opts, ...this.opts }
  }

  async fetch<T = any>(url: string, userOpts?: fetchOpts) {
    let opts: fetchOpts = { ...userOpts }
    let method = opts.method ? opts.method.toLocaleUpperCase() : 'GET'
    opts.method = method

    let params = opts.params
    if (params) {
      delete opts.params
      switch (method) {
        case 'GET':
          {
            const paramsStr = Object.keys(params).map(key => `${key}=${encodeURIComponent(params![key])}`).join("&")
            if (paramsStr) {
              url += (url.indexOf('?') >= 0 ? '&' : '?') + paramsStr
            }
          }
          break
        case 'POST':
          opts = {
            ...opts,
            body: JSON.stringify(params),
            headers: {
              'Content-Type': 'application/json;charset=utf-8'
            },
          }
          break
      }
    }

    let baseUrl = this.opts.baseUrl
    if (baseUrl) {
      url = `${baseUrl}${url}`
    }

    let next: any = () => this.request(url, opts)
    let middleList = [...this.middleware]
    for (let i = middleList.length - 1; i >= 0; i--) {
      const fn = middleList[i]
      const lastNext = next
      next = () => Promise.resolve(fn(opts, lastNext))
    }

    return next()
  }

  private async request<T = any>(url: string, opts: fetchOpts) {
    const response = await fetch(url, opts)
    let data!: T
    if (opts.parseJson !== false) {
      data = await response.clone().json()
    }
    return {
      response,
      data
    }
  }
}