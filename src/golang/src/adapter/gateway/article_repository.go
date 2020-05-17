package gateway

import (
	"github.com/jinzhu/gorm"

	"github.com/blacknikka/go-blog/domain/article"
)

type (
	ArticleRepository struct {
		Conn *gorm.DB
	}

	Article struct {
		gorm.Model
		Title string `gorm:"size:20;not null"`
		Body  string `gorm:"size:100;not null"`
	}
)

func (r *ArticleRepository) Store(a article.Article) (*article.ArticleID, error) {
	articleData := &Article{
		Title: a.Title,
		Body:  a.Body,
	}

	if err := r.Conn.Create(articleData).Error; err != nil {
		return nil, err
	}

	return &article.ArticleID{ID: articleData.ID}, nil
}

func (r *ArticleRepository) FindByID(id article.ArticleID) (article.Articles, error) {
	articles := []article.Article{}
	return articles, nil
}

func (r *ArticleRepository) FindAll() (article.Articles, error) {
	articles := []article.Article{}
	return articles, nil
}
