package book

import (
	"remakech7/module/handler/book"

	"github.com/gin-gonic/gin"
)

func NewBookRouter(gr *gin.RouterGroup, bookHdl book.BookHandler) {
	g := gr.Group("book")

	g.GET("/all", bookHdl.FindAllBookssHdl)
	g.GET("/:id", bookHdl.FindBookByIdHdl)
	g.POST("", bookHdl.InsertBookHdl)
	g.PUT("/:id", bookHdl.UpdateBookHdl)
	g.DELETE("/:id", bookHdl.DeleteBookByIdHdl)
}