package repository

import (
	"github.com/chenhg5/collection"
	"ms_novel/global"
	"ms_novel/model"
	"ms_novel/utils"
)

// 定义小说导航栏目的接口，便于查看所有方法
type NavigationRepositoryInterface interface {
	// 创建
	Create(navigationName string, sortNum, isDisplayIndex int) error

	// 更新
	Modify(navigation *model.NovelNavigation, navigationName string, sortNum, isDisplayIndex int) error

	// 删除
	Delete(navigation *model.NovelNavigation) error

	// 获取单条记录
	Get(navigationId int, navigationName string, preLoads []string) (*model.NovelNavigation, error)

	// 后台列表
	GetList(page int, pageSize int, preLoads []string) (list []*model.NovelNavigation, total int, err error)

	// 获取所有
	GetAll(isDisplayIndex int, preLoads []string) (list []*model.NovelNavigation, err error)
}

type NavigationRepository struct{}

// 小说导航栏目创建
func (p *NavigationRepository) Create(
	navigationName string,
	sortNum, isDisplayIndex int,
) error {
	navigation := model.NovelNavigation{
		NavigationName: navigationName,
		SortNum:        sortNum,
		IsDisplayIndex: isDisplayIndex,
	}
	return global.Gorm.Create(&navigation).Error
}

// 小说导航栏目更新
func (p *NavigationRepository) Modify(
	navigation *model.NovelNavigation,
	navigationName string,
	sortNum, isDisplayIndex int,
) error {
	updates := map[string]interface{}{
		"NavigationName": navigationName,
		"SortNum":        sortNum,
		"IsDisplayIndex": isDisplayIndex,
	}
	err := global.Gorm.Model(&navigation).Updates(updates).Error
	return err
}

// 导航栏目获取单条记录
func (p *NavigationRepository) Get(
	navigationId int,
	navigationName string,
	preLoads []string,
) (*model.NovelNavigation, error) {
	var navigation model.NovelNavigation
	query := global.Gorm.Model(&model.NovelNavigation{})
	if preLoads != nil {
		for _, preload := range preLoads {
			query = query.Preload(preload)
		}
	}

	if navigationId > 0 {
		query = query.Where("id = ?", navigationId)
	}

	if navigationName != "" {
		query = query.Where("navigation_name = ?", navigationName)
	}
	err := query.First(&navigation).Error
	return &navigation, err
}

// 导航栏目删除
func (p *NavigationRepository) Delete(navigation *model.NovelNavigation) error {
	return global.Gorm.Delete(&navigation).Error
}

// 导航栏目列表
func (p *NavigationRepository) GetList(page int, pageSize int, preLoads []string) (list []*model.NovelNavigation, total int, err error) {
	// 分页数据
	page, pageSize, offset := utils.GetPageData(page, pageSize)
	// 查询
	query := global.Gorm.Table(model.NovelNavigation{}.TableName())
	if preLoads != nil {
		for _, preload := range preLoads {
			query = query.Preload(preload)
		}
	}

	// 总数
	query.Count(&total)
	// 查询
	query.Order("id desc").Offset(offset).Limit(pageSize).Find(&list)

	return list, total, query.Error
}

// 导航栏目列表
func (p *NavigationRepository) GetAll(isDisplayIndex int, preLoads []string) (list []*model.NovelNavigation, err error) {
	// 查询
	query := global.Gorm.Table(model.NovelNavigation{}.TableName())
	if preLoads != nil {
		for _, preload := range preLoads {
			query = query.Preload(preload)
		}
	}

	if collection.Collect(model.GetCommonValues()).Contains(isDisplayIndex) {
		query = query.Where("is_display_index = ?", isDisplayIndex)
	}

	// 查询
	query.Order("sort_num desc").Find(&list)

	return list, query.Error
}
