package model

import "time"

type TagModel struct {
	Id               uint      `json:"id" xorm:"'id' pk notnull unique"`
	TagId            string    `json:"tag_id" xorm:"index notnull unique"`
	TagName          string    `json:"tag_name" xorm:"index notnull"`
	Color            string    `json:"color"`
	Icon             string    `json:"icon"`
	BackGround       string    `json:"back_ground"`
	Ctime            uint      `json:"ctime"`
	Mtime            uint      `json:"mtime"`
	Status           int       `json:"status"`
	CreatorId        uint      `json:"creator_id"`
	UserName         string    `json:"user_name"`
	PostArticleCount uint      `json:"post_article_count"`
	ConcernUserCount uint      `json:"concern_user_count"`
	Isfollowed       bool      `json:"isfollowed"`
	IsHasIn          bool      `json:"is_has_in"`
	CreateTime       time.Time `json:"create_time" xorm:"created"`
	UpdateTime       time.Time `json:"update_time" xorm:"updated"`
}

func (m *TagModel) TableName() string {
	return "tags"
}
