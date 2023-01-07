import Message from "../components/Message";

interface fetchOpts extends RequestInit {
  params?: Record<any, any>
  parseJson?: boolean
  baseUrl?: string
}

// // ReturnType
// interface Callable<R> {
//   (...args: any[]): R;
// }
// type GenericReturnType<R, X> = X extends Callable<R> ? R : never;
type PromiseResolved<T> = T extends Promise<infer R> ? R : never;
type ReturnedPromiseResolvedType<T extends (...args: any) => any> = PromiseResolved<ReturnType<T>>


type fetchRetureP = ReturnedPromiseResolvedType<FetchRequest['request']>
interface fetchReture<T> extends fetchRetureP {
  data: ResponseData<T>
};
type fetchProperty<T> = {
  config: fetchOpts
  response?: fetchReture<T>
  reject?: Error
};

type middleFun<T = any> = (property: fetchProperty<T>, next: () => Promise<void>) => Promise<any>

type ResponseData<T = any> = {
  status: boolean
  data: T
  msg: string;
}

class ApiError extends Error {
  response: Response
  data: ResponseData

  constructor(msg: string, data: ResponseData, response: Response) {
    super(msg)
    this.response = response
    this.data = data
  }
}



export default class FetchRequest {

  middleware: middleFun[] = []
  defaultOpts: fetchOpts = {}

  constructor(opts: fetchOpts = {}) {
    this.defaultOpts = { ...this.defaultOpts, ...opts }
  }

  use(fun: middleFun) {
    this.middleware.push(fun)
  }

  async fetch<T = any>(url: string, userOpts?: fetchOpts) {
    let opts: fetchOpts = { ...this.defaultOpts, ...userOpts }
    let property: fetchProperty<T> = {
      config: opts
    }

    let isEmit = false
    let next = async () => {
      property.response = await this.request<T>(url, opts)
      isEmit = true
    }
    let middleList = [...this.middleware]
    for (let i = middleList.length - 1; i >= 0; i--) {
      const fn = middleList[i]
      const lastNext = next
      next = () => fn(property, lastNext)
    }

    return new Promise<fetchReture<T>>(async (resolve, reject) => {
      await next()
      if (!isEmit && !property.response) {
        throw new Error("no Emit fetch, check middle has use and await next?")
      }
      if (property.response) {
        resolve(property.response)
      } else {
        throw new Error("no response")
      }
    })
  }

  private async request<T = any>(url: string, opts: fetchOpts) {
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

    let baseUrl = opts.baseUrl
    if (baseUrl) {
      url = `${baseUrl}${url}`
    }

    const response = await fetch(url, opts)
    let data!: ResponseData<T>
    let type = response.headers.get('content-type')?.toLocaleLowerCase() || ""
    if (opts.parseJson !== false && type.indexOf("application/json") >= 0) {
      data = await response.clone().json()
    }

    if (!data.status) {
      Message.msg(JSON.stringify(data))
      throw new ApiError(`Api Error ${data?.msg || ''}`, data, response)
    }

    return {
      response,
      data
    }
  }
}