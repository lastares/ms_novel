package model

import (
	"ms_novel/utils"
)

var (
	NOVEL_STATUS_SERIALING = 1 // 连载中
	NOVEL_STATUS_FINISH    = 2 // 已完结
)

type Novel struct {
	Id                int            `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	NovelTitle        string         `gorm:"column:novel_title;not null;default:'';size:256;index"`
	NovelAuthor       string         `gorm:"column:novel_author;not null;default:'';size:128"`
	NovelDescription  string         `gorm:"column:novel_description;not null;default:'';size:512"`
	NovelCover        string         `gorm:"column:novel_cover;not null;default:'';size:512;comment:'小说封面'"`
	NovelBannerCover  string         `gorm:"column:novel_banner_cover;not null;default:'';size:256;comment:'小说banner封面'"`
	MainRole          string         `gorm:"column:main_role;not null;default:'';size:256;comment:'小说主角'"`
	ChapterPreview    string         `gorm:"column:chapter_preview;comment:'章节预览';"`
	NovelStatus       int            `gorm:"column:novel_status;not null;size:4;default:1;comment:'小说状态:1-连载中 2-已完结';index;"`
	NovelNavigationId int            `gorm:"column:novel_navigation_id;not null;default:0;size:10;comment:'导航栏目ID';index"`
	NovelCategoryId   int            `gorm:"column:novel_category_id;not null;default:0;size:10;comment:'小说分类ID';index"`
	ViewNum           int            `gorm:"column:view_num;not null;default:0;size:10;comment:小说阅读次数"`
	IsNew             int            `gorm:"column:is_new;not null;default:0;size:4;comment:'是否是最新：0-否 1-是'"`
	IsHot             int            `gorm:"column:is_hot;not null;default:0;size:4;comment:'是否是最热：0-否 1-是'"`
	IsRecommend       int            `gorm:"column:is_recommend;not null;default:0;size:4;comment:'是否推荐：0-否 1-是'"`
	IsDisplayIndex    int            `gorm:"column:is_display_index;not null;default:0;size:4;comment:'是否首页展示：0-否 1-是'"`
	SortNum           int            `gorm:"column:sort_num;not null;default:0;size:10"`
	CreatedAt         utils.Datetime `gorm:"column:created_at;"`
	UpdatedAt         utils.Datetime `gorm:"column:updated_at;"`

	NovelNavigation *NovelNavigation `gorm:"association_autoupdate:false"`
	NovelCategory   *NovelCategory   `gorm:"association_autoupdate:false"`
	NovelThemes     []*NovelTheme    `gorm:"many2many:novel_theme_relation;association_autoupdate:false"`
}

// 表名
func (Novel) TableName() string {
	return "novel"
}

func (p *Novel) GetNovelStatusTypes() map[int]string {
	return map[int]string{
		NOVEL_STATUS_SERIALING: "连载中",
		NOVEL_STATUS_FINISH:    "已完结",
	}
}

// 获取小说状态
func (p *Novel) GetNovelStatusTitle(novelStatus int) string {
	novelStatusTypes := p.GetNovelStatusTypes()
	if _, ok := novelStatusTypes[novelStatus]; ok {
		return novelStatusTypes[novelStatus]
	}
	return ""
}

func (p *Novel) GetIsValueTitle(isNew int) (valueTitle string) {
	valueTitle = "否"
	if isNew == COMMON_ONE {
		valueTitle = "是"
	}
	return valueTitle
}
