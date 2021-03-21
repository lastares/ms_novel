package admin

import (
	"github.com/gin-gonic/gin"
	"ms_novel/controller"
	"ms_novel/global"
	"ms_novel/request"
	"ms_novel/response"
)

var NovelArticleController = novelArticleController{}

type novelArticleController struct {
	controller.Controller
}

// 文章创建
func (p *novelArticleController) Create(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelArticleCreate
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	code, err := p.GetArticleService().Create(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 文章修改
func (p *novelArticleController) Modify(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelArticleModify
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 修改
	code, err := p.GetArticleService().Modify(form)

	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}
	response.Json(c)
}

// 文章删除
func (p *novelArticleController) Delete(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelArticleDelete
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 删除
	code, err := p.GetArticleService().Delete(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 文章详情
func (p *novelArticleController) GetDetail(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelArticleDetail
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	data, code, err := p.GetArticleService().GetDetail(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 文章列表
func (p *novelArticleController) GetList(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelArticleList
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 获取列表
	data, code, err := p.GetArticleService().GetList(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}
