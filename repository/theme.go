package repository

import (
	"ms_novel/global"
	"ms_novel/model"
	"ms_novel/utils"
)

// 定义小说专题的接口，便于查看所有方法
type NovelThemeRepositoryInterface interface {
	// 创建
	Create(themeName, themeSubTitle, themeDescription, themeCover string, sortNum int) error

	// 修改
	Modify(theme *model.NovelTheme, themeName, themeSubTitle, themeDescription, themeCover string, sortNum int) error

	// 删除
	Delete(theme *model.NovelTheme) error

	// 获取单条分类
	Get(themeId int, themeName string, themeSubTitle string) (*model.NovelTheme, error)

	// 列表
	GetList(ThemeName, ThemeSubtitle string, page, pageSize int) (list []*model.NovelTheme, total int, err error)

	// 获取所有，不分页
	GetAllList(novelThemeIds []int, preLoads []string) (list []*model.NovelTheme, total int, err error)
}

type NovelThemeRepository struct{}

// 小说分类创建
func (p *NovelThemeRepository) Create(
	themeName, themeSubTitle, themeDescription, themeCover string,
	sortNum int,
) error {
	novelTheme := model.NovelTheme{
		ThemeName:        themeName,
		ThemeSubtitle:    themeSubTitle,
		ThemeDescription: themeDescription,
		ThemeCover:       themeCover,
		SortNum:          sortNum,
	}
	return global.Gorm.Create(&novelTheme).Error
}

// 小说主题更新
func (p *NovelThemeRepository) Modify(
	theme *model.NovelTheme,
	themeName, themeSubTitle, themeDescription, themeCover string,
	sortNum int,
) error {
	updates := map[string]interface{}{
		"ThemeName":        themeName,
		"ThemeSubtitle":    themeSubTitle,
		"ThemeDescription": themeDescription,
		"ThemeCover":       themeCover,
		"SortNum":          sortNum,
	}
	err := global.Gorm.Model(&theme).Updates(updates).Error
	return err
}

// 主题删除
func (p *NovelThemeRepository) Delete(theme *model.NovelTheme) error {
	return global.Gorm.Delete(&theme).Error
}

// 获取专题单条记录
func (p *NovelThemeRepository) Get(themeId int, themeName string, themeSubTitle string) (*model.NovelTheme, error) {
	var theme model.NovelTheme
	query := global.Gorm.Model(&model.NovelTheme{})
	if themeId > 0 {
		query = query.Where("id = ?", themeId)
	}

	if themeName != "" {
		query = query.Where("theme_name = ?", themeName)
	}
	if themeSubTitle != "" {
		query = query.Where("theme_subtitle = ?", themeSubTitle)
	}
	err := query.First(&theme).Error
	return &theme, err
}

// 专题列表
func (p *NovelThemeRepository) GetList(
	ThemeName, ThemeSubtitle string,
	page, pageSize int,
) (list []*model.NovelTheme, total int, err error) {
	// 分页数据
	page, pageSize, offset := utils.GetPageData(page, pageSize)
	// 查询
	query := global.Gorm.Table(model.NovelTheme{}.TableName())
	if ThemeName != "" {
		query = query.Where("theme_name like ?", "%"+ThemeName+"%")
	}

	if ThemeSubtitle != "" {
		query = query.Where("theme_subtitle like ?", "%"+ThemeSubtitle+"%")
	}

	// 总数
	query.Count(&total)
	// 查询
	query.Order("id desc").Offset(offset).Limit(pageSize).Find(&list)
	return list, total, query.Error
}

// 获取所有专题，部分页
func (p *NovelThemeRepository) GetAllList(
	novelThemeIds []int,
	preLoads []string,
) (list []*model.NovelTheme, total int, err error) {
	// 查询
	query := global.Gorm.Table(model.NovelTheme{}.TableName())
	if preLoads != nil {
		for _, value := range preLoads {
			query = query.Preload(value)
		}
	}

	if len(novelThemeIds) > 0 {
		query = query.Where("id in (?)", novelThemeIds)
	}

	// 总数
	query.Count(&total)
	// 查询
	query.Order("id desc").Find(&list)
	return list, total, query.Error
}
