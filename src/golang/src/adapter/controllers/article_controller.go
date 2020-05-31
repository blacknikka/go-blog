package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/blacknikka/go-blog/domain/article"
	"github.com/blacknikka/go-blog/usecase"
)

type ArticleController struct {
	Interactor usecase.ArticleInteractor
}

func NewArticleController(interactor usecase.ArticleInteractor) *ArticleController {
	return &ArticleController{
		Interactor: interactor,
	}
}

func (controller *ArticleController) Create(c *gin.Context) {
	type (
		Request struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}
		Response struct {
			ArticleID uint `json:"article_id"`
		}
	)
	req := Request{}
	c.BindJSON(&req)
	articleData := article.Article{Title: req.Title, Body: req.Body}

	_, err := controller.Interactor.Add(articleData, c)
	if err != nil {
		return
	}
}
