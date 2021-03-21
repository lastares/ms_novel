package request

// 小说分类添加校验结构体
type NovelCategoryCreate struct {
	CategoryName      string `validate:"required,gte=1,lt=256"` // 分类名称
	CategoryIcon      string `validate:"required,gte=1,lt=256"` // 分类icon地址
	IsDisplayIndex    int    `validate:"omitempty,gte=0,lte=1"` // 是否首页展示：0-否 1-是
	SortNum           int    `validate:"omitempty,gte=0"`       // 排序值
	NovelNavigationId int    `validate:"required,gt=0"`         // 导航栏目ID
}

// 小说分类修改校验结构体
type NovelCategoryModify struct {
	NovelCategoryId   int    `validate:"required,gte=1"`        // 分类ID
	CategoryName      string `validate:"required,gte=1,lt=256"` // 分类名称
	CategoryIcon      string `validate:"required,gte=1,lt=256"` // 分类icon地址
	IsDisplayIndex    int    `validate:"omitempty,gte=0,lte=1"` // 是否首页展示：0-否 1-是
	SortNum           int    `validate:"omitempty,gte=0"`       // 排序值
	NovelNavigationId int    `validate:"required,gt=0"`         // 导航栏目ID
}

// 小说分类删除校验结构体
type NovelCategoryDelete struct {
	NovelCategoryId int `validate:"required,gt=0"` // 分类ID
}

// 栏目导航列表
type NovelCategoryList struct {
	CategoryName string `validate:"omitempty,gt=0"` // 分类名称
	Page         int    `validate:"omitempty,gt=0"` // 当前页码
	PageSize     int    `validate:"omitempty,gt=0"` // 每页显示的条数
}

// 分类详情
type NovelCategoryDetail struct {
	NovelCategoryId int `validate:"required,gte=1"` // 分类ID
}
