
interface fetchOpts extends RequestInit {
  params: Record<any, any>
}

class FetchRequest {
  constructor() { }

  async fetch(url: string, userOpts?: fetchOpts) {
    let opts = { ...userOpts }
    let method = opts.method ? opts.method.toLocaleUpperCase() : 'GET'
    opts.method = method

    let params = opts.params
    if (params) {
      switch (method) {
        case 'GET':
          {
            const paramsStr = Object.keys(params).map(key => `${key}=${encodeURIComponent(params[key])}`).join("&")
            if (paramsStr) {
              url += (url.indexOf('?') >= 0 ? '&' : '?') + paramsStr
            }
          }
        case 'POST':
          opts = {
            ...opts,
            body: JSON.stringify(params),
            headers: {
              'Content-Type': 'application/json;charset=utf-8'
            },
          }
      }
    }

    const response = await fetch(url, opts)
  }
}