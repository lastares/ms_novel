package controller

import "ms_novel/service"

type Controller struct {
}

/*
* 相关service定义
 */
// 首页service
func (c *Controller) GetIndexService() (serviceInstance service.IndexServiceInterface) {
	serviceInstance = new(service.IndexService)
	return
}

// 小说service
func (c *Controller) GetNovelService() (serviceInstance service.NovelServiceInterface) {
	serviceInstance = new(service.NovelService)
	return
}

// 小说文章service
func (c *Controller) GetArticleService() (serviceInstance service.NovelArticleServiceInterface) {
	serviceInstance = new(service.NovelArticleService)
	return
}

// 小说分类service
func (c *Controller) GetCategoryService() (serviceInstance service.NovelCategoryInterface) {
	serviceInstance = new(service.NovelCategoryService)
	return
}

// 小说章节service
func (c *Controller) GetChapterService() (serviceInstance service.NovelChapterServiceInterface) {
	serviceInstance = new(service.NovelChapterService)
	return
}

// 小说友情链接service
func (c *Controller) GetFriendLinkService() (serviceInstance service.NovelFriendLinkServiceInterface) {
	serviceInstance = new(service.NovelFriendLinkService)
	return
}

// 小说导航栏目service
func (c *Controller) GetNavigationService() (serviceInstance service.NovelNavigationInterface) {
	serviceInstance = new(service.NovelNavigationService)
	return
}

// 小说专题service
func (c *Controller) GetThemeService() (serviceInstance service.NovelThemeServiceInterface) {
	serviceInstance = new(service.NovelThemeService)
	return
}
