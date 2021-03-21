package admin

import (
	"github.com/gin-gonic/gin"
	"ms_novel/controller"
	"ms_novel/global"
	"ms_novel/request"
	"ms_novel/response"
)

var NovelFriendLinkController = friendLinkController{}

type friendLinkController struct {
	controller.Controller
}

// 友情链接创建
func (p *friendLinkController) Create(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelFriendLinkCreate
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	code, err := p.GetFriendLinkService().Create(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 友情链接修改
func (p *friendLinkController) Modify(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelFriendLinkModify
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 修改
	code, err := p.GetFriendLinkService().Modify(form)

	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}
	response.Json(c)
}

// 友情链接删除
func (p *friendLinkController) Delete(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelFriendLinkDelete
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 删除
	code, err := p.GetFriendLinkService().Delete(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// 友情链接详情
func (p *friendLinkController) GetDetail(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelFriendLinkDetail
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	data, code, err := p.GetFriendLinkService().GetDetail(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// 友情链接列表
func (p *friendLinkController) GetList(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelFriendLinkList
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 获取列表
	data, code, err := p.GetFriendLinkService().GetList(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}
