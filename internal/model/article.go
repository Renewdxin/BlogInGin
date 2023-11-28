package model

import "BloginGin/pkg/app"

// define Article
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type CountArticleRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	State     uint8  `form:"state, default=1" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state, default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (A Article) TableName() string {
	return "blog_article_tag"
}
