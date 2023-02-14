package service

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

type CountArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy string `form:"create_by" binding:"max=100"`
}

type GetArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type CreateArticleRequest struct {
	Title string `form:"title" binding:"required,min=3,max=100"`
	CreatedBy string `form:"create_by" binding:"required,min=3,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
	Desc string `form:"desc" binding:"max=1000"`
	Content string `form:"content" binding:"required"`
	CoverImageUrl string `form:"coverimageurl" binding:"max=100"`
}

type UpdateArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	Title string `form:"title" binding:"min=0,max=100"`
	State uint8 `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
	Desc string `form:"desc" binding:"max=1000"`
	Content string `form:"content" binding:"required"`
	CoverImageUrl string `form:"cover_image_url" binding:"max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type AddTagRequest struct {
	TagID uint32 `form:"tag_id" binding:"required,gte=1"`
	ArticleID uint32 `form:"article_id" binding:"required,gte=1"`
}

type DeleteTagFromArticleRequest struct {
	TagID uint32 `form:"tag_id" binding:"required,gte=1"`
	ArticleID uint32 `form:"article_id" binding:"required,gte=1"`
}

func (svc *Service) CountArticle(param *CountArticleRequest) (int64, error) {
	return svc.dao.CountArticle(param.Title, param.ModifiedBy, param.State)
}

func (svc *Service) ListArticle(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.ListArticle(param.Title, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) GetArticle(param *GetArticleRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID)
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.State, param.Desc, param.Content, param.CoverImageUrl, param.CreatedBy)
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.State, param.Desc, param.Content, param.CoverImageUrl, param.ModifiedBy)
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}

func (svc *Service) AddTag(param *AddTagRequest)  {
	return
}

