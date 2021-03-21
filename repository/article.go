package repository

import (
	"ms_novel/global"
	"ms_novel/model"
	"ms_novel/utils"
)

// 定义接口，便于查看所有方法
type NovelArticleRepositoryInterface interface {
	// 创建
	Create(articleTitle, articleAuthor, articleContent string, sortNum int, novel *model.Novel) error

	// 修改
	Modify(article *model.NovelArticle, articleTitle, articleAuthor, articleContent string, sortNum int, novel *model.Novel) error

	// 删除
	Delete(article *model.NovelArticle) error

	// 获取单条
	Get(novelArticleId int, articleTitle string, preLoads []string) (*model.NovelArticle, error)

	// 列表
	GetList(articleTitle string, page, pageSize int, preLoads []string) (list []*model.NovelArticle, total int, err error)
}

type NovelArticleRepository struct{}

// 文章创建
func (p *NovelArticleRepository) Create(articleTitle, articleAuthor, articleContent string, sortNum int, novel *model.Novel) error {
	novelArticle := model.NovelArticle{
		ArticleTitle:   articleTitle,
		ArticleAuthor:  articleAuthor,
		ArticleContent: articleContent,
		Novel:          novel,
		SortNum:        sortNum,
	}
	return global.Gorm.Create(&novelArticle).Error
}

// 文章更新
func (p *NovelArticleRepository) Modify(article *model.NovelArticle, articleTitle, articleAuthor, articleContent string, sortNum int, novel *model.Novel) error {
	updates := map[string]interface{}{
		"ArticleTitle":   articleTitle,
		"ArticleAuthor":  articleAuthor,
		"ArticleContent": articleContent,
		"Novel":          novel,
		"SortNum":        sortNum,
	}
	err := global.Gorm.Model(&article).Updates(updates).Error
	return err
}

// 文章删除
func (p *NovelArticleRepository) Delete(article *model.NovelArticle) error {
	return global.Gorm.Delete(&article).Error
}

// 文章单条记录
func (p *NovelArticleRepository) Get(novelArticleId int, articleTitle string, preLoads []string) (*model.NovelArticle, error) {
	var novelArticle model.NovelArticle
	query := global.Gorm.Model(&model.NovelArticle{})

	if preLoads != nil {
		for _, preload := range preLoads {
			query = query.Preload(preload)
		}
	}

	if novelArticleId > 0 {
		query = query.Where("id = ?", novelArticleId)
	}

	if articleTitle != "" {
		query = query.Where("article_title like ?", "%"+articleTitle+"%")
	}

	err := query.First(&novelArticle).Error
	return &novelArticle, err
}

// 文章列表
func (p *NovelArticleRepository) GetList(articleTitle string, page, pageSize int, preLoads []string) (list []*model.NovelArticle, total int, err error) {
	// 分页数据
	page, pageSize, offset := utils.GetPageData(page, pageSize)

	// 查询
	query := global.Gorm.Table(model.NovelArticle{}.TableName())
	if preLoads != nil {
		for _, preload := range preLoads {
			query = query.Preload(preload)
		}
	}

	if articleTitle != "" {
		query = query.Where("article_title like ?", "%"+articleTitle+"%")
	}

	// 总数
	query.Count(&total)
	// 查询
	query.Order("id desc").Offset(offset).Limit(pageSize).Find(&list)
	return list, total, query.Error
}
