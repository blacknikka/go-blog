package infra

import (
	"github.com/blacknikka/go-blog/adapter/controllers"
	"github.com/blacknikka/go-blog/infra/db/postgresql"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	conn := postgresql.Connect()
	articleController := controllers.NewArticleController(conn)

	router.POST("/article", func(c *gin.Context) { articleController.Create(c) })

	Router = router
}
