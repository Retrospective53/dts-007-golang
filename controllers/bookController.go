package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"challenge-7/module/repository/book"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Desc string `json:"desc"`
}

var bookDatas = []Book{{
	BookID: 1,
  Title: "The Call of Cthulhu",
  Author: "H.P Lovecraft",
  Desc: "lolidk",
},
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	bookSent := book.CreateBook(newBook.Title, newBook.Author, newBook.Desc)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": bookSent,
	})
}

func UpdateBook(ctx *gin.Context) {
	BookIDStr := ctx.Param("BookID")
	BookID, err := strconv.Atoi(BookIDStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var updatedBook Book
	updatedBook.BookID = BookID

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book.UpdateBook(BookID , updatedBook.Title, updatedBook.Author, updatedBook.Desc)


	// for i, book := range bookDatas {
	// 	if BookID == book.BookID {
	// 		condition = true
	// 		bookDatas[i] = updatedBook
	// 		bookDatas[i].BookID = BookID
	// 		break
	// 	}
	// }


	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been succesfully updated", BookID),
		"updatedBook": updatedBook,
	})
}

func GetBook (ctx *gin.Context) {
	BookIDStr := ctx.Param("BookID")
	BookID, err := strconv.Atoi(BookIDStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	bookData := book.GetBook(BookID)



	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func GetAllBook (ctx *gin.Context) {
	booksDatas := book.GetBooks()
	ctx.JSON(http.StatusOK, gin.H{
		"books": booksDatas,
	})
}

func DeleteBook (ctx *gin.Context) {
	BookIDStr := ctx.Param("BookID")
	BookID, err := strconv.Atoi(BookIDStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	book.DeleteBook(BookID)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been succesfuly deleted", BookID),
	})
}