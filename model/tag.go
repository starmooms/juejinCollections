package model

import "time"

// v2 改为了 TagId/TagName 先兼容 collection_id/collection_name
// UpdateTime 存在冲突

type Tag struct {
	Id               uint      `json:"id" xorm:"'id' pk notnull unique"`
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
