package service

import (
	"errors"
	"ms_novel/DTO/articleDTO"
	"ms_novel/request"
	"ms_novel/response"
)

// 定义文章相关接口，便于查看所有方法
type NovelArticleServiceInterface interface {
	// 创建
	Create(form request.NovelArticleCreate) (code int, err error)

	// 修改
	Modify(form request.NovelArticleModify) (code int, err error)

	// 删除
	Delete(form request.NovelArticleDelete) (code int, err error)

	// 详情
	GetDetail(form request.NovelArticleDetail) (data *articleDTO.NovelArticleCopy, code int, err error)

	// 列表
	GetList(form request.NovelArticleList) (data articleDTO.NovelArticleAdminListDTO, code int, err error)
}

type NovelArticleService struct {
	CommonService
}

// 文章创建
func (p *NovelArticleService) Create(form request.NovelArticleCreate) (code int, err error) {
	novel, err := p.GetNovelRepository().Get(
		form.NovelId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Novel not exist.")
	}

	// 添加
	err = p.GetArticleRepository().Create(
		form.ArticleTitle,
		form.ArticleAuthor,
		form.ArticleContent,
		form.SortNum,
		novel,
	)
	if err != nil {
		return response.Failure, errors.New("Article create failed.")
	}
	return response.Ok, nil
}

// 文章修改
func (p *NovelArticleService) Modify(form request.NovelArticleModify) (code int, err error) {
	novel, err := p.GetNovelRepository().Get(
		form.NovelId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Novel not exist.")
	}
	// 查询文章是否存在
	article, err := p.GetArticleRepository().Get(
		form.NovelArticleId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Article not existed.")
	}

	// 更新
	err = p.GetArticleRepository().Modify(
		article,
		form.ArticleTitle,
		form.ArticleAuthor,
		form.ArticleContent,
		form.SortNum,
		novel,
	)

	if err != nil {
		return response.Failure, errors.New("Article modify failed.")
	}
	return response.Ok, nil
}

// 文章删除
func (p *NovelArticleService) Delete(form request.NovelArticleDelete) (code int, err error) {
	article, err := p.GetArticleRepository().Get(
		form.NovelArticleId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Article not existed.")
	}
	// 删除
	err = p.GetArticleRepository().Delete(article)
	if err != nil {
		return response.Failure, errors.New("Article delete failed.")
	}
	return response.Ok, nil
}

// 文章列表
func (p *NovelArticleService) GetList(form request.NovelArticleList) (data articleDTO.NovelArticleAdminListDTO, code int, err error) {
	// 获取源数据
	list, total, err := p.GetArticleRepository().GetList(
		form.ArticleTitle,
		form.Page,
		form.PageSize,
		[]string{"Novel"},
	)

	if err != nil {
		return articleDTO.NovelArticleAdminListDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var articleDTOs []*articleDTO.NovelArticleCopy
	for _, article := range list {
		articleDTOs = append(articleDTOs, &articleDTO.NovelArticleCopy{*article})
	}

	// 返回DTO对象
	data = articleDTO.NovelArticleAdminListDTO{
		List:  articleDTOs,
		Total: total,
	}
	return data, response.Ok, nil
}

// 文章详情
func (p *NovelArticleService) GetDetail(form request.NovelArticleDetail) (data *articleDTO.NovelArticleCopy, code int, err error) {
	article, err := p.GetArticleRepository().Get(
		form.NovelArticleId,
		"",
		[]string{"Novel"},
	)
	if err != nil {
		return &articleDTO.NovelArticleCopy{}, response.Failure, errors.New("Article not existed.")
	}
	return &articleDTO.NovelArticleCopy{*article}, response.Ok, nil
}
