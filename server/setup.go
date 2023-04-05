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
	// no gorm
	// pgconn := config.NewPostgresConfig()
	// bookRepo := bookrepo.NewBookPgRepo(pgconn)

	// gorm
	pgconn := config.NewPostgresGormConn()
	bookRepo := bookrepo.NewBookGormRepo(pgconn)
	
	bookSvc := booksvc.NewBookSvc(bookRepo)
	bookHdl := bookhdl.NewBookHandler(bookSvc)

	return handlers{
		bookHdl: bookHdl,
	}
}