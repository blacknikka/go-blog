package infra

import (
	"github.com/blacknikka/go-blog/adapter/controllers"
	"github.com/blacknikka/go-blog/adapter/gateway"
	"github.com/blacknikka/go-blog/infra/db/postgresql"
	"github.com/blacknikka/go-blog/usecase"
	"github.com/blacknikka/go-blog/usecase/presenter"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	conn := postgresql.Connect()

	articleInteractor := usecase.ArticleInteractor{
		ArticleRepository: &gateway.ArticleRepository{
			Conn: conn,
		},
		ArticlePresenter: presenter.NewArticleCretePresenterWithJSON(),
	}
	articleController := controllers.NewArticleController(articleInteractor)

	router.POST("/article", func(c *gin.Context) { articleController.Create(c) })

	Router = router
}
