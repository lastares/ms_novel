package model

import (
	"ms_novel/utils"
)

type NovelArticle struct {
	Id             int            `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	ArticleTitle   string         `gorm:"column:article_title;not null;default:'';size:512"`
	ArticleAuthor  string         `gorm:"column:article_author;not null;default:'';size:128"`
	ArticleContent string         `gorm:"column:article_content;"`
	NovelId        int            `gorm:"column:novel_id;not null;size:10;index"`
	SortNum        int            `gorm:"column:sort_num;not null;size:10"`
	CreatedAt      utils.Datetime `gorm:"column:created_at;"`
	UpdatedAt      utils.Datetime `gorm:"column:updated_at;"`

	Novel *Novel `gorm:"association_autoupdate:false"`
}

// 表名
func (NovelArticle) TableName() string {
	return "novel_article"
}
