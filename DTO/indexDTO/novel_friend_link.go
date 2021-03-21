package indexDTO

import (
	"encoding/json"
	"ms_novel/model"
)

// 总输出DTO
type IndexFriendLinkDataDTO struct {
	List []*IndexFriendLinkCopy
}

type IndexFriendLinkCopy struct {
	model.FriendLink
}

type IndexFriendLinkDTO struct {
	Id        int    `json:"id"`
	LinkTitle string `json:"linkTitle"`
	Link      string `json:"link"`
	SortNum   int    `json:"SortNum"`
}

func (p *IndexFriendLinkCopy) MarshalJSON() ([]byte, error) {
	indexFriendLinkDTO := IndexFriendLinkDTO{
		Id:        p.Id,
		LinkTitle: p.LinkTitle,
		Link:      p.Link,
		SortNum:   p.SortNum,
	}
	return json.Marshal(indexFriendLinkDTO)
}
