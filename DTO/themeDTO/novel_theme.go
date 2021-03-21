package themeDTO

import (
	"encoding/json"
	"ms_novel/model"
)

// 后台专题列表DTO
type NovelThemeAdminListDTO struct {
	List  []*NovelThemeCopy `json:"list"`
	Total int               `json:"total"`
}

// 新的自定义输出结构
type NovelThemeCopy struct {
	model.NovelTheme
}

// `json:"list"`该标签中的json tag 对应的拓展包是encoding/json包，这里可以重写MarshalJSON方法，实现我们想要的自定义输出结构
func (p *NovelThemeCopy) MarshalJSON() ([]byte, error) {
	novelThemeDTO := NovelThemeDTO{
		Id:               p.Id,
		ThemeName:        p.ThemeName,
		ThemeSubtitle:    p.ThemeSubtitle,
		ThemeDescription: p.ThemeDescription,
		ThemeCover:       p.ThemeCover,
		SortNum:          p.SortNum,
	}
	return json.Marshal(novelThemeDTO)
}

// 专题DTO
type NovelThemeDTO struct {
	Id               int    `json:"id"`
	ThemeName        string `json:"themeName"`
	ThemeSubtitle    string `json:"themeSubtitle"`
	ThemeDescription string `json:"themeDescription"`
	ThemeCover       string `json:"themeCover"`
	SortNum          int    `json:"sortNum"`
}
