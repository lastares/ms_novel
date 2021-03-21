package friendLinkDTO

import (
	"encoding/json"
	"ms_novel/model"
)

// 后台友情链接列表DTO
type FriendLinkAdminListDTO struct {
	List  []*FriendLinkCopy `json:"list"`
	Total int               `json:"total"`
}

// 新的自定义输出结构
type FriendLinkCopy struct {
	model.FriendLink
}

// `json:"list"`该标签中的json tag 对应的拓展包是encoding/json包，这里可以重写MarshalJSON方法，实现我们想要的自定义输出结构
func (p *FriendLinkCopy) MarshalJSON() ([]byte, error) {
	friendLinkDTO := FriendLinkDTO{
		Id:        p.Id,
		LinkTitle: p.LinkTitle,
		Link:      p.Link,
		SortNum:   p.SortNum,
	}
	return json.Marshal(friendLinkDTO)
}

// 专题DTO
type FriendLinkDTO struct {
	Id        int    `json:"id"`
	LinkTitle string `json:"linkTitle"`
	Link      string `json:"link"`
	SortNum   int    `json:"sortNum"`
}
