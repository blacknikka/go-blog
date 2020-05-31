package usecase

import (
	"github.com/blacknikka/go-blog/domain/article"
	"github.com/blacknikka/go-blog/usecase/interfaces"
	"github.com/gin-gonic/gin"
)

type ArticleInteractor struct {
	ArticleRepository interfaces.ArticleRepository
	ArticlePresenter  interfaces.ArticlePresenter
}

func (interactor *ArticleInteractor) Add(a article.Article, c *gin.Context) (*article.ArticleID, error) {
	id, err := interactor.ArticleRepository.Store(a)

	a.ID = *id
	interactor.ArticlePresenter.Complete(a, c)
	return id, err
}
