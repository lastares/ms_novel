package service

import (
	"errors"
	"ms_novel/DTO/chapterDTO"
	"ms_novel/request"
	"ms_novel/response"
)

// 定义章节相关接口，便于查看所有方法
type NovelChapterServiceInterface interface {
	// 创建
	Create(form request.NovelChapterCreate) (code int, err error)

	// 修改
	Modify(form request.NovelChapterModify) (code int, err error)

	// 删除
	Delete(form request.NovelChapterDelete) (code int, err error)

	// 详情
	GetDetail(form request.NovelChapterDetail) (data *chapterDTO.NovelChapterCopy, code int, err error)

	// 列表
	GetList(form request.NovelChapterList) (data chapterDTO.NovelChapterAdminListDTO, code int, err error)
}

type NovelChapterService struct {
	CommonService
}

// 专题创建
func (p *NovelChapterService) Create(
	form request.NovelChapterCreate,
) (code int, err error) {
	// 判断小说是否存在
	novel, err := p.GetNovelRepository().Get(
		form.NovelId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Novel not existed.")
	}

	// 添加
	err = p.GetChapterRepository().Create(
		form.ChapterTitle,
		form.ChapterNo,
		novel,
	)
	if err != nil {
		return response.Failure, errors.New("Novel chapter create failed.")
	}
	return response.Ok, nil
}

// 章节修改
func (p *NovelChapterService) Modify(form request.NovelChapterModify) (code int, err error) {
	// 判断章节是否存在
	novelChapter, err := p.GetChapterRepository().Get(form.NovelChapterId)
	if err != nil {
		return response.Failure, errors.New("Novel chapter not existed.")
	}

	// 判断小说是否存在
	novel, err := p.GetNovelRepository().Get(
		form.NovelId,
		"",
		nil,
	)
	if err != nil {
		return response.Failure, errors.New("Novel not existed.")
	}

	// 更新
	err = p.GetChapterRepository().Modify(
		novelChapter,
		form.ChapterTitle,
		form.ChapterNo,
		novel,
	)
	if err != nil {
		// 判断是否触发了msql的唯一性索引的错误
		if p.isUniqueConstraintError(err) {
			return response.Failure, errors.New("Novel chapter existed.")
		}
		return response.Failure, errors.New("Novel chapter modify failed.")
	}

	return response.Ok, nil
}

// 章节删除
func (p *NovelChapterService) Delete(form request.NovelChapterDelete) (code int, err error) {
	novelChapter, err := p.GetChapterRepository().Get(form.NovelChapterId)
	if err != nil {
		return response.Failure, errors.New("Novel chapter not existed.")
	}
	// 删除
	err = p.GetChapterRepository().Delete(novelChapter)
	if err != nil {
		return response.Failure, errors.New("Novel chapter delete failed.")
	}
	return response.Ok, nil
}

// 章节列表
func (p *NovelChapterService) GetList(form request.NovelChapterList) (data chapterDTO.NovelChapterAdminListDTO, code int, err error) {
	// 获取源数据
	list, total, err := p.GetChapterRepository().GetList(
		form.ChapterTitle,
		form.Page,
		form.PageSize,
	)

	if err != nil {
		return chapterDTO.NovelChapterAdminListDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var chapterDTOs []*chapterDTO.NovelChapterCopy
	for _, novelChapter := range list {
		chapterDTOs = append(chapterDTOs, &chapterDTO.NovelChapterCopy{*novelChapter})
	}

	// 返回DTO对象
	data = chapterDTO.NovelChapterAdminListDTO{
		List:  chapterDTOs,
		Total: total,
	}
	return data, response.Ok, nil
}

// 章节详情
func (p *NovelChapterService) GetDetail(form request.NovelChapterDetail) (data *chapterDTO.NovelChapterCopy, code int, err error) {
	novelChapter, err := p.GetChapterRepository().Get(form.NovelChapterId)
	if err != nil {
		return &chapterDTO.NovelChapterCopy{}, response.Failure, errors.New("Novel chapter not existed.")
	}
	return &chapterDTO.NovelChapterCopy{*novelChapter}, response.Ok, nil
}
