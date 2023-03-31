package book

import (
	"challenge-7/config"
	"challenge-7/module/model"
	"database/sql"
	"fmt"
)

var (
	db *sql.DB
	err error
)

func init() {
	db = config.NewPostgresConfig()
}

func CreateBook() {
	var book = model.Book{}

	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	RETURNING id, title, author, description
	`

	err = db.QueryRow(sqlStatement, "Call of Cthulhu", "H.P Lovecraft", "Horror").
	Scan(&book.ID, &book.Title, &book.Author, &book.Description)

	if err != nil {
		panic(err)
	}

	fmt.Printf("new book data : %+v\n", book)
}

func GetBook() {
	var book = model.Book{}

	sqlStatement := `
	SELECT id, title, author, description FROM books
	WHERE id = $1
	`

	row := db.QueryRow(sqlStatement, 6)

	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

	if err != nil {
		panic(err)
	}

	fmt.Printf("book data : %+v\n", book)
}

func GetBooks() {
	var results = []model.Book{}

	sqlStatement := `
	SELECT id, title, author, description FROM books
	`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = model.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	fmt.Println("books data :", results)
}

func UpdateBook() {
	sqlStatement := `
	UPDATE books 
	SET title = $2, author = $3, description = $4
	WHERE id = $1
	RETURNING id, title, author, description
	`

	res, err := db.Exec(sqlStatement, 6, "Banki Banki", "stacks", "rocknrocnnbeat")
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("updated data:", count)
}

func DeleteBook() {
	sqlStatement := `
	DELETE from books
	where id = $1
	`

	res, err := db.Exec(sqlStatement, 7)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("deleted data amount:", count)
}