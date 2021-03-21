package admin

import (
	"github.com/gin-gonic/gin"
	"ms_novel/controller"
	"ms_novel/global"
	"ms_novel/request"
	"ms_novel/response"
)

var NovelCategoryController = novelCategoryController{}

type novelCategoryController struct {
	controller.Controller
}

// 小说分类创建
func (p *novelCategoryController) Create(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelCategoryCreate
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 调用service,创建栏目
	code, err := p.GetCategoryService().Create(form)

	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 导航栏目修改
func (p *novelCategoryController) Modify(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelCategoryModify
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	//创建栏目
	code, err := p.GetCategoryService().Modify(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 导航栏目删除
func (p *novelCategoryController) Delete(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelCategoryDelete
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 删除
	code, err := p.GetCategoryService().Delete(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 分类详情
func (p *novelCategoryController) GetDetail(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelCategoryDetail
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	data, code, err := p.GetCategoryService().GetDetail(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(
		c,
		data,
	)
}

// 分类列表
func (p *novelCategoryController) GetList(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelCategoryList
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	data, code, err := p.GetCategoryService().GetList(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(
		c,
		data,
	)
}
