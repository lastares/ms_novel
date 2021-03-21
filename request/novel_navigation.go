package request

// 栏目导航增加
type NovelNavigationCreate struct {
	NavigationName string `validate:"required,gte=1"`      // 导航栏目名称
	SortNum        int    `validate:"omitempty,gte=0"`     // 排序值
	IsDisplayIndex int    `validate:"omitempty,oneof=0 1"` // 是否首页展示：0-否 1-是
}

// 栏目导航修改
type NovelNavigationModify struct {
	NavigationId   int    `validate:"required,gt=0"`       // 导航栏目名称
	NavigationName string `validate:"required,gte=1"`      // 导航栏目名称
	SortNum        int    `validate:"omitempty,gte=0"`     // 排序值
	IsDisplayIndex int    `validate:"omitempty,oneof=0 1"` // 是否首页展示：0-否 1-是
}

// 栏目导航删除
type NovelNavigationDelete struct {
	NavigationId int `validate:"required,gt=0"` // 导航栏目ID
}

// 栏目导航列表
type NovelNavigationList struct {
	Page     int `validate:"omitempty,gt=0"` // 当前页码
	PageSize int `validate:"omitempty,gt=0"` // 每页显示的条数
}

// 栏目导航详情
type NovelNavigationDetail struct {
	NovelNavigationId int `validate:"required,gt=0"` // 导航栏目名称
}
