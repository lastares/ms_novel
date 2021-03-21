package indexDTO

import (
	"encoding/json"
	"ms_novel/model"
	"os"
)

// 总输出DTO
type NovelDataDTO struct {
	List []*NovelCopy `json:"list"`
}

// 推荐小说DTO
type NovelCopy struct {
	model.Novel
}

type NovelDTO struct {
	Id               int          `json:"id"`
	NovelTitle       string       `json:"novelTitle"`
	NovelAuthor      string       `json:"novelAuthor"`
	NovelDescription string       `json:"novelDescription"`
	NovelCover       string       `json:"novelCover"`
	Category         *CategoryDTO `json:"category"`
}

type CategoryDTO struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func (p *NovelCopy) MarshalJSON() ([]byte, error) {
	novelAuthor := p.NovelAuthor
	if novelAuthor == "" {
		novelAuthor = "佚名"
	}
	novelDTO := &NovelDTO{
		Id:               p.Id,
		NovelTitle:       p.NovelTitle,
		NovelDescription: p.NovelDescription,
		NovelAuthor:      novelAuthor,
		NovelCover:       os.Getenv("FS_HOST") + p.NovelCover,
	}

	if p.NovelCategory != nil {
		novelDTO.Category = &CategoryDTO{
			Id:           p.NovelCategoryId,
			CategoryName: p.NovelCategory.CategoryName,
		}
	}
	return json.Marshal(novelDTO)
}
