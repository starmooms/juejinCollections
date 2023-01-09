package controller

type GetArticleParams struct {
	ArticleId string `form:"articleId" json:"articleId" binding:"required"`
}
