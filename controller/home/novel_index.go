package home

import (
	"github.com/gin-gonic/gin"
	"ms_novel/controller"
	"ms_novel/global"
	"ms_novel/request"
	"ms_novel/response"
)

var NovelIndexController = novelIndexController{}

type novelIndexController struct {
	controller.Controller
}

// 导航栏目列表
func (p *novelIndexController) GetNavList(c *gin.Context) {
	// 获取列表
	data, code, err := p.GetNavigationService().GetIndexList()
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 小说banner相关列表
func (p *novelIndexController) GetAllFocusNovel(c *gin.Context) {
	// 获取列表
	data, code, err := p.GetIndexService().GetBannerData()
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 首页推荐的小说数据
func (p *novelIndexController) GetAllRecommendList(c *gin.Context) {
	// 获取列表
	data, code, err := p.GetIndexService().GetAllRecommendList()
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 根据导航栏目获取小说数据
func (p *novelIndexController) GetNovelByNavId(c *gin.Context) {
	// 获取列表
	data, code, err := p.GetIndexService().GetNavByNavigationId()
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 根据导航栏目获取小说数据接口的第二种方法
func (p *novelIndexController) GetNovelByNavId2(c *gin.Context) {
	// 获取列表
	data, code, err := p.GetIndexService().GetNavByNavigationId2()
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 首页：获取分类下的小说数据
func (p *novelIndexController) GetNovelByCategory(c *gin.Context) {
	// 获取列表
	data, code, err := p.GetIndexService().GetNovelByCategory()
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 首页：获取友情链接
func (p *novelIndexController) GetFriendLinkList(c *gin.Context) {
	// 获取列表
	data, code, err := p.GetIndexService().GetFriendLinkList()
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 首页：右边小说数据
func (p *novelIndexController) GetRightNovelData(c *gin.Context) {
	// 参数绑定结构体
	var form request.IndexRightNovel
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 获取列表
	data, code, err := p.GetIndexService().GetRightNovelList(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}
