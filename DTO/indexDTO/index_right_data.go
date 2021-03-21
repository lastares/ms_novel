package indexDTO

import (
	"encoding/json"
	"ms_novel/model"
)

var QUALITY_GOODS_LIMIT = 10

// 最新资讯
type IndexLatestArticleDataDTO struct {
	LatestArticles []*IndexLatestArticleCopy
}

type IndexLatestArticleCopy struct {
	model.NovelArticle
}

type IndexNovelArticleDTO struct {
	Id             int    `json:"id"`
	ArticleTitle   string `json:"articleTitle"`
	ArticleAuthor  string `json:"articleAuthor"`
	ArticleContent string `json:"articleContent"`
	SortNum        int    `json:"SortNum"`
}

func (p *IndexLatestArticleCopy) MarshalJSON() ([]byte, error) {
	indexNovelArticleDTO := IndexNovelArticleDTO{
		Id:             p.Id,
		ArticleTitle:   p.ArticleTitle,
		ArticleAuthor:  p.ArticleAuthor,
		ArticleContent: p.ArticleContent,
		SortNum:        p.SortNum,
	}
	return json.Marshal(indexNovelArticleDTO)
}

// 精品、男频排行榜、女频排行榜DTO
type IndexRightNovelDataDTO struct {
	QualityNovels   []*IndexRightNovelCopy `json:"qualityNovel"`
	ManRankNovels   []*IndexRightNovelCopy `json:"manRankNovels"`
	WomenRankNovels []*IndexRightNovelCopy `json:"womenRankNovels"`
}

type IndexRightNovelCopy struct {
	model.Novel
}

type IndexRightNovelDTO struct {
	Id           int    `json:"id"`
	NovelTitle   string `json:"novelTitle"`
	CategoryName string `json:"categoryName"`
}

func (p *IndexRightNovelCopy) MarshalJSON() ([]byte, error) {
	indexRightNovelDTO := IndexRightNovelDTO{
		Id:         p.Id,
		NovelTitle: p.NovelTitle,
	}
	if p.NovelCategory != nil {
		indexRightNovelDTO.CategoryName = p.NovelCategory.CategoryName
	}
	return json.Marshal(indexRightNovelDTO)
}
