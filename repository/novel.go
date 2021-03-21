package repository

import (
	"github.com/chenhg5/collection"
	"ms_novel/global"
	"ms_novel/model"
	"ms_novel/utils"
)

// 定义接口，便于查看所有方法
type NovelRepositoryInterface interface {
	// 创建
	Create(
		novelTitle, novelAuthor, novelDescription, novelCover, novelBannerCover, mainRole, chapterPreview string,
		novelStatus, sortNum int,
		novelNavigation *model.NovelNavigation,
		novelCategory *model.NovelCategory,
		novelThemes []*model.NovelTheme,
	) error
	// 修改
	Modify(
		novel *model.Novel,
		novelStatus, sortNum int,
		novelTitle, novelAuthor, novelDescription, novelCover, novelBannerCover, mainRole, chapterPreview string,
		novelNavigation *model.NovelNavigation,
		novelCategory *model.NovelCategory,
		novelThemes []*model.NovelTheme,
	) error
	// 删除
	Delete(novel *model.Novel) error

	// 获取单条分类
	Get(novelId int, novelTitle string, preLoads []string) (*model.Novel, error)
	// 列表
	GetList(novelTitle string, novelStatus int, page, pageSize int, preLoads []string) (list []*model.Novel, total int, err error)
	// 获取所有数据
	GetAll(
		selectFields []string,
		isDisplayIndex, isRecommend, categoryId, navigationId, limit int,
		navigationIds []int,
		preLoads, orderBys []string,
	) (list []*model.Novel, err error)
}

type NovelRepository struct{}

// 小说分类创建
func (p *NovelRepository) Create(
	novelTitle, novelAuthor, novelDescription, novelCover, novelBannerCover, mainRole, chapterPreview string,
	novelStatus, sortNum int,
	novelNavigation *model.NovelNavigation,
	novelCategory *model.NovelCategory,
	novelThemes []*model.NovelTheme,
) error {
	novel := model.Novel{
		NovelTitle:       novelTitle,
		NovelAuthor:      novelAuthor,
		NovelDescription: novelDescription,
		NovelCover:       novelCover,
		NovelBannerCover: novelBannerCover,
		MainRole:         mainRole,
		NovelStatus:      novelStatus,
		NovelNavigation:  novelNavigation,
		NovelCategory:    novelCategory,
		ChapterPreview:   chapterPreview,
		SortNum:          sortNum,
		NovelThemes:      novelThemes,
	}
	return global.Gorm.Create(&novel).Error
}

// 小说更新
func (p *NovelRepository) Modify(
	novel *model.Novel,
	novelStatus, sortNum int,
	novelTitle, novelAuthor, novelDescription, novelCover, novelBannerCover, mainRole, chapterPreview string,
	novelNavigation *model.NovelNavigation,
	novelCategory *model.NovelCategory,
	novelThemes []*model.NovelTheme,
) error {
	updates := map[string]interface{}{
		"NovelTitle":       novelTitle,
		"NovelAuthor":      novelAuthor,
		"NovelDescription": novelDescription,
		"NovelCover":       novelCover,
		"NovelBannerCover": novelBannerCover,
		"MainRole":         mainRole,
		"NovelStatus":      novelStatus,
		"NovelNavigation":  novelNavigation,
		"NovelCategory":    novelCategory,
		"ChapterPreview":   chapterPreview,
		"NovelThemes":      novelThemes,
		"SortNum":          sortNum,
	}
	// 清空关联表数据
	var deleteNovelThemes []*model.NovelTheme
	global.Gorm.Model(&novel).Association("NovelThemes").Replace(deleteNovelThemes)
	// 更新所有数据
	err := global.Gorm.Model(&novel).Updates(updates).Error
	return err
}

// 小说删除
func (p *NovelRepository) Delete(novel *model.Novel) error {
	// 清空关联表数据
	var deleteNovelThemes []*model.NovelTheme
	global.Gorm.Model(&novel).Association("NovelThemes").Replace(deleteNovelThemes)
	return global.Gorm.Delete(&novel).Error
}

// 获取单条记录
func (p *NovelRepository) Get(
	novelId int,
	novelTitle string,
	preLoads []string,
) (*model.Novel, error) {
	var novel model.Novel
	query := global.Gorm.Table(model.Novel{}.TableName())
	if preLoads != nil {
		for _, preload := range preLoads {
			query = query.Preload(preload)
		}
	}

	if novelId > 0 {
		query = query.Where("id = ?", novelId)
	}

	if novelTitle != "" {
		query = query.Where("novel_title = ?", novelTitle)
	}
	err := query.First(&novel).Error
	return &novel, err
}

// 小说列表
func (p *NovelRepository) GetList(
	novelTitle string,
	novelStatus int, page, pageSize int,
	preLoads []string,
) (list []*model.Novel, total int, err error) {
	// 分页数据
	page, pageSize, offset := utils.GetPageData(page, pageSize)
	// 查询
	query := global.Gorm.Table(model.Novel{}.TableName())
	if preLoads != nil {
		for _, preload := range preLoads {
			query = query.Preload(preload)
		}
	}

	if novelTitle != "" {
		query = query.Where("novel_title like ?", "%"+novelTitle+"%")
	}

	if novelStatus != model.COMMON_ZERO {
		query = query.Where("novel_status = ?", novelStatus)
	}

	// 总数
	query.Count(&total)
	// 查询
	query.Order("id desc").Offset(offset).Limit(pageSize).Find(&list)
	return list, total, query.Error
}

// 小说列表
func (p *NovelRepository) GetAll(
	selectFields []string,
	isDisplayIndex, isRecommend, categoryId, navigationId, limit int,
	navigationIds []int,
	preLoads, orderBys []string,
) (list []*model.Novel, err error) {
	query := global.Gorm.Table(model.Novel{}.TableName()).Select(selectFields)
	if preLoads != nil {
		for _, preload := range preLoads {
			query = query.Preload(preload)
		}
	}

	if collection.Collect(model.GetCommonValues()).Contains(isDisplayIndex) {
		query = query.Where("is_display_index = ?", isDisplayIndex)
	}

	if collection.Collect(model.GetCommonValues()).Contains(isRecommend) {
		query = query.Where("is_recommend = ?", isRecommend)
	}

	if categoryId > model.COMMON_ZERO {
		query = query.Where("novel_category_id = ?", categoryId)
	}

	if navigationId > model.COMMON_ZERO {
		query = query.Where("novel_navigation_id = ?", navigationId)
	}

	if navigationIds != nil {
		query = query.Where("novel_navigation_id = (?)", navigationIds)
	}

	// 排序
	if orderBys != nil {
		for _, orderBy := range orderBys {
			query = query.Order(orderBy)
		}
	}
	if limit != 0 {
		query = query.Limit(limit)
	}
	query = query.Find(&list)
	return list, query.Error
}
