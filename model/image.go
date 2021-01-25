package model

import "time"

type Image struct {
	Id         uint      `json:"id" xorm:"'id' pk autoincr notnull unique"`
	ArticleId  string    `json:"article_id" xorm:"index notnull"`
	Url        string    `json:"url" xorm:"index notnull"`
	Ctype      string    `json:"ctype"`
	Code       int       `json:"code"`
	Data       []byte    `json:"data"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (m *Image) TableName() string {
	return "images_db"
}
