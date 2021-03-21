package model

import (
	"ms_novel/utils"
)

type NovelThemeRelation struct {
	Id           int            `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	NovelThemeId int            `gorm:"column:novel_theme_id;int(10);not null;default:0;size:10;comment:'小说专题ID'"`
	NovelId      int            `gorm:"column:novel_id;int(10);not null;default:0;size:10;comment:'小说ID'"`
	CreatedAt    utils.Datetime `gorm:"column:created_at;"`
	UpdatedAt    utils.Datetime `gorm:"column:updated_at;"`
}

// 表名
func (NovelThemeRelation) TableName() string {
	return "novel_theme_relation"
}
