package navigationDTO

import (
	"encoding/json"
	"ms_novel/model"
)

// 后台导航栏目列表DTO
type NovelNavigationAdminListDTO struct {
	List  []*NovelNavigationCopy `json:"list"`
	Total int                    `json:"total"`
}

// 前台导航栏目列表DTO
type NovelNavigationIndexListDTO struct {
	List []*NovelNavigationCopy `json:"list"`
}

// 新的自定义输出结构
type NovelNavigationCopy struct {
	model.NovelNavigation
}

// `json:"list"`该标签中的json tag 对应的拓展包是encoding/json包，这里可以重写MarshalJSON方法，实现我们想要的自定义输出结构
func (p *NovelNavigationCopy) MarshalJSON() ([]byte, error) {
	novelNavigationDTO := NovelNavigationDTO{
		Id:             p.Id,
		NavigationName: p.NavigationName,
	}

	if p.NovelCategories != nil {
		for _, category := range p.NovelCategories {
			novelNavigationDTO.NovelCategories = append(novelNavigationDTO.NovelCategories, &CategoryDTO{
				Id:           category.Id,
				CategoryName: category.CategoryName,
			})
		}
	}
	return json.Marshal(novelNavigationDTO)
}

// 导航栏目DTO
type NovelNavigationDTO struct {
	Id              int            `json:"id"`
	NavigationName  string         `json:"navigationName"`
	NovelCategories []*CategoryDTO `json:"novelCategories"`
}

// 分类DTO
type CategoryDTO struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}
