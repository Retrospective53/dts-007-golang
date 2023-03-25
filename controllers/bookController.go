package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Desc string `json:"desc"`
}

var bookDatas = []Book{{
	BookID: "b1",
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

	newBook.BookID = fmt.Sprintf("b%d", len(bookDatas)+1)
	bookDatas = append(bookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	BookID:= ctx.Param("BookID")
	condition := false
	var updatedBook Book
	updatedBook.BookID = BookID

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range bookDatas {
		if BookID == book.BookID {
			condition = true
			bookDatas[i] = updatedBook
			bookDatas[i].BookID = BookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", BookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been succesfully updated", BookID),
		"updatedBook": updatedBook,
	})
}

func GetBook (ctx *gin.Context) {
	BookID := ctx.Param("BookID")
	condition := false
	var bookData Book

	for i, book := range bookDatas {
		if BookID == book.BookID {
			condition = true
			bookData = bookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not Found",
			"error_message": fmt.Sprintf("book with id %v not found", BookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func GetAllBook (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"books": bookDatas,
	})
}

func DeleteBook (ctx *gin.Context) {
	BookID := ctx.Param("BookID")
	condition := false
	var bookIndex int

	for i, book := range bookDatas {
		if BookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status": "Data not Found",
			"error_message": fmt.Sprintf("book with id %v not found", BookID),
		})
		return
	}

	copy(bookDatas[bookIndex:], bookDatas[bookIndex+1:])
	bookDatas[len(bookDatas)-1] = Book{}
	bookDatas = bookDatas[:len(bookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been succesfuly deleted", BookID),
	})
}