package service

import (
	"errors"
	"ms_novel/DTO/categoryDTO"
	"ms_novel/request"
	"ms_novel/response"
)

// 定义接口，便于查看所有方法
type NovelCategoryInterface interface {
	Create(form request.NovelCategoryCreate) (code int, err error)
	Modify(form request.NovelCategoryModify) (code int, err error)
	Delete(form request.NovelCategoryDelete) (code int, err error)

	GetDetail(form request.NovelCategoryDetail) (data *categoryDTO.NovelCategoryCopy, code int, err error)
	GetList(form request.NovelCategoryList) (data categoryDTO.NovelCategoryAdminListDTO, code int, err error)
}

type NovelCategoryService struct {
	CommonService
}

// 分类创建
func (p *NovelCategoryService) Create(
	form request.NovelCategoryCreate,
) (code int, err error) {
	// 判断分类所属导航栏目是否存在
	novelNavigation, _ := p.GetNavigationRepository().Get(
		form.NovelNavigationId,
		"",
		nil,
	)
	if novelNavigation.Id == 0 {
		return response.Failure, errors.New("Navigation not existed.")
	}
	// 判断该分类是否存在
	novelCategory, _ := p.GetCategoryRepository().Get(
		0,
		form.CategoryName,
		nil,
	)
	// 如果查到数据了，说明该分类已存在
	if novelCategory.Id > 0 {
		return response.Failure, errors.New("Category existed.")
	}

	// 添加
	err = p.GetCategoryRepository().Create(
		form.CategoryName,
		form.CategoryIcon,
		form.IsDisplayIndex,
		form.SortNum,
		novelNavigation,
	)
	if err != nil {
		return response.Failure, errors.New("Category creation failed.")
	}
	return response.Ok, nil
}

// 分类修改
func (p *NovelCategoryService) Modify(
	form request.NovelCategoryModify,
) (code int, err error) {
	// 查询分类
	category, err := p.GetCategoryRepository().Get(
		form.NovelCategoryId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Category not existed.")
	}

	// 判断分类所属导航栏目是否存在
	novelNavigation, _ := p.GetNavigationRepository().Get(
		form.NovelNavigationId,
		"",
		nil,
	)
	if novelNavigation.Id == 0 {
		return response.Failure, errors.New("Navigation not existed.")
	}

	category.NovelNavigationId = form.NovelNavigationId

	// 更新
	err = p.GetCategoryRepository().Modify(
		category,
		form.CategoryName,
		form.CategoryIcon,
		form.IsDisplayIndex,
		form.SortNum,
		novelNavigation,
	)

	if err != nil {
		// 判断是否是触发了msql的唯一性索引的错误，如果是则提示分类存在，其他错误走下面的异常抛出
		if p.isUniqueConstraintError(err) {
			return response.Failure, errors.New("Category existed.")
		}
		return response.Failure, errors.New("Category modify failed.")
	}
	return response.Ok, nil
}

// 分类删除
func (p *NovelCategoryService) Delete(form request.NovelCategoryDelete) (code int, err error) {
	category, err := p.GetCategoryRepository().Get(
		form.NovelCategoryId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Category not existed.")
	}
	// 删除
	err = p.GetCategoryRepository().Delete(category)
	if err != nil {
		return response.Failure, errors.New("Category delete failed.")
	}
	return response.Ok, nil
}

// 分类列表
func (p *NovelCategoryService) GetList(form request.NovelCategoryList) (result categoryDTO.NovelCategoryAdminListDTO, code int, err error) {
	// 获取源数据
	list, total, err := p.GetCategoryRepository().GetList(
		form.CategoryName,
		form.Page,
		form.PageSize,
	)

	if err != nil {
		return categoryDTO.NovelCategoryAdminListDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var data []*categoryDTO.NovelCategoryCopy
	for _, novelCategory := range list {
		data = append(data, &categoryDTO.NovelCategoryCopy{*novelCategory})
	}

	// 返回DTO对象
	result = categoryDTO.NovelCategoryAdminListDTO{
		List:  data,
		Total: total,
	}
	return result, response.Ok, nil
}

// 分类详情
func (p *NovelCategoryService) GetDetail(form request.NovelCategoryDetail) (data *categoryDTO.NovelCategoryCopy, code int, err error) {
	novelCategory, err := p.GetCategoryRepository().Get(
		form.NovelCategoryId,
		"",
		[]string{"NovelNavigation"},
	)
	if err != nil {
		return &categoryDTO.NovelCategoryCopy{}, response.Failure, errors.New("Category not existed.")
	}
	return &categoryDTO.NovelCategoryCopy{*novelCategory}, response.Ok, nil
}
