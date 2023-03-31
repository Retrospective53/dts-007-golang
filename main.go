package main

import (
	"challenge-7/config"
	"challenge-7/routers"
	"fmt"
)



func main() {
	db := config.NewPostgresConfig()
	defer db.Close()
	fmt.Println("succesfully connectted to the database")
	var PORT = ":4000"


	routers.StartServer().Run(PORT)
	
}