package request

// 小说专题增加
type NovelThemeCreate struct {
	ThemeName        string `validate:"required,gte=1"`         // 专题名称
	ThemeSubtitle    string `validate:"required,gte=1,lte=512"` // 专题子标题
	ThemeDescription string `validate:"required,gte=1,lte=512"` // 专题描述
	ThemeCover       string `validate:"required,gte=1,lte=256"` // 专题封面地址
	SortNum          int    `validate:"omitempty,gte=0"`        // 排序值
}

// 小说专题修改
type NovelThemeModify struct {
	NovelThemeId     int    `validate:"required,gte=1"`         // 专题ID
	ThemeName        string `validate:"required,gte=1"`         // 专题名称
	ThemeSubtitle    string `validate:"required,gte=1,lte=512"` // 专题子标题
	ThemeDescription string `validate:"required,gte=1,lte=512"` // 专题描述
	ThemeCover       string `validate:"required,gte=1,lte=256"` // 专题封面地址
	SortNum          int    `validate:"omitempty,gte=0"`        // 排序值
}

// 主题删除
type NovelThemeDelete struct {
	NovelThemeId int `validate:"required,gt=0"` // 专题ID
}

// 主题列表
type NovelThemeList struct {
	ThemeName     string `validate:"omitempty,gt=0"` // 专题名称
	ThemeSubtitle string `validate:"omitempty,gt=0"` // 专题子标题
	Page          int    `validate:"omitempty,gt=0"` // 当前页码
	PageSize      int    `validate:"omitempty,gt=0"` // 每页显示的条数
}

// 专题详情
type NovelThemeDetail struct {
	NovelThemeId int `validate:"required,gt=0"` // 专题ID
}
