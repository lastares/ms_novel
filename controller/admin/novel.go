package admin

import (
	"github.com/gin-gonic/gin"
	"ms_novel/controller"
	"ms_novel/global"
	"ms_novel/request"
	"ms_novel/response"
)

var NovelController = novelController{}

type novelController struct {
	controller.Controller
}

// @Summary 创建
// @Tags 小说（Novel）
// @Accept application/json
// @Produce application/json
// @Param object body request.NovelCreate true "参数"
// @Success 200 {object} response.Succeed
// @Router /novel/create [post]
func (p *novelController) Create(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelCreate
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	code, err := p.GetNovelService().Create(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// @Summary 修改
// @Tags 小说（Novel）
// @Accept application/json
// @Produce application/json
// @Param object body request.NovelModify true "参数"
// @Success 200 {object} response.Succeed
// @Router /novel/modify [post]
func (p *novelController) Modify(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelModify
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 修改
	code, err := p.GetNovelService().Modify(form)

	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}
	response.Json(c)
}

// @Summary 删除
// @Tags 小说（Novel）
// @Accept application/json
// @Produce application/json
// @Param object body request.NovelDelete true "参数"
// @Success 200 {object} response.Succeed
// @Router /novel/delete [post]
func (p *novelController) Delete(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelDelete
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 删除
	code, err := p.GetNovelService().Delete(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.Json(c)
}

// @Summary 详情
// @Tags 小说（Novel）
// @Accept application/json
// @Produce application/json
// @Param object body request.NovelDetail true "参数"
// @Success 200 {object} novelDTO.NovelDTO
// @Router /novel/get-detail [post]
func (p *novelController) GetDetail(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelDetail
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	data, code, err := p.GetNovelService().GetDetail(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

func (p *novelController) GetDetail2(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelDetail
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	data, code, err := p.GetNovelService().GetDetail2(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}

// @Summary 列表
// @Tags 小说（Novel）
// @Accept application/json
// @Produce application/json
// @Param object body request.NovelList true "参数"
// @Success 200 {object} novelDTO.NovelAdminListDTO
// @Router /novel/get-list [post]
func (p *novelController) GetList(c *gin.Context) {
	// 参数绑定结构体
	var form request.NovelList
	c.ShouldBindJSON(&form)

	// 校验结构体参数
	err := global.Validate.Struct(form)
	if err != nil {
		request.FormValidateException(c, err)
		return
	}

	// 获取列表
	data, code, err := p.GetNovelService().GetList(form)
	if err != nil {
		response.CustomerException(err.Error(), c, code)
		return
	}

	response.JsonData(c, data)
}
