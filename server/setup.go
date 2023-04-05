package server

import (
	"remakech7/config"
	bookhdl "remakech7/module/handler/book"
	bookrepo "remakech7/module/repository/books"
	booksvc "remakech7/module/service/book"
)

type handlers struct {
	bookHdl bookhdl.BookHandler
}

func initDI() handlers {
	pgconn := config.NewPostgresConfig()
	bookRepo := bookrepo.NewBookPgRepo(pgconn)
	bookSvc := booksvc.NewBookSvc(bookRepo)
	bookHdl := bookhdl.NewBookHandler(bookSvc)

	return handlers{
		bookHdl: bookHdl,
	}
}