package chapterDTO

import (
	"encoding/json"
	"ms_novel/model"
)

// 后台章节列表DTO
type NovelChapterAdminListDTO struct {
	List  []*NovelChapterCopy `json:"list"`
	Total int                 `json:"total"`
}

// 新的自定义输出结构
type NovelChapterCopy struct {
	model.NovelChapter
}

// `json:"list"`该标签中的json tag 对应的拓展包是encoding/json包，这里可以重写MarshalJSON方法，实现我们想要的自定义输出结构
func (p *NovelChapterCopy) MarshalJSON() ([]byte, error) {
	novelThemeDTO := NovelChapterDTO{
		Id:           p.Id,
		ChapterTitle: p.ChapterTitle,
		ChapterNo:    p.ChapterNo,
		Novel: &NovelDTO{
			Id:         p.NovelId,
			NovelTitle: p.Novel.NovelTitle,
		},
	}
	return json.Marshal(novelThemeDTO)
}

// 小说章节DTO
type NovelChapterDTO struct {
	Id           int    `json:"id"`
	ChapterTitle string `json:"chapterTitle"`
	ChapterNo    int    `json:"chapterNo"`
	Novel        *NovelDTO
}

// 小说DTO
type NovelDTO struct {
	Id         int    `json:"id"`
	NovelTitle string `json:"novelTitle"`
}
