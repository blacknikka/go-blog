package interfaces

import "github.com/blacknikka/go-blog/domain/article"

type ArticleRepository interface {
	Store(article.Article) (*article.ArticleID, error)
	FindByID(article.ArticleID) (article.Articles, error)
	FindAll() (article.Articles, error)
}
