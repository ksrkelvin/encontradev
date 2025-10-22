package internal

import (
	"encontradev/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./internal/static")

	r.LoadHTMLGlob("internal/templates/*")

	r.GET("/", handlers.Home)
	r.GET("/partial", handlers.PartialExample)

	return r
}
