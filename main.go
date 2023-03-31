package main

import (
	"challenge-7/config"
	"challenge-7/module/repository/book"
	"challenge-7/routers"
	"fmt"
)



func main() {
	var PORT = ":4000"
	config.NewPostgresConfig()
	fmt.Println("succesfully connectted to the database")
	book.CreateBook()

	routers.StartServer().Run(PORT)
	
}