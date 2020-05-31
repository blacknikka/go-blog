package presenter

import (
	"net/http"

	"github.com/blacknikka/go-blog/domain/article"
	"github.com/gin-gonic/gin"
)

func NewArticleCretePresenterWithJSON() *articleCretePresenterWithJSON {
	return &articleCretePresenterWithJSON{}
}

type articleCretePresenterWithJSON struct{}

type Response struct {
	ArticleID uint `json:"article_id"`
}

func (p *articleCretePresenterWithJSON) Complete(a article.Article, c *gin.Context) error {
	c.JSON(http.StatusOK, &Response{ArticleID: a.ID.ID})
	return nil
}
