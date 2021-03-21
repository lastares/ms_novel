package categoryDTO

import (
	"encoding/json"
	"ms_novel/model"
)

// 后台分类列表DTO
type NovelCategoryAdminListDTO struct {
	List  []*NovelCategoryCopy `json:"list"`
	Total int                  `json:"total"`
}

// 新的自定义输出结构
type NovelCategoryCopy struct {
	model.NovelCategory
}

// `json:"list"`该标签中的json tag 对应的拓展包是encoding/json包，这里可以重写MarshalJSON方法，实现我们想要的自定义输出结构
func (p *NovelCategoryCopy) MarshalJSON() ([]byte, error) {
	novelCategoryDTO := NovelCategoryDTO{
		Id:             p.Id,
		CategoryName:   p.CategoryName,
		CategoryIcon:   p.CategoryIcon,
		IsDisplayIndex: p.IsDisplayIndex,
		SortNum:        p.SortNum,
	}

	if p.NovelNavigationId != 0 {
		novelCategoryDTO.NovelNavigation = &NavigationDTO{
			Id:             p.NovelNavigationId,
			NavigationName: p.NovelNavigation.NavigationName,
		}
	}
	return json.Marshal(novelCategoryDTO)
}

// 分类DTO
type NovelCategoryDTO struct {
	Id              int            `json:"id"`
	CategoryName    string         `json:"categoryName"`
	CategoryIcon    string         `json:"categoryIcon"`
	IsDisplayIndex  int            `json:"isDisplayIndex"`
	SortNum         int            `json:"sortNum"`
	NovelNavigation *NavigationDTO `json:"novelNavigation,omitempty"`
}

// 导航栏目DTO
type NavigationDTO struct {
	Id             int    `json:"id"`
	NavigationName string `json:"navigationName"`
}
