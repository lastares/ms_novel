package request

// 小说增加
type NovelChapterCreate struct {
	ChapterTitle string `validate:"required,gte=1"` // 章节标题
	ChapterNo    int    `validate:"required,gte=1"` // 章节编码
	NovelId      int    `validate:"required,gte=1"` // 小说ID
}

// 小说章节修改
type NovelChapterModify struct {
	NovelChapterId int    `validate:"required,gte=1"` // 章节ID
	ChapterTitle   string `validate:"required,gte=1"` // 章节标题
	ChapterNo      int    `validate:"required,gte=1"` // 章节编码
	NovelId        int    `validate:"required,gte=1"` // 小说ID
}

// 主题删除
type NovelChapterDelete struct {
	NovelChapterId int `validate:"required,gt=0"` // 章节ID
}

// 章节列表
type NovelChapterList struct {
	ChapterTitle string `validate:"omitempty,gt=0"` // 章节标题
	Page         int    `validate:"omitempty,gt=0"` // 当前页码
	PageSize     int    `validate:"omitempty,gt=0"` // 每页显示的条数
}

//
// 章节详情
type NovelChapterDetail struct {
	NovelChapterId int `validate:"required,gt=0"` // 章节ID
}
