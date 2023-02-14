package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"gorm.io/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"Content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	List []*Article
	Pager *app.Pager
}

func (a Article) TableName() string{
	return "blog_article"
}

func (a Article) Count(db *gorm.DB) (int64, error) {
	var count int64
	if a.Title != "" {
		db.Where("title = ?", a.Title)
	}
	db.Where("state = ?", a.State)
	if a.ModifiedBy != "" {
		db.Where("modified_by = ?", a.ModifiedBy)
	}
	if err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0,err
	}
	return count, nil
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articleList []*Article

	if pageSize > 0 && pageOffset >= 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	if err := db.Model(&a).Where("is_del = ?", 0).Find(&articleList).Error; err != nil {
		return nil, err
	}
	return articleList, nil
}

func (a  Article) Get(db *gorm.DB) (*Article, error) {
	if err := db.Model(&a).Where("id = ? and is_del = ?", a.ID, 0).Find(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error
}

