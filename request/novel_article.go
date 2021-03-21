package request

// 小说文章增加
type NovelArticleCreate struct {
	ArticleTitle   string `validate:"required,gte=1"`  // 文章标题
	ArticleAuthor  string `validate:"required,gte=1"`  // 文章作者
	ArticleContent string `validate:"required,gte=1"`  // 文章内容
	NovelId        int    `validate:"omitempty,gte=0"` // 小说ID
	SortNum        int    `validate:"omitempty,gte=0"` // 排序值
}

// 小说文章修改
type NovelArticleModify struct {
	NovelArticleId int    `validate:"required,gte=1"`  // 文章ID
	ArticleTitle   string `validate:"required,gte=1"`  // 文章标题
	ArticleAuthor  string `validate:"required,gte=1"`  // 文章作者
	ArticleContent string `validate:"required,gte=1"`  // 文章内容
	NovelId        int    `validate:"omitempty,gte=0"` // 小说ID
	SortNum        int    `validate:"omitempty,gte=0"` // 排序值
}

// 小说文章删除
type NovelArticleDelete struct {
	NovelArticleId int `validate:"required,gt=0"` // 文章ID
}

// 小说文章列表
type NovelArticleList struct {
	ArticleTitle string `validate:"omitempty,gt=0"` // 文章标题
	Page         int    `validate:"omitempty,gt=0"` // 当前页码
	PageSize     int    `validate:"omitempty,gt=0"` // 每页显示的条数
}

// 小说文章详情
type NovelArticleDetail struct {
	NovelArticleId int `validate:"required,gt=0"` // 文章ID
}
