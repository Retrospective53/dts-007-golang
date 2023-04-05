package book

import (
	"net/http"
	"remakech7/module/model"
	"remakech7/module/service/book"
	"remakech7/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHdlImpl struct {
	bookSvc book.BookService
}

func NewBookHandler(bookSvc book.BookService) BookHandler {
	return &BookHdlImpl{
		bookSvc: bookSvc,
	}
}

func (b *BookHdlImpl) FindBookByIdHdl(ctx *gin.Context) {
	// id := ctx.Query("id")
	// if id == "" {
	// 	ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
	// 		Message:  "failed to find book",
	// 		ErrorMsg: response.InvalidQuery,
	// 	})
	// 	return
	// }

	id := ctx.Param("id")

	// transform id string to uint64
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to find book",
			ErrorMsg: response.InvalidParam,
		})
		return
	}

	// call service
	book, err := b.bookSvc.FindBookByIdSvc(ctx, idUint)
	if err != nil {
		if err.Error() == "book is not found" {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse{
				Message:  "failed to find book",
				ErrorMsg: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to find book",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success find book",
		Data:    book,
	})
}

func (b *BookHdlImpl) FindAllBookssHdl(ctx *gin.Context)  {
	books, err := b.bookSvc.FindAllBooksSvc(ctx)
	if err != nil {
		// bad code, should be wrapped in other package
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to get books",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "success get books",
		Data:    books,
	})
}

func (b *BookHdlImpl) InsertBookHdl(ctx *gin.Context) {
	var bookIn model.Book
	
	if err := ctx.Bind(&bookIn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to insert book",
			ErrorMsg: response.InvalidBody,
		})
		return
	}

	//validate
	if bookIn.Title == "" || bookIn.Author == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to insert book",
			ErrorMsg: response.InvalidParam,
		})
		return
	}

	insertedBook, err := b.bookSvc.InsertBookSvc(ctx, bookIn)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to insert book",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success create book",
		Data:    insertedBook,
	})
	
}

func (b *BookHdlImpl) UpdateBookHdl(ctx *gin.Context) {
	idUint, err := b.getIdFromParam(ctx)
	if err != nil {
		return
	}

	// binding payload
	var bookIn model.Book
	if err := ctx.Bind(&bookIn); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to update book",
			ErrorMsg: response.InvalidBody,
		})
		return
	}
	
	bookIn.BookID = idUint

	// validate name
	if bookIn.Title == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to update book",
			ErrorMsg: response.InvalidParam,
		})
		return
	}

	book, err := b.bookSvc.UpdateBookSvc(ctx, bookIn); 
	if err != nil {
		if err.Error() == "book is not found" {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse{
				Message:  "failed to update book",
				ErrorMsg: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to update book",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success update book",
		Data: book,
	})
}

func (b *BookHdlImpl) DeleteBookByIdHdl(ctx *gin.Context) {
	idUint, err := b.getIdFromParam(ctx)
	if err != nil {
		return
	}
	_, err = b.bookSvc.DeleteBookByIdSvc(ctx, idUint)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message:  "failed to delete book",
			ErrorMsg: response.SomethingWentWrong,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, response.SuccessResponse{
		Message: "success delete book",
		Data:    "none",
	})
}

func (b *BookHdlImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
	id := ctx.Param("id")

	// transform id string to uint64
	idUint, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message:  "failed to find book",
			ErrorMsg: response.InvalidParam,
		})
		return
	}

	return
}