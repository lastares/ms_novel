package indexDTO

import (
	"encoding/json"
	"ms_novel/model"
	"os"
)

var INDEX_CATEGORY_NOVELS_LIMIT = 10

// 总输出DTO
type NovelByNavCategoryDataDTO struct {
	List []*IndexCategoryNovelCopy
}

type IndexCategoryDTO struct {
	Id           int                      `json:"id"`
	CategoryName string                   `json:"CategoryName"`
	Novels       []*IndexCategoryNovelDTO `json:"novels"`
}

type IndexCategoryNovelCopy struct {
	model.NovelCategory
}

type IndexCategoryNovelDTO struct {
	Id               int    `json:"id"`
	NovelCover       string `json:"novelCover"`
	NovelAuthor      string `json:"novelAuthor"`
	NovelTitle       string `json:"novelTitle"`
	NovelDescription string `json:"novelDescription"`
}

func (p *IndexCategoryNovelCopy) MarshalJSON() ([]byte, error) {
	var novels []*IndexCategoryNovelDTO
	for _, value := range p.Novels {
		novelAuthor := value.NovelAuthor
		if novelAuthor == "" {
			novelAuthor = "佚名"
		}
		novel := IndexCategoryNovelDTO{
			Id:               value.Id,
			NovelCover:       os.Getenv("FS_HOST") + value.NovelCover,
			NovelAuthor:      novelAuthor,
			NovelTitle:       value.NovelTitle,
			NovelDescription: value.NovelDescription,
		}
		novels = append(novels, &novel)
		// 首页每个分类下最多展示10个小说
		if len(novels) == INDEX_CATEGORY_NOVELS_LIMIT {
			break
		}
	}
	indexCategoryDTO := IndexCategoryDTO{
		Id:           p.Id,
		CategoryName: p.CategoryName,
		Novels:       novels,
	}
	return json.Marshal(indexCategoryDTO)
}
