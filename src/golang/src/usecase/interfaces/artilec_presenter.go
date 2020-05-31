package interfaces

import (
	"github.com/blacknikka/go-blog/domain/article"
	"github.com/gin-gonic/gin"
)

type ArticlePresenter interface {
	Complete(article.Article, *gin.Context) error
}
