package repository

import (
	"ms_novel/global"
	"ms_novel/model"
	"ms_novel/utils"
)

// 定义小说专题的接口，便于查看所有方法
type NovelFriendLinkRepositoryInterface interface {
	// 创建
	Create(linkTitle, link string, sortNum int) error
	// 修改
	Modify(friendLink *model.FriendLink, linkTitle, link string, sortNum int) error
	// 删除
	Delete(theme *model.FriendLink) error

	// 获取单条
	Get(friendLinkId int, linkTitle string) (*model.FriendLink, error)
	// 列表
	GetList(linkTitle string, page, pageSize int) (list []*model.FriendLink, total int, err error)

	GetAll() (list []*model.FriendLink, err error)
}

type NovelFriendLinkRepository struct{}

// 友情链接创建
func (p *NovelFriendLinkRepository) Create(linkTitle, link string, sortNum int) error {
	friendLink := model.FriendLink{
		LinkTitle: linkTitle,
		Link:      link,
		SortNum:   sortNum,
	}
	return global.Gorm.Create(&friendLink).Error
}

// 友情链接更新
func (p *NovelFriendLinkRepository) Modify(friendLink *model.FriendLink, linkTitle, link string, sortNum int) error {
	updates := map[string]interface{}{
		"LinkTitle": linkTitle,
		"Link":      link,
		"SortNum":   sortNum,
	}
	err := global.Gorm.Model(&friendLink).Updates(updates).Error
	return err
}

// 友情链接删除
func (p *NovelFriendLinkRepository) Delete(friendLink *model.FriendLink) error {
	return global.Gorm.Delete(&friendLink).Error
}

// 友情链接单条记录
func (p *NovelFriendLinkRepository) Get(friendLinkId int, linkTitle string) (*model.FriendLink, error) {
	var friendLink model.FriendLink
	query := global.Gorm.Model(&model.FriendLink{})
	if friendLinkId > 0 {
		query = query.Where("id = ?", friendLinkId)
	}

	if linkTitle != "" {
		query = query.Where("link_title like ?", "%"+linkTitle+"%")
	}

	err := query.First(&friendLink).Error
	return &friendLink, err
}

// 友情链接列表
func (p *NovelFriendLinkRepository) GetList(linkTitle string, page, pageSize int) (list []*model.FriendLink, total int, err error) {
	// 分页数据
	page, pageSize, offset := utils.GetPageData(page, pageSize)
	// 查询
	query := global.Gorm.Table(model.FriendLink{}.TableName())
	if linkTitle != "" {
		query = query.Where("link_title like ?", "%"+linkTitle+"%")
	}

	// 总数
	query.Count(&total)
	// 查询
	query.Order("id desc").Offset(offset).Limit(pageSize).Find(&list)
	return list, total, query.Error
}

// 获取列表
func (p *NovelFriendLinkRepository) GetAll() (list []*model.FriendLink, err error) {
	// 查询
	query := global.Gorm.Table(model.FriendLink{}.TableName())

	// 查询
	query = query.Select([]string{"id, link_title, link, sort_num"}).Order("sort_num desc").Order("id desc").Find(&list)

	return list, query.Error
}
