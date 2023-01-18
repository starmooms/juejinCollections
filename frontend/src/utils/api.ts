import { ApiStatus, Article } from "../type"
import FetchRequest from "./FetchRequest"

const request = new FetchRequest()

request.use(async (ctx, next) => {
  try {
    await next()
    const response = ctx.response
    if (response) {
      if (response.data) {
        const data = response.data
        if (data.status !== true) {
          console.error(data.msg)
        }
      } else if (response.response.status >= 400) {
        console.error(`Request Err: ${response.response.status} ${response.response.statusText}`)
      }
    }
  } catch (err) {
    console.error("Other Err:", err)
    if (!ctx.response) {
      ctx.response = {
        data: {
          status: false,
          msg: `Other Err: ${(err as any)?.message}`,
          err: err
        }
      } as never
    }
  }
})

request.use((ctx, next) => {
  ctx.config.headers = {
    'X-Requested-With': 'XMLHttpRequest',
    ...ctx.config.headers
  }
  return next()
})

export const getArticle = async (params: any) => {
  return await request.fetch<Article>("/api/getArticle", {
    params: params,
    method: "GET"
  })
}

export const postSyncCollection = async () => {
  return await request.fetch<ApiStatus>("/api/syncCollection", {
    method: "POST"
  })
}

export const searchArticleList = async (params: any) => {
  return await request.fetch<ApiStatus>("/api/searchArticle", {
    method: "GET",
    params,
  })
}
