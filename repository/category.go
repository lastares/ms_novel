package repository

import (
	"github.com/chenhg5/collection"
	"ms_novel/global"
	"ms_novel/model"
	"ms_novel/utils"
)

// 定义小说导航栏目的仓储接口，便于查看所有方法
type NovelCategoryRepositoryInterface interface {
	// 创建
	Create(categoryName, categoryIcon string, isDisplayIndex, sortNum int, novelNavigation *model.NovelNavigation) error

	// 修改
	Modify(category *model.NovelCategory, categoryName, categoryIcon string, isDisplayIndex, sortNum int, novelNavigation *model.NovelNavigation) error

	// 删除
	Delete(category *model.NovelCategory) error

	// 获取单条分类
	Get(categoryId int, categoryName string, preLoads []string) (*model.NovelCategory, error)

	// 列表
	GetList(categoryName string, page, pageSize int) (list []*model.NovelCategory, total int, err error)

	// 获取所有数据
	GetAll(isDisplayIndex int, preLoads []string) (list []*model.NovelCategory, err error)
}

type NovelCategoryRepository struct{}

// 小说分类创建
func (p *NovelCategoryRepository) Create(categoryName, categoryIcon string, isDisplayIndex, sortNum int, novelNavigation *model.NovelNavigation) error {
	novelCategory := model.NovelCategory{
		CategoryName:    categoryName,
		CategoryIcon:    categoryIcon,
		IsDisplayIndex:  isDisplayIndex,
		NovelNavigation: novelNavigation,
	}
	return global.Gorm.Create(&novelCategory).Error
}

// 小说分类更新
func (p *NovelCategoryRepository) Modify(
	category *model.NovelCategory,
	categoryName, categoryIcon string,
	isDisplayIndex, sortNum int,
	novelNavigation *model.NovelNavigation,
) error {
	updates := map[string]interface{}{
		"CategoryName":    categoryName,
		"CategoryIcon":    categoryIcon,
		"IsDisplayIndex":  isDisplayIndex,
		"NovelNavigation": novelNavigation,
		"SortNum":         sortNum,
	}
	err := global.Gorm.Model(&category).Updates(updates).Error
	return err
}

// 分类删除
func (p *NovelCategoryRepository) Delete(category *model.NovelCategory) error {
	return global.Gorm.Delete(&category).Error
}

// 导航栏目获取单条记录
func (p *NovelCategoryRepository) Get(
	categoryId int,
	categoryName string,
	preLoads []string,
) (*model.NovelCategory, error) {
	var category model.NovelCategory
	query := global.Gorm.Model(&model.NovelCategory{})
	if preLoads != nil {
		for _, preload := range preLoads {
			query = query.Preload(preload)
		}
	}

	if categoryId > 0 {
		query = query.Where("id = ?", categoryId)
	}

	if categoryName != "" {
		query = query.Where("category_name = ?", categoryName)
	}
	err := query.First(&category).Error
	return &category, err
}

// 分类列表
func (p *NovelCategoryRepository) GetList(
	categoryName string,
	page, pageSize int,
) (list []*model.NovelCategory, total int, err error) {
	// 分页数据
	page, pageSize, offset := utils.GetPageData(page, pageSize)
	// 查询
	query := global.Gorm.Table(model.NovelCategory{}.TableName()).Preload("NovelNavigation")
	if categoryName != "" {
		query = query.Where("category_name like ?", "%"+categoryName+"%")
	}

	// 总数
	query.Count(&total)
	// 查询
	query.Order("id desc").Offset(offset).Limit(pageSize).Find(&list)
	return list, total, query.Error
}

// 分类列表,不分页
func (p *NovelCategoryRepository) GetAll(isDisplayIndex int, preLoads []string) (list []*model.NovelCategory, err error) {
	// 查询
	query := global.Gorm.Table(model.NovelCategory{}.TableName())
	if preLoads != nil {
		for _, preload := range preLoads {
			query = query.Preload(preload)
		}
	}

	if collection.Collect(model.GetCommonValues()).Contains(isDisplayIndex) {
		query = query.Where("is_display_index = ?", isDisplayIndex)
	}

	// 查询
	query.Order("id desc").Find(&list)
	return list, query.Error
}
