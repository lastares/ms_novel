package model

import (
	"ms_novel/utils"
)

type NovelChapter struct {
	Id           int            `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	ChapterNo    int            `gorm:"column:chapter_no;not null;default:0;size:10"`
	ChapterTitle string         `gorm:"column:chapter_title;comment:'章节标题';"`
	CreatedAt    utils.Datetime `gorm:"column:created_at;"`
	UpdatedAt    utils.Datetime `gorm:"column:updated_at;"`

	NovelId int    `gorm:"novel_id;index"`
	Novel   *Novel `gorm:"association_autoupdate:false"`
}

// 表名
func (NovelChapter) TableName() string {
	return "novel_chapter"
}
