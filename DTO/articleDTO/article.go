package articleDTO

import (
	"encoding/json"
	"ms_novel/model"
)

// 后台文章列表DTO
type NovelArticleAdminListDTO struct {
	List  []*NovelArticleCopy `json:"list"`
	Total int                 `json:"total"`
}

// 新的自定义输出结构
type NovelArticleCopy struct {
	model.NovelArticle
}

// `json:"list"`该标签中的json tag 对应的拓展包是encoding/json包，这里可以重写MarshalJSON方法，实现我们想要的自定义输出结构
func (p *NovelArticleCopy) MarshalJSON() ([]byte, error) {
	novelArticleDTO := NovelArticleDTO{
		Id:             p.Id,
		ArticleTitle:   p.ArticleTitle,
		ArticleAuthor:  p.ArticleAuthor,
		ArticleContent: p.ArticleContent,
		SortNum:        p.SortNum,
	}

	if p.Novel != nil {
		novelArticleDTO.Novel = &NovelDTO{
			p.NovelId,
			p.Novel.NovelTitle,
		}
	}
	return json.Marshal(novelArticleDTO)
}

// 小说文章DTO
type NovelArticleDTO struct {
	Id             int       `json:"id"`
	ArticleTitle   string    `json:"articleTitle"`
	ArticleAuthor  string    `json:"articleAuthor"`
	ArticleContent string    `json:"articleContent"`
	SortNum        int       `json:"sortNum"`
	Novel          *NovelDTO `json:"novel"`
}

type NovelDTO struct {
	Id         int    `json:"id"`
	NovelTitle string `json:"novelTitle"`
}
