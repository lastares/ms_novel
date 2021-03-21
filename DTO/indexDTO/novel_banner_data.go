package indexDTO

import (
	"encoding/json"
	"ms_novel/model"
	"os"
)

// 总输出DTO
type NovelBannerDataDTO struct {
	Novels          []*NovelBannerCopy           `json:"novels"`
	NovelCategories []*NovelBannerCategoriesCopy `json:"novelCategories"`
}

// 小说banner的DTO
type NovelBannerCopy struct {
	model.Novel
}

type NovelBannerDTO struct {
	Id               int    `json:"id"`
	NovelTitle       string `json:"novelTitle"`
	NovelDescription string `json:"novelDescription"`
	NovelBannerCover string `json:"novelBannerCover"`
}

func (p *NovelBannerCopy) MarshalJSON() ([]byte, error) {
	novelBannerDTO := &NovelBannerDTO{
		Id:               p.Id,
		NovelTitle:       p.NovelTitle,
		NovelDescription: p.NovelDescription,
		NovelBannerCover: os.Getenv("FS_HOST") + p.NovelBannerCover,
	}
	return json.Marshal(novelBannerDTO)
}

// 小说分类的DTO
type NovelBannerCategoriesCopy struct {
	model.NovelCategory
}

type NovelCategoryDTO struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
	CategoryIcon string `json:"categoryIcon"`
}

func (p *NovelBannerCategoriesCopy) MarshalJSON() ([]byte, error) {
	novelCategoryDTO := NovelCategoryDTO{
		Id:           p.Id,
		CategoryName: p.CategoryName,
		CategoryIcon: os.Getenv("FS_HOST") + p.CategoryIcon,
	}
	return json.Marshal(novelCategoryDTO)
}
