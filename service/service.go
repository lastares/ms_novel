package service

import (
	"github.com/go-sql-driver/mysql"
	"ms_novel/repository"
)

type CommonService struct {
}

const (
	ErrMySQLDupEntry            = 1062 // 唯一索引错误
	ErrMySQLDupEntryWithKeyName = 1586 // 主键冲突
)

// 判断mysql唯一索引错误问题
func (p *CommonService) isUniqueConstraintError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == ErrMySQLDupEntry ||
			mysqlErr.Number == ErrMySQLDupEntryWithKeyName {
			return true
		}
	}
	return false
}

// 小说repository
func (c *CommonService) GetNovelRepository() (repositoryInterface repository.NovelRepositoryInterface) {
	repositoryInterface = new(repository.NovelRepository)
	return
}

// 小说文章repository
func (c *CommonService) GetArticleRepository() (repositoryInterface repository.NovelArticleRepositoryInterface) {
	repositoryInterface = new(repository.NovelArticleRepository)
	return
}

// 小说分类repository
func (c *CommonService) GetCategoryRepository() (repositoryInterface repository.NovelCategoryRepositoryInterface) {
	repositoryInterface = new(repository.NovelCategoryRepository)
	return
}

// 小说友情链接repository
func (c *CommonService) GetFriendLinkRepository() (repositoryInterface repository.NovelFriendLinkRepositoryInterface) {
	repositoryInterface = new(repository.NovelFriendLinkRepository)
	return
}

// 小说导航栏目repository
func (c *CommonService) GetNavigationRepository() (repositoryInterface repository.NavigationRepositoryInterface) {
	repositoryInterface = new(repository.NavigationRepository)
	return
}

// 小说专题repository
func (c *CommonService) GetThemeRepository() (repositoryInterface repository.NovelThemeRepositoryInterface) {
	repositoryInterface = new(repository.NovelThemeRepository)
	return
}

// 小说章节repository
func (c *CommonService) GetChapterRepository() (repositoryInterface repository.NovelChapterRepositoryInterface) {
	repositoryInterface = new(repository.NovelChapterRepository)
	return
}
