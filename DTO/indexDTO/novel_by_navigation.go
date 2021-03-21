package indexDTO

import (
	"encoding/json"
	"ms_novel/model"
	"os"
)

var INDEX_NAVIGATION_NOVELS_LIMIT = 17

// 总输出DTO
type NovelByNavDataDTO struct {
	List []*NovelByNavCopy
}

type NovelNavigationDTO struct {
	Id             int              `json:"id"`
	NavigationName string           `json:"navigationName"`
	Novels         []*NovelByNavDTO `json:"novels"`
}

type NovelByNavCopy struct {
	model.NovelNavigation
}

type NovelByNavDTO struct {
	Id               int    `json:"id"`
	NovelCover       string `json:"novelCover"`
	NovelAuthor      string `json:"novelAuthor"`
	NovelTitle       string `json:"novelTitle"`
	NovelDescription string `json:"novelDescription"`
}

func (p *NovelByNavCopy) MarshalJSON() ([]byte, error) {
	var novels []*NovelByNavDTO
	for _, value := range p.Novels {
		novelAuthor := value.NovelAuthor
		if novelAuthor == "" {
			novelAuthor = "佚名"
		}
		novel := NovelByNavDTO{
			Id:               value.Id,
			NovelCover:       os.Getenv("FS_HOST") + value.NovelCover,
			NovelAuthor:      novelAuthor,
			NovelTitle:       value.NovelTitle,
			NovelDescription: value.NovelDescription,
		}
		novels = append(novels, &novel)
		// 首页每个导航栏目下最多展示17个小说
		if len(novels) == INDEX_NAVIGATION_NOVELS_LIMIT {
			break
		}
	}
	novelNavigationDTO := NovelNavigationDTO{
		Id:             p.Id,
		NavigationName: p.NavigationName,
		Novels:         novels,
	}
	return json.Marshal(novelNavigationDTO)
}
