package service

import (
	"errors"
	"ms_novel/DTO/friendLinkDTO"
	"ms_novel/request"
	"ms_novel/response"
)

// 定义友情链接的所有接口，便于查看所有方法
type NovelFriendLinkServiceInterface interface {
	// 创建
	Create(form request.NovelFriendLinkCreate) (code int, err error)

	// 修改
	Modify(form request.NovelFriendLinkModify) (code int, err error)

	// 删除
	Delete(form request.NovelFriendLinkDelete) (code int, err error)

	// 详情
	GetDetail(form request.NovelFriendLinkDetail) (data *friendLinkDTO.FriendLinkCopy, code int, err error)

	// 列表
	GetList(form request.NovelFriendLinkList) (data friendLinkDTO.FriendLinkAdminListDTO, code int, err error)
}

type NovelFriendLinkService struct {
	CommonService
}

// 友情链接创建
func (p *NovelFriendLinkService) Create(form request.NovelFriendLinkCreate) (code int, err error) {
	// 判断专题是否存在
	friendLink, _ := p.GetFriendLinkRepository().Get(
		0,
		form.LinkTitle,
	)
	if friendLink.Id > 0 {
		return response.Failure, errors.New("Friend link existed.")
	}

	// 添加
	err = p.GetFriendLinkRepository().Create(
		form.LinkTitle,
		form.Link,
		form.SortNum,
	)
	if err != nil {
		return response.Failure, errors.New("Friend link failed.")
	}
	return response.Ok, nil
}

// 友情链接修改
func (p *NovelFriendLinkService) Modify(form request.NovelFriendLinkModify) (code int, err error) {
	// 查询友情链接
	friendLink, err := p.GetFriendLinkRepository().Get(
		form.FriendLinkId,
		"",
	)
	if err != nil {
		return response.Failure, errors.New("Friend link not existed.")
	}

	// 更新
	err = p.GetFriendLinkRepository().Modify(
		friendLink,
		form.LinkTitle,
		form.Link,
		form.SortNum,
	)

	if err != nil {
		// 判断是否是触发了msql的唯一性索引的错误，如果是则提示友情链接存在，其他错误走下面的异常抛出
		if p.isUniqueConstraintError(err) {
			return response.Failure, errors.New("Friend link existed.")
		}
		return response.Failure, errors.New("Friend link modify failed.")
	}
	return response.Ok, nil
}

// 友情链接删除
func (p *NovelFriendLinkService) Delete(form request.NovelFriendLinkDelete) (code int, err error) {
	friendLink, err := p.GetFriendLinkRepository().Get(
		form.FriendLinkId,
		"",
	)
	if err != nil {
		return response.Failure, errors.New("Friend link not existed.")
	}
	// 删除
	err = p.GetFriendLinkRepository().Delete(friendLink)
	if err != nil {
		return response.Failure, errors.New("Friend link delete failed.")
	}
	return response.Ok, nil
}

// 友情链接列表
func (p *NovelFriendLinkService) GetList(form request.NovelFriendLinkList) (data friendLinkDTO.FriendLinkAdminListDTO, code int, err error) {
	// 获取源数据
	list, total, err := p.GetFriendLinkRepository().GetList(
		form.LinkTitle,
		form.Page,
		form.PageSize,
	)

	if err != nil {
		return friendLinkDTO.FriendLinkAdminListDTO{}, response.Failure, errors.New("Failed to get data.")
	}

	// 实现自定义输出结构
	var friendLinkDTOs []*friendLinkDTO.FriendLinkCopy
	for _, friendLink := range list {
		friendLinkDTOs = append(friendLinkDTOs, &friendLinkDTO.FriendLinkCopy{*friendLink})
	}

	// 返回DTO对象
	data = friendLinkDTO.FriendLinkAdminListDTO{
		List:  friendLinkDTOs,
		Total: total,
	}
	return data, response.Ok, nil
}

// 友情链接详情
func (p *NovelFriendLinkService) GetDetail(form request.NovelFriendLinkDetail) (data *friendLinkDTO.FriendLinkCopy, code int, err error) {
	friendLink, err := p.GetFriendLinkRepository().Get(
		form.FriendLinkId,
		"",
	)
	if err != nil {
		return &friendLinkDTO.FriendLinkCopy{}, response.Failure, errors.New("Friend link not existed.")
	}
	return &friendLinkDTO.FriendLinkCopy{*friendLink}, response.Ok, nil
}
