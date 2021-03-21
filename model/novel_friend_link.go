package model

import (
	"ms_novel/utils"
)

type FriendLink struct {
	Id        int            `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	LinkTitle string         `gorm:"column:link_title;not null;default:'';size:128"`
	Link      string         `gorm:"column:link;not null;default:'';size:256"`
	SortNum   int            `gorm:"column:sort_num;not null;size:10"`
	CreatedAt utils.Datetime `gorm:"column:created_at;"`
	UpdatedAt utils.Datetime `gorm:"column:updated_at;"`
}

// 表名
func (FriendLink) TableName() string {
	return "novel_friend_link"
}
