package model

import "time"

// v2 改为了 TagId/TagName 先兼容 collection_id/collection_name

type Tag struct {
	CollectionId     string    `json:"collection_id" xorm:"index notnull unique"`
	CollectionName   string    `json:"collection_name" xorm:"index notnull"`
	PostArticleCount int       `json:"post_article_count"`
	ConcernUserCount int       `json:"concern_user_count"`
	CreatorId        string    `json:"creator_id"`
	CreateDate       time.Time `json:"create_date" xorm:"created"`
	UpdateDate       time.Time `json:"update_date" xorm:"updated"`
}

func (m *Tag) TableName() string {
	return "tags"
}
