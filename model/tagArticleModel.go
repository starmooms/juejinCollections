package model

import "time"

type TagArticleModel struct {
	TagId      string    `json:"tag_id" xorm:"pk notnull"`
	ArticleId  string    `json:"article_id" xorm:"pk notnull"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (m *TagArticleModel) TableName() string {
	return "tag_article"
}
