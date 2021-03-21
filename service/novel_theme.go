package service

import (
	"errors"
	"ms_novel/DTO/themeDTO"
	"ms_novel/request"
	"ms_novel/response"
)

// 定义小说专题接口，便于查看所有方法
type NovelThemeServiceInterface interface {
	// 创建
	Create(form request.NovelThemeCreate) (code int, err error)

	// 修改
	Modify(form request.NovelThemeModify) (code int, err error)

	// 删除
	Delete(form request.NovelThemeDelete) (code int, err error)

	// 详情
	GetDetail(form request.NovelThemeDetail) (data *themeDTO.NovelThemeCopy, code int, err error)

	// 列表
	GetList(form request.NovelThemeList) (data themeDTO.NovelThemeAdminListDTO, code int, err error)
}

type NovelThemeService struct {
	CommonService
}

// 专题创建
func (p *NovelThemeService) Create(
	form request.NovelThemeCreate,
) (code int, err error) {
	// 判断专题是否存在
	novelTheme, _ := p.GetThemeRepository().Get(
		0,
		form.ThemeName,
		"",
	)
	if novelTheme.Id > 0 {
		return response.Failure, errors.New("Theme existed.")
	}

	// 添加
	err = p.GetThemeRepository().Create(
		form.ThemeName,
		form.ThemeSubtitle,
		form.ThemeDescription,
		form.ThemeCover,
		form.SortNum,
	)
	if err != nil {
		return response.Failure, errors.New("Theme create failed.")
	}
	return response.Ok, nil
}

// 专题修改
func (p *NovelThemeService) Modify(
	form request.NovelThemeModify,
) (code int, err error) {
	// 查询专题
	novelTheme, err := p.GetThemeRepository().Get(
		form.NovelThemeId,
		"",
		"",
	)
	if err != nil {
		return response.Failure, errors.New("Theme not existed.")
	}

	// 更新
	err = p.GetThemeRepository().Modify(
		novelTheme,
		form.ThemeName,
		form.ThemeSubtitle,
		form.ThemeDescription,
		form.ThemeCover,
		form.SortNum,
	)

	if err != nil {
		// 判断是否触发了msql的唯一性索引的错误
		if p.isUniqueConstraintError(err) {
			return response.Failure, errors.New("Theme existed.")
		}
		return response.Failure, errors.New("Theme modify failed.")
	}
	return response.Ok, nil
}

// 专题删除
func (p *NovelThemeService) Delete(form request.NovelThemeDelete) (code int, err error) {
	novelTheme, err := p.GetThemeRepository().Get(
		form.NovelThemeId,
		"",
		"",
	)
	if err != nil {
		return response.Failure, errors.New("Theme not existed.")
	}
	// 删除
	err = p.GetThemeRepository().Delete(novelTheme)
	if err != nil {
		return response.Failure, errors.New("Theme delete failed.")
	}
	return response.Ok, nil
}

// 专题列表
func (p *NovelThemeService) GetList(form request.NovelThemeList) (data themeDTO.NovelThemeAdminListDTO, code int, err error) {
	// 获取源数据
	list, total, err := p.GetThemeRepository().GetList(
		form.ThemeName,
		form.ThemeSubtitle,
		form.Page,
		form.PageSize,
	)

	if err != nil {
		return themeDTO.NovelThemeAdminListDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var themeDTOs []*themeDTO.NovelThemeCopy
	for _, novelTheme := range list {
		themeDTOs = append(themeDTOs, &themeDTO.NovelThemeCopy{*novelTheme})
	}

	// 返回DTO对象
	data = themeDTO.NovelThemeAdminListDTO{
		List:  themeDTOs,
		Total: total,
	}
	return data, response.Ok, nil
}

// 专题详情
func (p *NovelThemeService) GetDetail(form request.NovelThemeDetail) (data *themeDTO.NovelThemeCopy, code int, err error) {
	novelTheme, err := p.GetThemeRepository().Get(
		form.NovelThemeId,
		"",
		"",
	)
	if err != nil {
		return &themeDTO.NovelThemeCopy{}, response.Failure, errors.New("Theme not existed.")
	}
	return &themeDTO.NovelThemeCopy{*novelTheme}, response.Ok, nil
}
