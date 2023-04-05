package book

import (
	"context"
	"log"
	"remakech7/module/model"
	book "remakech7/module/repository/books"
)

type BookSvcImpl struct {
	bookRepo book.BookRepo
}

func NewBookSvc(bookRepo book.BookRepo) BookService {
	return &BookSvcImpl{
		bookRepo: bookRepo,
	}
}

func (b *BookSvcImpl) FindBookByIdSvc(ctx context.Context, bookId uint64) (book model.Book, err error) {
	log.Printf("[INFO] %T FindUserById invoked\n", b)
	if book, err = b.bookRepo.FindBookById(ctx, bookId); err != nil {
		log.Printf("[ERROR] error FindBookById :%v\n", err)
	}
	return
}

func (b *BookSvcImpl) FindAllBooksSvc(ctx context.Context) (books []model.Book, err error) {
	log.Printf("[INFO] %T FindAllBooks invoked\n", b)
	if books, err = b.bookRepo.FindAllBooks(ctx); err != nil {
		log.Printf("[ERROR] error FindAllBooks :%v\n", err)
	}
	return
}

func (b *BookSvcImpl) InsertBookSvc(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	log.Printf("[INFO] %T InsertBooks invoked\n", b)
	if book, err = b.bookRepo.InsertBook(ctx, bookIn); err != nil {
		log.Printf("[ERROR] error InsertBook :%v\n", err)
	}
	return
}

func (b *BookSvcImpl) UpdateBookSvc(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	log.Printf("[INFO] %T UpdateBooks invoked\n", b)
	book, err = b.bookRepo.UpdateBook(ctx, bookIn); 
	if err != nil {
		log.Printf("[Error] error UpdateBook : %v\n", err)
	}
	return
}

func (b *BookSvcImpl) DeleteBookByIdSvc(ctx context.Context, bookId uint64) (book model.Book, err error) {
	log.Printf("[INFO] %T DeleteBookById invoked\n", b)
	if book, err = b.bookRepo.DeleteBookById(ctx, bookId); err != nil {
		log.Printf("[ERROR] error DeleteBookById :%v\n", err)
	}
	return
}