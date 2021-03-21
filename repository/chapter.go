package repository

import (
	"ms_novel/global"
	"ms_novel/model"
	"ms_novel/utils"
)

// 定义小说章节的所有接口，便于查看所有方法
type NovelChapterRepositoryInterface interface {
	// 创建
	Create(chapterTitle string, chapterNo int, novel *model.Novel) error

	// 修改
	Modify(novelChapter *model.NovelChapter, chapterTitle string, chapterNo int, novel *model.Novel) error

	//// 删除
	Delete(novelChapter *model.NovelChapter) error

	// 获取单条记录
	Get(novelChapterId int) (*model.NovelChapter, error)

	// 列表
	GetList(chapterTitle string, page, pageSize int) (list []*model.NovelChapter, total int, err error)
}

type NovelChapterRepository struct{}

// 小说章节创建
func (p *NovelChapterRepository) Create(chapterTitle string, chapterNo int, novel *model.Novel) error {
	novelChapter := model.NovelChapter{
		ChapterTitle: chapterTitle,
		ChapterNo:    chapterNo,
		Novel:        novel,
	}
	return global.Gorm.Create(&novelChapter).Error
}

// 小说章节更新
func (p *NovelChapterRepository) Modify(novelChapter *model.NovelChapter, chapterTitle string, chapterNo int, novel *model.Novel) error {
	updates := map[string]interface{}{
		"ChapterTitle": chapterTitle,
		"ChapterNo":    chapterNo,
		"novel":        novel,
	}
	err := global.Gorm.Model(&novelChapter).Updates(updates).Error
	return err
}

// 章节删除
func (p *NovelChapterRepository) Delete(novelChapter *model.NovelChapter) error {
	return global.Gorm.Delete(&novelChapter).Error
}

// 获取单条记录
func (p *NovelChapterRepository) Get(novelChapterId int) (*model.NovelChapter, error) {
	var novelChapter model.NovelChapter
	query := global.Gorm.Model(&model.NovelCategory{}).Preload("Novel")
	if novelChapterId > 0 {
		query = query.Where("id = ?", novelChapterId)
	}
	err := query.First(&novelChapter).Error
	return &novelChapter, err
}

// 章节列表
func (p *NovelChapterRepository) GetList(chapterTitle string, page, pageSize int) (list []*model.NovelChapter, total int, err error) {
	// 分页数据
	page, pageSize, offset := utils.GetPageData(page, pageSize)
	// 查询
	query := global.Gorm.Table(model.NovelChapter{}.TableName()).Preload("Novel")
	if chapterTitle != "" {
		query = query.Where("chapter_title like ?", "%"+chapterTitle+"%")
	}

	// 总数
	query.Count(&total)
	// 查询
	query.Order("chapter_no asc").Offset(offset).Limit(pageSize).Find(&list)
	return list, total, query.Error
}
