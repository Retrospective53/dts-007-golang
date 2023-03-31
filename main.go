package main

import (
	"challenge-7/config"
	"challenge-7/module/repository/book"
	"fmt"
)



func main() {
	config.NewPostgresConfig()
	fmt.Println("succesfully connectted to the database")
	book.DeleteBook()
}