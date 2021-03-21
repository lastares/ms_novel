package admin

import (
	"github.com/gin-gonic/gin"
	"ms_novel/controller"
	"ms_novel/global"
	"ms_novel/request"
	"ms_novel/response"
)

var NovelThemeController = novelThemeController{}

type novelThemeController struct {
	controller.Controller
}

// 专题创建
func (p *novelThemeController) Create(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelThemeCreate
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	code, err := p.GetThemeService().Create(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 专题修改
func (p *novelThemeController) Modify(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelThemeModify
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 修改
	code, err := p.GetThemeService().Modify(form)

	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}
	response.Json(c)
}

// 专题删除
func (p *novelThemeController) Delete(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelThemeDelete
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 删除
	code, err := p.GetThemeService().Delete(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 专题详情
func (p *novelThemeController) GetDetail(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelThemeDetail
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	data, code, err := p.GetThemeService().GetDetail(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 专题列表
func (p *novelThemeController) GetList(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelThemeList
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 获取列表
	data, code, err := p.GetThemeService().GetList(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}
