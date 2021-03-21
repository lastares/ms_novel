package admin

import (
	"github.com/gin-gonic/gin"
	"ms_novel/controller"
	"ms_novel/global"
	"ms_novel/request"
	"ms_novel/response"
)

var NovelChapterController = novelChapterController{}

type novelChapterController struct {
	controller.Controller
}

// 专题创建
func (p *novelChapterController) Create(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelChapterCreate
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	code, err := p.GetChapterService().Create(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 章节修改
func (p *novelChapterController) Modify(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelChapterModify
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 修改
	code, err := p.GetChapterService().Modify(form)

	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}
	response.Json(c)
}

// 专题删除
func (p *novelChapterController) Delete(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelChapterDelete
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 删除
	code, err := p.GetChapterService().Delete(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 章节详情
func (p *novelChapterController) GetDetail(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelChapterDetail
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	data, code, err := p.GetChapterService().GetDetail(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 章节列表
func (p *novelChapterController) GetList(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelChapterList
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 获取列表
	data, code, err := p.GetChapterService().GetList(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}
