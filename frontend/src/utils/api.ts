import FetchRequest from "./FetchRequest"

const request = new FetchRequest()

const getArticle = async (params: any) => {
  const { response } = await request.fetch("/api/getArticle", {
    params: params,
    method: "GET"
  })
  console.log(response)
  console.log(response.headers.get("content-type"))
  return response
}

getArticle({
  articleId: "6844903974378668039"
})