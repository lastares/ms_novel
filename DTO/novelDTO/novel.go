package novelDTO

import (
	"encoding/json"
	"ms_novel/model"
)

// 后台小说列表DTO
type NovelAdminListDTO struct {
	List  []*NovelCopy `json:"list"`  // 列表数据
	Total int          `json:"total"` // 总数
}

// 新的自定义输出结构
type NovelCopy struct {
	model.Novel
}

// `json:"list"`该标签中的json tag 对应的拓展包是encoding/json包，这里可以重写MarshalJSON方法，实现我们想要的自定义输出结构
func (p *NovelCopy) MarshalJSON() ([]byte, error) {
	novelDTO := NovelDTO{
		Id:               p.Id,
		NovelTitle:       p.NovelTitle,
		NovelAuthor:      p.NovelAuthor,
		MainRole:         p.MainRole,
		NovelStatus:      p.NovelStatus,
		NovelStatusTitle: p.GetNovelStatusTitle(p.NovelStatus),
		IsNew:            p.IsNew,
		IsNewTitle:       p.GetIsValueTitle(p.IsNew),
		IsHot:            p.IsHot,
		IsHotTitle:       p.GetIsValueTitle(p.IsHot),
		IsRecommend:      p.IsRecommend,
		IsRecommendTitle: p.GetIsValueTitle(p.IsRecommend),
		IsDisplayIndex:   p.IsDisplayIndex,
	}
	if p.NovelNavigation != nil {
		novelDTO.NovelNavigation = &NavigationDTO{
			Id:             p.NovelNavigationId,
			NavigationName: p.NovelNavigation.NavigationName,
		}
	}

	if p.NovelCategory != nil {
		novelDTO.NovelCategory = &CategoryDTO{
			Id:           p.NovelCategoryId,
			CategoryName: p.NovelCategory.CategoryName,
		}
	}

	if p.NovelThemes != nil {
		for _, value := range p.NovelThemes {
			novelDTO.NovelThemes = append(novelDTO.NovelThemes, &ThemeDTO{
				Id:        value.Id,
				ThemeName: value.ThemeName,
			})
		}
	}
	return json.Marshal(novelDTO)
}

// 小说DTO
type NovelDTO struct {
	Id               int            `json:"id"`
	NovelTitle       string         `json:"novelTitle"`
	NovelAuthor      string         `json:"novelAuthor"`
	MainRole         string         `json:"mainRole"`
	NovelStatus      int            `json:"novelStatus"`
	NovelStatusTitle string         `json:"novelStatusTitle"`
	IsNew            int            `json:"isNew"`
	IsNewTitle       string         `json:"isNewTitle"`
	IsHot            int            `json:"isHot"`
	IsHotTitle       string         `json:"isHotTitle"`
	IsRecommend      int            `json:"isRecommend"`
	IsRecommendTitle string         `json:"isRecommendTitle"`
	IsDisplayIndex   int            `json:"isDisplayIndex"`
	NovelNavigation  *NavigationDTO `json:"novelNavigation"`
	NovelCategory    *CategoryDTO   `json:"novelCategory"`
	NovelThemes      []*ThemeDTO    `json:"novelThemes"`
}

// 分类DTO
type CategoryDTO struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

// 导航栏目
type NavigationDTO struct {
	Id             int    `json:"id"`
	NavigationName string `json:"navigationName"`
}

type ThemeDTO struct {
	Id        int    `json:"id"`
	ThemeName string `json:"themeName"`
}
