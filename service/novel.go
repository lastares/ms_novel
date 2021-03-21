package service

import (
	"errors"
	"ms_novel/DTO/novelDTO"
	"ms_novel/request"
	"ms_novel/response"
)

// 定义小说相关接口，便于查看所有方法
type NovelServiceInterface interface {
	// 创建
	Create(form request.NovelCreate) (code int, err error)

	// 修改
	Modify(form request.NovelModify) (code int, err error)

	// 删除
	Delete(form request.NovelDelete) (code int, err error)

	// 详情
	GetDetail(form request.NovelDetail) (data *novelDTO.NovelCopy, code int, err error)
	GetDetail2(form request.NovelDetail) (data novelDTO.NovelDTO, code int, err error)

	// 列表
	GetList(form request.NovelList) (data novelDTO.NovelAdminListDTO, code int, err error)
}

type NovelService struct {
	CommonService
}

// 专题创建
func (p *NovelService) Create(
	form request.NovelCreate,
) (code int, err error) {
	// 判断导航栏目是否存在
	novelNavigation, err := p.GetNavigationRepository().Get(
		form.NovelNavigationId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Navigation not existed.")
	}

	// 判断分类是否存在
	novelCategory, err := p.GetCategoryRepository().Get(
		form.NovelCategoryId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Category not existed.")
	}

	// 判断专题是否存在
	themeList, _, err := p.GetThemeRepository().GetAllList(
		form.NovelThemeIds,
		[]string{"Novels"},
	)
	if err != nil {
		return response.Failure, errors.New("Failed to get data.")
	}

	// 添加
	err = p.GetNovelRepository().Create(
		form.NovelTitle,
		form.NovelAuthor,
		form.NovelDescription,
		form.NovelCover,
		form.NovelBannerCover,
		form.MainRole,
		form.ChapterPreview,
		form.NovelStatus,
		form.SortNum,
		novelNavigation,
		novelCategory,
		themeList,
	)
	if err != nil {
		return response.Failure, errors.New("Theme create failed.")
	}
	return response.Ok, nil
}

// 小说修改
func (p *NovelService) Modify(
	form request.NovelModify,
) (code int, err error) {
	// 判断小说是否存在
	novel, err := p.GetNovelRepository().Get(
		form.NovelId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Novel not existed.")
	}

	// 判断导航栏目是否存在
	novelNavigation, err := p.GetNavigationRepository().Get(
		form.NovelNavigationId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Navigation not existed.")
	}

	// 判断分类是否存在
	novelCategory, err := p.GetCategoryRepository().Get(
		form.NovelCategoryId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Category not existed.")
	}

	// 判断专题是否存在
	themeList, _, err := p.GetThemeRepository().GetAllList(
		form.NovelThemeIds,
		[]string{"Novels"},
	)
	if err != nil {
		return response.Failure, errors.New("Failed to get data.")
	}

	// 修改
	err = p.GetNovelRepository().Modify(
		novel,
		form.NovelStatus,
		form.SortNum,
		form.NovelTitle,
		form.NovelAuthor,
		form.NovelDescription,
		form.NovelCover,
		form.NovelBannerCover,
		form.MainRole,
		form.ChapterPreview,
		novelNavigation,
		novelCategory,
		themeList,
	)
	if err != nil {
		return response.Failure, errors.New("Novel modify failed.")
	}
	return response.Ok, nil
}

// 小说删除
func (p *NovelService) Delete(form request.NovelDelete) (code int, err error) {
	novel, err := p.GetNovelRepository().Get(
		form.NovelId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Novel not existed.")
	}
	// 删除
	err = p.GetNovelRepository().Delete(novel)
	if err != nil {
		return response.Failure, errors.New("Novel delete failed.")
	}
	return response.Ok, nil
}

// 小说列表
func (p *NovelService) GetList(form request.NovelList) (data novelDTO.NovelAdminListDTO, code int, err error) {
	// 获取源数据
	list, total, err := p.GetNovelRepository().GetList(
		form.NovelTitle,
		form.NovelStatus,
		form.Page,
		form.PageSize,
		[]string{"NovelNavigation", "NovelCategory", "NovelThemes"},
	)

	if err != nil {
		return novelDTO.NovelAdminListDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var novelDTOs []*novelDTO.NovelCopy
	for _, novel := range list {
		novelDTOs = append(novelDTOs, &novelDTO.NovelCopy{*novel})
	}

	// 返回DTO对象
	data = novelDTO.NovelAdminListDTO{
		List:  novelDTOs,
		Total: total,
	}
	return data, response.Ok, nil
}

// 专题详情
func (p *NovelService) GetDetail(form request.NovelDetail) (data *novelDTO.NovelCopy, code int, err error) {
	novel, err := p.GetNovelRepository().Get(
		form.NovelId,
		"",
		[]string{"NovelNavigation", "NovelCategory", "NovelThemes"},
	)
	if err != nil {
		return &novelDTO.NovelCopy{}, response.Failure, errors.New("Novel not existed.")
	}
	return &novelDTO.NovelCopy{*novel}, response.Ok, nil
}

func (p *NovelService) GetDetail2(form request.NovelDetail) (data novelDTO.NovelDTO, code int, err error) {
	novel, err := p.GetNovelRepository().Get(
		form.NovelId,
		"",
		[]string{"NovelNavigation", "NovelCategory", "NovelThemes"},
	)
	if err != nil {
		return novelDTO.NovelDTO{}, response.Failure, errors.New("Novel not existed.")
	}

	data = novelDTO.NovelDTO{
		Id:               novel.Id,
		NovelTitle:       novel.NovelTitle,
		NovelAuthor:      novel.NovelAuthor,
		MainRole:         novel.MainRole,
		NovelStatus:      novel.NovelStatus,
		NovelStatusTitle: novel.GetNovelStatusTitle(novel.NovelStatus),
		IsNew:            novel.IsNew,
		IsNewTitle:       novel.GetIsValueTitle(novel.IsNew),
		IsHot:            novel.IsHot,
		IsHotTitle:       novel.GetIsValueTitle(novel.IsHot),
		IsRecommend:      novel.IsRecommend,
		IsRecommendTitle: novel.GetIsValueTitle(novel.IsRecommend),
		IsDisplayIndex:   novel.IsDisplayIndex,
	}

	if novel.NovelNavigation != nil {
		data.NovelNavigation = &novelDTO.NavigationDTO{
			Id:             novel.NovelNavigationId,
			NavigationName: novel.NovelNavigation.NavigationName,
		}
	}

	if novel.NovelCategory != nil {
		data.NovelCategory = &novelDTO.CategoryDTO{
			Id:           novel.NovelCategoryId,
			CategoryName: novel.NovelCategory.CategoryName,
		}
	}

	if novel.NovelThemes != nil {
		for _, value := range novel.NovelThemes {
			data.NovelThemes = append(data.NovelThemes, &novelDTO.ThemeDTO{
				Id:        value.Id,
				ThemeName: value.ThemeName,
			})
		}
	}
	return data, response.Ok, nil
}
