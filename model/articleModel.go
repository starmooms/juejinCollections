package model

import "time"

type ArticleModel struct {
	ArticleId    string `json:"article_id" xorm:"pk notnull unique"`
	AuditStatus  int    `json:"audit_status"`
	BriefContent string `json:"brief_content"`
	CategoryId   string `json:"category_id"`
	CollectCount int    `json:"collect_count"`
	// CommentCount int    `json:"comment_count"`
	Content    string `json:"content"`
	CoverImage string `json:"cover_image"`
	Ctime      string `json:"ctime"`
	// DiggCount    int    `json:"digg_count"`
	// DraftId      string `json:"draft_id"`
	// HotIndex     int    `json:"hot_index"`
	// IsEnglish    int    `json:"is_english"`
	// IsGfw        int    `json:"is_gfw"`
	// IsHot        int    `json:"is_hot"`
	// IsOriginal   int    `json:"is_original"`
	// LinkUrl        int       `json:"link_url"`
	MarkContent    string `json:"mark_content"`
	Mtime          string `json:"mtime"`
	OriginalAuthor string `json:"original_author"`
	OriginalType   int    `json:"original_type"`
	// RankIndex      float64 `json:"rank_index"`
	Rtime  string `json:"rtime"`
	Status int    `json:"status"`
	// TagIds         []int     `json:"tag_ids"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
	// UserIndex    float64   `json:"user_index"`
	// VerifyStatus int       `json:"verify_status"`
	// ViewCount    int       `json:"view_count"`
	// VisibleLevel int       `json:"visible_level"`
	CreateTime time.Time `json:"CreateTime"`
	UpdateTime time.Time `json:"UpdateTime"`
}

func (m *ArticleModel) TableName() string {
	return "articles"
}
