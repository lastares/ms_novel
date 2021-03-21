package service

import (
	"errors"
	"ms_novel/DTO/indexDTO"
	"ms_novel/model"
	"ms_novel/request"
	"ms_novel/response"
)

// 前台首页相关接口，便于查看所有方法
type IndexServiceInterface interface {
	// banner附近相关数据
	GetBannerData() (data indexDTO.NovelBannerDataDTO, code int, err error)

	// 推荐小说
	GetAllRecommendList() (data indexDTO.NovelDataDTO, code int, err error)

	// 友情链接
	GetFriendLinkList() (data indexDTO.IndexFriendLinkDataDTO, code int, err error)

	// 获取右边小说数据
	GetRightNovelList(form request.IndexRightNovel) (data indexDTO.IndexRightNovelDataDTO, code int, err error)

	// 获取首页导航栏目下的小说
	GetNavByNavigationId() (data indexDTO.NovelByNavDataDTO, code int, err error)
	GetNavByNavigationId2() (data []IndexNovelNavDTO, code int, err error)

	// 获取首页分类下的小说
	GetNovelByCategory() (data indexDTO.NovelByNavCategoryDataDTO, code int, err error)
}

//var IndexService = indexService{}

type IndexService struct {
	CommonService
}

// 首页第一块banner相关数据
func (p *IndexService) GetBannerData() (data indexDTO.NovelBannerDataDTO, code int, err error) {
	novels, err := p.GetNovelRepository().GetAll(
		[]string{"id", "novel_title", "novel_description", "novel_banner_cover"},
		model.COMMON_ONE,
		model.COMMON_MINUS,
		model.COMMON_ZERO,
		model.COMMON_ZERO,
		4,
		nil,
		nil,
		[]string{"id desc"},
	)
	if err != nil {
		return indexDTO.NovelBannerDataDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	novelCategories, err := p.GetCategoryRepository().GetAll(
		model.COMMON_ONE,
		nil,
	)
	if err != nil {
		return indexDTO.NovelBannerDataDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var novelDTOs []*indexDTO.NovelBannerCopy
	for _, novel := range novels {
		novelDTOs = append(novelDTOs, &indexDTO.NovelBannerCopy{*novel})
	}

	var novelCategoryDTOs []*indexDTO.NovelBannerCategoriesCopy
	for _, novelCategory := range novelCategories {
		novelCategoryDTOs = append(novelCategoryDTOs, &indexDTO.NovelBannerCategoriesCopy{*novelCategory})
	}

	// 返回DTO对象
	data = indexDTO.NovelBannerDataDTO{
		Novels:          novelDTOs,
		NovelCategories: novelCategoryDTOs,
	}
	return data, response.Ok, nil
}

// 获取首页推荐的小说
func (p *IndexService) GetAllRecommendList() (data indexDTO.NovelDataDTO, code int, err error) {
	novels, err := p.GetNovelRepository().GetAll(
		[]string{"id", "novel_title", "novel_author", "novel_description", "novel_cover", "novel_category_id"},
		model.COMMON_MINUS,
		model.COMMON_ONE,
		model.COMMON_ZERO,
		model.COMMON_ZERO,
		17,
		nil,
		[]string{"NovelCategory"},
		[]string{"id desc"},
	)
	if err != nil {
		return indexDTO.NovelDataDTO{}, response.Failure, errors.New("Failed to get data.")
	}
	// 实现自定义输出结构
	var novelDTOs []*indexDTO.NovelCopy
	for _, novel := range novels {
		novelDTOs = append(novelDTOs, &indexDTO.NovelCopy{*novel})
	}
	// 返回DTO对象
	data = indexDTO.NovelDataDTO{
		List: novelDTOs,
	}
	return data, response.Ok, nil
}

// 首页：获取男生频道、女生频道导航栏目下的小说
func (p *IndexService) GetNavByNavigationId() (data indexDTO.NovelByNavDataDTO, code int, err error) {
	// 1.查询要展示到首页的导航小说
	navigations, err := p.GetNavigationRepository().GetAll(
		model.COMMON_ONE,
		[]string{"Novels"},
	)

	if err != nil {
		return indexDTO.NovelByNavDataDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var novelNavDTOs []*indexDTO.NovelByNavCopy
	for _, navigation := range navigations {
		novelNavDTOs = append(novelNavDTOs, &indexDTO.NovelByNavCopy{*navigation})
	}
	// 返回DTO对象
	data = indexDTO.NovelByNavDataDTO{
		List: novelNavDTOs,
	}
	return data, response.Ok, nil
}

// 首页：获取男生频道、女生频道导航栏目下的小说接口第二种返回数据结构的方法，本质上是一种，只是写法不同
type IndexNovelNavDTO struct {
	Id             int             `json:"id"`
	NavigationName string          `json:"navigationName"`
	Novels         []IndexNovelDTO `json:"novels"`
}

type IndexNovelDTO struct {
	Id               int    `json:"id"`
	NovelCover       string `json:"novelCover"`
	NovelAuthor      string `json:"novelAuthor"`
	NovelTitle       string `json:"novelTitle"`
	NovelDescription string `json:"novelDescription"`
}

func (p *IndexService) GetNavByNavigationId2() (data []IndexNovelNavDTO, code int, err error) {
	//func (p *indexService) GetNavByNavigationId() (data indexDTO.NovelByNavDataDTO, code int, err error) {
	// 1.查询要展示到首页的导航小说
	navigations, err := p.GetNavigationRepository().GetAll(
		model.COMMON_ONE,
		[]string{"Novels"},
	)

	if err != nil {
		return []IndexNovelNavDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	for _, value := range navigations {
		var novels []IndexNovelDTO
		for _, novelItem := range value.Novels {
			novel := IndexNovelDTO{
				Id:               novelItem.Id,
				NovelCover:       novelItem.NovelCover,
				NovelAuthor:      novelItem.NovelAuthor,
				NovelTitle:       novelItem.NovelTitle,
				NovelDescription: novelItem.NovelDescription,
			}
			novels = append(novels, novel)
		}
		indexNovelNavDTO := IndexNovelNavDTO{
			Id:             value.Id,
			NavigationName: value.NavigationName,
			Novels:         novels,
		}
		data = append(data, indexNovelNavDTO)
	}
	return data, response.Ok, nil
}

// 首页：获取男生频道、女生频道导航栏目下的小说
func (p *IndexService) GetNovelByCategory() (data indexDTO.NovelByNavCategoryDataDTO, code int, err error) {
	// 1.查询要展示到首页的导航小说
	categories, err := p.GetCategoryRepository().GetAll(
		model.COMMON_ONE,
		[]string{"Novels"},
	)

	if err != nil {
		return indexDTO.NovelByNavCategoryDataDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var categoryNovelDTOs []*indexDTO.IndexCategoryNovelCopy
	for _, value := range categories {
		categoryNovelDTOs = append(categoryNovelDTOs, &indexDTO.IndexCategoryNovelCopy{*value})
	}
	// 返回DTO对象
	data = indexDTO.NovelByNavCategoryDataDTO{
		List: categoryNovelDTOs,
	}
	return data, response.Ok, nil
}

// 首页：获取友情链接
func (p *IndexService) GetFriendLinkList() (data indexDTO.IndexFriendLinkDataDTO, code int, err error) {
	// 1.查询要展示到首页的导航小说
	list, err := p.GetFriendLinkRepository().GetAll()

	if err != nil {
		return indexDTO.IndexFriendLinkDataDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var friendLinkDTOs []*indexDTO.IndexFriendLinkCopy
	for _, value := range list {
		friendLinkDTOs = append(friendLinkDTOs, &indexDTO.IndexFriendLinkCopy{*value})
	}
	// 返回DTO对象
	data = indexDTO.IndexFriendLinkDataDTO{
		List: friendLinkDTOs,
	}
	return data, response.Ok, nil
}

// 首页：获取友情链接
func (p *IndexService) GetRightNovelList(form request.IndexRightNovel) (data indexDTO.IndexRightNovelDataDTO, code int, err error) {
	// 精品小说
	qualityNovelDTOs, _, _ := p.getRightNovel(model.COMMON_ZERO)
	// 男频精品小说
	manRankNovelDTOS, _, _ := p.getRightNovel(form.ManNavigationId)
	// 女频精品小说
	womenRankNovelDTOS, _, _ := p.getRightNovel(form.WomenNavigationId)
	// 返回DTO对象
	data = indexDTO.IndexRightNovelDataDTO{
		QualityNovels:   qualityNovelDTOs,
		ManRankNovels:   manRankNovelDTOS,
		WomenRankNovels: womenRankNovelDTOS,
	}
	return data, response.Ok, nil
}

// 获取右边小说数据
func (p *IndexService) getRightNovel(
	navigationId int,
) (data []*indexDTO.IndexRightNovelCopy, code int, err error) {
	list, err := p.GetNovelRepository().GetAll(
		[]string{"id", "novel_title", "novel_category_id"},
		model.COMMON_MINUS,
		model.COMMON_MINUS,
		model.COMMON_ZERO,
		navigationId,
		indexDTO.QUALITY_GOODS_LIMIT,
		nil,
		[]string{"NovelCategory"},
		[]string{"view_num desc"},
	)
	if err != nil {
		return []*indexDTO.IndexRightNovelCopy{}, response.Failure, errors.New("Failed to get data.")
	}

	var qualityGoodsDTOs []*indexDTO.IndexRightNovelCopy
	for _, value := range list {
		qualityGoodsDTOs = append(qualityGoodsDTOs, &indexDTO.IndexRightNovelCopy{*value})
	}
	return qualityGoodsDTOs, response.Ok, nil
}
