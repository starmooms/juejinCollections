package model

import "time"

type Article struct {
	ArticleId      string    `json:"article_id" xorm:"pk notnull unique"`
	AuditStatus    int       `json:"audit_status"`
	BriefContent   string    `json:"brief_content"`
	CategoryId     string    `json:"category_id"`
	CollectCount   int       `json:"collect_count"`
	Content        string    `json:"content"`
	CoverImage     string    `json:"cover_image"`
	Ctime          string    `json:"ctime"`
	MarkContent    string    `json:"mark_content"`
	Mtime          string    `json:"mtime"`
	OriginalAuthor string    `json:"original_author"`
	OriginalType   int       `json:"original_type"`
	Rtime          string    `json:"rtime"`
	Status         int       `json:"status"`
	Title          string    `json:"title"`
	UserId         string    `json:"user_id"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
}

func (m *Article) TableName() string {
	return "articles"
}
