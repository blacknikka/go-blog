package usecase

import (
	"github.com/blacknikka/go-blog/domain/article"
	"github.com/blacknikka/go-blog/usecase/interfaces"
)

type ArticleInteractor struct {
	ArticleRepository interfaces.ArticleRepository
}

func (interactor *ArticleInteractor) Add(a article.Article) (*article.ArticleID, error) {
	return interactor.ArticleRepository.Store(a)
}
