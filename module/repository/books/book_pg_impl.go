package book

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"remakech7/module/model"
)

type BookPgRepoImpl struct {
	db *sql.DB
}

func NewBookPgRepo(db *sql.DB) BookRepo {
	return &BookPgRepoImpl{
		db: db,
	}
}

func (b *BookPgRepoImpl) FindBookById(ctx context.Context, bookId uint64) (book model.Book, err error) {
	sqlStatement := `
	SELECT id, title, author, description FROM books
	WHERE id = $1
	`

	stmt, err := b.db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.QueryContext(ctx, bookId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Desc); err != nil {
			return
		}
	}

	if book.BookID <= 0 {
		err = errors.New("book not found")
	}
	
	fmt.Printf("book data : %+v\n", book)
	return 
}
func (b *BookPgRepoImpl) FindAllBooks(ctx context.Context) (books []model.Book, err error) {
	sqlStatement := `
	SELECT id, title, author, description FROM books
	`

	stmt, err := b.db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		if err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Desc); err != nil {
			return
		}
		books = append(books, book)
	}

	return
}
func (b *BookPgRepoImpl) InsertBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	RETURNING id, title, author, description
	`

	stmt, err := b.db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.QueryContext(ctx,
		bookIn.Title,
		bookIn.Author,
		bookIn.Desc)
	if err != nil {
		return
	}


	for rows.Next() {
		if err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Desc); err != nil {
			return
		}
	}

	return
}
func (b *BookPgRepoImpl) UpdateBook(ctx context.Context, bookIn model.Book) (err error) {
	sqlStatement := `
	UPDATE books 
	SET title = $2, author = $3, description = $4
	WHERE id = $1
	RETURNING id, title, author, description
	`

	stmt, err := b.db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		bookIn.BookID,
		bookIn.Title,
		bookIn.Author,
		bookIn.Desc)
	if err != nil {
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affected <= 0 {
		err = errors.New("user is not found")
		return
	}
	return
}

func (b *BookPgRepoImpl) DeleteBookById(ctx context.Context, bookId uint64) (book model.Book, err error) {
	sqlStatement := `
	DELETE from books
	where id = $1
	`

	stmt, err := b.db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		return
	}

	rows, err := stmt.QueryContext(ctx,
		bookId)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Desc); err != nil {
			return
		}
	}
	// if book.BookID <= 0 {
	// 	err = errors.New("user is not found")
	// }
	return
}