package model

import (
	"ms_novel/utils"
)

type NovelTheme struct {
	Id               int            `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	ThemeName        string         `gorm:"column:theme_name;not null;default:'';size:64"`
	ThemeSubtitle    string         `gorm:"column:theme_subtitle;not null;default:'';size:512"`
	ThemeDescription string         `gorm:"column:theme_description;not null;default:'';size:512"`
	ThemeCover       string         `gorm:"column:theme_cover;not null;default:'';size:256"`
	SortNum          int            `gorm:"column:sort_num;not null;size:10"`
	CreatedAt        utils.Datetime `gorm:"column:created_at;"`
	UpdatedAt        utils.Datetime `gorm:"column:updated_at;"`

	Novels []*Novel `gorm:"many2many:novel_theme_relation;"`
}

// 表名
func (NovelTheme) TableName() string {
	return "novel_theme"
}
