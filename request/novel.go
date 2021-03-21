package request

// 小说增加
type NovelCreate struct {
	NovelTitle        string `validate:"required,max=256" example:"盘龙"`                    // 小说标题
	NovelAuthor       string `validate:"required,max=128" example:"天蚕土豆"`                  // 小说作者
	NovelDescription  string `validate:"required,gte=1" example:"这是一段美好的描述"`               // 小说描述
	NovelCover        string `validate:"required,gte=1" example:"http://baidu.com/a.jpg"`  // 小说小封面
	NovelBannerCover  string `validate:"omitempty,gte=1" example:"http://baidu.com/a.jpg"` // 小说banner
	MainRole          string `validate:"required,gte=1" example:"张三丰"`                     // 小说角色
	ChapterPreview    string `validate:"required,gte=1" example:"xxxx"`                    // 章节预览
	NovelStatus       int    `validate:"omitempty,oneof=1 2" example:"1"`                  // 小说状态：1-连载中 2-已完结
	IsNew             int    `validate:"omitempty,oneof=0 1"`                              // 是否是最新：0-否 1-是
	IsHot             int    `validate:"omitempty,oneof=0 1"`                              // 是否热门：0-否 1-是
	IsRecommend       int    `validate:"omitempty,oneof=0 1"`                              // 是否推荐：0-否 1-是
	IsDisplayIndex    int    `validate:"omitempty,oneof=0 1"`                              // 是否首页展示：1-是 0-否
	SortNum           int    `validate:"omitempty,gte=0"`                                  // 排序值
	NovelNavigationId int    `validate:"required,gte=1" example:"1"`                       // 小说导航ID
	NovelCategoryId   int    `validate:"required,gte=1"`                                   // 小说分类ID
	NovelThemeIds     []int  `validate:"required"`                                         // 小说专题id数组
}

// 小说修改
type NovelModify struct {
	NovelId           int    `validate:"required,gt=0" example:"1"`                        // 小说ID
	NovelTitle        string `validate:"required,max=256" example:"盘龙"`                    // 小说标题
	NovelAuthor       string `validate:"required,max=128" example:"天蚕土豆"`                  // 小说作者
	NovelDescription  string `validate:"required,gte=1" example:"这是一段美好的描述"`               // 小说描述
	NovelCover        string `validate:"required,gte=1" example:"http://baidu.com/a.jpg"`  // 小说小封面
	NovelBannerCover  string `validate:"omitempty,gte=1" example:"http://baidu.com/a.jpg"` // 小说banner
	MainRole          string `validate:"required,gte=1" example:"张三丰"`                     // 小说角色
	ChapterPreview    string `validate:"required,gte=1"`                                   // 章节预览
	NovelStatus       int    `validate:"omitempty,oneof=1 2"`                              // 小说状态：1-连载中 2-已完结
	IsNew             int    `validate:"omitempty,oneof=0 1"`                              // 是否是最新：0-否 1-是
	IsHot             int    `validate:"omitempty,oneof=0 1"`                              // 是否热门：0-否 1-是
	IsRecommend       int    `validate:"omitempty,oneof=0 1"`                              // 是否推荐：0-否 1-是
	IsDisplayIndex    int    `validate:"omitempty,oneof=0 1"`                              // 是否首页展示：1-是 0-否
	SortNum           int    `validate:"omitempty,gte=0"`                                  // 排序值
	NovelNavigationId int    `validate:"required,gte=1"`                                   // 小说导航ID
	NovelCategoryId   int    `validate:"required,gte=1"`                                   // 小说分类ID
	NovelThemeIds     []int  `validate:"required"`                                         // 小说专题id数组
}

// 小说删除
type NovelDelete struct {
	NovelId int `validate:"required,gt=0"` // 小说ID
}

// 小说列表
type NovelList struct {
	NovelTitle  string `validate:"omitempty,gt=0"`      // 小说标题
	NovelStatus int    `validate:"omitempty,oneof=1 2"` // 小说状态
	Page        int    `validate:"omitempty,gt=0"`      // 当前页码
	PageSize    int    `validate:"omitempty,gt=0"`      // 每页显示的条数
}

// 小说详情
type NovelDetail struct {
	NovelId int `validate:"required,gt=0"` // 小说ID
}
