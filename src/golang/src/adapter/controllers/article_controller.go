package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/blacknikka/go-blog/adapter/gateway"
	"github.com/blacknikka/go-blog/domain/article"
	"github.com/blacknikka/go-blog/usecase"
)

type ArticleController struct {
	Interactor usecase.ArticleInteractor
}

func NewArticleController(conn *gorm.DB) *ArticleController {
	return &ArticleController{
		Interactor: usecase.ArticleInteractor{
			ArticleRepository: &gateway.ArticleRepository{
				Conn: conn,
			},
		},
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

	id, err := controller.Interactor.Add(articleData)
	if err != nil {
		return
	}
	_ = Response{ArticleID: id.ID}
}
