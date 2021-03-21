package service

import (
	"errors"
	"ms_novel/DTO/navigationDTO"
	"ms_novel/model"
	"ms_novel/request"
	"ms_novel/response"
)

// 定义导航栏目接口，便于查看所有方法
type NovelNavigationInterface interface {
	// 创建
	Create(form request.NovelNavigationCreate) (code int, err error)

	// 修改
	Modify(form request.NovelNavigationModify) (code int, err error)

	// 删除
	Delete(navigationId int) (code int, err error)

	// 详情
	Get(form request.NovelNavigationDetail) (data *navigationDTO.NovelNavigationCopy, code int, err error)

	// 后台列表
	GetList(form request.NovelNavigationList) (result navigationDTO.NovelNavigationAdminListDTO, code int, err error)

	// 前台导航栏目数据
	GetIndexList() (result navigationDTO.NovelNavigationIndexListDTO, code int, err error)
}

type NovelNavigationService struct {
	CommonService
}

// 导航栏目创建
func (p *NovelNavigationService) Create(form request.NovelNavigationCreate) (code int, err error) {
	// 判断改导航栏目是否存在
	novelNavigation, err := p.GetNavigationRepository().Get(
		0,
		form.NavigationName,
		nil,
	)
	if novelNavigation.Id > 0 {
		return response.Failure, errors.New("Navigation existed.")
	}

	err = p.GetNavigationRepository().Create(
		form.NavigationName,
		form.SortNum,
		form.IsDisplayIndex,
	)
	if err != nil {
		return response.Failure, errors.New("Navigation create failed.")
	}
	return response.Ok, nil
}

// 导航栏目修改
func (p *NovelNavigationService) Modify(form request.NovelNavigationModify) (code int, err error) {
	novelNavigation, err := p.GetNavigationRepository().Get(
		form.NavigationId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Navigation not existed.")
	}

	// 更新
	err = p.GetNavigationRepository().Modify(
		novelNavigation,
		form.NavigationName,
		form.SortNum,
		form.IsDisplayIndex,
	)

	if err != nil {
		// 判断是否是触发了msql的唯一性索引的错误，如果是则提示分类存在，其他错误走下面的异常抛出
		if p.isUniqueConstraintError(err) {
			return response.Failure, errors.New("Navigation existed.")
		}
		return response.Failure, errors.New("Navigation modify failed.")
	}
	return response.Ok, nil
}

// 导航栏目删除
func (p *NovelNavigationService) Delete(navigationId int) (code int, err error) {
	novelNavigation, err := p.GetNavigationRepository().Get(
		navigationId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Navigation not existed.")
	}
	// 删除
	err = p.GetNavigationRepository().Delete(novelNavigation)
	if err != nil {
		return response.Failure, errors.New("Navigation delete failed.")
	}
	return response.Ok, nil
}

// 导航栏目后台列表
func (p *NovelNavigationService) GetList(form request.NovelNavigationList) (result navigationDTO.NovelNavigationAdminListDTO, code int, err error) {
	list, total, err := p.GetNavigationRepository().GetList(
		form.Page,
		form.PageSize,
		nil,
	)
	if err != nil {
		return navigationDTO.NovelNavigationAdminListDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var data []*navigationDTO.NovelNavigationCopy
	for _, novelNavigation := range list {
		data = append(data, &navigationDTO.NovelNavigationCopy{*novelNavigation})
	}

	result = navigationDTO.NovelNavigationAdminListDTO{
		List:  data,
		Total: total,
	}
	return result, response.Ok, nil
}

// 首页的导航栏目数据
func (p *NovelNavigationService) GetIndexList() (result navigationDTO.NovelNavigationIndexListDTO, code int, err error) {
	list, err := p.GetNavigationRepository().GetAll(
		model.COMMON_MINUS,
		nil,
	)
	if err != nil {
		return navigationDTO.NovelNavigationIndexListDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var data []*navigationDTO.NovelNavigationCopy
	for _, novelNavigation := range list {
		data = append(data, &navigationDTO.NovelNavigationCopy{*novelNavigation})
	}

	result = navigationDTO.NovelNavigationIndexListDTO{
		List: data,
	}
	return result, response.Ok, nil
}

// 导航栏目详情
func (p *NovelNavigationService) Get(form request.NovelNavigationDetail) (data *navigationDTO.NovelNavigationCopy, code int, err error) {
	novelNavigation, err := p.GetNavigationRepository().Get(
		form.NovelNavigationId,
		"",
		[]string{"NovelCategories"},
	)
	if err != nil {
		return &navigationDTO.NovelNavigationCopy{}, response.Failure, errors.New("Navigation not existed.")
	}
	return &navigationDTO.NovelNavigationCopy{*novelNavigation}, response.Ok, nil
}
