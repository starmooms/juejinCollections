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
          msg: `Other Err: ${err.message}`,
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

const getArticle = async (params: any) => {
  const { response, data } = await request.fetch("/api/getArticle", {
    params: params,
    method: "GET"
  })
  console.log(response, data)
  return response
}

getArticle({
  articleId: ""
})

document.addEventListener("click", () => {
  getArticle({
    articleId: "684490397437866803923"
  })

})