package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

func (d *Dao) CountArticle(title string, modifiedBy string,state uint8) (int64, error) {
	article := model.Article{
		Title: title,
		State: state,
		Model: &model.Model{ModifiedBy: modifiedBy},
	}
	return article.Count(d.engine)
}

func (d *Dao) ListArticle(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{
		Title: title,
		State: state,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetArticle(id uint32) (*model.Article, error) {
	article := model.Article{
		Model: &model.Model{ID: id},
	}
	return article.Get(d.engine)
}

func (d *Dao) CreateArticle(title string, state uint8, desc string, content string, coverImageUrl string, createdBy string) error {
	article := model.Article{
		Title: title,
		State: state,
		Desc: desc,
		Content: content,
		CoverImageUrl: coverImageUrl,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title string, state uint8, desc string, content string, coverImageUrl string, modifiedBy string) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}


	values := map[string]interface{} {
		"state": state,
		"modified_by": modifiedBy,
		"content": content,
	}

	if title != ""{
		values["title"] = title
	}

	if desc != ""{
		values["desc"] = desc
	}

	if coverImageUrl != ""{
		values["cover_image_url"] = coverImageUrl
	}

	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}

