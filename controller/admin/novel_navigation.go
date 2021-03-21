package admin

import (
	"github.com/gin-gonic/gin"
	"ms_novel/controller"
	"ms_novel/global"
	"ms_novel/request"
	"ms_novel/response"
)

var NovelNavigationController = novelNavigationController{}

type novelNavigationController struct {
	controller.Controller
}

// 导航栏目创建
func (p *novelNavigationController) Create(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelNavigationCreate
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}
	// 调用service,创建栏目
	code, err := p.GetNavigationService().Create(form)

	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}
	response.Json(c)
}

// 导航栏目修改
func (p *novelNavigationController) Modify(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelNavigationModify
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 修改
	code, err := p.GetNavigationService().Modify(form)

	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}
	response.Json(c)
}

// 导航栏目删除
func (p *novelNavigationController) Delete(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelNavigationDelete
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 删除
	code, err := p.GetNavigationService().Delete(form.NavigationId)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 导航栏目详情
func (p *novelNavigationController) GetDetail(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelNavigationDetail
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	data, code, err := p.GetNavigationService().Get(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 导航栏目列表
func (p *novelNavigationController) GetList(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelNavigationList
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 获取列表
	data, code, err := p.GetNavigationService().GetList(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}
