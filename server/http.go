package server

import (
	"remakech7/config"
	"remakech7/module/router/book"

	"github.com/gin-gonic/gin"
)

func NewHttpServer() {
	hdls := initDI()
	ginServer := gin.Default()

	ginServer.Use(
		// gin.Logger(),
		gin.Recovery(),
	)

	api := ginServer.Group("/api")
	book.NewBookRouter(api, hdls.bookHdl)
	ginServer.Run(config.PORT)
}