package request

// 获取右边小说数据
type IndexRightNovel struct {
	ManNavigationId   int `validate:"omitempty,min=1"` // 男生频道导航栏目ID
	WomenNavigationId int `validate:"omitempty,min=1"` // 女生频道导航栏目ID
}
