package model

import "ms_novel/utils"

type NovelNavigation struct {
	Id             int            `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	NavigationName string         `gorm:"column:navigation_name;not null;default:'';size:64"`
	SortNum        int            `gorm:"column:sort_num;not null;size:10"`
	IsDisplayIndex int            `gorm:"column:is_display_index;not null;default:0;size:4;comment:'是否首页展示：0-否 1-是'"`
	CreatedAt      utils.Datetime `gorm:"column:created_at"`
	UpdatedAt      utils.Datetime `gorm:"column:updated_at;"`

	NovelCategories []NovelCategory
	Novels          []Novel
}

// 表名
func (NovelNavigation) TableName() string {
	return "novel_navigation"
}
