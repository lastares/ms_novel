package routers

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"ms_novel/controller/admin"
	"ms_novel/controller/home"
	_ "ms_novel/docs"
	"ms_novel/response"
	"net/http"
)

func InitRouter(router *gin.Engine) *gin.Engine {
	// swagger
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 后台相关接口
	novel := router.Group("/novel")
	{
		// 小说栏目导航
		novel.POST("/navigation/create", admin.NovelNavigationController.Create)
		novel.POST("/navigation/modify", admin.NovelNavigationController.Modify)
		novel.POST("/navigation/delete", admin.NovelNavigationController.Delete)
		novel.POST("/navigation/get-detail", admin.NovelNavigationController.GetDetail)
		novel.POST("/navigation/get-list", admin.NovelNavigationController.GetList)

		// 小说分类
		novel.POST("/category/create", admin.NovelCategoryController.Create)
		novel.POST("/category/modify", admin.NovelCategoryController.Modify)
		novel.POST("/category/delete", admin.NovelCategoryController.Delete)
		novel.POST("/category/get-list", admin.NovelCategoryController.GetList)
		novel.POST("/category/get-detail", admin.NovelCategoryController.GetDetail)

		// 小说专题
		novel.POST("/theme/create", admin.NovelThemeController.Create)
		novel.POST("/theme/modify", admin.NovelThemeController.Modify)
		novel.POST("/theme/delete", admin.NovelThemeController.Delete)
		novel.POST("/theme/get-detail", admin.NovelThemeController.GetDetail)
		novel.POST("/theme/get-list", admin.NovelThemeController.GetList)

		// 小说友情链接
		novel.POST("/friend-link/create", admin.NovelFriendLinkController.Create)
		novel.POST("/friend-link/modify", admin.NovelFriendLinkController.Modify)
		novel.POST("/friend-link/delete", admin.NovelFriendLinkController.Delete)
		novel.POST("/friend-link/get-detail", admin.NovelFriendLinkController.GetDetail)
		novel.POST("/friend-link/get-list", admin.NovelFriendLinkController.GetList)

		// 小说文章资讯
		novel.POST("/article/create", admin.NovelArticleController.Create)
		novel.POST("/article/get-detail", admin.NovelArticleController.GetDetail)
		novel.POST("/article/modify", admin.NovelArticleController.Modify)
		novel.POST("/article/delete", admin.NovelArticleController.Delete)
		novel.POST("/article/get-list", admin.NovelArticleController.GetList)

		// 小说
		novel.POST("/create", admin.NovelController.Create)
		novel.POST("/modify", admin.NovelController.Modify)
		novel.POST("/delete", admin.NovelController.Delete)
		novel.POST("/get-detail", admin.NovelController.GetDetail)
		novel.POST("/get-detail2", admin.NovelController.GetDetail2)
		novel.POST("/get-list", admin.NovelController.GetList)

		// 小说章节
		novel.POST("/chapter/create", admin.NovelChapterController.Create)
		novel.POST("/chapter/modify", admin.NovelChapterController.Modify)
		novel.POST("/chapter/get-detail", admin.NovelChapterController.GetDetail)
		novel.POST("/chapter/get-list", admin.NovelChapterController.GetList)
		novel.POST("/chapter/delete", admin.NovelChapterController.Delete)
	}

	// 前台相关接口
	index := router.Group("/home")
	{
		// 导航数据接口
		index.POST("/novel/navigation/get-list", home.NovelIndexController.GetNavList)
		// 首页第一块相关数据
		index.POST("/novel/get-all-focus", home.NovelIndexController.GetAllFocusNovel)
		// 获取推荐的小说
		index.POST("/novel/get-all-recommend", home.NovelIndexController.GetAllRecommendList)
		// 获取导航栏目下的小说
		index.POST("/novel/get-all-by-nav", home.NovelIndexController.GetNovelByNavId)
		index.POST("/novel/get-all-by-nav2", home.NovelIndexController.GetNovelByNavId2)
		// 获取分类下的小说
		index.POST("/novel/get-all-by-category", home.NovelIndexController.GetNovelByCategory)
		// 获取友情链接
		index.POST("/novel/get-friend-link", home.NovelIndexController.GetFriendLinkList)
		// 获取右边小说数据
		index.POST("/novel/get-right-novel", home.NovelIndexController.GetRightNovelData)
	}

	// 未知路由处理
	router.NoRoute(func(c *gin.Context) {
		response.CustomerException(
			"No router.",
			c,
			http.StatusNotFound,
		)
		return
	})
	// 未知方法处理
	router.NoMethod(func(c *gin.Context) {
		response.CustomerException(
			"No method.",
			c,
			http.StatusMethodNotAllowed,
		)
		return
	})
	return router
}
