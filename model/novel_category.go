package model

import (
	"ms_novel/utils"
)

type NovelCategory struct {
	Id                int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CategoryName      string `gorm:"column:category_name;not null;default:'';size:256"`
	CategoryIcon      string `gorm:"column:category_icon;not null;default:'';size:256"`
	IsDisplayIndex    int    `gorm:"column:is_display_index;not null;size:4"`
	SortNum           int    `gorm:"column:sort_num;not null;size:10"`
	NovelNavigationId int    `gorm:"column:novel_navigation_id;not null;default:'0';size:10;index"`

	CreatedAt utils.Datetime `gorm:"column:created_at;"`
	UpdatedAt utils.Datetime `gorm:"column:updated_at;"`

	NovelNavigation *NovelNavigation `gorm:"association_autoupdate:false"`
	Novels          []Novel
}

// 表名
func (NovelCategory) TableName() string {
	return "novel_category"
}
