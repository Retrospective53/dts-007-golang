package book

import (
	"context"
	"remakech7/module/model"
)

type BookRepo interface {
	FindBookById(ctx context.Context, bookId uint64) (book model.Book, err error)
	FindAllBooks(ctx context.Context) (books []model.Book, err error)
	InsertBook(ctx context.Context, bookIn model.Book) (book model.Book, err error)
	UpdateBook(ctx context.Context, bookIn model.Book) (err error)
	DeleteBookById(ctx context.Context, bookId uint64) (book model.Book, err error)
}