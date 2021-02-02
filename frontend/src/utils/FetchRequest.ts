
interface fetchOpts extends RequestInit {
  params?: Record<any, any>
  parseJson?: boolean
}

interface Opts {
  baseUrl: string
}

export default class FetchRequest {
  opts: Opts = {
    baseUrl: ""
  }

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