package request

// 友情链接增加
type NovelFriendLinkCreate struct {
	LinkTitle string `validate:"required,gte=1"`  // 友情链接名称
	Link      string `validate:"required,gte=1"`  // 链接地址
	SortNum   int    `validate:"omitempty,gte=0"` // 排序值
}

// 友情链接修改
type NovelFriendLinkModify struct {
	FriendLinkId int    `validate:"required,gte=1"`  // 友情链接ID
	LinkTitle    string `validate:"required,gte=1"`  // 友情链接名称
	Link         string `validate:"required,gte=1"`  // 链接地址
	SortNum      int    `validate:"omitempty,gte=0"` // 排序值
}

// 友情链接删除
type NovelFriendLinkDelete struct {
	FriendLinkId int `validate:"required,gt=0"` // 友情链接ID
}

// 友情链接列表
type NovelFriendLinkList struct {
	LinkTitle string `validate:"omitempty,gt=0"` // 友情链接ID
	Page      int    `validate:"omitempty,gt=0"` // 当前页码
	PageSize  int    `validate:"omitempty,gt=0"` // 每页显示的条数
}

// 友情链接详情
type NovelFriendLinkDetail struct {
	FriendLinkId int `validate:"required,gt=0"` // 友情链接ID
}
