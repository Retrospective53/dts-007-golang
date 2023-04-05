package book

import (
	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	FindBookByIdHdl(ctx *gin.Context)
	FindAllBookssHdl(ctx *gin.Context)
	InsertBookHdl(ctx *gin.Context)
	UpdateBookHdl(ctx *gin.Context)
	DeleteBookByIdHdl(ctx *gin.Context)
}