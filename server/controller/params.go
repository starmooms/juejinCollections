package controller

type CursorBase struct {
	Limt   int `form:"limt" json:"limt"`
	Cursor int `form:"cursor" json:"cursor"`
}

type GetArticleParams struct {
	ArticleId string `form:"articleId" json:"articleId" binding:"required"`
}

type SearchArticleParams struct {
	Keyword string `form:"keyword" json:"keyword" binding:"required"`
	CursorBase
}
